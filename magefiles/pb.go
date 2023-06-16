package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gosimple/slug"
	_ "github.com/mattn/go-sqlite3"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/magefile/mage/mg"
)

type Pb mg.Namespace

type Symptom struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Common struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Image struct {
	Id          string `json:"id"`
	SourceUrl   string `json:"source_url"`
	Attribution string `json:"attribution"`
	License     string `json:"license"`
}

type Plant struct {
	Pid          string   `json:"pid"`
	Name         string   `json:"name"`
	Symptoms     []string `json:"symptoms"`
	Images       []string `json:"images"`
	Common       []string `json:"common"`
	Animals      []string `json:"animals"`
	WikipediaUrl string   `json:"wikipedia_url"`
	Kingdom      string   `json:"kingdom"`
	Clades       []string `json:"clades"`
	Family       string   `json:"family"`
	Genus        string   `json:"genus"`
	Order        string   `json:"order"`
	Species      string   `json:"species"`
}

type MarshalSymptom struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type MarshalCommon struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type MarshalImage struct {
	SourceUrl    string `json:"source_url"`
	Attribution  string `json:"attribution"`
	License      string `json:"license"`
	RelativePath string `json:"relative_path"`
}

type MarshalPlant struct {
	Pid             string                `json:"pid"`
	Name            string                `json:"name"`
	Animals         []interface{}         `json:"animals"`
	Common          []MarshalCommon       `json:"common"`
	Symptoms        []MarshalSymptom      `json:"symptoms"`
	Images          []MarshalImage        `json:"images"`
	WikipediaUrl    string                `json:"wikipedia_url"`
	DateLastUpdated string                `json:"date_last_updated"`
}

func (Pb) Export(ctx context.Context) error {
	response, err := http.Get("http://0.0.0.0:8090/api/collections/plants/records?expand=images,symptoms,common&sort=pid&perPage=1000")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var recordJson interface{}
	err = json.Unmarshal(body, &recordJson)
	if err != nil {
		fmt.Println(err)
	}

	var (
		exported []MarshalPlant
		common   []MarshalCommon
		symptoms []MarshalSymptom
		images   []MarshalImage
	)

	returnById := func(id string, haystack []interface{}) (interface{}, bool) {
		for _, record := range haystack {
			if id == record.(map[string]interface{})["id"].(string) {
				return record, true
			}
		}
		return nil, false
	}

	for _, plant := range recordJson.(map[string]interface{})["items"].([]interface{}) {
		for _, cId := range plant.(map[string]interface{})["common"].([]interface{}) {
			if found, ok := returnById(string(cId.(string)), plant.(map[string]interface{})["expand"].(map[string]interface{})["common"].([]interface{})); ok {
				common = append(common, MarshalCommon{
					Slug: found.(map[string]interface{})["slug"].(string),
					Name: found.(map[string]interface{})["name"].(string),
				})
			}
		}

		for _, sId := range plant.(map[string]interface{})["symptoms"].([]interface{}) {
			if found, ok := returnById(string(sId.(string)), plant.(map[string]interface{})["expand"].(map[string]interface{})["symptoms"].([]interface{})); ok {
				symptoms = append(symptoms, MarshalSymptom{
					Slug: found.(map[string]interface{})["slug"].(string),
					Name: found.(map[string]interface{})["name"].(string),
				})
			}
		}

		for _, iId := range plant.(map[string]interface{})["images"].([]interface{}) {
			if found, ok := returnById(iId.(string), plant.(map[string]interface{})["expand"].(map[string]interface{})["images"].([]interface{})); ok {
				imageNameHash := md5.Sum([]byte(found.(map[string]interface{})["source_url"].(string)))
				images = append(images, MarshalImage{
					Attribution:  found.(map[string]interface{})["attribution"].(string),
					SourceUrl:    found.(map[string]interface{})["source_url"].(string),
					License:      found.(map[string]interface{})["license"].(string),
					RelativePath: fmt.Sprintf("plants/%s/%s/large.jpg", plant.(map[string]interface{})["pid"].(string), hex.EncodeToString(imageNameHash[:])),
				})
			}
		}

		exported = append(exported, MarshalPlant{
			Pid:             plant.(map[string]interface{})["pid"].(string),
			Name:            plant.(map[string]interface{})["name"].(string),
			Animals:         plant.(map[string]interface{})["animals"].([]interface{}),
			WikipediaUrl:    plant.(map[string]interface{})["wikipedia_url"].(string),
			DateLastUpdated: plant.(map[string]interface{})["updated"].(string),
			Common:          common,
			Symptoms:        symptoms,
			Images:          images,
		})

		common = nil
		symptoms = nil
		images = nil
	}

	if err := json.NewEncoder(os.Stdout).Encode(exported); err != nil {
		panic(err)
	}

	return nil
}

func (Pb) Import(ctx context.Context, sourcePath, databasePath string) error {
	fmt.Printf("Unmarshaling: %s\n", sourcePath)
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	fmt.Printf("Connecting to database: %s\n", databasePath)
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		images   collection.Collection[*Image]
		common   collection.Collection[*Common]
		symptoms collection.Collection[*Symptom]
	)

	fmt.Println("Collecting images, common names and symptoms ...")
	for _, p := range plants.Data {
		for _, i := range p.Images {
			images.PushDistinct(&Image{
				SourceUrl:   i.SourceUrl,
				Attribution: i.Attribution,
				License:     i.License,
			})
		}

		for _, c := range p.Common {
			common.PushDistinct(&Common{
				Name: cases.Title(language.AmericanEnglish, cases.NoLower).String(c.Name),
				Slug: slug.Make(c.Name),
			})
		}

		for _, s := range p.Symptoms {
			symptoms.PushDistinct(&Symptom{
				Name: cases.Title(language.AmericanEnglish, cases.NoLower).String(s.Name),
				Slug: slug.Make(s.Name),
			})
		}
	}

	fmt.Printf("Upserting %d Common Names ...\n", common.Length())
	bar := progressbar.Default(int64(common.Length()))
	common.Each(func(i int, c *Common) bool {
		defer func() {
			bar.Add(1)
		}()

		body, _ := json.Marshal(c)

		response, err := http.Post("http://0.0.0.0:8090/api/collections/common/records", "application/json", bytes.NewBuffer(body))

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var recordJson interface{}
		err = json.Unmarshal(body, &recordJson)
		if err != nil {
			fmt.Println(err)
		}

		var id string

		if recordJson.(map[string]interface{})["code"] != nil && recordJson.(map[string]interface{})["data"].(map[string]interface{})["slug"].(map[string]interface{})["code"] == "validation_not_unique" {
			statement, _ := db.Prepare("SELECT id FROM common WHERE slug = ?")
			defer statement.Close()
			statement.QueryRow(c.Slug).Scan(&id)
		} else {
			id = recordJson.(map[string]interface{})["id"].(string)
		}

		c.Id = id

		return false
	})

	fmt.Printf("Upserting %d Symptoms ...\n", symptoms.Length())
	bar = progressbar.Default(int64(symptoms.Length()))
	symptoms.Each(func(i int, s *Symptom) bool {
		defer func() {
			bar.Add(1)
		}()
		body, _ := json.Marshal(s)

		response, err := http.Post("http://0.0.0.0:8090/api/collections/symptoms/records", "application/json", bytes.NewBuffer(body))

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var recordJson interface{}
		err = json.Unmarshal(body, &recordJson)
		if err != nil {
			fmt.Println(err)
		}

		var id string
		if recordJson.(map[string]interface{})["code"] != nil && recordJson.(map[string]interface{})["data"].(map[string]interface{})["name"].(map[string]interface{})["code"] == "validation_not_unique" {
			statement, _ := db.Prepare("SELECT id FROM symptoms WHERE slug = ?")
			defer statement.Close()
			statement.QueryRow(s.Slug).Scan(&id)
		} else {
			id = recordJson.(map[string]interface{})["id"].(string)
		}

		s.Id = id

		return false
	})

	fmt.Printf("Upserting %d images ...\n", images.Length())
	bar = progressbar.Default(int64(images.Length()))
	images.Each(func(i int, image *Image) bool {
		defer func() {
			bar.Add(1)
		}()
		body, _ := json.Marshal(image)

		response, err := http.Post("http://0.0.0.0:8090/api/collections/images/records", "application/json", bytes.NewBuffer(body))

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var recordJson interface{}
		err = json.Unmarshal(body, &recordJson)
		if err != nil {
			fmt.Println(err)
		}

		var id string
		if recordJson.(map[string]interface{})["code"] != nil && recordJson.(map[string]interface{})["data"].(map[string]interface{})["source_url"].(map[string]interface{})["code"] == "validation_not_unique" {
			statement, _ := db.Prepare("SELECT id FROM images WHERE source_url = ?")
			defer statement.Close()
			statement.QueryRow(image.SourceUrl).Scan(&id)
		} else {
			id = recordJson.(map[string]interface{})["id"].(string)
		}

		image.Id = id

		return false
	})

	var symptomIds []string
	var imageIds []string
	var commonIds []string

	fmt.Printf("Upserting %d plants ...\n", len(plants.Data))
	bar = progressbar.Default(int64(len(plants.Data)))
	for _, plant := range plants.Data {
		for _, s := range plant.Symptoms {
			found := symptoms.Find(func(i int, item *Symptom) bool {
				return item.Slug == slug.Make(s.Name)
			})

			if found != nil {
				symptomIds = append(symptomIds, found.Id)
			}
		}

		for _, c := range plant.Common {
			found := common.Find(func(i int, item *Common) bool {
				return item.Slug == slug.Make(c.Name)
			})

			if found != nil {
				commonIds = append(commonIds, found.Id)
			}
		}

		for _, i := range plant.Images {
			found := images.Find(func(_ int, item *Image) bool {
				return item.SourceUrl == i.SourceUrl
			})

			if found != nil {
				imageIds = append(imageIds, found.Id)
			}
		}

		body, _ := json.Marshal(Plant{
			Pid:          plant.Pid,
			Name:         plant.Name,
			Animals:      plant.Animals,
			WikipediaUrl: plant.WikipediaUrl,
			Symptoms:     symptomIds,
			Common:       commonIds,
			Images:       imageIds,
		})

		symptomIds = nil
		commonIds = nil
		imageIds = nil

		response, err := http.Post("http://0.0.0.0:8090/api/collections/plants/records", "application/json", bytes.NewBuffer(body))

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var recordJson interface{}
		err = json.Unmarshal(body, &recordJson)
		if err != nil {
			fmt.Println(err)
		}

		bar.Add(1)
	}

	return nil
}

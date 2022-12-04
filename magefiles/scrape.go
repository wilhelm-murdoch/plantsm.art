package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/plantsm.art/models"
)

const baseUrl = "https://www.aspca.org/pet-care/animal-poison-control/%s-plant-list"

type Scrape mg.Namespace

func (Scrape) Bootstrap(ctx context.Context, path string) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("delete the scaffold file %s and try again", path)
	}

	data := &models.Data{
		PlantsByCat: collection.New[string](),
		PlantsByDog: collection.New[string](),
		Plants:      collection.New[*models.Plant](),
	}

	if err := importPlantsForAnimal("cats", data); err != nil {
		return err
	}

	if err := importPlantsForAnimal("dogs", data); err != nil {
		return err
	}

	data.Plants.Each(func(i int, p *models.Plant) bool {
		document, err := getDocumentFromUrl(p.SourceUrl)
		if err != nil {
			return true
		}

		symptoms := document.Find("div.field-name-field-clinical-signs span.values").Text()

		p.Symptoms = strings.Split(symptoms, ",")

		return false
	})

	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func importPlantsForAnimal(animal string, data *models.Data) error {
	if animal != "dogs" && animal != "cats" {
		return fmt.Errorf("expected \"cats\" or \"dogs\", but got \"%s\" instead", animal)
	}

	document, err := getDocumentFromUrl(fmt.Sprintf(baseUrl, animal))
	if err != nil {
		return err
	}

	document.Find("div.view-content").First().Find("div.views-row span.field-content").Each(func(i int, s *goquery.Selection) {
		plant := models.NewPlant(animal, s)

		switch {
		case animal == "cats":
			data.PlantsByCat.PushDistinct(plant.Id)
		case animal == "dogs":
			data.PlantsByDog.PushDistinct(plant.Id)
		}

		data.Plants.PushDistinct(plant)
	})

	return nil
}

type UnmarshalData struct {
	PlantsByCat      []string   `json:"plants_by_cat"`
	PlantsByDog      []string   `json:"plants_by_dog"`
	PlantsBySymptoms [][]string `json:"plants_by_symptoms"`
	Plants           []struct {
		Id             string   `json:"id"`
		Name           string   `json:"name"`
		ScientificName string   `json:"scientific_name"`
		Common         []string `json:"common"`
		Images         []struct {
			SourceUrl   string `json:"source_url"`
			Url         string `json:"url"`
			License     string `json:"license"`
			Attribution string `json:"attribution"`
		} `json:"images"`
		Symptoms       []string `json:"symptoms"`
		SourceUrl      string   `json:"source_url"`
		Classification struct {
			Kingdom string   `json:"kingdom"`
			Clades  []string `json:"clades"`
			Order   string   `json:"order"`
			Family  string   `json:"family"`
			Genus   string   `json:"genus"`
			Species string   `json:"species"`
		} `json:"classification"`
	} `json:"plants"`
}

type Photo struct {
	License     string `json:"license_code"`
	Attribution string `json:"attribution"`
	Url         string `json:"url"`
	UrlMedium   string `json:"medium_url"`
	UrlLarge    string `json:"large_url"`
	UrlOriginal string `json:"original_url"`
}

type UnmarshalRecord struct {
	Id           int   `json:"id"`
	DefaultPhoto Photo `json:"default_photo"`
	Photos       []struct {
		Photo Photo `json:"photo"`
	} `json:"taxon_photos"`
}

type UnmarshalSearchResults struct {
	Results []struct {
		Score  float32         `json:"score"`
		Record UnmarshalRecord `json:"record"`
	} `json:"results"`
}

type SortableResult struct {
	Id     string  `json:"id"`
	Score  float32 `json:"score"`
	Term   string  `json:"term"`
	Photos []Photo `json:"photos"`
}

type Search struct {
	Id    string
	Terms *collection.Collection[string]
}

func NewSearch(id, name, scientific_name string, common []string) *Search {
	terms := collection.New(common...)
	terms.Push(name, scientific_name)
	return &Search{
		Id:    id,
		Terms: terms,
	}
}

func (Scrape) Images(ctx context.Context, sourcePath, savePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var parsed UnmarshalData
	err = json.Unmarshal(content, &parsed)

	var searches collection.Collection[*Search]

	for _, p := range parsed.Plants {
		searches.Push(NewSearch(p.Id, p.Name, p.ScientificName, p.Common))
	}

	var sortable_results collection.Collection[*SortableResult]

	searches.Batch(func(b1, j1 int, s1 *Search) {
		s1.Terms.Batch(func(b2, j2 int, s2 string) {
			results, _ := imageSearch(s2)

			var parsed UnmarshalSearchResults
			json.Unmarshal(results, &parsed)

			for _, r := range parsed.Results {
				var photos []Photo
				for _, p := range r.Record.Photos {
					photos = append(photos, p.Photo)
				}

				photos = append(photos, r.Record.DefaultPhoto)

				sortable_results.Push(&SortableResult{
					Id:     s1.Id,
					Score:  r.Score,
					Term:   s2,
					Photos: photos,
				})
			}
		}, s1.Terms.Length())
	}, 2)

	file, _ := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(sortable_results.Items()); err != nil {
		return err
	}

	return err
}

func imageSearch(term string) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	request, _ := http.NewRequest("GET", fmt.Sprintf("https://api.inaturalist.org/v1/search?sources=taxa&q=%s", term), nil)
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		b, _ := io.ReadAll(response.Body)
		return b, nil
	}

	return nil, fmt.Errorf("status code %d", response.StatusCode)
}

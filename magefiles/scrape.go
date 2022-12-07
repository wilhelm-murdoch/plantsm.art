package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

type UnmarshalledPlants struct {
	Data []struct {
		Id     string   `json:"id"`
		Name   string   `json:"name"`
		Common []string `json:"common"`
		Images []struct {
			Url         string `json:"url"`
			SourceUrl   string `json:"source_url"`
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
	}
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
	Id, Term string
}

func NewSearch(id, name string) *Search {
	return &Search{
		Id:   id,
		Term: name,
	}
}

// func (Scrape) Images(ctx context.Context, sourcePath, savePath string) error {
// 	if _, err := os.Stat(sourcePath); err != nil {
// 		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
// 	}

// 	content, err := os.ReadFile(sourcePath)
// 	if err != nil {
// 		return err
// 	}

// 	var parsed UnmarshalledPlants
// 	err = json.Unmarshal(content, &parsed)

// 	var searches collection.Collection[*Search]

// 	for _, p := range parsed.Data {
// 		searches.Push(NewSearch(p.Id, p.Name))
// 	}

// 	var sortable_results collection.Collection[*SortableResult]

// 	limiter := rate.NewLimiter(rate.Every(1*time.Minute/60), 60)

// 	searches.Each(func(i int, s *Search) bool {
// 		limiter.Wait(context.Background())

// 		fmt.Printf("`%s` ( %s ) searching\n", s.Id, s.Term)
// 		results, _ := imageSearch(s.Term)

// 		var parsed UnmarshalSearchResults
// 		json.Unmarshal(results, &parsed)

// 		for _, r := range parsed.Results {
// 			var photos []Photo
// 			for _, p := range r.Record.Photos {
// 				photos = append(photos, p.Photo)
// 			}

// 			photos = append(photos, r.Record.DefaultPhoto)

// 			sortable_results.Push(&SortableResult{
// 				Id:     s.Id,
// 				Score:  r.Score,
// 				Term:   s.Term,
// 				Photos: photos,
// 			})
// 		}
// 		return false
// 	})

// 	file, _ := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	if err := encoder.Encode(sortable_results.Items()); err != nil {
// 		return err
// 	}

// 	return err
// }

type ImageDownload struct {
	Source, Destination string
}

// func (Scrape) DownloadImages(ctx context.Context, sourcePath string) error {
// 	content, err := os.ReadFile(sourcePath)
// 	if err != nil {
// 		return err
// 	}

// 	var parsed UnmarshalledPlants
// 	if err = json.Unmarshal(content, &parsed); err != nil {
// 		return err
// 	}

// 	// index[0] remote path
// 	// index[1] local path
// 	var images collection.Collection[[]string]
// 	for _, plant := range parsed.Data {
// 		for _, image := range *plant.Images {
// 			images.Push([]string{image.SourceUrl, image.Url})
// 		}
// 	}

// 	bar := progressbar.Default(int64(images.Length()))
// 	images.Batch(func(b, j int, image []string) {
// 		defer func() {
// 			bar.Add(1)
// 		}()

// 		request, _ := http.NewRequest("GET", image[0], nil)
// 		response, _ := http.DefaultClient.Do(request)
// 		defer response.Body.Close()

// 		file, _ := os.OpenFile(fmt.Sprintf("data/%s", image[1]), os.O_CREATE|os.O_WRONLY, 0644)
// 		defer file.Close()

// 		io.Copy(file, response.Body)
// 	}, 10)

// 	return nil
// }

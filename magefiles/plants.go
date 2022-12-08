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

type Plants mg.Namespace

func (Plants) Bootstrap(ctx context.Context, path string) error {
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
		WikipediaUrl   string   `json:"wikipedia_url"`
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

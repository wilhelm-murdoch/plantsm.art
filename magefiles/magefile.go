package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	"github.com/wilhelm-murdoch/go-collection"
)

const (
	baseUrlRoot = "https://www.aspca.org"
	baseUrl     = "https://www.aspca.org/pet-care/animal-poison-control/%s-plant-list"
)

func getDocumentFromUrl(url string) (*goquery.Document, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var backoff time.Duration

	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		response, err := client.Get(url)
		if err != nil {
			break
		}

		switch response.StatusCode {
		case 429:
			if attempt >= maxAttempts {
				log.Printf("could not get url %s after %d attempts; skipping ...", url, attempt)
				return goquery.NewDocumentFromReader(response.Body)
			}

			backoff = time.Duration(attempt) * time.Second
			log.Printf("got rate-limited on url %s; waiting another %d seconds", url, (backoff / time.Second))
			time.Sleep(backoff)
		case 200:
			return goquery.NewDocumentFromReader(response.Body)
		case 404:
			log.Printf("url %s could not be found; skipping ...", url)
			return goquery.NewDocumentFromReader(response.Body)
		}

		defer response.Body.Close()
	}

	return nil, errors.New("this should not happen, but here we are")
}

type Classification struct {
	Kingdom string   `json:"kingdom"`
	Clades  []string `json:"clades"`
	Order   string   `json:"order"`
	Family  string   `json:"family"`
	Genus   string   `json:"genus"`
	Species string   `json:"species"`
}

type Plant struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	ScientificName string         `json:"scientific_name"`
	Common         []string       `json:"common"`
	Images         []string       `json:"images"`
	Symptoms       []string       `json:"symptoms"`
	SourceUrl      string         `json:"source_url"`
	Classification Classification `json:"classification"`
}

type Data struct {
	PlantsByCat      *collection.Collection[string] `json:"plants_by_cat"`
	PlantsByDog      *collection.Collection[string] `json:"plants_by_dog"`
	PlantsBySymptoms [][]string                     `json:"plants_by_symptoms"`
	Plants           *collection.Collection[*Plant] `json:"plants"`
}

func (d *Data) NewPlantForAnimal(animal string, s *goquery.Selection) *Plant {
	source_url, _ := s.Find("a").First().Attr("href")

	pattern := regexp.MustCompile(`^(.+?)\s\((.*?)\).+Scientific\sNames:\s?(.+?)?\s\|\sFamily:(.+?)?$`)
	matches := pattern.FindStringSubmatch(s.Text())

	var common_names []string
	for _, name := range strings.Split(matches[2], ",") {
		if name != "" {
			common_names = append(common_names, strings.TrimSpace(name))
		}
	}

	return &Plant{
		Id:             slug.Make(matches[1]),
		Name:           strings.TrimSpace(matches[1]),
		ScientificName: strings.TrimSpace(matches[3]),
		Common:         common_names,
		SourceUrl:      fmt.Sprintf("%s%s", baseUrlRoot, source_url),
		Classification: Classification{
			Family: strings.TrimSpace(matches[4]),
		},
	}
}

func Import(ctx context.Context) error {
	data := &Data{
		PlantsByCat: collection.New[string](),
		PlantsByDog: collection.New[string](),
		Plants:      collection.New[*Plant](),
	}

	if err := importPlantsForAnimal("cats", data); err != nil {
		return err
	}

	if err := importPlantsForAnimal("dogs", data); err != nil {
		return err
	}

	_, err := data.Plants.Batch(func(b, j int, plant *Plant) (*Plant, error) {
		fmt.Println("Processing ...", plant.Id)

		document, err := getDocumentFromUrl(plant.SourceUrl)
		if err != nil {
			return nil, err
		}

		symptoms := document.Find("div.field-name-field-clinical-signs span.values").Text()

		plant.Symptoms = strings.Split(symptoms, ",")

		return nil, nil
	}, 5)

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func importPlantsForAnimal(animal string, data *Data) error {
	if animal != "dogs" && animal != "cats" {
		return fmt.Errorf("expected \"cats\" or \"dogs\", but got \"%s\" instead", animal)
	}

	document, err := getDocumentFromUrl(fmt.Sprintf(baseUrl, animal))
	if err != nil {
		return err
	}

	document.Find("div.view-content").First().Find("div.views-row span.field-content").Each(func(i int, s *goquery.Selection) {
		plant := data.NewPlantForAnimal(animal, s)

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

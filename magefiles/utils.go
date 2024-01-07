package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

type MarshalClassification struct {
	Kingdom string        `json:"kingdom"`
	Clades  []interface{} `json:"clades"`
	Order   string        `json:"order"`
	Family  string        `json:"family"`
	Genus   string        `json:"genus"`
	Species string        `json:"species"`
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
	Toxicity        string                `json:"toxicity"`
	Classification  MarshalClassification `json:"classification"`
}

var supportedAnimals = []string{"cats", "dogs", "horses", "birds", "reptiles", "small-mammals", "fish"}

var containsAnimal = func(animal string, animals []string) bool {
	for _, a := range animals {
		if a == animal {
			return true
		}
	}
	return false
}

type UnmarshalledPlants struct {
	Data []struct {
		Pid    string `json:"pid"`
		Name   string `json:"name"`
		Common []struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"common"`
		Symptoms []struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"symptoms"`
		DateLastUpdated string   `json:"date_last_updated"`
		Animals         []string `json:"animals"`
		Images          []struct {
			RelativePath string `json:"relative_path"`
			SourceUrl    string `json:"source_url"`
			License      string `json:"license"`
			Attribution  string `json:"attribution"`
		} `json:"images"`
		WikipediaUrl   string `json:"wikipedia_url"`
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

func unmarshalPlantsFromSource(sourcePath string) (UnmarshalledPlants, error) {
	if _, err := os.Stat(sourcePath); err != nil {
		return UnmarshalledPlants{}, err
	}

	plantsJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return UnmarshalledPlants{}, err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	return plants, nil
}

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

	return nil, nil
}

func imageSearch(term string) ([]byte, error) {
	url := fmt.Sprintf("https://api.inaturalist.org/v1/search?sources=taxa&q=%s", url.QueryEscape(term))
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var backoff time.Duration
	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		response, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return body, err
		}

		switch response.StatusCode {
		case 429:
			if attempt >= maxAttempts {
				return nil, fmt.Errorf(fmt.Sprintf("could not find %s after %d attempts; skipping ...", term, attempt))
			}

			backoff = time.Duration(attempt) * time.Second
			log.Printf("got rate-limited on term %s; waiting another %d seconds", term, (backoff / time.Second))
			time.Sleep(backoff)
		case 200:
			return body, err
		default:
			return nil, fmt.Errorf("status code %d", response.StatusCode)
		}
	}

	return nil, nil
}

package main

import (
	"encoding/json"
	"os"
)

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

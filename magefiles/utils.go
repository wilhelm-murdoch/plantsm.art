package main

import (
	"encoding/json"
	"os"
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

type MarshalPlant struct {
	Pid          string           `json:"pid"`
	Name         string           `json:"name"`
	Animals      []interface{}    `json:"animals"`
	Common       []MarshalCommon  `json:"common"`
	Symptoms     []MarshalSymptom `json:"symptoms"`
	Images       []MarshalImage   `json:"images"`
	WikipediaUrl string           `json:"wikipedia_url"`
	Toxicity     string           `json:"toxicity"`
	Family       string           `json:"family"`
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
		Animals []string `json:"animals"`
		Images  []struct {
			RelativePath string `json:"relative_path"`
			SourceUrl    string `json:"source_url"`
			License      string `json:"license"`
			Attribution  string `json:"attribution"`
		} `json:"images"`
		WikipediaUrl string `json:"wikipedia_url"`
		Family       string `json:"family"`
	}
}

type UnmarshalledSeverity struct {
	Data []*struct {
		Label    string   `json:"label"`
		Slug     string   `json:"slug"`
		Level    int      `json:"level"`
		Symptoms []string `json:"symptoms"`
		Plants   []string `json:"plants"`
	}
}

func unmarshalSeverityFromSource(sourcePath string) (*UnmarshalledSeverity, error) {
	if _, err := os.Stat(sourcePath); err != nil {
		return &UnmarshalledSeverity{}, err
	}

	severityJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return &UnmarshalledSeverity{}, err
	}

	var severity UnmarshalledSeverity
	json.Unmarshal(severityJson, &severity.Data)

	return &severity, nil
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

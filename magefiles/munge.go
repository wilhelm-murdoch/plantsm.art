package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

type Munge mg.Namespace

func (Munge) Symptoms(ctx context.Context, sourcePath, writePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	plantsJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	var symptoms collection.Collection[string]
	for _, plant := range plants.Data {
		symptoms.PushDistinct(plant.Symptoms...)
	}

	symptoms_by_plant := make(map[string][]string)
	symptoms.Each(func(i int, symptom string) bool {
		for _, plant := range plants.Data {
			for _, s := range plant.Symptoms {
				if symptom == s {
					symptoms_by_plant[symptom] = append(symptoms_by_plant[symptom], plant.Id)
				}
			}
		}
		return false
	})

	file, _ := os.OpenFile(writePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(symptoms_by_plant); err != nil {
		return err
	}

	return nil
}

func (Munge) Clean(ctx context.Context, animal, sourcePath string) error {
	if animal != "dogs" && animal != "cats" {
		return fmt.Errorf("expected \"cats\" or \"dogs\", but got \"%s\" instead", animal)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	dataJson, err := os.ReadFile(fmt.Sprintf("data/plants-by-%s.json", animal))
	if err != nil {
		return err
	}

	var data []string
	json.Unmarshal(dataJson, &data)

	var results []string
	for _, id := range data {
		for _, plant := range plants.Data {
			if id == plant.Id {
				results = append(results, id)
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	out, _ := json.Marshal(results)

	fmt.Print(string(out))

	return nil
}

func (Munge) Wikipedia(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	for _, plant := range plants.Data {
		if plant.WikipediaUrl == "" {
			fmt.Println(plant.Id)
		}
	}

	return nil
}

func (Munge) AssignAnimals(ctx context.Context, animal, sourcePath string) error {
	if animal != "dog" && animal != "cat" {
		return fmt.Errorf("expected \"cat\" or \"dog\", but got \"%s\" instead", animal)
	}

	plantsJson, err := os.ReadFile("data/plants.json")
	if err != nil {
		return err
	}

	var plants interface{}
	json.Unmarshal(plantsJson, &plants)

	dataJson, err := os.ReadFile(fmt.Sprintf("data/plants-by-%s.json", animal))
	if err != nil {
		return err
	}

	var data []string
	json.Unmarshal(dataJson, &data)

	for _, id := range data {
		for _, plant := range plants.([]interface{}) {
			if id == plant.(map[string]interface{})["id"].(string) {
				plant.(map[string]interface{})["animals"] = append(plant.(map[string]interface{})["animals"].([]interface{}), animal)
			}
		}
	}

	file, _ := os.OpenFile(sourcePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(plants); err != nil {
		return err
	}

	return nil
}

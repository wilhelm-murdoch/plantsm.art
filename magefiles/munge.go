package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
		for _, symptom := range plant.Symptoms {
			symptoms.PushDistinct(symptom.Name)
		}
	}

	symptoms_by_plant := make(map[string][]string)
	symptoms.Each(func(i int, symptom string) bool {
		for _, plant := range plants.Data {
			for _, s := range plant.Symptoms {
				if symptom == s.Name {
					symptoms_by_plant[symptom] = append(symptoms_by_plant[symptom], plant.Pid)
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

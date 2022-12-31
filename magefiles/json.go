package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

type Json mg.Namespace

func (Json) Pages(ctx context.Context, sourcePath, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	for _, plant := range plants.Data {
		file, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", toPath, plant.Pid), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(plant); err != nil {
			return err
		}

		fmt.Printf("Wrote: %s/%s.json\n", toPath, plant.Pid)
	}

	return nil
}

type SymptomItem struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Plants int    `json:"plants"`
}

func (Json) Symptoms(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var symptoms collection.Collection[*SymptomItem]

	for _, plant := range plants.Data {
		for _, symptom := range plant.Symptoms {
			if found := symptoms.Find(func(i int, s *SymptomItem) bool {
				return s.Slug == symptom.Slug
			}); found != nil {
				found.Plants += 1
			} else {
				symptoms.PushDistinct(&SymptomItem{
					Name:   symptom.Name,
					Slug:   symptom.Slug,
					Plants: 1,
				})
			}
		}
	}

	if err := json.NewEncoder(os.Stdout).Encode(symptoms.Items()); err != nil {
		panic(err)
	}

	return nil
}

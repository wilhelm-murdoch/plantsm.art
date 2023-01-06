package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

var supportedAnimals = []string{"cats", "dogs", "horses", "rats", "birds", "rabbits", "reptiles", "hamsters"}

var found = func(animal string, animals []string) bool {
	for _, a := range animals {
		if a == animal {
			return true
		}
	}
	return false
}

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
	Name   string                        `json:"name"`
	Slug   string                        `json:"slug"`
	Count  int                           `json:"count"`
	Plants collection.Collection[string] `json:"plants"`
}

func (Json) Symptoms(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var symptoms collection.Collection[*SymptomItem]

	for _, plant := range plants.Data {
		for _, symptom := range plant.Symptoms {
			symptoms.PushDistinct(&SymptomItem{
				Name: symptom.Name,
				Slug: symptom.Slug,
			})
		}
	}

	symptoms.Each(func(i int, s *SymptomItem) bool {
		for _, plant := range plants.Data {
			for _, symptom := range plant.Symptoms {
				if symptom.Slug == s.Slug {
					s.Count += 1
					s.Plants.PushDistinct(plant.Pid)
				}
			}
		}
		return false
	})

	if err := json.NewEncoder(os.Stdout).Encode(symptoms.Items()); err != nil {
		panic(err)
	}

	return nil
}

type AnimalItem struct {
	Animal string   `json:"animal"`
	Count  int      `json:"count"`
	Plants []string `json:"plants"`
}

func (Json) Animals(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var animalItems []AnimalItem

	var animalItem AnimalItem
	for _, animal := range supportedAnimals {
		animalItem = AnimalItem{
			Animal: animal,
		}

		for _, plant := range plants.Data {
			if found(animal, plant.Animals) {
				animalItem.Count += 1
				animalItem.Plants = append(animalItem.Plants, plant.Pid)
			}
		}

		animalItems = append(animalItems, animalItem)
	}

	if err := json.NewEncoder(os.Stdout).Encode(animalItems); err != nil {
		panic(err)
	}

	return nil
}

func (Json) Animal(ctx context.Context, sourcePath, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var animalItem AnimalItem
	for _, animal := range supportedAnimals {
		animalItem = AnimalItem{
			Animal: animal,
		}

		for _, plant := range plants.Data {
			if found(animal, plant.Animals) {
				animalItem.Count += 1
				animalItem.Plants = append(animalItem.Plants, plant.Pid)
			}
		}

		file, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", toPath, animal), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(animalItem); err != nil {
			return err
		}

		fmt.Printf("Wrote: %s/%s.json\n", toPath, animal)
	}

	return nil
}

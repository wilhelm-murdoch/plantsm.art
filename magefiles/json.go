package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gosimple/slug"
	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

type Json mg.Namespace

type SlimPlant struct {
	Pid           string   `json:"pid"`
	Name          string   `json:"name"`
	Common        []string `json:"common"`
	CommonTotal   int      `json:"common_total"`
	Symptoms      []string `json:"symptoms"`
	SymptomsTotal int      `json:"symptoms_total"`
	Animals       []string `json:"animals"`
	CoverImageUrl string   `json:"cover_image_url"`
	ImageTotal    int      `json:"image_total"`
	SearchIndex   string   `json:"search_index"`
	IsDeadly      bool     `json:"is_deadly"`
	Family        string   `json:"family"`
}

func (Json) Slim(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var convertToSlim = func(items interface{}) []string {
		var output []string

		for _, i := range items.([]struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		}) {
			output = append(output, i.Name)
		}

		return output
	}

	var SlimPlants []SlimPlant
	var IsDeadly bool
	for _, plant := range plants.Data {
		slimCommon := convertToSlim(plant.Common)
		slimSymptoms := convertToSlim(plant.Symptoms)

		var searchIndex = strings.ToLower(plant.Classification.Family)

		searchIndex += " " + strings.ToLower(plant.Name)

		for _, a := range plant.Animals {
			searchIndex += " " + strings.ToLower(a)
		}

		for _, s := range slimSymptoms {
			if strings.ToLower(s) == "death" {
				IsDeadly = true
			}

			searchIndex += " " + strings.ToLower(s)
		}

		for _, c := range slimCommon {
			searchIndex += " " + strings.ToLower(c)
		}

		if len(slimCommon) >= 3 {
			slimCommon = slimCommon[0:3]
		}

		if len(slimSymptoms) >= 3 {
			slimSymptoms = slimSymptoms[0:3]
		}

		SlimPlants = append(SlimPlants, SlimPlant{
			Pid:           plant.Pid,
			Name:          plant.Name,
			Common:        slimCommon,
			CommonTotal:   len(plant.Common),
			Symptoms:      slimSymptoms,
			SymptomsTotal: len(plant.Symptoms),
			Animals:       plant.Animals,
			CoverImageUrl: plant.Images[0].RelativePath,
			ImageTotal:    len(plant.Images),
			SearchIndex:   searchIndex,
			IsDeadly:      IsDeadly,
			Family:        plant.Classification.Family,
		})

		IsDeadly = false
	}

	if err := json.NewEncoder(os.Stdout).Encode(SlimPlants); err != nil {
		panic(err)
	}

	return nil
}

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

type FamilyItem struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Count int    `json:"count"`
}

func (Json) Families(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var families collection.Collection[*FamilyItem]

	for _, plant := range plants.Data {
		families.PushDistinct(&FamilyItem{
			Name: plant.Classification.Family,
			Slug: slug.Make(plant.Classification.Family),
		})
	}

	families.Each(func(i int, s *FamilyItem) bool {
		for _, plant := range plants.Data {
			if slug.Make(plant.Classification.Family) == s.Slug {
				s.Count += 1
			}
		}

		return false
	})

	if err := json.NewEncoder(os.Stdout).Encode(families.Items()); err != nil {
		panic(err)
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
			if containsAnimal(animal, plant.Animals) {
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
			if containsAnimal(animal, plant.Animals) {
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

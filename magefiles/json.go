package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	Severity      Severity `json:"severity"`
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
	for _, plant := range plants.Data {
		slimCommon := convertToSlim(plant.Common)
		slimSymptoms := convertToSlim(plant.Symptoms)

		var searchIndex = strings.ToLower(plant.Family)

		searchIndex += " " + strings.ToLower(plant.Name)

		for _, a := range plant.Animals {
			searchIndex += " " + strings.ToLower(a)
		}

		for _, s := range slimSymptoms {
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
			Severity:      plant.Severity,
			Family:        plant.Family,
		})
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
			Name: plant.Family,
			Slug: slug.Make(plant.Family),
		})
	}

	families.Each(func(i int, s *FamilyItem) bool {
		for _, plant := range plants.Data {
			if slug.Make(plant.Family) == s.Slug {
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

func (Json) Severity(ctx context.Context, sourcePath, severityPath string) error {
	severity, err := unmarshalSeverityFromSource(severityPath)
	if err != nil {
		return err
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	for _, tox := range severity.Data {
		for _, symptom := range tox.Symptoms {
			for _, plant := range plants.Data {
				if containsSymptom(symptom, plant.Symptoms) && !containsPlant(plant.Pid, tox.Plants) {
					tox.Plants = append(tox.Plants, plant.Pid)
				}
			}
		}
	}

	if err := json.NewEncoder(os.Stdout).Encode(severity.Data); err != nil {
		panic(err)
	}

	return nil
}

type SortPlantsBySeverity []*struct {
	Label    string   `json:"label"`
	Slug     string   `json:"slug"`
	Level    int      `json:"level"`
	Symptoms []string `json:"symptoms"`
	Plants   []string `json:"plants"`
}

func (s SortPlantsBySeverity) Len() int           { return len(s) }
func (s SortPlantsBySeverity) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortPlantsBySeverity) Less(i, j int) bool { return s[i].Level < s[j].Level }

func (Json) SeverityToPlant(ctx context.Context, sourcePath, severityPath string) error {
	severities, err := unmarshalSeverityFromSource(severityPath)
	if err != nil {
		return err
	}

	sort.Sort(SortPlantsBySeverity(severities.Data))

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var getSeverityFromPlant = func(plant *UnmarshalledPlant) Severity {
		out := Severity{
			Label: severities.Data[0].Label,
			Slug:  severities.Data[0].Slug,
			Level: severities.Data[0].Level,
		}

		for _, severity := range severities.Data {
			if containsPlant(plant.Pid, severity.Plants) && severity.Level > out.Level {
				out.Label = severity.Label
				out.Slug = severity.Slug
				out.Level = severity.Level
			}
		}

		return out
	}

	for idx, plant := range plants.Data {
		severity := getSeverityFromPlant(plant)

		plants.Data[idx].Severity = severity
	}

	if err := json.NewEncoder(os.Stdout).Encode(plants.Data); err != nil {
		panic(err)
	}

	return nil
}

func containsPlant(pid string, pids []string) bool {
	for _, p := range pids {
		if p == pid {
			return true
		}
	}
	return false
}

func containsSymptom(symptom string, symptoms []struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}) bool {
	for _, s := range symptoms {
		if s.Slug == slug.Make(symptom) {
			return true
		}
	}

	return false
}

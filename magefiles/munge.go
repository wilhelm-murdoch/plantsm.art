package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gosimple/slug"
	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

type Munge mg.Namespace

type UnmarshalPhoto struct {
	Photos []SortableResult `json:"photos"`
}

func (Munge) Photos(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var parsed UnmarshalPhoto
	json.Unmarshal(content, &parsed)

	var photos collection.Collection[SortableResult]
	photos.Push(parsed.Photos...)

	plantsJson, err := os.ReadFile("data/plants.json")
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants)

	var ids collection.Collection[string]
	for _, plant := range plants.Data {
		ids.Push(plant.Id)
	}

	plant_id_to_photos := make(map[string][]Photo)

	ids.Each(func(i int, id string) bool {
		skip := false

		filtered := photos.Filter(func(sr SortableResult) bool {
			return sr.Id == id
		})

		if filtered.Length() == 0 {
			skip = true
		}

		filtered.Sort(func(i, j int) bool {
			left, _ := filtered.At(i)
			right, _ := filtered.At(j)
			return left.Score > right.Score
		})

		var photos []Photo

		if first, ok := filtered.At(0); ok && first.Score >= 5 {
			photos = append(photos, first.Photos...)
		}

		if second, ok := filtered.At(1); ok && second.Score >= 5 {
			photos = append(photos, second.Photos...)
		}

		if third, ok := filtered.At(2); ok && third.Score >= 5 {
			photos = append(photos, third.Photos...)
		}

		p := collection.New(photos...).Filter(func(p Photo) bool {
			return p.License == "cc-by-nc" || p.License == "cc0" || p.License == "cc-by-sa" || p.License == "cc-by-nc-sa" || p.License == "cc-by"
		})

		if p.Length() <= 0 {
			skip = true
		}

		p = p.Filter(func(p Photo) bool {
			return p.UrlMedium != ""
		})

		if p.Length() < 3 {
			skip = true
		}

		if skip {
			fmt.Printf("no images found for %s\n", id)
		} else {
			plant_id_to_photos[id] = p.Items()
		}

		return false
	})

	var derp interface{}
	json.Unmarshal(plantsJson, &derp)

	for _, plant := range derp.(map[string]interface{})["plants"].([]interface{}) {
		id := plant.(map[string]interface{})["id"].(string)
		name := plant.(map[string]interface{})["scientific_name"].(string)
		if photos, ok := plant_id_to_photos[id]; ok {
			for i, photo := range photos {
				plant.(map[string]interface{})["images"] = append(plant.(map[string]interface{})["images"].([]interface{}), struct {
					Url         string `json:"url"`
					SourceUrl   string `json:"source_url"`
					License     string `json:"license"`
					Attribution string `json:"attribution"`
				}{
					Url:         fmt.Sprintf("plants/%s/images/%d.jpg", slug.Make(name), i+1),
					SourceUrl:   photo.UrlMedium,
					License:     photo.License,
					Attribution: photo.Attribution,
				})
			}
		}
	}

	file, _ := os.OpenFile("data/plants-photos.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(derp); err != nil {
		return err
	}

	return nil
}

func (Munge) PlantsBySymptom(ctx context.Context) error {
	plantsJson, err := os.ReadFile("data/plants.json")
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	file, _ := os.OpenFile("data/plants.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	var symptoms collection.Collection[string]
	for _, plant := range plants.Data {
		symptoms.PushDistinct(plant.Symptoms...)
	}

	var symptoms_by_plant map[string][]string
	symptoms.Each(func(i int, symptom string) bool {
		for _, plant := range plants.Data {
			for _, s := range plant.Symptoms {
				if symptom == s {
					// symptoms_by_plant[symptom] = append(symptoms_by_plant[symptom], plant.Id)
				}
			}
		}
		return false
	})

	file, _ = os.OpenFile("data/plants-by-symptom.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(symptoms_by_plant); err != nil {
		return err
	}

	return nil
}

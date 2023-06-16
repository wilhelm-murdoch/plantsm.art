package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gosimple/slug"
	"github.com/magefile/mage/mg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Reset mg.Namespace

type ResetSymptom struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ResetCommon struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ResetImage struct {
	SourceUrl    string `json:"source_url"`
	Attribution  string `json:"attribution"`
	License      string `json:"license"`
	RelativePath string `json:"relative_path"`
}

type ResetPlant struct {
	Pid             string              `json:"pid"`
	Name            string              `json:"name"`
	Animals         []string            `json:"animals"`
	Common          []ResetCommon       `json:"common"`
	Symptoms        []ResetSymptom      `json:"symptoms"`
	Images          []ResetImage        `json:"images"`
	WikipediaUrl    string              `json:"wikipedia_url"`
	DateLastUpdated string              `json:"date_last_updated"`
}

func (Reset) All(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var resetPlantsOutput []ResetPlant
	for _, plant := range plants.Data {
		fmt.Println("Resetting images and common names for:", plant.Name)
		var symptoms []ResetSymptom

		for _, s := range plant.Symptoms {
			symptoms = append(symptoms, ResetSymptom{
				Name: s.Name,
				Slug: s.Slug,
			})
		}

		resetPlant := &ResetPlant{
			Pid:             plant.Pid,
			Name:            plant.Name,
			Animals:         plant.Animals,
			Symptoms:        symptoms,
			WikipediaUrl:    plant.WikipediaUrl,
			DateLastUpdated: plant.DateLastUpdated,
		}

		resetPlant, err = getiNaturalistImagesReset(resetPlant)
		if err != nil {
			return err
		}

		resetPlantsOutput = append(resetPlantsOutput, *resetPlant)
	}

	if err := json.NewEncoder(os.Stdout).Encode(resetPlantsOutput); err != nil {
		panic(err)
	}

	return nil
}

func getiNaturalistImagesReset(plant *ResetPlant) (*ResetPlant, error) {
	results, err := imageSearch(strings.TrimSuffix(strings.TrimSuffix(plant.Name, "sp."), "spp."))
	if err != nil {
		return nil, err
	}

	var naturalist interface{}
	json.Unmarshal(results, &naturalist)

	records := naturalist.(map[string]interface{})["results"].([]interface{})
	for _, r := range records {
		record := r.(map[string]interface{})["record"].(map[string]interface{})
		iconicTaxonName := record["iconic_taxon_name"]
		if iconicTaxonName == nil {
			continue
		}

		if iconicTaxonName.(string) == "Plantae" {
			if record["preferred_common_name"] != nil {
				plant.Common = append(plant.Common, ResetCommon{
					Name: cases.Title(language.AmericanEnglish, cases.NoLower).String(record["preferred_common_name"].(string)),
					Slug: slug.Make(record["preferred_common_name"].(string)),
				})
			}

			if plant.WikipediaUrl == "" && record["wikipedia_url"] != nil {
				plant.WikipediaUrl = record["wikipedia_url"].(string)
			}

			for _, photo := range record["taxon_photos"].([]interface{}) {
				licenseCode := photo.(map[string]interface{})["photo"].(map[string]interface{})["license_code"]
				if licenseCode == nil {
					continue
				}

				sourceUrl := photo.(map[string]interface{})["photo"].(map[string]interface{})["large_url"]
				if sourceUrl == nil {
					sourceUrl = photo.(map[string]interface{})["photo"].(map[string]interface{})["medium_url"]
				}

				attribution := photo.(map[string]interface{})["photo"].(map[string]interface{})["attribution"]
				if attribution == nil {
					continue
				}

				plant.Images = append(plant.Images, ResetImage{
					License:     licenseCode.(string),
					Attribution: attribution.(string),
					SourceUrl:   sourceUrl.(string),
				})
			}
		}
	}

	if len(plant.Images) >= 10 {
		plant.Images = plant.Images[0:9]
	}

	return plant, nil
}

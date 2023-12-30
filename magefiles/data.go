package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gosimple/slug"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/magefile/mage/mg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Data mg.Namespace

var plantExistsByName = func(plantName string, plants UnmarshalledPlants) (string, bool) {
	for _, plant := range plants.Data {
		if plant.Name == plantName || fuzzy.RankMatchFold(plantName, plant.Name) != -1 {
			return plant.Pid, true
		}
	}

	return "", false
}

var symptomItemsFromString = func(symptoms string) []MarshalSymptom {
	var symptomItems []MarshalSymptom
	for _, symptom := range strings.Split(symptoms, ",") {
		symptomItems = append(symptomItems, MarshalSymptom{
			Name: cases.Title(language.AmericanEnglish, cases.NoLower).String(symptom),
			Slug: slug.Make(symptom),
		})
	}
	return symptomItems
}

func (Data) AddAnimalToPlantsFromFile(ctx context.Context, sourcePath, plantFilePath, animal string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}
	if !containsAnimal(animal, supportedAnimals) {
		return fmt.Errorf("animal %s not supported", animal)
	}

	plantFileBytes, err := ioutil.ReadFile(plantFilePath)
	if err != nil {
		return err
	}

	plantNames := strings.Split(string(plantFileBytes), "\n")

	var nameInNames = func(name string, names []string) bool {
		for _, n := range names {
			if name == n {
				return true
			}
		}

		return false
	}

	var resetPlantsOutput []ResetPlant

	for _, plant := range plants.Data {
		var symptoms []ResetSymptom
		for _, s := range plant.Symptoms {
			symptoms = append(symptoms, ResetSymptom{
				Name: s.Name,
				Slug: s.Slug,
			})
		}

		var images []ResetImage
		for _, i := range plant.Images {
			images = append(images, ResetImage{
				SourceUrl:    i.SourceUrl,
				Attribution:  i.Attribution,
				License:      i.License,
				RelativePath: i.RelativePath,
			})
		}

		var common []ResetCommon
		for _, c := range plant.Common {
			common = append(common, ResetCommon{
				Name: c.Name,
				Slug: c.Slug,
			})
		}

		if nameInNames(plant.Name, plantNames) && !containsAnimal(animal, plant.Animals) {
			plant.Animals = append(plant.Animals, animal)
		}

		resetPlant := &ResetPlant{
			Pid:             plant.Pid,
			Name:            plant.Name,
			Animals:         plant.Animals,
			Common:          common,
			Symptoms:        symptoms,
			WikipediaUrl:    plant.WikipediaUrl,
			DateLastUpdated: plant.DateLastUpdated,
			Images:          images,
		}

		resetPlantsOutput = append(resetPlantsOutput, *resetPlant)
	}

	if err := json.NewEncoder(os.Stdout).Encode(resetPlantsOutput); err != nil {
		panic(err)
	}

	return nil
}

func (Data) File(ctx context.Context, sourcePath, plantFilePath, animal, symptoms string) error {
	if _, err := os.Stat(plantFilePath); err != nil {
		return fmt.Errorf("cannot plant source path %s", plantFilePath)
	}

	if !containsAnimal(animal, supportedAnimals) {
		return fmt.Errorf("animal %s not supported", animal)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	plantFile, err := os.Open(plantFilePath)
	if err != nil {
		return err
	}
	defer plantFile.Close()

	var newPlantsOutput []MarshalPlant

	scanner := bufio.NewScanner(plantFile)
	for scanner.Scan() {
		plantName := scanner.Text()
		if _, found := plantExistsByName(strings.Replace(plantName, "sp", "", -1), plants); found {
			continue
		}

		newPlant := &MarshalPlant{
			Pid:      slug.Make(plantName),
			Name:     plantName,
			Animals:  []interface{}{animal},
			Symptoms: symptomItemsFromString(symptoms),
		}

		newPlant, err = getiNaturalistImages(newPlant)
		if err != nil {
			return err
		}

		newPlant, err = getWikipediaDetails(newPlant)
		if err != nil {
			return err
		}

		newPlantsOutput = append(newPlantsOutput, *newPlant)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(newPlantsOutput); err != nil {
		panic(err)
	}

	return nil
}

func (Data) Scaffold(ctx context.Context, sourcePath, animal, plantName string) error {
	if !containsAnimal(animal, supportedAnimals) {
		return fmt.Errorf("animal %s not supported", animal)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	pid, found := plantExistsByName(strings.Replace(plantName, "sp", "", -1), plants)
	if found {
		return fmt.Errorf("plant exists with %s for %s", pid, plantName)
	}

	newPlant := &MarshalPlant{
		Pid:     slug.Make(plantName),
		Name:    plantName,
		Animals: []interface{}{animal},
	}

	newPlant, err = getiNaturalistImages(newPlant)
	if err != nil {
		return err
	}

	newPlant, err = getWikipediaDetails(newPlant)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(os.Stdout).Encode(newPlant); err != nil {
		panic(err)
	}

	return nil
}

type SortableImageResults struct {
	Id     float64
	Score  float64
	Images []MarshalImage
}

func getiNaturalistImages(plant *MarshalPlant) (*MarshalPlant, error) {
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
				plant.Common = append(plant.Common, MarshalCommon{
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

				attribution := photo.(map[string]interface{})["photo"].(map[string]interface{})["attribution"]
				if attribution == nil {
					continue
				}

				imageNameHash := md5.Sum([]byte(sourceUrl.(string)))

				plant.Images = append(plant.Images, MarshalImage{
					License:      licenseCode.(string),
					Attribution:  attribution.(string),
					SourceUrl:    sourceUrl.(string),
					RelativePath: fmt.Sprintf("plants/%s/%s/large.jpg", plant.Pid, hex.EncodeToString(imageNameHash[:])),
				})
			}
		}
	}

	if len(plant.Images) >= 10 {
		plant.Images = plant.Images[0:9]
	}

	return plant, nil
}

func getWikipediaDetails(plant *MarshalPlant) (*MarshalPlant, error) {
	wikipediaUrl := plant.WikipediaUrl
	if wikipediaUrl == "" {
		wikipediaUrl = fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.TrimSuffix(strings.TrimSuffix(plant.Name, "sp."), "spp."))
	}

	return plant, nil
}

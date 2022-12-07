package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Classify mg.Namespace

func (Classify) Plants(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	plantsJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	var names collection.Collection[string]
	for _, plant := range plants.Data {
		names.PushDistinct(TrimSuffix(plant.Name, " sp."))
	}

	var derp interface{}
	json.Unmarshal(plantsJson, &derp)

	bar := progressbar.Default(int64(names.Length()))
	names.Batch(func(i1, i2 int, name string) {
		defer func() {
			bar.Add(1)
		}()

		document, _ := getDocumentFromUrl(fmt.Sprintf("https://en.wikipedia.org/wiki/%s", name))

		var clades []string
		document.Find("table.biota").Each(func(i int, s *goquery.Selection) {
			s.Find("tr:contains('Clade:') td").Each(func(i int, s *goquery.Selection) {
				clade := strings.TrimSpace(s.First().Next().Text())
				if clade != "" {
					clades = append(clades, clade)
				}
			})

			for _, plant := range derp.([]interface{}) {
				if plant.(map[string]interface{})["id"].(string) == slug.Make(name) {
					plant.(map[string]interface{})["wikipedia_url"] = fmt.Sprintf("https://en.wikipedia.org/wiki/%s", name)
					plant.(map[string]interface{})["classification"] = struct {
						Kingdom string   `json:"kingdom"`
						Clades  []string `json:"clades"`
						Order   string   `json:"order"`
						Family  string   `json:"family"`
						Genus   string   `json:"genus"`
						Species string   `json:"species"`
					}{
						Kingdom: strings.TrimSpace(s.Find("tr:contains('Kingdom:') td").First().Next().Text()),
						Order:   strings.TrimSpace(s.Find("tr:contains('Order:') td").First().Next().Text()),
						Family:  strings.TrimSpace(s.Find("tr:contains('Family:') td").First().Next().Text()),
						Genus:   strings.TrimSpace(s.Find("tr:contains('Genus:') td").First().Next().Text()),
						Species: strings.TrimSpace(s.Find("tr:contains('Species:') td").First().Next().Text()),
						Clades:  clades,
					}
				}
			}
		})
	}, 10)

	file, _ := os.OpenFile("data/plants-classified.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(derp); err != nil {
		return err
	}

	return nil
}

func (Classify) Missing(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	plantsJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	var missing []string
	var count int
	for _, plant := range plants.Data {
		if plant.Classification.Family == "" {
			missing = append(missing, "family")
		}

		if plant.Classification.Kingdom == "" {
			missing = append(missing, "kingdom")
		}

		if plant.Classification.Order == "" {
			missing = append(missing, "order")
		}

		if plant.Classification.Genus == "" {
			missing = append(missing, "genus")
		}

		if plant.Classification.Species == "" {
			missing = append(missing, "species")
		}

		if len(plant.Classification.Clades) == 0 {
			missing = append(missing, "clades")
		}

		if len(missing) > 0 {
			fmt.Println(plant.Name, "is missing:")
			for _, m := range missing {
				fmt.Println("-", m)
			}

			count++
		}

		missing = nil
	}

	fmt.Println(count, "with missing classification data")

	return nil
}

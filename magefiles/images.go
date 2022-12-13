package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
	"golang.org/x/time/rate"
)

type Images mg.Namespace

func (Images) Bootstrap(ctx context.Context, sourcePath, savePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var parsed UnmarshalledPlants
	err = json.Unmarshal(content, &parsed.Data)

	var searches collection.Collection[*Search]

	for _, p := range parsed.Data {
		if len(p.Images) == 0 {
			searches.Push(NewSearch(p.Id, p.Name))
		}
	}

	var sortable_results collection.Collection[*SortableResult]

	limiter := rate.NewLimiter(rate.Every(1*time.Minute/60), 60)

	searches.Each(func(i int, s *Search) bool {
		limiter.Wait(context.Background())

		fmt.Printf("`%s` ( %s ) searching\n", s.Id, s.Term)
		results, _ := imageSearch(strings.TrimSuffix(s.Term, "sp."))

		var parsed UnmarshalSearchResults
		json.Unmarshal(results, &parsed)

		fmt.Printf("`%s` ( %s ) found %d\n", s.Id, s.Term, len(parsed.Results))

		for _, r := range parsed.Results {
			var photos []Photo
			for _, p := range r.Record.Photos {
				photos = append(photos, p.Photo)
			}

			photos = append(photos, r.Record.DefaultPhoto)

			sortable_results.Push(&SortableResult{
				Id:     s.Id,
				Score:  r.Score,
				Term:   s.Term,
				Photos: photos,
			})
		}

		return false
	})

	file, _ := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(sortable_results.Items()); err != nil {
		return err
	}

	return err
}

type UnmarshalPhoto struct {
	Photos []SortableResult `json:"photos"`
}

func (Images) Filter(ctx context.Context, sourcePath, writePath string) error {
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
	json.Unmarshal(plantsJson, &plants.Data)

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
			return p.License == "cc-by-nc-nd" || p.License == "cc-by-nc" || p.License == "cc0" || p.License == "cc-by-sa" || p.License == "cc-by-nc-sa" || p.License == "cc-by"
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

	for _, plant := range derp.([]interface{}) {
		id := plant.(map[string]interface{})["id"].(string)
		name := plant.(map[string]interface{})["name"].(string)
		if len(plant.(map[string]interface{})["images"].([]interface{})) <= 0 {
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
	}

	file, _ := os.OpenFile(writePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(derp); err != nil {
		return err
	}

	return nil
}

type ImageDownload struct {
	Source, Destination string
}

func (Images) Download(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	// index[0] remote path
	// index[1] local path
	var images collection.Collection[[]string]
	for _, plant := range plants.Data {
		for _, image := range plant.Images {
			if u, err := url.Parse(image.Url); err == nil {
				parts := strings.Split(u.Path, ".")
				path := parts[0]
				ext := parts[1]

				images.Push([]string{image.SourceUrl, fmt.Sprintf("%s-square.%s", path, ext)})
			}
		}
	}

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image []string) {
		defer func() {
			bar.Add(1)
		}()

		if _, err := os.Stat(fmt.Sprintf("data/%s", image[1])); errors.Is(err, os.ErrNotExist) {
			request, _ := http.NewRequest("GET", image[0], nil)
			response, _ := http.DefaultClient.Do(request)
			defer response.Body.Close()

			file, _ := os.OpenFile(fmt.Sprintf("data/%s", image[1]), os.O_CREATE|os.O_WRONLY, 0644)
			defer file.Close()

			io.Copy(file, response.Body)
		}
	}, 10)

	return nil
}

func (Images) Missing(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate scaffold file %s; regenerate and try again", sourcePath)
	}

	plantsJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var plants UnmarshalledPlants
	json.Unmarshal(plantsJson, &plants.Data)

	var count int
	for _, plant := range plants.Data {
		if len(plant.Images) == 0 {
			fmt.Println(plant.Name, plant.Id)
			count++
		}
	}

	fmt.Println("found", count, "plants without images")

	return nil
}

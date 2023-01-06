package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Images mg.Namespace

func (Images) Download(ctx context.Context, sourcePath, writePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	// index[0] remote path
	// index[1] local path
	var images collection.Collection[[]string]
	for _, plant := range plants.Data {
		for _, image := range plant.Images {
			if err := os.MkdirAll(fmt.Sprintf("%s/%s", writePath, strings.TrimSuffix(image.RelativePath, "/original.jpg")), os.ModePerm); err != nil {
				fmt.Println(err)
			}

			images.Push([]string{image.SourceUrl, fmt.Sprintf("%s/%s", writePath, image.RelativePath)})
		}
	}

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image []string) {
		defer func() {
			bar.Add(1)
		}()

		if _, err := os.Stat(image[1]); errors.Is(err, os.ErrNotExist) {
			response, err := http.Get(image[0])
			if err != nil {
				fmt.Println(err)
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				fmt.Println(err)
			}

			var data bytes.Buffer
			_, err = io.Copy(&data, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			if err := os.WriteFile(image[1], data.Bytes(), 0644); err != nil {
				fmt.Println(err)
			}
		}
	}, 10)

	return nil
}

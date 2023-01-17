package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Images mg.Namespace

type CloudflareImage struct {
	Url, Id string
}

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
			if err := os.MkdirAll(fmt.Sprintf("%s/%s", writePath, strings.TrimSuffix(image.RelativePath, "/large.jpg")), os.ModePerm); err != nil {
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

func (Images) Cloudflare(ctx context.Context, sourcePath string) error {
	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	var images collection.Collection[CloudflareImage]
	for _, plant := range plants.Data {
		for _, image := range plant.Images {
			images.Push(CloudflareImage{image.SourceUrl, fmt.Sprintf("images/%s", strings.TrimSuffix(image.RelativePath, "/large.jpg"))})
		}
	}

	client := &http.Client{}

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image CloudflareImage) {
		defer func() {
			bar.Add(1)
		}()

		request, err := makeRequest(image.Url, image.Id)
		if err != nil {
			panic(err)
		}

		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if response.StatusCode != 200 {
			fmt.Printf("Failed: %s\n%s\n\n", image.Id, string(body))
		}
	}, 25)

	return nil
}

func makeRequest(url, id string) (*http.Request, error) {
	var data bytes.Buffer
	var fw io.Writer
	form := multipart.NewWriter(&data)

	fw, err := form.CreateFormField("url")
	if err != nil {
		return nil, err
	}
	fw.Write([]byte(url))

	fw, err = form.CreateFormField("id")
	if err != nil {
		return nil, err
	}
	fw.Write([]byte(id))

	form.Close()

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/images/v1", os.Getenv("CLOUDFLARE_ACCOUNT_ID")), &data)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_API_TOKEN")))
	request.Header.Set("Content-Type", form.FormDataContentType())

	return request, nil
}

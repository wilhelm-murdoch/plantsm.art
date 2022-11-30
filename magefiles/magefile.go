//go:build mage
// +build mage

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl = "https://www.aspca.org/pet-care/animal-poison-control/%s-plant-list"
)

func getDocumentFromUrl(url string) (*goquery.Document, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var backoff time.Duration

	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		response, err := client.Get(url)
		if err != nil {
			break
		}

		switch response.StatusCode {
		case 429:
			if attempt >= maxAttempts {
				log.Printf("could not get url %s after %d attempts; skipping ...", url, attempt)
				return goquery.NewDocumentFromReader(response.Body)
			}

			backoff = time.Duration(attempt) * time.Second
			log.Printf("got rate-limited on url %s; waiting another %d seconds", url, (backoff / time.Second))
			time.Sleep(backoff)
		case 200:
			return goquery.NewDocumentFromReader(response.Body)
		case 404:
			log.Printf("url %s could not be found; skipping ...", url)
			return goquery.NewDocumentFromReader(response.Body)
		}

		defer response.Body.Close()
	}

	return nil, errors.New("this should not happen, but here we are")
}

type Plant struct {
	Id             string
	Name           string
	Common         []string
	Images         []string
	Classification struct {
		Kingdom string
		Clade   string
		Order   string
		Family  string
		Genus   string
		Species string
	}
}

type Data struct {
	PlantsByCat []string
	PlantsByDog []string
	Plants      []Plant
}

func ImportCats() error {
	document, err := getDocumentFromUrl(fmt.Sprintf(baseUrl, "cats"))
	if err != nil {
		return err
	}

	document.Find("div.views-row span.field-content").Each(func(i int, s *goquery.Selection) {
		fmt.Print(strings.TrimSpace(s.Text()))
	})

	return nil
}

func Import(ctx context.Context, animal string) error {
	if animal != "dogs" && animal != "cats" {
		return fmt.Errorf("expected \"cats\" or \"dogs\", but got \"%s\" instead", animal)
	}

	return ImportCats()
}

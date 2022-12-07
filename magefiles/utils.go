package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

	return nil, nil
}

func imageSearch(term string) ([]byte, error) {
	url := fmt.Sprintf("https://api.inaturalist.org/v1/search?sources=taxa&q=%s", url.QueryEscape(term))
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var backoff time.Duration
	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		response, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return body, err
		}

		switch response.StatusCode {
		case 429:
			if attempt >= maxAttempts {
				return nil, fmt.Errorf(fmt.Sprintf("could not find %s after %d attempts; skipping ...", term, attempt))
			}

			backoff = time.Duration(attempt) * time.Second
			log.Printf("got rate-limited on term %s; waiting another %d seconds", term, (backoff / time.Second))
			time.Sleep(backoff)
		case 200:
			return body, err
		default:
			return nil, fmt.Errorf("status code %d", response.StatusCode)
		}
	}

	return nil, nil
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

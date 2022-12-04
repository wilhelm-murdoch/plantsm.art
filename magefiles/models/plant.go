package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
)

const baseUrlRoot = "https://www.aspca.org"

type Plant struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	ScientificName string         `json:"scientific_name"`
	Common         []string       `json:"common"`
	Images         []string       `json:"images"`
	Symptoms       []string       `json:"symptoms"`
	SourceUrl      string         `json:"source_url"`
	Classification Classification `json:"classification"`
}

func NewPlant(animal string, s *goquery.Selection) *Plant {
	source_url, _ := s.Find("a").First().Attr("href")

	pattern := regexp.MustCompile(`^(.+?)\s\((.*?)\).+Scientific\sNames:\s?(.+?)?\s\|\sFamily:(.+?)?$`)
	matches := pattern.FindStringSubmatch(s.Text())

	var common_names []string
	for _, name := range strings.Split(matches[2], ",") {
		if name != "" {
			common_names = append(common_names, strings.TrimSpace(name))
		}
	}

	return &Plant{
		Id:             slug.Make(matches[1]),
		Name:           strings.TrimSpace(matches[1]),
		ScientificName: strings.TrimSpace(matches[3]),
		Common:         common_names,
		SourceUrl:      fmt.Sprintf("%s%s", baseUrlRoot, source_url),
		Classification: Classification{
			Family: strings.TrimSpace(matches[4]),
		},
	}
}

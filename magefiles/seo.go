package main

import (
	"context"
	"fmt"
	"html"
	"html/template"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
)

type Seo mg.Namespace

type Sitemap struct {
	UrlSet []SitemapUrl
}

type SitemapUrl struct {
	Loc     string
	LastMod string
}

func NewSitemap() *Sitemap {
	return &Sitemap{}
}

func (s *Sitemap) AddUrl(url string, lastMod string) {
	s.UrlSet = append(s.UrlSet, SitemapUrl{url, lastMod})
}

func (s Sitemap) UpsertRobots(toPath string) error {
	var (
		robotsPath = fmt.Sprintf("%s/robots.txt", toPath)
		directive  = fmt.Sprintf("Sitemap: %s/sitemap.xml\n", "https://plantsm.art")
	)

	file, err := os.OpenFile(robotsPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	data := make([]byte, stat.Size())
	file.Read(data)
	if err != nil {
		return err
	}

	pattern := regexp.MustCompile(`(?m)^Sitemap:\s+[(http(s)?):\/\/(www\.)?[a-zA-Z0-9-]{2,256}\.[a-z]{2,6}\/[a-zA-Z0-9-_]+.xml`)
	if len(pattern.FindAllString(string(data), -1)) > 0 {
		return nil
	}

	if _, err := file.Write([]byte(directive)); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func (s Sitemap) Save(path string) error {
	sitemap, err := template.New("").Parse(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">{{ range . }}
	<url>
		<loc>{{ .Loc }}</loc>{{ if .LastMod }}
		<lastmod>{{ .LastMod }}</lastmod>{{ end }}
	</url>{{ end }}
</urlset>`)

	if err != nil {
		return err
	}

	var buffer strings.Builder
	if sitemap.Execute(&buffer, s.UrlSet); err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%s/sitemap.xml", path))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(html.UnescapeString(buffer.String()))
	if err != nil {
		return err
	}

	return nil
}

func (Seo) Write(ctx context.Context, sourcePath, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	sitemap := NewSitemap()
	currentTime := time.Now()
	currentTimeFormatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	for _, plant := range plants.Data {
		sitemap.AddUrl(fmt.Sprintf("https://plantsm.art/plant/%s/", plant.Pid), currentTimeFormatted)
	}

	sitemap.AddUrl("https://plantsm.art/", currentTimeFormatted)
	sitemap.AddUrl("https://plantsm.art/why", currentTimeFormatted)
	sitemap.AddUrl("https://plantsm.art/contribute", currentTimeFormatted)
	sitemap.AddUrl("https://plantsm.art/api", currentTimeFormatted)
	sitemap.AddUrl("https://plantsm.art/updates", currentTimeFormatted)

	sitemap.Save(toPath)
	if err := sitemap.UpsertRobots(toPath); err != nil {
		return err
	}

	return nil
}

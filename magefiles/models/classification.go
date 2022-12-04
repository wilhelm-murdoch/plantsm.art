package models

type Classification struct {
	Kingdom string   `json:"kingdom"`
	Clades  []string `json:"clades"`
	Order   string   `json:"order"`
	Family  string   `json:"family"`
	Genus   string   `json:"genus"`
	Species string   `json:"species"`
}

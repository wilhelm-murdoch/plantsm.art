package models

import "github.com/wilhelm-murdoch/go-collection"

type Data struct {
	PlantsByCat      *collection.Collection[string] `json:"plants_by_cat"`
	PlantsByDog      *collection.Collection[string] `json:"plants_by_dog"`
	PlantsBySymptoms [][]string                     `json:"plants_by_symptoms"`
	Plants           *collection.Collection[*Plant] `json:"plants"`
}

type UnmarshalData struct {
	PlantsByCat      []string   `json:"plants_by_cat"`
	PlantsByDog      []string   `json:"plants_by_dog"`
	PlantsBySymptoms [][]string `json:"plants_by_symptoms"`
	Plants           []Plant    `json:"plants"`
}

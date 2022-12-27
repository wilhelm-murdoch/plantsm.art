package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type Json mg.Namespace

func (Json) Pages(ctx context.Context, sourcePath, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	plants, err := unmarshalPlantsFromSource(sourcePath)
	if err != nil {
		return err
	}

	for _, plant := range plants.Data {
		file, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", toPath, plant.Pid), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(plant); err != nil {
			return err
		}

		fmt.Printf("Wrote: %s/%s.json\n", toPath, plant.Pid)
	}

	return nil
}

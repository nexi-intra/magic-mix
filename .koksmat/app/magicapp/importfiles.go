package magicapp

import (
	"fmt"
	"os"

	"github.com/magicbutton/magic-mix/services/endpoints/importdata"
	"github.com/magicbutton/magic-mix/services/models/importdatamodel"
)

func ImportFiles(filepath string) (*importdatamodel.ImportData, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {

		return nil, fmt.Errorf("could not read file: %w", err)
	}

	OpenDatabase()
	text := string(data)

	importRecord := importdatamodel.ImportData{
		Name:        filepath,
		Description: "Imported file",

		Data: []byte(text),
	}
	return importdata.ImportDataCreate(importRecord)

}

package collect

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/magicbutton/magic-mix/services/endpoints/importdata"
	"github.com/magicbutton/magic-mix/services/models/importdatamodel"
)

// EnumerateFiles walks through the given directory path and calls the callback function with the path of each file found.
func EnumerateFiles(root string, callback func(path string) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			return callback(path)
		}
		return nil
	})
}

func UploadBatch(rootPath string) {
	callback := func(filePath string) error {
		log.Println("Uploading", filePath)
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal("could not read file: %w", err)
		}

		_, importErr := importdata.ImportDataCreate(importdatamodel.ImportData{
			Name:        filePath,
			Description: "Imported file",

			Data: string(data),
		})
		return importErr
	}

	if err := EnumerateFiles(rootPath, callback); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

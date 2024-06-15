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
	inserter := func(filePath string, chunk *string) error {
		log.Println("Uploading", filePath, len(*chunk), "bytes")
		//return nil
		_, importErr := importdata.ImportDataCreate(importdatamodel.ImportData{
			Name:        filePath,
			Description: "Imported file",

			Data: chunk,
		})
		return importErr
	}

	callback := func(filePath string) error {
		//log.Println("Uploading", filePath)
		return LoadJSON(filePath, inserter)
	}

	if err := EnumerateFiles(rootPath, callback); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

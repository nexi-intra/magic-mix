package collect

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CombineJsonFiles(dir string, prefix string, deleteAfterCombine bool) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	var combinedData []map[string]interface{}
	var filesToDelete []string

	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) && strings.HasSuffix(file.Name(), ".json") {
			filePath := filepath.Join(dir, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
				continue
			}

			var jsonData map[string]interface{}
			if err := json.Unmarshal(data, &jsonData); err != nil {
				fmt.Printf("Error unmarshalling JSON from file %s: %v\n", file.Name(), err)
				continue
			}

			fileData := map[string]interface{}{
				"file": filePath,
				"data": jsonData,
			}

			combinedData = append(combinedData, fileData)
			filesToDelete = append(filesToDelete, filePath)
		}
	}

	outputFileName := filepath.Join(dir, prefix+".json")
	outputData, err := json.MarshalIndent(combinedData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling combined JSON data: %w", err)
	}

	if err := os.WriteFile(outputFileName, outputData, 0644); err != nil {
		return fmt.Errorf("error writing to output file %s: %w", outputFileName, err)
	}

	fmt.Printf("Combined JSON data has been written to %s\n", outputFileName)

	if deleteAfterCombine {
		for _, filePath := range filesToDelete {
			if err := os.Remove(filePath); err != nil {
				fmt.Printf("Error deleting file %s: %v\n", filePath, err)
			} else {
				// fmt.Printf("Deleted file %s\n", filePath)
			}
		}
	}

	return nil
}

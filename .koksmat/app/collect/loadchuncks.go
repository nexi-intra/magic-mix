package collect

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	maxInsertSize = 100000000 // Adjust based on your PostgreSQL settings
	chunkSize     = maxInsertSize / 2
)

// ChunkCallback defines the type for the callback function
type ChunkCallback func(filePath string, chunk *string) error

// LoadJSON loads a JSON file and processes its content using the provided callback function.
// If the content is a JSON object, it processes the entire content as one chunk.
// If the content is a JSON array, it processes the content in chunks.
func LoadJSON(filePath string, callback ChunkCallback) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data []byte
	data, err = io.ReadAll(file)
	if err != nil {
		return err
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	switch v := jsonData.(type) {
	case map[string]interface{}:
		// It's a JSON object
		chunkData, err := json.Marshal(v)
		if err != nil {
			return err
		}
		chunk := string(chunkData)
		return callback(filePath, &chunk)

	case []interface{}:
		// It's a JSON array
		jsonArray := make([]map[string]interface{}, len(v))
		for i, item := range v {
			jsonArray[i] = item.(map[string]interface{})
		}

		var chunk []map[string]interface{}
		currentSize := 0

		for _, item := range jsonArray {
			itemData, _ := json.Marshal(item)
			itemSize := len(itemData)

			if currentSize+itemSize > chunkSize {
				chunkData, err := json.Marshal(chunk)
				if err != nil {
					return err
				}
				chunkString := string(chunkData)
				if err := callback(filePath, &chunkString); err != nil {
					return err
				}
				chunk = []map[string]interface{}{}
				currentSize = 0
			}

			chunk = append(chunk, item)
			currentSize += itemSize
		}

		if len(chunk) > 0 {
			chunkData, err := json.Marshal(chunk)
			if err != nil {
				return err
			}
			chunkString := string(chunkData)
			if err := callback(filePath, &chunkString); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unsupported JSON format")
	}

	return nil
}

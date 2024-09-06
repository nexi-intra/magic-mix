package officegraph

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

func FilterJSONArray(data []map[string]interface{}, queryStr string) []map[string]interface{} {
	// Convert []map[string]interface{} to []interface{} to match gojq's expected input type
	var dataInterface []interface{}
	for _, item := range data {
		dataInterface = append(dataInterface, item)
	}

	query, err := gojq.Parse(queryStr)
	if err != nil {
		panic(err)
	}

	iter := query.Run(dataInterface)

	var result []map[string]interface{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			fmt.Println("Error:", err)
			return nil
		}

		// Type assertion to convert interface{} back to map[string]interface{}
		if mapResult, ok := v.(map[string]interface{}); ok {
			result = append(result, mapResult)
		} else {
			fmt.Println("Unexpected type:", v)
		}
	}

	return result
}
func JogqJSONfilterSample() {
	jsonArray := `[
        {"name": "x", "age": 20},
        {"name": "y", "age": 30},
        {"name": "z", "age": 40}
    ]`
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonArray), &data)
	if err != nil {
		panic(err)
	}
	filtered := FilterJSONArray(data, "[.[] | select(.name != \"x\" and .name != \"y\")]")

	fmt.Printf("Filtered Result: %+v\n", filtered)
}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3
package app

// noma2
import (
	"fmt"

	"github.com/magicbutton/magic-mix/query"
)

func Dictionary(args []string) (*SelectResponse, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Expected 1 arguments")
	}

	area := args[0]
	database := args[1]

	result := &SelectResponse{}
	switch area {
	case "storedprocedures":
		r, err := query.GetStoredProcedures(database)
		if err != nil {
			return nil, fmt.Errorf("Failed to get stored procedures")
		}
		result.Result = *r
		// macd.1
	// case "storedprocedure":
	// 	// macd.2
	default:
		return nil, fmt.Errorf("Unknown area")
	}

	return result, nil
}

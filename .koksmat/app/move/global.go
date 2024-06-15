package move

import (
	"fmt"
	"log"

	"github.com/magicbutton/magic-mix/services/endpoints/connection"
)

func GetConnectionString(name string) (*string, error) {
	connectionRecord, err := connection.ConnectionSearch(name + "%")
	if err != nil {
		log.Println("failed to get connection:", err)
		return nil, err
	}
	if connectionRecord.TotalItems != 1 {
		log.Println("failed to get connection: ", err)

		// go a new go error
		return nil, fmt.Errorf("failed to get connection: %v", err)

	}

	return &connectionRecord.Items[0].Connectionstring, nil
}

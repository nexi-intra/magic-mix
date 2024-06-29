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
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/magicbutton/magic-mix/services/endpoints/connection"
	"github.com/magicbutton/magic-mix/utils"
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

func Select2(args []string) (*SelectResponse, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Expected 1 arguments")
	}
	dbName := args[0]

	conn, err := GetConnectionString(dbName)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", *conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := args[1]

	jsonsql := fmt.Sprintf(`
	SELECT json_agg(json_data) AS result
	FROM (
		%s
	) AS json_data;
		
	`, sql)
	ctx := context.Background()

	rows, err := db.QueryContext(ctx, jsonsql)
	if err != nil {
		return nil, err
	}
	result := []SelectResponse{}
	err = utils.Db.ScanRows(ctx, rows, &result)
	if len(result) != 1 {
		return nil, fmt.Errorf("Unknown result")
	}

	return &result[0], nil
}

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
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/magicbutton/magic-mix/utils"
	"github.com/spf13/viper"
)

type Result struct {
	Data string `json:"data"`
}

func call(connectionString string, procName string, who string, payload json.RawMessage) (string, error) {
	//

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		return "", pingErr

	}
	var payloadStr string = string(payload)

	sqlStatement := fmt.Sprintf(`
	    SELECT * from  proc.%s('%s', '%s');
	
		`, procName, who, payloadStr)
	//sqlStatement := "call proc.simple_procedure()"
	log.Println(sqlStatement)
	//x, err := db.Exec(sqlStatement)
	var pOutput string
	//query := fmt.Sprintf(`CALL proc.%s($1, $2);SELECT p_output;`, procName)
	//rows, err := db.Query(query, who, payloadStr)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call stored procedure: %v\n", err)
		return "", err
	}
	defer rows.Close()

	// Assuming the output is in the first column of the first row
	if rows.Next() {
		if err := rows.Scan(&pOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to retrieve output: %v\n", err)
			return "", err
		}
	}

	// Wrap the result into a JSON object

	// Scan the JSON into a string variable
	return fmt.Sprintf(`{"OK":true,"ID":%s}`, pOutput), nil

}

func Process(args []string) (*SelectResponse, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("Expected arguments")
	}
	jwt := args[1]
	if jwt == "" {
		return nil, fmt.Errorf("Expected JWT")
	}
	claims, err := utils.DecodeAndValidateMicrosoftJWT(jwt)
	if err != nil {
		return nil, err
	}
	connectionString := viper.GetString("POSTGRES_DB")
	upn := claims["upn"].(string)
	log.Println("Process", args[0], upn, json.RawMessage(args[2]))
	rows, err := call(connectionString, args[0], upn, json.RawMessage(args[2]))
	if err != nil {
		return nil, err
	}

	result := SelectResponse{
		Result: json.RawMessage(rows),
	}

	return &result, nil
}

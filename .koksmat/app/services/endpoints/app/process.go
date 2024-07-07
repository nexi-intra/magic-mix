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
		DO $$
	DECLARE
	    p_id INTEGER;
	    p_actor_name VARCHAR := '%s';
	    p_params JSONB := '%s';
	BEGIN
	    -- Call the procedure
	    CALL proc.%s(p_actor_name, p_params, p_id);

	    -- Output the resulting ID
	    RAISE NOTICE ' ID: %s', p_id;
	END $$;

		`, who, payloadStr, procName, "%")
	//sqlStatement := "call proc.simple_procedure()"
	log.Println(sqlStatement)
	x, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call stored procedure: %v\n", err)
		return "", err
	}

	if x != nil {
		log.Println("Executed successfully", x)
	}

	// Wrap the result into a JSON object

	// Scan the JSON into a string variable
	return `{"OK":true}`, nil

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

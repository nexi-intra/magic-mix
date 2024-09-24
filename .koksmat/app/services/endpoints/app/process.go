package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/magicbutton/magic-mix/utils"
	"github.com/spf13/viper"
)

type Result struct {
	Data string `json:"data"`
}

func callWithNotification(connectionString string, procName string, who string, payload json.RawMessage) (string, error) {

	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}
	db := stdlib.OpenDB(*config)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		return "", pingErr
	}

	sqlStatement := `SELECT * FROM proc.update_sqlquery($1::text, $2::jsonb, $3::jsonb)`

	rows, err := db.Query(sqlStatement, who, payload, nil)
	if err != nil {
		log.Printf("Failed to call stored procedure: %v\n", err)
		log.Printf("Executing query: %s with params: %s, %s", sqlStatement, who, payload)
		return "", err
	}
	defer rows.Close()

	var pOutput string
	if rows.Next() {
		if err := rows.Scan(&pOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to retrieve output: %v\n", err)
			return "", err
		}
	}

	return pOutput, nil
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
	//log.Println("Process", args[0], upn, json.RawMessage(args[2]))
	rows, err := callWithNotification(connectionString, args[0], upn, json.RawMessage(args[2]))
	if err != nil {
		return nil, err
	}

	result := SelectResponse{
		Result: json.RawMessage(rows),
	}

	return &result, nil
}

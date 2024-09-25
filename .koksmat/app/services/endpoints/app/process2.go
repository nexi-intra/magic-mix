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
	"encoding/json"
	"fmt"
	"log"

	"github.com/magicbutton/magic-mix/utils"
	"github.com/nats-io/nats.go"
)

func Process2(args []string, nc *nats.Conn) (*SelectResponse, error) {
	if len(args) < 4 {
		return nil, fmt.Errorf("Expected arguments")
	}
	jwt := args[2]
	if jwt == "" {
		return nil, fmt.Errorf("Expected JWT")
	}
	claims, err := utils.DecodeAndValidateMicrosoftJWT(jwt)
	if err != nil {
		return nil, err
	}

	dbName := args[0]

	conn, err := GetConnectionString(dbName)
	if err != nil {
		return nil, err
	}

	upn := claims["upn"].(string)
	log.Println("calling callWithNotification", args[1], upn)
	rows, err := callWithNotification(*conn, args[1], upn, json.RawMessage(args[3]), nc, dbName)
	if err != nil {
		return nil, err
	}

	result := SelectResponse{
		Result: json.RawMessage(rows),
	}

	return &result, nil
}

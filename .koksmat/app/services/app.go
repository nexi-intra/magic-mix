/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
// macd.1
package services

import (
	"encoding/json"

	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-mix/services/endpoints/app"
	. "github.com/magicbutton/magic-mix/utils"
)

func HandleAppRequests(req micro.Request) {

	rawRequest := string(req.Data())
	if rawRequest == "ping" {
		req.Respond([]byte("pong"))
		return

	}

	var payload ServiceRequest
	_ = json.Unmarshal([]byte(req.Data()), &payload)
	if len(payload.Args) < 1 {
		ServiceResponseError(req, "missing command")
		return

	}
	switch payload.Args[0] {

	// macd.2

	case "select":
		ProcessAppRequest(req, app.Select)
	case "dictionary":
		ProcessAppRequest(req, app.Dictionary)
	default:
		ServiceResponseError(req, "Unknown command")
	}
}

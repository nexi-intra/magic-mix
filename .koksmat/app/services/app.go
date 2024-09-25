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

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-mix/services/endpoints/app"
	. "github.com/magicbutton/magic-mix/utils"
)

func HandleAppRequests(req micro.Request, nc *nats.Conn) {

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
		ProcessAppRequest(req, nc, app.Select)
	case "query":
		ProcessAppRequest(req, nc, app.Select2)
	case "dictionary":
		ProcessAppRequest(req, nc, app.Dictionary)
	case "process":
		ProcessAppRequest(req, nc, app.Process)
	case "execute":
		ProcessAppRequest(req, nc, app.Process2)
	default:
		ServiceResponseError(req, "Unknown command")
	}
}

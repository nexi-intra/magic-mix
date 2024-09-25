/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
package magicapp

import (
	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-mix/services"
)

func RegisterServiceEndpoints(root micro.Group, nc *nats.Conn) {
	//root.AddEndpoint("app", micro.HandlerFunc(services.HandleAppRequests))
	root.AddEndpoint("app", micro.HandlerFunc(func(req micro.Request) {
		services.HandleAppRequests(req, nc)
	}))

	root.AddEndpoint("connection", micro.HandlerFunc(func(req micro.Request) {
		services.HandleConnectionRequests(req, nc)
	}))
	root.AddEndpoint("importdata", micro.HandlerFunc(func(req micro.Request) {
		services.HandleImportDataRequests(req, nc)
	}))

}

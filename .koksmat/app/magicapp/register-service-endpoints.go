/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
package magicapp

import (
	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-mix/services"
)

func RegisterServiceEndpoints(root micro.Group) {
	root.AddEndpoint("app", micro.HandlerFunc(services.HandleAppRequests))
	root.AddEndpoint("connection", micro.HandlerFunc(services.HandleConnectionRequests))
	root.AddEndpoint("importdata", micro.HandlerFunc(services.HandleImportDataRequests))
}

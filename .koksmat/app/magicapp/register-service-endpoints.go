/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-mix/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
	root.AddEndpoint("app", micro.HandlerFunc(services.HandleAppRequests))

	root.AddEndpoint("connection", micro.HandlerFunc(services.HandleConnectionRequests))
	root.AddEndpoint("transformer", micro.HandlerFunc(services.HandleTransformerRequests))
	root.AddEndpoint("dataset", micro.HandlerFunc(services.HandleDatasetRequests))
	root.AddEndpoint("column", micro.HandlerFunc(services.HandleColumnRequests))
	root.AddEndpoint("mapper", micro.HandlerFunc(services.HandleMapperRequests))
	root.AddEndpoint("transformation", micro.HandlerFunc(services.HandleTransformationRequests))
	root.AddEndpoint("processlog", micro.HandlerFunc(services.HandleProcessLogRequests))
	root.AddEndpoint("importdata", micro.HandlerFunc(services.HandleImportDataRequests))
}

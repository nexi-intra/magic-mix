/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
// macd.1
package services
import (
	"encoding/json"
    "fmt"
	"log"
    "github.com/magicbutton/magic-mix/services/endpoints/importdata"
    "github.com/magicbutton/magic-mix/services/models/importdatamodel"

	. "github.com/magicbutton/magic-mix/utils"
	"github.com/nats-io/nats.go/micro"
)

func HandleImportDataRequests(req micro.Request) {

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
case "read":
if (len(payload.Args) < 2) {
    log.Println("Expected 2 arguments, got %d", len(payload.Args))
    ServiceResponseError(req, "Expected 1 arguments")
    return
}


    
    result,err := importdata.ImportDataRead(payload.Args[1])
    if (err != nil) {
        log.Println("Error", err)
        ServiceResponseError(req, fmt.Sprintf("Error calling ImportDataRead: %s", err))


        return
    }

    ServiceResponse(req, result)

// macd.2
case "create":
if (len(payload.Args) < 2) {
    log.Println("Expected 2 arguments, got %d", len(payload.Args))
    ServiceResponseError(req, "Expected 1 arguments")
    return
}


                // transformer v1
            object := importdatamodel.ImportData{}
            body := ""

            json.Unmarshal([]byte(payload.Args[1]), &body)
            err := json.Unmarshal([]byte(body), &object)
    
            if err != nil {
                log.Println("Error", err)
                ServiceResponseError(req, "Error unmarshalling importdata")
                return
            }
                     
    result,err := importdata.ImportDataCreate(object)
    if (err != nil) {
        log.Println("Error", err)
        ServiceResponseError(req, fmt.Sprintf("Error calling ImportDataCreate: %s", err))


        return
    }

    ServiceResponse(req, result)

// macd.2
case "update":
if (len(payload.Args) < 2) {
    log.Println("Expected 2 arguments, got %d", len(payload.Args))
    ServiceResponseError(req, "Expected 1 arguments")
    return
}


                // transformer v1
            object := importdatamodel.ImportData{}
            body := ""

            json.Unmarshal([]byte(payload.Args[1]), &body)
            err := json.Unmarshal([]byte(body), &object)
    
            if err != nil {
                log.Println("Error", err)
                ServiceResponseError(req, "Error unmarshalling importdata")
                return
            }
                     
    result,err := importdata.ImportDataUpdate(object)
    if (err != nil) {
        log.Println("Error", err)
        ServiceResponseError(req, fmt.Sprintf("Error calling ImportDataUpdate: %s", err))


        return
    }

    ServiceResponse(req, result)

// macd.2
case "delete":
if (len(payload.Args) < 2) {
    log.Println("Expected 2 arguments, got %d", len(payload.Args))
    ServiceResponseError(req, "Expected 1 arguments")
    return
}


            err :=  importdata.ImportDataDelete(payload.Args[1])
            if (err != nil) {
                log.Println("Error", err)
                ServiceResponseError(req, fmt.Sprintf("Error calling ImportDataDelete: %s", err))


                return
            }
            ServiceResponse(req, "")

// macd.2
case "search":
if (len(payload.Args) < 2) {
    log.Println("Expected 2 arguments, got %d", len(payload.Args))
    ServiceResponseError(req, "Expected 1 arguments")
    return
}


    
    result,err := importdata.ImportDataSearch(payload.Args[1])
    if (err != nil) {
        log.Println("Error", err)
        ServiceResponseError(req, fmt.Sprintf("Error calling ImportDataSearch: %s", err))


        return
    }

    ServiceResponse(req, result)

default:
ServiceResponseError(req, "Unknown command")
}
}

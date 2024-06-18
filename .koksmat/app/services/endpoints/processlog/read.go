/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package processlog

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/processlogmodel"
)

func ProcessLogRead(arg0 string) (*processlogmodel.ProcessLog, error) {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling ProcessLogread")

	return applogic.Read[database.ProcessLog, processlogmodel.ProcessLog](id, applogic.MapProcessLogOutgoing)

}

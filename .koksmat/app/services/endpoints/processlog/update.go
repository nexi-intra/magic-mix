/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package processlog

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/processlogmodel"
)

func ProcessLogUpdate(item processlogmodel.ProcessLog) (*processlogmodel.ProcessLog, error) {
	log.Println("Calling ProcessLogupdate")

	return applogic.Update[database.ProcessLog, processlogmodel.ProcessLog](item.ID, item, applogic.MapProcessLogIncoming, applogic.MapProcessLogOutgoing)

}

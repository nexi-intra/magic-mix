/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package schedule

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/schedulemodel"
)

func ScheduleRead(arg0 string) (*schedulemodel.Schedule, error) {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling Scheduleread")

	return applogic.Read[database.Schedule, schedulemodel.Schedule](id, applogic.MapScheduleOutgoing)

}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package schedule

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/schedulemodel"

)

func ScheduleUpdate(item schedulemodel.Schedule) (*schedulemodel.Schedule, error) {
    log.Println("Calling Scheduleupdate")

    return applogic.Update[database.Schedule, schedulemodel.Schedule](item.ID,item, applogic.MapScheduleIncoming, applogic.MapScheduleOutgoing)

}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package schedule

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/schedulemodel"
    . "github.com/magicbutton/magic-mix/utils"
)

func ScheduleSearch(query string) (*Page[schedulemodel.Schedule], error) {
    log.Println("Calling Schedulesearch")

    return applogic.Search[database.Schedule, schedulemodel.Schedule]("searchindex", query, applogic.MapScheduleOutgoing)

}

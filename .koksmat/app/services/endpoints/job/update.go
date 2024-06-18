/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package job

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/jobmodel"

)

func JobUpdate(item jobmodel.Job) (*jobmodel.Job, error) {
    log.Println("Calling Jobupdate")

    return applogic.Update[database.Job, jobmodel.Job](item.ID,item, applogic.MapJobIncoming, applogic.MapJobOutgoing)

}

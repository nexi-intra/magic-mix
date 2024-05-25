/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package processlog

import (
    "log"
   
    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/processlogmodel"

)

func ProcessLogCreate(item processlogmodel.ProcessLog) (*processlogmodel.ProcessLog, error) {
    log.Println("Calling ProcessLogcreate")

    return applogic.Create[database.ProcessLog, processlogmodel.ProcessLog](item, applogic.MapProcessLogIncoming, applogic.MapProcessLogOutgoing)

}

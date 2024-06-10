/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package importdata

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/importdatamodel"

)

func ImportDataRead(arg0 string) (*importdatamodel.ImportData, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ImportDataread")

    return applogic.Read[database.ImportData, importdatamodel.ImportData](id, applogic.MapImportDataOutgoing)

}

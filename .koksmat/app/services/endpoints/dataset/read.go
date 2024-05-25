/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package dataset

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/datasetmodel"

)

func DatasetRead(arg0 string) (*datasetmodel.Dataset, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Datasetread")

    return applogic.Read[database.Dataset, datasetmodel.Dataset](id, applogic.MapDatasetOutgoing)

}

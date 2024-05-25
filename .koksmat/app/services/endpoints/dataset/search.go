/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package dataset

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/datasetmodel"
    . "github.com/magicbutton/magic-mix/utils"
)

func DatasetSearch(query string) (*Page[datasetmodel.Dataset], error) {
    log.Println("Calling Datasetsearch")

    return applogic.Search[database.Dataset, datasetmodel.Dataset]("searchindex", query, applogic.MapDatasetOutgoing)

}

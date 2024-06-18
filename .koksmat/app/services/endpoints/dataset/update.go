/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package dataset

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/datasetmodel"
)

func DatasetUpdate(item datasetmodel.Dataset) (*datasetmodel.Dataset, error) {
	log.Println("Calling Datasetupdate")

	return applogic.Update[database.Dataset, datasetmodel.Dataset](item.ID, item, applogic.MapDatasetIncoming, applogic.MapDatasetOutgoing)

}

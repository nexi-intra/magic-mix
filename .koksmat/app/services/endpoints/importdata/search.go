/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importdata

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/importdatamodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func ImportDataSearch(query string) (*Page[importdatamodel.ImportData], error) {
	log.Println("Calling ImportDatasearch")

	return applogic.Search[database.ImportData, importdatamodel.ImportData]("searchindex", query, applogic.MapImportDataOutgoing)

}

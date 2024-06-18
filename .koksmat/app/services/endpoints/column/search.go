/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package column

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/columnmodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func ColumnSearch(query string) (*Page[columnmodel.Column], error) {
	log.Println("Calling Columnsearch")

	return applogic.Search[database.Column, columnmodel.Column]("searchindex", query, applogic.MapColumnOutgoing)

}

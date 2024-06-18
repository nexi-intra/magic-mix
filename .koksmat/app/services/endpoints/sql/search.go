/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package sql

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/sqlmodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func SQLSearch(query string) (*Page[sqlmodel.SQL], error) {
	log.Println("Calling SQLsearch")

	return applogic.Search[database.SQL, sqlmodel.SQL]("searchindex", query, applogic.MapSQLOutgoing)

}

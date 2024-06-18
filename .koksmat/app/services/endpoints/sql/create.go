/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package sql

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/sqlmodel"
)

func SQLCreate(item sqlmodel.SQL) (*sqlmodel.SQL, error) {
	log.Println("Calling SQLcreate")

	return applogic.Create[database.SQL, sqlmodel.SQL](item, applogic.MapSQLIncoming, applogic.MapSQLOutgoing)

}

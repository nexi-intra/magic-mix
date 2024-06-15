/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package sql

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/sqlmodel"

)

func SQLUpdate(item sqlmodel.SQL) (*sqlmodel.SQL, error) {
    log.Println("Calling SQLupdate")

    return applogic.Update[database.SQL, sqlmodel.SQL](item.ID,item, applogic.MapSQLIncoming, applogic.MapSQLOutgoing)

}

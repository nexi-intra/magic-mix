/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package column

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/columnmodel"

)

func ColumnUpdate(item columnmodel.Column) (*columnmodel.Column, error) {
    log.Println("Calling Columnupdate")

    return applogic.Update[database.Column, columnmodel.Column](item.ID,item, applogic.MapColumnIncoming, applogic.MapColumnOutgoing)

}

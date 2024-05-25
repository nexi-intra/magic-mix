/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package connection

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/connectionmodel"

)

func ConnectionUpdate(item connectionmodel.Connection) (*connectionmodel.Connection, error) {
    log.Println("Calling Connectionupdate")

    return applogic.Update[database.Connection, connectionmodel.Connection](item.ID,item, applogic.MapConnectionIncoming, applogic.MapConnectionOutgoing)

}

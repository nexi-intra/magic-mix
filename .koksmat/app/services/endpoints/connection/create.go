/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package connection

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/connectionmodel"
)

func ConnectionCreate(item connectionmodel.Connection) (*connectionmodel.Connection, error) {
	log.Println("Calling Connectioncreate")

	return applogic.Create[database.Connection, connectionmodel.Connection](item, applogic.MapConnectionIncoming, applogic.MapConnectionOutgoing)

}

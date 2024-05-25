/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package connection

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/connectionmodel"
    . "github.com/magicbutton/magic-mix/utils"
)

func ConnectionSearch(query string) (*Page[connectionmodel.Connection], error) {
    log.Println("Calling Connectionsearch")

    return applogic.Search[database.Connection, connectionmodel.Connection]("searchindex", query, applogic.MapConnectionOutgoing)

}

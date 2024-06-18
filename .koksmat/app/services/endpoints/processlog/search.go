/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package processlog

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/processlogmodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func ProcessLogSearch(query string) (*Page[processlogmodel.ProcessLog], error) {
	log.Println("Calling ProcessLogsearch")

	return applogic.Search[database.ProcessLog, processlogmodel.ProcessLog]("searchindex", query, applogic.MapProcessLogOutgoing)

}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package mapper

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/mappermodel"
)

func MapperRead(arg0 string) (*mappermodel.Mapper, error) {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling Mapperread")

	return applogic.Read[database.Mapper, mappermodel.Mapper](id, applogic.MapMapperOutgoing)

}

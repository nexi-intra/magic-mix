/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package mapper

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/mappermodel"
)

func MapperUpdate(item mappermodel.Mapper) (*mappermodel.Mapper, error) {
	log.Println("Calling Mapperupdate")

	return applogic.Update[database.Mapper, mappermodel.Mapper](item.ID, item, applogic.MapMapperIncoming, applogic.MapMapperOutgoing)

}

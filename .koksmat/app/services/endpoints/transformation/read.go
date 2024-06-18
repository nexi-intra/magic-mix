/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package transformation

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformationmodel"
)

func TransformationRead(arg0 string) (*transformationmodel.Transformation, error) {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling Transformationread")

	return applogic.Read[database.Transformation, transformationmodel.Transformation](id, applogic.MapTransformationOutgoing)

}

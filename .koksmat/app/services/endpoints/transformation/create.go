/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package transformation

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformationmodel"
)

func TransformationCreate(item transformationmodel.Transformation) (*transformationmodel.Transformation, error) {
	log.Println("Calling Transformationcreate")

	return applogic.Create[database.Transformation, transformationmodel.Transformation](item, applogic.MapTransformationIncoming, applogic.MapTransformationOutgoing)

}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package transformation

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/transformationmodel"

)

func TransformationUpdate(item transformationmodel.Transformation) (*transformationmodel.Transformation, error) {
    log.Println("Calling Transformationupdate")

    return applogic.Update[database.Transformation, transformationmodel.Transformation](item.ID,item, applogic.MapTransformationIncoming, applogic.MapTransformationOutgoing)

}

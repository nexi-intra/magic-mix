/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package transformer

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/transformermodel"

)

func TransformerUpdate(item transformermodel.Transformer) (*transformermodel.Transformer, error) {
    log.Println("Calling Transformerupdate")

    return applogic.Update[database.Transformer, transformermodel.Transformer](item.ID,item, applogic.MapTransformerIncoming, applogic.MapTransformerOutgoing)

}

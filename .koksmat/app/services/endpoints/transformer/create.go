/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package transformer

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformermodel"
)

func TransformerCreate(item transformermodel.Transformer) (*transformermodel.Transformer, error) {
	log.Println("Calling Transformercreate")

	return applogic.Create[database.Transformer, transformermodel.Transformer](item, applogic.MapTransformerIncoming, applogic.MapTransformerOutgoing)

}

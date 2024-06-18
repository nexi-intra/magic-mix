/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package transformation

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformationmodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func TransformationSearch(query string) (*Page[transformationmodel.Transformation], error) {
	log.Println("Calling Transformationsearch")

	return applogic.Search[database.Transformation, transformationmodel.Transformation]("searchindex", query, applogic.MapTransformationOutgoing)

}

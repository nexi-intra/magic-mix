/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package transformer

import (
	"log"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformermodel"
	. "github.com/magicbutton/magic-mix/utils"
)

func TransformerSearch(query string) (*Page[transformermodel.Transformer], error) {
	log.Println("Calling Transformersearch")

	return applogic.Search[database.Transformer, transformermodel.Transformer]("searchindex", query, applogic.MapTransformerOutgoing)

}

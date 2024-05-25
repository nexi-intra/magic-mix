/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package mapper

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/mappermodel"
    . "github.com/magicbutton/magic-mix/utils"
)

func MapperSearch(query string) (*Page[mappermodel.Mapper], error) {
    log.Println("Calling Mappersearch")

    return applogic.Search[database.Mapper, mappermodel.Mapper]("searchindex", query, applogic.MapMapperOutgoing)

}

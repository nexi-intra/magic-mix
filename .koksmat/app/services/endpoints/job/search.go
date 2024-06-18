/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package job

import (
    "log"

    "github.com/magicbutton/magic-mix/applogic"
    "github.com/magicbutton/magic-mix/database"
    "github.com/magicbutton/magic-mix/services/models/jobmodel"
    . "github.com/magicbutton/magic-mix/utils"
)

func JobSearch(query string) (*Page[jobmodel.Job], error) {
    log.Println("Calling Jobsearch")

    return applogic.Search[database.Job, jobmodel.Job]("searchindex", query, applogic.MapJobOutgoing)

}

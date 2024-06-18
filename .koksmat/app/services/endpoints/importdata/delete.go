/*
   File have been automatically created. To prevent the file from getting overwritten
   set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
   ---
   keep: false
   ---
*/
//generator:  noma3.delete.v2
package importdata

import (
	"log"
	"strconv"

	"github.com/magicbutton/magic-mix/applogic"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/importdatamodel"
)

func ImportDataDelete(arg0 string) error {
	id, _ := strconv.Atoi(arg0)
	log.Println("Calling ImportDatadelete")

	return applogic.Delete[database.ImportData, importdatamodel.ImportData](id)

}

/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-mix/database"
	"github.com/magicbutton/magic-mix/services/models/transformationmodel"
   
)


func MapTransformationOutgoing(db database.Transformation) transformationmodel.Transformation {
    return transformationmodel.Transformation{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Input_id : db.Input_id,
                Output_id : db.Output_id,

    }
}

func MapTransformationIncoming(in transformationmodel.Transformation) database.Transformation {
    return database.Transformation{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Input_id : in.Input_id,
                Output_id : in.Output_id,
        Searchindex : in.Name,

    }
}

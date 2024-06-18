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
	"github.com/magicbutton/magic-mix/services/models/schedulemodel"
   
)


func MapScheduleOutgoing(db database.Schedule) schedulemodel.Schedule {
    return schedulemodel.Schedule{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Cron : db.Cron,
                Job_id : db.Job_id,
        Data : db.Data,

    }
}

func MapScheduleIncoming(in schedulemodel.Schedule) database.Schedule {
    return database.Schedule{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Cron : in.Cron,
                Job_id : in.Job_id,
        Data : in.Data,
        Searchindex : in.Name,

    }
}

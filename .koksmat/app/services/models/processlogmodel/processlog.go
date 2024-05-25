/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package processlogmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-mix/database/databasetypes"
)

func UnmarshalProcessLog(data []byte) (ProcessLog, error) {
	var r ProcessLog
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProcessLog) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProcessLog struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Transformation_id int `json:"transformation_id"`
    Status string `json:"status"`
    Message string `json:"message"`

}


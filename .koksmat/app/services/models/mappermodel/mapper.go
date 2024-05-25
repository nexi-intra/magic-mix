/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package mappermodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-mix/database/databasetypes"
)

func UnmarshalMapper(data []byte) (Mapper, error) {
	var r Mapper
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Mapper) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Mapper struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Source_id int `json:"source_id"`
    Transformation_id int `json:"transformation_id"`
    Target_id int `json:"target_id"`

}


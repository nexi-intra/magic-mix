/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package transformationmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-mix/database/databasetypes"
)

func UnmarshalTransformation(data []byte) (Transformation, error) {
	var r Transformation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Transformation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Transformation struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Input_id int `json:"input_id"`
    Output_id int `json:"output_id"`

}


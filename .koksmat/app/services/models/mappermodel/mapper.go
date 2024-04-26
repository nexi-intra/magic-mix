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
    "github.com/magicbutton/magic-mix/database/databasetypes"
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
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Source databasetypes.Reference `json:"source"`
    Transformation databasetypes.Reference `json:"transformation"`
    Target databasetypes.Reference `json:"target"`

}


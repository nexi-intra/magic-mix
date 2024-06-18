/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package routemodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-mix/database/databasetypes"
)

func UnmarshalRoute(data []byte) (Route, error) {
	var r Route
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Route) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Route struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Method string `json:"method"`
    Slug string `json:"slug"`
    Api_id int `json:"api_id"`

}


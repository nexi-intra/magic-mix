/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateGoModel v2
package apimodel

import (
	"encoding/json"
	"time"
)

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type API struct {
	ID          int         `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatedBy   string      `json:"created_by"`
	UpdatedAt   time.Time   `json:"updated_at"`
	UpdatedBy   string      `json:"updated_by"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Method      string      `json:"method"`
	Source      interface{} `json:"source"`
	Schema      interface{} `json:"schema"`
}

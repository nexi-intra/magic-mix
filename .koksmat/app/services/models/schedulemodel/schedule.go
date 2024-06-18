/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateGoModel v2
package schedulemodel

import (
	"encoding/json"
	"time"
)

func UnmarshalSchedule(data []byte) (Schedule, error) {
	var r Schedule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Schedule) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Schedule struct {
	ID          int         `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatedBy   string      `json:"created_by"`
	UpdatedAt   time.Time   `json:"updated_at"`
	UpdatedBy   string      `json:"updated_by"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Cron        string      `json:"cron"`
	Job_id      int         `json:"job_id"`
	Data        interface{} `json:"data"`
}

package events

import (
	"encoding/json"
	"time"
)

type Event struct {
	Timestamp time.Time       `json:"timestamp"`
	Subject   string          `json:"subject"`
	Data      json.RawMessage `json:"data"`
}

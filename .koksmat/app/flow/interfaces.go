package flow

import "encoding/json"

// Storage interface for saving and loading flows
type Storage interface {
	Save(id string, flowJSON json.RawMessage) error
	Load(id string) (json.RawMessage, error)
	GetEvents() ([]interface{}, error)
}

// Emitter interface for emitting events
type Emitter interface {
	Emit(event string, data interface{})
}

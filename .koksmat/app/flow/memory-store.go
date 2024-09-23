package flow

import (
	"encoding/json"
	"errors"
)

// InMemoryStorage implements Storage using a map
type InMemoryStorage struct {
	flows map[string]json.RawMessage
}

// NewInMemoryStorage creates a new instance of InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		flows: make(map[string]json.RawMessage),
	}
}

// Save stores the flow JSON by ID
func (s *InMemoryStorage) Save(id string, flowJSON json.RawMessage) error {
	s.flows[id] = flowJSON
	return nil
}

// Load retrieves the flow JSON by ID
func (s *InMemoryStorage) Load(id string) (json.RawMessage, error) {
	flowJSON, exists := s.flows[id]
	if !exists {
		return nil, errors.New("flow not found")
	}
	return flowJSON, nil
}

func (s *InMemoryStorage) GetEvents() ([]interface{}, error) {
	return nil, nil
}

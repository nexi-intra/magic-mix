package flow

import "errors"

// InMemoryStorage implements Storage using a map
type InMemoryStorage struct {
	flows map[string]string
}

// NewInMemoryStorage creates a new instance of InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		flows: make(map[string]string),
	}
}

// Save stores the flow JSON by ID
func (s *InMemoryStorage) Save(id string, flowJSON string) error {
	s.flows[id] = flowJSON
	return nil
}

// Load retrieves the flow JSON by ID
func (s *InMemoryStorage) Load(id string) (string, error) {
	flowJSON, exists := s.flows[id]
	if !exists {
		return "", errors.New("flow not found")
	}
	return flowJSON, nil
}

package flow

// Storage interface for saving and loading flows
type Storage interface {
	Save(id string, flowJSON string) error
	Load(id string) (string, error)
}

// Emitter interface for emitting events
type Emitter interface {
	Emit(event string, data interface{})
}

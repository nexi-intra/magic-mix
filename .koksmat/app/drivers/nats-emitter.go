package drivers

import (
	"log"
)

// NATSEmitter implements the flow.Emitter interface using NATS
type NATSEmitter struct {
	//nc *nats.Conn
}

// NewNATSEmitter creates a new NATSEmitter instance
func NewNATSEmitter() *NATSEmitter {
	return &NATSEmitter{}
}

// Emit sends an event to a NATS subject
func (e *NATSEmitter) Emit(event string, data interface{}) {
	log.Println("not implemented")
	// dataJSON, err := json.Marshal(data)
	// if err != nil {
	// 	log.Printf("failed to marshal event data: %v", err)
	// 	return
	// }

	// err = e.nc.Publish(event, dataJSON)
	// if err != nil {
	// 	log.Printf("failed to emit event: %v", err)
	// }
}

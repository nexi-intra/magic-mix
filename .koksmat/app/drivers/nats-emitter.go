package drivers

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// NATSEmitter implements the flow.Emitter interface using NATS
type NATSEmitter struct {
	nc *nats.Conn
}

// NewNATSEmitter creates a new NATSEmitter instance
func NewNATSEmitter(nc *nats.Conn) *NATSEmitter {
	return &NATSEmitter{
		nc: nc,
	}
}

// Emit sends an event to a NATS subject
func (e *NATSEmitter) Emit(event string, data interface{}) {

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal event data: %v", err)
		return
	}

	err = e.nc.Publish("workflow.event."+event, dataJSON)
	if err != nil {
		log.Printf("failed to emit event: %v", err)
	}
}

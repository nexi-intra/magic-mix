package flow

import "fmt"

// ConsoleEmitter implements Emitter by printing events to the console
type ConsoleEmitter struct{}

// Emit logs the event and associated data to the console
func (e *ConsoleEmitter) Emit(event string, data interface{}) {
	fmt.Printf("Event: %s, Data: %v\n", event, data)
}

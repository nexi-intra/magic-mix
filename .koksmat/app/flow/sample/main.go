package main

import (
	"fmt"

	"github.com/magicbutton/magic-mix/flow"
)

func main() {
	// Initialize storage and emitter
	storage := flow.NewInMemoryStorage()
	emitter := &flow.ConsoleEmitter{}

	// Create flow engine
	engine := flow.NewFlowEngine(storage, emitter)

	// Add a new flow
	err := engine.AddFlow("flow1", []byte(`{"name": "example flow"}`))
	if err != nil {
		fmt.Println("Error adding flow:", err)
		return
	}

	// Start the flow
	err = engine.StartFlow("flow1")
	if err != nil {
		fmt.Println("Error starting flow:", err)
		return
	}

	// Save the flow
	err = engine.SaveFlow("flow1")
	if err != nil {
		fmt.Println("Error saving flow:", err)
		return
	}

	// Load the flow back
	err = engine.LoadFlow("flow1")
	if err != nil {
		fmt.Println("Error loading flow:", err)
		return
	}

	// Get the flow status
	flowInstance, _ := engine.GetFlow("flow1")
	fmt.Println("Flow status:", flowInstance.GetStatus())
}

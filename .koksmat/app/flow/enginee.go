package flow

import (
	"encoding/json"
	"errors"
	"sync"
)

// FlowEngine manages multiple flows
type FlowEngine struct {
	flows   map[string]*Flow
	storage Storage
	emitter Emitter
	mu      sync.RWMutex
	loops   int
}

// NewFlowEngine creates a new FlowEngine instance
func NewFlowEngine(storage Storage, emitter Emitter) *FlowEngine {
	return &FlowEngine{
		flows:   make(map[string]*Flow),
		storage: storage,
		emitter: emitter,
		mu:      sync.RWMutex{},
	}
}

// AddFlow adds a new flow to the engine and emits an event
func (fe *FlowEngine) AddFlow(id string, flowJSON json.RawMessage) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	if _, exists := fe.flows[id]; exists {
		return errors.New("flow with the given ID already exists")
	}

	receipe := RecipeV1{}

	json.Unmarshal(flowJSON, &receipe.Definition)

	flow := &Flow{
		ID:       id,
		FlowJSON: flowJSON,
		Status:   StatusStopped,
		Recipe:   receipe,
	}

	fe.flows[id] = flow
	fe.emitter.Emit("FlowAdded", flow)
	return nil
}

// StartFlow starts a flow by ID and emits an event
func (fe *FlowEngine) StartFlow(id string) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	flow, exists := fe.flows[id]
	if !exists {
		return errors.New("flow not found")
	}

	err := flow.StartFlow()
	if err != nil {
		return err
	}

	fe.emitter.Emit("FlowStarted", flow)
	return nil
}

// PauseFlow pauses a flow by ID and emits an event
func (fe *FlowEngine) PauseFlow(id string) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	flow, exists := fe.flows[id]
	if !exists {
		return errors.New("flow not found")
	}

	err := flow.PauseFlow()
	if err != nil {
		return err
	}

	fe.emitter.Emit("FlowPaused", flow)
	return nil
}

// StopFlow stops a flow by ID and emits an event
func (fe *FlowEngine) StopFlow(id string) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	flow, exists := fe.flows[id]
	if !exists {
		return errors.New("flow not found")
	}

	err := flow.StopFlow()
	if err != nil {
		return err
	}

	fe.emitter.Emit("FlowStopped", flow)
	return nil
}

// DeleteFlow deletes a flow by ID and emits an event
func (fe *FlowEngine) DeleteFlow(id string) error {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	_, exists := fe.flows[id]
	if !exists {
		return errors.New("flow not found")
	}

	delete(fe.flows, id)
	fe.emitter.Emit("FlowDeleted", id)
	return nil
}

// GetFlow returns the flow by ID
func (fe *FlowEngine) GetFlow(id string) (*Flow, error) {
	fe.mu.RLock()
	defer fe.mu.RUnlock()

	flow, exists := fe.flows[id]
	if !exists {
		return nil, errors.New("flow not found")
	}

	return flow, nil
}

// SaveFlow uses the storage to save a flow by ID
func (fe *FlowEngine) SaveFlow(id string) error {
	fe.mu.RLock()
	defer fe.mu.RUnlock()

	flow, exists := fe.flows[id]
	if !exists {
		return errors.New("flow not found")
	}

	return fe.storage.Save(id, flow.FlowJSON)
}

// LoadFlow loads a flow using storage and adds it to the engine
func (fe *FlowEngine) LoadFlow(id string) error {
	flowJSON, err := fe.storage.Load(id)
	if err != nil {
		return err
	}

	return fe.AddFlow(id, flowJSON)
}

func (fe *FlowEngine) GetFlows() map[string]*Flow {
	fe.mu.RLock()
	defer fe.mu.RUnlock()

	return fe.flows
}

func (fe *FlowEngine) Tick() error {
	if (fe.loops % 100) == 0 {
		fe.emitter.Emit("EngineeTick", fe.loops)
		fe.loops = 0 //reset the loop counter
	}
	fe.loops++

	events, err := fe.storage.GetEvents()
	if err != nil {
		return err
	}

	for _, event := range events {
		fe.emitter.Emit("Event", event)
	}

	//for each flow in the engine
	for _, flow := range fe.flows {
		if flow.Status == StatusRunning {
			flow.onTick()

			// err := nil
			// if err != nil {
			// 	return err
			// }
		}
	}
	return nil

}

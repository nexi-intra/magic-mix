package flow

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

// Flow represents a single flow instance
type Flow struct {
	ID         string          // Unique ID for the flow
	Status     string          // Current status of the flow
	Lease      time.Duration   // Lease duration for the flow
	LeaseStart time.Time       // Start time of the lease
	FlowJSON   json.RawMessage // Flow definition in JSON format
	Recipe     RecipeV1        // Flow recipe
}

// Possible statuses for a flow
const (
	StatusRunning = "running"
	StatusPaused  = "paused"
	StatusStopped = "stopped"
	StatusDeleted = "deleted"
)

// NewFlow creates a new flow instance with default values
func NewFlow(id string, flowJSON json.RawMessage) *Flow {
	receipe := RecipeV1{}
	json.Unmarshal([]byte(flowJSON), &receipe)
	return &Flow{
		ID:       id,
		Status:   StatusStopped,
		FlowJSON: flowJSON,
		Recipe:   receipe,
	}
}

// StartFlow starts the workflow, setting the status to running
func (f *Flow) StartFlow() error {
	if f.Status == StatusRunning {
		return errors.New("flow is already running")
	}
	f.Status = StatusRunning
	f.LeaseStart = time.Now()
	return nil
}

// PauseFlow pauses the workflow, changing the status to paused
func (f *Flow) PauseFlow() error {
	if f.Status != StatusRunning {
		return errors.New("cannot pause a non-running flow")
	}
	f.Status = StatusPaused
	return nil
}

// StopFlow stops the workflow, changing the status to stopped
func (f *Flow) StopFlow() error {
	if f.Status == StatusStopped {
		return errors.New("flow is already stopped")
	}
	f.Status = StatusStopped
	return nil
}

// DeleteFlow marks the workflow as deleted
func (f *Flow) DeleteFlow() error {
	if f.Status == StatusDeleted {
		return errors.New("flow is already deleted")
	}
	f.Status = StatusDeleted
	return nil
}

// GetStatus returns the current status of the flow
func (f *Flow) GetStatus() string {
	return f.Status
}

// GetLease returns the current lease duration for the flow
func (f *Flow) GetLease() time.Duration {
	return f.Lease
}

// ExtendLease extends the lease duration for the flow
func (f *Flow) ExtendLease(extension time.Duration) error {
	if f.Status == StatusStopped || f.Status == StatusDeleted {
		return errors.New("cannot extend lease for stopped or deleted flow")
	}
	f.Lease += extension
	return nil
}

// ReleaseLease resets the lease, ending the current lease
func (f *Flow) ReleaseLease() {
	f.Lease = 0
	f.LeaseStart = time.Time{}
}

// LoadFlowJSON loads a flow configuration from a JSON string
func (f *Flow) LoadFlowJSON(json json.RawMessage) error {
	if f.Status == StatusRunning {
		return errors.New("cannot load flow while running")
	}
	f.FlowJSON = json
	return nil
}

// GetFlowJSON returns the flow configuration as a JSON string
func (f *Flow) GetFlowJSON() json.RawMessage {
	return f.FlowJSON
}

// onTick is called on every tick of the flow

// use for evaluating it the flow should continue  being in memory or put to sleep

func (f *Flow) onTick() error {
	log.Println("Tick", f.Recipe.Definition.ID, f.Recipe.Definition.Name)
	return nil
}

// This function is used to decide if a flow should come active or not
// The koksmat_event table decides this

/*

Example case to illustrate the flow using the Pizza order process

Phase 1 - Feeling hungry - Registering a new order in table Orders with status "Posted" triggers the flow to start

The koksmat trigger will make a lookup in the koksmat_model table to find if the is a flow for the table, if so it will
insert a new event in the koksmat_event table with the name of the table, the id and a snapshot of the record as a JSONB structure.

SQL trigger on Order insert/update in Orders table add event to koksmat_event table with status "Active" and event "OrderPosted"
Ticker runs select * from koksmat_event where event = "OrderPosted" and status = "Active"


*/

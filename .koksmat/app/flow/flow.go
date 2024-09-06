package flow

import (
	"errors"
	"time"
)

// Flow represents a single flow instance
type Flow struct {
	ID         string        // Unique ID for the flow
	Status     string        // Current status of the flow
	Lease      time.Duration // Lease duration for the flow
	LeaseStart time.Time     // Start time of the lease
	FlowJSON   string        // Flow definition in JSON format
}

// Possible statuses for a flow
const (
	StatusRunning = "running"
	StatusPaused  = "paused"
	StatusStopped = "stopped"
	StatusDeleted = "deleted"
)

// NewFlow creates a new flow instance with default values
func NewFlow(id string, flowJSON string) *Flow {
	return &Flow{
		ID:       id,
		Status:   StatusStopped,
		FlowJSON: flowJSON,
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
func (f *Flow) LoadFlowJSON(json string) error {
	if f.Status == StatusRunning {
		return errors.New("cannot load flow while running")
	}
	f.FlowJSON = json
	return nil
}

// GetFlowJSON returns the flow configuration as a JSON string
func (f *Flow) GetFlowJSON() string {
	return f.FlowJSON
}

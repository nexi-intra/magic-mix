package flow

import (
	"testing"
	"time"
)

// TestNewFlow checks if a new flow is initialized correctly
func TestNewFlow(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	if f.ID != "123" {
		t.Errorf("Expected ID '123', got %s", f.ID)
	}
	if f.Status != StatusStopped {
		t.Errorf("Expected status 'stopped', got %s", f.Status)
	}
	if f.FlowJSON != `{"name":"test-flow"}` {
		t.Errorf("Expected FlowJSON `{\"name\":\"test-flow\"}`, got %s", f.FlowJSON)
	}
}

// TestStartFlow checks if a flow is started correctly
func TestStartFlow(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	err := f.StartFlow()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.Status != StatusRunning {
		t.Errorf("Expected status 'running', got %s", f.Status)
	}
	if f.LeaseStart.IsZero() {
		t.Error("Expected LeaseStart to be set, but it is zero")
	}
}

// TestPauseFlow checks if a flow is paused correctly
func TestPauseFlow(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	f.StartFlow()
	err := f.PauseFlow()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.Status != StatusPaused {
		t.Errorf("Expected status 'paused', got %s", f.Status)
	}
}

// TestStopFlow checks if a flow is stopped correctly
func TestStopFlow(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	f.StartFlow()
	err := f.StopFlow()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.Status != StatusStopped {
		t.Errorf("Expected status 'stopped', got %s", f.Status)
	}
}

// TestDeleteFlow checks if a flow is deleted correctly
func TestDeleteFlow(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	err := f.DeleteFlow()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.Status != StatusDeleted {
		t.Errorf("Expected status 'deleted', got %s", f.Status)
	}
}

// TestExtendLease checks if lease extension works correctly
func TestExtendLease(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	f.StartFlow()
	f.Lease = 10 * time.Second
	err := f.ExtendLease(5 * time.Second)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.Lease != 15*time.Second {
		t.Errorf("Expected lease duration to be 15s, got %v", f.Lease)
	}
}

// TestReleaseLease checks if lease is released correctly
func TestReleaseLease(t *testing.T) {
	f := NewFlow("123", `{"name":"test-flow"}`)
	f.StartFlow()
	f.ReleaseLease()
	if f.Lease != 0 {
		t.Errorf("Expected lease to be 0, got %v", f.Lease)
	}
	if !f.LeaseStart.IsZero() {
		t.Error("Expected LeaseStart to be reset to zero, but it's not")
	}
}

// TestLoadFlowJSON checks if loading a flow configuration works correctly
func TestLoadFlowJSON(t *testing.T) {
	f := NewFlow("123", `{"name":"initial-flow"}`)
	err := f.LoadFlowJSON(`{"name":"updated-flow"}`)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if f.FlowJSON != `{"name":"updated-flow"}` {
		t.Errorf("Expected FlowJSON `{\"name\":\"updated-flow\"}`, got %s", f.FlowJSON)
	}
}

// TestLoadFlowJSONWhileRunning checks error when loading JSON while flow is running
func TestLoadFlowJSONWhileRunning(t *testing.T) {
	f := NewFlow("123", `{"name":"initial-flow"}`)
	f.StartFlow()
	err := f.LoadFlowJSON(`{"name":"updated-flow"}`)
	if err == nil {
		t.Error("Expected an error when loading flow JSON while running, got nil")
	}
}

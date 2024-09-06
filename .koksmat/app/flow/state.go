package flow

import "fmt"

// State represents the workflow state
type State string

// Define some sample states
const (
	StateIdle    State = "Idle"
	StateRunning State = "Running"
	StatePaused  State = "Paused"
	StateStopped State = "Stopped"
)

// Event represents the event that can trigger a state transition
type Event string

// Define some sample events
const (
	EventStart  Event = "Start"
	EventPause  Event = "Pause"
	EventResume Event = "Resume"
	EventStop   Event = "Stop"
)

// StateMachine struct holds the current state and possible transitions
type StateMachine struct {
	CurrentState State
	Transitions  map[State]map[Event]State
}

// NewStateMachine initializes a new StateMachine with possible transitions
func NewStateMachine() *StateMachine {
	return &StateMachine{
		CurrentState: StateIdle, // Initial state
		Transitions: map[State]map[Event]State{
			StateIdle: {
				EventStart: StateRunning,
			},
			StateRunning: {
				EventPause: StatePaused,
				EventStop:  StateStopped,
			},
			StatePaused: {
				EventResume: StateRunning,
				EventStop:   StateStopped,
			},
			StateStopped: {
				EventStart: StateIdle,
			},
		},
	}
}

// Trigger processes an event and updates the current state if a valid transition exists
func (sm *StateMachine) Trigger(event Event) error {
	if nextState, ok := sm.Transitions[sm.CurrentState][event]; ok {
		fmt.Printf("Transitioning from %s to %s on event %s\n", sm.CurrentState, nextState, event)
		sm.CurrentState = nextState
		return nil
	}
	return fmt.Errorf("invalid transition from %s on event %s", sm.CurrentState, event)
}

// GetCurrentState returns the current state of the StateMachine
func (sm *StateMachine) GetCurrentState() State {
	return sm.CurrentState
}

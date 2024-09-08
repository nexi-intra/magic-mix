package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStorage is a mock implementation of the Storage interface
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Save(id string, flowJSON string) error {
	args := m.Called(id, flowJSON)
	return args.Error(0)
}

func (m *MockStorage) Load(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

// MockEmitter is a mock implementation of the Emitter interface
type MockEmitter struct {
	mock.Mock
}

func (m *MockEmitter) Emit(eventType string, data interface{}) {
	m.Called(eventType, data)
}

func TestFlowEngine_AddFlow(t *testing.T) {
	storage := new(MockStorage)
	emitter := new(MockEmitter)
	engine := NewFlowEngine(storage, emitter)

	flowID := "test-flow"
	flowJSON := `{"steps":[{"name":"step1"}]}`
	emitter.On("Emit", "FlowAdded", mock.Anything).Once()

	err := engine.AddFlow(flowID, flowJSON)
	assert.NoError(t, err)

	// Check that the flow was added correctly
	flowInstance, err := engine.GetFlow(flowID)
	assert.NoError(t, err)
	assert.Equal(t, flowID, flowInstance.ID)
	assert.Equal(t, flowJSON, flowInstance.FlowJSON)

	emitter.AssertCalled(t, "Emit", "FlowAdded", mock.Anything)
}

func TestFlowEngine_AddFlow_DuplicateFlow(t *testing.T) {
	storage := new(MockStorage)
	emitter := new(MockEmitter)
	engine := NewFlowEngine(storage, emitter)

	flowID := "duplicate-flow"
	flowJSON := `{"steps":[{"name":"step1"}]}`
	emitter.On("Emit", "FlowAdded", mock.Anything).Once()
	err := engine.AddFlow(flowID, flowJSON)
	assert.NoError(t, err)

	// Trying to add a flow with the same ID should return an error
	err = engine.AddFlow(flowID, flowJSON)
	assert.Error(t, err)
	assert.Equal(t, "flow with the given ID already exists", err.Error())
}

func TestFlowEngine_StartFlow(t *testing.T) {
	storage := new(MockStorage)
	emitter := new(MockEmitter)
	engine := NewFlowEngine(storage, emitter)

	flowID := "start-flow"
	flowJSON := `{"steps":[{"name":"step1"}]}`
	emitter.On("Emit", "FlowAdded", mock.Anything).Once()
	engine.AddFlow(flowID, flowJSON)

	flowInstance, _ := engine.GetFlow(flowID)
	// flowInstance.StartFlow = func() error {
	// 	flowInstance.Status = StatusRunning
	// 	return nil
	// }

	emitter.On("Emit", "FlowStarted", mock.Anything).Once()

	err := engine.StartFlow(flowID)
	assert.NoError(t, err)

	flowInstance, err = engine.GetFlow(flowID)
	assert.NoError(t, err)
	assert.Equal(t, StatusRunning, flowInstance.Status)

	emitter.AssertCalled(t, "Emit", "FlowStarted", mock.Anything)
}

func TestFlowEngine_DeleteFlow(t *testing.T) {
	storage := new(MockStorage)
	emitter := new(MockEmitter)
	engine := NewFlowEngine(storage, emitter)

	flowID := "delete-flow"
	flowJSON := `{"steps":[{"name":"step1"}]}`
	emitter.On("Emit", "FlowAdded", mock.Anything).Once()
	engine.AddFlow(flowID, flowJSON)

	emitter.On("Emit", "FlowDeleted", flowID).Once()

	err := engine.DeleteFlow(flowID)
	assert.NoError(t, err)

	// The flow should no longer exist
	_, err = engine.GetFlow(flowID)
	assert.Error(t, err)
	assert.Equal(t, "flow not found", err.Error())

	emitter.AssertCalled(t, "Emit", "FlowDeleted", flowID)
}

func TestFlowEngine_GetFlows(t *testing.T) {
	storage := new(MockStorage)
	emitter := new(MockEmitter)
	engine := NewFlowEngine(storage, emitter)

	// Add some flows
	flowID1 := "flow1"
	flowJSON1 := `{"steps":[{"name":"step1"}]}`
	flowID2 := "flow2"
	flowJSON2 := `{"steps":[{"name":"step2"}]}`

	emitter.On("Emit", "FlowAdded", mock.Anything).Twice()

	err := engine.AddFlow(flowID1, flowJSON1)
	assert.NoError(t, err)

	err = engine.AddFlow(flowID2, flowJSON2)
	assert.NoError(t, err)

	// Now retrieve the flows
	flows := engine.GetFlows()

	// Check if the number of flows is correct
	assert.Equal(t, 2, len(flows))

	// Check if the correct flows are returned
	assert.Contains(t, flows, flowID1)
	assert.Contains(t, flows, flowID2)

	// Check the contents of the flows
	assert.Equal(t, flowID1, flows[flowID1].ID)
	assert.Equal(t, flowJSON1, flows[flowID1].FlowJSON)
	assert.Equal(t, flowID2, flows[flowID2].ID)
	assert.Equal(t, flowJSON2, flows[flowID2].FlowJSON)

	// Verify the events were emitted as expected
	emitter.AssertCalled(t, "Emit", "FlowAdded", mock.Anything)
}

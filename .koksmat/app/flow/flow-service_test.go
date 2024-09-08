package flow

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowEngineService_AddFlow(t *testing.T) {
	storage := &MockStorage{}
	emitter := &MockEmitter{}
	engine := NewFlowEngine(storage, emitter)
	service := NewFlowEngineService(engine)

	// Create the add flow request
	request := `{
		"type": "add_flow",
		"payload": {
			"id": "flow1",
			"flow_json": "{\"steps\":[{\"name\":\"step1\"}]}"
		}
	}`

	// Call handleRequest
	response, err := service.HandleRequest([]byte(request))
	assert.NoError(t, err)

	// Parse the response
	var respMap map[string]string
	err = json.Unmarshal(response, &respMap)
	assert.NoError(t, err)

	// Assert the expected message
	assert.Equal(t, "flow added successfully", respMap["message"])
}

func TestFlowEngineService_AddDuplicateFlow(t *testing.T) {
	storage := &MockStorage{}
	emitter := &MockEmitter{}
	engine := NewFlowEngine(storage, emitter)
	service := NewFlowEngineService(engine)

	// Add the flow first
	addFlowRequest := `{
		"type": "add_flow",
		"payload": {
			"id": "flow1",
			"flow_json": "{\"steps\":[{\"name\":\"step1\"}]}"
		}
	}`
	service.HandleRequest([]byte(addFlowRequest))

	// Try adding the same flow again
	duplicateFlowRequest := `{
		"type": "add_flow",
		"payload": {
			"id": "flow1",
			"flow_json": "{\"steps\":[{\"name\":\"step1\"}]}"
		}
	}`

	// Call handleRequest and expect an error
	response, err := service.HandleRequest([]byte(duplicateFlowRequest))
	assert.NoError(t, err)

	// Parse the response
	var respMap map[string]string
	err = json.Unmarshal(response, &respMap)
	assert.NoError(t, err)

	// Assert the expected error message
	assert.Contains(t, respMap["error"], "flow with the given ID already exists")
}

func TestFlowEngineService_StartFlow(t *testing.T) {
	storage := &MockStorage{}
	emitter := &MockEmitter{}
	engine := NewFlowEngine(storage, emitter)
	service := NewFlowEngineService(engine)

	// Add the flow
	addFlowRequest := `{
		"type": "add_flow",
		"payload": {
			"id": "flow1",
			"flow_json": "{\"steps\":[{\"name\":\"step1\"}]}"
		}
	}`
	service.HandleRequest([]byte(addFlowRequest))

	// Create the start flow request
	startFlowRequest := `{
		"type": "start_flow",
		"payload": {
			"id": "flow1"
		}
	}`

	// Call handleRequest
	response, err := service.HandleRequest([]byte(startFlowRequest))
	assert.NoError(t, err)

	// Parse the response
	var respMap map[string]string
	err = json.Unmarshal(response, &respMap)
	assert.NoError(t, err)

	// Assert the expected message
	assert.Equal(t, "flow started successfully", respMap["message"])
}

func TestFlowEngineService_UnknownRequestType(t *testing.T) {
	storage := &MockStorage{}
	emitter := &MockEmitter{}
	engine := NewFlowEngine(storage, emitter)
	service := NewFlowEngineService(engine)

	// Create a request with an unknown type
	request := `{
		"type": "unknown_type",
		"payload": {}
	}`

	// Call handleRequest and expect an error
	response, err := service.HandleRequest([]byte(request))
	assert.NoError(t, err)

	// Parse the response
	var respMap map[string]string
	err = json.Unmarshal(response, &respMap)
	assert.NoError(t, err)

	// Assert the expected error message
	assert.Contains(t, respMap["error"], "unknown request type")
}

package flow

import (
	"encoding/json"
	"errors"
	"log"
)

// Request struct to determine which FlowEngine method to call
type Request struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// AddFlowPayload represents the payload for adding a flow
type AddFlowPayload struct {
	ID       string `json:"id"`
	FlowJSON string `json:"flow_json"`
}

// StartFlowPayload represents the payload for starting a flow
type StartFlowPayload struct {
	ID string `json:"id"`
}

// FlowEngineService wraps the FlowEngine functionality
type FlowEngineService struct {
	engine *FlowEngine
}

// NewFlowEngineService creates a new service wrapping the FlowEngine
func NewFlowEngineService(engine *FlowEngine) *FlowEngineService {
	return &FlowEngineService{
		engine: engine,
	}
}

// HandleRequest processes the incoming request and returns a response as a JSON object or an error
func (s *FlowEngineService) HandleRequest(requestJSON []byte) ([]byte, error) {
	var req Request

	// Unmarshal the generic request
	err := json.Unmarshal(requestJSON, &req)
	if err != nil {
		return nil, err
	}
	log.Println("Flow Request type: ", req.Type)
	// Switch case on request type and get the response from each handler
	switch req.Type {
	case "add_flow":
		return s.handleAddFlow(req.Payload)
	case "start_flow":
		return s.handleStartFlow(req.Payload)
	case "pause_flow":
		return s.handlePauseFlow(req.Payload)
	case "stop_flow":
		return s.handleStopFlow(req.Payload)
	case "delete_flow":
		return s.handleDeleteFlow(req.Payload)
	case "get_flow":
		return s.handleGetFlow(req.Payload)
	case "get_all_flows":
		return s.handleGetAllFlows()
	default:
		return nil, errors.New("unknown request type")
	}
}

// Handler functions for each request, returning JSON and error

func (s *FlowEngineService) handleAddFlow(payload json.RawMessage) ([]byte, error) {
	var addFlowPayload AddFlowPayload
	err := json.Unmarshal(payload, &addFlowPayload)
	if err != nil {
		return nil, err
	}

	err = s.engine.AddFlow(addFlowPayload.ID, addFlowPayload.FlowJSON)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return a JSON response
	resp := map[string]string{"message": "flow added successfully"}
	return json.Marshal(resp)
}

func (s *FlowEngineService) handleStartFlow(payload json.RawMessage) ([]byte, error) {
	var startFlowPayload StartFlowPayload
	err := json.Unmarshal(payload, &startFlowPayload)
	if err != nil {
		return nil, err
	}

	err = s.engine.StartFlow(startFlowPayload.ID)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return a JSON response
	resp := map[string]string{"message": "flow started successfully"}
	return json.Marshal(resp)
}

func (s *FlowEngineService) handlePauseFlow(payload json.RawMessage) ([]byte, error) {
	var pauseFlowPayload StartFlowPayload
	err := json.Unmarshal(payload, &pauseFlowPayload)
	if err != nil {
		return nil, err
	}

	err = s.engine.PauseFlow(pauseFlowPayload.ID)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return a JSON response
	resp := map[string]string{"message": "flow paused successfully"}
	return json.Marshal(resp)
}

func (s *FlowEngineService) handleStopFlow(payload json.RawMessage) ([]byte, error) {
	var stopFlowPayload StartFlowPayload
	err := json.Unmarshal(payload, &stopFlowPayload)
	if err != nil {
		return nil, err
	}

	err = s.engine.StopFlow(stopFlowPayload.ID)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return a JSON response
	resp := map[string]string{"message": "flow stopped successfully"}
	return json.Marshal(resp)
}

func (s *FlowEngineService) handleDeleteFlow(payload json.RawMessage) ([]byte, error) {
	var deleteFlowPayload StartFlowPayload
	err := json.Unmarshal(payload, &deleteFlowPayload)
	if err != nil {
		return nil, err
	}

	err = s.engine.DeleteFlow(deleteFlowPayload.ID)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return a JSON response
	resp := map[string]string{"message": "flow deleted successfully"}
	return json.Marshal(resp)
}

func (s *FlowEngineService) handleGetFlow(payload json.RawMessage) ([]byte, error) {
	var getFlowPayload StartFlowPayload
	err := json.Unmarshal(payload, &getFlowPayload)
	if err != nil {
		return nil, err
	}

	flow, err := s.engine.GetFlow(getFlowPayload.ID)
	if err != nil {
		resp := map[string]string{"error": err.Error()}
		return json.Marshal(resp)
	}

	// Return the flow as JSON
	return json.Marshal(flow)
}

func (s *FlowEngineService) handleGetAllFlows() ([]byte, error) {
	flows := s.engine.GetFlows()

	// Return all flows as JSON
	return json.Marshal(flows)
}

package work

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test JWT Authentication Middleware
func TestJWTAuthMiddleware(t *testing.T) {
	// Create a dummy handler that always returns success
	handler := jwtAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Create a valid token for the test
	token, err := createToken("test_user")
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Test a request with the valid token
	req := httptest.NewRequest("GET", "/update", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	// Test a request without a token
	req = httptest.NewRequest("GET", "/update", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status Unauthorized, got %v", w.Code)
	}
}

// Test object state persistence after update
func TestObjectPersistence(t *testing.T) {
	// Initialize the object
	obj := &Object{
		subObject: make(map[string]*SubObject),
		leases:    make(map[string]time.Time),
	}

	// Acquire lease and update the object
	err := obj.acquireLease("subObject1", time.Minute)
	if err != nil {
		t.Fatalf("Failed to acquire lease: %v", err)
	}

	err = obj.updateSubObject("subObject1", "newData", "testLeaseID")
	if err != nil {
		t.Fatalf("Failed to update sub-object: %v", err)
	}

	// Check if the object was persisted (this will verify the state was written to the file)
	err = persistState(obj)
	if err != nil {
		t.Errorf("Failed to persist object: %v", err)
	}
}

// Test Undo functionality
func TestUndo(t *testing.T) {
	obj := &Object{
		subObject: make(map[string]*SubObject),
		leases:    make(map[string]time.Time),
	}

	// Update the object
	obj.subObject["subObject1"] = &SubObject{Data: "originalData"}
	obj.pushUndo()

	// Make a new change and undo it
	obj.subObject["subObject1"] = &SubObject{Data: "newData"}
	err := obj.undo()
	if err != nil {
		t.Errorf("Failed to undo: %v", err)
	}

	// Check that the original data was restored
	if obj.subObject["subObject1"].Data != "originalData" {
		t.Errorf("Expected originalData, got %v", obj.subObject["subObject1"].Data)
	}
}

// Test Lease acquisition and expiration
func TestLeaseAcquisitionAndExpiration(t *testing.T) {
	obj := &Object{
		subObject: make(map[string]*SubObject),
		leases:    make(map[string]time.Time),
	}

	// Acquire a lease
	err := obj.acquireLease("subObject1", 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to acquire lease: %v", err)
	}

	// Ensure lease is valid
	if expiration, exists := obj.leases["subObject1"]; !exists || time.Now().After(expiration) {
		t.Errorf("Expected valid lease for subObject1")
	}

	// Wait for the lease to expire
	time.Sleep(3 * time.Second)

	// Ensure the lease has expired
	if expiration, exists := obj.leases["subObject1"]; exists && time.Now().Before(expiration) {
		t.Errorf("Expected lease to have expired")
	}
}

package github

import (
	"context"
	"log"
	"testing"
)

// TestNewFlow checks if a new flow is initialized correctly
func TestAuthenticate(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Errorf("Error authenticating: %s", err)
	}
	ctx := context.Background()

	a, b, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	log.Printf("Repositories: %v", a)
	log.Printf("Response: %v", b)

}

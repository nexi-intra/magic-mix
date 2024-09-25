package app

import (
	"log"
	"testing"
)

func TestProcessNull(t *testing.T) {

	_, err := Process([]string{}, nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

func TestProcessUnknown(t *testing.T) {
	log.Println("TestProcessUnknown")

	_, err := Process([]string{"unknown"}, nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

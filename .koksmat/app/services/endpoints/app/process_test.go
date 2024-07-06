package app

import (
	"log"
	"testing"
)

func TestProcessNull(t *testing.T) {

	_, err := Work([]string{})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

func TestProcessUnknown(t *testing.T) {
	log.Println("TestProcessUnknown")

	_, err := Work([]string{"unknown"})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

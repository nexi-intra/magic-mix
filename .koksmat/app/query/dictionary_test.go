package query

import "testing"

func TestGetstoredProcedures(t *testing.T) {
	records, err := GetStoredProcedures("booking")
	if err != nil {
		t.Errorf("GetStoredProcedures() failed: %v", err)
	}
	if len(*records) == 0 {
		t.Errorf("GetStoredProcedures() failed: %v", err)
	}
}

func TestGetstoredProcedure(t *testing.T) {
	records, err := GetStoredProcedures("booking")
	if err != nil {
		t.Errorf("GetStoredProcedures() failed: %v", err)
	}
	if len(*records) == 0 {
		t.Errorf("GetStoredProcedures() failed: %v", err)
	}

}

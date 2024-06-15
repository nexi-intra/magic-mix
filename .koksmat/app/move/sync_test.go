package move

import "testing"

func TestMove(t *testing.T) {
	err := Move("mix", "files", "sharepoint.pageviews")
	if err != nil {
		t.Errorf("Move() failed: %v", err)
	}

}

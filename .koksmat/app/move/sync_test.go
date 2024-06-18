package move

import "testing"

func TestMove(t *testing.T) {
	err := Move("mix", "files", "sharepoint.pageviews", "events", nil, &MoveOptions{dryRun: true})
	if err != nil {
		t.Errorf("Move() failed: %v", err)
	}

}

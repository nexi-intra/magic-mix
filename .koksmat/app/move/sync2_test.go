package move

import (
	"testing"
)

func TestSync2(t *testing.T) {
	err := SyncPageViews2("mix", "files")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

package officegraph

import (
	"testing"
)

func TestDownload1(t *testing.T) {
	token, _ := GetAuthToken()
	Download("https://graph.microsoft.com/v1.0/sites", token, 1)

}

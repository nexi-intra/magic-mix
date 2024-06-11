package move

import (
	"fmt"
	"testing"

	"github.com/magicbutton/magic-mix/services/endpoints/connection"
)

func TestHandshake(t *testing.T) {
	connections, err := connection.ConnectionSearch("%")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	//log.Println(connections.TotalItems, "Connections")
	for _, connection := range connections.Items {
		fmt.Println(connection.Name)
	}
}

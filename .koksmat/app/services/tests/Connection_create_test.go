/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma4.1
package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/magicbutton/magic-mix/services/endpoints/connection"
	"github.com/magicbutton/magic-mix/services/models/connectionmodel"
)

func TestConnectioncreate(t *testing.T) {
	// transformer v1
	object := connectionmodel.Connection{}

	result, err := connection.ConnectionCreate(object)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	assert.NotNil(t, result)

}

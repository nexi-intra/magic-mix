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

	"github.com/magicbutton/magic-mix/services/endpoints/importdata"
)

func TestImportDatasearch(t *testing.T) {

	result, err := importdata.ImportDataSearch(".")
	if err != nil {
		t.Errorf("Error %s", err)
	}
	assert.NotNil(t, result)

}

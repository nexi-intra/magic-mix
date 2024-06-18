package applogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var filename = "/Users/nielsgregersjohansen/kitchens/magic-apps/.koksmat/workdir/Estrazione Catalogo NEAR_20240404.xlsx"
// var sheetname = "Applicazioni v2"

var filename = "/Users/nielsgregersjohansen/kitchens/magic-people/.koksmat/workdir/Global Mail List (04-2024).xlsx"

var sheetname = "identities"

func TestReadSheetNames(t *testing.T) {

	names, err := SheetNames(filename)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, names)
}

func TestReadSheet(t *testing.T) {

	sheet, err := ReadSheet(filename, sheetname)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, sheet)
}

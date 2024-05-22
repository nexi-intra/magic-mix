package applogic

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLCreateAndInsert(t *testing.T) {

	sheet, err := ReadSheet(filename, sheetname)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, sheet)

	createtablesql := SheetToInsertCreateTable(sheet, sheetname)

	os.WriteFile("createtablesql.sql", []byte(createtablesql), 0644)
	batch := 0
	batchsize := 5000
	startIndex := 1 + (batchsize * batch)

	for startIndex < len(sheet.Rows) {
		inserttablesql := SheetToInsertCreateBatch(sheet, sheetname, startIndex, startIndex+batchsize)

		os.WriteFile(fmt.Sprintf("inserttablesql_%d.sql", batch), []byte(inserttablesql), 0644)
		batch++
		startIndex += batchsize
	}

}

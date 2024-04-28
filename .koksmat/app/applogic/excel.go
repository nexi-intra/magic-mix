package applogic

import (
	"sort"

	"github.com/tealeg/xlsx/v3"
)

func SheetNames(filename string) ([]string, error) {
	// open an existing file
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sort.Slice(wb.Sheets, func(i, j int) bool {
		return wb.Sheets[i].Name < wb.Sheets[j].Name
	})
	result := make([]string, 0)
	for _, sh := range wb.Sheets {
		result = append(result, sh.Name)

	}
	return result, nil
}

func ReadSheet(filename string, sheetName string) (*Sheet, error) {
	// open an existing file

	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	// get the first sheet
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		return nil, nil
	}

	result := &Sheet{Name: sh.Name}
	var row = &Row{}
	var cellVisitor xlsx.CellVisitorFunc = func(c *xlsx.Cell) error {

		value, err := c.FormattedValue()
		if err != nil {
			return err
		} else {
			colnumber, rownumber := c.GetCoordinates()
			row.Cells = append(row.Cells, Cell{Row: rownumber, Column: colnumber, Value: value})

		}
		return nil
	}

	var visitor xlsx.RowVisitor = func(r *xlsx.Row) error {
		row = &Row{}
		err := r.ForEachCell(cellVisitor)
		if err != nil {
			return err
		}
		result.Rows = append(result.Rows, *row)
		return nil
	}
	err = sh.ForEachRow(visitor, xlsx.SkipEmptyRows)
	if err != nil {
		return nil, err
	}

	return result, nil
}

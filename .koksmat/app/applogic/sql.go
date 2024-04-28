package applogic

import (
	"fmt"
	"strings"
)

func SheetToInsertCreateTable(sheet *Sheet, namespace string) string {

	var sql = fmt.Sprintf(`
	/*
	Create script
	
	*/
	


	CREATE SCHEMA IF NOT EXISTS "%s"
    AUTHORIZATION ps_user;
-- DROP TABLE IF EXISTS "%s"."%s";

CREATE TABLE IF NOT EXISTS "%s"."%s"
(
    id SERIAL PRIMARY KEY,
	`, namespace, namespace, sheet.Name, namespace, sheet.Name)

	for i, cell := range sheet.Rows[0].Cells {

		sql += fmt.Sprintf(`"col%s (column %d)" character varying COLLATE pg_catalog."default"
	`, cell.Value, i)

		if i < (len(sheet.Rows[0].Cells) - 1) {
			sql += ","
		}

	}
	sql += ");"

	return sql

}

func SheetToInsertCreateBatch(sheet *Sheet, namespace string, startIndex int, endIndex int) string {

	var sql = `
	/*
	Insert script
	
	*/
	

	`

	sql += fmt.Sprintf(`
	
INSERT INTO "%s"."%s"
(id,	
`, namespace, sheet.Name)
	for i, cell := range sheet.Rows[0].Cells {

		sql += fmt.Sprintf(`"col%s (column %d)"
		`, cell.Value, i)

		if i < (len(sheet.Rows[0].Cells) - 1) {
			sql += ","
		}

	}

	sql += ") VALUES "

	for rowIndex, row := range sheet.Rows {
		if rowIndex == 0 {
			continue
		}
		if rowIndex < startIndex {
			continue
		}
		if rowIndex >= endIndex {
			break
		}
		sql += "( DEFAULT,"

		for colIndex, cell := range row.Cells {
			sql += fmt.Sprintf(` -- %d 
			`, colIndex)

			sql += fmt.Sprintf("'%s'", strings.ReplaceAll(cell.Value, "'", "''"))
			if colIndex < (len(row.Cells) - 1) {
				sql += ","
			}

		}
		sql += ")"
		if rowIndex < (endIndex - 1) {
			sql += ","
		}
	}

	sql += ";"
	return sql

}

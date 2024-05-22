package applogic

import (
	"fmt"
	"strings"
)

func SheetToInsertCreateTable(sheet *Sheet, tablename string) string {
	namespace := "excelimport"
	var sql = fmt.Sprintf(`
	/*
	Create script
	
	*/
	


CREATE SCHEMA IF NOT EXISTS "%s";

DROP TABLE IF EXISTS "%s"."%s";

CREATE TABLE IF NOT EXISTS "%s"."%s"
(
    id SERIAL PRIMARY KEY,
	`, namespace, namespace, tablename, namespace, tablename)

	for i, cell := range sheet.Rows[0].Cells {

		sql += fmt.Sprintf(`"%s (column %d)" character varying COLLATE pg_catalog."default"
	`, cell.Value, i)

		if i < (len(sheet.Rows[0].Cells) - 1) {
			sql += ","
		}

	}
	sql += ");"

	return sql

}

func SheetToInsertCreateBatch(sheet *Sheet, tablename string, startIndex int, endIndex int) string {
	namespace := "excelimport"
	var sql = `
	/*
	Insert script
	
	*/
	

	`

	sql += fmt.Sprintf(`
	
INSERT INTO "%s"."%s"
(id,	
`, namespace, tablename)
	for i, cell := range sheet.Rows[0].Cells {

		sql += fmt.Sprintf(`"%s (column %d)"
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

	sql = strings.TrimSuffix(sql, ",")
	sql += ";"
	return sql

}

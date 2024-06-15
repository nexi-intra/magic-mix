package applogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/tealeg/xlsx/v3"
)

func ConvertExcelToSQL(filename string, sheetName string, tablename string, batchsize int) error {
	log.Println("Converting Excel to SQL")
	sheet, err := ReadSheet(filename, sheetName)
	if err != nil {
		log.Println("Error writing file", err)
		return err
	}
	if sheet == nil {
		log.Println("Sheet not found")
		return errors.New("Sheet not found")
	}
	createtablesql := SheetToInsertCreateTable(sheet, tablename)

	os.WriteFile(tablename+".createtablesql.sql", []byte(createtablesql), 0644)
	log.Println("Creation SQL file created")
	batch := 0

	startIndex := 1 + (batchsize * batch)

	for startIndex < len(sheet.Rows) {
		log.Println("Creating batch", batch)
		inserttablesql := SheetToInsertCreateBatch(sheet, tablename, startIndex, startIndex+batchsize)

		err = os.WriteFile(fmt.Sprintf(tablename+".inserttablesql_%d.sql", batch), []byte(inserttablesql), 0644)
		if err != nil {
			log.Println("Error writing file", err)
			return err
		}
		batch++
		startIndex += batchsize
	}
	log.Println("Conversion completed")
	return nil
}
func ensureBatchFolderExists(folder string) error {
	// Check if folder exists
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		// Folder doesn't exist, create it
		if err := os.Mkdir(folder, 0755); err != nil {
			return err
		}
	}
	return nil
}
func ConvertExcelToSQLJson(excelFilename string, outputFolder string) error {
	log.Println("Converting Excel to SQL")

	// Open the Excel file
	file, err := xlsx.OpenFile(excelFilename)
	if err != nil {
		return err

	}

	ensureBatchFolderExists(outputFolder)
	// Iterate through each sheet
	for _, sh := range file.Sheets {
		log.Println("Converting sheet", sh.Name)
		sheet, err := ReadSheet(excelFilename, sh.Name)
		if err != nil {
			return err
		}
		var rows []map[string]interface{}
		startIndex := 0
		// Iterate through the rows of the sheet
		//for rowIndex, row := range sheet.Rows {
		for startIndex < len(sheet.Rows) {
			// Create a map to store the row data
			row := sheet.Rows[startIndex]
			rowData := make(map[string]interface{})
			rowData["rownumber"] = startIndex + 1
			// Iterate through the cells of the row
			for cellIndex, cell := range row.Cells {
				// Get the value of the cell
				value := cell.Value

				// Create the key using the first row value appended with the column number
				if startIndex == 0 {
					// Skip the first row as it is used for keys

					continue
				}
				key := fmt.Sprintf("%s (column %d)", sheet.Rows[0].Cells[cellIndex].Value, cellIndex+1)

				// Add the value to the row data map
				rowData[key] = value
			}

			// Skip adding the first row to the rows array
			if startIndex > 0 {
				rows = append(rows, rowData)
			}
			startIndex++
		}

		// Convert the rows array to JSON
		jsonData, err := json.Marshal(rows)
		if err != nil {
			log.Fatalf("Failed to convert to JSON: %v", err)
		}

		// // Print or save the JSON data
		// fmt.Printf("Sheet: %s\n%s\n", sheet.Name, string(jsonData))
		err = os.WriteFile(path.Join(outputFolder, sh.Name+".json"), jsonData, 0644)

	}

	return nil
}

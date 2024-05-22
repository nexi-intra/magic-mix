package applogic

import (
	"errors"
	"fmt"
	"log"
	"os"
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

package collect

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/magicbutton/magic-mix/services/endpoints/importdata"
	"github.com/magicbutton/magic-mix/services/models/importdatamodel"
)

// EnumerateFiles walks through the given directory path and calls the callback function with the path of each file found.
func EnumerateFiles(root string, callback func(path string) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			return callback(path)
		}
		return nil
	})
}

func UploadBatch(rootPath string) {
	inserter := func(filePath string, chunk *string) error {
		log.Println("Uploading", filePath, len(*chunk), "bytes")
		//return nil

		_, importErr := importdata.ImportDataCreate(importdatamodel.ImportData{
			Name:        filePath,
			Description: "Imported file",

			Data: chunk,
		})
		return importErr
	}

	callback := func(filePath string) error {
		//log.Println("Uploading", filePath)
		return LoadJSON(filePath, inserter)
	}

	if err := EnumerateFiles(rootPath, callback); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

func UploadBatch2(rootPath string, db *sql.DB) {
	inserter := func(filePath string, chunk *string) error {
		log.Println("Uploading", filePath, len(*chunk), "bytes")
		ctx := context.Background()

		sqlstatement := fmt.Sprintf(`
		INSERT INTO importdata ( id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        data)
		VALUES (DEFAULT,
        DEFAULT,
        DEFAULT,
        'system', 
        'system', 
        '',
        '',
        $1,
        '',
        $2);
		`)
		_, err := db.ExecContext(ctx, sqlstatement, filePath, *chunk)
		if err != nil {
			log.Println("Error:", err)
			log.Printf("SQL: %s ...\n", sqlstatement[:1024])
			return err
		}

		//rowsAffected, _ := raw.RowsAffected()
		return nil
	}

	callback := func(filePath string) error {
		//log.Println("Uploading", filePath)
		return LoadJSON(filePath, inserter)
	}

	if err := EnumerateFiles(rootPath, callback); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

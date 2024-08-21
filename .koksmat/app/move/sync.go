package move

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const batchSize = 10000

type Record struct {
	// Define the structure based on the view's output
	// Example fields:

	Data json.RawMessage `json:"data"`
}

// ConnectDB initializes a connection to the database using a connection string
func ConnectDB(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

// GetDistinctValues fetches distinct values from the specified column in the given table.
func GetDistinctValues(db *sql.DB, tableName string, columnName string) (map[string]string, error) {
	// Define the query to get distinct values from the column
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s", columnName, tableName)
	log.Println("Executing query:", query)
	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store distinct values
	var values = make(map[string]string)

	// Iterate over the result set and append each distinct value to the slice
	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return nil, err
		}

		values[value] = value
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

// ExtractBatchNames retrieves distinct batch names from the specified table
func ExtractBatchNames(ctx context.Context, db *sql.DB, query string) (map[string]struct{}, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query batch names: %w", err)
	}
	defer rows.Close()

	batchNames := make(map[string]struct{})
	for rows.Next() {
		var batchName string
		if err := rows.Scan(&batchName); err != nil {
			return nil, fmt.Errorf("failed to scan batch name: %w", err)
		}
		batchNames[batchName] = struct{}{}
	}
	return batchNames, nil
}

// ExtractDataAsJSON retrieves data for a specific batch name and returns it as JSON
func ExtractDataAsJSON(ctx context.Context, db *sql.DB, batchName string) (string, error) {
	query := "SELECT * FROM sharepoint.pageviews WHERE batchname = $1"
	rows, err := db.QueryContext(ctx, query, batchName)
	if err != nil {
		return "", fmt.Errorf("failed to query pageviews for batchname %s: %w", batchName, err)
	}
	defer rows.Close()

	var results []map[string]interface{}
	cols, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("failed to get columns: %w", err)
	}

	for rows.Next() {
		row := make(map[string]interface{})
		vals := make([]interface{}, len(cols))
		for i := range vals {
			vals[i] = new(interface{})
		}
		if err := rows.Scan(vals...); err != nil {
			return "", fmt.Errorf("failed to scan row: %w", err)
		}
		for i, col := range cols {
			row[col] = *(vals[i].(*interface{}))
		}
		results = append(results, row)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(jsonData), nil
}

// InsertIntoImportTable inserts JSON data into the import table
func InsertIntoImportTable(ctx context.Context, db *sql.DB, batchName, jsonData string) error {
	query := `
        INSERT INTO public.importdata (
            created_by, updated_by, tenant, searchindex, name, data
        ) VALUES (
            '', '', '', '', '', $1
        )`
	_, err := db.ExecContext(ctx, query, jsonData)
	if err != nil {
		return fmt.Errorf("failed to insert into import table: %w", err)
	}
	return nil
}

func GetBatchNotSynced(sourceDB *sql.DB, targetDB *sql.DB, fromTableName string, toTableName string) ([]string, error) {

	// Extract batch names
	mixerBatchNames, err := GetDistinctValues(sourceDB, fromTableName, "batchname")
	if err != nil {
		return nil, err
	}

	filesBatchNames, err := GetDistinctValues(targetDB, toTableName, "batchname")
	if err != nil {
		return nil, err
	}

	unSyncedBatchNames := []string{}

	// Identify batch names only in target
	for batchName := range mixerBatchNames {
		if _, exists := filesBatchNames[batchName]; !exists {
			unSyncedBatchNames = append(unSyncedBatchNames, batchName)

		}
	}

	return unSyncedBatchNames, nil
}

type MoveOptions struct {
	dryRun bool
}

func Move(fromDatabase string, toDatabase string, fromTableName string, toTableName string, storedProcedure *string, options *MoveOptions) error {
	fromConnection, err := GetConnectionString(fromDatabase)
	if err != nil {
		return err
	}

	toConnection, err := GetConnectionString(toDatabase)
	if err != nil {
		return err
	}
	sourceDB, err := sql.Open("postgres", *fromConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceDB.Close()

	destDB, err := sql.Open("postgres", *toConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer destDB.Close()

	batchNamestoSync, err := GetBatchNotSynced(sourceDB, destDB, fromTableName, toTableName)
	if err != nil {
		return err
	}
	if options != nil && options.dryRun {
		log.Println("Dry run mode")
		log.Println("Batches to sync", batchNamestoSync)
		return nil
	}
	//return nil
	return MoveData(batchNamestoSync, sourceDB, fromTableName, destDB, storedProcedure)

}

func Copy(fromDatabase string, toDatabase string, fromTableName string, batchName string) error {
	fromConnection, err := GetConnectionString(fromDatabase)
	if err != nil {
		return err
	}

	toConnection, err := GetConnectionString(toDatabase)
	if err != nil {
		return err
	}
	sourceDB, err := sql.Open("postgres", *fromConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceDB.Close()

	destDB, err := sql.Open("postgres", *toConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer destDB.Close()

	//return nil
	return CopyData(batchName, sourceDB, fromTableName, destDB)

}

func MoveData(batchNamestoSync []string, sourceDB *sql.DB, sourceTable string, destDB *sql.DB, storedProcedure *string) error {
	for _, batchName := range batchNamestoSync {
		tx, err := destDB.BeginTx(context.Background(), nil)
		if err != nil {
			return err
		}
		offset := 0
		log.Println("Reading data for batch", batchName)
		for {
			sql := fmt.Sprintf("SELECT  row_to_json(d) as data FROM %s as d WHERE batchname = '%s' LIMIT %d OFFSET %d ", sourceTable, batchName, batchSize, offset)
			log.Println("Executing", sql)
			rows, err := sourceDB.Query(sql)
			if err != nil {
				return fmt.Errorf("Reading data using %s fails with : %w", sql, err)
			}

			var records []Record
			for rows.Next() {
				var record Record
				if err := rows.Scan(&record.Data); err != nil {
					return fmt.Errorf("Scanning data : %w", err)
				}
				records = append(records, record)
			}
			rows.Close()

			if len(records) == 0 {
				break
			}
			log.Println("Copying", len(records), "rows")
			jsonData, err := json.Marshal(records)
			if err != nil {
				return fmt.Errorf("Marshalling data : %w", err)

			}

			_, err = destDB.Exec(`
		INSERT INTO public.importdata(
	id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, searchindex, name, description, data)
	VALUES (DEFAULT, DEFAULT, '', DEFAULT, '', DEFAULT, '', $1, $2, '', $3)`, batchName, batchName, string(jsonData))
			if err != nil {
				return fmt.Errorf("Inserting data : %w", err)
			}

			offset += batchSize

			if err != nil {
				err = tx.Rollback()
				return fmt.Errorf("Rolling back batch %s: %w", batchName, err)
			}

		}
		if storedProcedure != nil {
			log.Println("Executing stored procedure", *storedProcedure)
			storedProcedureSQL := fmt.Sprintf("CALL %s()", *storedProcedure)
			_, err = destDB.Exec(storedProcedureSQL)

			if err != nil {
				err = tx.Rollback()
				return fmt.Errorf("Error executing stored procedure %s: %w - Rolling back", *storedProcedure, err)
			}
		}
		log.Println("Committing batch", batchName)
		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("Committing batch %s: %w", batchName, err)
		}

	}
	return nil
}

func CopyData(batchName string, sourceDB *sql.DB, sourceTable string, destDB *sql.DB) error {

	tx, err := destDB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	offset := 0
	log.Println("Reading data for batch", batchName)
	for {
		sql := fmt.Sprintf("SELECT  row_to_json(d) as data FROM %s as d WHERE batchname = '%s' LIMIT %d OFFSET %d ", sourceTable, batchName, batchSize, offset)
		log.Println("Executing", sql)
		rows, err := sourceDB.Query(sql)
		if err != nil {
			return fmt.Errorf("Reading data using %s fails with : %w", sql, err)
		}

		var records []Record
		for rows.Next() {
			var record Record
			if err := rows.Scan(&record.Data); err != nil {
				return fmt.Errorf("Scanning data : %w", err)
			}
			records = append(records, record)
		}
		rows.Close()

		if len(records) == 0 {
			break
		}
		log.Println("Copying", len(records), "rows")
		jsonData, err := json.Marshal(records)
		if err != nil {
			return fmt.Errorf("Marshalling data : %w", err)

		}

		_, err = destDB.Exec(`
		INSERT INTO public.importdata(
	id, created_at, created_by, updated_at, updated_by, deleted_at, tenant, searchindex, name, description, data)
	VALUES (DEFAULT, DEFAULT, '', DEFAULT, '', DEFAULT, '', $1, $2, '', $3)`, batchName, batchName, string(jsonData))
		if err != nil {
			return fmt.Errorf("Inserting data : %w", err)
		}

		offset += batchSize

		if err != nil {
			err = tx.Rollback()
			return fmt.Errorf("Rolling back batch %s: %w", batchName, err)
		}

	}

	log.Println("Committing batch", batchName)
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Committing batch %s: %w", batchName, err)
	}

	return nil
}

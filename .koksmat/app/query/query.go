package query

import (
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

func QueryGetJSON(database string, sqlStatement string) (*json.RawMessage, error) {
	connection, err := GetConnectionString(database)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", *connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlWrappedStatement := fmt.Sprintf(`
	SELECT json_agg(json_data) AS result
	FROM (
		%s
	) AS json_data;
		
	`, sqlStatement)
	rows, err := db.Query(sqlWrappedStatement)
	if err != nil {
		return nil, fmt.Errorf("Reading data using %s fails with : %w", sqlWrappedStatement, err)
	}

	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.Data); err != nil {
			return nil, fmt.Errorf("Scanning data : %w", err)
		}
		records = append(records, record)
	}
	if len(records) != 1 {
		return nil, fmt.Errorf("Expected 1 record, got %d", len(records))
	}
	rows.Close()

	//res := json.Unmarshal(records[0].Data, &records)
	return &records[0].Data, nil

}

/*

sql := fmt.Sprintf("SELECT  row_to_json(d) as data FROM %s as d WHERE batchname = '%s' LIMIT %d OFFSET %d ", sourceTable, batchName, batchSize, offset)
	log.Println("Executing", sql)
	rows, err := db.Query(sql)
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

	// if storedProcedure != nil {
	// 	log.Println("Executing stored procedure", *storedProcedure)
	// 	storedProcedureSQL := fmt.Sprintf("CALL %s()", *storedProcedure)
	// 	_, err = destDB.Exec(storedProcedureSQL)

	// 	if err != nil {
	// 		err = tx.Rollback()
	// 		return fmt.Errorf("Error executing stored procedure %s: %w - Rolling back", *storedProcedure, err)
	// 	}
	// }

*/

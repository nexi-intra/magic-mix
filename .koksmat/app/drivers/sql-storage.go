package drivers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// DBStorage implements the flow.Storage interface using a SQL database
type DBStorage struct {
	db *sql.DB
}

// NewDBStorage creates a new DBStorage instance
func NewDBStorage(initdb *sql.DB) *DBStorage {
	return &DBStorage{
		db: initdb,
	}
}

// Save stores flow JSON data in the database
func (s *DBStorage) Save(id string, flowJSON json.RawMessage) error {
	return fmt.Errorf("not implemented")
	// query := "INSERT INTO flows (id, flow_data) VALUES (?, ?) ON DUPLICATE KEY UPDATE flow_data = VALUES(flow_data)"
	// _, err := s.db.Exec(query, id, flowJSON)
	// if err != nil {
	// 	return fmt.Errorf("failed to save flow: %v", err)
	// }
	// return nil
}

// Load retrieves flow JSON data from the database
func (s *DBStorage) Load(id string) (json.RawMessage, error) {
	return nil, fmt.Errorf("not implemented")
	// var flowJSON string
	// query := "SELECT flow_data FROM flows WHERE id = ?"
	// err := s.db.QueryRow(query, id).Scan(&flowJSON)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return "", fmt.Errorf("flow not found")
	// 	}
	// 	return "", fmt.Errorf("failed to load flow: %v", err)
	// }
	// return flowJSON, nil
}

func (s *DBStorage) GetEvents() ([]interface{}, error) {
	var result []interface{}
	// Query the table `koksmat` for records where `data->>'type' = 'event'` and `data->>'processed' IS NULL`
	rows, err := s.db.Query(`
		SELECT id,data FROM koksmat 
		WHERE data->>'type' = 'event' 
		AND data->>'processed' IS NULL
	`)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	// Loop through the rows and print results
	for rows.Next() {
		// Assuming the structure of koksmat includes an ID and JSONB column 'data'
		var id int
		var data interface{}

		err := rows.Scan(&id, &data)
		result = append(result, data)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		fmt.Printf("ID: %d, Data: %s\n", id, data)
		// Set the 'processed' field to the current timestamp
		currentTime := time.Now().Format(time.RFC3339)

		// Update the JSONB column `data` to set 'processed' to the current time
		_, err = s.db.Exec(`
				UPDATE koksmat
				SET data = jsonb_set(data, '{processed}', to_jsonb($2::text), true)
				WHERE id = $1
			`, id, currentTime)
		if err := rows.Err(); err != nil {
			log.Fatalf("Error during update  : %v\n", err)
		}

	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v\n", err)
	}
	return result, nil
}

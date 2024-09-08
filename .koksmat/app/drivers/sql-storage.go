package drivers

import (
	"fmt"
)

// DBStorage implements the flow.Storage interface using a SQL database
type DBStorage struct {
	// db *sql.DB
}

// NewDBStorage creates a new DBStorage instance
func NewDBStorage() *DBStorage {
	return &DBStorage{
		//	db: db
	}
}

// Save stores flow JSON data in the database
func (s *DBStorage) Save(id string, flowJSON string) error {
	return fmt.Errorf("not implemented")
	// query := "INSERT INTO flows (id, flow_data) VALUES (?, ?) ON DUPLICATE KEY UPDATE flow_data = VALUES(flow_data)"
	// _, err := s.db.Exec(query, id, flowJSON)
	// if err != nil {
	// 	return fmt.Errorf("failed to save flow: %v", err)
	// }
	// return nil
}

// Load retrieves flow JSON data from the database
func (s *DBStorage) Load(id string) (string, error) {
	return "", fmt.Errorf("not implemented")
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

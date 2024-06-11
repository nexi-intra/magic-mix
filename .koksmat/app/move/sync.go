package move

import (
	"context"
)

type Data struct {
	ID    int
	Value string
}

// SyncData function to sync data from master to target
func SyncData(ctx context.Context, masterDB, targetDB *DBConnection, fetchQuery, insertQuery, updateQuery string) error {
	var data []Data
	err := masterDB.FetchData(ctx, fetchQuery, &data)
	if err != nil {
		return err
	}

	for _, row := range data {
		rowData := []interface{}{row.ID, row.Value}
		err = targetDB.InsertOrUpdate(ctx, insertQuery, updateQuery, rowData...)
		if err != nil {
			return err
		}
	}

	return nil
}

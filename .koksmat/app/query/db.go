package query

import (
	"context"
)

// FetchData function to get data from the master database
func (db *DBConnection) FetchData(ctx context.Context, query string, dest interface{}) error {
	return db.DB.NewRaw(query).Scan(ctx, dest)
}

// CheckExistence function to check if a row exists in the target database
func (db *DBConnection) CheckExistence(ctx context.Context, query string, args ...interface{}) (bool, error) {
	var exists bool
	err := db.DB.NewRaw(query, args...).Scan(ctx, &exists)
	return exists, err
}

// InsertOrUpdate function to insert or update data in the target database
func (db *DBConnection) InsertOrUpdate(ctx context.Context, insertQuery, updateQuery string, data ...interface{}) error {
	exists, err := db.CheckExistence(ctx, updateQuery, data[0])
	if err != nil {
		return err
	}

	if exists {
		_, err = db.DB.NewRaw(updateQuery, data...).Exec(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err = db.DB.NewRaw(insertQuery, data...).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

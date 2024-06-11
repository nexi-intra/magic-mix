package move

import (
	"context"
	"log"
)

func moveIt() {
	ctx := context.Background()
	connStr1 := "postgres://username:password@dbserver:5432/db1"
	connStr2 := "postgres://username:password@dbserver:5432/db2"

	masterDB, err := Connect(connStr1)
	if err != nil {
		log.Fatalf("Unable to connect to db1: %v\n", err)
	}
	defer masterDB.Close()

	targetDB, err := Connect(connStr2)
	if err != nil {
		log.Fatalf("Unable to connect to db2: %v\n", err)
	}
	defer targetDB.Close()

	fetchQuery := "SELECT id, value FROM your_table"
	insertQuery := "INSERT INTO your_table (id, value) VALUES ($1, $2)"
	updateQuery := "UPDATE your_table SET value=$1 WHERE id=$2"

	err = SyncData(ctx, masterDB, targetDB, fetchQuery, insertQuery, updateQuery)
	if err != nil {
		log.Fatalf("Data synchronization failed: %v\n", err)
	}

	log.Println("Data synchronization completed successfully!")
}

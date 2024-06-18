package magicapp

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/move"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterMoveCmd() {
	moveCmd := &cobra.Command{
		Use:   "move [from] [to] [source] [destination] [storedprocedure]",
		Short: "Move batch from one database to another",
		Example: `
magic-mix move mix files sharepoint.pageviews events insert_audit_records	

		`,
		Long: `
Move Batch from One Database to Another

This command is used to move batches of records from one database table to another.

Required Parameters:

from: The source database connection string.
to: The destination database connection string.
source: The source table name.
destination: The destination table name.
storedprocedure (optional): The stored procedure to execute at the destination after moving the batch.

Batch Identification: Both the source and destination tables must have a column named batchname. This column is used to identify the batch of records that are being moved.

Batch Movement Criteria: The destination table is queried for distinct batchnames. The source table is queried for distinct batchnames. If the source table has a batchname that is not present in the destination table, that batch is moved.

Batch Size: Data is always moved in batches of a maximum of 10,000 records.

Destination Table Requirement: Data will always be moved to a destination table named importdata. If the specified destination table (parameter 4) is not named importdata, the task will fail unless a stored procedure (parameter 5) is provided. The purpose of the stored procedure is to take data from the importdata table and move it to the final destination table.


`,
		Args: cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			from := args[0]
			to := args[1]
			source := args[2]
			destination := args[3]
			var storedprocedure string
			if len(args) > 4 {
				storedprocedure = args[4]
			}
			if (strings.ToLower(destination) != "importdata") && (len(args) == 4) {
				log.Fatal("Destination table different from importdata requires a stored procedure to be provided")
			}
			OpenDatabase()
			defer utils.Db.Close()
			err := move.Move(from, to, source, destination, &storedprocedure, nil)
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	utils.RootCmd.AddCommand(moveCmd)
}

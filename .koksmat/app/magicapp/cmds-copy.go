package magicapp

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/move"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterCopyCmd() {
	copyCmd := &cobra.Command{
		Use:   "copy [from] [to] [source] [destination] [storedprocedure]",
		Short: "Copy from one database to another",
		Example: `
magic-mix copy mix files sharepoint.pageviews

		`,
		Long: `
Copy data from one Database to Another

This command is used to move batches of records from one database table to another.

Required Parameters:

from: The source database connection string.
to: The destination database connection string.
source: The source table (or view) name.


Page Size: Data is always moved in batches of a maximum of 10,000 records.

Destination Table Requirement: Data will always be copy to a destination table named importdata, the records will be named 'sourcetable|tablename' . If the specified destination table (parameter 4) is not named importdata, the task will fail unless a stored procedure (parameter 5) is provided. The purpose of the stored procedure is to take data from the importdata table and move it to the final destination table.


`,
		Args: cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			from := args[0]
			to := args[1]
			source := args[2]
			batchName := from + "|" + source
			OpenDatabase()
			defer utils.Db.Close()
			err := move.Copy(from, to, source, batchName)
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	utils.RootCmd.AddCommand(copyCmd)
}

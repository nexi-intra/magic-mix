package magicapp

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect"
	"github.com/magicbutton/magic-mix/utils"

	"github.com/magicbutton/magic-mix/move"
)

func RegisterUploadCmd() {
	natsCmd := &cobra.Command{
		Use:   "upload [sourcefolder] [targetdatabase]",
		Short: "Upload a batch",
		Example: `
magic-mix upload users meetings

		`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			batchID := args[0]
			if len(args) > 1 {
				OpenDatabase()
				defer utils.Db.Close()

				toConnection, err := move.GetConnectionString(args[1])
				if err != nil {
					log.Fatal(err)
				}
				db, err := sql.Open("postgres", *toConnection)
				if err != nil {
					log.Fatal(err)
				}
				collect.UploadBatch2(batchID, db)

				defer db.Close()
			} else {
				OpenDatabase()
				defer utils.Db.Close()
				collect.UploadBatch(batchID)
			}
		},
	}
	utils.RootCmd.AddCommand(natsCmd)
}

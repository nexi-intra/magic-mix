package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterUploadCmd() {
	natsCmd := &cobra.Command{
		Use:   "upload [sourcefolder] ",
		Short: "Upload a batch",
		Example: `
magic-mix upload users	

		`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			batchID := args[0]
			OpenDatabase()
			defer utils.Db.Close()
			collect.UploadBatch(batchID)

		},
	}
	utils.RootCmd.AddCommand(natsCmd)
}

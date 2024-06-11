package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterCombineCmd() {
	natsCmd := &cobra.Command{
		Use:   "combine sourcefolder prefix [delete]",
		Short: "Combina a number of json files into one",
		Example: `
magic-mix combine groups owners	

will combine all files in the "groups" folder into one file named "owners" containing an array of the content of all files
matching owners*.json. If the third argument is "delete" the files will be deleted after combining.

		`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var deleteFiles bool = false
			batchID := args[0]
			prefix := args[1]
			if len(args) == 3 {
				deleteFiles = (args[2] == "delete")
			}
			collect.CombineJsonFiles(batchID, prefix, deleteFiles)

		},
	}
	utils.RootCmd.AddCommand(natsCmd)
}

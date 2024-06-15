package magicapp

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/move"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterMoveCmd() {
	moveCmd := &cobra.Command{
		Use:   "move [from] [to] [source] [destination]",
		Short: "Move batch from one database to another",
		Example: `
magic-mix move mix files sharepoint.pageviews events	

		`,
		Args: cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			from := args[0]
			to := args[1]
			source := args[2]
			destination := args[3]
			OpenDatabase()
			defer utils.Db.Close()
			err := move.Move(from, to, source, destination, nil)
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	utils.RootCmd.AddCommand(moveCmd)
}

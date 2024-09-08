package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/utils"
)

func RegisterFlowCmd() {
	natsCmd := &cobra.Command{
		Use:   "flow",
		Short: "Start the Flow Service responder",

		Run: func(cmd *cobra.Command, args []string) {
			StartFlowService()
		},
	}
	utils.RootCmd.AddCommand(natsCmd)
}

package magicapp

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/collect/officegraph"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterParserCmd() {
	parseCmd := &cobra.Command{
		Use: "parse ",
	}
	parseFile := &cobra.Command{
		Use:   "file [format] [batchfile.json]",
		Short: "Runs a parser",
		Example: `


		`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			//parseFormat := args[0]
			inputFile := strings.ToLower(args[1])
			officegraph.ParseMeetings(inputFile)

		},
	}

	parseCmd.AddCommand(parseFile)

	utils.RootCmd.AddCommand(parseCmd)
}

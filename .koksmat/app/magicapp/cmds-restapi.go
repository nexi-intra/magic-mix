package magicapp

import (
	"fmt"
	"log"

	"github.com/magicbutton/magic-mix/restapi"
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/utils"
)

func RegisterRestAPICmd() {
	apiManagementCmd := &cobra.Command{
		Use:   "api",
		Short: "Manage API",
	}
	idTokensCmd := &cobra.Command{
		Use:   "idtoken",
		Short: "Manage ID Tokens",
	}

	issueIdTokenCmd := &cobra.Command{
		Use: "issue AppId Permissions",

		Example: `magic-mix api idtoken issue 1 *`,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			result, err := restapi.IssueIdToken(args[0], args[1])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(result)

		},
	}
	apiManagementCmd.AddCommand(idTokensCmd)
	idTokensCmd.AddCommand(issueIdTokenCmd)

	utils.RootCmd.AddCommand(apiManagementCmd)
}

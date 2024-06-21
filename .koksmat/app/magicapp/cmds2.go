package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-mix/cmds"
	"github.com/magicbutton/magic-mix/utils"
)

func RegisterCmds() {
	utils.RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  ``,
	}
	HealthPingPostCmd := &cobra.Command{
		Use:   "ping  pong",
		Short: "Ping",
		Long:  `Simple ping endpoint`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthPingPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthPingPostCmd)
	HealthCoreversionPostCmd := &cobra.Command{
		Use:   "coreversion ",
		Short: "Core Version",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthCoreversionPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthCoreversionPostCmd)

	utils.RootCmd.AddCommand(healthCmd)
	extractCmd := &cobra.Command{
		Use:   "extract",
		Short: "Extract",
		Long:  ``,
	}
	ExtractHubsiteSpokesPagesPostCmd := &cobra.Command{
		Use:   "hubsite-spokes-pages ",
		Short: "Get Hub Site Spokes Pages",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ExtractHubsiteSpokesPagesPost(ctx, args)
		},
	}
	extractCmd.AddCommand(ExtractHubsiteSpokesPagesPostCmd)

	utils.RootCmd.AddCommand(extractCmd)
}

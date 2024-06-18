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
	processCmd := &cobra.Command{
		Use:   "process",
		Short: "Process",
		Long:  ``,
	}
	ProcessAuditPostCmd := &cobra.Command{
		Use:   "audit  year month day",
		Short: "Download audit logs",
		Long:  `Download audit logs from the auditlog service`,
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProcessAuditPost(ctx, args)
		},
	}
	processCmd.AddCommand(ProcessAuditPostCmd)
	ProcessBatchPostCmd := &cobra.Command{
		Use:   "batch  batchname batchdefition",
		Short: "Download audit logs",
		Long:  `Download audit logs from the auditlog service`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProcessBatchPost(ctx, args)
		},
	}
	processCmd.AddCommand(ProcessBatchPostCmd)

	utils.RootCmd.AddCommand(processCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  ``,
	}
	ProvisionAppdeployproductionPostCmd := &cobra.Command{
		Use:   "appdeployproduction ",
		Short: "App deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionAppdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionAppdeployproductionPostCmd)
	ProvisionWebdeployproductionPostCmd := &cobra.Command{
		Use:   "webdeployproduction ",
		Short: "Web deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeployproductionPostCmd)
	ProvisionWebdeploytestPostCmd := &cobra.Command{
		Use:   "webdeploytest ",
		Short: "Web deploy to Test",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeploytestPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeploytestPostCmd)

	utils.RootCmd.AddCommand(provisionCmd)
}

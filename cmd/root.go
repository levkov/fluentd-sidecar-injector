package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is cobra command.
var RootCmd = &cobra.Command{
	Use:           "fluentd-sidecar-injector",
	Short:         "fluentd-sidecar-injector is a webhook server to inject fluentd sidecar",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		serverCmd(),
		versionCmd(),
	)
}

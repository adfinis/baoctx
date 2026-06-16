package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "Show current installed version of target-cli",
	Long:          `Show current installed version of target-cli.`,
	Example:       `baoctx version`,
	Args:          cobra.NoArgs,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Current Version:", version)
		cmd.Println("")

		return nil
	},
}

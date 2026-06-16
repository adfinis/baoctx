package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteOpenbaoCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete removes a context profile",
	Long:    `delete a context with the delete command.`,
	Example: `baoctx delete example`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		nameToDelete := args[0]

		if _, exists := c.OpenBao[args[0]]; exists {
			delete(c.OpenBao, args[0])
			viper.Set("openbao", c.OpenBao)
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Error writing config: %v\n", err)
				return
			}
			fmt.Printf("Deleted OpenBao profile '%s'\n", nameToDelete)
		} else {
			fmt.Printf("OpenBao profile '%s' not found\n", nameToDelete)
		}
	},
}

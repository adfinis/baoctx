package cmd

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/spf13/cobra"
)

var listOpenbaoCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for OpenBao",
	Long:    `list all context profiles for OpenBao using the list command`,
	Example: `baoctx list`,
	Run: func(cmd *cobra.Command, args []string) {
		t := table.New().
			Border(lipgloss.NormalBorder()).
			Headers("Profile Name", "Endpoint")

		for name, e := range c.OpenBao {
			t.Row(name, e.Endpoint)
		}

		fmt.Println(t)
	},
}

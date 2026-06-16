package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listOpenbaoCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all context profiles for OpenBao",
	Long:    `list all context profiles for OpenBao using the list command`,
	Example: `baoctx list`,
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Profile Name",
			"Endpoint",
		})
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})

		for i, e := range c.OpenBao {
			data := []string{
				i,
				e.Endpoint,
			}
			table.Append(data)
		}
		table.Render()
	},
}

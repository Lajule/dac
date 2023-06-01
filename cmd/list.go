package cmd

import (
	"github.com/Lajule/dac/ent"
	"github.com/Lajule/dac/table"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Typing training sessions listing",
		Long:  `Provides a listing of all training sessions.`,
		Run: func(cmd *cobra.Command, args []string) {
			table.Print(cmd.Context())
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

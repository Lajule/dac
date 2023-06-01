package cmd

import (
	"strings"

	"github.com/Lajule/dac/ent"
	"github.com/Lajule/dac/graph"
	"github.com/spf13/cobra"
)

var (
	statistics string

	statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Typing training sessions statistics",
		Long:  `Provides some statistics of your training sessions by drawing graphs.`,
		Run: func(cmd *cobra.Command, args []string) {
			s := &graph.Statistics{
				Fields: strings.Split(statistics, ","),
			}
			s.Plot(cmd.Context())
		},
	}
)

func init() {
	statsCmd.Flags().StringVarP(&statistics, "statistics", "s", "speed,accuracy,progress", "Statistics to display")
	rootCmd.AddCommand(statsCmd)
}

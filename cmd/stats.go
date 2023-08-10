package cmd

import (
	"log"
	"strings"

	"github.com/Lajule/dac/app/graph"
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

			if err := s.Plot(cmd.Context()); err != nil {
				log.Fatalf("failed plotting data: %v", err)
			}
		},
	}
)

func init() {
	statsCmd.Flags().StringVarP(&statistics, "statistics", "s", "speed,accuracy,progress", "Statistics to display")
	rootCmd.AddCommand(statsCmd)
}

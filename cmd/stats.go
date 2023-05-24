package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Lajule/dac/graph"
	"github.com/Lajule/dac/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	statistics string

	statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
			if err != nil {
				log.Fatalf("failed opening connection to sqlite: %v", err)
			}
			defer client.Close()
			if err := client.Schema.Create(context.Background()); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}

			s := graph.Statistics{
				Fields: strings.Split(statistics, ","),
				Client: client,
			}
			s.Plot()
		},
	}
)

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVarP(&statistics, "statistics", "s", "speed,accuracy,progress", "Statistics to display")
}

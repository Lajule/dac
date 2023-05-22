package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Lajule/dac/ent"
	"github.com/guptarohit/asciigraph"
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
			flields := strings.Split(statistics, ",")
			var values []struct {
				Speed    float64 `json:"speed"`
				Accuracy float64 `json:"accuracy"`
				Progress float64 `json:"progress"`
			}
			if err := client.Training.
				Query().
				Select(flields...).
				Scan(context.Background(), &values); err != nil {
				log.Fatalf("failed selecting data: %v", err)
			}
			data := [][]float64{[]float64{}, []float64{}, []float64{}}
			for _, value := range values {
				data[0] = append(data[0], value.Speed)
				data[1] = append(data[1], value.Accuracy)
				data[2] = append(data[2], value.Progress)
			}
			graph := asciigraph.PlotMany(data, asciigraph.Height(10), asciigraph.SeriesColors(
				asciigraph.Blue,
				asciigraph.Orange,
				asciigraph.Cyan,
			))
			fmt.Println(graph)
		},
	}
)

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVarP(&statistics, "statistics", "s", "speed,accuracy,progress", "Statistics to display")
}

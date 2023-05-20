package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
	"github.com/guptarohit/asciigraph"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	dbFile string

	duration time.Duration

	closable bool

	rootCmd = &cobra.Command{
		Use:   "dac",
		Short: "Typing training sessions",
		Long:  `Dac is typing training sessions program, it's help you to improve your typing skills.`,
		Run: func(cmd *cobra.Command, args []string) {
			if closable && duration == 0 {
				log.Fatal("a duration must be defined")
			}
			var input io.Reader
			if len(args) > 0 {
				file, err := os.Open(args[0])
				if err != nil {
					log.Fatal(err)
				}
				input = file
			} else {
				input = os.Stdin
			}
			b, err := io.ReadAll(input)
			if err != nil {
				log.Fatalf("failed reading input: %v", err)
			}
			if len(b) == 0 {
				log.Fatal("input is empty")
			}
			client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
			if err != nil {
				log.Fatalf("failed opening connection to sqlite: %v", err)
			}
			defer client.Close()
			if err := client.Schema.Create(context.Background()); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}
			t := client.Training.Create().
				SetDuration(duration.Seconds()).
				SetClosable(closable)
			d, err := app.NewDac(string(b))
			if err != nil {
				log.Fatalf("failed creating app: %v", err)
			}
			d.Start(t.Mutation())
			if _, err := t.Save(context.Background()); err != nil {
				log.Fatalf("failed updating training: %v", err)
			}
			speeds, err := client.Training.
				Query().
				Select("speed").
				Float64s(context.Background())
			if err != nil {
				log.Fatalf("failed selecting speed: %v", err)
			}
			accuracies, err := client.Training.
				Query().
				Select("accuracy").
				Float64s(context.Background())
			if err != nil {
				log.Fatalf("failed selecting accuracy: %v", err)
			}
			graph := asciigraph.PlotMany([][]float64{speeds, accuracies}, asciigraph.Height(10), asciigraph.SeriesColors(
				asciigraph.Black,
				asciigraph.Blue,
			))
			fmt.Println(graph)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed executing root command: %v", err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbFile, "database", "dac.db", "Database file (default is dac.db)")
	rootCmd.Flags().DurationVarP(&duration, "duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolVarP(&closable, "closable", "c", false, "Close on session timeout")
}

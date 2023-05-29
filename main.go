package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
	"github.com/Lajule/dac/graph"
	"github.com/Lajule/dac/table"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	client *ent.Client

	duration time.Duration

	closable bool

	statistic string

	statistics string

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

			d, err := app.NewDac(string(b))
			if err != nil {
				log.Fatalf("failed creating app: %v", err)
			}

			t := client.Training.Create().
				SetDuration(duration.Seconds()).
				SetClosable(closable)

			d.Start(t.Mutation())

			if _, err := t.Save(context.Background()); err != nil {
				log.Fatalf("failed updating training: %v", err)
			}

			s := graph.Statistic{
				Field:  statistic,
				Client: client,
			}
			s.Plot()
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Typing training sessions listing",
		Long:  `Provides a listing of all training sessions.`,
		Run: func(cmd *cobra.Command, args []string) {
			l := table.Listing{
				Client: client,
			}
			l.Print()
		},
	}

	statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Typing training sessions statistics",
		Long:  `Provides some statistics of your training sessions by drawing graphs.`,
		Run: func(cmd *cobra.Command, args []string) {
			s := graph.Statistics{
				Fields: strings.Split(statistics, ","),
				Client: client,
			}
			s.Plot()
		},
	}
)

func init() {
	rootCmd.Flags().DurationVarP(&duration, "duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolVarP(&closable, "closable", "c", false, "Close on session timeout")
	rootCmd.Flags().StringVarP(&statistic, "statistic", "s", "speed", "Statistic to display")
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVarP(&statistics, "statistics", "s", "speed,accuracy,progress", "Statistics to display")
}

func main() {
	var dbFile string
	if val, ok := os.LookupEnv("DAC_DB_FILE"); ok {
		dbFile = val
	} else {
		dbFile = "dac.db"
	}

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed executing root command: %v", err)
	}
}

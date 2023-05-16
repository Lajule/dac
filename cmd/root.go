package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	"os"

	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	dbFile string

	theme string

	duration time.Duration

	close bool

	save bool

	rootCmd = &cobra.Command{
		Use:   "dac",
		Short: "Typing training sessions",
		Long:  `Dac is typing training sessions program, it's help you to improve your typing skills.`,
		Run: func(cmd *cobra.Command, args []string) {
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
				log.Println("input is empty")
				os.Exit(1)
			}

			if close && duration == 0 {
				log.Println("a duration must be defined")
				os.Exit(1)
			}

			client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
			if err != nil {
				log.Fatalf("failed opening connection to sqlite: %v", err)
			}
			defer client.Close()
			if err := client.Schema.Create(context.Background()); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}

			d, err := app.NewDac(duration, close, string(b))
			if err != nil {
				log.Fatalf("failed creating ui: %v", err)
			}

			d.Start()
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
	rootCmd.PersistentFlags().StringVar(&theme, "theme", "green", "Color theme (default is green)")
	rootCmd.Flags().DurationVarP(&duration, "duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolVarP(&close, "close", "c", false, "Close on session timeout")
	rootCmd.Flags().BoolVarP(&save, "save", "s", true, "Save session in database")
}

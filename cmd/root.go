package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	dbFile string

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

			duration, err := cmd.Flags().GetDuration("duration")
			if err != nil {
				log.Fatalf("failed getting duration: %v", err)
			}

			close, err := cmd.Flags().GetBool("close")
			if err != nil {
				log.Fatalf("failed getting close: %v", err)
			}

			if close && duration == 0 {
				log.Println("a duration must be defined")
				os.Exit(1)
			}

			d, err := app.NewDac(duration, close, string(b))
			if err != nil {
				log.Fatalf("failed creating ui: %v", err)
			}

			d.Start()

			save, err := cmd.Flags().GetBool("close")
			if err != nil {
				log.Fatalf("failed getting save: %v", err)
			}

			if save {
				client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", dbFile))
				if err != nil {
					log.Fatalf("failed opening connection to sqlite: %v", err)
				}
				defer client.Close()
				if err := client.Schema.Create(context.Background()); err != nil {
					log.Fatalf("failed creating schema resources: %v", err)
				}
				if _, err := client.Training.
					Create().
					SetDuration(int(d.Duration)).
					SetSpeed(d.Speed).
					SetPrecision(d.Precision).
					Save(context.Background()); err != nil {
					log.Fatalf("failed creating training: %v", err)
				}
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed executing root command: %v", err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbFile, "db", "dac.db", "Database file (default is dac.db)")
	rootCmd.Flags().DurationP("duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolP("close", "c", false, "Close on session timeout")
	rootCmd.Flags().BoolP("save", "s", true, "Save session in database")
}

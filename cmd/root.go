package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/Lajule/dac/app"
	"github.com/Lajule/dac/ent"
	"github.com/Lajule/dac/graph"
	"github.com/spf13/cobra"
)

var (
	duration time.Duration

	closable bool

	statistic string

	rootCmd = &cobra.Command{
		Use:   "dac",
		Short: "Typing training sessions",
		Long:  `Dac is typing training sessions program, it's help you to improve your typing skills.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			client := ctx.Value("client").(*ent.Client)
			t := client.Training.Create().
				SetDuration(duration.Seconds()).
				SetClosable(closable)

			var reader io.Reader
			if len(args) > 0 {
				file, err := os.Open(args[0])
				if err != nil {
					log.Fatal(err)
				}

				t.SetInput(file.Name())
				reader = file
			} else {
				reader = os.Stdin
			}

			b, err := io.ReadAll(reader)
			if err != nil {
				log.Fatalf("failed reading input: %v", err)
			}
			if len(b) == 0 {
				log.Fatal("input is empty")
			}
			t.SetLength(len(b))

			ts, err := app.NewTrainingSession(string(b))
			if err != nil {
				log.Fatalf("failed creating app: %v", err)
			}

			ts.Start(t.Mutation())

			if _, err := t.Save(ctx); err != nil {
				log.Fatalf("failed updating training: %v", err)
			}

			s := &graph.Statistic{
				Field:  statistic,
			}
			s.Plot(ctx)
		},
	}
)

func init() {
	rootCmd.Flags().DurationVarP(&duration, "duration", "d", 0, "Duration of the training session")
	rootCmd.Flags().BoolVarP(&closable, "closable", "c", false, "Close on session timeout")
	rootCmd.Flags().StringVarP(&statistic, "statistic", "s", "speed", "Statistic to display")
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}

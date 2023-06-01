package table

import (
	"context"
	"fmt"

	//"github.com/fatih/color"
	"github.com/gosuri/uitable"

	"github.com/Lajule/dac/ent"
)

func Print(ctx context.Context) error {
	client := ctx.Value("client").(*ent.Client)

	trainings, err := client.Training.
		Query().
		All(context.Background())
	if err != nil {
		return fmt.Errorf("failed selecting data: %w", err)
	}

	table := uitable.New()
	table.AddRow("CREATED_AT", "DURATION", "CLOSABLE", "STOPWATCH", "PROGRESS", "ACCURACY", "SPEED", "INPUT")

	for _, training := range trainings {
		table.AddRow(training.CreatedAt.String(),
			"DURATION",
			"CLOSABLE",
			"STOPWATCH",
			"PROGRESS",
			"ACCURACY",
			"SPEED",
			"INPUT")
	}

	fmt.Println(table)

	return nil
}

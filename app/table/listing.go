package table

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gosuri/uitable"

	dac "github.com/Lajule/dac/context"
)

func Print(ctx context.Context) error {
	val := ctx.Value(dac.Key).(dac.ValueType)

	trainings, err := val.Client.Training.
		Query().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed selecting data: %w", err)
	}

	table := uitable.New()
	table.AddRow("CREATED_AT", "DURATION", "CLOSABLE", "STOPWATCH", "PROGRESS", "ACCURACY", "SPEED", "INPUT", "LENGTH")

	for _, training := range trainings {
		table.AddRow(training.CreatedAt.Format(time.RFC3339),
			time.Duration(time.Duration(training.Duration)*time.Second).String(),
			strconv.FormatBool(training.Closable),
			time.Duration(time.Duration(training.Stopwatch)*time.Second).String(),
			fmt.Sprintf("%.0f%%", training.Progress),
			fmt.Sprintf("%.0f%%", training.Accuracy),
			fmt.Sprintf("%.0fw/m", training.Speed),
			training.Input,
			training.Length)
	}

	fmt.Println(table)

	return nil
}

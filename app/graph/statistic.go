package graph

import (
	"context"
	"fmt"

	"github.com/guptarohit/asciigraph"

	dac "github.com/Lajule/dac/context"
)

type Statistic struct {
	Field string
}

func (s *Statistic) Plot(ctx context.Context) error {
	val := ctx.Value(dac.KeyName).(dac.Value)

	data, err := val.Client.Training.
		Query().
		Select(s.Field).
		Float64s(ctx)
	if err != nil {
		return fmt.Errorf("failed selecting data: %w", err)
	}

	fmt.Println(asciigraph.Plot(data, asciigraph.Height(10), asciigraph.SeriesColors(
		asciigraph.Blue,
	)))

	return nil
}

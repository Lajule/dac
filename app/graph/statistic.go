package graph

import (
	"context"
	"fmt"

	"github.com/Lajule/dac/ent"
	"github.com/guptarohit/asciigraph"
)

type Statistic struct {
	Field string
}

func (s *Statistic) Plot(ctx context.Context) error {
	values := ctx.Value("values").(map[string]any)

	client := values["client"].(*ent.Client)
	data, err := client.Training.
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

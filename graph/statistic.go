package graph

import (
	"context"
	"fmt"

	"github.com/Lajule/dac/ent"
	"github.com/guptarohit/asciigraph"
)

type Statistic struct {
	Field string

	Client *ent.Client
}

func (s *Statistic) Plot() error {
	data, err := s.Client.Training.
		Query().
		Select(s.Field).
		Float64s(context.Background())
	if err != nil {
		return fmt.Errorf("failed selecting data: %w", err)
	}

	fmt.Println(asciigraph.Plot(data, asciigraph.Height(10), asciigraph.SeriesColors(
		asciigraph.Blue,
	)))

	return nil
}

package graph

import (
	"context"
	"fmt"

	"github.com/Lajule/dac/ent"
	"github.com/guptarohit/asciigraph"
)

type Statistics struct {
	Fields []string
}

func (s *Statistics) Plot(ctx context.Context) error {
	client := ctx.Value("client").(*ent.Client)

	var values []struct {
		Speed    float64 `json:"speed"`
		Accuracy float64 `json:"accuracy"`
		Progress float64 `json:"progress"`
	}

	if err := client.Training.
		Query().
		Select(s.Fields...).
		Scan(ctx, &values); err != nil {
		return fmt.Errorf("failed selecting data: %w", err)
	}

	if len(values) > 0 {
		data := [][]float64{[]float64{}, []float64{}, []float64{}}
		for _, value := range values {
			data[0] = append(data[0], value.Speed)
			data[1] = append(data[1], value.Accuracy)
			data[2] = append(data[2], value.Progress)
		}

		fmt.Println(asciigraph.PlotMany(data, asciigraph.Height(10), asciigraph.SeriesColors(
			asciigraph.Blue,
			asciigraph.Orange,
			asciigraph.Cyan,
		)))
	}

	return nil
}

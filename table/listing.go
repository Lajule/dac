package table

import (
	//"context"
	"fmt"

	//"github.com/fatih/color"
	"github.com/gosuri/uitable"

	"github.com/Lajule/dac/ent"
)

type Listing struct {
	Client *ent.Client
}

func (s *Listing) Print() error {
	table := uitable.New()

	table.AddRow("CREATED_AT", "DURATION", "CLOSABLE", "STOPWATCH", "PROGRESS", "ACCURACY", "SPEED", "INPUT")
	fmt.Println(table)

	return nil
}

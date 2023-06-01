package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Training holds the schema definition for the Training entity.
type Training struct {
	ent.Schema
}

// Fields of the Training.
func (Training) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Float("duration").Default(0.0),
		field.Bool("closable").Default(false),
		field.Float("stopwatch").Default(0.0),
		field.Float("progress").Default(0.0),
		field.Float("accuracy").Default(0.0),
		field.Float("speed").Default(0.0),
		field.String("input").Optional(),
	}
}

// Edges of the Training.
func (Training) Edges() []ent.Edge {
	return nil
}

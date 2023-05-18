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
		field.Int("duration").Default(0),
		field.Bool("closable").Default(false),
		field.Int("stopwatch").Default(0),
		field.Int("progress").Default(0),
		field.Int("accuracy").Default(0),
		field.Int("speed").Default(0),
	}
}

// Edges of the Training.
func (Training) Edges() []ent.Edge {
	return nil
}

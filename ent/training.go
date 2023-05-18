// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Lajule/dac/ent/training"
)

// Training is the model entity for the Training schema.
type Training struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Closable holds the value of the "closable" field.
	Closable bool `json:"closable,omitempty"`
	// Stopwatch holds the value of the "stopwatch" field.
	Stopwatch int `json:"stopwatch,omitempty"`
	// Progress holds the value of the "progress" field.
	Progress int `json:"progress,omitempty"`
	// Accuracy holds the value of the "accuracy" field.
	Accuracy int `json:"accuracy,omitempty"`
	// Speed holds the value of the "speed" field.
	Speed        int `json:"speed,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Training) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case training.FieldClosable:
			values[i] = new(sql.NullBool)
		case training.FieldID, training.FieldDuration, training.FieldStopwatch, training.FieldProgress, training.FieldAccuracy, training.FieldSpeed:
			values[i] = new(sql.NullInt64)
		case training.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Training fields.
func (t *Training) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case training.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case training.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case training.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				t.Duration = int(value.Int64)
			}
		case training.FieldClosable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field closable", values[i])
			} else if value.Valid {
				t.Closable = value.Bool
			}
		case training.FieldStopwatch:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stopwatch", values[i])
			} else if value.Valid {
				t.Stopwatch = int(value.Int64)
			}
		case training.FieldProgress:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field progress", values[i])
			} else if value.Valid {
				t.Progress = int(value.Int64)
			}
		case training.FieldAccuracy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field accuracy", values[i])
			} else if value.Valid {
				t.Accuracy = int(value.Int64)
			}
		case training.FieldSpeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field speed", values[i])
			} else if value.Valid {
				t.Speed = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Training.
// This includes values selected through modifiers, order, etc.
func (t *Training) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Training.
// Note that you need to call Training.Unwrap() before calling this method if this Training
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Training) Update() *TrainingUpdateOne {
	return NewTrainingClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Training entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Training) Unwrap() *Training {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Training is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Training) String() string {
	var builder strings.Builder
	builder.WriteString("Training(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", t.Duration))
	builder.WriteString(", ")
	builder.WriteString("closable=")
	builder.WriteString(fmt.Sprintf("%v", t.Closable))
	builder.WriteString(", ")
	builder.WriteString("stopwatch=")
	builder.WriteString(fmt.Sprintf("%v", t.Stopwatch))
	builder.WriteString(", ")
	builder.WriteString("progress=")
	builder.WriteString(fmt.Sprintf("%v", t.Progress))
	builder.WriteString(", ")
	builder.WriteString("accuracy=")
	builder.WriteString(fmt.Sprintf("%v", t.Accuracy))
	builder.WriteString(", ")
	builder.WriteString("speed=")
	builder.WriteString(fmt.Sprintf("%v", t.Speed))
	builder.WriteByte(')')
	return builder.String()
}

// Trainings is a parsable slice of Training.
type Trainings []*Training

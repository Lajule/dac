// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TrainingsColumns holds the columns for the "trainings" table.
	TrainingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "duration", Type: field.TypeFloat64, Default: 0},
		{Name: "closable", Type: field.TypeBool, Default: false},
		{Name: "stopwatch", Type: field.TypeFloat64, Default: 0},
		{Name: "progress", Type: field.TypeFloat64, Default: 0},
		{Name: "accuracy", Type: field.TypeFloat64, Default: 0},
		{Name: "speed", Type: field.TypeFloat64, Default: 0},
		{Name: "input", Type: field.TypeString, Nullable: true},
		{Name: "length", Type: field.TypeInt, Default: 0},
	}
	// TrainingsTable holds the schema information for the "trainings" table.
	TrainingsTable = &schema.Table{
		Name:       "trainings",
		Columns:    TrainingsColumns,
		PrimaryKey: []*schema.Column{TrainingsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TrainingsTable,
	}
)

func init() {
}

// Code generated by ent, DO NOT EDIT.

package training

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the training type in the database.
	Label = "training"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldClosable holds the string denoting the closable field in the database.
	FieldClosable = "closable"
	// FieldStopwatch holds the string denoting the stopwatch field in the database.
	FieldStopwatch = "stopwatch"
	// FieldProgress holds the string denoting the progress field in the database.
	FieldProgress = "progress"
	// FieldAccuracy holds the string denoting the accuracy field in the database.
	FieldAccuracy = "accuracy"
	// FieldSpeed holds the string denoting the speed field in the database.
	FieldSpeed = "speed"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// Table holds the table name of the training in the database.
	Table = "trainings"
)

// Columns holds all SQL columns for training fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldDuration,
	FieldClosable,
	FieldStopwatch,
	FieldProgress,
	FieldAccuracy,
	FieldSpeed,
	FieldInput,
	FieldLength,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration float64
	// DefaultClosable holds the default value on creation for the "closable" field.
	DefaultClosable bool
	// DefaultStopwatch holds the default value on creation for the "stopwatch" field.
	DefaultStopwatch float64
	// DefaultProgress holds the default value on creation for the "progress" field.
	DefaultProgress float64
	// DefaultAccuracy holds the default value on creation for the "accuracy" field.
	DefaultAccuracy float64
	// DefaultSpeed holds the default value on creation for the "speed" field.
	DefaultSpeed float64
	// DefaultLength holds the default value on creation for the "length" field.
	DefaultLength int
)

// OrderOption defines the ordering options for the Training queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByClosable orders the results by the closable field.
func ByClosable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClosable, opts...).ToFunc()
}

// ByStopwatch orders the results by the stopwatch field.
func ByStopwatch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStopwatch, opts...).ToFunc()
}

// ByProgress orders the results by the progress field.
func ByProgress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgress, opts...).ToFunc()
}

// ByAccuracy orders the results by the accuracy field.
func ByAccuracy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccuracy, opts...).ToFunc()
}

// BySpeed orders the results by the speed field.
func BySpeed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpeed, opts...).ToFunc()
}

// ByInput orders the results by the input field.
func ByInput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInput, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

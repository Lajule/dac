// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Lajule/dac/ent/predicate"
	"github.com/Lajule/dac/ent/training"
)

// TrainingUpdate is the builder for updating Training entities.
type TrainingUpdate struct {
	config
	hooks    []Hook
	mutation *TrainingMutation
}

// Where appends a list predicates to the TrainingUpdate builder.
func (tu *TrainingUpdate) Where(ps ...predicate.Training) *TrainingUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDuration sets the "duration" field.
func (tu *TrainingUpdate) SetDuration(f float64) *TrainingUpdate {
	tu.mutation.ResetDuration()
	tu.mutation.SetDuration(f)
	return tu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableDuration(f *float64) *TrainingUpdate {
	if f != nil {
		tu.SetDuration(*f)
	}
	return tu
}

// AddDuration adds f to the "duration" field.
func (tu *TrainingUpdate) AddDuration(f float64) *TrainingUpdate {
	tu.mutation.AddDuration(f)
	return tu
}

// SetClosable sets the "closable" field.
func (tu *TrainingUpdate) SetClosable(b bool) *TrainingUpdate {
	tu.mutation.SetClosable(b)
	return tu
}

// SetNillableClosable sets the "closable" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableClosable(b *bool) *TrainingUpdate {
	if b != nil {
		tu.SetClosable(*b)
	}
	return tu
}

// SetStopwatch sets the "stopwatch" field.
func (tu *TrainingUpdate) SetStopwatch(f float64) *TrainingUpdate {
	tu.mutation.ResetStopwatch()
	tu.mutation.SetStopwatch(f)
	return tu
}

// SetNillableStopwatch sets the "stopwatch" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableStopwatch(f *float64) *TrainingUpdate {
	if f != nil {
		tu.SetStopwatch(*f)
	}
	return tu
}

// AddStopwatch adds f to the "stopwatch" field.
func (tu *TrainingUpdate) AddStopwatch(f float64) *TrainingUpdate {
	tu.mutation.AddStopwatch(f)
	return tu
}

// SetProgress sets the "progress" field.
func (tu *TrainingUpdate) SetProgress(f float64) *TrainingUpdate {
	tu.mutation.ResetProgress()
	tu.mutation.SetProgress(f)
	return tu
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableProgress(f *float64) *TrainingUpdate {
	if f != nil {
		tu.SetProgress(*f)
	}
	return tu
}

// AddProgress adds f to the "progress" field.
func (tu *TrainingUpdate) AddProgress(f float64) *TrainingUpdate {
	tu.mutation.AddProgress(f)
	return tu
}

// SetAccuracy sets the "accuracy" field.
func (tu *TrainingUpdate) SetAccuracy(f float64) *TrainingUpdate {
	tu.mutation.ResetAccuracy()
	tu.mutation.SetAccuracy(f)
	return tu
}

// SetNillableAccuracy sets the "accuracy" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableAccuracy(f *float64) *TrainingUpdate {
	if f != nil {
		tu.SetAccuracy(*f)
	}
	return tu
}

// AddAccuracy adds f to the "accuracy" field.
func (tu *TrainingUpdate) AddAccuracy(f float64) *TrainingUpdate {
	tu.mutation.AddAccuracy(f)
	return tu
}

// SetSpeed sets the "speed" field.
func (tu *TrainingUpdate) SetSpeed(f float64) *TrainingUpdate {
	tu.mutation.ResetSpeed()
	tu.mutation.SetSpeed(f)
	return tu
}

// SetNillableSpeed sets the "speed" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableSpeed(f *float64) *TrainingUpdate {
	if f != nil {
		tu.SetSpeed(*f)
	}
	return tu
}

// AddSpeed adds f to the "speed" field.
func (tu *TrainingUpdate) AddSpeed(f float64) *TrainingUpdate {
	tu.mutation.AddSpeed(f)
	return tu
}

// SetInput sets the "input" field.
func (tu *TrainingUpdate) SetInput(s string) *TrainingUpdate {
	tu.mutation.SetInput(s)
	return tu
}

// Mutation returns the TrainingMutation object of the builder.
func (tu *TrainingUpdate) Mutation() *TrainingMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TrainingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TrainingUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TrainingUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TrainingUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TrainingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(training.Table, training.Columns, sqlgraph.NewFieldSpec(training.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Duration(); ok {
		_spec.SetField(training.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedDuration(); ok {
		_spec.AddField(training.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Closable(); ok {
		_spec.SetField(training.FieldClosable, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Stopwatch(); ok {
		_spec.SetField(training.FieldStopwatch, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedStopwatch(); ok {
		_spec.AddField(training.FieldStopwatch, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Progress(); ok {
		_spec.SetField(training.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedProgress(); ok {
		_spec.AddField(training.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Accuracy(); ok {
		_spec.SetField(training.FieldAccuracy, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedAccuracy(); ok {
		_spec.AddField(training.FieldAccuracy, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Speed(); ok {
		_spec.SetField(training.FieldSpeed, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedSpeed(); ok {
		_spec.AddField(training.FieldSpeed, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.Input(); ok {
		_spec.SetField(training.FieldInput, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{training.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TrainingUpdateOne is the builder for updating a single Training entity.
type TrainingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrainingMutation
}

// SetDuration sets the "duration" field.
func (tuo *TrainingUpdateOne) SetDuration(f float64) *TrainingUpdateOne {
	tuo.mutation.ResetDuration()
	tuo.mutation.SetDuration(f)
	return tuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableDuration(f *float64) *TrainingUpdateOne {
	if f != nil {
		tuo.SetDuration(*f)
	}
	return tuo
}

// AddDuration adds f to the "duration" field.
func (tuo *TrainingUpdateOne) AddDuration(f float64) *TrainingUpdateOne {
	tuo.mutation.AddDuration(f)
	return tuo
}

// SetClosable sets the "closable" field.
func (tuo *TrainingUpdateOne) SetClosable(b bool) *TrainingUpdateOne {
	tuo.mutation.SetClosable(b)
	return tuo
}

// SetNillableClosable sets the "closable" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableClosable(b *bool) *TrainingUpdateOne {
	if b != nil {
		tuo.SetClosable(*b)
	}
	return tuo
}

// SetStopwatch sets the "stopwatch" field.
func (tuo *TrainingUpdateOne) SetStopwatch(f float64) *TrainingUpdateOne {
	tuo.mutation.ResetStopwatch()
	tuo.mutation.SetStopwatch(f)
	return tuo
}

// SetNillableStopwatch sets the "stopwatch" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableStopwatch(f *float64) *TrainingUpdateOne {
	if f != nil {
		tuo.SetStopwatch(*f)
	}
	return tuo
}

// AddStopwatch adds f to the "stopwatch" field.
func (tuo *TrainingUpdateOne) AddStopwatch(f float64) *TrainingUpdateOne {
	tuo.mutation.AddStopwatch(f)
	return tuo
}

// SetProgress sets the "progress" field.
func (tuo *TrainingUpdateOne) SetProgress(f float64) *TrainingUpdateOne {
	tuo.mutation.ResetProgress()
	tuo.mutation.SetProgress(f)
	return tuo
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableProgress(f *float64) *TrainingUpdateOne {
	if f != nil {
		tuo.SetProgress(*f)
	}
	return tuo
}

// AddProgress adds f to the "progress" field.
func (tuo *TrainingUpdateOne) AddProgress(f float64) *TrainingUpdateOne {
	tuo.mutation.AddProgress(f)
	return tuo
}

// SetAccuracy sets the "accuracy" field.
func (tuo *TrainingUpdateOne) SetAccuracy(f float64) *TrainingUpdateOne {
	tuo.mutation.ResetAccuracy()
	tuo.mutation.SetAccuracy(f)
	return tuo
}

// SetNillableAccuracy sets the "accuracy" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableAccuracy(f *float64) *TrainingUpdateOne {
	if f != nil {
		tuo.SetAccuracy(*f)
	}
	return tuo
}

// AddAccuracy adds f to the "accuracy" field.
func (tuo *TrainingUpdateOne) AddAccuracy(f float64) *TrainingUpdateOne {
	tuo.mutation.AddAccuracy(f)
	return tuo
}

// SetSpeed sets the "speed" field.
func (tuo *TrainingUpdateOne) SetSpeed(f float64) *TrainingUpdateOne {
	tuo.mutation.ResetSpeed()
	tuo.mutation.SetSpeed(f)
	return tuo
}

// SetNillableSpeed sets the "speed" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableSpeed(f *float64) *TrainingUpdateOne {
	if f != nil {
		tuo.SetSpeed(*f)
	}
	return tuo
}

// AddSpeed adds f to the "speed" field.
func (tuo *TrainingUpdateOne) AddSpeed(f float64) *TrainingUpdateOne {
	tuo.mutation.AddSpeed(f)
	return tuo
}

// SetInput sets the "input" field.
func (tuo *TrainingUpdateOne) SetInput(s string) *TrainingUpdateOne {
	tuo.mutation.SetInput(s)
	return tuo
}

// Mutation returns the TrainingMutation object of the builder.
func (tuo *TrainingUpdateOne) Mutation() *TrainingMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TrainingUpdate builder.
func (tuo *TrainingUpdateOne) Where(ps ...predicate.Training) *TrainingUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TrainingUpdateOne) Select(field string, fields ...string) *TrainingUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Training entity.
func (tuo *TrainingUpdateOne) Save(ctx context.Context) (*Training, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TrainingUpdateOne) SaveX(ctx context.Context) *Training {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TrainingUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TrainingUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TrainingUpdateOne) sqlSave(ctx context.Context) (_node *Training, err error) {
	_spec := sqlgraph.NewUpdateSpec(training.Table, training.Columns, sqlgraph.NewFieldSpec(training.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Training.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, training.FieldID)
		for _, f := range fields {
			if !training.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != training.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Duration(); ok {
		_spec.SetField(training.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedDuration(); ok {
		_spec.AddField(training.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Closable(); ok {
		_spec.SetField(training.FieldClosable, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Stopwatch(); ok {
		_spec.SetField(training.FieldStopwatch, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedStopwatch(); ok {
		_spec.AddField(training.FieldStopwatch, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Progress(); ok {
		_spec.SetField(training.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedProgress(); ok {
		_spec.AddField(training.FieldProgress, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Accuracy(); ok {
		_spec.SetField(training.FieldAccuracy, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedAccuracy(); ok {
		_spec.AddField(training.FieldAccuracy, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Speed(); ok {
		_spec.SetField(training.FieldSpeed, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedSpeed(); ok {
		_spec.AddField(training.FieldSpeed, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.Input(); ok {
		_spec.SetField(training.FieldInput, field.TypeString, value)
	}
	_node = &Training{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{training.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}

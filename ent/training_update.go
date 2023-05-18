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
func (tu *TrainingUpdate) SetDuration(i int) *TrainingUpdate {
	tu.mutation.ResetDuration()
	tu.mutation.SetDuration(i)
	return tu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableDuration(i *int) *TrainingUpdate {
	if i != nil {
		tu.SetDuration(*i)
	}
	return tu
}

// AddDuration adds i to the "duration" field.
func (tu *TrainingUpdate) AddDuration(i int) *TrainingUpdate {
	tu.mutation.AddDuration(i)
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
func (tu *TrainingUpdate) SetStopwatch(i int) *TrainingUpdate {
	tu.mutation.ResetStopwatch()
	tu.mutation.SetStopwatch(i)
	return tu
}

// SetNillableStopwatch sets the "stopwatch" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableStopwatch(i *int) *TrainingUpdate {
	if i != nil {
		tu.SetStopwatch(*i)
	}
	return tu
}

// AddStopwatch adds i to the "stopwatch" field.
func (tu *TrainingUpdate) AddStopwatch(i int) *TrainingUpdate {
	tu.mutation.AddStopwatch(i)
	return tu
}

// SetProgress sets the "progress" field.
func (tu *TrainingUpdate) SetProgress(i int) *TrainingUpdate {
	tu.mutation.ResetProgress()
	tu.mutation.SetProgress(i)
	return tu
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableProgress(i *int) *TrainingUpdate {
	if i != nil {
		tu.SetProgress(*i)
	}
	return tu
}

// AddProgress adds i to the "progress" field.
func (tu *TrainingUpdate) AddProgress(i int) *TrainingUpdate {
	tu.mutation.AddProgress(i)
	return tu
}

// SetAccuracy sets the "accuracy" field.
func (tu *TrainingUpdate) SetAccuracy(i int) *TrainingUpdate {
	tu.mutation.ResetAccuracy()
	tu.mutation.SetAccuracy(i)
	return tu
}

// SetNillableAccuracy sets the "accuracy" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableAccuracy(i *int) *TrainingUpdate {
	if i != nil {
		tu.SetAccuracy(*i)
	}
	return tu
}

// AddAccuracy adds i to the "accuracy" field.
func (tu *TrainingUpdate) AddAccuracy(i int) *TrainingUpdate {
	tu.mutation.AddAccuracy(i)
	return tu
}

// SetSpeed sets the "speed" field.
func (tu *TrainingUpdate) SetSpeed(i int) *TrainingUpdate {
	tu.mutation.ResetSpeed()
	tu.mutation.SetSpeed(i)
	return tu
}

// SetNillableSpeed sets the "speed" field if the given value is not nil.
func (tu *TrainingUpdate) SetNillableSpeed(i *int) *TrainingUpdate {
	if i != nil {
		tu.SetSpeed(*i)
	}
	return tu
}

// AddSpeed adds i to the "speed" field.
func (tu *TrainingUpdate) AddSpeed(i int) *TrainingUpdate {
	tu.mutation.AddSpeed(i)
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
		_spec.SetField(training.FieldDuration, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedDuration(); ok {
		_spec.AddField(training.FieldDuration, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Closable(); ok {
		_spec.SetField(training.FieldClosable, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Stopwatch(); ok {
		_spec.SetField(training.FieldStopwatch, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedStopwatch(); ok {
		_spec.AddField(training.FieldStopwatch, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Progress(); ok {
		_spec.SetField(training.FieldProgress, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedProgress(); ok {
		_spec.AddField(training.FieldProgress, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Accuracy(); ok {
		_spec.SetField(training.FieldAccuracy, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedAccuracy(); ok {
		_spec.AddField(training.FieldAccuracy, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Speed(); ok {
		_spec.SetField(training.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedSpeed(); ok {
		_spec.AddField(training.FieldSpeed, field.TypeInt, value)
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
func (tuo *TrainingUpdateOne) SetDuration(i int) *TrainingUpdateOne {
	tuo.mutation.ResetDuration()
	tuo.mutation.SetDuration(i)
	return tuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableDuration(i *int) *TrainingUpdateOne {
	if i != nil {
		tuo.SetDuration(*i)
	}
	return tuo
}

// AddDuration adds i to the "duration" field.
func (tuo *TrainingUpdateOne) AddDuration(i int) *TrainingUpdateOne {
	tuo.mutation.AddDuration(i)
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
func (tuo *TrainingUpdateOne) SetStopwatch(i int) *TrainingUpdateOne {
	tuo.mutation.ResetStopwatch()
	tuo.mutation.SetStopwatch(i)
	return tuo
}

// SetNillableStopwatch sets the "stopwatch" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableStopwatch(i *int) *TrainingUpdateOne {
	if i != nil {
		tuo.SetStopwatch(*i)
	}
	return tuo
}

// AddStopwatch adds i to the "stopwatch" field.
func (tuo *TrainingUpdateOne) AddStopwatch(i int) *TrainingUpdateOne {
	tuo.mutation.AddStopwatch(i)
	return tuo
}

// SetProgress sets the "progress" field.
func (tuo *TrainingUpdateOne) SetProgress(i int) *TrainingUpdateOne {
	tuo.mutation.ResetProgress()
	tuo.mutation.SetProgress(i)
	return tuo
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableProgress(i *int) *TrainingUpdateOne {
	if i != nil {
		tuo.SetProgress(*i)
	}
	return tuo
}

// AddProgress adds i to the "progress" field.
func (tuo *TrainingUpdateOne) AddProgress(i int) *TrainingUpdateOne {
	tuo.mutation.AddProgress(i)
	return tuo
}

// SetAccuracy sets the "accuracy" field.
func (tuo *TrainingUpdateOne) SetAccuracy(i int) *TrainingUpdateOne {
	tuo.mutation.ResetAccuracy()
	tuo.mutation.SetAccuracy(i)
	return tuo
}

// SetNillableAccuracy sets the "accuracy" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableAccuracy(i *int) *TrainingUpdateOne {
	if i != nil {
		tuo.SetAccuracy(*i)
	}
	return tuo
}

// AddAccuracy adds i to the "accuracy" field.
func (tuo *TrainingUpdateOne) AddAccuracy(i int) *TrainingUpdateOne {
	tuo.mutation.AddAccuracy(i)
	return tuo
}

// SetSpeed sets the "speed" field.
func (tuo *TrainingUpdateOne) SetSpeed(i int) *TrainingUpdateOne {
	tuo.mutation.ResetSpeed()
	tuo.mutation.SetSpeed(i)
	return tuo
}

// SetNillableSpeed sets the "speed" field if the given value is not nil.
func (tuo *TrainingUpdateOne) SetNillableSpeed(i *int) *TrainingUpdateOne {
	if i != nil {
		tuo.SetSpeed(*i)
	}
	return tuo
}

// AddSpeed adds i to the "speed" field.
func (tuo *TrainingUpdateOne) AddSpeed(i int) *TrainingUpdateOne {
	tuo.mutation.AddSpeed(i)
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
		_spec.SetField(training.FieldDuration, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedDuration(); ok {
		_spec.AddField(training.FieldDuration, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Closable(); ok {
		_spec.SetField(training.FieldClosable, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Stopwatch(); ok {
		_spec.SetField(training.FieldStopwatch, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedStopwatch(); ok {
		_spec.AddField(training.FieldStopwatch, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Progress(); ok {
		_spec.SetField(training.FieldProgress, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedProgress(); ok {
		_spec.AddField(training.FieldProgress, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Accuracy(); ok {
		_spec.SetField(training.FieldAccuracy, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedAccuracy(); ok {
		_spec.AddField(training.FieldAccuracy, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Speed(); ok {
		_spec.SetField(training.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedSpeed(); ok {
		_spec.AddField(training.FieldSpeed, field.TypeInt, value)
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

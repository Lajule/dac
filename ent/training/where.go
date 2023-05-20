// Code generated by ent, DO NOT EDIT.

package training

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Lajule/dac/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldCreatedAt, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldDuration, v))
}

// Closable applies equality check predicate on the "closable" field. It's identical to ClosableEQ.
func Closable(v bool) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldClosable, v))
}

// Stopwatch applies equality check predicate on the "stopwatch" field. It's identical to StopwatchEQ.
func Stopwatch(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldStopwatch, v))
}

// Progress applies equality check predicate on the "progress" field. It's identical to ProgressEQ.
func Progress(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldProgress, v))
}

// Accuracy applies equality check predicate on the "accuracy" field. It's identical to AccuracyEQ.
func Accuracy(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldAccuracy, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldSpeed, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldCreatedAt, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v float64) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v float64) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldDuration, v))
}

// ClosableEQ applies the EQ predicate on the "closable" field.
func ClosableEQ(v bool) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldClosable, v))
}

// ClosableNEQ applies the NEQ predicate on the "closable" field.
func ClosableNEQ(v bool) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldClosable, v))
}

// StopwatchEQ applies the EQ predicate on the "stopwatch" field.
func StopwatchEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldStopwatch, v))
}

// StopwatchNEQ applies the NEQ predicate on the "stopwatch" field.
func StopwatchNEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldStopwatch, v))
}

// StopwatchIn applies the In predicate on the "stopwatch" field.
func StopwatchIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldStopwatch, vs...))
}

// StopwatchNotIn applies the NotIn predicate on the "stopwatch" field.
func StopwatchNotIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldStopwatch, vs...))
}

// StopwatchGT applies the GT predicate on the "stopwatch" field.
func StopwatchGT(v float64) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldStopwatch, v))
}

// StopwatchGTE applies the GTE predicate on the "stopwatch" field.
func StopwatchGTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldStopwatch, v))
}

// StopwatchLT applies the LT predicate on the "stopwatch" field.
func StopwatchLT(v float64) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldStopwatch, v))
}

// StopwatchLTE applies the LTE predicate on the "stopwatch" field.
func StopwatchLTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldStopwatch, v))
}

// ProgressEQ applies the EQ predicate on the "progress" field.
func ProgressEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldProgress, v))
}

// ProgressNEQ applies the NEQ predicate on the "progress" field.
func ProgressNEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldProgress, v))
}

// ProgressIn applies the In predicate on the "progress" field.
func ProgressIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldProgress, vs...))
}

// ProgressNotIn applies the NotIn predicate on the "progress" field.
func ProgressNotIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldProgress, vs...))
}

// ProgressGT applies the GT predicate on the "progress" field.
func ProgressGT(v float64) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldProgress, v))
}

// ProgressGTE applies the GTE predicate on the "progress" field.
func ProgressGTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldProgress, v))
}

// ProgressLT applies the LT predicate on the "progress" field.
func ProgressLT(v float64) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldProgress, v))
}

// ProgressLTE applies the LTE predicate on the "progress" field.
func ProgressLTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldProgress, v))
}

// AccuracyEQ applies the EQ predicate on the "accuracy" field.
func AccuracyEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldAccuracy, v))
}

// AccuracyNEQ applies the NEQ predicate on the "accuracy" field.
func AccuracyNEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldAccuracy, v))
}

// AccuracyIn applies the In predicate on the "accuracy" field.
func AccuracyIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldAccuracy, vs...))
}

// AccuracyNotIn applies the NotIn predicate on the "accuracy" field.
func AccuracyNotIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldAccuracy, vs...))
}

// AccuracyGT applies the GT predicate on the "accuracy" field.
func AccuracyGT(v float64) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldAccuracy, v))
}

// AccuracyGTE applies the GTE predicate on the "accuracy" field.
func AccuracyGTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldAccuracy, v))
}

// AccuracyLT applies the LT predicate on the "accuracy" field.
func AccuracyLT(v float64) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldAccuracy, v))
}

// AccuracyLTE applies the LTE predicate on the "accuracy" field.
func AccuracyLTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldAccuracy, v))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v float64) predicate.Training {
	return predicate.Training(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...float64) predicate.Training {
	return predicate.Training(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v float64) predicate.Training {
	return predicate.Training(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v float64) predicate.Training {
	return predicate.Training(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v float64) predicate.Training {
	return predicate.Training(sql.FieldLTE(FieldSpeed, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Training) predicate.Training {
	return predicate.Training(func(s *sql.Selector) {
		p(s.Not())
	})
}

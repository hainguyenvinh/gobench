// Code generated by entc, DO NOT EDIT.

package application

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gobench-io/gobench/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Scenario applies equality check predicate on the "scenario" field. It's identical to ScenarioEQ.
func Scenario(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScenario), v))
	})
}

// Gomod applies equality check predicate on the "gomod" field. It's identical to GomodEQ.
func Gomod(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGomod), v))
	})
}

// Gosum applies equality check predicate on the "gosum" field. It's identical to GosumEQ.
func Gosum(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGosum), v))
	})
}

// Tags applies equality check predicate on the "tags" field. It's identical to TagsEQ.
func Tags(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTags), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ScenarioEQ applies the EQ predicate on the "scenario" field.
func ScenarioEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScenario), v))
	})
}

// ScenarioNEQ applies the NEQ predicate on the "scenario" field.
func ScenarioNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScenario), v))
	})
}

// ScenarioIn applies the In predicate on the "scenario" field.
func ScenarioIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScenario), v...))
	})
}

// ScenarioNotIn applies the NotIn predicate on the "scenario" field.
func ScenarioNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScenario), v...))
	})
}

// ScenarioGT applies the GT predicate on the "scenario" field.
func ScenarioGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScenario), v))
	})
}

// ScenarioGTE applies the GTE predicate on the "scenario" field.
func ScenarioGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScenario), v))
	})
}

// ScenarioLT applies the LT predicate on the "scenario" field.
func ScenarioLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScenario), v))
	})
}

// ScenarioLTE applies the LTE predicate on the "scenario" field.
func ScenarioLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScenario), v))
	})
}

// ScenarioContains applies the Contains predicate on the "scenario" field.
func ScenarioContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScenario), v))
	})
}

// ScenarioHasPrefix applies the HasPrefix predicate on the "scenario" field.
func ScenarioHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScenario), v))
	})
}

// ScenarioHasSuffix applies the HasSuffix predicate on the "scenario" field.
func ScenarioHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScenario), v))
	})
}

// ScenarioEqualFold applies the EqualFold predicate on the "scenario" field.
func ScenarioEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScenario), v))
	})
}

// ScenarioContainsFold applies the ContainsFold predicate on the "scenario" field.
func ScenarioContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScenario), v))
	})
}

// GomodEQ applies the EQ predicate on the "gomod" field.
func GomodEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGomod), v))
	})
}

// GomodNEQ applies the NEQ predicate on the "gomod" field.
func GomodNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGomod), v))
	})
}

// GomodIn applies the In predicate on the "gomod" field.
func GomodIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGomod), v...))
	})
}

// GomodNotIn applies the NotIn predicate on the "gomod" field.
func GomodNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGomod), v...))
	})
}

// GomodGT applies the GT predicate on the "gomod" field.
func GomodGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGomod), v))
	})
}

// GomodGTE applies the GTE predicate on the "gomod" field.
func GomodGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGomod), v))
	})
}

// GomodLT applies the LT predicate on the "gomod" field.
func GomodLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGomod), v))
	})
}

// GomodLTE applies the LTE predicate on the "gomod" field.
func GomodLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGomod), v))
	})
}

// GomodContains applies the Contains predicate on the "gomod" field.
func GomodContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGomod), v))
	})
}

// GomodHasPrefix applies the HasPrefix predicate on the "gomod" field.
func GomodHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGomod), v))
	})
}

// GomodHasSuffix applies the HasSuffix predicate on the "gomod" field.
func GomodHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGomod), v))
	})
}

// GomodEqualFold applies the EqualFold predicate on the "gomod" field.
func GomodEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGomod), v))
	})
}

// GomodContainsFold applies the ContainsFold predicate on the "gomod" field.
func GomodContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGomod), v))
	})
}

// GosumEQ applies the EQ predicate on the "gosum" field.
func GosumEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGosum), v))
	})
}

// GosumNEQ applies the NEQ predicate on the "gosum" field.
func GosumNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGosum), v))
	})
}

// GosumIn applies the In predicate on the "gosum" field.
func GosumIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGosum), v...))
	})
}

// GosumNotIn applies the NotIn predicate on the "gosum" field.
func GosumNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGosum), v...))
	})
}

// GosumGT applies the GT predicate on the "gosum" field.
func GosumGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGosum), v))
	})
}

// GosumGTE applies the GTE predicate on the "gosum" field.
func GosumGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGosum), v))
	})
}

// GosumLT applies the LT predicate on the "gosum" field.
func GosumLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGosum), v))
	})
}

// GosumLTE applies the LTE predicate on the "gosum" field.
func GosumLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGosum), v))
	})
}

// GosumContains applies the Contains predicate on the "gosum" field.
func GosumContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGosum), v))
	})
}

// GosumHasPrefix applies the HasPrefix predicate on the "gosum" field.
func GosumHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGosum), v))
	})
}

// GosumHasSuffix applies the HasSuffix predicate on the "gosum" field.
func GosumHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGosum), v))
	})
}

// GosumEqualFold applies the EqualFold predicate on the "gosum" field.
func GosumEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGosum), v))
	})
}

// GosumContainsFold applies the ContainsFold predicate on the "gosum" field.
func GosumContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGosum), v))
	})
}

// TagsEQ applies the EQ predicate on the "tags" field.
func TagsEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTags), v))
	})
}

// TagsNEQ applies the NEQ predicate on the "tags" field.
func TagsNEQ(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTags), v))
	})
}

// TagsIn applies the In predicate on the "tags" field.
func TagsIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTags), v...))
	})
}

// TagsNotIn applies the NotIn predicate on the "tags" field.
func TagsNotIn(vs ...string) predicate.Application {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Application(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTags), v...))
	})
}

// TagsGT applies the GT predicate on the "tags" field.
func TagsGT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTags), v))
	})
}

// TagsGTE applies the GTE predicate on the "tags" field.
func TagsGTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTags), v))
	})
}

// TagsLT applies the LT predicate on the "tags" field.
func TagsLT(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTags), v))
	})
}

// TagsLTE applies the LTE predicate on the "tags" field.
func TagsLTE(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTags), v))
	})
}

// TagsContains applies the Contains predicate on the "tags" field.
func TagsContains(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTags), v))
	})
}

// TagsHasPrefix applies the HasPrefix predicate on the "tags" field.
func TagsHasPrefix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTags), v))
	})
}

// TagsHasSuffix applies the HasSuffix predicate on the "tags" field.
func TagsHasSuffix(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTags), v))
	})
}

// TagsEqualFold applies the EqualFold predicate on the "tags" field.
func TagsEqualFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTags), v))
	})
}

// TagsContainsFold applies the ContainsFold predicate on the "tags" field.
func TagsContainsFold(v string) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTags), v))
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
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
func Not(p predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		p(s.Not())
	})
}

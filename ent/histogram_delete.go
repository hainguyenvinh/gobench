// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/predicate"
)

// HistogramDelete is the builder for deleting a Histogram entity.
type HistogramDelete struct {
	config
	hooks      []Hook
	mutation   *HistogramMutation
	predicates []predicate.Histogram
}

// Where adds a new predicate to the delete builder.
func (hd *HistogramDelete) Where(ps ...predicate.Histogram) *HistogramDelete {
	hd.predicates = append(hd.predicates, ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HistogramDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hd.hooks) == 0 {
		affected, err = hd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistogramMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hd.mutation = mutation
			affected, err = hd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hd.hooks) - 1; i >= 0; i-- {
			mut = hd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HistogramDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HistogramDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: histogram.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: histogram.FieldID,
			},
		},
	}
	if ps := hd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
}

// HistogramDeleteOne is the builder for deleting a single Histogram entity.
type HistogramDeleteOne struct {
	hd *HistogramDelete
}

// Exec executes the deletion query.
func (hdo *HistogramDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{histogram.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HistogramDeleteOne) ExecX(ctx context.Context) {
	hdo.hd.ExecX(ctx)
}

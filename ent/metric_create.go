// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/counter"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"
)

// MetricCreate is the builder for creating a Metric entity.
type MetricCreate struct {
	config
	mutation *MetricMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (mc *MetricCreate) SetTitle(s string) *MetricCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetType sets the type field.
func (mc *MetricCreate) SetType(s string) *MetricCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetGraphID sets the graph edge to Graph by id.
func (mc *MetricCreate) SetGraphID(id int) *MetricCreate {
	mc.mutation.SetGraphID(id)
	return mc
}

// SetNillableGraphID sets the graph edge to Graph by id if the given value is not nil.
func (mc *MetricCreate) SetNillableGraphID(id *int) *MetricCreate {
	if id != nil {
		mc = mc.SetGraphID(*id)
	}
	return mc
}

// SetGraph sets the graph edge to Graph.
func (mc *MetricCreate) SetGraph(g *Graph) *MetricCreate {
	return mc.SetGraphID(g.ID)
}

// AddHistogramIDs adds the histograms edge to Histogram by ids.
func (mc *MetricCreate) AddHistogramIDs(ids ...int) *MetricCreate {
	mc.mutation.AddHistogramIDs(ids...)
	return mc
}

// AddHistograms adds the histograms edges to Histogram.
func (mc *MetricCreate) AddHistograms(h ...*Histogram) *MetricCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return mc.AddHistogramIDs(ids...)
}

// AddCounterIDs adds the counters edge to Counter by ids.
func (mc *MetricCreate) AddCounterIDs(ids ...int) *MetricCreate {
	mc.mutation.AddCounterIDs(ids...)
	return mc
}

// AddCounters adds the counters edges to Counter.
func (mc *MetricCreate) AddCounters(c ...*Counter) *MetricCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddCounterIDs(ids...)
}

// AddGaugeIDs adds the gauges edge to Gauge by ids.
func (mc *MetricCreate) AddGaugeIDs(ids ...int) *MetricCreate {
	mc.mutation.AddGaugeIDs(ids...)
	return mc
}

// AddGauges adds the gauges edges to Gauge.
func (mc *MetricCreate) AddGauges(g ...*Gauge) *MetricCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mc.AddGaugeIDs(ids...)
}

// Mutation returns the MetricMutation object of the builder.
func (mc *MetricCreate) Mutation() *MetricMutation {
	return mc.mutation
}

// Save creates the Metric in the database.
func (mc *MetricCreate) Save(ctx context.Context) (*Metric, error) {
	var (
		err  error
		node *Metric
	)
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetricCreate) SaveX(ctx context.Context) *Metric {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetricCreate) check() error {
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	return nil
}

func (mc *MetricCreate) sqlSave(ctx context.Context) (*Metric, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MetricCreate) createSpec() (*Metric, *sqlgraph.CreateSpec) {
	var (
		_node = &Metric{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metric.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldType,
		})
		_node.Type = value
	}
	if nodes := mc.mutation.GraphIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.GraphTable,
			Columns: []string{metric.GraphColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: graph.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.HistogramsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metric.HistogramsTable,
			Columns: []string{metric.HistogramsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: histogram.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.CountersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metric.CountersTable,
			Columns: []string{metric.CountersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: counter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.GaugesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metric.GaugesTable,
			Columns: []string{metric.GaugesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gauge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MetricCreateBulk is the builder for creating a bulk of Metric entities.
type MetricCreateBulk struct {
	config
	builders []*MetricCreate
}

// Save creates the Metric entities in the database.
func (mcb *MetricCreateBulk) Save(ctx context.Context) ([]*Metric, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metric, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MetricCreateBulk) SaveX(ctx context.Context) []*Metric {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/counter"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"
	"github.com/gobench-io/gobench/ent/predicate"
)

// MetricUpdate is the builder for updating Metric entities.
type MetricUpdate struct {
	config
	hooks      []Hook
	mutation   *MetricMutation
	predicates []predicate.Metric
}

// Where adds a new predicate for the builder.
func (mu *MetricUpdate) Where(ps ...predicate.Metric) *MetricUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetType sets the type field.
func (mu *MetricUpdate) SetType(s string) *MetricUpdate {
	mu.mutation.SetType(s)
	return mu
}

// SetGraphID sets the graph edge to Graph by id.
func (mu *MetricUpdate) SetGraphID(id int) *MetricUpdate {
	mu.mutation.SetGraphID(id)
	return mu
}

// SetNillableGraphID sets the graph edge to Graph by id if the given value is not nil.
func (mu *MetricUpdate) SetNillableGraphID(id *int) *MetricUpdate {
	if id != nil {
		mu = mu.SetGraphID(*id)
	}
	return mu
}

// SetGraph sets the graph edge to Graph.
func (mu *MetricUpdate) SetGraph(g *Graph) *MetricUpdate {
	return mu.SetGraphID(g.ID)
}

// AddHistogramIDs adds the histograms edge to Histogram by ids.
func (mu *MetricUpdate) AddHistogramIDs(ids ...int) *MetricUpdate {
	mu.mutation.AddHistogramIDs(ids...)
	return mu
}

// AddHistograms adds the histograms edges to Histogram.
func (mu *MetricUpdate) AddHistograms(h ...*Histogram) *MetricUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return mu.AddHistogramIDs(ids...)
}

// AddCounterIDs adds the counters edge to Counter by ids.
func (mu *MetricUpdate) AddCounterIDs(ids ...int) *MetricUpdate {
	mu.mutation.AddCounterIDs(ids...)
	return mu
}

// AddCounters adds the counters edges to Counter.
func (mu *MetricUpdate) AddCounters(c ...*Counter) *MetricUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.AddCounterIDs(ids...)
}

// AddGaugeIDs adds the gauges edge to Gauge by ids.
func (mu *MetricUpdate) AddGaugeIDs(ids ...int) *MetricUpdate {
	mu.mutation.AddGaugeIDs(ids...)
	return mu
}

// AddGauges adds the gauges edges to Gauge.
func (mu *MetricUpdate) AddGauges(g ...*Gauge) *MetricUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.AddGaugeIDs(ids...)
}

// Mutation returns the MetricMutation object of the builder.
func (mu *MetricUpdate) Mutation() *MetricMutation {
	return mu.mutation
}

// ClearGraph clears the graph edge to Graph.
func (mu *MetricUpdate) ClearGraph() *MetricUpdate {
	mu.mutation.ClearGraph()
	return mu
}

// RemoveHistogramIDs removes the histograms edge to Histogram by ids.
func (mu *MetricUpdate) RemoveHistogramIDs(ids ...int) *MetricUpdate {
	mu.mutation.RemoveHistogramIDs(ids...)
	return mu
}

// RemoveHistograms removes histograms edges to Histogram.
func (mu *MetricUpdate) RemoveHistograms(h ...*Histogram) *MetricUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return mu.RemoveHistogramIDs(ids...)
}

// RemoveCounterIDs removes the counters edge to Counter by ids.
func (mu *MetricUpdate) RemoveCounterIDs(ids ...int) *MetricUpdate {
	mu.mutation.RemoveCounterIDs(ids...)
	return mu
}

// RemoveCounters removes counters edges to Counter.
func (mu *MetricUpdate) RemoveCounters(c ...*Counter) *MetricUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.RemoveCounterIDs(ids...)
}

// RemoveGaugeIDs removes the gauges edge to Gauge by ids.
func (mu *MetricUpdate) RemoveGaugeIDs(ids ...int) *MetricUpdate {
	mu.mutation.RemoveGaugeIDs(ids...)
	return mu
}

// RemoveGauges removes gauges edges to Gauge.
func (mu *MetricUpdate) RemoveGauges(g ...*Gauge) *MetricUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.RemoveGaugeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MetricUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetricUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetricUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetricUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MetricUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldType,
		})
	}
	if mu.mutation.GraphCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.GraphIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedHistogramsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.HistogramsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedCountersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CountersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedGaugesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.GaugesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MetricUpdateOne is the builder for updating a single Metric entity.
type MetricUpdateOne struct {
	config
	hooks    []Hook
	mutation *MetricMutation
}

// SetType sets the type field.
func (muo *MetricUpdateOne) SetType(s string) *MetricUpdateOne {
	muo.mutation.SetType(s)
	return muo
}

// SetGraphID sets the graph edge to Graph by id.
func (muo *MetricUpdateOne) SetGraphID(id int) *MetricUpdateOne {
	muo.mutation.SetGraphID(id)
	return muo
}

// SetNillableGraphID sets the graph edge to Graph by id if the given value is not nil.
func (muo *MetricUpdateOne) SetNillableGraphID(id *int) *MetricUpdateOne {
	if id != nil {
		muo = muo.SetGraphID(*id)
	}
	return muo
}

// SetGraph sets the graph edge to Graph.
func (muo *MetricUpdateOne) SetGraph(g *Graph) *MetricUpdateOne {
	return muo.SetGraphID(g.ID)
}

// AddHistogramIDs adds the histograms edge to Histogram by ids.
func (muo *MetricUpdateOne) AddHistogramIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.AddHistogramIDs(ids...)
	return muo
}

// AddHistograms adds the histograms edges to Histogram.
func (muo *MetricUpdateOne) AddHistograms(h ...*Histogram) *MetricUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return muo.AddHistogramIDs(ids...)
}

// AddCounterIDs adds the counters edge to Counter by ids.
func (muo *MetricUpdateOne) AddCounterIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.AddCounterIDs(ids...)
	return muo
}

// AddCounters adds the counters edges to Counter.
func (muo *MetricUpdateOne) AddCounters(c ...*Counter) *MetricUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.AddCounterIDs(ids...)
}

// AddGaugeIDs adds the gauges edge to Gauge by ids.
func (muo *MetricUpdateOne) AddGaugeIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.AddGaugeIDs(ids...)
	return muo
}

// AddGauges adds the gauges edges to Gauge.
func (muo *MetricUpdateOne) AddGauges(g ...*Gauge) *MetricUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.AddGaugeIDs(ids...)
}

// Mutation returns the MetricMutation object of the builder.
func (muo *MetricUpdateOne) Mutation() *MetricMutation {
	return muo.mutation
}

// ClearGraph clears the graph edge to Graph.
func (muo *MetricUpdateOne) ClearGraph() *MetricUpdateOne {
	muo.mutation.ClearGraph()
	return muo
}

// RemoveHistogramIDs removes the histograms edge to Histogram by ids.
func (muo *MetricUpdateOne) RemoveHistogramIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.RemoveHistogramIDs(ids...)
	return muo
}

// RemoveHistograms removes histograms edges to Histogram.
func (muo *MetricUpdateOne) RemoveHistograms(h ...*Histogram) *MetricUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return muo.RemoveHistogramIDs(ids...)
}

// RemoveCounterIDs removes the counters edge to Counter by ids.
func (muo *MetricUpdateOne) RemoveCounterIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.RemoveCounterIDs(ids...)
	return muo
}

// RemoveCounters removes counters edges to Counter.
func (muo *MetricUpdateOne) RemoveCounters(c ...*Counter) *MetricUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.RemoveCounterIDs(ids...)
}

// RemoveGaugeIDs removes the gauges edge to Gauge by ids.
func (muo *MetricUpdateOne) RemoveGaugeIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.RemoveGaugeIDs(ids...)
	return muo
}

// RemoveGauges removes gauges edges to Gauge.
func (muo *MetricUpdateOne) RemoveGauges(g ...*Gauge) *MetricUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.RemoveGaugeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MetricUpdateOne) Save(ctx context.Context) (*Metric, error) {

	var (
		err  error
		node *Metric
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetricUpdateOne) SaveX(ctx context.Context) *Metric {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MetricUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetricUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MetricUpdateOne) sqlSave(ctx context.Context) (m *Metric, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Metric.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldType,
		})
	}
	if muo.mutation.GraphCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.GraphIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedHistogramsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.HistogramsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedCountersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CountersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedGaugesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.GaugesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Metric{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}

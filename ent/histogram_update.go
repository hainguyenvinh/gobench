// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"
	"github.com/gobench-io/gobench/ent/predicate"
)

// HistogramUpdate is the builder for updating Histogram entities.
type HistogramUpdate struct {
	config
	hooks      []Hook
	mutation   *HistogramMutation
	predicates []predicate.Histogram
}

// Where adds a new predicate for the builder.
func (hu *HistogramUpdate) Where(ps ...predicate.Histogram) *HistogramUpdate {
	hu.predicates = append(hu.predicates, ps...)
	return hu
}

// SetTime sets the time field.
func (hu *HistogramUpdate) SetTime(i int64) *HistogramUpdate {
	hu.mutation.ResetTime()
	hu.mutation.SetTime(i)
	return hu
}

// AddTime adds i to time.
func (hu *HistogramUpdate) AddTime(i int64) *HistogramUpdate {
	hu.mutation.AddTime(i)
	return hu
}

// SetCount sets the count field.
func (hu *HistogramUpdate) SetCount(i int64) *HistogramUpdate {
	hu.mutation.ResetCount()
	hu.mutation.SetCount(i)
	return hu
}

// AddCount adds i to count.
func (hu *HistogramUpdate) AddCount(i int64) *HistogramUpdate {
	hu.mutation.AddCount(i)
	return hu
}

// SetMin sets the min field.
func (hu *HistogramUpdate) SetMin(i int64) *HistogramUpdate {
	hu.mutation.ResetMin()
	hu.mutation.SetMin(i)
	return hu
}

// AddMin adds i to min.
func (hu *HistogramUpdate) AddMin(i int64) *HistogramUpdate {
	hu.mutation.AddMin(i)
	return hu
}

// SetMax sets the max field.
func (hu *HistogramUpdate) SetMax(i int64) *HistogramUpdate {
	hu.mutation.ResetMax()
	hu.mutation.SetMax(i)
	return hu
}

// AddMax adds i to max.
func (hu *HistogramUpdate) AddMax(i int64) *HistogramUpdate {
	hu.mutation.AddMax(i)
	return hu
}

// SetMean sets the mean field.
func (hu *HistogramUpdate) SetMean(f float64) *HistogramUpdate {
	hu.mutation.ResetMean()
	hu.mutation.SetMean(f)
	return hu
}

// AddMean adds f to mean.
func (hu *HistogramUpdate) AddMean(f float64) *HistogramUpdate {
	hu.mutation.AddMean(f)
	return hu
}

// SetStddev sets the stddev field.
func (hu *HistogramUpdate) SetStddev(f float64) *HistogramUpdate {
	hu.mutation.ResetStddev()
	hu.mutation.SetStddev(f)
	return hu
}

// AddStddev adds f to stddev.
func (hu *HistogramUpdate) AddStddev(f float64) *HistogramUpdate {
	hu.mutation.AddStddev(f)
	return hu
}

// SetMedian sets the median field.
func (hu *HistogramUpdate) SetMedian(f float64) *HistogramUpdate {
	hu.mutation.ResetMedian()
	hu.mutation.SetMedian(f)
	return hu
}

// AddMedian adds f to median.
func (hu *HistogramUpdate) AddMedian(f float64) *HistogramUpdate {
	hu.mutation.AddMedian(f)
	return hu
}

// SetP75 sets the p75 field.
func (hu *HistogramUpdate) SetP75(f float64) *HistogramUpdate {
	hu.mutation.ResetP75()
	hu.mutation.SetP75(f)
	return hu
}

// AddP75 adds f to p75.
func (hu *HistogramUpdate) AddP75(f float64) *HistogramUpdate {
	hu.mutation.AddP75(f)
	return hu
}

// SetP95 sets the p95 field.
func (hu *HistogramUpdate) SetP95(f float64) *HistogramUpdate {
	hu.mutation.ResetP95()
	hu.mutation.SetP95(f)
	return hu
}

// AddP95 adds f to p95.
func (hu *HistogramUpdate) AddP95(f float64) *HistogramUpdate {
	hu.mutation.AddP95(f)
	return hu
}

// SetP99 sets the p99 field.
func (hu *HistogramUpdate) SetP99(f float64) *HistogramUpdate {
	hu.mutation.ResetP99()
	hu.mutation.SetP99(f)
	return hu
}

// AddP99 adds f to p99.
func (hu *HistogramUpdate) AddP99(f float64) *HistogramUpdate {
	hu.mutation.AddP99(f)
	return hu
}

// SetP999 sets the p999 field.
func (hu *HistogramUpdate) SetP999(f float64) *HistogramUpdate {
	hu.mutation.ResetP999()
	hu.mutation.SetP999(f)
	return hu
}

// AddP999 adds f to p999.
func (hu *HistogramUpdate) AddP999(f float64) *HistogramUpdate {
	hu.mutation.AddP999(f)
	return hu
}

// SetWID sets the wID field.
func (hu *HistogramUpdate) SetWID(s string) *HistogramUpdate {
	hu.mutation.SetWID(s)
	return hu
}

// SetMetricID sets the metric edge to Metric by id.
func (hu *HistogramUpdate) SetMetricID(id int) *HistogramUpdate {
	hu.mutation.SetMetricID(id)
	return hu
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (hu *HistogramUpdate) SetNillableMetricID(id *int) *HistogramUpdate {
	if id != nil {
		hu = hu.SetMetricID(*id)
	}
	return hu
}

// SetMetric sets the metric edge to Metric.
func (hu *HistogramUpdate) SetMetric(m *Metric) *HistogramUpdate {
	return hu.SetMetricID(m.ID)
}

// Mutation returns the HistogramMutation object of the builder.
func (hu *HistogramUpdate) Mutation() *HistogramMutation {
	return hu.mutation
}

// ClearMetric clears the metric edge to Metric.
func (hu *HistogramUpdate) ClearMetric() *HistogramUpdate {
	hu.mutation.ClearMetric()
	return hu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (hu *HistogramUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistogramMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistogramUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistogramUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistogramUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HistogramUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   histogram.Table,
			Columns: histogram.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: histogram.FieldID,
			},
		},
	}
	if ps := hu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldTime,
		})
	}
	if value, ok := hu.mutation.AddedTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldTime,
		})
	}
	if value, ok := hu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldCount,
		})
	}
	if value, ok := hu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldCount,
		})
	}
	if value, ok := hu.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMin,
		})
	}
	if value, ok := hu.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMin,
		})
	}
	if value, ok := hu.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMax,
		})
	}
	if value, ok := hu.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMax,
		})
	}
	if value, ok := hu.mutation.Mean(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMean,
		})
	}
	if value, ok := hu.mutation.AddedMean(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMean,
		})
	}
	if value, ok := hu.mutation.Stddev(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldStddev,
		})
	}
	if value, ok := hu.mutation.AddedStddev(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldStddev,
		})
	}
	if value, ok := hu.mutation.Median(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMedian,
		})
	}
	if value, ok := hu.mutation.AddedMedian(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMedian,
		})
	}
	if value, ok := hu.mutation.P75(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP75,
		})
	}
	if value, ok := hu.mutation.AddedP75(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP75,
		})
	}
	if value, ok := hu.mutation.P95(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP95,
		})
	}
	if value, ok := hu.mutation.AddedP95(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP95,
		})
	}
	if value, ok := hu.mutation.P99(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP99,
		})
	}
	if value, ok := hu.mutation.AddedP99(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP99,
		})
	}
	if value, ok := hu.mutation.P999(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP999,
		})
	}
	if value, ok := hu.mutation.AddedP999(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP999,
		})
	}
	if value, ok := hu.mutation.WID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: histogram.FieldWID,
		})
	}
	if hu.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   histogram.MetricTable,
			Columns: []string{histogram.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   histogram.MetricTable,
			Columns: []string{histogram.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{histogram.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HistogramUpdateOne is the builder for updating a single Histogram entity.
type HistogramUpdateOne struct {
	config
	hooks    []Hook
	mutation *HistogramMutation
}

// SetTime sets the time field.
func (huo *HistogramUpdateOne) SetTime(i int64) *HistogramUpdateOne {
	huo.mutation.ResetTime()
	huo.mutation.SetTime(i)
	return huo
}

// AddTime adds i to time.
func (huo *HistogramUpdateOne) AddTime(i int64) *HistogramUpdateOne {
	huo.mutation.AddTime(i)
	return huo
}

// SetCount sets the count field.
func (huo *HistogramUpdateOne) SetCount(i int64) *HistogramUpdateOne {
	huo.mutation.ResetCount()
	huo.mutation.SetCount(i)
	return huo
}

// AddCount adds i to count.
func (huo *HistogramUpdateOne) AddCount(i int64) *HistogramUpdateOne {
	huo.mutation.AddCount(i)
	return huo
}

// SetMin sets the min field.
func (huo *HistogramUpdateOne) SetMin(i int64) *HistogramUpdateOne {
	huo.mutation.ResetMin()
	huo.mutation.SetMin(i)
	return huo
}

// AddMin adds i to min.
func (huo *HistogramUpdateOne) AddMin(i int64) *HistogramUpdateOne {
	huo.mutation.AddMin(i)
	return huo
}

// SetMax sets the max field.
func (huo *HistogramUpdateOne) SetMax(i int64) *HistogramUpdateOne {
	huo.mutation.ResetMax()
	huo.mutation.SetMax(i)
	return huo
}

// AddMax adds i to max.
func (huo *HistogramUpdateOne) AddMax(i int64) *HistogramUpdateOne {
	huo.mutation.AddMax(i)
	return huo
}

// SetMean sets the mean field.
func (huo *HistogramUpdateOne) SetMean(f float64) *HistogramUpdateOne {
	huo.mutation.ResetMean()
	huo.mutation.SetMean(f)
	return huo
}

// AddMean adds f to mean.
func (huo *HistogramUpdateOne) AddMean(f float64) *HistogramUpdateOne {
	huo.mutation.AddMean(f)
	return huo
}

// SetStddev sets the stddev field.
func (huo *HistogramUpdateOne) SetStddev(f float64) *HistogramUpdateOne {
	huo.mutation.ResetStddev()
	huo.mutation.SetStddev(f)
	return huo
}

// AddStddev adds f to stddev.
func (huo *HistogramUpdateOne) AddStddev(f float64) *HistogramUpdateOne {
	huo.mutation.AddStddev(f)
	return huo
}

// SetMedian sets the median field.
func (huo *HistogramUpdateOne) SetMedian(f float64) *HistogramUpdateOne {
	huo.mutation.ResetMedian()
	huo.mutation.SetMedian(f)
	return huo
}

// AddMedian adds f to median.
func (huo *HistogramUpdateOne) AddMedian(f float64) *HistogramUpdateOne {
	huo.mutation.AddMedian(f)
	return huo
}

// SetP75 sets the p75 field.
func (huo *HistogramUpdateOne) SetP75(f float64) *HistogramUpdateOne {
	huo.mutation.ResetP75()
	huo.mutation.SetP75(f)
	return huo
}

// AddP75 adds f to p75.
func (huo *HistogramUpdateOne) AddP75(f float64) *HistogramUpdateOne {
	huo.mutation.AddP75(f)
	return huo
}

// SetP95 sets the p95 field.
func (huo *HistogramUpdateOne) SetP95(f float64) *HistogramUpdateOne {
	huo.mutation.ResetP95()
	huo.mutation.SetP95(f)
	return huo
}

// AddP95 adds f to p95.
func (huo *HistogramUpdateOne) AddP95(f float64) *HistogramUpdateOne {
	huo.mutation.AddP95(f)
	return huo
}

// SetP99 sets the p99 field.
func (huo *HistogramUpdateOne) SetP99(f float64) *HistogramUpdateOne {
	huo.mutation.ResetP99()
	huo.mutation.SetP99(f)
	return huo
}

// AddP99 adds f to p99.
func (huo *HistogramUpdateOne) AddP99(f float64) *HistogramUpdateOne {
	huo.mutation.AddP99(f)
	return huo
}

// SetP999 sets the p999 field.
func (huo *HistogramUpdateOne) SetP999(f float64) *HistogramUpdateOne {
	huo.mutation.ResetP999()
	huo.mutation.SetP999(f)
	return huo
}

// AddP999 adds f to p999.
func (huo *HistogramUpdateOne) AddP999(f float64) *HistogramUpdateOne {
	huo.mutation.AddP999(f)
	return huo
}

// SetWID sets the wID field.
func (huo *HistogramUpdateOne) SetWID(s string) *HistogramUpdateOne {
	huo.mutation.SetWID(s)
	return huo
}

// SetMetricID sets the metric edge to Metric by id.
func (huo *HistogramUpdateOne) SetMetricID(id int) *HistogramUpdateOne {
	huo.mutation.SetMetricID(id)
	return huo
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (huo *HistogramUpdateOne) SetNillableMetricID(id *int) *HistogramUpdateOne {
	if id != nil {
		huo = huo.SetMetricID(*id)
	}
	return huo
}

// SetMetric sets the metric edge to Metric.
func (huo *HistogramUpdateOne) SetMetric(m *Metric) *HistogramUpdateOne {
	return huo.SetMetricID(m.ID)
}

// Mutation returns the HistogramMutation object of the builder.
func (huo *HistogramUpdateOne) Mutation() *HistogramMutation {
	return huo.mutation
}

// ClearMetric clears the metric edge to Metric.
func (huo *HistogramUpdateOne) ClearMetric() *HistogramUpdateOne {
	huo.mutation.ClearMetric()
	return huo
}

// Save executes the query and returns the updated entity.
func (huo *HistogramUpdateOne) Save(ctx context.Context) (*Histogram, error) {

	var (
		err  error
		node *Histogram
	)
	if len(huo.hooks) == 0 {
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistogramMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistogramUpdateOne) SaveX(ctx context.Context) *Histogram {
	h, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return h
}

// Exec executes the query on the entity.
func (huo *HistogramUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistogramUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HistogramUpdateOne) sqlSave(ctx context.Context) (h *Histogram, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   histogram.Table,
			Columns: histogram.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: histogram.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Histogram.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := huo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldTime,
		})
	}
	if value, ok := huo.mutation.AddedTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldTime,
		})
	}
	if value, ok := huo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldCount,
		})
	}
	if value, ok := huo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldCount,
		})
	}
	if value, ok := huo.mutation.Min(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMin,
		})
	}
	if value, ok := huo.mutation.AddedMin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMin,
		})
	}
	if value, ok := huo.mutation.Max(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMax,
		})
	}
	if value, ok := huo.mutation.AddedMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMax,
		})
	}
	if value, ok := huo.mutation.Mean(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMean,
		})
	}
	if value, ok := huo.mutation.AddedMean(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMean,
		})
	}
	if value, ok := huo.mutation.Stddev(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldStddev,
		})
	}
	if value, ok := huo.mutation.AddedStddev(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldStddev,
		})
	}
	if value, ok := huo.mutation.Median(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMedian,
		})
	}
	if value, ok := huo.mutation.AddedMedian(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMedian,
		})
	}
	if value, ok := huo.mutation.P75(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP75,
		})
	}
	if value, ok := huo.mutation.AddedP75(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP75,
		})
	}
	if value, ok := huo.mutation.P95(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP95,
		})
	}
	if value, ok := huo.mutation.AddedP95(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP95,
		})
	}
	if value, ok := huo.mutation.P99(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP99,
		})
	}
	if value, ok := huo.mutation.AddedP99(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP99,
		})
	}
	if value, ok := huo.mutation.P999(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP999,
		})
	}
	if value, ok := huo.mutation.AddedP999(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP999,
		})
	}
	if value, ok := huo.mutation.WID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: histogram.FieldWID,
		})
	}
	if huo.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   histogram.MetricTable,
			Columns: []string{histogram.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   histogram.MetricTable,
			Columns: []string{histogram.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	h = &Histogram{config: huo.config}
	_spec.Assign = h.assignValues
	_spec.ScanValues = h.scanValues()
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{histogram.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return h, nil
}

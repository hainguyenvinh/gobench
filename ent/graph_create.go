// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/group"
	"github.com/gobench-io/gobench/ent/metric"
)

// GraphCreate is the builder for creating a Graph entity.
type GraphCreate struct {
	config
	mutation *GraphMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (gc *GraphCreate) SetTitle(s string) *GraphCreate {
	gc.mutation.SetTitle(s)
	return gc
}

// SetUnit sets the unit field.
func (gc *GraphCreate) SetUnit(s string) *GraphCreate {
	gc.mutation.SetUnit(s)
	return gc
}

// SetGroupID sets the group edge to Group by id.
func (gc *GraphCreate) SetGroupID(id int) *GraphCreate {
	gc.mutation.SetGroupID(id)
	return gc
}

// SetNillableGroupID sets the group edge to Group by id if the given value is not nil.
func (gc *GraphCreate) SetNillableGroupID(id *int) *GraphCreate {
	if id != nil {
		gc = gc.SetGroupID(*id)
	}
	return gc
}

// SetGroup sets the group edge to Group.
func (gc *GraphCreate) SetGroup(g *Group) *GraphCreate {
	return gc.SetGroupID(g.ID)
}

// AddMetricIDs adds the metrics edge to Metric by ids.
func (gc *GraphCreate) AddMetricIDs(ids ...int) *GraphCreate {
	gc.mutation.AddMetricIDs(ids...)
	return gc
}

// AddMetrics adds the metrics edges to Metric.
func (gc *GraphCreate) AddMetrics(m ...*Metric) *GraphCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMetricIDs(ids...)
}

// Mutation returns the GraphMutation object of the builder.
func (gc *GraphCreate) Mutation() *GraphMutation {
	return gc.mutation
}

// Save creates the Graph in the database.
func (gc *GraphCreate) Save(ctx context.Context) (*Graph, error) {
	if err := gc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Graph
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GraphMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GraphCreate) SaveX(ctx context.Context) *Graph {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GraphCreate) preSave() error {
	if _, ok := gc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := gc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New("ent: missing required field \"unit\"")}
	}
	return nil
}

func (gc *GraphCreate) sqlSave(ctx context.Context) (*Graph, error) {
	gr, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	gr.ID = int(id)
	return gr, nil
}

func (gc *GraphCreate) createSpec() (*Graph, *sqlgraph.CreateSpec) {
	var (
		gr    = &Graph{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: graph.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: graph.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: graph.FieldTitle,
		})
		gr.Title = value
	}
	if value, ok := gc.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: graph.FieldUnit,
		})
		gr.Unit = value
	}
	if nodes := gc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   graph.GroupTable,
			Columns: []string{graph.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   graph.MetricsTable,
			Columns: []string{graph.MetricsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return gr, _spec
}

// GraphCreateBulk is the builder for creating a bulk of Graph entities.
type GraphCreateBulk struct {
	config
	builders []*GraphCreate
}

// Save creates the Graph entities in the database.
func (gcb *GraphCreateBulk) Save(ctx context.Context) ([]*Graph, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Graph, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*GraphMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (gcb *GraphCreateBulk) SaveX(ctx context.Context) []*Graph {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

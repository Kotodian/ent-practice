// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/orderevent"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// OrderEventDelete is the builder for deleting a OrderEvent entity.
type OrderEventDelete struct {
	config
	hooks    []Hook
	mutation *OrderEventMutation
}

// Where appends a list predicates to the OrderEventDelete builder.
func (oed *OrderEventDelete) Where(ps ...predicate.OrderEvent) *OrderEventDelete {
	oed.mutation.Where(ps...)
	return oed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oed *OrderEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oed.hooks) == 0 {
		affected, err = oed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oed.mutation = mutation
			affected, err = oed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oed.hooks) - 1; i >= 0; i-- {
			if oed.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oed *OrderEventDelete) ExecX(ctx context.Context) int {
	n, err := oed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oed *OrderEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderevent.FieldID,
			},
		},
	}
	if ps := oed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrderEventDeleteOne is the builder for deleting a single OrderEvent entity.
type OrderEventDeleteOne struct {
	oed *OrderEventDelete
}

// Exec executes the deletion query.
func (oedo *OrderEventDeleteOne) Exec(ctx context.Context) error {
	n, err := oedo.oed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oedo *OrderEventDeleteOne) ExecX(ctx context.Context) {
	oedo.oed.ExecX(ctx)
}

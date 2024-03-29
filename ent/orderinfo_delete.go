// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/cwmodel/ent/predicate"
)

// OrderInfoDelete is the builder for deleting a OrderInfo entity.
type OrderInfoDelete struct {
	config
	hooks    []Hook
	mutation *OrderInfoMutation
}

// Where appends a list predicates to the OrderInfoDelete builder.
func (oid *OrderInfoDelete) Where(ps ...predicate.OrderInfo) *OrderInfoDelete {
	oid.mutation.Where(ps...)
	return oid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oid *OrderInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oid.hooks) == 0 {
		affected, err = oid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oid.mutation = mutation
			affected, err = oid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oid.hooks) - 1; i >= 0; i-- {
			if oid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oid *OrderInfoDelete) ExecX(ctx context.Context) int {
	n, err := oid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oid *OrderInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderinfo.FieldID,
			},
		},
	}
	if ps := oid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrderInfoDeleteOne is the builder for deleting a single OrderInfo entity.
type OrderInfoDeleteOne struct {
	oid *OrderInfoDelete
}

// Exec executes the deletion query.
func (oido *OrderInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := oido.oid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oido *OrderInfoDeleteOne) ExecX(ctx context.Context) {
	oido.oid.ExecX(ctx)
}

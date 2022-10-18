// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipmentiot"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// EquipmentIotDelete is the builder for deleting a EquipmentIot entity.
type EquipmentIotDelete struct {
	config
	hooks    []Hook
	mutation *EquipmentIotMutation
}

// Where appends a list predicates to the EquipmentIotDelete builder.
func (eid *EquipmentIotDelete) Where(ps ...predicate.EquipmentIot) *EquipmentIotDelete {
	eid.mutation.Where(ps...)
	return eid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eid *EquipmentIotDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eid.hooks) == 0 {
		affected, err = eid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentIotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eid.mutation = mutation
			affected, err = eid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eid.hooks) - 1; i >= 0; i-- {
			if eid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (eid *EquipmentIotDelete) ExecX(ctx context.Context) int {
	n, err := eid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eid *EquipmentIotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: equipmentiot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentiot.FieldID,
			},
		},
	}
	if ps := eid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// EquipmentIotDeleteOne is the builder for deleting a single EquipmentIot entity.
type EquipmentIotDeleteOne struct {
	eid *EquipmentIotDelete
}

// Exec executes the deletion query.
func (eido *EquipmentIotDeleteOne) Exec(ctx context.Context) error {
	n, err := eido.eid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{equipmentiot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eido *EquipmentIotDeleteOne) ExecX(ctx context.Context) {
	eido.eid.ExecX(ctx)
}

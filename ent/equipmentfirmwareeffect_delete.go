// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// EquipmentFirmwareEffectDelete is the builder for deleting a EquipmentFirmwareEffect entity.
type EquipmentFirmwareEffectDelete struct {
	config
	hooks    []Hook
	mutation *EquipmentFirmwareEffectMutation
}

// Where appends a list predicates to the EquipmentFirmwareEffectDelete builder.
func (efed *EquipmentFirmwareEffectDelete) Where(ps ...predicate.EquipmentFirmwareEffect) *EquipmentFirmwareEffectDelete {
	efed.mutation.Where(ps...)
	return efed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (efed *EquipmentFirmwareEffectDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(efed.hooks) == 0 {
		affected, err = efed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentFirmwareEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			efed.mutation = mutation
			affected, err = efed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(efed.hooks) - 1; i >= 0; i-- {
			if efed.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = efed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, efed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (efed *EquipmentFirmwareEffectDelete) ExecX(ctx context.Context) int {
	n, err := efed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (efed *EquipmentFirmwareEffectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: equipmentfirmwareeffect.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentfirmwareeffect.FieldID,
			},
		},
	}
	if ps := efed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, efed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// EquipmentFirmwareEffectDeleteOne is the builder for deleting a single EquipmentFirmwareEffect entity.
type EquipmentFirmwareEffectDeleteOne struct {
	efed *EquipmentFirmwareEffectDelete
}

// Exec executes the deletion query.
func (efedo *EquipmentFirmwareEffectDeleteOne) Exec(ctx context.Context) error {
	n, err := efedo.efed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (efedo *EquipmentFirmwareEffectDeleteOne) ExecX(ctx context.Context) {
	efedo.efed.ExecX(ctx)
}

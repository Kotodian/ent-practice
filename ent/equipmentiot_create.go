// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentiot"
)

// EquipmentIotCreate is the builder for creating a EquipmentIot entity.
type EquipmentIotCreate struct {
	config
	mutation *EquipmentIotMutation
	hooks    []Hook
}

// SetIccid sets the "iccid" field.
func (eic *EquipmentIotCreate) SetIccid(s string) *EquipmentIotCreate {
	eic.mutation.SetIccid(s)
	return eic
}

// SetNillableIccid sets the "iccid" field if the given value is not nil.
func (eic *EquipmentIotCreate) SetNillableIccid(s *string) *EquipmentIotCreate {
	if s != nil {
		eic.SetIccid(*s)
	}
	return eic
}

// SetImei sets the "imei" field.
func (eic *EquipmentIotCreate) SetImei(s string) *EquipmentIotCreate {
	eic.mutation.SetImei(s)
	return eic
}

// SetNillableImei sets the "imei" field if the given value is not nil.
func (eic *EquipmentIotCreate) SetNillableImei(s *string) *EquipmentIotCreate {
	if s != nil {
		eic.SetImei(*s)
	}
	return eic
}

// SetRemoteAddress sets the "remote_address" field.
func (eic *EquipmentIotCreate) SetRemoteAddress(s string) *EquipmentIotCreate {
	eic.mutation.SetRemoteAddress(s)
	return eic
}

// SetNillableRemoteAddress sets the "remote_address" field if the given value is not nil.
func (eic *EquipmentIotCreate) SetNillableRemoteAddress(s *string) *EquipmentIotCreate {
	if s != nil {
		eic.SetRemoteAddress(*s)
	}
	return eic
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eic *EquipmentIotCreate) SetEquipmentID(id int) *EquipmentIotCreate {
	eic.mutation.SetEquipmentID(id)
	return eic
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (eic *EquipmentIotCreate) SetNillableEquipmentID(id *int) *EquipmentIotCreate {
	if id != nil {
		eic = eic.SetEquipmentID(*id)
	}
	return eic
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eic *EquipmentIotCreate) SetEquipment(e *Equipment) *EquipmentIotCreate {
	return eic.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentIotMutation object of the builder.
func (eic *EquipmentIotCreate) Mutation() *EquipmentIotMutation {
	return eic.mutation
}

// Save creates the EquipmentIot in the database.
func (eic *EquipmentIotCreate) Save(ctx context.Context) (*EquipmentIot, error) {
	var (
		err  error
		node *EquipmentIot
	)
	if len(eic.hooks) == 0 {
		if err = eic.check(); err != nil {
			return nil, err
		}
		node, err = eic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentIotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eic.check(); err != nil {
				return nil, err
			}
			eic.mutation = mutation
			if node, err = eic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eic.hooks) - 1; i >= 0; i-- {
			if eic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EquipmentIot)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EquipmentIotMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eic *EquipmentIotCreate) SaveX(ctx context.Context) *EquipmentIot {
	v, err := eic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eic *EquipmentIotCreate) Exec(ctx context.Context) error {
	_, err := eic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eic *EquipmentIotCreate) ExecX(ctx context.Context) {
	if err := eic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eic *EquipmentIotCreate) check() error {
	return nil
}

func (eic *EquipmentIotCreate) sqlSave(ctx context.Context) (*EquipmentIot, error) {
	_node, _spec := eic.createSpec()
	if err := sqlgraph.CreateNode(ctx, eic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (eic *EquipmentIotCreate) createSpec() (*EquipmentIot, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentIot{config: eic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmentiot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentiot.FieldID,
			},
		}
	)
	if value, ok := eic.mutation.Iccid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldIccid,
		})
		_node.Iccid = value
	}
	if value, ok := eic.mutation.Imei(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldImei,
		})
		_node.Imei = value
	}
	if value, ok := eic.mutation.RemoteAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldRemoteAddress,
		})
		_node.RemoteAddress = value
	}
	if nodes := eic.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentiot.EquipmentTable,
			Columns: []string{equipmentiot.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.equipment_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentIotCreateBulk is the builder for creating many EquipmentIot entities in bulk.
type EquipmentIotCreateBulk struct {
	config
	builders []*EquipmentIotCreate
}

// Save creates the EquipmentIot entities in the database.
func (eicb *EquipmentIotCreateBulk) Save(ctx context.Context) ([]*EquipmentIot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eicb.builders))
	nodes := make([]*EquipmentIot, len(eicb.builders))
	mutators := make([]Mutator, len(eicb.builders))
	for i := range eicb.builders {
		func(i int, root context.Context) {
			builder := eicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentIotMutation)
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
					_, err = mutators[i+1].Mutate(root, eicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eicb *EquipmentIotCreateBulk) SaveX(ctx context.Context) []*EquipmentIot {
	v, err := eicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eicb *EquipmentIotCreateBulk) Exec(ctx context.Context) error {
	_, err := eicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eicb *EquipmentIotCreateBulk) ExecX(ctx context.Context) {
	if err := eicb.Exec(ctx); err != nil {
		panic(err)
	}
}

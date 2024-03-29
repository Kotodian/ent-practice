// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/connector"
	"github.com/Kotodian/cwmodel/ent/equipment"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/cwmodel/ent/smartchargingeffect"
	"github.com/Kotodian/cwmodel/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

// SmartChargingEffectCreate is the builder for creating a SmartChargingEffect entity.
type SmartChargingEffectCreate struct {
	config
	mutation *SmartChargingEffectMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (scec *SmartChargingEffectCreate) SetVersion(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetVersion(i)
	return scec
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableVersion(i *int64) *SmartChargingEffectCreate {
	if i != nil {
		scec.SetVersion(*i)
	}
	return scec
}

// SetCreatedBy sets the "created_by" field.
func (scec *SmartChargingEffectCreate) SetCreatedBy(d datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetCreatedBy(d)
	return scec
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableCreatedBy(d *datasource.UUID) *SmartChargingEffectCreate {
	if d != nil {
		scec.SetCreatedBy(*d)
	}
	return scec
}

// SetCreatedAt sets the "created_at" field.
func (scec *SmartChargingEffectCreate) SetCreatedAt(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetCreatedAt(i)
	return scec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableCreatedAt(i *int64) *SmartChargingEffectCreate {
	if i != nil {
		scec.SetCreatedAt(*i)
	}
	return scec
}

// SetUpdatedBy sets the "updated_by" field.
func (scec *SmartChargingEffectCreate) SetUpdatedBy(d datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetUpdatedBy(d)
	return scec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableUpdatedBy(d *datasource.UUID) *SmartChargingEffectCreate {
	if d != nil {
		scec.SetUpdatedBy(*d)
	}
	return scec
}

// SetUpdatedAt sets the "updated_at" field.
func (scec *SmartChargingEffectCreate) SetUpdatedAt(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetUpdatedAt(i)
	return scec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableUpdatedAt(i *int64) *SmartChargingEffectCreate {
	if i != nil {
		scec.SetUpdatedAt(*i)
	}
	return scec
}

// SetSmartID sets the "smart_id" field.
func (scec *SmartChargingEffectCreate) SetSmartID(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetSmartID(i)
	return scec
}

// SetStartTime sets the "start_time" field.
func (scec *SmartChargingEffectCreate) SetStartTime(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetStartTime(i)
	return scec
}

// SetPid sets the "pid" field.
func (scec *SmartChargingEffectCreate) SetPid(d datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetPid(d)
	return scec
}

// SetUnit sets the "unit" field.
func (scec *SmartChargingEffectCreate) SetUnit(s string) *SmartChargingEffectCreate {
	scec.mutation.SetUnit(s)
	return scec
}

// SetEquipmentSn sets the "equipment_sn" field.
func (scec *SmartChargingEffectCreate) SetEquipmentSn(s string) *SmartChargingEffectCreate {
	scec.mutation.SetEquipmentSn(s)
	return scec
}

// SetValidFrom sets the "valid_from" field.
func (scec *SmartChargingEffectCreate) SetValidFrom(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetValidFrom(i)
	return scec
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableValidFrom(i *int64) *SmartChargingEffectCreate {
	if i != nil {
		scec.SetValidFrom(*i)
	}
	return scec
}

// SetValidTo sets the "valid_to" field.
func (scec *SmartChargingEffectCreate) SetValidTo(i int64) *SmartChargingEffectCreate {
	scec.mutation.SetValidTo(i)
	return scec
}

// SetNillableValidTo sets the "valid_to" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableValidTo(i *int64) *SmartChargingEffectCreate {
	if i != nil {
		scec.SetValidTo(*i)
	}
	return scec
}

// SetSpec sets the "spec" field.
func (scec *SmartChargingEffectCreate) SetSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEffectCreate {
	scec.mutation.SetSpec(tsp)
	return scec
}

// SetID sets the "id" field.
func (scec *SmartChargingEffectCreate) SetID(d datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetID(d)
	return scec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableID(d *datasource.UUID) *SmartChargingEffectCreate {
	if d != nil {
		scec.SetID(*d)
	}
	return scec
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (scec *SmartChargingEffectCreate) SetEquipmentID(id datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetEquipmentID(id)
	return scec
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (scec *SmartChargingEffectCreate) SetEquipment(e *Equipment) *SmartChargingEffectCreate {
	return scec.SetEquipmentID(e.ID)
}

// SetConnectorID sets the "connector" edge to the Connector entity by ID.
func (scec *SmartChargingEffectCreate) SetConnectorID(id datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetConnectorID(id)
	return scec
}

// SetConnector sets the "connector" edge to the Connector entity.
func (scec *SmartChargingEffectCreate) SetConnector(c *Connector) *SmartChargingEffectCreate {
	return scec.SetConnectorID(c.ID)
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (scec *SmartChargingEffectCreate) SetOrderInfoID(id datasource.UUID) *SmartChargingEffectCreate {
	scec.mutation.SetOrderInfoID(id)
	return scec
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (scec *SmartChargingEffectCreate) SetNillableOrderInfoID(id *datasource.UUID) *SmartChargingEffectCreate {
	if id != nil {
		scec = scec.SetOrderInfoID(*id)
	}
	return scec
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (scec *SmartChargingEffectCreate) SetOrderInfo(o *OrderInfo) *SmartChargingEffectCreate {
	return scec.SetOrderInfoID(o.ID)
}

// Mutation returns the SmartChargingEffectMutation object of the builder.
func (scec *SmartChargingEffectCreate) Mutation() *SmartChargingEffectMutation {
	return scec.mutation
}

// Save creates the SmartChargingEffect in the database.
func (scec *SmartChargingEffectCreate) Save(ctx context.Context) (*SmartChargingEffect, error) {
	var (
		err  error
		node *SmartChargingEffect
	)
	scec.defaults()
	if len(scec.hooks) == 0 {
		if err = scec.check(); err != nil {
			return nil, err
		}
		node, err = scec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scec.check(); err != nil {
				return nil, err
			}
			scec.mutation = mutation
			if node, err = scec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(scec.hooks) - 1; i >= 0; i-- {
			if scec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, scec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SmartChargingEffect)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SmartChargingEffectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (scec *SmartChargingEffectCreate) SaveX(ctx context.Context) *SmartChargingEffect {
	v, err := scec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scec *SmartChargingEffectCreate) Exec(ctx context.Context) error {
	_, err := scec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scec *SmartChargingEffectCreate) ExecX(ctx context.Context) {
	if err := scec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scec *SmartChargingEffectCreate) defaults() {
	if _, ok := scec.mutation.Version(); !ok {
		v := smartchargingeffect.DefaultVersion
		scec.mutation.SetVersion(v)
	}
	if _, ok := scec.mutation.CreatedBy(); !ok {
		v := smartchargingeffect.DefaultCreatedBy
		scec.mutation.SetCreatedBy(v)
	}
	if _, ok := scec.mutation.CreatedAt(); !ok {
		v := smartchargingeffect.DefaultCreatedAt
		scec.mutation.SetCreatedAt(v)
	}
	if _, ok := scec.mutation.UpdatedBy(); !ok {
		v := smartchargingeffect.DefaultUpdatedBy
		scec.mutation.SetUpdatedBy(v)
	}
	if _, ok := scec.mutation.UpdatedAt(); !ok {
		v := smartchargingeffect.DefaultUpdatedAt
		scec.mutation.SetUpdatedAt(v)
	}
	if _, ok := scec.mutation.ID(); !ok {
		v := smartchargingeffect.DefaultID
		scec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scec *SmartChargingEffectCreate) check() error {
	if _, ok := scec.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "SmartChargingEffect.version"`)}
	}
	if _, ok := scec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "SmartChargingEffect.created_by"`)}
	}
	if _, ok := scec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SmartChargingEffect.created_at"`)}
	}
	if _, ok := scec.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "SmartChargingEffect.updated_by"`)}
	}
	if _, ok := scec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SmartChargingEffect.updated_at"`)}
	}
	if _, ok := scec.mutation.SmartID(); !ok {
		return &ValidationError{Name: "smart_id", err: errors.New(`ent: missing required field "SmartChargingEffect.smart_id"`)}
	}
	if _, ok := scec.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "SmartChargingEffect.start_time"`)}
	}
	if _, ok := scec.mutation.Pid(); !ok {
		return &ValidationError{Name: "pid", err: errors.New(`ent: missing required field "SmartChargingEffect.pid"`)}
	}
	if _, ok := scec.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "SmartChargingEffect.unit"`)}
	}
	if _, ok := scec.mutation.EquipmentSn(); !ok {
		return &ValidationError{Name: "equipment_sn", err: errors.New(`ent: missing required field "SmartChargingEffect.equipment_sn"`)}
	}
	if _, ok := scec.mutation.Spec(); !ok {
		return &ValidationError{Name: "spec", err: errors.New(`ent: missing required field "SmartChargingEffect.spec"`)}
	}
	if _, ok := scec.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "SmartChargingEffect.equipment"`)}
	}
	if _, ok := scec.mutation.ConnectorID(); !ok {
		return &ValidationError{Name: "connector", err: errors.New(`ent: missing required edge "SmartChargingEffect.connector"`)}
	}
	return nil
}

func (scec *SmartChargingEffectCreate) sqlSave(ctx context.Context) (*SmartChargingEffect, error) {
	_node, _spec := scec.createSpec()
	if err := sqlgraph.CreateNode(ctx, scec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = datasource.UUID(id)
	}
	return _node, nil
}

func (scec *SmartChargingEffectCreate) createSpec() (*SmartChargingEffect, *sqlgraph.CreateSpec) {
	var (
		_node = &SmartChargingEffect{config: scec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: smartchargingeffect.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: smartchargingeffect.FieldID,
			},
		}
	)
	if id, ok := scec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := scec.mutation.Version(); ok {
		_spec.SetField(smartchargingeffect.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := scec.mutation.CreatedBy(); ok {
		_spec.SetField(smartchargingeffect.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := scec.mutation.CreatedAt(); ok {
		_spec.SetField(smartchargingeffect.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := scec.mutation.UpdatedBy(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := scec.mutation.UpdatedAt(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := scec.mutation.SmartID(); ok {
		_spec.SetField(smartchargingeffect.FieldSmartID, field.TypeInt64, value)
		_node.SmartID = value
	}
	if value, ok := scec.mutation.StartTime(); ok {
		_spec.SetField(smartchargingeffect.FieldStartTime, field.TypeInt64, value)
		_node.StartTime = value
	}
	if value, ok := scec.mutation.Pid(); ok {
		_spec.SetField(smartchargingeffect.FieldPid, field.TypeUint64, value)
		_node.Pid = value
	}
	if value, ok := scec.mutation.Unit(); ok {
		_spec.SetField(smartchargingeffect.FieldUnit, field.TypeString, value)
		_node.Unit = value
	}
	if value, ok := scec.mutation.EquipmentSn(); ok {
		_spec.SetField(smartchargingeffect.FieldEquipmentSn, field.TypeString, value)
		_node.EquipmentSn = value
	}
	if value, ok := scec.mutation.ValidFrom(); ok {
		_spec.SetField(smartchargingeffect.FieldValidFrom, field.TypeInt64, value)
		_node.ValidFrom = &value
	}
	if value, ok := scec.mutation.ValidTo(); ok {
		_spec.SetField(smartchargingeffect.FieldValidTo, field.TypeInt64, value)
		_node.ValidTo = &value
	}
	if value, ok := scec.mutation.Spec(); ok {
		_spec.SetField(smartchargingeffect.FieldSpec, field.TypeJSON, value)
		_node.Spec = value
	}
	if nodes := scec.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.EquipmentTable,
			Columns: []string{smartchargingeffect.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
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
	if nodes := scec.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.ConnectorTable,
			Columns: []string{smartchargingeffect.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.connector_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scec.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   smartchargingeffect.OrderInfoTable,
			Columns: []string{smartchargingeffect.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SmartChargingEffectCreateBulk is the builder for creating many SmartChargingEffect entities in bulk.
type SmartChargingEffectCreateBulk struct {
	config
	builders []*SmartChargingEffectCreate
}

// Save creates the SmartChargingEffect entities in the database.
func (scecb *SmartChargingEffectCreateBulk) Save(ctx context.Context) ([]*SmartChargingEffect, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scecb.builders))
	nodes := make([]*SmartChargingEffect, len(scecb.builders))
	mutators := make([]Mutator, len(scecb.builders))
	for i := range scecb.builders {
		func(i int, root context.Context) {
			builder := scecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SmartChargingEffectMutation)
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
					_, err = mutators[i+1].Mutate(root, scecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = datasource.UUID(id)
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
		if _, err := mutators[0].Mutate(ctx, scecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scecb *SmartChargingEffectCreateBulk) SaveX(ctx context.Context) []*SmartChargingEffect {
	v, err := scecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scecb *SmartChargingEffectCreateBulk) Exec(ctx context.Context) error {
	_, err := scecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scecb *SmartChargingEffectCreateBulk) ExecX(ctx context.Context) {
	if err := scecb.Exec(ctx); err != nil {
		panic(err)
	}
}

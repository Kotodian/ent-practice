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
	"github.com/Kotodian/cwmodel/ent/equipmentalarm"
	"github.com/Kotodian/cwmodel/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/cwmodel/ent/equipmentinfo"
	"github.com/Kotodian/cwmodel/ent/equipmentiot"
	"github.com/Kotodian/cwmodel/ent/equipmentlog"
	"github.com/Kotodian/cwmodel/ent/evse"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/cwmodel/ent/reservation"
	"github.com/Kotodian/cwmodel/ent/smartchargingeffect"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentCreate is the builder for creating a Equipment entity.
type EquipmentCreate struct {
	config
	mutation *EquipmentMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (ec *EquipmentCreate) SetVersion(i int64) *EquipmentCreate {
	ec.mutation.SetVersion(i)
	return ec
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableVersion(i *int64) *EquipmentCreate {
	if i != nil {
		ec.SetVersion(*i)
	}
	return ec
}

// SetCreatedBy sets the "created_by" field.
func (ec *EquipmentCreate) SetCreatedBy(d datasource.UUID) *EquipmentCreate {
	ec.mutation.SetCreatedBy(d)
	return ec
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableCreatedBy(d *datasource.UUID) *EquipmentCreate {
	if d != nil {
		ec.SetCreatedBy(*d)
	}
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EquipmentCreate) SetCreatedAt(i int64) *EquipmentCreate {
	ec.mutation.SetCreatedAt(i)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableCreatedAt(i *int64) *EquipmentCreate {
	if i != nil {
		ec.SetCreatedAt(*i)
	}
	return ec
}

// SetUpdatedBy sets the "updated_by" field.
func (ec *EquipmentCreate) SetUpdatedBy(d datasource.UUID) *EquipmentCreate {
	ec.mutation.SetUpdatedBy(d)
	return ec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableUpdatedBy(d *datasource.UUID) *EquipmentCreate {
	if d != nil {
		ec.SetUpdatedBy(*d)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EquipmentCreate) SetUpdatedAt(i int64) *EquipmentCreate {
	ec.mutation.SetUpdatedAt(i)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableUpdatedAt(i *int64) *EquipmentCreate {
	if i != nil {
		ec.SetUpdatedAt(*i)
	}
	return ec
}

// SetSn sets the "sn" field.
func (ec *EquipmentCreate) SetSn(s string) *EquipmentCreate {
	ec.mutation.SetSn(s)
	return ec
}

// SetOperatorID sets the "operator_id" field.
func (ec *EquipmentCreate) SetOperatorID(d datasource.UUID) *EquipmentCreate {
	ec.mutation.SetOperatorID(d)
	return ec
}

// SetStationID sets the "station_id" field.
func (ec *EquipmentCreate) SetStationID(d datasource.UUID) *EquipmentCreate {
	ec.mutation.SetStationID(d)
	return ec
}

// SetID sets the "id" field.
func (ec *EquipmentCreate) SetID(d datasource.UUID) *EquipmentCreate {
	ec.mutation.SetID(d)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableID(d *datasource.UUID) *EquipmentCreate {
	if d != nil {
		ec.SetID(*d)
	}
	return ec
}

// SetEquipmentInfoID sets the "equipment_info" edge to the EquipmentInfo entity by ID.
func (ec *EquipmentCreate) SetEquipmentInfoID(id datasource.UUID) *EquipmentCreate {
	ec.mutation.SetEquipmentInfoID(id)
	return ec
}

// SetNillableEquipmentInfoID sets the "equipment_info" edge to the EquipmentInfo entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEquipmentInfoID(id *datasource.UUID) *EquipmentCreate {
	if id != nil {
		ec = ec.SetEquipmentInfoID(*id)
	}
	return ec
}

// SetEquipmentInfo sets the "equipment_info" edge to the EquipmentInfo entity.
func (ec *EquipmentCreate) SetEquipmentInfo(e *EquipmentInfo) *EquipmentCreate {
	return ec.SetEquipmentInfoID(e.ID)
}

// AddEvseIDs adds the "evse" edge to the Evse entity by IDs.
func (ec *EquipmentCreate) AddEvseIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddEvseIDs(ids...)
	return ec
}

// AddEvse adds the "evse" edges to the Evse entity.
func (ec *EquipmentCreate) AddEvse(e ...*Evse) *EquipmentCreate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEvseIDs(ids...)
}

// AddConnectorIDs adds the "connector" edge to the Connector entity by IDs.
func (ec *EquipmentCreate) AddConnectorIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddConnectorIDs(ids...)
	return ec
}

// AddConnector adds the "connector" edges to the Connector entity.
func (ec *EquipmentCreate) AddConnector(c ...*Connector) *EquipmentCreate {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddConnectorIDs(ids...)
}

// AddEquipmentAlarmIDs adds the "equipment_alarm" edge to the EquipmentAlarm entity by IDs.
func (ec *EquipmentCreate) AddEquipmentAlarmIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddEquipmentAlarmIDs(ids...)
	return ec
}

// AddEquipmentAlarm adds the "equipment_alarm" edges to the EquipmentAlarm entity.
func (ec *EquipmentCreate) AddEquipmentAlarm(e ...*EquipmentAlarm) *EquipmentCreate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentAlarmIDs(ids...)
}

// SetEquipmentIotID sets the "equipment_iot" edge to the EquipmentIot entity by ID.
func (ec *EquipmentCreate) SetEquipmentIotID(id datasource.UUID) *EquipmentCreate {
	ec.mutation.SetEquipmentIotID(id)
	return ec
}

// SetNillableEquipmentIotID sets the "equipment_iot" edge to the EquipmentIot entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEquipmentIotID(id *datasource.UUID) *EquipmentCreate {
	if id != nil {
		ec = ec.SetEquipmentIotID(*id)
	}
	return ec
}

// SetEquipmentIot sets the "equipment_iot" edge to the EquipmentIot entity.
func (ec *EquipmentCreate) SetEquipmentIot(e *EquipmentIot) *EquipmentCreate {
	return ec.SetEquipmentIotID(e.ID)
}

// AddEquipmentFirmwareEffectIDs adds the "equipment_firmware_effect" edge to the EquipmentFirmwareEffect entity by IDs.
func (ec *EquipmentCreate) AddEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddEquipmentFirmwareEffectIDs(ids...)
	return ec
}

// AddEquipmentFirmwareEffect adds the "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (ec *EquipmentCreate) AddEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *EquipmentCreate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentFirmwareEffectIDs(ids...)
}

// AddOrderInfoIDs adds the "order_info" edge to the OrderInfo entity by IDs.
func (ec *EquipmentCreate) AddOrderInfoIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddOrderInfoIDs(ids...)
	return ec
}

// AddOrderInfo adds the "order_info" edges to the OrderInfo entity.
func (ec *EquipmentCreate) AddOrderInfo(o ...*OrderInfo) *EquipmentCreate {
	ids := make([]datasource.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ec.AddOrderInfoIDs(ids...)
}

// AddReservationIDs adds the "reservation" edge to the Reservation entity by IDs.
func (ec *EquipmentCreate) AddReservationIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddReservationIDs(ids...)
	return ec
}

// AddReservation adds the "reservation" edges to the Reservation entity.
func (ec *EquipmentCreate) AddReservation(r ...*Reservation) *EquipmentCreate {
	ids := make([]datasource.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddReservationIDs(ids...)
}

// AddEquipmentLogIDs adds the "equipment_log" edge to the EquipmentLog entity by IDs.
func (ec *EquipmentCreate) AddEquipmentLogIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddEquipmentLogIDs(ids...)
	return ec
}

// AddEquipmentLog adds the "equipment_log" edges to the EquipmentLog entity.
func (ec *EquipmentCreate) AddEquipmentLog(e ...*EquipmentLog) *EquipmentCreate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentLogIDs(ids...)
}

// AddSmartChargingEffectIDs adds the "smart_charging_effect" edge to the SmartChargingEffect entity by IDs.
func (ec *EquipmentCreate) AddSmartChargingEffectIDs(ids ...datasource.UUID) *EquipmentCreate {
	ec.mutation.AddSmartChargingEffectIDs(ids...)
	return ec
}

// AddSmartChargingEffect adds the "smart_charging_effect" edges to the SmartChargingEffect entity.
func (ec *EquipmentCreate) AddSmartChargingEffect(s ...*SmartChargingEffect) *EquipmentCreate {
	ids := make([]datasource.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ec.AddSmartChargingEffectIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (ec *EquipmentCreate) Mutation() *EquipmentMutation {
	return ec.mutation
}

// Save creates the Equipment in the database.
func (ec *EquipmentCreate) Save(ctx context.Context) (*Equipment, error) {
	var (
		err  error
		node *Equipment
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Equipment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EquipmentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EquipmentCreate) SaveX(ctx context.Context) *Equipment {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EquipmentCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EquipmentCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EquipmentCreate) defaults() {
	if _, ok := ec.mutation.Version(); !ok {
		v := equipment.DefaultVersion
		ec.mutation.SetVersion(v)
	}
	if _, ok := ec.mutation.CreatedBy(); !ok {
		v := equipment.DefaultCreatedBy
		ec.mutation.SetCreatedBy(v)
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := equipment.DefaultCreatedAt
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedBy(); !ok {
		v := equipment.DefaultUpdatedBy
		ec.mutation.SetUpdatedBy(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := equipment.DefaultUpdatedAt
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := equipment.DefaultID
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EquipmentCreate) check() error {
	if _, ok := ec.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Equipment.version"`)}
	}
	if _, ok := ec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Equipment.created_by"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Equipment.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Equipment.updated_by"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Equipment.updated_at"`)}
	}
	if _, ok := ec.mutation.Sn(); !ok {
		return &ValidationError{Name: "sn", err: errors.New(`ent: missing required field "Equipment.sn"`)}
	}
	if v, ok := ec.mutation.Sn(); ok {
		if err := equipment.SnValidator(v); err != nil {
			return &ValidationError{Name: "sn", err: fmt.Errorf(`ent: validator failed for field "Equipment.sn": %w`, err)}
		}
	}
	if _, ok := ec.mutation.OperatorID(); !ok {
		return &ValidationError{Name: "operator_id", err: errors.New(`ent: missing required field "Equipment.operator_id"`)}
	}
	if _, ok := ec.mutation.StationID(); !ok {
		return &ValidationError{Name: "station_id", err: errors.New(`ent: missing required field "Equipment.station_id"`)}
	}
	return nil
}

func (ec *EquipmentCreate) sqlSave(ctx context.Context) (*Equipment, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *EquipmentCreate) createSpec() (*Equipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Equipment{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipment.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Version(); ok {
		_spec.SetField(equipment.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := ec.mutation.CreatedBy(); ok {
		_spec.SetField(equipment.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(equipment.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedBy(); ok {
		_spec.SetField(equipment.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(equipment.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.Sn(); ok {
		_spec.SetField(equipment.FieldSn, field.TypeString, value)
		_node.Sn = value
	}
	if value, ok := ec.mutation.OperatorID(); ok {
		_spec.SetField(equipment.FieldOperatorID, field.TypeUint64, value)
		_node.OperatorID = value
	}
	if value, ok := ec.mutation.StationID(); ok {
		_spec.SetField(equipment.FieldStationID, field.TypeUint64, value)
		_node.StationID = value
	}
	if nodes := ec.mutation.EquipmentInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.EquipmentInfoTable,
			Columns: []string{equipment.EquipmentInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EvseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.EvseTable,
			Columns: []string{equipment.EvseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: evse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.ConnectorTable,
			Columns: []string{equipment.ConnectorColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentAlarmIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.EquipmentAlarmTable,
			Columns: []string{equipment.EquipmentAlarmColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentalarm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentIotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.EquipmentIotTable,
			Columns: []string{equipment.EquipmentIotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentiot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentFirmwareEffectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.EquipmentFirmwareEffectTable,
			Columns: []string{equipment.EquipmentFirmwareEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentfirmwareeffect.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.OrderInfoTable,
			Columns: []string{equipment.OrderInfoColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ReservationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.ReservationTable,
			Columns: []string{equipment.ReservationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: reservation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.EquipmentLogTable,
			Columns: []string{equipment.EquipmentLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.SmartChargingEffectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipment.SmartChargingEffectTable,
			Columns: []string{equipment.SmartChargingEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: smartchargingeffect.FieldID,
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

// EquipmentCreateBulk is the builder for creating many Equipment entities in bulk.
type EquipmentCreateBulk struct {
	config
	builders []*EquipmentCreate
}

// Save creates the Equipment entities in the database.
func (ecb *EquipmentCreateBulk) Save(ctx context.Context) ([]*Equipment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Equipment, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) SaveX(ctx context.Context) []*Equipment {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EquipmentCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

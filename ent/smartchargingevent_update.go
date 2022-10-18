// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/ent-practice/ent/smartchargingevent"
	"github.com/Kotodian/ent-practice/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

// SmartChargingEventUpdate is the builder for updating SmartChargingEvent entities.
type SmartChargingEventUpdate struct {
	config
	hooks    []Hook
	mutation *SmartChargingEventMutation
}

// Where appends a list predicates to the SmartChargingEventUpdate builder.
func (sceu *SmartChargingEventUpdate) Where(ps ...predicate.SmartChargingEvent) *SmartChargingEventUpdate {
	sceu.mutation.Where(ps...)
	return sceu
}

// SetSmartID sets the "smart_id" field.
func (sceu *SmartChargingEventUpdate) SetSmartID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.ResetSmartID()
	sceu.mutation.SetSmartID(d)
	return sceu
}

// AddSmartID adds d to the "smart_id" field.
func (sceu *SmartChargingEventUpdate) AddSmartID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.AddSmartID(d)
	return sceu
}

// SetEquipmentID sets the "equipment_id" field.
func (sceu *SmartChargingEventUpdate) SetEquipmentID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.ResetEquipmentID()
	sceu.mutation.SetEquipmentID(d)
	return sceu
}

// AddEquipmentID adds d to the "equipment_id" field.
func (sceu *SmartChargingEventUpdate) AddEquipmentID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.AddEquipmentID(d)
	return sceu
}

// SetConnectorID sets the "connector_id" field.
func (sceu *SmartChargingEventUpdate) SetConnectorID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.ResetConnectorID()
	sceu.mutation.SetConnectorID(d)
	return sceu
}

// AddConnectorID adds d to the "connector_id" field.
func (sceu *SmartChargingEventUpdate) AddConnectorID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.AddConnectorID(d)
	return sceu
}

// SetOrderID sets the "order_id" field.
func (sceu *SmartChargingEventUpdate) SetOrderID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.ResetOrderID()
	sceu.mutation.SetOrderID(d)
	return sceu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (sceu *SmartChargingEventUpdate) SetNillableOrderID(d *datasource.UUID) *SmartChargingEventUpdate {
	if d != nil {
		sceu.SetOrderID(*d)
	}
	return sceu
}

// AddOrderID adds d to the "order_id" field.
func (sceu *SmartChargingEventUpdate) AddOrderID(d datasource.UUID) *SmartChargingEventUpdate {
	sceu.mutation.AddOrderID(d)
	return sceu
}

// ClearOrderID clears the value of the "order_id" field.
func (sceu *SmartChargingEventUpdate) ClearOrderID() *SmartChargingEventUpdate {
	sceu.mutation.ClearOrderID()
	return sceu
}

// SetUnit sets the "unit" field.
func (sceu *SmartChargingEventUpdate) SetUnit(s string) *SmartChargingEventUpdate {
	sceu.mutation.SetUnit(s)
	return sceu
}

// SetValidFrom sets the "valid_from" field.
func (sceu *SmartChargingEventUpdate) SetValidFrom(i int64) *SmartChargingEventUpdate {
	sceu.mutation.ResetValidFrom()
	sceu.mutation.SetValidFrom(i)
	return sceu
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (sceu *SmartChargingEventUpdate) SetNillableValidFrom(i *int64) *SmartChargingEventUpdate {
	if i != nil {
		sceu.SetValidFrom(*i)
	}
	return sceu
}

// AddValidFrom adds i to the "valid_from" field.
func (sceu *SmartChargingEventUpdate) AddValidFrom(i int64) *SmartChargingEventUpdate {
	sceu.mutation.AddValidFrom(i)
	return sceu
}

// ClearValidFrom clears the value of the "valid_from" field.
func (sceu *SmartChargingEventUpdate) ClearValidFrom() *SmartChargingEventUpdate {
	sceu.mutation.ClearValidFrom()
	return sceu
}

// SetValidTo sets the "valid_to" field.
func (sceu *SmartChargingEventUpdate) SetValidTo(i int64) *SmartChargingEventUpdate {
	sceu.mutation.ResetValidTo()
	sceu.mutation.SetValidTo(i)
	return sceu
}

// SetNillableValidTo sets the "valid_to" field if the given value is not nil.
func (sceu *SmartChargingEventUpdate) SetNillableValidTo(i *int64) *SmartChargingEventUpdate {
	if i != nil {
		sceu.SetValidTo(*i)
	}
	return sceu
}

// AddValidTo adds i to the "valid_to" field.
func (sceu *SmartChargingEventUpdate) AddValidTo(i int64) *SmartChargingEventUpdate {
	sceu.mutation.AddValidTo(i)
	return sceu
}

// ClearValidTo clears the value of the "valid_to" field.
func (sceu *SmartChargingEventUpdate) ClearValidTo() *SmartChargingEventUpdate {
	sceu.mutation.ClearValidTo()
	return sceu
}

// SetSpec sets the "spec" field.
func (sceu *SmartChargingEventUpdate) SetSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEventUpdate {
	sceu.mutation.SetSpec(tsp)
	return sceu
}

// Mutation returns the SmartChargingEventMutation object of the builder.
func (sceu *SmartChargingEventUpdate) Mutation() *SmartChargingEventMutation {
	return sceu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sceu *SmartChargingEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sceu.hooks) == 0 {
		affected, err = sceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sceu.mutation = mutation
			affected, err = sceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sceu.hooks) - 1; i >= 0; i-- {
			if sceu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sceu *SmartChargingEventUpdate) SaveX(ctx context.Context) int {
	affected, err := sceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sceu *SmartChargingEventUpdate) Exec(ctx context.Context) error {
	_, err := sceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sceu *SmartChargingEventUpdate) ExecX(ctx context.Context) {
	if err := sceu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sceu *SmartChargingEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smartchargingevent.Table,
			Columns: smartchargingevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: smartchargingevent.FieldID,
			},
		},
	}
	if ps := sceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sceu.mutation.SmartID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldSmartID,
		})
	}
	if value, ok := sceu.mutation.AddedSmartID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldSmartID,
		})
	}
	if value, ok := sceu.mutation.EquipmentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldEquipmentID,
		})
	}
	if value, ok := sceu.mutation.AddedEquipmentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldEquipmentID,
		})
	}
	if value, ok := sceu.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldConnectorID,
		})
	}
	if value, ok := sceu.mutation.AddedConnectorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldConnectorID,
		})
	}
	if value, ok := sceu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if value, ok := sceu.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if sceu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if value, ok := sceu.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smartchargingevent.FieldUnit,
		})
	}
	if value, ok := sceu.mutation.ValidFrom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if value, ok := sceu.mutation.AddedValidFrom(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if sceu.mutation.ValidFromCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if value, ok := sceu.mutation.ValidTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if value, ok := sceu.mutation.AddedValidTo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if sceu.mutation.ValidToCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if value, ok := sceu.mutation.Spec(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: smartchargingevent.FieldSpec,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smartchargingevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SmartChargingEventUpdateOne is the builder for updating a single SmartChargingEvent entity.
type SmartChargingEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SmartChargingEventMutation
}

// SetSmartID sets the "smart_id" field.
func (sceuo *SmartChargingEventUpdateOne) SetSmartID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetSmartID()
	sceuo.mutation.SetSmartID(d)
	return sceuo
}

// AddSmartID adds d to the "smart_id" field.
func (sceuo *SmartChargingEventUpdateOne) AddSmartID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddSmartID(d)
	return sceuo
}

// SetEquipmentID sets the "equipment_id" field.
func (sceuo *SmartChargingEventUpdateOne) SetEquipmentID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetEquipmentID()
	sceuo.mutation.SetEquipmentID(d)
	return sceuo
}

// AddEquipmentID adds d to the "equipment_id" field.
func (sceuo *SmartChargingEventUpdateOne) AddEquipmentID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddEquipmentID(d)
	return sceuo
}

// SetConnectorID sets the "connector_id" field.
func (sceuo *SmartChargingEventUpdateOne) SetConnectorID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetConnectorID()
	sceuo.mutation.SetConnectorID(d)
	return sceuo
}

// AddConnectorID adds d to the "connector_id" field.
func (sceuo *SmartChargingEventUpdateOne) AddConnectorID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddConnectorID(d)
	return sceuo
}

// SetOrderID sets the "order_id" field.
func (sceuo *SmartChargingEventUpdateOne) SetOrderID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetOrderID()
	sceuo.mutation.SetOrderID(d)
	return sceuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (sceuo *SmartChargingEventUpdateOne) SetNillableOrderID(d *datasource.UUID) *SmartChargingEventUpdateOne {
	if d != nil {
		sceuo.SetOrderID(*d)
	}
	return sceuo
}

// AddOrderID adds d to the "order_id" field.
func (sceuo *SmartChargingEventUpdateOne) AddOrderID(d datasource.UUID) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddOrderID(d)
	return sceuo
}

// ClearOrderID clears the value of the "order_id" field.
func (sceuo *SmartChargingEventUpdateOne) ClearOrderID() *SmartChargingEventUpdateOne {
	sceuo.mutation.ClearOrderID()
	return sceuo
}

// SetUnit sets the "unit" field.
func (sceuo *SmartChargingEventUpdateOne) SetUnit(s string) *SmartChargingEventUpdateOne {
	sceuo.mutation.SetUnit(s)
	return sceuo
}

// SetValidFrom sets the "valid_from" field.
func (sceuo *SmartChargingEventUpdateOne) SetValidFrom(i int64) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetValidFrom()
	sceuo.mutation.SetValidFrom(i)
	return sceuo
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (sceuo *SmartChargingEventUpdateOne) SetNillableValidFrom(i *int64) *SmartChargingEventUpdateOne {
	if i != nil {
		sceuo.SetValidFrom(*i)
	}
	return sceuo
}

// AddValidFrom adds i to the "valid_from" field.
func (sceuo *SmartChargingEventUpdateOne) AddValidFrom(i int64) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddValidFrom(i)
	return sceuo
}

// ClearValidFrom clears the value of the "valid_from" field.
func (sceuo *SmartChargingEventUpdateOne) ClearValidFrom() *SmartChargingEventUpdateOne {
	sceuo.mutation.ClearValidFrom()
	return sceuo
}

// SetValidTo sets the "valid_to" field.
func (sceuo *SmartChargingEventUpdateOne) SetValidTo(i int64) *SmartChargingEventUpdateOne {
	sceuo.mutation.ResetValidTo()
	sceuo.mutation.SetValidTo(i)
	return sceuo
}

// SetNillableValidTo sets the "valid_to" field if the given value is not nil.
func (sceuo *SmartChargingEventUpdateOne) SetNillableValidTo(i *int64) *SmartChargingEventUpdateOne {
	if i != nil {
		sceuo.SetValidTo(*i)
	}
	return sceuo
}

// AddValidTo adds i to the "valid_to" field.
func (sceuo *SmartChargingEventUpdateOne) AddValidTo(i int64) *SmartChargingEventUpdateOne {
	sceuo.mutation.AddValidTo(i)
	return sceuo
}

// ClearValidTo clears the value of the "valid_to" field.
func (sceuo *SmartChargingEventUpdateOne) ClearValidTo() *SmartChargingEventUpdateOne {
	sceuo.mutation.ClearValidTo()
	return sceuo
}

// SetSpec sets the "spec" field.
func (sceuo *SmartChargingEventUpdateOne) SetSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEventUpdateOne {
	sceuo.mutation.SetSpec(tsp)
	return sceuo
}

// Mutation returns the SmartChargingEventMutation object of the builder.
func (sceuo *SmartChargingEventUpdateOne) Mutation() *SmartChargingEventMutation {
	return sceuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sceuo *SmartChargingEventUpdateOne) Select(field string, fields ...string) *SmartChargingEventUpdateOne {
	sceuo.fields = append([]string{field}, fields...)
	return sceuo
}

// Save executes the query and returns the updated SmartChargingEvent entity.
func (sceuo *SmartChargingEventUpdateOne) Save(ctx context.Context) (*SmartChargingEvent, error) {
	var (
		err  error
		node *SmartChargingEvent
	)
	if len(sceuo.hooks) == 0 {
		node, err = sceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sceuo.mutation = mutation
			node, err = sceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sceuo.hooks) - 1; i >= 0; i-- {
			if sceuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sceuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sceuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SmartChargingEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SmartChargingEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sceuo *SmartChargingEventUpdateOne) SaveX(ctx context.Context) *SmartChargingEvent {
	node, err := sceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sceuo *SmartChargingEventUpdateOne) Exec(ctx context.Context) error {
	_, err := sceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sceuo *SmartChargingEventUpdateOne) ExecX(ctx context.Context) {
	if err := sceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sceuo *SmartChargingEventUpdateOne) sqlSave(ctx context.Context) (_node *SmartChargingEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smartchargingevent.Table,
			Columns: smartchargingevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: smartchargingevent.FieldID,
			},
		},
	}
	id, ok := sceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SmartChargingEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smartchargingevent.FieldID)
		for _, f := range fields {
			if !smartchargingevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != smartchargingevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sceuo.mutation.SmartID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldSmartID,
		})
	}
	if value, ok := sceuo.mutation.AddedSmartID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldSmartID,
		})
	}
	if value, ok := sceuo.mutation.EquipmentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldEquipmentID,
		})
	}
	if value, ok := sceuo.mutation.AddedEquipmentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldEquipmentID,
		})
	}
	if value, ok := sceuo.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldConnectorID,
		})
	}
	if value, ok := sceuo.mutation.AddedConnectorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldConnectorID,
		})
	}
	if value, ok := sceuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if value, ok := sceuo.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if sceuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: smartchargingevent.FieldOrderID,
		})
	}
	if value, ok := sceuo.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smartchargingevent.FieldUnit,
		})
	}
	if value, ok := sceuo.mutation.ValidFrom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if value, ok := sceuo.mutation.AddedValidFrom(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if sceuo.mutation.ValidFromCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: smartchargingevent.FieldValidFrom,
		})
	}
	if value, ok := sceuo.mutation.ValidTo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if value, ok := sceuo.mutation.AddedValidTo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if sceuo.mutation.ValidToCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: smartchargingevent.FieldValidTo,
		})
	}
	if value, ok := sceuo.mutation.Spec(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: smartchargingevent.FieldSpec,
		})
	}
	_node = &SmartChargingEvent{config: sceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smartchargingevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

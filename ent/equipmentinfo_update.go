// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentinfo"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentInfoUpdate is the builder for updating EquipmentInfo entities.
type EquipmentInfoUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentInfoMutation
}

// Where appends a list predicates to the EquipmentInfoUpdate builder.
func (eiu *EquipmentInfoUpdate) Where(ps ...predicate.EquipmentInfo) *EquipmentInfoUpdate {
	eiu.mutation.Where(ps...)
	return eiu
}

// SetEquipmentSn sets the "equipment_sn" field.
func (eiu *EquipmentInfoUpdate) SetEquipmentSn(s string) *EquipmentInfoUpdate {
	eiu.mutation.SetEquipmentSn(s)
	return eiu
}

// SetModelID sets the "model_id" field.
func (eiu *EquipmentInfoUpdate) SetModelID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.ResetModelID()
	eiu.mutation.SetModelID(d)
	return eiu
}

// AddModelID adds d to the "model_id" field.
func (eiu *EquipmentInfoUpdate) AddModelID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.AddModelID(d)
	return eiu
}

// SetManufacturerID sets the "manufacturer_id" field.
func (eiu *EquipmentInfoUpdate) SetManufacturerID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.ResetManufacturerID()
	eiu.mutation.SetManufacturerID(d)
	return eiu
}

// AddManufacturerID adds d to the "manufacturer_id" field.
func (eiu *EquipmentInfoUpdate) AddManufacturerID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.AddManufacturerID(d)
	return eiu
}

// SetFirmwareID sets the "firmware_id" field.
func (eiu *EquipmentInfoUpdate) SetFirmwareID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.ResetFirmwareID()
	eiu.mutation.SetFirmwareID(d)
	return eiu
}

// AddFirmwareID adds d to the "firmware_id" field.
func (eiu *EquipmentInfoUpdate) AddFirmwareID(d datasource.UUID) *EquipmentInfoUpdate {
	eiu.mutation.AddFirmwareID(d)
	return eiu
}

// SetAccessPod sets the "access_pod" field.
func (eiu *EquipmentInfoUpdate) SetAccessPod(s string) *EquipmentInfoUpdate {
	eiu.mutation.SetAccessPod(s)
	return eiu
}

// SetState sets the "state" field.
func (eiu *EquipmentInfoUpdate) SetState(b bool) *EquipmentInfoUpdate {
	eiu.mutation.SetState(b)
	return eiu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eiu *EquipmentInfoUpdate) SetEquipmentID(id int) *EquipmentInfoUpdate {
	eiu.mutation.SetEquipmentID(id)
	return eiu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eiu *EquipmentInfoUpdate) SetEquipment(e *Equipment) *EquipmentInfoUpdate {
	return eiu.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentInfoMutation object of the builder.
func (eiu *EquipmentInfoUpdate) Mutation() *EquipmentInfoMutation {
	return eiu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eiu *EquipmentInfoUpdate) ClearEquipment() *EquipmentInfoUpdate {
	eiu.mutation.ClearEquipment()
	return eiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eiu *EquipmentInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eiu.hooks) == 0 {
		if err = eiu.check(); err != nil {
			return 0, err
		}
		affected, err = eiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eiu.check(); err != nil {
				return 0, err
			}
			eiu.mutation = mutation
			affected, err = eiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eiu.hooks) - 1; i >= 0; i-- {
			if eiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eiu *EquipmentInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := eiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eiu *EquipmentInfoUpdate) Exec(ctx context.Context) error {
	_, err := eiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiu *EquipmentInfoUpdate) ExecX(ctx context.Context) {
	if err := eiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eiu *EquipmentInfoUpdate) check() error {
	if _, ok := eiu.mutation.EquipmentID(); eiu.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentInfo.equipment"`)
	}
	return nil
}

func (eiu *EquipmentInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentinfo.Table,
			Columns: equipmentinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentinfo.FieldID,
			},
		},
	}
	if ps := eiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiu.mutation.EquipmentSn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldEquipmentSn,
		})
	}
	if value, ok := eiu.mutation.ModelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldModelID,
		})
	}
	if value, ok := eiu.mutation.AddedModelID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldModelID,
		})
	}
	if value, ok := eiu.mutation.ManufacturerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldManufacturerID,
		})
	}
	if value, ok := eiu.mutation.AddedManufacturerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldManufacturerID,
		})
	}
	if value, ok := eiu.mutation.FirmwareID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldFirmwareID,
		})
	}
	if value, ok := eiu.mutation.AddedFirmwareID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldFirmwareID,
		})
	}
	if value, ok := eiu.mutation.AccessPod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldAccessPod,
		})
	}
	if value, ok := eiu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: equipmentinfo.FieldState,
		})
	}
	if eiu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentinfo.EquipmentTable,
			Columns: []string{equipmentinfo.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentinfo.EquipmentTable,
			Columns: []string{equipmentinfo.EquipmentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EquipmentInfoUpdateOne is the builder for updating a single EquipmentInfo entity.
type EquipmentInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentInfoMutation
}

// SetEquipmentSn sets the "equipment_sn" field.
func (eiuo *EquipmentInfoUpdateOne) SetEquipmentSn(s string) *EquipmentInfoUpdateOne {
	eiuo.mutation.SetEquipmentSn(s)
	return eiuo
}

// SetModelID sets the "model_id" field.
func (eiuo *EquipmentInfoUpdateOne) SetModelID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.ResetModelID()
	eiuo.mutation.SetModelID(d)
	return eiuo
}

// AddModelID adds d to the "model_id" field.
func (eiuo *EquipmentInfoUpdateOne) AddModelID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.AddModelID(d)
	return eiuo
}

// SetManufacturerID sets the "manufacturer_id" field.
func (eiuo *EquipmentInfoUpdateOne) SetManufacturerID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.ResetManufacturerID()
	eiuo.mutation.SetManufacturerID(d)
	return eiuo
}

// AddManufacturerID adds d to the "manufacturer_id" field.
func (eiuo *EquipmentInfoUpdateOne) AddManufacturerID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.AddManufacturerID(d)
	return eiuo
}

// SetFirmwareID sets the "firmware_id" field.
func (eiuo *EquipmentInfoUpdateOne) SetFirmwareID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.ResetFirmwareID()
	eiuo.mutation.SetFirmwareID(d)
	return eiuo
}

// AddFirmwareID adds d to the "firmware_id" field.
func (eiuo *EquipmentInfoUpdateOne) AddFirmwareID(d datasource.UUID) *EquipmentInfoUpdateOne {
	eiuo.mutation.AddFirmwareID(d)
	return eiuo
}

// SetAccessPod sets the "access_pod" field.
func (eiuo *EquipmentInfoUpdateOne) SetAccessPod(s string) *EquipmentInfoUpdateOne {
	eiuo.mutation.SetAccessPod(s)
	return eiuo
}

// SetState sets the "state" field.
func (eiuo *EquipmentInfoUpdateOne) SetState(b bool) *EquipmentInfoUpdateOne {
	eiuo.mutation.SetState(b)
	return eiuo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eiuo *EquipmentInfoUpdateOne) SetEquipmentID(id int) *EquipmentInfoUpdateOne {
	eiuo.mutation.SetEquipmentID(id)
	return eiuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eiuo *EquipmentInfoUpdateOne) SetEquipment(e *Equipment) *EquipmentInfoUpdateOne {
	return eiuo.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentInfoMutation object of the builder.
func (eiuo *EquipmentInfoUpdateOne) Mutation() *EquipmentInfoMutation {
	return eiuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eiuo *EquipmentInfoUpdateOne) ClearEquipment() *EquipmentInfoUpdateOne {
	eiuo.mutation.ClearEquipment()
	return eiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eiuo *EquipmentInfoUpdateOne) Select(field string, fields ...string) *EquipmentInfoUpdateOne {
	eiuo.fields = append([]string{field}, fields...)
	return eiuo
}

// Save executes the query and returns the updated EquipmentInfo entity.
func (eiuo *EquipmentInfoUpdateOne) Save(ctx context.Context) (*EquipmentInfo, error) {
	var (
		err  error
		node *EquipmentInfo
	)
	if len(eiuo.hooks) == 0 {
		if err = eiuo.check(); err != nil {
			return nil, err
		}
		node, err = eiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eiuo.check(); err != nil {
				return nil, err
			}
			eiuo.mutation = mutation
			node, err = eiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eiuo.hooks) - 1; i >= 0; i-- {
			if eiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eiuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EquipmentInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EquipmentInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eiuo *EquipmentInfoUpdateOne) SaveX(ctx context.Context) *EquipmentInfo {
	node, err := eiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eiuo *EquipmentInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := eiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiuo *EquipmentInfoUpdateOne) ExecX(ctx context.Context) {
	if err := eiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eiuo *EquipmentInfoUpdateOne) check() error {
	if _, ok := eiuo.mutation.EquipmentID(); eiuo.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentInfo.equipment"`)
	}
	return nil
}

func (eiuo *EquipmentInfoUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentinfo.Table,
			Columns: equipmentinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentinfo.FieldID,
			},
		},
	}
	id, ok := eiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentinfo.FieldID)
		for _, f := range fields {
			if !equipmentinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmentinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiuo.mutation.EquipmentSn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldEquipmentSn,
		})
	}
	if value, ok := eiuo.mutation.ModelID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldModelID,
		})
	}
	if value, ok := eiuo.mutation.AddedModelID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldModelID,
		})
	}
	if value, ok := eiuo.mutation.ManufacturerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldManufacturerID,
		})
	}
	if value, ok := eiuo.mutation.AddedManufacturerID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldManufacturerID,
		})
	}
	if value, ok := eiuo.mutation.FirmwareID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldFirmwareID,
		})
	}
	if value, ok := eiuo.mutation.AddedFirmwareID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: equipmentinfo.FieldFirmwareID,
		})
	}
	if value, ok := eiuo.mutation.AccessPod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldAccessPod,
		})
	}
	if value, ok := eiuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: equipmentinfo.FieldState,
		})
	}
	if eiuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentinfo.EquipmentTable,
			Columns: []string{equipmentinfo.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentinfo.EquipmentTable,
			Columns: []string{equipmentinfo.EquipmentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EquipmentInfo{config: eiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

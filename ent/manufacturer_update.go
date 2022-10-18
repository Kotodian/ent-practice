// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/manufacturer"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// ManufacturerUpdate is the builder for updating Manufacturer entities.
type ManufacturerUpdate struct {
	config
	hooks    []Hook
	mutation *ManufacturerMutation
}

// Where appends a list predicates to the ManufacturerUpdate builder.
func (mu *ManufacturerUpdate) Where(ps ...predicate.Manufacturer) *ManufacturerUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCode sets the "code" field.
func (mu *ManufacturerUpdate) SetCode(s string) *ManufacturerUpdate {
	mu.mutation.SetCode(s)
	return mu
}

// SetName sets the "name" field.
func (mu *ManufacturerUpdate) SetName(s string) *ManufacturerUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *ManufacturerUpdate) SetNillableName(s *string) *ManufacturerUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// ClearName clears the value of the "name" field.
func (mu *ManufacturerUpdate) ClearName() *ManufacturerUpdate {
	mu.mutation.ClearName()
	return mu
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (mu *ManufacturerUpdate) AddEquipmentIDs(ids ...int) *ManufacturerUpdate {
	mu.mutation.AddEquipmentIDs(ids...)
	return mu
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (mu *ManufacturerUpdate) AddEquipment(e ...*Equipment) *ManufacturerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mu.AddEquipmentIDs(ids...)
}

// Mutation returns the ManufacturerMutation object of the builder.
func (mu *ManufacturerUpdate) Mutation() *ManufacturerMutation {
	return mu.mutation
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (mu *ManufacturerUpdate) ClearEquipment() *ManufacturerUpdate {
	mu.mutation.ClearEquipment()
	return mu
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (mu *ManufacturerUpdate) RemoveEquipmentIDs(ids ...int) *ManufacturerUpdate {
	mu.mutation.RemoveEquipmentIDs(ids...)
	return mu
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (mu *ManufacturerUpdate) RemoveEquipment(e ...*Equipment) *ManufacturerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mu.RemoveEquipmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ManufacturerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ManufacturerUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ManufacturerUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ManufacturerUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *ManufacturerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manufacturer.Table,
			Columns: manufacturer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manufacturer.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldCode,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldName,
		})
	}
	if mu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: manufacturer.FieldName,
		})
	}
	if mu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
	if nodes := mu.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !mu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manufacturer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ManufacturerUpdateOne is the builder for updating a single Manufacturer entity.
type ManufacturerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ManufacturerMutation
}

// SetCode sets the "code" field.
func (muo *ManufacturerUpdateOne) SetCode(s string) *ManufacturerUpdateOne {
	muo.mutation.SetCode(s)
	return muo
}

// SetName sets the "name" field.
func (muo *ManufacturerUpdateOne) SetName(s string) *ManufacturerUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *ManufacturerUpdateOne) SetNillableName(s *string) *ManufacturerUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// ClearName clears the value of the "name" field.
func (muo *ManufacturerUpdateOne) ClearName() *ManufacturerUpdateOne {
	muo.mutation.ClearName()
	return muo
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (muo *ManufacturerUpdateOne) AddEquipmentIDs(ids ...int) *ManufacturerUpdateOne {
	muo.mutation.AddEquipmentIDs(ids...)
	return muo
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (muo *ManufacturerUpdateOne) AddEquipment(e ...*Equipment) *ManufacturerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return muo.AddEquipmentIDs(ids...)
}

// Mutation returns the ManufacturerMutation object of the builder.
func (muo *ManufacturerUpdateOne) Mutation() *ManufacturerMutation {
	return muo.mutation
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (muo *ManufacturerUpdateOne) ClearEquipment() *ManufacturerUpdateOne {
	muo.mutation.ClearEquipment()
	return muo
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (muo *ManufacturerUpdateOne) RemoveEquipmentIDs(ids ...int) *ManufacturerUpdateOne {
	muo.mutation.RemoveEquipmentIDs(ids...)
	return muo
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (muo *ManufacturerUpdateOne) RemoveEquipment(e ...*Equipment) *ManufacturerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return muo.RemoveEquipmentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ManufacturerUpdateOne) Select(field string, fields ...string) *ManufacturerUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Manufacturer entity.
func (muo *ManufacturerUpdateOne) Save(ctx context.Context) (*Manufacturer, error) {
	var (
		err  error
		node *Manufacturer
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ManufacturerUpdateOne) SaveX(ctx context.Context) *Manufacturer {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ManufacturerUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ManufacturerUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *ManufacturerUpdateOne) sqlSave(ctx context.Context) (_node *Manufacturer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manufacturer.Table,
			Columns: manufacturer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manufacturer.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Manufacturer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, manufacturer.FieldID)
		for _, f := range fields {
			if !manufacturer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != manufacturer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldCode,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldName,
		})
	}
	if muo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: manufacturer.FieldName,
		})
	}
	if muo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
	if nodes := muo.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !muo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.EquipmentTable,
			Columns: []string{manufacturer.EquipmentColumn},
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
	_node = &Manufacturer{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manufacturer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/orderevent"
	"github.com/Kotodian/ent-practice/ent/orderinfo"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// OrderEventUpdate is the builder for updating OrderEvent entities.
type OrderEventUpdate struct {
	config
	hooks    []Hook
	mutation *OrderEventMutation
}

// Where appends a list predicates to the OrderEventUpdate builder.
func (oeu *OrderEventUpdate) Where(ps ...predicate.OrderEvent) *OrderEventUpdate {
	oeu.mutation.Where(ps...)
	return oeu
}

// SetVersion sets the "version" field.
func (oeu *OrderEventUpdate) SetVersion(i int64) *OrderEventUpdate {
	oeu.mutation.ResetVersion()
	oeu.mutation.SetVersion(i)
	return oeu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (oeu *OrderEventUpdate) SetNillableVersion(i *int64) *OrderEventUpdate {
	if i != nil {
		oeu.SetVersion(*i)
	}
	return oeu
}

// AddVersion adds i to the "version" field.
func (oeu *OrderEventUpdate) AddVersion(i int64) *OrderEventUpdate {
	oeu.mutation.AddVersion(i)
	return oeu
}

// SetUpdatedBy sets the "updated_by" field.
func (oeu *OrderEventUpdate) SetUpdatedBy(d datasource.UUID) *OrderEventUpdate {
	oeu.mutation.ResetUpdatedBy()
	oeu.mutation.SetUpdatedBy(d)
	return oeu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oeu *OrderEventUpdate) SetNillableUpdatedBy(d *datasource.UUID) *OrderEventUpdate {
	if d != nil {
		oeu.SetUpdatedBy(*d)
	}
	return oeu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (oeu *OrderEventUpdate) AddUpdatedBy(d datasource.UUID) *OrderEventUpdate {
	oeu.mutation.AddUpdatedBy(d)
	return oeu
}

// SetUpdatedAt sets the "updated_at" field.
func (oeu *OrderEventUpdate) SetUpdatedAt(i int64) *OrderEventUpdate {
	oeu.mutation.ResetUpdatedAt()
	oeu.mutation.SetUpdatedAt(i)
	return oeu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oeu *OrderEventUpdate) SetNillableUpdatedAt(i *int64) *OrderEventUpdate {
	if i != nil {
		oeu.SetUpdatedAt(*i)
	}
	return oeu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (oeu *OrderEventUpdate) AddUpdatedAt(i int64) *OrderEventUpdate {
	oeu.mutation.AddUpdatedAt(i)
	return oeu
}

// SetOrderID sets the "order_id" field.
func (oeu *OrderEventUpdate) SetOrderID(d datasource.UUID) *OrderEventUpdate {
	oeu.mutation.ResetOrderID()
	oeu.mutation.SetOrderID(d)
	return oeu
}

// AddOrderID adds d to the "order_id" field.
func (oeu *OrderEventUpdate) AddOrderID(d datasource.UUID) *OrderEventUpdate {
	oeu.mutation.AddOrderID(d)
	return oeu
}

// SetContent sets the "content" field.
func (oeu *OrderEventUpdate) SetContent(s string) *OrderEventUpdate {
	oeu.mutation.SetContent(s)
	return oeu
}

// SetOccurrence sets the "occurrence" field.
func (oeu *OrderEventUpdate) SetOccurrence(i int64) *OrderEventUpdate {
	oeu.mutation.ResetOccurrence()
	oeu.mutation.SetOccurrence(i)
	return oeu
}

// AddOccurrence adds i to the "occurrence" field.
func (oeu *OrderEventUpdate) AddOccurrence(i int64) *OrderEventUpdate {
	oeu.mutation.AddOccurrence(i)
	return oeu
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (oeu *OrderEventUpdate) SetOrderInfoID(id datasource.UUID) *OrderEventUpdate {
	oeu.mutation.SetOrderInfoID(id)
	return oeu
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (oeu *OrderEventUpdate) SetNillableOrderInfoID(id *datasource.UUID) *OrderEventUpdate {
	if id != nil {
		oeu = oeu.SetOrderInfoID(*id)
	}
	return oeu
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (oeu *OrderEventUpdate) SetOrderInfo(o *OrderInfo) *OrderEventUpdate {
	return oeu.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderEventMutation object of the builder.
func (oeu *OrderEventUpdate) Mutation() *OrderEventMutation {
	return oeu.mutation
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (oeu *OrderEventUpdate) ClearOrderInfo() *OrderEventUpdate {
	oeu.mutation.ClearOrderInfo()
	return oeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oeu *OrderEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oeu.hooks) == 0 {
		affected, err = oeu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oeu.mutation = mutation
			affected, err = oeu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oeu.hooks) - 1; i >= 0; i-- {
			if oeu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oeu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oeu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oeu *OrderEventUpdate) SaveX(ctx context.Context) int {
	affected, err := oeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oeu *OrderEventUpdate) Exec(ctx context.Context) error {
	_, err := oeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeu *OrderEventUpdate) ExecX(ctx context.Context) {
	if err := oeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oeu *OrderEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderevent.Table,
			Columns: orderevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderevent.FieldID,
			},
		},
	}
	if ps := oeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldVersion,
		})
	}
	if value, ok := oeu.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldVersion,
		})
	}
	if value, ok := oeu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldUpdatedBy,
		})
	}
	if value, ok := oeu.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldUpdatedBy,
		})
	}
	if value, ok := oeu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldUpdatedAt,
		})
	}
	if value, ok := oeu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldUpdatedAt,
		})
	}
	if value, ok := oeu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldOrderID,
		})
	}
	if value, ok := oeu.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldOrderID,
		})
	}
	if value, ok := oeu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderevent.FieldContent,
		})
	}
	if value, ok := oeu.mutation.Occurrence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldOccurrence,
		})
	}
	if value, ok := oeu.mutation.AddedOccurrence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldOccurrence,
		})
	}
	if oeu.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderevent.OrderInfoTable,
			Columns: []string{orderevent.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeu.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderevent.OrderInfoTable,
			Columns: []string{orderevent.OrderInfoColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderEventUpdateOne is the builder for updating a single OrderEvent entity.
type OrderEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderEventMutation
}

// SetVersion sets the "version" field.
func (oeuo *OrderEventUpdateOne) SetVersion(i int64) *OrderEventUpdateOne {
	oeuo.mutation.ResetVersion()
	oeuo.mutation.SetVersion(i)
	return oeuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (oeuo *OrderEventUpdateOne) SetNillableVersion(i *int64) *OrderEventUpdateOne {
	if i != nil {
		oeuo.SetVersion(*i)
	}
	return oeuo
}

// AddVersion adds i to the "version" field.
func (oeuo *OrderEventUpdateOne) AddVersion(i int64) *OrderEventUpdateOne {
	oeuo.mutation.AddVersion(i)
	return oeuo
}

// SetUpdatedBy sets the "updated_by" field.
func (oeuo *OrderEventUpdateOne) SetUpdatedBy(d datasource.UUID) *OrderEventUpdateOne {
	oeuo.mutation.ResetUpdatedBy()
	oeuo.mutation.SetUpdatedBy(d)
	return oeuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oeuo *OrderEventUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *OrderEventUpdateOne {
	if d != nil {
		oeuo.SetUpdatedBy(*d)
	}
	return oeuo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (oeuo *OrderEventUpdateOne) AddUpdatedBy(d datasource.UUID) *OrderEventUpdateOne {
	oeuo.mutation.AddUpdatedBy(d)
	return oeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (oeuo *OrderEventUpdateOne) SetUpdatedAt(i int64) *OrderEventUpdateOne {
	oeuo.mutation.ResetUpdatedAt()
	oeuo.mutation.SetUpdatedAt(i)
	return oeuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oeuo *OrderEventUpdateOne) SetNillableUpdatedAt(i *int64) *OrderEventUpdateOne {
	if i != nil {
		oeuo.SetUpdatedAt(*i)
	}
	return oeuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (oeuo *OrderEventUpdateOne) AddUpdatedAt(i int64) *OrderEventUpdateOne {
	oeuo.mutation.AddUpdatedAt(i)
	return oeuo
}

// SetOrderID sets the "order_id" field.
func (oeuo *OrderEventUpdateOne) SetOrderID(d datasource.UUID) *OrderEventUpdateOne {
	oeuo.mutation.ResetOrderID()
	oeuo.mutation.SetOrderID(d)
	return oeuo
}

// AddOrderID adds d to the "order_id" field.
func (oeuo *OrderEventUpdateOne) AddOrderID(d datasource.UUID) *OrderEventUpdateOne {
	oeuo.mutation.AddOrderID(d)
	return oeuo
}

// SetContent sets the "content" field.
func (oeuo *OrderEventUpdateOne) SetContent(s string) *OrderEventUpdateOne {
	oeuo.mutation.SetContent(s)
	return oeuo
}

// SetOccurrence sets the "occurrence" field.
func (oeuo *OrderEventUpdateOne) SetOccurrence(i int64) *OrderEventUpdateOne {
	oeuo.mutation.ResetOccurrence()
	oeuo.mutation.SetOccurrence(i)
	return oeuo
}

// AddOccurrence adds i to the "occurrence" field.
func (oeuo *OrderEventUpdateOne) AddOccurrence(i int64) *OrderEventUpdateOne {
	oeuo.mutation.AddOccurrence(i)
	return oeuo
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (oeuo *OrderEventUpdateOne) SetOrderInfoID(id datasource.UUID) *OrderEventUpdateOne {
	oeuo.mutation.SetOrderInfoID(id)
	return oeuo
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (oeuo *OrderEventUpdateOne) SetNillableOrderInfoID(id *datasource.UUID) *OrderEventUpdateOne {
	if id != nil {
		oeuo = oeuo.SetOrderInfoID(*id)
	}
	return oeuo
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (oeuo *OrderEventUpdateOne) SetOrderInfo(o *OrderInfo) *OrderEventUpdateOne {
	return oeuo.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderEventMutation object of the builder.
func (oeuo *OrderEventUpdateOne) Mutation() *OrderEventMutation {
	return oeuo.mutation
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (oeuo *OrderEventUpdateOne) ClearOrderInfo() *OrderEventUpdateOne {
	oeuo.mutation.ClearOrderInfo()
	return oeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oeuo *OrderEventUpdateOne) Select(field string, fields ...string) *OrderEventUpdateOne {
	oeuo.fields = append([]string{field}, fields...)
	return oeuo
}

// Save executes the query and returns the updated OrderEvent entity.
func (oeuo *OrderEventUpdateOne) Save(ctx context.Context) (*OrderEvent, error) {
	var (
		err  error
		node *OrderEvent
	)
	if len(oeuo.hooks) == 0 {
		node, err = oeuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oeuo.mutation = mutation
			node, err = oeuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oeuo.hooks) - 1; i >= 0; i-- {
			if oeuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oeuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oeuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oeuo *OrderEventUpdateOne) SaveX(ctx context.Context) *OrderEvent {
	node, err := oeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oeuo *OrderEventUpdateOne) Exec(ctx context.Context) error {
	_, err := oeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeuo *OrderEventUpdateOne) ExecX(ctx context.Context) {
	if err := oeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oeuo *OrderEventUpdateOne) sqlSave(ctx context.Context) (_node *OrderEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderevent.Table,
			Columns: orderevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderevent.FieldID,
			},
		},
	}
	id, ok := oeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderevent.FieldID)
		for _, f := range fields {
			if !orderevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeuo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldVersion,
		})
	}
	if value, ok := oeuo.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldVersion,
		})
	}
	if value, ok := oeuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldUpdatedBy,
		})
	}
	if value, ok := oeuo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldUpdatedBy,
		})
	}
	if value, ok := oeuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldUpdatedAt,
		})
	}
	if value, ok := oeuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldUpdatedAt,
		})
	}
	if value, ok := oeuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldOrderID,
		})
	}
	if value, ok := oeuo.mutation.AddedOrderID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldOrderID,
		})
	}
	if value, ok := oeuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderevent.FieldContent,
		})
	}
	if value, ok := oeuo.mutation.Occurrence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldOccurrence,
		})
	}
	if value, ok := oeuo.mutation.AddedOccurrence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldOccurrence,
		})
	}
	if oeuo.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderevent.OrderInfoTable,
			Columns: []string{orderevent.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oeuo.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderevent.OrderInfoTable,
			Columns: []string{orderevent.OrderInfoColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderEvent{config: oeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

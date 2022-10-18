// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/orderevent"
	"github.com/Kotodian/ent-practice/ent/orderinfo"
	"github.com/Kotodian/gokit/datasource"
)

// OrderEventCreate is the builder for creating a OrderEvent entity.
type OrderEventCreate struct {
	config
	mutation *OrderEventMutation
	hooks    []Hook
}

// SetOrderID sets the "order_id" field.
func (oec *OrderEventCreate) SetOrderID(d datasource.UUID) *OrderEventCreate {
	oec.mutation.SetOrderID(d)
	return oec
}

// SetContent sets the "content" field.
func (oec *OrderEventCreate) SetContent(s string) *OrderEventCreate {
	oec.mutation.SetContent(s)
	return oec
}

// SetOccurrence sets the "occurrence" field.
func (oec *OrderEventCreate) SetOccurrence(i int64) *OrderEventCreate {
	oec.mutation.SetOccurrence(i)
	return oec
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (oec *OrderEventCreate) SetOrderInfoID(id int) *OrderEventCreate {
	oec.mutation.SetOrderInfoID(id)
	return oec
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (oec *OrderEventCreate) SetNillableOrderInfoID(id *int) *OrderEventCreate {
	if id != nil {
		oec = oec.SetOrderInfoID(*id)
	}
	return oec
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (oec *OrderEventCreate) SetOrderInfo(o *OrderInfo) *OrderEventCreate {
	return oec.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderEventMutation object of the builder.
func (oec *OrderEventCreate) Mutation() *OrderEventMutation {
	return oec.mutation
}

// Save creates the OrderEvent in the database.
func (oec *OrderEventCreate) Save(ctx context.Context) (*OrderEvent, error) {
	var (
		err  error
		node *OrderEvent
	)
	if len(oec.hooks) == 0 {
		if err = oec.check(); err != nil {
			return nil, err
		}
		node, err = oec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oec.check(); err != nil {
				return nil, err
			}
			oec.mutation = mutation
			if node, err = oec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oec.hooks) - 1; i >= 0; i-- {
			if oec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oec *OrderEventCreate) SaveX(ctx context.Context) *OrderEvent {
	v, err := oec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oec *OrderEventCreate) Exec(ctx context.Context) error {
	_, err := oec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oec *OrderEventCreate) ExecX(ctx context.Context) {
	if err := oec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oec *OrderEventCreate) check() error {
	if _, ok := oec.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderEvent.order_id"`)}
	}
	if _, ok := oec.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "OrderEvent.content"`)}
	}
	if _, ok := oec.mutation.Occurrence(); !ok {
		return &ValidationError{Name: "occurrence", err: errors.New(`ent: missing required field "OrderEvent.occurrence"`)}
	}
	return nil
}

func (oec *OrderEventCreate) sqlSave(ctx context.Context) (*OrderEvent, error) {
	_node, _spec := oec.createSpec()
	if err := sqlgraph.CreateNode(ctx, oec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oec *OrderEventCreate) createSpec() (*OrderEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderEvent{config: oec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderevent.FieldID,
			},
		}
	)
	if value, ok := oec.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderevent.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := oec.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderevent.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := oec.mutation.Occurrence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderevent.FieldOccurrence,
		})
		_node.Occurrence = value
	}
	if nodes := oec.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderevent.OrderInfoTable,
			Columns: []string{orderevent.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_info_order_event = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderEventCreateBulk is the builder for creating many OrderEvent entities in bulk.
type OrderEventCreateBulk struct {
	config
	builders []*OrderEventCreate
}

// Save creates the OrderEvent entities in the database.
func (oecb *OrderEventCreateBulk) Save(ctx context.Context) ([]*OrderEvent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oecb.builders))
	nodes := make([]*OrderEvent, len(oecb.builders))
	mutators := make([]Mutator, len(oecb.builders))
	for i := range oecb.builders {
		func(i int, root context.Context) {
			builder := oecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderEventMutation)
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
					_, err = mutators[i+1].Mutate(root, oecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oecb *OrderEventCreateBulk) SaveX(ctx context.Context) []*OrderEvent {
	v, err := oecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oecb *OrderEventCreateBulk) Exec(ctx context.Context) error {
	_, err := oecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oecb *OrderEventCreateBulk) ExecX(ctx context.Context) {
	if err := oecb.Exec(ctx); err != nil {
		panic(err)
	}
}
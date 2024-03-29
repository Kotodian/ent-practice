// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/appmoduleinfo"
	"github.com/Kotodian/cwmodel/ent/predicate"
)

// AppModuleInfoQuery is the builder for querying AppModuleInfo entities.
type AppModuleInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppModuleInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppModuleInfoQuery builder.
func (amiq *AppModuleInfoQuery) Where(ps ...predicate.AppModuleInfo) *AppModuleInfoQuery {
	amiq.predicates = append(amiq.predicates, ps...)
	return amiq
}

// Limit adds a limit step to the query.
func (amiq *AppModuleInfoQuery) Limit(limit int) *AppModuleInfoQuery {
	amiq.limit = &limit
	return amiq
}

// Offset adds an offset step to the query.
func (amiq *AppModuleInfoQuery) Offset(offset int) *AppModuleInfoQuery {
	amiq.offset = &offset
	return amiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amiq *AppModuleInfoQuery) Unique(unique bool) *AppModuleInfoQuery {
	amiq.unique = &unique
	return amiq
}

// Order adds an order step to the query.
func (amiq *AppModuleInfoQuery) Order(o ...OrderFunc) *AppModuleInfoQuery {
	amiq.order = append(amiq.order, o...)
	return amiq
}

// First returns the first AppModuleInfo entity from the query.
// Returns a *NotFoundError when no AppModuleInfo was found.
func (amiq *AppModuleInfoQuery) First(ctx context.Context) (*AppModuleInfo, error) {
	nodes, err := amiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appmoduleinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) FirstX(ctx context.Context) *AppModuleInfo {
	node, err := amiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppModuleInfo ID from the query.
// Returns a *NotFoundError when no AppModuleInfo ID was found.
func (amiq *AppModuleInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appmoduleinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := amiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppModuleInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppModuleInfo entity is found.
// Returns a *NotFoundError when no AppModuleInfo entities are found.
func (amiq *AppModuleInfoQuery) Only(ctx context.Context) (*AppModuleInfo, error) {
	nodes, err := amiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appmoduleinfo.Label}
	default:
		return nil, &NotSingularError{appmoduleinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) OnlyX(ctx context.Context) *AppModuleInfo {
	node, err := amiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppModuleInfo ID in the query.
// Returns a *NotSingularError when more than one AppModuleInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (amiq *AppModuleInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appmoduleinfo.Label}
	default:
		err = &NotSingularError{appmoduleinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := amiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppModuleInfos.
func (amiq *AppModuleInfoQuery) All(ctx context.Context) ([]*AppModuleInfo, error) {
	if err := amiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return amiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) AllX(ctx context.Context) []*AppModuleInfo {
	nodes, err := amiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppModuleInfo IDs.
func (amiq *AppModuleInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := amiq.Select(appmoduleinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := amiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amiq *AppModuleInfoQuery) Count(ctx context.Context) (int, error) {
	if err := amiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return amiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) CountX(ctx context.Context) int {
	count, err := amiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amiq *AppModuleInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := amiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return amiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (amiq *AppModuleInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := amiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppModuleInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amiq *AppModuleInfoQuery) Clone() *AppModuleInfoQuery {
	if amiq == nil {
		return nil
	}
	return &AppModuleInfoQuery{
		config:     amiq.config,
		limit:      amiq.limit,
		offset:     amiq.offset,
		order:      append([]OrderFunc{}, amiq.order...),
		predicates: append([]predicate.AppModuleInfo{}, amiq.predicates...),
		// clone intermediate query.
		sql:    amiq.sql.Clone(),
		path:   amiq.path,
		unique: amiq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppModuleInfo.Query().
//		GroupBy(appmoduleinfo.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (amiq *AppModuleInfoQuery) GroupBy(field string, fields ...string) *AppModuleInfoGroupBy {
	grbuild := &AppModuleInfoGroupBy{config: amiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := amiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return amiq.sqlQuery(ctx), nil
	}
	grbuild.label = appmoduleinfo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.AppModuleInfo.Query().
//		Select(appmoduleinfo.FieldName).
//		Scan(ctx, &v)
func (amiq *AppModuleInfoQuery) Select(fields ...string) *AppModuleInfoSelect {
	amiq.fields = append(amiq.fields, fields...)
	selbuild := &AppModuleInfoSelect{AppModuleInfoQuery: amiq}
	selbuild.label = appmoduleinfo.Label
	selbuild.flds, selbuild.scan = &amiq.fields, selbuild.Scan
	return selbuild
}

func (amiq *AppModuleInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range amiq.fields {
		if !appmoduleinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if amiq.path != nil {
		prev, err := amiq.path(ctx)
		if err != nil {
			return err
		}
		amiq.sql = prev
	}
	return nil
}

func (amiq *AppModuleInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppModuleInfo, error) {
	var (
		nodes = []*AppModuleInfo{}
		_spec = amiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppModuleInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppModuleInfo{config: amiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, amiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (amiq *AppModuleInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := amiq.querySpec()
	_spec.Node.Columns = amiq.fields
	if len(amiq.fields) > 0 {
		_spec.Unique = amiq.unique != nil && *amiq.unique
	}
	return sqlgraph.CountNodes(ctx, amiq.driver, _spec)
}

func (amiq *AppModuleInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := amiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (amiq *AppModuleInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appmoduleinfo.Table,
			Columns: appmoduleinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appmoduleinfo.FieldID,
			},
		},
		From:   amiq.sql,
		Unique: true,
	}
	if unique := amiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := amiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appmoduleinfo.FieldID)
		for i := range fields {
			if fields[i] != appmoduleinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := amiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amiq *AppModuleInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amiq.driver.Dialect())
	t1 := builder.Table(appmoduleinfo.Table)
	columns := amiq.fields
	if len(columns) == 0 {
		columns = appmoduleinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if amiq.sql != nil {
		selector = amiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if amiq.unique != nil && *amiq.unique {
		selector.Distinct()
	}
	for _, p := range amiq.predicates {
		p(selector)
	}
	for _, p := range amiq.order {
		p(selector)
	}
	if offset := amiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppModuleInfoGroupBy is the group-by builder for AppModuleInfo entities.
type AppModuleInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amigb *AppModuleInfoGroupBy) Aggregate(fns ...AggregateFunc) *AppModuleInfoGroupBy {
	amigb.fns = append(amigb.fns, fns...)
	return amigb
}

// Scan applies the group-by query and scans the result into the given value.
func (amigb *AppModuleInfoGroupBy) Scan(ctx context.Context, v any) error {
	query, err := amigb.path(ctx)
	if err != nil {
		return err
	}
	amigb.sql = query
	return amigb.sqlScan(ctx, v)
}

func (amigb *AppModuleInfoGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range amigb.fields {
		if !appmoduleinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := amigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amigb *AppModuleInfoGroupBy) sqlQuery() *sql.Selector {
	selector := amigb.sql.Select()
	aggregation := make([]string, 0, len(amigb.fns))
	for _, fn := range amigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(amigb.fields)+len(amigb.fns))
		for _, f := range amigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(amigb.fields...)...)
}

// AppModuleInfoSelect is the builder for selecting fields of AppModuleInfo entities.
type AppModuleInfoSelect struct {
	*AppModuleInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (amis *AppModuleInfoSelect) Scan(ctx context.Context, v any) error {
	if err := amis.prepareQuery(ctx); err != nil {
		return err
	}
	amis.sql = amis.AppModuleInfoQuery.sqlQuery(ctx)
	return amis.sqlScan(ctx, v)
}

func (amis *AppModuleInfoSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := amis.sql.Query()
	if err := amis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

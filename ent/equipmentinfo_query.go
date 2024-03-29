// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/equipment"
	"github.com/Kotodian/cwmodel/ent/equipmentinfo"
	"github.com/Kotodian/cwmodel/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentInfoQuery is the builder for querying EquipmentInfo entities.
type EquipmentInfoQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.EquipmentInfo
	withEquipment *EquipmentQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentInfoQuery builder.
func (eiq *EquipmentInfoQuery) Where(ps ...predicate.EquipmentInfo) *EquipmentInfoQuery {
	eiq.predicates = append(eiq.predicates, ps...)
	return eiq
}

// Limit adds a limit step to the query.
func (eiq *EquipmentInfoQuery) Limit(limit int) *EquipmentInfoQuery {
	eiq.limit = &limit
	return eiq
}

// Offset adds an offset step to the query.
func (eiq *EquipmentInfoQuery) Offset(offset int) *EquipmentInfoQuery {
	eiq.offset = &offset
	return eiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eiq *EquipmentInfoQuery) Unique(unique bool) *EquipmentInfoQuery {
	eiq.unique = &unique
	return eiq
}

// Order adds an order step to the query.
func (eiq *EquipmentInfoQuery) Order(o ...OrderFunc) *EquipmentInfoQuery {
	eiq.order = append(eiq.order, o...)
	return eiq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (eiq *EquipmentInfoQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: eiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentinfo.Table, equipmentinfo.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, equipmentinfo.EquipmentTable, equipmentinfo.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentInfo entity from the query.
// Returns a *NotFoundError when no EquipmentInfo was found.
func (eiq *EquipmentInfoQuery) First(ctx context.Context) (*EquipmentInfo, error) {
	nodes, err := eiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) FirstX(ctx context.Context) *EquipmentInfo {
	node, err := eiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentInfo ID from the query.
// Returns a *NotFoundError when no EquipmentInfo ID was found.
func (eiq *EquipmentInfoQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := eiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentInfo entity is found.
// Returns a *NotFoundError when no EquipmentInfo entities are found.
func (eiq *EquipmentInfoQuery) Only(ctx context.Context) (*EquipmentInfo, error) {
	nodes, err := eiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentinfo.Label}
	default:
		return nil, &NotSingularError{equipmentinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) OnlyX(ctx context.Context) *EquipmentInfo {
	node, err := eiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentInfo ID in the query.
// Returns a *NotSingularError when more than one EquipmentInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (eiq *EquipmentInfoQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentinfo.Label}
	default:
		err = &NotSingularError{equipmentinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := eiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentInfos.
func (eiq *EquipmentInfoQuery) All(ctx context.Context) ([]*EquipmentInfo, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) AllX(ctx context.Context) []*EquipmentInfo {
	nodes, err := eiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentInfo IDs.
func (eiq *EquipmentInfoQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := eiq.Select(equipmentinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := eiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eiq *EquipmentInfoQuery) Count(ctx context.Context) (int, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) CountX(ctx context.Context) int {
	count, err := eiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eiq *EquipmentInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eiq *EquipmentInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := eiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eiq *EquipmentInfoQuery) Clone() *EquipmentInfoQuery {
	if eiq == nil {
		return nil
	}
	return &EquipmentInfoQuery{
		config:        eiq.config,
		limit:         eiq.limit,
		offset:        eiq.offset,
		order:         append([]OrderFunc{}, eiq.order...),
		predicates:    append([]predicate.EquipmentInfo{}, eiq.predicates...),
		withEquipment: eiq.withEquipment.Clone(),
		// clone intermediate query.
		sql:    eiq.sql.Clone(),
		path:   eiq.path,
		unique: eiq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *EquipmentInfoQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentInfoQuery {
	query := &EquipmentQuery{config: eiq.config}
	for _, opt := range opts {
		opt(query)
	}
	eiq.withEquipment = query
	return eiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentInfo.Query().
//		GroupBy(equipmentinfo.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eiq *EquipmentInfoQuery) GroupBy(field string, fields ...string) *EquipmentInfoGroupBy {
	grbuild := &EquipmentInfoGroupBy{config: eiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eiq.sqlQuery(ctx), nil
	}
	grbuild.label = equipmentinfo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version,omitempty"`
//	}
//
//	client.EquipmentInfo.Query().
//		Select(equipmentinfo.FieldVersion).
//		Scan(ctx, &v)
func (eiq *EquipmentInfoQuery) Select(fields ...string) *EquipmentInfoSelect {
	eiq.fields = append(eiq.fields, fields...)
	selbuild := &EquipmentInfoSelect{EquipmentInfoQuery: eiq}
	selbuild.label = equipmentinfo.Label
	selbuild.flds, selbuild.scan = &eiq.fields, selbuild.Scan
	return selbuild
}

func (eiq *EquipmentInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eiq.fields {
		if !equipmentinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eiq.path != nil {
		prev, err := eiq.path(ctx)
		if err != nil {
			return err
		}
		eiq.sql = prev
	}
	return nil
}

func (eiq *EquipmentInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EquipmentInfo, error) {
	var (
		nodes       = []*EquipmentInfo{}
		withFKs     = eiq.withFKs
		_spec       = eiq.querySpec()
		loadedTypes = [1]bool{
			eiq.withEquipment != nil,
		}
	)
	if eiq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentinfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EquipmentInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EquipmentInfo{config: eiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eiq.withEquipment; query != nil {
		if err := eiq.loadEquipment(ctx, query, nodes, nil,
			func(n *EquipmentInfo, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eiq *EquipmentInfoQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*EquipmentInfo, init func(*EquipmentInfo), assign func(*EquipmentInfo, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*EquipmentInfo)
	for i := range nodes {
		if nodes[i].equipment_id == nil {
			continue
		}
		fk := *nodes[i].equipment_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(equipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eiq *EquipmentInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eiq.querySpec()
	_spec.Node.Columns = eiq.fields
	if len(eiq.fields) > 0 {
		_spec.Unique = eiq.unique != nil && *eiq.unique
	}
	return sqlgraph.CountNodes(ctx, eiq.driver, _spec)
}

func (eiq *EquipmentInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eiq *EquipmentInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentinfo.Table,
			Columns: equipmentinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentinfo.FieldID,
			},
		},
		From:   eiq.sql,
		Unique: true,
	}
	if unique := eiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentinfo.FieldID)
		for i := range fields {
			if fields[i] != equipmentinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eiq *EquipmentInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eiq.driver.Dialect())
	t1 := builder.Table(equipmentinfo.Table)
	columns := eiq.fields
	if len(columns) == 0 {
		columns = equipmentinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eiq.sql != nil {
		selector = eiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eiq.unique != nil && *eiq.unique {
		selector.Distinct()
	}
	for _, p := range eiq.predicates {
		p(selector)
	}
	for _, p := range eiq.order {
		p(selector)
	}
	if offset := eiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentInfoGroupBy is the group-by builder for EquipmentInfo entities.
type EquipmentInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eigb *EquipmentInfoGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentInfoGroupBy {
	eigb.fns = append(eigb.fns, fns...)
	return eigb
}

// Scan applies the group-by query and scans the result into the given value.
func (eigb *EquipmentInfoGroupBy) Scan(ctx context.Context, v any) error {
	query, err := eigb.path(ctx)
	if err != nil {
		return err
	}
	eigb.sql = query
	return eigb.sqlScan(ctx, v)
}

func (eigb *EquipmentInfoGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range eigb.fields {
		if !equipmentinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := eigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eigb *EquipmentInfoGroupBy) sqlQuery() *sql.Selector {
	selector := eigb.sql.Select()
	aggregation := make([]string, 0, len(eigb.fns))
	for _, fn := range eigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(eigb.fields)+len(eigb.fns))
		for _, f := range eigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(eigb.fields...)...)
}

// EquipmentInfoSelect is the builder for selecting fields of EquipmentInfo entities.
type EquipmentInfoSelect struct {
	*EquipmentInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (eis *EquipmentInfoSelect) Scan(ctx context.Context, v any) error {
	if err := eis.prepareQuery(ctx); err != nil {
		return err
	}
	eis.sql = eis.EquipmentInfoQuery.sqlQuery(ctx)
	return eis.sqlScan(ctx, v)
}

func (eis *EquipmentInfoSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := eis.sql.Query()
	if err := eis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

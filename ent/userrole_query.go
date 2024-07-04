// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/user"
	"trec/ent/userrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserRoleQuery is the builder for querying UserRole entities.
type UserRoleQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.UserRole
	withUserEdge *UserQuery
	withRoleEdge *RoleQuery
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*UserRole) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserRoleQuery builder.
func (urq *UserRoleQuery) Where(ps ...predicate.UserRole) *UserRoleQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit adds a limit step to the query.
func (urq *UserRoleQuery) Limit(limit int) *UserRoleQuery {
	urq.limit = &limit
	return urq
}

// Offset adds an offset step to the query.
func (urq *UserRoleQuery) Offset(offset int) *UserRoleQuery {
	urq.offset = &offset
	return urq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urq *UserRoleQuery) Unique(unique bool) *UserRoleQuery {
	urq.unique = &unique
	return urq
}

// Order adds an order step to the query.
func (urq *UserRoleQuery) Order(o ...OrderFunc) *UserRoleQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// QueryUserEdge chains the current query on the "user_edge" edge.
func (urq *UserRoleQuery) QueryUserEdge() *UserQuery {
	query := &UserQuery{config: urq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userrole.UserEdgeTable, userrole.UserEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoleEdge chains the current query on the "role_edge" edge.
func (urq *UserRoleQuery) QueryRoleEdge() *RoleQuery {
	query := &RoleQuery{config: urq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userrole.RoleEdgeTable, userrole.RoleEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserRole entity from the query.
// Returns a *NotFoundError when no UserRole was found.
func (urq *UserRoleQuery) First(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UserRoleQuery) FirstX(ctx context.Context) *UserRole {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserRole ID from the query.
// Returns a *NotFoundError when no UserRole ID was found.
func (urq *UserRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UserRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserRole entity is found.
// Returns a *NotFoundError when no UserRole entities are found.
func (urq *UserRoleQuery) Only(ctx context.Context) (*UserRole, error) {
	nodes, err := urq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userrole.Label}
	default:
		return nil, &NotSingularError{userrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyX(ctx context.Context) *UserRole {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserRole ID in the query.
// Returns a *NotSingularError when more than one UserRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (urq *UserRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = urq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userrole.Label}
	default:
		err = &NotSingularError{userrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UserRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserRoles.
func (urq *UserRoleQuery) All(ctx context.Context) ([]*UserRole, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return urq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (urq *UserRoleQuery) AllX(ctx context.Context) []*UserRole {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserRole IDs.
func (urq *UserRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := urq.Select(userrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UserRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UserRoleQuery) Count(ctx context.Context) (int, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return urq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UserRoleQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UserRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return urq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UserRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UserRoleQuery) Clone() *UserRoleQuery {
	if urq == nil {
		return nil
	}
	return &UserRoleQuery{
		config:       urq.config,
		limit:        urq.limit,
		offset:       urq.offset,
		order:        append([]OrderFunc{}, urq.order...),
		predicates:   append([]predicate.UserRole{}, urq.predicates...),
		withUserEdge: urq.withUserEdge.Clone(),
		withRoleEdge: urq.withRoleEdge.Clone(),
		// clone intermediate query.
		sql:    urq.sql.Clone(),
		path:   urq.path,
		unique: urq.unique,
	}
}

// WithUserEdge tells the query-builder to eager-load the nodes that are connected to
// the "user_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (urq *UserRoleQuery) WithUserEdge(opts ...func(*UserQuery)) *UserRoleQuery {
	query := &UserQuery{config: urq.config}
	for _, opt := range opts {
		opt(query)
	}
	urq.withUserEdge = query
	return urq
}

// WithRoleEdge tells the query-builder to eager-load the nodes that are connected to
// the "role_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (urq *UserRoleQuery) WithRoleEdge(opts ...func(*RoleQuery)) *UserRoleQuery {
	query := &RoleQuery{config: urq.config}
	for _, opt := range opts {
		opt(query)
	}
	urq.withRoleEdge = query
	return urq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserRole.Query().
//		GroupBy(userrole.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (urq *UserRoleQuery) GroupBy(field string, fields ...string) *UserRoleGroupBy {
	grbuild := &UserRoleGroupBy{config: urq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return urq.sqlQuery(ctx), nil
	}
	grbuild.label = userrole.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserRole.Query().
//		Select(userrole.FieldCreatedAt).
//		Scan(ctx, &v)
func (urq *UserRoleQuery) Select(fields ...string) *UserRoleSelect {
	urq.fields = append(urq.fields, fields...)
	selbuild := &UserRoleSelect{UserRoleQuery: urq}
	selbuild.label = userrole.Label
	selbuild.flds, selbuild.scan = &urq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a UserRoleSelect configured with the given aggregations.
func (urq *UserRoleQuery) Aggregate(fns ...AggregateFunc) *UserRoleSelect {
	return urq.Select().Aggregate(fns...)
}

func (urq *UserRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range urq.fields {
		if !userrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	return nil
}

func (urq *UserRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserRole, error) {
	var (
		nodes       = []*UserRole{}
		_spec       = urq.querySpec()
		loadedTypes = [2]bool{
			urq.withUserEdge != nil,
			urq.withRoleEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserRole{config: urq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := urq.withUserEdge; query != nil {
		if err := urq.loadUserEdge(ctx, query, nodes, nil,
			func(n *UserRole, e *User) { n.Edges.UserEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := urq.withRoleEdge; query != nil {
		if err := urq.loadRoleEdge(ctx, query, nodes, nil,
			func(n *UserRole, e *Role) { n.Edges.RoleEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range urq.loadTotal {
		if err := urq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (urq *UserRoleQuery) loadUserEdge(ctx context.Context, query *UserQuery, nodes []*UserRole, init func(*UserRole), assign func(*UserRole, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserRole)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (urq *UserRoleQuery) loadRoleEdge(ctx context.Context, query *RoleQuery, nodes []*UserRole, init func(*UserRole), assign func(*UserRole, *Role)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserRole)
	for i := range nodes {
		fk := nodes[i].RoleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(role.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "role_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (urq *UserRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	_spec.Node.Columns = urq.fields
	if len(urq.fields) > 0 {
		_spec.Unique = urq.unique != nil && *urq.unique
	}
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UserRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := urq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (urq *UserRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
		From:   urq.sql,
		Unique: true,
	}
	if unique := urq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := urq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrole.FieldID)
		for i := range fields {
			if fields[i] != userrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urq *UserRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(userrole.Table)
	columns := urq.fields
	if len(columns) == 0 {
		columns = userrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urq.unique != nil && *urq.unique {
		selector.Distinct()
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector)
	}
	if offset := urq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserRoleGroupBy is the group-by builder for UserRole entities.
type UserRoleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UserRoleGroupBy) Aggregate(fns ...AggregateFunc) *UserRoleGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the group-by query and scans the result into the given value.
func (urgb *UserRoleGroupBy) Scan(ctx context.Context, v any) error {
	query, err := urgb.path(ctx)
	if err != nil {
		return err
	}
	urgb.sql = query
	return urgb.sqlScan(ctx, v)
}

func (urgb *UserRoleGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range urgb.fields {
		if !userrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := urgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (urgb *UserRoleGroupBy) sqlQuery() *sql.Selector {
	selector := urgb.sql.Select()
	aggregation := make([]string, 0, len(urgb.fns))
	for _, fn := range urgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(urgb.fields)+len(urgb.fns))
		for _, f := range urgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(urgb.fields...)...)
}

// UserRoleSelect is the builder for selecting fields of UserRole entities.
type UserRoleSelect struct {
	*UserRoleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (urs *UserRoleSelect) Aggregate(fns ...AggregateFunc) *UserRoleSelect {
	urs.fns = append(urs.fns, fns...)
	return urs
}

// Scan applies the selector query and scans the result into the given value.
func (urs *UserRoleSelect) Scan(ctx context.Context, v any) error {
	if err := urs.prepareQuery(ctx); err != nil {
		return err
	}
	urs.sql = urs.UserRoleQuery.sqlQuery(ctx)
	return urs.sqlScan(ctx, v)
}

func (urs *UserRoleSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(urs.fns))
	for _, fn := range urs.fns {
		aggregation = append(aggregation, fn(urs.sql))
	}
	switch n := len(*urs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		urs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		urs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := urs.sql.Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

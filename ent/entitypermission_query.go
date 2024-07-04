// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"trec/ent/entitypermission"
	"trec/ent/permission"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntityPermissionQuery is the builder for querying EntityPermission entities.
type EntityPermissionQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.EntityPermission
	withPermissionEdges *PermissionQuery
	withUserEdge        *UserQuery
	withRoleEdge        *RoleQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*EntityPermission) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntityPermissionQuery builder.
func (epq *EntityPermissionQuery) Where(ps ...predicate.EntityPermission) *EntityPermissionQuery {
	epq.predicates = append(epq.predicates, ps...)
	return epq
}

// Limit adds a limit step to the query.
func (epq *EntityPermissionQuery) Limit(limit int) *EntityPermissionQuery {
	epq.limit = &limit
	return epq
}

// Offset adds an offset step to the query.
func (epq *EntityPermissionQuery) Offset(offset int) *EntityPermissionQuery {
	epq.offset = &offset
	return epq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (epq *EntityPermissionQuery) Unique(unique bool) *EntityPermissionQuery {
	epq.unique = &unique
	return epq
}

// Order adds an order step to the query.
func (epq *EntityPermissionQuery) Order(o ...OrderFunc) *EntityPermissionQuery {
	epq.order = append(epq.order, o...)
	return epq
}

// QueryPermissionEdges chains the current query on the "permission_edges" edge.
func (epq *EntityPermissionQuery) QueryPermissionEdges() *PermissionQuery {
	query := &PermissionQuery{config: epq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitypermission.Table, entitypermission.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entitypermission.PermissionEdgesTable, entitypermission.PermissionEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(epq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserEdge chains the current query on the "user_edge" edge.
func (epq *EntityPermissionQuery) QueryUserEdge() *UserQuery {
	query := &UserQuery{config: epq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitypermission.Table, entitypermission.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entitypermission.UserEdgeTable, entitypermission.UserEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(epq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoleEdge chains the current query on the "role_edge" edge.
func (epq *EntityPermissionQuery) QueryRoleEdge() *RoleQuery {
	query := &RoleQuery{config: epq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := epq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := epq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entitypermission.Table, entitypermission.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entitypermission.RoleEdgeTable, entitypermission.RoleEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(epq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntityPermission entity from the query.
// Returns a *NotFoundError when no EntityPermission was found.
func (epq *EntityPermissionQuery) First(ctx context.Context) (*EntityPermission, error) {
	nodes, err := epq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entitypermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epq *EntityPermissionQuery) FirstX(ctx context.Context) *EntityPermission {
	node, err := epq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntityPermission ID from the query.
// Returns a *NotFoundError when no EntityPermission ID was found.
func (epq *EntityPermissionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = epq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entitypermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (epq *EntityPermissionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := epq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntityPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntityPermission entity is found.
// Returns a *NotFoundError when no EntityPermission entities are found.
func (epq *EntityPermissionQuery) Only(ctx context.Context) (*EntityPermission, error) {
	nodes, err := epq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entitypermission.Label}
	default:
		return nil, &NotSingularError{entitypermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epq *EntityPermissionQuery) OnlyX(ctx context.Context) *EntityPermission {
	node, err := epq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntityPermission ID in the query.
// Returns a *NotSingularError when more than one EntityPermission ID is found.
// Returns a *NotFoundError when no entities are found.
func (epq *EntityPermissionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = epq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entitypermission.Label}
	default:
		err = &NotSingularError{entitypermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (epq *EntityPermissionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := epq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntityPermissions.
func (epq *EntityPermissionQuery) All(ctx context.Context) ([]*EntityPermission, error) {
	if err := epq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return epq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (epq *EntityPermissionQuery) AllX(ctx context.Context) []*EntityPermission {
	nodes, err := epq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntityPermission IDs.
func (epq *EntityPermissionQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := epq.Select(entitypermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epq *EntityPermissionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := epq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epq *EntityPermissionQuery) Count(ctx context.Context) (int, error) {
	if err := epq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return epq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (epq *EntityPermissionQuery) CountX(ctx context.Context) int {
	count, err := epq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epq *EntityPermissionQuery) Exist(ctx context.Context) (bool, error) {
	if err := epq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return epq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (epq *EntityPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := epq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntityPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epq *EntityPermissionQuery) Clone() *EntityPermissionQuery {
	if epq == nil {
		return nil
	}
	return &EntityPermissionQuery{
		config:              epq.config,
		limit:               epq.limit,
		offset:              epq.offset,
		order:               append([]OrderFunc{}, epq.order...),
		predicates:          append([]predicate.EntityPermission{}, epq.predicates...),
		withPermissionEdges: epq.withPermissionEdges.Clone(),
		withUserEdge:        epq.withUserEdge.Clone(),
		withRoleEdge:        epq.withRoleEdge.Clone(),
		// clone intermediate query.
		sql:    epq.sql.Clone(),
		path:   epq.path,
		unique: epq.unique,
	}
}

// WithPermissionEdges tells the query-builder to eager-load the nodes that are connected to
// the "permission_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (epq *EntityPermissionQuery) WithPermissionEdges(opts ...func(*PermissionQuery)) *EntityPermissionQuery {
	query := &PermissionQuery{config: epq.config}
	for _, opt := range opts {
		opt(query)
	}
	epq.withPermissionEdges = query
	return epq
}

// WithUserEdge tells the query-builder to eager-load the nodes that are connected to
// the "user_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (epq *EntityPermissionQuery) WithUserEdge(opts ...func(*UserQuery)) *EntityPermissionQuery {
	query := &UserQuery{config: epq.config}
	for _, opt := range opts {
		opt(query)
	}
	epq.withUserEdge = query
	return epq
}

// WithRoleEdge tells the query-builder to eager-load the nodes that are connected to
// the "role_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (epq *EntityPermissionQuery) WithRoleEdge(opts ...func(*RoleQuery)) *EntityPermissionQuery {
	query := &RoleQuery{config: epq.config}
	for _, opt := range opts {
		opt(query)
	}
	epq.withRoleEdge = query
	return epq
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
//	client.EntityPermission.Query().
//		GroupBy(entitypermission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (epq *EntityPermissionQuery) GroupBy(field string, fields ...string) *EntityPermissionGroupBy {
	grbuild := &EntityPermissionGroupBy{config: epq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := epq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return epq.sqlQuery(ctx), nil
	}
	grbuild.label = entitypermission.Label
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
//	client.EntityPermission.Query().
//		Select(entitypermission.FieldCreatedAt).
//		Scan(ctx, &v)
func (epq *EntityPermissionQuery) Select(fields ...string) *EntityPermissionSelect {
	epq.fields = append(epq.fields, fields...)
	selbuild := &EntityPermissionSelect{EntityPermissionQuery: epq}
	selbuild.label = entitypermission.Label
	selbuild.flds, selbuild.scan = &epq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a EntityPermissionSelect configured with the given aggregations.
func (epq *EntityPermissionQuery) Aggregate(fns ...AggregateFunc) *EntityPermissionSelect {
	return epq.Select().Aggregate(fns...)
}

func (epq *EntityPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range epq.fields {
		if !entitypermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if epq.path != nil {
		prev, err := epq.path(ctx)
		if err != nil {
			return err
		}
		epq.sql = prev
	}
	return nil
}

func (epq *EntityPermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntityPermission, error) {
	var (
		nodes       = []*EntityPermission{}
		_spec       = epq.querySpec()
		loadedTypes = [3]bool{
			epq.withPermissionEdges != nil,
			epq.withUserEdge != nil,
			epq.withRoleEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntityPermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntityPermission{config: epq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(epq.modifiers) > 0 {
		_spec.Modifiers = epq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, epq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := epq.withPermissionEdges; query != nil {
		if err := epq.loadPermissionEdges(ctx, query, nodes, nil,
			func(n *EntityPermission, e *Permission) { n.Edges.PermissionEdges = e }); err != nil {
			return nil, err
		}
	}
	if query := epq.withUserEdge; query != nil {
		if err := epq.loadUserEdge(ctx, query, nodes, nil,
			func(n *EntityPermission, e *User) { n.Edges.UserEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := epq.withRoleEdge; query != nil {
		if err := epq.loadRoleEdge(ctx, query, nodes, nil,
			func(n *EntityPermission, e *Role) { n.Edges.RoleEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range epq.loadTotal {
		if err := epq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (epq *EntityPermissionQuery) loadPermissionEdges(ctx context.Context, query *PermissionQuery, nodes []*EntityPermission, init func(*EntityPermission), assign func(*EntityPermission, *Permission)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntityPermission)
	for i := range nodes {
		fk := nodes[i].PermissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(permission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "permission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (epq *EntityPermissionQuery) loadUserEdge(ctx context.Context, query *UserQuery, nodes []*EntityPermission, init func(*EntityPermission), assign func(*EntityPermission, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntityPermission)
	for i := range nodes {
		fk := nodes[i].EntityID
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
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (epq *EntityPermissionQuery) loadRoleEdge(ctx context.Context, query *RoleQuery, nodes []*EntityPermission, init func(*EntityPermission), assign func(*EntityPermission, *Role)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntityPermission)
	for i := range nodes {
		fk := nodes[i].EntityID
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
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (epq *EntityPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := epq.querySpec()
	if len(epq.modifiers) > 0 {
		_spec.Modifiers = epq.modifiers
	}
	_spec.Node.Columns = epq.fields
	if len(epq.fields) > 0 {
		_spec.Unique = epq.unique != nil && *epq.unique
	}
	return sqlgraph.CountNodes(ctx, epq.driver, _spec)
}

func (epq *EntityPermissionQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := epq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (epq *EntityPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitypermission.Table,
			Columns: entitypermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entitypermission.FieldID,
			},
		},
		From:   epq.sql,
		Unique: true,
	}
	if unique := epq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := epq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitypermission.FieldID)
		for i := range fields {
			if fields[i] != entitypermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := epq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := epq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := epq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := epq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (epq *EntityPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(entitypermission.Table)
	columns := epq.fields
	if len(columns) == 0 {
		columns = entitypermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if epq.sql != nil {
		selector = epq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if epq.unique != nil && *epq.unique {
		selector.Distinct()
	}
	for _, p := range epq.predicates {
		p(selector)
	}
	for _, p := range epq.order {
		p(selector)
	}
	if offset := epq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntityPermissionGroupBy is the group-by builder for EntityPermission entities.
type EntityPermissionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epgb *EntityPermissionGroupBy) Aggregate(fns ...AggregateFunc) *EntityPermissionGroupBy {
	epgb.fns = append(epgb.fns, fns...)
	return epgb
}

// Scan applies the group-by query and scans the result into the given value.
func (epgb *EntityPermissionGroupBy) Scan(ctx context.Context, v any) error {
	query, err := epgb.path(ctx)
	if err != nil {
		return err
	}
	epgb.sql = query
	return epgb.sqlScan(ctx, v)
}

func (epgb *EntityPermissionGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range epgb.fields {
		if !entitypermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := epgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (epgb *EntityPermissionGroupBy) sqlQuery() *sql.Selector {
	selector := epgb.sql.Select()
	aggregation := make([]string, 0, len(epgb.fns))
	for _, fn := range epgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(epgb.fields)+len(epgb.fns))
		for _, f := range epgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(epgb.fields...)...)
}

// EntityPermissionSelect is the builder for selecting fields of EntityPermission entities.
type EntityPermissionSelect struct {
	*EntityPermissionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eps *EntityPermissionSelect) Aggregate(fns ...AggregateFunc) *EntityPermissionSelect {
	eps.fns = append(eps.fns, fns...)
	return eps
}

// Scan applies the selector query and scans the result into the given value.
func (eps *EntityPermissionSelect) Scan(ctx context.Context, v any) error {
	if err := eps.prepareQuery(ctx); err != nil {
		return err
	}
	eps.sql = eps.EntityPermissionQuery.sqlQuery(ctx)
	return eps.sqlScan(ctx, v)
}

func (eps *EntityPermissionSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(eps.fns))
	for _, fn := range eps.fns {
		aggregation = append(aggregation, fn(eps.sql))
	}
	switch n := len(*eps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		eps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		eps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := eps.sql.Query()
	if err := eps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

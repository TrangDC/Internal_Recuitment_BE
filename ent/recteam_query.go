// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/recteam"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecTeamQuery is the builder for querying RecTeam entities.
type RecTeamQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.RecTeam
	withRecMemberEdges       *UserQuery
	withRecTeamJobEdges      *HiringJobQuery
	withRecLeaderEdge        *UserQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*RecTeam) error
	withNamedRecMemberEdges  map[string]*UserQuery
	withNamedRecTeamJobEdges map[string]*HiringJobQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecTeamQuery builder.
func (rtq *RecTeamQuery) Where(ps ...predicate.RecTeam) *RecTeamQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit adds a limit step to the query.
func (rtq *RecTeamQuery) Limit(limit int) *RecTeamQuery {
	rtq.limit = &limit
	return rtq
}

// Offset adds an offset step to the query.
func (rtq *RecTeamQuery) Offset(offset int) *RecTeamQuery {
	rtq.offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *RecTeamQuery) Unique(unique bool) *RecTeamQuery {
	rtq.unique = &unique
	return rtq
}

// Order adds an order step to the query.
func (rtq *RecTeamQuery) Order(o ...OrderFunc) *RecTeamQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryRecMemberEdges chains the current query on the "rec_member_edges" edge.
func (rtq *RecTeamQuery) QueryRecMemberEdges() *UserQuery {
	query := &UserQuery{config: rtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recteam.Table, recteam.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, recteam.RecMemberEdgesTable, recteam.RecMemberEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRecTeamJobEdges chains the current query on the "rec_team_job_edges" edge.
func (rtq *RecTeamQuery) QueryRecTeamJobEdges() *HiringJobQuery {
	query := &HiringJobQuery{config: rtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recteam.Table, recteam.FieldID, selector),
			sqlgraph.To(hiringjob.Table, hiringjob.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, recteam.RecTeamJobEdgesTable, recteam.RecTeamJobEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRecLeaderEdge chains the current query on the "rec_leader_edge" edge.
func (rtq *RecTeamQuery) QueryRecLeaderEdge() *UserQuery {
	query := &UserQuery{config: rtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recteam.Table, recteam.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, recteam.RecLeaderEdgeTable, recteam.RecLeaderEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecTeam entity from the query.
// Returns a *NotFoundError when no RecTeam was found.
func (rtq *RecTeamQuery) First(ctx context.Context) (*RecTeam, error) {
	nodes, err := rtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recteam.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *RecTeamQuery) FirstX(ctx context.Context) *RecTeam {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecTeam ID from the query.
// Returns a *NotFoundError when no RecTeam ID was found.
func (rtq *RecTeamQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recteam.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *RecTeamQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecTeam entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecTeam entity is found.
// Returns a *NotFoundError when no RecTeam entities are found.
func (rtq *RecTeamQuery) Only(ctx context.Context) (*RecTeam, error) {
	nodes, err := rtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recteam.Label}
	default:
		return nil, &NotSingularError{recteam.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *RecTeamQuery) OnlyX(ctx context.Context) *RecTeam {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecTeam ID in the query.
// Returns a *NotSingularError when more than one RecTeam ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *RecTeamQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recteam.Label}
	default:
		err = &NotSingularError{recteam.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *RecTeamQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecTeams.
func (rtq *RecTeamQuery) All(ctx context.Context) ([]*RecTeam, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rtq *RecTeamQuery) AllX(ctx context.Context) []*RecTeam {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecTeam IDs.
func (rtq *RecTeamQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rtq.Select(recteam.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *RecTeamQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *RecTeamQuery) Count(ctx context.Context) (int, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *RecTeamQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *RecTeamQuery) Exist(ctx context.Context) (bool, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *RecTeamQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecTeamQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *RecTeamQuery) Clone() *RecTeamQuery {
	if rtq == nil {
		return nil
	}
	return &RecTeamQuery{
		config:              rtq.config,
		limit:               rtq.limit,
		offset:              rtq.offset,
		order:               append([]OrderFunc{}, rtq.order...),
		predicates:          append([]predicate.RecTeam{}, rtq.predicates...),
		withRecMemberEdges:  rtq.withRecMemberEdges.Clone(),
		withRecTeamJobEdges: rtq.withRecTeamJobEdges.Clone(),
		withRecLeaderEdge:   rtq.withRecLeaderEdge.Clone(),
		// clone intermediate query.
		sql:    rtq.sql.Clone(),
		path:   rtq.path,
		unique: rtq.unique,
	}
}

// WithRecMemberEdges tells the query-builder to eager-load the nodes that are connected to
// the "rec_member_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *RecTeamQuery) WithRecMemberEdges(opts ...func(*UserQuery)) *RecTeamQuery {
	query := &UserQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	rtq.withRecMemberEdges = query
	return rtq
}

// WithRecTeamJobEdges tells the query-builder to eager-load the nodes that are connected to
// the "rec_team_job_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *RecTeamQuery) WithRecTeamJobEdges(opts ...func(*HiringJobQuery)) *RecTeamQuery {
	query := &HiringJobQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	rtq.withRecTeamJobEdges = query
	return rtq
}

// WithRecLeaderEdge tells the query-builder to eager-load the nodes that are connected to
// the "rec_leader_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *RecTeamQuery) WithRecLeaderEdge(opts ...func(*UserQuery)) *RecTeamQuery {
	query := &UserQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	rtq.withRecLeaderEdge = query
	return rtq
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
//	client.RecTeam.Query().
//		GroupBy(recteam.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rtq *RecTeamQuery) GroupBy(field string, fields ...string) *RecTeamGroupBy {
	grbuild := &RecTeamGroupBy{config: rtq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rtq.sqlQuery(ctx), nil
	}
	grbuild.label = recteam.Label
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
//	client.RecTeam.Query().
//		Select(recteam.FieldCreatedAt).
//		Scan(ctx, &v)
func (rtq *RecTeamQuery) Select(fields ...string) *RecTeamSelect {
	rtq.fields = append(rtq.fields, fields...)
	selbuild := &RecTeamSelect{RecTeamQuery: rtq}
	selbuild.label = recteam.Label
	selbuild.flds, selbuild.scan = &rtq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a RecTeamSelect configured with the given aggregations.
func (rtq *RecTeamQuery) Aggregate(fns ...AggregateFunc) *RecTeamSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *RecTeamQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rtq.fields {
		if !recteam.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *RecTeamQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecTeam, error) {
	var (
		nodes       = []*RecTeam{}
		_spec       = rtq.querySpec()
		loadedTypes = [3]bool{
			rtq.withRecMemberEdges != nil,
			rtq.withRecTeamJobEdges != nil,
			rtq.withRecLeaderEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecTeam).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecTeam{config: rtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rtq.modifiers) > 0 {
		_spec.Modifiers = rtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rtq.withRecMemberEdges; query != nil {
		if err := rtq.loadRecMemberEdges(ctx, query, nodes,
			func(n *RecTeam) { n.Edges.RecMemberEdges = []*User{} },
			func(n *RecTeam, e *User) { n.Edges.RecMemberEdges = append(n.Edges.RecMemberEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := rtq.withRecTeamJobEdges; query != nil {
		if err := rtq.loadRecTeamJobEdges(ctx, query, nodes,
			func(n *RecTeam) { n.Edges.RecTeamJobEdges = []*HiringJob{} },
			func(n *RecTeam, e *HiringJob) { n.Edges.RecTeamJobEdges = append(n.Edges.RecTeamJobEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := rtq.withRecLeaderEdge; query != nil {
		if err := rtq.loadRecLeaderEdge(ctx, query, nodes, nil,
			func(n *RecTeam, e *User) { n.Edges.RecLeaderEdge = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range rtq.withNamedRecMemberEdges {
		if err := rtq.loadRecMemberEdges(ctx, query, nodes,
			func(n *RecTeam) { n.appendNamedRecMemberEdges(name) },
			func(n *RecTeam, e *User) { n.appendNamedRecMemberEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rtq.withNamedRecTeamJobEdges {
		if err := rtq.loadRecTeamJobEdges(ctx, query, nodes,
			func(n *RecTeam) { n.appendNamedRecTeamJobEdges(name) },
			func(n *RecTeam, e *HiringJob) { n.appendNamedRecTeamJobEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rtq.loadTotal {
		if err := rtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rtq *RecTeamQuery) loadRecMemberEdges(ctx context.Context, query *UserQuery, nodes []*RecTeam, init func(*RecTeam), assign func(*RecTeam, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*RecTeam)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(recteam.RecMemberEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RecTeamID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rec_team_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rtq *RecTeamQuery) loadRecTeamJobEdges(ctx context.Context, query *HiringJobQuery, nodes []*RecTeam, init func(*RecTeam), assign func(*RecTeam, *HiringJob)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*RecTeam)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.HiringJob(func(s *sql.Selector) {
		s.Where(sql.InValues(recteam.RecTeamJobEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RecTeamID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rec_team_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rtq *RecTeamQuery) loadRecLeaderEdge(ctx context.Context, query *UserQuery, nodes []*RecTeam, init func(*RecTeam), assign func(*RecTeam, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RecTeam)
	for i := range nodes {
		fk := nodes[i].LeaderID
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
			return fmt.Errorf(`unexpected foreign-key "leader_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rtq *RecTeamQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	if len(rtq.modifiers) > 0 {
		_spec.Modifiers = rtq.modifiers
	}
	_spec.Node.Columns = rtq.fields
	if len(rtq.fields) > 0 {
		_spec.Unique = rtq.unique != nil && *rtq.unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *RecTeamQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rtq *RecTeamQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recteam.Table,
			Columns: recteam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recteam.FieldID,
			},
		},
		From:   rtq.sql,
		Unique: true,
	}
	if unique := rtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recteam.FieldID)
		for i := range fields {
			if fields[i] != recteam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *RecTeamQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(recteam.Table)
	columns := rtq.fields
	if len(columns) == 0 {
		columns = recteam.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.unique != nil && *rtq.unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRecMemberEdges tells the query-builder to eager-load the nodes that are connected to the "rec_member_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rtq *RecTeamQuery) WithNamedRecMemberEdges(name string, opts ...func(*UserQuery)) *RecTeamQuery {
	query := &UserQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rtq.withNamedRecMemberEdges == nil {
		rtq.withNamedRecMemberEdges = make(map[string]*UserQuery)
	}
	rtq.withNamedRecMemberEdges[name] = query
	return rtq
}

// WithNamedRecTeamJobEdges tells the query-builder to eager-load the nodes that are connected to the "rec_team_job_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rtq *RecTeamQuery) WithNamedRecTeamJobEdges(name string, opts ...func(*HiringJobQuery)) *RecTeamQuery {
	query := &HiringJobQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rtq.withNamedRecTeamJobEdges == nil {
		rtq.withNamedRecTeamJobEdges = make(map[string]*HiringJobQuery)
	}
	rtq.withNamedRecTeamJobEdges[name] = query
	return rtq
}

// RecTeamGroupBy is the group-by builder for RecTeam entities.
type RecTeamGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *RecTeamGroupBy) Aggregate(fns ...AggregateFunc) *RecTeamGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rtgb *RecTeamGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rtgb.path(ctx)
	if err != nil {
		return err
	}
	rtgb.sql = query
	return rtgb.sqlScan(ctx, v)
}

func (rtgb *RecTeamGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rtgb.fields {
		if !recteam.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rtgb *RecTeamGroupBy) sqlQuery() *sql.Selector {
	selector := rtgb.sql.Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rtgb.fields)+len(rtgb.fns))
		for _, f := range rtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rtgb.fields...)...)
}

// RecTeamSelect is the builder for selecting fields of RecTeam entities.
type RecTeamSelect struct {
	*RecTeamQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *RecTeamSelect) Aggregate(fns ...AggregateFunc) *RecTeamSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *RecTeamSelect) Scan(ctx context.Context, v any) error {
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	rts.sql = rts.RecTeamQuery.sqlQuery(ctx)
	return rts.sqlScan(ctx, v)
}

func (rts *RecTeamSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(rts.sql))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rts.sql.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

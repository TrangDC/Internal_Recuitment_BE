// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/team"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringJobQuery is the builder for querying HiringJob entities.
type HiringJobQuery struct {
	config
	limit                      *int
	offset                     *int
	unique                     *bool
	order                      []OrderFunc
	fields                     []string
	predicates                 []predicate.HiringJob
	withOwnerEdge              *UserQuery
	withTeamEdge               *TeamQuery
	withCandidateJobEdges      *CandidateJobQuery
	modifiers                  []func(*sql.Selector)
	loadTotal                  []func(context.Context, []*HiringJob) error
	withNamedCandidateJobEdges map[string]*CandidateJobQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HiringJobQuery builder.
func (hjq *HiringJobQuery) Where(ps ...predicate.HiringJob) *HiringJobQuery {
	hjq.predicates = append(hjq.predicates, ps...)
	return hjq
}

// Limit adds a limit step to the query.
func (hjq *HiringJobQuery) Limit(limit int) *HiringJobQuery {
	hjq.limit = &limit
	return hjq
}

// Offset adds an offset step to the query.
func (hjq *HiringJobQuery) Offset(offset int) *HiringJobQuery {
	hjq.offset = &offset
	return hjq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hjq *HiringJobQuery) Unique(unique bool) *HiringJobQuery {
	hjq.unique = &unique
	return hjq
}

// Order adds an order step to the query.
func (hjq *HiringJobQuery) Order(o ...OrderFunc) *HiringJobQuery {
	hjq.order = append(hjq.order, o...)
	return hjq
}

// QueryOwnerEdge chains the current query on the "owner_edge" edge.
func (hjq *HiringJobQuery) QueryOwnerEdge() *UserQuery {
	query := &UserQuery{config: hjq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hiringjob.Table, hiringjob.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hiringjob.OwnerEdgeTable, hiringjob.OwnerEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(hjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeamEdge chains the current query on the "team_edge" edge.
func (hjq *HiringJobQuery) QueryTeamEdge() *TeamQuery {
	query := &TeamQuery{config: hjq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hiringjob.Table, hiringjob.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hiringjob.TeamEdgeTable, hiringjob.TeamEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(hjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateJobEdges chains the current query on the "candidate_job_edges" edge.
func (hjq *HiringJobQuery) QueryCandidateJobEdges() *CandidateJobQuery {
	query := &CandidateJobQuery{config: hjq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hiringjob.Table, hiringjob.FieldID, selector),
			sqlgraph.To(candidatejob.Table, candidatejob.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, hiringjob.CandidateJobEdgesTable, hiringjob.CandidateJobEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(hjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HiringJob entity from the query.
// Returns a *NotFoundError when no HiringJob was found.
func (hjq *HiringJobQuery) First(ctx context.Context) (*HiringJob, error) {
	nodes, err := hjq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hiringjob.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hjq *HiringJobQuery) FirstX(ctx context.Context) *HiringJob {
	node, err := hjq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HiringJob ID from the query.
// Returns a *NotFoundError when no HiringJob ID was found.
func (hjq *HiringJobQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hjq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hiringjob.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hjq *HiringJobQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := hjq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HiringJob entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HiringJob entity is found.
// Returns a *NotFoundError when no HiringJob entities are found.
func (hjq *HiringJobQuery) Only(ctx context.Context) (*HiringJob, error) {
	nodes, err := hjq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hiringjob.Label}
	default:
		return nil, &NotSingularError{hiringjob.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hjq *HiringJobQuery) OnlyX(ctx context.Context) *HiringJob {
	node, err := hjq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HiringJob ID in the query.
// Returns a *NotSingularError when more than one HiringJob ID is found.
// Returns a *NotFoundError when no entities are found.
func (hjq *HiringJobQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = hjq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hiringjob.Label}
	default:
		err = &NotSingularError{hiringjob.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hjq *HiringJobQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := hjq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HiringJobs.
func (hjq *HiringJobQuery) All(ctx context.Context) ([]*HiringJob, error) {
	if err := hjq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hjq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hjq *HiringJobQuery) AllX(ctx context.Context) []*HiringJob {
	nodes, err := hjq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HiringJob IDs.
func (hjq *HiringJobQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := hjq.Select(hiringjob.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hjq *HiringJobQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := hjq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hjq *HiringJobQuery) Count(ctx context.Context) (int, error) {
	if err := hjq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hjq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hjq *HiringJobQuery) CountX(ctx context.Context) int {
	count, err := hjq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hjq *HiringJobQuery) Exist(ctx context.Context) (bool, error) {
	if err := hjq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hjq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hjq *HiringJobQuery) ExistX(ctx context.Context) bool {
	exist, err := hjq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HiringJobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hjq *HiringJobQuery) Clone() *HiringJobQuery {
	if hjq == nil {
		return nil
	}
	return &HiringJobQuery{
		config:                hjq.config,
		limit:                 hjq.limit,
		offset:                hjq.offset,
		order:                 append([]OrderFunc{}, hjq.order...),
		predicates:            append([]predicate.HiringJob{}, hjq.predicates...),
		withOwnerEdge:         hjq.withOwnerEdge.Clone(),
		withTeamEdge:          hjq.withTeamEdge.Clone(),
		withCandidateJobEdges: hjq.withCandidateJobEdges.Clone(),
		// clone intermediate query.
		sql:    hjq.sql.Clone(),
		path:   hjq.path,
		unique: hjq.unique,
	}
}

// WithOwnerEdge tells the query-builder to eager-load the nodes that are connected to
// the "owner_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (hjq *HiringJobQuery) WithOwnerEdge(opts ...func(*UserQuery)) *HiringJobQuery {
	query := &UserQuery{config: hjq.config}
	for _, opt := range opts {
		opt(query)
	}
	hjq.withOwnerEdge = query
	return hjq
}

// WithTeamEdge tells the query-builder to eager-load the nodes that are connected to
// the "team_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (hjq *HiringJobQuery) WithTeamEdge(opts ...func(*TeamQuery)) *HiringJobQuery {
	query := &TeamQuery{config: hjq.config}
	for _, opt := range opts {
		opt(query)
	}
	hjq.withTeamEdge = query
	return hjq
}

// WithCandidateJobEdges tells the query-builder to eager-load the nodes that are connected to
// the "candidate_job_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (hjq *HiringJobQuery) WithCandidateJobEdges(opts ...func(*CandidateJobQuery)) *HiringJobQuery {
	query := &CandidateJobQuery{config: hjq.config}
	for _, opt := range opts {
		opt(query)
	}
	hjq.withCandidateJobEdges = query
	return hjq
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
//	client.HiringJob.Query().
//		GroupBy(hiringjob.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hjq *HiringJobQuery) GroupBy(field string, fields ...string) *HiringJobGroupBy {
	grbuild := &HiringJobGroupBy{config: hjq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hjq.sqlQuery(ctx), nil
	}
	grbuild.label = hiringjob.Label
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
//	client.HiringJob.Query().
//		Select(hiringjob.FieldCreatedAt).
//		Scan(ctx, &v)
func (hjq *HiringJobQuery) Select(fields ...string) *HiringJobSelect {
	hjq.fields = append(hjq.fields, fields...)
	selbuild := &HiringJobSelect{HiringJobQuery: hjq}
	selbuild.label = hiringjob.Label
	selbuild.flds, selbuild.scan = &hjq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a HiringJobSelect configured with the given aggregations.
func (hjq *HiringJobQuery) Aggregate(fns ...AggregateFunc) *HiringJobSelect {
	return hjq.Select().Aggregate(fns...)
}

func (hjq *HiringJobQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hjq.fields {
		if !hiringjob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hjq.path != nil {
		prev, err := hjq.path(ctx)
		if err != nil {
			return err
		}
		hjq.sql = prev
	}
	return nil
}

func (hjq *HiringJobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HiringJob, error) {
	var (
		nodes       = []*HiringJob{}
		_spec       = hjq.querySpec()
		loadedTypes = [3]bool{
			hjq.withOwnerEdge != nil,
			hjq.withTeamEdge != nil,
			hjq.withCandidateJobEdges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HiringJob).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HiringJob{config: hjq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hjq.modifiers) > 0 {
		_spec.Modifiers = hjq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hjq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hjq.withOwnerEdge; query != nil {
		if err := hjq.loadOwnerEdge(ctx, query, nodes, nil,
			func(n *HiringJob, e *User) { n.Edges.OwnerEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := hjq.withTeamEdge; query != nil {
		if err := hjq.loadTeamEdge(ctx, query, nodes, nil,
			func(n *HiringJob, e *Team) { n.Edges.TeamEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := hjq.withCandidateJobEdges; query != nil {
		if err := hjq.loadCandidateJobEdges(ctx, query, nodes,
			func(n *HiringJob) { n.Edges.CandidateJobEdges = []*CandidateJob{} },
			func(n *HiringJob, e *CandidateJob) { n.Edges.CandidateJobEdges = append(n.Edges.CandidateJobEdges, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hjq.withNamedCandidateJobEdges {
		if err := hjq.loadCandidateJobEdges(ctx, query, nodes,
			func(n *HiringJob) { n.appendNamedCandidateJobEdges(name) },
			func(n *HiringJob, e *CandidateJob) { n.appendNamedCandidateJobEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range hjq.loadTotal {
		if err := hjq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hjq *HiringJobQuery) loadOwnerEdge(ctx context.Context, query *UserQuery, nodes []*HiringJob, init func(*HiringJob), assign func(*HiringJob, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HiringJob)
	for i := range nodes {
		fk := nodes[i].CreatedBy
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
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hjq *HiringJobQuery) loadTeamEdge(ctx context.Context, query *TeamQuery, nodes []*HiringJob, init func(*HiringJob), assign func(*HiringJob, *Team)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HiringJob)
	for i := range nodes {
		fk := nodes[i].TeamID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hjq *HiringJobQuery) loadCandidateJobEdges(ctx context.Context, query *CandidateJobQuery, nodes []*HiringJob, init func(*HiringJob), assign func(*HiringJob, *CandidateJob)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*HiringJob)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.InValues(hiringjob.CandidateJobEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.HiringJobID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "hiring_job_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (hjq *HiringJobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hjq.querySpec()
	if len(hjq.modifiers) > 0 {
		_spec.Modifiers = hjq.modifiers
	}
	_spec.Node.Columns = hjq.fields
	if len(hjq.fields) > 0 {
		_spec.Unique = hjq.unique != nil && *hjq.unique
	}
	return sqlgraph.CountNodes(ctx, hjq.driver, _spec)
}

func (hjq *HiringJobQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := hjq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (hjq *HiringJobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hiringjob.Table,
			Columns: hiringjob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjob.FieldID,
			},
		},
		From:   hjq.sql,
		Unique: true,
	}
	if unique := hjq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hjq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hiringjob.FieldID)
		for i := range fields {
			if fields[i] != hiringjob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hjq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hjq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hjq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hjq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hjq *HiringJobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hjq.driver.Dialect())
	t1 := builder.Table(hiringjob.Table)
	columns := hjq.fields
	if len(columns) == 0 {
		columns = hiringjob.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hjq.sql != nil {
		selector = hjq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hjq.unique != nil && *hjq.unique {
		selector.Distinct()
	}
	for _, p := range hjq.predicates {
		p(selector)
	}
	for _, p := range hjq.order {
		p(selector)
	}
	if offset := hjq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hjq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCandidateJobEdges tells the query-builder to eager-load the nodes that are connected to the "candidate_job_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hjq *HiringJobQuery) WithNamedCandidateJobEdges(name string, opts ...func(*CandidateJobQuery)) *HiringJobQuery {
	query := &CandidateJobQuery{config: hjq.config}
	for _, opt := range opts {
		opt(query)
	}
	if hjq.withNamedCandidateJobEdges == nil {
		hjq.withNamedCandidateJobEdges = make(map[string]*CandidateJobQuery)
	}
	hjq.withNamedCandidateJobEdges[name] = query
	return hjq
}

// HiringJobGroupBy is the group-by builder for HiringJob entities.
type HiringJobGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hjgb *HiringJobGroupBy) Aggregate(fns ...AggregateFunc) *HiringJobGroupBy {
	hjgb.fns = append(hjgb.fns, fns...)
	return hjgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hjgb *HiringJobGroupBy) Scan(ctx context.Context, v any) error {
	query, err := hjgb.path(ctx)
	if err != nil {
		return err
	}
	hjgb.sql = query
	return hjgb.sqlScan(ctx, v)
}

func (hjgb *HiringJobGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range hjgb.fields {
		if !hiringjob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hjgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hjgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hjgb *HiringJobGroupBy) sqlQuery() *sql.Selector {
	selector := hjgb.sql.Select()
	aggregation := make([]string, 0, len(hjgb.fns))
	for _, fn := range hjgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(hjgb.fields)+len(hjgb.fns))
		for _, f := range hjgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(hjgb.fields...)...)
}

// HiringJobSelect is the builder for selecting fields of HiringJob entities.
type HiringJobSelect struct {
	*HiringJobQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hjs *HiringJobSelect) Aggregate(fns ...AggregateFunc) *HiringJobSelect {
	hjs.fns = append(hjs.fns, fns...)
	return hjs
}

// Scan applies the selector query and scans the result into the given value.
func (hjs *HiringJobSelect) Scan(ctx context.Context, v any) error {
	if err := hjs.prepareQuery(ctx); err != nil {
		return err
	}
	hjs.sql = hjs.HiringJobQuery.sqlQuery(ctx)
	return hjs.sqlScan(ctx, v)
}

func (hjs *HiringJobSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(hjs.fns))
	for _, fn := range hjs.fns {
		aggregation = append(aggregation, fn(hjs.sql))
	}
	switch n := len(*hjs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		hjs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		hjs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := hjs.sql.Query()
	if err := hjs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
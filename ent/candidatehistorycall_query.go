// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidatehistorycall"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateHistoryCallQuery is the builder for querying CandidateHistoryCall entities.
type CandidateHistoryCallQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.CandidateHistoryCall
	withAttachmentEdges      *AttachmentQuery
	withCandidateEdge        *CandidateQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*CandidateHistoryCall) error
	withNamedAttachmentEdges map[string]*AttachmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CandidateHistoryCallQuery builder.
func (chcq *CandidateHistoryCallQuery) Where(ps ...predicate.CandidateHistoryCall) *CandidateHistoryCallQuery {
	chcq.predicates = append(chcq.predicates, ps...)
	return chcq
}

// Limit adds a limit step to the query.
func (chcq *CandidateHistoryCallQuery) Limit(limit int) *CandidateHistoryCallQuery {
	chcq.limit = &limit
	return chcq
}

// Offset adds an offset step to the query.
func (chcq *CandidateHistoryCallQuery) Offset(offset int) *CandidateHistoryCallQuery {
	chcq.offset = &offset
	return chcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chcq *CandidateHistoryCallQuery) Unique(unique bool) *CandidateHistoryCallQuery {
	chcq.unique = &unique
	return chcq
}

// Order adds an order step to the query.
func (chcq *CandidateHistoryCallQuery) Order(o ...OrderFunc) *CandidateHistoryCallQuery {
	chcq.order = append(chcq.order, o...)
	return chcq
}

// QueryAttachmentEdges chains the current query on the "attachment_edges" edge.
func (chcq *CandidateHistoryCallQuery) QueryAttachmentEdges() *AttachmentQuery {
	query := &AttachmentQuery{config: chcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := chcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := chcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatehistorycall.Table, candidatehistorycall.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, candidatehistorycall.AttachmentEdgesTable, candidatehistorycall.AttachmentEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(chcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateEdge chains the current query on the "candidate_edge" edge.
func (chcq *CandidateHistoryCallQuery) QueryCandidateEdge() *CandidateQuery {
	query := &CandidateQuery{config: chcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := chcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := chcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatehistorycall.Table, candidatehistorycall.FieldID, selector),
			sqlgraph.To(candidate.Table, candidate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidatehistorycall.CandidateEdgeTable, candidatehistorycall.CandidateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(chcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CandidateHistoryCall entity from the query.
// Returns a *NotFoundError when no CandidateHistoryCall was found.
func (chcq *CandidateHistoryCallQuery) First(ctx context.Context) (*CandidateHistoryCall, error) {
	nodes, err := chcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{candidatehistorycall.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) FirstX(ctx context.Context) *CandidateHistoryCall {
	node, err := chcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CandidateHistoryCall ID from the query.
// Returns a *NotFoundError when no CandidateHistoryCall ID was found.
func (chcq *CandidateHistoryCallQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{candidatehistorycall.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := chcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CandidateHistoryCall entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CandidateHistoryCall entity is found.
// Returns a *NotFoundError when no CandidateHistoryCall entities are found.
func (chcq *CandidateHistoryCallQuery) Only(ctx context.Context) (*CandidateHistoryCall, error) {
	nodes, err := chcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{candidatehistorycall.Label}
	default:
		return nil, &NotSingularError{candidatehistorycall.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) OnlyX(ctx context.Context) *CandidateHistoryCall {
	node, err := chcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CandidateHistoryCall ID in the query.
// Returns a *NotSingularError when more than one CandidateHistoryCall ID is found.
// Returns a *NotFoundError when no entities are found.
func (chcq *CandidateHistoryCallQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = chcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{candidatehistorycall.Label}
	default:
		err = &NotSingularError{candidatehistorycall.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := chcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CandidateHistoryCalls.
func (chcq *CandidateHistoryCallQuery) All(ctx context.Context) ([]*CandidateHistoryCall, error) {
	if err := chcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return chcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) AllX(ctx context.Context) []*CandidateHistoryCall {
	nodes, err := chcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CandidateHistoryCall IDs.
func (chcq *CandidateHistoryCallQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := chcq.Select(candidatehistorycall.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := chcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chcq *CandidateHistoryCallQuery) Count(ctx context.Context) (int, error) {
	if err := chcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return chcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) CountX(ctx context.Context) int {
	count, err := chcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chcq *CandidateHistoryCallQuery) Exist(ctx context.Context) (bool, error) {
	if err := chcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return chcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (chcq *CandidateHistoryCallQuery) ExistX(ctx context.Context) bool {
	exist, err := chcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CandidateHistoryCallQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chcq *CandidateHistoryCallQuery) Clone() *CandidateHistoryCallQuery {
	if chcq == nil {
		return nil
	}
	return &CandidateHistoryCallQuery{
		config:              chcq.config,
		limit:               chcq.limit,
		offset:              chcq.offset,
		order:               append([]OrderFunc{}, chcq.order...),
		predicates:          append([]predicate.CandidateHistoryCall{}, chcq.predicates...),
		withAttachmentEdges: chcq.withAttachmentEdges.Clone(),
		withCandidateEdge:   chcq.withCandidateEdge.Clone(),
		// clone intermediate query.
		sql:    chcq.sql.Clone(),
		path:   chcq.path,
		unique: chcq.unique,
	}
}

// WithAttachmentEdges tells the query-builder to eager-load the nodes that are connected to
// the "attachment_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (chcq *CandidateHistoryCallQuery) WithAttachmentEdges(opts ...func(*AttachmentQuery)) *CandidateHistoryCallQuery {
	query := &AttachmentQuery{config: chcq.config}
	for _, opt := range opts {
		opt(query)
	}
	chcq.withAttachmentEdges = query
	return chcq
}

// WithCandidateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (chcq *CandidateHistoryCallQuery) WithCandidateEdge(opts ...func(*CandidateQuery)) *CandidateHistoryCallQuery {
	query := &CandidateQuery{config: chcq.config}
	for _, opt := range opts {
		opt(query)
	}
	chcq.withCandidateEdge = query
	return chcq
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
//	client.CandidateHistoryCall.Query().
//		GroupBy(candidatehistorycall.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (chcq *CandidateHistoryCallQuery) GroupBy(field string, fields ...string) *CandidateHistoryCallGroupBy {
	grbuild := &CandidateHistoryCallGroupBy{config: chcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := chcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return chcq.sqlQuery(ctx), nil
	}
	grbuild.label = candidatehistorycall.Label
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
//	client.CandidateHistoryCall.Query().
//		Select(candidatehistorycall.FieldCreatedAt).
//		Scan(ctx, &v)
func (chcq *CandidateHistoryCallQuery) Select(fields ...string) *CandidateHistoryCallSelect {
	chcq.fields = append(chcq.fields, fields...)
	selbuild := &CandidateHistoryCallSelect{CandidateHistoryCallQuery: chcq}
	selbuild.label = candidatehistorycall.Label
	selbuild.flds, selbuild.scan = &chcq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CandidateHistoryCallSelect configured with the given aggregations.
func (chcq *CandidateHistoryCallQuery) Aggregate(fns ...AggregateFunc) *CandidateHistoryCallSelect {
	return chcq.Select().Aggregate(fns...)
}

func (chcq *CandidateHistoryCallQuery) prepareQuery(ctx context.Context) error {
	for _, f := range chcq.fields {
		if !candidatehistorycall.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chcq.path != nil {
		prev, err := chcq.path(ctx)
		if err != nil {
			return err
		}
		chcq.sql = prev
	}
	return nil
}

func (chcq *CandidateHistoryCallQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CandidateHistoryCall, error) {
	var (
		nodes       = []*CandidateHistoryCall{}
		_spec       = chcq.querySpec()
		loadedTypes = [2]bool{
			chcq.withAttachmentEdges != nil,
			chcq.withCandidateEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CandidateHistoryCall).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CandidateHistoryCall{config: chcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(chcq.modifiers) > 0 {
		_spec.Modifiers = chcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := chcq.withAttachmentEdges; query != nil {
		if err := chcq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateHistoryCall) { n.Edges.AttachmentEdges = []*Attachment{} },
			func(n *CandidateHistoryCall, e *Attachment) {
				n.Edges.AttachmentEdges = append(n.Edges.AttachmentEdges, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := chcq.withCandidateEdge; query != nil {
		if err := chcq.loadCandidateEdge(ctx, query, nodes, nil,
			func(n *CandidateHistoryCall, e *Candidate) { n.Edges.CandidateEdge = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range chcq.withNamedAttachmentEdges {
		if err := chcq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateHistoryCall) { n.appendNamedAttachmentEdges(name) },
			func(n *CandidateHistoryCall, e *Attachment) { n.appendNamedAttachmentEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range chcq.loadTotal {
		if err := chcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (chcq *CandidateHistoryCallQuery) loadAttachmentEdges(ctx context.Context, query *AttachmentQuery, nodes []*CandidateHistoryCall, init func(*CandidateHistoryCall), assign func(*CandidateHistoryCall, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CandidateHistoryCall)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(candidatehistorycall.AttachmentEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RelationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (chcq *CandidateHistoryCallQuery) loadCandidateEdge(ctx context.Context, query *CandidateQuery, nodes []*CandidateHistoryCall, init func(*CandidateHistoryCall), assign func(*CandidateHistoryCall, *Candidate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateHistoryCall)
	for i := range nodes {
		fk := nodes[i].CandidateID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "candidate_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (chcq *CandidateHistoryCallQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chcq.querySpec()
	if len(chcq.modifiers) > 0 {
		_spec.Modifiers = chcq.modifiers
	}
	_spec.Node.Columns = chcq.fields
	if len(chcq.fields) > 0 {
		_spec.Unique = chcq.unique != nil && *chcq.unique
	}
	return sqlgraph.CountNodes(ctx, chcq.driver, _spec)
}

func (chcq *CandidateHistoryCallQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := chcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (chcq *CandidateHistoryCallQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatehistorycall.Table,
			Columns: candidatehistorycall.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatehistorycall.FieldID,
			},
		},
		From:   chcq.sql,
		Unique: true,
	}
	if unique := chcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := chcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidatehistorycall.FieldID)
		for i := range fields {
			if fields[i] != candidatehistorycall.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chcq *CandidateHistoryCallQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chcq.driver.Dialect())
	t1 := builder.Table(candidatehistorycall.Table)
	columns := chcq.fields
	if len(columns) == 0 {
		columns = candidatehistorycall.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chcq.sql != nil {
		selector = chcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chcq.unique != nil && *chcq.unique {
		selector.Distinct()
	}
	for _, p := range chcq.predicates {
		p(selector)
	}
	for _, p := range chcq.order {
		p(selector)
	}
	if offset := chcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAttachmentEdges tells the query-builder to eager-load the nodes that are connected to the "attachment_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (chcq *CandidateHistoryCallQuery) WithNamedAttachmentEdges(name string, opts ...func(*AttachmentQuery)) *CandidateHistoryCallQuery {
	query := &AttachmentQuery{config: chcq.config}
	for _, opt := range opts {
		opt(query)
	}
	if chcq.withNamedAttachmentEdges == nil {
		chcq.withNamedAttachmentEdges = make(map[string]*AttachmentQuery)
	}
	chcq.withNamedAttachmentEdges[name] = query
	return chcq
}

// CandidateHistoryCallGroupBy is the group-by builder for CandidateHistoryCall entities.
type CandidateHistoryCallGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chcgb *CandidateHistoryCallGroupBy) Aggregate(fns ...AggregateFunc) *CandidateHistoryCallGroupBy {
	chcgb.fns = append(chcgb.fns, fns...)
	return chcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (chcgb *CandidateHistoryCallGroupBy) Scan(ctx context.Context, v any) error {
	query, err := chcgb.path(ctx)
	if err != nil {
		return err
	}
	chcgb.sql = query
	return chcgb.sqlScan(ctx, v)
}

func (chcgb *CandidateHistoryCallGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range chcgb.fields {
		if !candidatehistorycall.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := chcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (chcgb *CandidateHistoryCallGroupBy) sqlQuery() *sql.Selector {
	selector := chcgb.sql.Select()
	aggregation := make([]string, 0, len(chcgb.fns))
	for _, fn := range chcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(chcgb.fields)+len(chcgb.fns))
		for _, f := range chcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(chcgb.fields...)...)
}

// CandidateHistoryCallSelect is the builder for selecting fields of CandidateHistoryCall entities.
type CandidateHistoryCallSelect struct {
	*CandidateHistoryCallQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chcs *CandidateHistoryCallSelect) Aggregate(fns ...AggregateFunc) *CandidateHistoryCallSelect {
	chcs.fns = append(chcs.fns, fns...)
	return chcs
}

// Scan applies the selector query and scans the result into the given value.
func (chcs *CandidateHistoryCallSelect) Scan(ctx context.Context, v any) error {
	if err := chcs.prepareQuery(ctx); err != nil {
		return err
	}
	chcs.sql = chcs.CandidateHistoryCallQuery.sqlQuery(ctx)
	return chcs.sqlScan(ctx, v)
}

func (chcs *CandidateHistoryCallSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(chcs.fns))
	for _, fn := range chcs.fns {
		aggregation = append(aggregation, fn(chcs.sql))
	}
	switch n := len(*chcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		chcs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		chcs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := chcs.sql.Query()
	if err := chcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

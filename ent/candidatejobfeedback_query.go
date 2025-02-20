// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/attachment"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateJobFeedbackQuery is the builder for querying CandidateJobFeedback entities.
type CandidateJobFeedbackQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.CandidateJobFeedback
	withCreatedByEdge        *UserQuery
	withCandidateJobEdge     *CandidateJobQuery
	withAttachmentEdges      *AttachmentQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*CandidateJobFeedback) error
	withNamedAttachmentEdges map[string]*AttachmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CandidateJobFeedbackQuery builder.
func (cjfq *CandidateJobFeedbackQuery) Where(ps ...predicate.CandidateJobFeedback) *CandidateJobFeedbackQuery {
	cjfq.predicates = append(cjfq.predicates, ps...)
	return cjfq
}

// Limit adds a limit step to the query.
func (cjfq *CandidateJobFeedbackQuery) Limit(limit int) *CandidateJobFeedbackQuery {
	cjfq.limit = &limit
	return cjfq
}

// Offset adds an offset step to the query.
func (cjfq *CandidateJobFeedbackQuery) Offset(offset int) *CandidateJobFeedbackQuery {
	cjfq.offset = &offset
	return cjfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cjfq *CandidateJobFeedbackQuery) Unique(unique bool) *CandidateJobFeedbackQuery {
	cjfq.unique = &unique
	return cjfq
}

// Order adds an order step to the query.
func (cjfq *CandidateJobFeedbackQuery) Order(o ...OrderFunc) *CandidateJobFeedbackQuery {
	cjfq.order = append(cjfq.order, o...)
	return cjfq
}

// QueryCreatedByEdge chains the current query on the "created_by_edge" edge.
func (cjfq *CandidateJobFeedbackQuery) QueryCreatedByEdge() *UserQuery {
	query := &UserQuery{config: cjfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cjfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cjfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatejobfeedback.Table, candidatejobfeedback.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidatejobfeedback.CreatedByEdgeTable, candidatejobfeedback.CreatedByEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cjfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateJobEdge chains the current query on the "candidate_job_edge" edge.
func (cjfq *CandidateJobFeedbackQuery) QueryCandidateJobEdge() *CandidateJobQuery {
	query := &CandidateJobQuery{config: cjfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cjfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cjfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatejobfeedback.Table, candidatejobfeedback.FieldID, selector),
			sqlgraph.To(candidatejob.Table, candidatejob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidatejobfeedback.CandidateJobEdgeTable, candidatejobfeedback.CandidateJobEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cjfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachmentEdges chains the current query on the "attachment_edges" edge.
func (cjfq *CandidateJobFeedbackQuery) QueryAttachmentEdges() *AttachmentQuery {
	query := &AttachmentQuery{config: cjfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cjfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cjfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatejobfeedback.Table, candidatejobfeedback.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, candidatejobfeedback.AttachmentEdgesTable, candidatejobfeedback.AttachmentEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cjfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CandidateJobFeedback entity from the query.
// Returns a *NotFoundError when no CandidateJobFeedback was found.
func (cjfq *CandidateJobFeedbackQuery) First(ctx context.Context) (*CandidateJobFeedback, error) {
	nodes, err := cjfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{candidatejobfeedback.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) FirstX(ctx context.Context) *CandidateJobFeedback {
	node, err := cjfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CandidateJobFeedback ID from the query.
// Returns a *NotFoundError when no CandidateJobFeedback ID was found.
func (cjfq *CandidateJobFeedbackQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cjfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{candidatejobfeedback.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cjfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CandidateJobFeedback entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CandidateJobFeedback entity is found.
// Returns a *NotFoundError when no CandidateJobFeedback entities are found.
func (cjfq *CandidateJobFeedbackQuery) Only(ctx context.Context) (*CandidateJobFeedback, error) {
	nodes, err := cjfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{candidatejobfeedback.Label}
	default:
		return nil, &NotSingularError{candidatejobfeedback.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) OnlyX(ctx context.Context) *CandidateJobFeedback {
	node, err := cjfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CandidateJobFeedback ID in the query.
// Returns a *NotSingularError when more than one CandidateJobFeedback ID is found.
// Returns a *NotFoundError when no entities are found.
func (cjfq *CandidateJobFeedbackQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cjfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{candidatejobfeedback.Label}
	default:
		err = &NotSingularError{candidatejobfeedback.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cjfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CandidateJobFeedbacks.
func (cjfq *CandidateJobFeedbackQuery) All(ctx context.Context) ([]*CandidateJobFeedback, error) {
	if err := cjfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cjfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) AllX(ctx context.Context) []*CandidateJobFeedback {
	nodes, err := cjfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CandidateJobFeedback IDs.
func (cjfq *CandidateJobFeedbackQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cjfq.Select(candidatejobfeedback.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cjfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cjfq *CandidateJobFeedbackQuery) Count(ctx context.Context) (int, error) {
	if err := cjfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cjfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) CountX(ctx context.Context) int {
	count, err := cjfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cjfq *CandidateJobFeedbackQuery) Exist(ctx context.Context) (bool, error) {
	if err := cjfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cjfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cjfq *CandidateJobFeedbackQuery) ExistX(ctx context.Context) bool {
	exist, err := cjfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CandidateJobFeedbackQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cjfq *CandidateJobFeedbackQuery) Clone() *CandidateJobFeedbackQuery {
	if cjfq == nil {
		return nil
	}
	return &CandidateJobFeedbackQuery{
		config:               cjfq.config,
		limit:                cjfq.limit,
		offset:               cjfq.offset,
		order:                append([]OrderFunc{}, cjfq.order...),
		predicates:           append([]predicate.CandidateJobFeedback{}, cjfq.predicates...),
		withCreatedByEdge:    cjfq.withCreatedByEdge.Clone(),
		withCandidateJobEdge: cjfq.withCandidateJobEdge.Clone(),
		withAttachmentEdges:  cjfq.withAttachmentEdges.Clone(),
		// clone intermediate query.
		sql:    cjfq.sql.Clone(),
		path:   cjfq.path,
		unique: cjfq.unique,
	}
}

// WithCreatedByEdge tells the query-builder to eager-load the nodes that are connected to
// the "created_by_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (cjfq *CandidateJobFeedbackQuery) WithCreatedByEdge(opts ...func(*UserQuery)) *CandidateJobFeedbackQuery {
	query := &UserQuery{config: cjfq.config}
	for _, opt := range opts {
		opt(query)
	}
	cjfq.withCreatedByEdge = query
	return cjfq
}

// WithCandidateJobEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_job_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (cjfq *CandidateJobFeedbackQuery) WithCandidateJobEdge(opts ...func(*CandidateJobQuery)) *CandidateJobFeedbackQuery {
	query := &CandidateJobQuery{config: cjfq.config}
	for _, opt := range opts {
		opt(query)
	}
	cjfq.withCandidateJobEdge = query
	return cjfq
}

// WithAttachmentEdges tells the query-builder to eager-load the nodes that are connected to
// the "attachment_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (cjfq *CandidateJobFeedbackQuery) WithAttachmentEdges(opts ...func(*AttachmentQuery)) *CandidateJobFeedbackQuery {
	query := &AttachmentQuery{config: cjfq.config}
	for _, opt := range opts {
		opt(query)
	}
	cjfq.withAttachmentEdges = query
	return cjfq
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
//	client.CandidateJobFeedback.Query().
//		GroupBy(candidatejobfeedback.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cjfq *CandidateJobFeedbackQuery) GroupBy(field string, fields ...string) *CandidateJobFeedbackGroupBy {
	grbuild := &CandidateJobFeedbackGroupBy{config: cjfq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cjfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cjfq.sqlQuery(ctx), nil
	}
	grbuild.label = candidatejobfeedback.Label
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
//	client.CandidateJobFeedback.Query().
//		Select(candidatejobfeedback.FieldCreatedAt).
//		Scan(ctx, &v)
func (cjfq *CandidateJobFeedbackQuery) Select(fields ...string) *CandidateJobFeedbackSelect {
	cjfq.fields = append(cjfq.fields, fields...)
	selbuild := &CandidateJobFeedbackSelect{CandidateJobFeedbackQuery: cjfq}
	selbuild.label = candidatejobfeedback.Label
	selbuild.flds, selbuild.scan = &cjfq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CandidateJobFeedbackSelect configured with the given aggregations.
func (cjfq *CandidateJobFeedbackQuery) Aggregate(fns ...AggregateFunc) *CandidateJobFeedbackSelect {
	return cjfq.Select().Aggregate(fns...)
}

func (cjfq *CandidateJobFeedbackQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cjfq.fields {
		if !candidatejobfeedback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cjfq.path != nil {
		prev, err := cjfq.path(ctx)
		if err != nil {
			return err
		}
		cjfq.sql = prev
	}
	return nil
}

func (cjfq *CandidateJobFeedbackQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CandidateJobFeedback, error) {
	var (
		nodes       = []*CandidateJobFeedback{}
		_spec       = cjfq.querySpec()
		loadedTypes = [3]bool{
			cjfq.withCreatedByEdge != nil,
			cjfq.withCandidateJobEdge != nil,
			cjfq.withAttachmentEdges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CandidateJobFeedback).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CandidateJobFeedback{config: cjfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cjfq.modifiers) > 0 {
		_spec.Modifiers = cjfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cjfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cjfq.withCreatedByEdge; query != nil {
		if err := cjfq.loadCreatedByEdge(ctx, query, nodes, nil,
			func(n *CandidateJobFeedback, e *User) { n.Edges.CreatedByEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := cjfq.withCandidateJobEdge; query != nil {
		if err := cjfq.loadCandidateJobEdge(ctx, query, nodes, nil,
			func(n *CandidateJobFeedback, e *CandidateJob) { n.Edges.CandidateJobEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := cjfq.withAttachmentEdges; query != nil {
		if err := cjfq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateJobFeedback) { n.Edges.AttachmentEdges = []*Attachment{} },
			func(n *CandidateJobFeedback, e *Attachment) {
				n.Edges.AttachmentEdges = append(n.Edges.AttachmentEdges, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range cjfq.withNamedAttachmentEdges {
		if err := cjfq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateJobFeedback) { n.appendNamedAttachmentEdges(name) },
			func(n *CandidateJobFeedback, e *Attachment) { n.appendNamedAttachmentEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cjfq.loadTotal {
		if err := cjfq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cjfq *CandidateJobFeedbackQuery) loadCreatedByEdge(ctx context.Context, query *UserQuery, nodes []*CandidateJobFeedback, init func(*CandidateJobFeedback), assign func(*CandidateJobFeedback, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateJobFeedback)
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
func (cjfq *CandidateJobFeedbackQuery) loadCandidateJobEdge(ctx context.Context, query *CandidateJobQuery, nodes []*CandidateJobFeedback, init func(*CandidateJobFeedback), assign func(*CandidateJobFeedback, *CandidateJob)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateJobFeedback)
	for i := range nodes {
		fk := nodes[i].CandidateJobID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidatejob.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "candidate_job_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cjfq *CandidateJobFeedbackQuery) loadAttachmentEdges(ctx context.Context, query *AttachmentQuery, nodes []*CandidateJobFeedback, init func(*CandidateJobFeedback), assign func(*CandidateJobFeedback, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CandidateJobFeedback)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(candidatejobfeedback.AttachmentEdgesColumn, fks...))
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

func (cjfq *CandidateJobFeedbackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cjfq.querySpec()
	if len(cjfq.modifiers) > 0 {
		_spec.Modifiers = cjfq.modifiers
	}
	_spec.Node.Columns = cjfq.fields
	if len(cjfq.fields) > 0 {
		_spec.Unique = cjfq.unique != nil && *cjfq.unique
	}
	return sqlgraph.CountNodes(ctx, cjfq.driver, _spec)
}

func (cjfq *CandidateJobFeedbackQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cjfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cjfq *CandidateJobFeedbackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatejobfeedback.Table,
			Columns: candidatejobfeedback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejobfeedback.FieldID,
			},
		},
		From:   cjfq.sql,
		Unique: true,
	}
	if unique := cjfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cjfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidatejobfeedback.FieldID)
		for i := range fields {
			if fields[i] != candidatejobfeedback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cjfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cjfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cjfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cjfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cjfq *CandidateJobFeedbackQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cjfq.driver.Dialect())
	t1 := builder.Table(candidatejobfeedback.Table)
	columns := cjfq.fields
	if len(columns) == 0 {
		columns = candidatejobfeedback.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cjfq.sql != nil {
		selector = cjfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cjfq.unique != nil && *cjfq.unique {
		selector.Distinct()
	}
	for _, p := range cjfq.predicates {
		p(selector)
	}
	for _, p := range cjfq.order {
		p(selector)
	}
	if offset := cjfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cjfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAttachmentEdges tells the query-builder to eager-load the nodes that are connected to the "attachment_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cjfq *CandidateJobFeedbackQuery) WithNamedAttachmentEdges(name string, opts ...func(*AttachmentQuery)) *CandidateJobFeedbackQuery {
	query := &AttachmentQuery{config: cjfq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cjfq.withNamedAttachmentEdges == nil {
		cjfq.withNamedAttachmentEdges = make(map[string]*AttachmentQuery)
	}
	cjfq.withNamedAttachmentEdges[name] = query
	return cjfq
}

// CandidateJobFeedbackGroupBy is the group-by builder for CandidateJobFeedback entities.
type CandidateJobFeedbackGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cjfgb *CandidateJobFeedbackGroupBy) Aggregate(fns ...AggregateFunc) *CandidateJobFeedbackGroupBy {
	cjfgb.fns = append(cjfgb.fns, fns...)
	return cjfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cjfgb *CandidateJobFeedbackGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cjfgb.path(ctx)
	if err != nil {
		return err
	}
	cjfgb.sql = query
	return cjfgb.sqlScan(ctx, v)
}

func (cjfgb *CandidateJobFeedbackGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cjfgb.fields {
		if !candidatejobfeedback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cjfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cjfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cjfgb *CandidateJobFeedbackGroupBy) sqlQuery() *sql.Selector {
	selector := cjfgb.sql.Select()
	aggregation := make([]string, 0, len(cjfgb.fns))
	for _, fn := range cjfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cjfgb.fields)+len(cjfgb.fns))
		for _, f := range cjfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cjfgb.fields...)...)
}

// CandidateJobFeedbackSelect is the builder for selecting fields of CandidateJobFeedback entities.
type CandidateJobFeedbackSelect struct {
	*CandidateJobFeedbackQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cjfs *CandidateJobFeedbackSelect) Aggregate(fns ...AggregateFunc) *CandidateJobFeedbackSelect {
	cjfs.fns = append(cjfs.fns, fns...)
	return cjfs
}

// Scan applies the selector query and scans the result into the given value.
func (cjfs *CandidateJobFeedbackSelect) Scan(ctx context.Context, v any) error {
	if err := cjfs.prepareQuery(ctx); err != nil {
		return err
	}
	cjfs.sql = cjfs.CandidateJobFeedbackQuery.sqlQuery(ctx)
	return cjfs.sqlScan(ctx, v)
}

func (cjfs *CandidateJobFeedbackSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cjfs.fns))
	for _, fn := range cjfs.fns {
		aggregation = append(aggregation, fn(cjfs.sql))
	}
	switch n := len(*cjfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cjfs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cjfs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cjfs.sql.Query()
	if err := cjfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

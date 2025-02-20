// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidatenote"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateNoteQuery is the builder for querying CandidateNote entities.
type CandidateNoteQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.CandidateNote
	withCandidateEdge        *CandidateQuery
	withCreatedByEdge        *UserQuery
	withAttachmentEdges      *AttachmentQuery
	modifiers                []func(*sql.Selector)
	loadTotal                []func(context.Context, []*CandidateNote) error
	withNamedAttachmentEdges map[string]*AttachmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CandidateNoteQuery builder.
func (cnq *CandidateNoteQuery) Where(ps ...predicate.CandidateNote) *CandidateNoteQuery {
	cnq.predicates = append(cnq.predicates, ps...)
	return cnq
}

// Limit adds a limit step to the query.
func (cnq *CandidateNoteQuery) Limit(limit int) *CandidateNoteQuery {
	cnq.limit = &limit
	return cnq
}

// Offset adds an offset step to the query.
func (cnq *CandidateNoteQuery) Offset(offset int) *CandidateNoteQuery {
	cnq.offset = &offset
	return cnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cnq *CandidateNoteQuery) Unique(unique bool) *CandidateNoteQuery {
	cnq.unique = &unique
	return cnq
}

// Order adds an order step to the query.
func (cnq *CandidateNoteQuery) Order(o ...OrderFunc) *CandidateNoteQuery {
	cnq.order = append(cnq.order, o...)
	return cnq
}

// QueryCandidateEdge chains the current query on the "candidate_edge" edge.
func (cnq *CandidateNoteQuery) QueryCandidateEdge() *CandidateQuery {
	query := &CandidateQuery{config: cnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatenote.Table, candidatenote.FieldID, selector),
			sqlgraph.To(candidate.Table, candidate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidatenote.CandidateEdgeTable, candidatenote.CandidateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedByEdge chains the current query on the "created_by_edge" edge.
func (cnq *CandidateNoteQuery) QueryCreatedByEdge() *UserQuery {
	query := &UserQuery{config: cnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatenote.Table, candidatenote.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidatenote.CreatedByEdgeTable, candidatenote.CreatedByEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachmentEdges chains the current query on the "attachment_edges" edge.
func (cnq *CandidateNoteQuery) QueryAttachmentEdges() *AttachmentQuery {
	query := &AttachmentQuery{config: cnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidatenote.Table, candidatenote.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, candidatenote.AttachmentEdgesTable, candidatenote.AttachmentEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CandidateNote entity from the query.
// Returns a *NotFoundError when no CandidateNote was found.
func (cnq *CandidateNoteQuery) First(ctx context.Context) (*CandidateNote, error) {
	nodes, err := cnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{candidatenote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cnq *CandidateNoteQuery) FirstX(ctx context.Context) *CandidateNote {
	node, err := cnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CandidateNote ID from the query.
// Returns a *NotFoundError when no CandidateNote ID was found.
func (cnq *CandidateNoteQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{candidatenote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cnq *CandidateNoteQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CandidateNote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CandidateNote entity is found.
// Returns a *NotFoundError when no CandidateNote entities are found.
func (cnq *CandidateNoteQuery) Only(ctx context.Context) (*CandidateNote, error) {
	nodes, err := cnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{candidatenote.Label}
	default:
		return nil, &NotSingularError{candidatenote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cnq *CandidateNoteQuery) OnlyX(ctx context.Context) *CandidateNote {
	node, err := cnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CandidateNote ID in the query.
// Returns a *NotSingularError when more than one CandidateNote ID is found.
// Returns a *NotFoundError when no entities are found.
func (cnq *CandidateNoteQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{candidatenote.Label}
	default:
		err = &NotSingularError{candidatenote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cnq *CandidateNoteQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CandidateNotes.
func (cnq *CandidateNoteQuery) All(ctx context.Context) ([]*CandidateNote, error) {
	if err := cnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cnq *CandidateNoteQuery) AllX(ctx context.Context) []*CandidateNote {
	nodes, err := cnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CandidateNote IDs.
func (cnq *CandidateNoteQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cnq.Select(candidatenote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cnq *CandidateNoteQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cnq *CandidateNoteQuery) Count(ctx context.Context) (int, error) {
	if err := cnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cnq *CandidateNoteQuery) CountX(ctx context.Context) int {
	count, err := cnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cnq *CandidateNoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := cnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cnq *CandidateNoteQuery) ExistX(ctx context.Context) bool {
	exist, err := cnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CandidateNoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cnq *CandidateNoteQuery) Clone() *CandidateNoteQuery {
	if cnq == nil {
		return nil
	}
	return &CandidateNoteQuery{
		config:              cnq.config,
		limit:               cnq.limit,
		offset:              cnq.offset,
		order:               append([]OrderFunc{}, cnq.order...),
		predicates:          append([]predicate.CandidateNote{}, cnq.predicates...),
		withCandidateEdge:   cnq.withCandidateEdge.Clone(),
		withCreatedByEdge:   cnq.withCreatedByEdge.Clone(),
		withAttachmentEdges: cnq.withAttachmentEdges.Clone(),
		// clone intermediate query.
		sql:    cnq.sql.Clone(),
		path:   cnq.path,
		unique: cnq.unique,
	}
}

// WithCandidateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (cnq *CandidateNoteQuery) WithCandidateEdge(opts ...func(*CandidateQuery)) *CandidateNoteQuery {
	query := &CandidateQuery{config: cnq.config}
	for _, opt := range opts {
		opt(query)
	}
	cnq.withCandidateEdge = query
	return cnq
}

// WithCreatedByEdge tells the query-builder to eager-load the nodes that are connected to
// the "created_by_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (cnq *CandidateNoteQuery) WithCreatedByEdge(opts ...func(*UserQuery)) *CandidateNoteQuery {
	query := &UserQuery{config: cnq.config}
	for _, opt := range opts {
		opt(query)
	}
	cnq.withCreatedByEdge = query
	return cnq
}

// WithAttachmentEdges tells the query-builder to eager-load the nodes that are connected to
// the "attachment_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (cnq *CandidateNoteQuery) WithAttachmentEdges(opts ...func(*AttachmentQuery)) *CandidateNoteQuery {
	query := &AttachmentQuery{config: cnq.config}
	for _, opt := range opts {
		opt(query)
	}
	cnq.withAttachmentEdges = query
	return cnq
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
//	client.CandidateNote.Query().
//		GroupBy(candidatenote.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cnq *CandidateNoteQuery) GroupBy(field string, fields ...string) *CandidateNoteGroupBy {
	grbuild := &CandidateNoteGroupBy{config: cnq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cnq.sqlQuery(ctx), nil
	}
	grbuild.label = candidatenote.Label
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
//	client.CandidateNote.Query().
//		Select(candidatenote.FieldCreatedAt).
//		Scan(ctx, &v)
func (cnq *CandidateNoteQuery) Select(fields ...string) *CandidateNoteSelect {
	cnq.fields = append(cnq.fields, fields...)
	selbuild := &CandidateNoteSelect{CandidateNoteQuery: cnq}
	selbuild.label = candidatenote.Label
	selbuild.flds, selbuild.scan = &cnq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CandidateNoteSelect configured with the given aggregations.
func (cnq *CandidateNoteQuery) Aggregate(fns ...AggregateFunc) *CandidateNoteSelect {
	return cnq.Select().Aggregate(fns...)
}

func (cnq *CandidateNoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cnq.fields {
		if !candidatenote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cnq.path != nil {
		prev, err := cnq.path(ctx)
		if err != nil {
			return err
		}
		cnq.sql = prev
	}
	return nil
}

func (cnq *CandidateNoteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CandidateNote, error) {
	var (
		nodes       = []*CandidateNote{}
		_spec       = cnq.querySpec()
		loadedTypes = [3]bool{
			cnq.withCandidateEdge != nil,
			cnq.withCreatedByEdge != nil,
			cnq.withAttachmentEdges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CandidateNote).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CandidateNote{config: cnq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cnq.modifiers) > 0 {
		_spec.Modifiers = cnq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cnq.withCandidateEdge; query != nil {
		if err := cnq.loadCandidateEdge(ctx, query, nodes, nil,
			func(n *CandidateNote, e *Candidate) { n.Edges.CandidateEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := cnq.withCreatedByEdge; query != nil {
		if err := cnq.loadCreatedByEdge(ctx, query, nodes, nil,
			func(n *CandidateNote, e *User) { n.Edges.CreatedByEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := cnq.withAttachmentEdges; query != nil {
		if err := cnq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateNote) { n.Edges.AttachmentEdges = []*Attachment{} },
			func(n *CandidateNote, e *Attachment) { n.Edges.AttachmentEdges = append(n.Edges.AttachmentEdges, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cnq.withNamedAttachmentEdges {
		if err := cnq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateNote) { n.appendNamedAttachmentEdges(name) },
			func(n *CandidateNote, e *Attachment) { n.appendNamedAttachmentEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cnq.loadTotal {
		if err := cnq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cnq *CandidateNoteQuery) loadCandidateEdge(ctx context.Context, query *CandidateQuery, nodes []*CandidateNote, init func(*CandidateNote), assign func(*CandidateNote, *Candidate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateNote)
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
func (cnq *CandidateNoteQuery) loadCreatedByEdge(ctx context.Context, query *UserQuery, nodes []*CandidateNote, init func(*CandidateNote), assign func(*CandidateNote, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateNote)
	for i := range nodes {
		fk := nodes[i].CreatedByID
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
			return fmt.Errorf(`unexpected foreign-key "created_by_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cnq *CandidateNoteQuery) loadAttachmentEdges(ctx context.Context, query *AttachmentQuery, nodes []*CandidateNote, init func(*CandidateNote), assign func(*CandidateNote, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CandidateNote)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(candidatenote.AttachmentEdgesColumn, fks...))
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

func (cnq *CandidateNoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cnq.querySpec()
	if len(cnq.modifiers) > 0 {
		_spec.Modifiers = cnq.modifiers
	}
	_spec.Node.Columns = cnq.fields
	if len(cnq.fields) > 0 {
		_spec.Unique = cnq.unique != nil && *cnq.unique
	}
	return sqlgraph.CountNodes(ctx, cnq.driver, _spec)
}

func (cnq *CandidateNoteQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cnq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cnq *CandidateNoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatenote.Table,
			Columns: candidatenote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatenote.FieldID,
			},
		},
		From:   cnq.sql,
		Unique: true,
	}
	if unique := cnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidatenote.FieldID)
		for i := range fields {
			if fields[i] != candidatenote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cnq *CandidateNoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cnq.driver.Dialect())
	t1 := builder.Table(candidatenote.Table)
	columns := cnq.fields
	if len(columns) == 0 {
		columns = candidatenote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cnq.sql != nil {
		selector = cnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cnq.unique != nil && *cnq.unique {
		selector.Distinct()
	}
	for _, p := range cnq.predicates {
		p(selector)
	}
	for _, p := range cnq.order {
		p(selector)
	}
	if offset := cnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAttachmentEdges tells the query-builder to eager-load the nodes that are connected to the "attachment_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cnq *CandidateNoteQuery) WithNamedAttachmentEdges(name string, opts ...func(*AttachmentQuery)) *CandidateNoteQuery {
	query := &AttachmentQuery{config: cnq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cnq.withNamedAttachmentEdges == nil {
		cnq.withNamedAttachmentEdges = make(map[string]*AttachmentQuery)
	}
	cnq.withNamedAttachmentEdges[name] = query
	return cnq
}

// CandidateNoteGroupBy is the group-by builder for CandidateNote entities.
type CandidateNoteGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cngb *CandidateNoteGroupBy) Aggregate(fns ...AggregateFunc) *CandidateNoteGroupBy {
	cngb.fns = append(cngb.fns, fns...)
	return cngb
}

// Scan applies the group-by query and scans the result into the given value.
func (cngb *CandidateNoteGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cngb.path(ctx)
	if err != nil {
		return err
	}
	cngb.sql = query
	return cngb.sqlScan(ctx, v)
}

func (cngb *CandidateNoteGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cngb.fields {
		if !candidatenote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cngb *CandidateNoteGroupBy) sqlQuery() *sql.Selector {
	selector := cngb.sql.Select()
	aggregation := make([]string, 0, len(cngb.fns))
	for _, fn := range cngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cngb.fields)+len(cngb.fns))
		for _, f := range cngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cngb.fields...)...)
}

// CandidateNoteSelect is the builder for selecting fields of CandidateNote entities.
type CandidateNoteSelect struct {
	*CandidateNoteQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cns *CandidateNoteSelect) Aggregate(fns ...AggregateFunc) *CandidateNoteSelect {
	cns.fns = append(cns.fns, fns...)
	return cns
}

// Scan applies the selector query and scans the result into the given value.
func (cns *CandidateNoteSelect) Scan(ctx context.Context, v any) error {
	if err := cns.prepareQuery(ctx); err != nil {
		return err
	}
	cns.sql = cns.CandidateNoteQuery.sqlQuery(ctx)
	return cns.sqlScan(ctx, v)
}

func (cns *CandidateNoteSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cns.fns))
	for _, fn := range cns.fns {
		aggregation = append(aggregation, fn(cns.sql))
	}
	switch n := len(*cns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cns.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cns.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cns.sql.Query()
	if err := cns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/attachment"
	"trec/ent/candidateinterview"
	"trec/ent/candidateinterviewer"
	"trec/ent/candidatejob"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateInterviewQuery is the builder for querying CandidateInterview entities.
type CandidateInterviewQuery struct {
	config
	limit                     *int
	offset                    *int
	unique                    *bool
	order                     []OrderFunc
	fields                    []string
	predicates                []predicate.CandidateInterview
	withCandidateJobEdge      *CandidateJobQuery
	withAttachmentEdges       *AttachmentQuery
	withInterviewerEdges      *UserQuery
	withCreatedByEdge         *UserQuery
	withUserInterviewers      *CandidateInterviewerQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*CandidateInterview) error
	withNamedAttachmentEdges  map[string]*AttachmentQuery
	withNamedInterviewerEdges map[string]*UserQuery
	withNamedUserInterviewers map[string]*CandidateInterviewerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CandidateInterviewQuery builder.
func (ciq *CandidateInterviewQuery) Where(ps ...predicate.CandidateInterview) *CandidateInterviewQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CandidateInterviewQuery) Limit(limit int) *CandidateInterviewQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CandidateInterviewQuery) Offset(offset int) *CandidateInterviewQuery {
	ciq.offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CandidateInterviewQuery) Unique(unique bool) *CandidateInterviewQuery {
	ciq.unique = &unique
	return ciq
}

// Order adds an order step to the query.
func (ciq *CandidateInterviewQuery) Order(o ...OrderFunc) *CandidateInterviewQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryCandidateJobEdge chains the current query on the "candidate_job_edge" edge.
func (ciq *CandidateInterviewQuery) QueryCandidateJobEdge() *CandidateJobQuery {
	query := &CandidateJobQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateinterview.Table, candidateinterview.FieldID, selector),
			sqlgraph.To(candidatejob.Table, candidatejob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidateinterview.CandidateJobEdgeTable, candidateinterview.CandidateJobEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachmentEdges chains the current query on the "attachment_edges" edge.
func (ciq *CandidateInterviewQuery) QueryAttachmentEdges() *AttachmentQuery {
	query := &AttachmentQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateinterview.Table, candidateinterview.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, candidateinterview.AttachmentEdgesTable, candidateinterview.AttachmentEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInterviewerEdges chains the current query on the "interviewer_edges" edge.
func (ciq *CandidateInterviewQuery) QueryInterviewerEdges() *UserQuery {
	query := &UserQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateinterview.Table, candidateinterview.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, candidateinterview.InterviewerEdgesTable, candidateinterview.InterviewerEdgesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedByEdge chains the current query on the "created_by_edge" edge.
func (ciq *CandidateInterviewQuery) QueryCreatedByEdge() *UserQuery {
	query := &UserQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateinterview.Table, candidateinterview.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidateinterview.CreatedByEdgeTable, candidateinterview.CreatedByEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserInterviewers chains the current query on the "user_interviewers" edge.
func (ciq *CandidateInterviewQuery) QueryUserInterviewers() *CandidateInterviewerQuery {
	query := &CandidateInterviewerQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateinterview.Table, candidateinterview.FieldID, selector),
			sqlgraph.To(candidateinterviewer.Table, candidateinterviewer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, candidateinterview.UserInterviewersTable, candidateinterview.UserInterviewersColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CandidateInterview entity from the query.
// Returns a *NotFoundError when no CandidateInterview was found.
func (ciq *CandidateInterviewQuery) First(ctx context.Context) (*CandidateInterview, error) {
	nodes, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{candidateinterview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) FirstX(ctx context.Context) *CandidateInterview {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CandidateInterview ID from the query.
// Returns a *NotFoundError when no CandidateInterview ID was found.
func (ciq *CandidateInterviewQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{candidateinterview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CandidateInterview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CandidateInterview entity is found.
// Returns a *NotFoundError when no CandidateInterview entities are found.
func (ciq *CandidateInterviewQuery) Only(ctx context.Context) (*CandidateInterview, error) {
	nodes, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{candidateinterview.Label}
	default:
		return nil, &NotSingularError{candidateinterview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) OnlyX(ctx context.Context) *CandidateInterview {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CandidateInterview ID in the query.
// Returns a *NotSingularError when more than one CandidateInterview ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CandidateInterviewQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{candidateinterview.Label}
	default:
		err = &NotSingularError{candidateinterview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CandidateInterviews.
func (ciq *CandidateInterviewQuery) All(ctx context.Context) ([]*CandidateInterview, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) AllX(ctx context.Context) []*CandidateInterview {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CandidateInterview IDs.
func (ciq *CandidateInterviewQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ciq.Select(candidateinterview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CandidateInterviewQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CandidateInterviewQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CandidateInterviewQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CandidateInterviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CandidateInterviewQuery) Clone() *CandidateInterviewQuery {
	if ciq == nil {
		return nil
	}
	return &CandidateInterviewQuery{
		config:               ciq.config,
		limit:                ciq.limit,
		offset:               ciq.offset,
		order:                append([]OrderFunc{}, ciq.order...),
		predicates:           append([]predicate.CandidateInterview{}, ciq.predicates...),
		withCandidateJobEdge: ciq.withCandidateJobEdge.Clone(),
		withAttachmentEdges:  ciq.withAttachmentEdges.Clone(),
		withInterviewerEdges: ciq.withInterviewerEdges.Clone(),
		withCreatedByEdge:    ciq.withCreatedByEdge.Clone(),
		withUserInterviewers: ciq.withUserInterviewers.Clone(),
		// clone intermediate query.
		sql:    ciq.sql.Clone(),
		path:   ciq.path,
		unique: ciq.unique,
	}
}

// WithCandidateJobEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_job_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithCandidateJobEdge(opts ...func(*CandidateJobQuery)) *CandidateInterviewQuery {
	query := &CandidateJobQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCandidateJobEdge = query
	return ciq
}

// WithAttachmentEdges tells the query-builder to eager-load the nodes that are connected to
// the "attachment_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithAttachmentEdges(opts ...func(*AttachmentQuery)) *CandidateInterviewQuery {
	query := &AttachmentQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withAttachmentEdges = query
	return ciq
}

// WithInterviewerEdges tells the query-builder to eager-load the nodes that are connected to
// the "interviewer_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithInterviewerEdges(opts ...func(*UserQuery)) *CandidateInterviewQuery {
	query := &UserQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withInterviewerEdges = query
	return ciq
}

// WithCreatedByEdge tells the query-builder to eager-load the nodes that are connected to
// the "created_by_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithCreatedByEdge(opts ...func(*UserQuery)) *CandidateInterviewQuery {
	query := &UserQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCreatedByEdge = query
	return ciq
}

// WithUserInterviewers tells the query-builder to eager-load the nodes that are connected to
// the "user_interviewers" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithUserInterviewers(opts ...func(*CandidateInterviewerQuery)) *CandidateInterviewQuery {
	query := &CandidateInterviewerQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withUserInterviewers = query
	return ciq
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
//	client.CandidateInterview.Query().
//		GroupBy(candidateinterview.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ciq *CandidateInterviewQuery) GroupBy(field string, fields ...string) *CandidateInterviewGroupBy {
	grbuild := &CandidateInterviewGroupBy{config: ciq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(ctx), nil
	}
	grbuild.label = candidateinterview.Label
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
//	client.CandidateInterview.Query().
//		Select(candidateinterview.FieldCreatedAt).
//		Scan(ctx, &v)
func (ciq *CandidateInterviewQuery) Select(fields ...string) *CandidateInterviewSelect {
	ciq.fields = append(ciq.fields, fields...)
	selbuild := &CandidateInterviewSelect{CandidateInterviewQuery: ciq}
	selbuild.label = candidateinterview.Label
	selbuild.flds, selbuild.scan = &ciq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CandidateInterviewSelect configured with the given aggregations.
func (ciq *CandidateInterviewQuery) Aggregate(fns ...AggregateFunc) *CandidateInterviewSelect {
	return ciq.Select().Aggregate(fns...)
}

func (ciq *CandidateInterviewQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ciq.fields {
		if !candidateinterview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CandidateInterviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CandidateInterview, error) {
	var (
		nodes       = []*CandidateInterview{}
		_spec       = ciq.querySpec()
		loadedTypes = [5]bool{
			ciq.withCandidateJobEdge != nil,
			ciq.withAttachmentEdges != nil,
			ciq.withInterviewerEdges != nil,
			ciq.withCreatedByEdge != nil,
			ciq.withUserInterviewers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CandidateInterview).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CandidateInterview{config: ciq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ciq.withCandidateJobEdge; query != nil {
		if err := ciq.loadCandidateJobEdge(ctx, query, nodes, nil,
			func(n *CandidateInterview, e *CandidateJob) { n.Edges.CandidateJobEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := ciq.withAttachmentEdges; query != nil {
		if err := ciq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateInterview) { n.Edges.AttachmentEdges = []*Attachment{} },
			func(n *CandidateInterview, e *Attachment) {
				n.Edges.AttachmentEdges = append(n.Edges.AttachmentEdges, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := ciq.withInterviewerEdges; query != nil {
		if err := ciq.loadInterviewerEdges(ctx, query, nodes,
			func(n *CandidateInterview) { n.Edges.InterviewerEdges = []*User{} },
			func(n *CandidateInterview, e *User) { n.Edges.InterviewerEdges = append(n.Edges.InterviewerEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := ciq.withCreatedByEdge; query != nil {
		if err := ciq.loadCreatedByEdge(ctx, query, nodes, nil,
			func(n *CandidateInterview, e *User) { n.Edges.CreatedByEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := ciq.withUserInterviewers; query != nil {
		if err := ciq.loadUserInterviewers(ctx, query, nodes,
			func(n *CandidateInterview) { n.Edges.UserInterviewers = []*CandidateInterviewer{} },
			func(n *CandidateInterview, e *CandidateInterviewer) {
				n.Edges.UserInterviewers = append(n.Edges.UserInterviewers, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range ciq.withNamedAttachmentEdges {
		if err := ciq.loadAttachmentEdges(ctx, query, nodes,
			func(n *CandidateInterview) { n.appendNamedAttachmentEdges(name) },
			func(n *CandidateInterview, e *Attachment) { n.appendNamedAttachmentEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ciq.withNamedInterviewerEdges {
		if err := ciq.loadInterviewerEdges(ctx, query, nodes,
			func(n *CandidateInterview) { n.appendNamedInterviewerEdges(name) },
			func(n *CandidateInterview, e *User) { n.appendNamedInterviewerEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ciq.withNamedUserInterviewers {
		if err := ciq.loadUserInterviewers(ctx, query, nodes,
			func(n *CandidateInterview) { n.appendNamedUserInterviewers(name) },
			func(n *CandidateInterview, e *CandidateInterviewer) { n.appendNamedUserInterviewers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ciq.loadTotal {
		if err := ciq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ciq *CandidateInterviewQuery) loadCandidateJobEdge(ctx context.Context, query *CandidateJobQuery, nodes []*CandidateInterview, init func(*CandidateInterview), assign func(*CandidateInterview, *CandidateJob)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateInterview)
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
func (ciq *CandidateInterviewQuery) loadAttachmentEdges(ctx context.Context, query *AttachmentQuery, nodes []*CandidateInterview, init func(*CandidateInterview), assign func(*CandidateInterview, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CandidateInterview)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(candidateinterview.AttachmentEdgesColumn, fks...))
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
func (ciq *CandidateInterviewQuery) loadInterviewerEdges(ctx context.Context, query *UserQuery, nodes []*CandidateInterview, init func(*CandidateInterview), assign func(*CandidateInterview, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*CandidateInterview)
	nids := make(map[uuid.UUID]map[*CandidateInterview]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(candidateinterview.InterviewerEdgesTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(candidateinterview.InterviewerEdgesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(candidateinterview.InterviewerEdgesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(candidateinterview.InterviewerEdgesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*CandidateInterview]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "interviewer_edges" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ciq *CandidateInterviewQuery) loadCreatedByEdge(ctx context.Context, query *UserQuery, nodes []*CandidateInterview, init func(*CandidateInterview), assign func(*CandidateInterview, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateInterview)
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
func (ciq *CandidateInterviewQuery) loadUserInterviewers(ctx context.Context, query *CandidateInterviewerQuery, nodes []*CandidateInterview, init func(*CandidateInterview), assign func(*CandidateInterview, *CandidateInterviewer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CandidateInterview)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.CandidateInterviewer(func(s *sql.Selector) {
		s.Where(sql.InValues(candidateinterview.UserInterviewersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CandidateInterviewID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "candidate_interview_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ciq *CandidateInterviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	_spec.Node.Columns = ciq.fields
	if len(ciq.fields) > 0 {
		_spec.Unique = ciq.unique != nil && *ciq.unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CandidateInterviewQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ciq *CandidateInterviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidateinterview.Table,
			Columns: candidateinterview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateinterview.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if unique := ciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidateinterview.FieldID)
		for i := range fields {
			if fields[i] != candidateinterview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CandidateInterviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(candidateinterview.Table)
	columns := ciq.fields
	if len(columns) == 0 {
		columns = candidateinterview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.unique != nil && *ciq.unique {
		selector.Distinct()
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAttachmentEdges tells the query-builder to eager-load the nodes that are connected to the "attachment_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithNamedAttachmentEdges(name string, opts ...func(*AttachmentQuery)) *CandidateInterviewQuery {
	query := &AttachmentQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	if ciq.withNamedAttachmentEdges == nil {
		ciq.withNamedAttachmentEdges = make(map[string]*AttachmentQuery)
	}
	ciq.withNamedAttachmentEdges[name] = query
	return ciq
}

// WithNamedInterviewerEdges tells the query-builder to eager-load the nodes that are connected to the "interviewer_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithNamedInterviewerEdges(name string, opts ...func(*UserQuery)) *CandidateInterviewQuery {
	query := &UserQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	if ciq.withNamedInterviewerEdges == nil {
		ciq.withNamedInterviewerEdges = make(map[string]*UserQuery)
	}
	ciq.withNamedInterviewerEdges[name] = query
	return ciq
}

// WithNamedUserInterviewers tells the query-builder to eager-load the nodes that are connected to the "user_interviewers"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ciq *CandidateInterviewQuery) WithNamedUserInterviewers(name string, opts ...func(*CandidateInterviewerQuery)) *CandidateInterviewQuery {
	query := &CandidateInterviewerQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	if ciq.withNamedUserInterviewers == nil {
		ciq.withNamedUserInterviewers = make(map[string]*CandidateInterviewerQuery)
	}
	ciq.withNamedUserInterviewers[name] = query
	return ciq
}

// CandidateInterviewGroupBy is the group-by builder for CandidateInterview entities.
type CandidateInterviewGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CandidateInterviewGroupBy) Aggregate(fns ...AggregateFunc) *CandidateInterviewGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scans the result into the given value.
func (cigb *CandidateInterviewGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

func (cigb *CandidateInterviewGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cigb.fields {
		if !candidateinterview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CandidateInterviewGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql.Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
		for _, f := range cigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cigb.fields...)...)
}

// CandidateInterviewSelect is the builder for selecting fields of CandidateInterview entities.
type CandidateInterviewSelect struct {
	*CandidateInterviewQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cis *CandidateInterviewSelect) Aggregate(fns ...AggregateFunc) *CandidateInterviewSelect {
	cis.fns = append(cis.fns, fns...)
	return cis
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CandidateInterviewSelect) Scan(ctx context.Context, v any) error {
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	cis.sql = cis.CandidateInterviewQuery.sqlQuery(ctx)
	return cis.sqlScan(ctx, v)
}

func (cis *CandidateInterviewSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cis.fns))
	for _, fn := range cis.fns {
		aggregation = append(aggregation, fn(cis.sql))
	}
	switch n := len(*cis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cis.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cis.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cis.sql.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

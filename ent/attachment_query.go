// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateaward"
	"trec/ent/candidatecertificate"
	"trec/ent/candidateeducate"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatenote"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttachmentQuery is the builder for querying Attachment entities.
type AttachmentQuery struct {
	config
	limit                        *int
	offset                       *int
	unique                       *bool
	order                        []OrderFunc
	fields                       []string
	predicates                   []predicate.Attachment
	withCandidateJobEdge         *CandidateJobQuery
	withCandidateJobFeedbackEdge *CandidateJobFeedbackQuery
	withCandidateInterviewEdge   *CandidateInterviewQuery
	withCandidateEdge            *CandidateQuery
	withCandidateEducateEdge     *CandidateEducateQuery
	withCandidateAwardEdge       *CandidateAwardQuery
	withCandidateCertificateEdge *CandidateCertificateQuery
	withCandidateHistoryCallEdge *CandidateHistoryCallQuery
	withCandidateNoteEdge        *CandidateNoteQuery
	modifiers                    []func(*sql.Selector)
	loadTotal                    []func(context.Context, []*Attachment) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttachmentQuery builder.
func (aq *AttachmentQuery) Where(ps ...predicate.Attachment) *AttachmentQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AttachmentQuery) Limit(limit int) *AttachmentQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AttachmentQuery) Offset(offset int) *AttachmentQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttachmentQuery) Unique(unique bool) *AttachmentQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AttachmentQuery) Order(o ...OrderFunc) *AttachmentQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryCandidateJobEdge chains the current query on the "candidate_job_edge" edge.
func (aq *AttachmentQuery) QueryCandidateJobEdge() *CandidateJobQuery {
	query := &CandidateJobQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidatejob.Table, candidatejob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateJobEdgeTable, attachment.CandidateJobEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateJobFeedbackEdge chains the current query on the "candidate_job_feedback_edge" edge.
func (aq *AttachmentQuery) QueryCandidateJobFeedbackEdge() *CandidateJobFeedbackQuery {
	query := &CandidateJobFeedbackQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidatejobfeedback.Table, candidatejobfeedback.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateJobFeedbackEdgeTable, attachment.CandidateJobFeedbackEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateInterviewEdge chains the current query on the "candidate_interview_edge" edge.
func (aq *AttachmentQuery) QueryCandidateInterviewEdge() *CandidateInterviewQuery {
	query := &CandidateInterviewQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidateinterview.Table, candidateinterview.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateInterviewEdgeTable, attachment.CandidateInterviewEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateEdge chains the current query on the "candidate_edge" edge.
func (aq *AttachmentQuery) QueryCandidateEdge() *CandidateQuery {
	query := &CandidateQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidate.Table, candidate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateEdgeTable, attachment.CandidateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateEducateEdge chains the current query on the "candidate_educate_edge" edge.
func (aq *AttachmentQuery) QueryCandidateEducateEdge() *CandidateEducateQuery {
	query := &CandidateEducateQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidateeducate.Table, candidateeducate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateEducateEdgeTable, attachment.CandidateEducateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateAwardEdge chains the current query on the "candidate_award_edge" edge.
func (aq *AttachmentQuery) QueryCandidateAwardEdge() *CandidateAwardQuery {
	query := &CandidateAwardQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidateaward.Table, candidateaward.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateAwardEdgeTable, attachment.CandidateAwardEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateCertificateEdge chains the current query on the "candidate_certificate_edge" edge.
func (aq *AttachmentQuery) QueryCandidateCertificateEdge() *CandidateCertificateQuery {
	query := &CandidateCertificateQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidatecertificate.Table, candidatecertificate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateCertificateEdgeTable, attachment.CandidateCertificateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateHistoryCallEdge chains the current query on the "candidate_history_call_edge" edge.
func (aq *AttachmentQuery) QueryCandidateHistoryCallEdge() *CandidateHistoryCallQuery {
	query := &CandidateHistoryCallQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidatehistorycall.Table, candidatehistorycall.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateHistoryCallEdgeTable, attachment.CandidateHistoryCallEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateNoteEdge chains the current query on the "candidate_note_edge" edge.
func (aq *AttachmentQuery) QueryCandidateNoteEdge() *CandidateNoteQuery {
	query := &CandidateNoteQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(candidatenote.Table, candidatenote.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.CandidateNoteEdgeTable, attachment.CandidateNoteEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Attachment entity from the query.
// Returns a *NotFoundError when no Attachment was found.
func (aq *AttachmentQuery) First(ctx context.Context) (*Attachment, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attachment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttachmentQuery) FirstX(ctx context.Context) *Attachment {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attachment ID from the query.
// Returns a *NotFoundError when no Attachment ID was found.
func (aq *AttachmentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attachment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttachmentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attachment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attachment entity is found.
// Returns a *NotFoundError when no Attachment entities are found.
func (aq *AttachmentQuery) Only(ctx context.Context) (*Attachment, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attachment.Label}
	default:
		return nil, &NotSingularError{attachment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttachmentQuery) OnlyX(ctx context.Context) *Attachment {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attachment ID in the query.
// Returns a *NotSingularError when more than one Attachment ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttachmentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attachment.Label}
	default:
		err = &NotSingularError{attachment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttachmentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attachments.
func (aq *AttachmentQuery) All(ctx context.Context) ([]*Attachment, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttachmentQuery) AllX(ctx context.Context) []*Attachment {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attachment IDs.
func (aq *AttachmentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aq.Select(attachment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttachmentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttachmentQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttachmentQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttachmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttachmentQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttachmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttachmentQuery) Clone() *AttachmentQuery {
	if aq == nil {
		return nil
	}
	return &AttachmentQuery{
		config:                       aq.config,
		limit:                        aq.limit,
		offset:                       aq.offset,
		order:                        append([]OrderFunc{}, aq.order...),
		predicates:                   append([]predicate.Attachment{}, aq.predicates...),
		withCandidateJobEdge:         aq.withCandidateJobEdge.Clone(),
		withCandidateJobFeedbackEdge: aq.withCandidateJobFeedbackEdge.Clone(),
		withCandidateInterviewEdge:   aq.withCandidateInterviewEdge.Clone(),
		withCandidateEdge:            aq.withCandidateEdge.Clone(),
		withCandidateEducateEdge:     aq.withCandidateEducateEdge.Clone(),
		withCandidateAwardEdge:       aq.withCandidateAwardEdge.Clone(),
		withCandidateCertificateEdge: aq.withCandidateCertificateEdge.Clone(),
		withCandidateHistoryCallEdge: aq.withCandidateHistoryCallEdge.Clone(),
		withCandidateNoteEdge:        aq.withCandidateNoteEdge.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithCandidateJobEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_job_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateJobEdge(opts ...func(*CandidateJobQuery)) *AttachmentQuery {
	query := &CandidateJobQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateJobEdge = query
	return aq
}

// WithCandidateJobFeedbackEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_job_feedback_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateJobFeedbackEdge(opts ...func(*CandidateJobFeedbackQuery)) *AttachmentQuery {
	query := &CandidateJobFeedbackQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateJobFeedbackEdge = query
	return aq
}

// WithCandidateInterviewEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_interview_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateInterviewEdge(opts ...func(*CandidateInterviewQuery)) *AttachmentQuery {
	query := &CandidateInterviewQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateInterviewEdge = query
	return aq
}

// WithCandidateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateEdge(opts ...func(*CandidateQuery)) *AttachmentQuery {
	query := &CandidateQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateEdge = query
	return aq
}

// WithCandidateEducateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_educate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateEducateEdge(opts ...func(*CandidateEducateQuery)) *AttachmentQuery {
	query := &CandidateEducateQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateEducateEdge = query
	return aq
}

// WithCandidateAwardEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_award_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateAwardEdge(opts ...func(*CandidateAwardQuery)) *AttachmentQuery {
	query := &CandidateAwardQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateAwardEdge = query
	return aq
}

// WithCandidateCertificateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_certificate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateCertificateEdge(opts ...func(*CandidateCertificateQuery)) *AttachmentQuery {
	query := &CandidateCertificateQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateCertificateEdge = query
	return aq
}

// WithCandidateHistoryCallEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_history_call_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateHistoryCallEdge(opts ...func(*CandidateHistoryCallQuery)) *AttachmentQuery {
	query := &CandidateHistoryCallQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateHistoryCallEdge = query
	return aq
}

// WithCandidateNoteEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_note_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithCandidateNoteEdge(opts ...func(*CandidateNoteQuery)) *AttachmentQuery {
	query := &CandidateNoteQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCandidateNoteEdge = query
	return aq
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
//	client.Attachment.Query().
//		GroupBy(attachment.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttachmentQuery) GroupBy(field string, fields ...string) *AttachmentGroupBy {
	grbuild := &AttachmentGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = attachment.Label
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
//	client.Attachment.Query().
//		Select(attachment.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AttachmentQuery) Select(fields ...string) *AttachmentSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AttachmentSelect{AttachmentQuery: aq}
	selbuild.label = attachment.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AttachmentSelect configured with the given aggregations.
func (aq *AttachmentQuery) Aggregate(fns ...AggregateFunc) *AttachmentSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AttachmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !attachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttachmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attachment, error) {
	var (
		nodes       = []*Attachment{}
		_spec       = aq.querySpec()
		loadedTypes = [9]bool{
			aq.withCandidateJobEdge != nil,
			aq.withCandidateJobFeedbackEdge != nil,
			aq.withCandidateInterviewEdge != nil,
			aq.withCandidateEdge != nil,
			aq.withCandidateEducateEdge != nil,
			aq.withCandidateAwardEdge != nil,
			aq.withCandidateCertificateEdge != nil,
			aq.withCandidateHistoryCallEdge != nil,
			aq.withCandidateNoteEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Attachment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Attachment{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withCandidateJobEdge; query != nil {
		if err := aq.loadCandidateJobEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateJob) { n.Edges.CandidateJobEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateJobFeedbackEdge; query != nil {
		if err := aq.loadCandidateJobFeedbackEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateJobFeedback) { n.Edges.CandidateJobFeedbackEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateInterviewEdge; query != nil {
		if err := aq.loadCandidateInterviewEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateInterview) { n.Edges.CandidateInterviewEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateEdge; query != nil {
		if err := aq.loadCandidateEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *Candidate) { n.Edges.CandidateEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateEducateEdge; query != nil {
		if err := aq.loadCandidateEducateEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateEducate) { n.Edges.CandidateEducateEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateAwardEdge; query != nil {
		if err := aq.loadCandidateAwardEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateAward) { n.Edges.CandidateAwardEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateCertificateEdge; query != nil {
		if err := aq.loadCandidateCertificateEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateCertificate) { n.Edges.CandidateCertificateEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateHistoryCallEdge; query != nil {
		if err := aq.loadCandidateHistoryCallEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateHistoryCall) { n.Edges.CandidateHistoryCallEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCandidateNoteEdge; query != nil {
		if err := aq.loadCandidateNoteEdge(ctx, query, nodes, nil,
			func(n *Attachment, e *CandidateNote) { n.Edges.CandidateNoteEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AttachmentQuery) loadCandidateJobEdge(ctx context.Context, query *CandidateJobQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateJob)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
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
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateJobFeedbackEdge(ctx context.Context, query *CandidateJobFeedbackQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateJobFeedback)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidatejobfeedback.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateInterviewEdge(ctx context.Context, query *CandidateInterviewQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateInterview)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidateinterview.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateEdge(ctx context.Context, query *CandidateQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *Candidate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
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
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateEducateEdge(ctx context.Context, query *CandidateEducateQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateEducate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidateeducate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateAwardEdge(ctx context.Context, query *CandidateAwardQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateAward)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidateaward.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateCertificateEdge(ctx context.Context, query *CandidateCertificateQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateCertificate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidatecertificate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateHistoryCallEdge(ctx context.Context, query *CandidateHistoryCallQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateHistoryCall)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidatehistorycall.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadCandidateNoteEdge(ctx context.Context, query *CandidateNoteQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *CandidateNote)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		fk := nodes[i].RelationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(candidatenote.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relation_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AttachmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttachmentQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AttachmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for i := range fields {
			if fields[i] != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AttachmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attachment.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = attachment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttachmentGroupBy is the group-by builder for Attachment entities.
type AttachmentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttachmentGroupBy) Aggregate(fns ...AggregateFunc) *AttachmentGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AttachmentGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AttachmentGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !attachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AttachmentGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AttachmentSelect is the builder for selecting fields of Attachment entities.
type AttachmentSelect struct {
	*AttachmentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AttachmentSelect) Aggregate(fns ...AggregateFunc) *AttachmentSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttachmentSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AttachmentQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AttachmentSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(as.sql))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		as.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		as.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

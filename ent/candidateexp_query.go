// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"trec/ent/candidate"
	"trec/ent/candidateexp"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateExpQuery is the builder for querying CandidateExp entities.
type CandidateExpQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.CandidateExp
	withCandidateEdge *CandidateQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*CandidateExp) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CandidateExpQuery builder.
func (ceq *CandidateExpQuery) Where(ps ...predicate.CandidateExp) *CandidateExpQuery {
	ceq.predicates = append(ceq.predicates, ps...)
	return ceq
}

// Limit adds a limit step to the query.
func (ceq *CandidateExpQuery) Limit(limit int) *CandidateExpQuery {
	ceq.limit = &limit
	return ceq
}

// Offset adds an offset step to the query.
func (ceq *CandidateExpQuery) Offset(offset int) *CandidateExpQuery {
	ceq.offset = &offset
	return ceq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ceq *CandidateExpQuery) Unique(unique bool) *CandidateExpQuery {
	ceq.unique = &unique
	return ceq
}

// Order adds an order step to the query.
func (ceq *CandidateExpQuery) Order(o ...OrderFunc) *CandidateExpQuery {
	ceq.order = append(ceq.order, o...)
	return ceq
}

// QueryCandidateEdge chains the current query on the "candidate_edge" edge.
func (ceq *CandidateExpQuery) QueryCandidateEdge() *CandidateQuery {
	query := &CandidateQuery{config: ceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ceq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(candidateexp.Table, candidateexp.FieldID, selector),
			sqlgraph.To(candidate.Table, candidate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, candidateexp.CandidateEdgeTable, candidateexp.CandidateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CandidateExp entity from the query.
// Returns a *NotFoundError when no CandidateExp was found.
func (ceq *CandidateExpQuery) First(ctx context.Context) (*CandidateExp, error) {
	nodes, err := ceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{candidateexp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ceq *CandidateExpQuery) FirstX(ctx context.Context) *CandidateExp {
	node, err := ceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CandidateExp ID from the query.
// Returns a *NotFoundError when no CandidateExp ID was found.
func (ceq *CandidateExpQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{candidateexp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ceq *CandidateExpQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CandidateExp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CandidateExp entity is found.
// Returns a *NotFoundError when no CandidateExp entities are found.
func (ceq *CandidateExpQuery) Only(ctx context.Context) (*CandidateExp, error) {
	nodes, err := ceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{candidateexp.Label}
	default:
		return nil, &NotSingularError{candidateexp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ceq *CandidateExpQuery) OnlyX(ctx context.Context) *CandidateExp {
	node, err := ceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CandidateExp ID in the query.
// Returns a *NotSingularError when more than one CandidateExp ID is found.
// Returns a *NotFoundError when no entities are found.
func (ceq *CandidateExpQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{candidateexp.Label}
	default:
		err = &NotSingularError{candidateexp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ceq *CandidateExpQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CandidateExps.
func (ceq *CandidateExpQuery) All(ctx context.Context) ([]*CandidateExp, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ceq *CandidateExpQuery) AllX(ctx context.Context) []*CandidateExp {
	nodes, err := ceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CandidateExp IDs.
func (ceq *CandidateExpQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ceq.Select(candidateexp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ceq *CandidateExpQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ceq *CandidateExpQuery) Count(ctx context.Context) (int, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ceq *CandidateExpQuery) CountX(ctx context.Context) int {
	count, err := ceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ceq *CandidateExpQuery) Exist(ctx context.Context) (bool, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ceq *CandidateExpQuery) ExistX(ctx context.Context) bool {
	exist, err := ceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CandidateExpQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ceq *CandidateExpQuery) Clone() *CandidateExpQuery {
	if ceq == nil {
		return nil
	}
	return &CandidateExpQuery{
		config:            ceq.config,
		limit:             ceq.limit,
		offset:            ceq.offset,
		order:             append([]OrderFunc{}, ceq.order...),
		predicates:        append([]predicate.CandidateExp{}, ceq.predicates...),
		withCandidateEdge: ceq.withCandidateEdge.Clone(),
		// clone intermediate query.
		sql:    ceq.sql.Clone(),
		path:   ceq.path,
		unique: ceq.unique,
	}
}

// WithCandidateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (ceq *CandidateExpQuery) WithCandidateEdge(opts ...func(*CandidateQuery)) *CandidateExpQuery {
	query := &CandidateQuery{config: ceq.config}
	for _, opt := range opts {
		opt(query)
	}
	ceq.withCandidateEdge = query
	return ceq
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
//	client.CandidateExp.Query().
//		GroupBy(candidateexp.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ceq *CandidateExpQuery) GroupBy(field string, fields ...string) *CandidateExpGroupBy {
	grbuild := &CandidateExpGroupBy{config: ceq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(ctx), nil
	}
	grbuild.label = candidateexp.Label
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
//	client.CandidateExp.Query().
//		Select(candidateexp.FieldCreatedAt).
//		Scan(ctx, &v)
func (ceq *CandidateExpQuery) Select(fields ...string) *CandidateExpSelect {
	ceq.fields = append(ceq.fields, fields...)
	selbuild := &CandidateExpSelect{CandidateExpQuery: ceq}
	selbuild.label = candidateexp.Label
	selbuild.flds, selbuild.scan = &ceq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CandidateExpSelect configured with the given aggregations.
func (ceq *CandidateExpQuery) Aggregate(fns ...AggregateFunc) *CandidateExpSelect {
	return ceq.Select().Aggregate(fns...)
}

func (ceq *CandidateExpQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ceq.fields {
		if !candidateexp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ceq.path != nil {
		prev, err := ceq.path(ctx)
		if err != nil {
			return err
		}
		ceq.sql = prev
	}
	return nil
}

func (ceq *CandidateExpQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CandidateExp, error) {
	var (
		nodes       = []*CandidateExp{}
		_spec       = ceq.querySpec()
		loadedTypes = [1]bool{
			ceq.withCandidateEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CandidateExp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CandidateExp{config: ceq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ceq.modifiers) > 0 {
		_spec.Modifiers = ceq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ceq.withCandidateEdge; query != nil {
		if err := ceq.loadCandidateEdge(ctx, query, nodes, nil,
			func(n *CandidateExp, e *Candidate) { n.Edges.CandidateEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range ceq.loadTotal {
		if err := ceq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ceq *CandidateExpQuery) loadCandidateEdge(ctx context.Context, query *CandidateQuery, nodes []*CandidateExp, init func(*CandidateExp), assign func(*CandidateExp, *Candidate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CandidateExp)
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

func (ceq *CandidateExpQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ceq.querySpec()
	if len(ceq.modifiers) > 0 {
		_spec.Modifiers = ceq.modifiers
	}
	_spec.Node.Columns = ceq.fields
	if len(ceq.fields) > 0 {
		_spec.Unique = ceq.unique != nil && *ceq.unique
	}
	return sqlgraph.CountNodes(ctx, ceq.driver, _spec)
}

func (ceq *CandidateExpQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ceq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ceq *CandidateExpQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidateexp.Table,
			Columns: candidateexp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateexp.FieldID,
			},
		},
		From:   ceq.sql,
		Unique: true,
	}
	if unique := ceq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ceq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidateexp.FieldID)
		for i := range fields {
			if fields[i] != candidateexp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ceq *CandidateExpQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ceq.driver.Dialect())
	t1 := builder.Table(candidateexp.Table)
	columns := ceq.fields
	if len(columns) == 0 {
		columns = candidateexp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ceq.sql != nil {
		selector = ceq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ceq.unique != nil && *ceq.unique {
		selector.Distinct()
	}
	for _, p := range ceq.predicates {
		p(selector)
	}
	for _, p := range ceq.order {
		p(selector)
	}
	if offset := ceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CandidateExpGroupBy is the group-by builder for CandidateExp entities.
type CandidateExpGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cegb *CandidateExpGroupBy) Aggregate(fns ...AggregateFunc) *CandidateExpGroupBy {
	cegb.fns = append(cegb.fns, fns...)
	return cegb
}

// Scan applies the group-by query and scans the result into the given value.
func (cegb *CandidateExpGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cegb.path(ctx)
	if err != nil {
		return err
	}
	cegb.sql = query
	return cegb.sqlScan(ctx, v)
}

func (cegb *CandidateExpGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cegb.fields {
		if !candidateexp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cegb *CandidateExpGroupBy) sqlQuery() *sql.Selector {
	selector := cegb.sql.Select()
	aggregation := make([]string, 0, len(cegb.fns))
	for _, fn := range cegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cegb.fields)+len(cegb.fns))
		for _, f := range cegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cegb.fields...)...)
}

// CandidateExpSelect is the builder for selecting fields of CandidateExp entities.
type CandidateExpSelect struct {
	*CandidateExpQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ces *CandidateExpSelect) Aggregate(fns ...AggregateFunc) *CandidateExpSelect {
	ces.fns = append(ces.fns, fns...)
	return ces
}

// Scan applies the selector query and scans the result into the given value.
func (ces *CandidateExpSelect) Scan(ctx context.Context, v any) error {
	if err := ces.prepareQuery(ctx); err != nil {
		return err
	}
	ces.sql = ces.CandidateExpQuery.sqlQuery(ctx)
	return ces.sqlScan(ctx, v)
}

func (ces *CandidateExpSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ces.fns))
	for _, fn := range ces.fns {
		aggregation = append(aggregation, fn(ces.sql))
	}
	switch n := len(*ces.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ces.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ces.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ces.sql.Query()
	if err := ces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

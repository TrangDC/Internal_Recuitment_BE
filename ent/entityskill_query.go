// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"trec/ent/candidate"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/skill"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntitySkillQuery is the builder for querying EntitySkill entities.
type EntitySkillQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.EntitySkill
	withSkillEdge     *SkillQuery
	withHiringJobEdge *HiringJobQuery
	withCandidateEdge *CandidateQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*EntitySkill) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntitySkillQuery builder.
func (esq *EntitySkillQuery) Where(ps ...predicate.EntitySkill) *EntitySkillQuery {
	esq.predicates = append(esq.predicates, ps...)
	return esq
}

// Limit adds a limit step to the query.
func (esq *EntitySkillQuery) Limit(limit int) *EntitySkillQuery {
	esq.limit = &limit
	return esq
}

// Offset adds an offset step to the query.
func (esq *EntitySkillQuery) Offset(offset int) *EntitySkillQuery {
	esq.offset = &offset
	return esq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (esq *EntitySkillQuery) Unique(unique bool) *EntitySkillQuery {
	esq.unique = &unique
	return esq
}

// Order adds an order step to the query.
func (esq *EntitySkillQuery) Order(o ...OrderFunc) *EntitySkillQuery {
	esq.order = append(esq.order, o...)
	return esq
}

// QuerySkillEdge chains the current query on the "skill_edge" edge.
func (esq *EntitySkillQuery) QuerySkillEdge() *SkillQuery {
	query := &SkillQuery{config: esq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entityskill.Table, entityskill.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entityskill.SkillEdgeTable, entityskill.SkillEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(esq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHiringJobEdge chains the current query on the "hiring_job_edge" edge.
func (esq *EntitySkillQuery) QueryHiringJobEdge() *HiringJobQuery {
	query := &HiringJobQuery{config: esq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entityskill.Table, entityskill.FieldID, selector),
			sqlgraph.To(hiringjob.Table, hiringjob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entityskill.HiringJobEdgeTable, entityskill.HiringJobEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(esq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCandidateEdge chains the current query on the "candidate_edge" edge.
func (esq *EntitySkillQuery) QueryCandidateEdge() *CandidateQuery {
	query := &CandidateQuery{config: esq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entityskill.Table, entityskill.FieldID, selector),
			sqlgraph.To(candidate.Table, candidate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entityskill.CandidateEdgeTable, entityskill.CandidateEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(esq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntitySkill entity from the query.
// Returns a *NotFoundError when no EntitySkill was found.
func (esq *EntitySkillQuery) First(ctx context.Context) (*EntitySkill, error) {
	nodes, err := esq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entityskill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esq *EntitySkillQuery) FirstX(ctx context.Context) *EntitySkill {
	node, err := esq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntitySkill ID from the query.
// Returns a *NotFoundError when no EntitySkill ID was found.
func (esq *EntitySkillQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = esq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entityskill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esq *EntitySkillQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := esq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntitySkill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntitySkill entity is found.
// Returns a *NotFoundError when no EntitySkill entities are found.
func (esq *EntitySkillQuery) Only(ctx context.Context) (*EntitySkill, error) {
	nodes, err := esq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entityskill.Label}
	default:
		return nil, &NotSingularError{entityskill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esq *EntitySkillQuery) OnlyX(ctx context.Context) *EntitySkill {
	node, err := esq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntitySkill ID in the query.
// Returns a *NotSingularError when more than one EntitySkill ID is found.
// Returns a *NotFoundError when no entities are found.
func (esq *EntitySkillQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = esq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entityskill.Label}
	default:
		err = &NotSingularError{entityskill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esq *EntitySkillQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := esq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntitySkills.
func (esq *EntitySkillQuery) All(ctx context.Context) ([]*EntitySkill, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return esq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (esq *EntitySkillQuery) AllX(ctx context.Context) []*EntitySkill {
	nodes, err := esq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntitySkill IDs.
func (esq *EntitySkillQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := esq.Select(entityskill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esq *EntitySkillQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := esq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esq *EntitySkillQuery) Count(ctx context.Context) (int, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return esq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (esq *EntitySkillQuery) CountX(ctx context.Context) int {
	count, err := esq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esq *EntitySkillQuery) Exist(ctx context.Context) (bool, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return esq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (esq *EntitySkillQuery) ExistX(ctx context.Context) bool {
	exist, err := esq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntitySkillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esq *EntitySkillQuery) Clone() *EntitySkillQuery {
	if esq == nil {
		return nil
	}
	return &EntitySkillQuery{
		config:            esq.config,
		limit:             esq.limit,
		offset:            esq.offset,
		order:             append([]OrderFunc{}, esq.order...),
		predicates:        append([]predicate.EntitySkill{}, esq.predicates...),
		withSkillEdge:     esq.withSkillEdge.Clone(),
		withHiringJobEdge: esq.withHiringJobEdge.Clone(),
		withCandidateEdge: esq.withCandidateEdge.Clone(),
		// clone intermediate query.
		sql:    esq.sql.Clone(),
		path:   esq.path,
		unique: esq.unique,
	}
}

// WithSkillEdge tells the query-builder to eager-load the nodes that are connected to
// the "skill_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (esq *EntitySkillQuery) WithSkillEdge(opts ...func(*SkillQuery)) *EntitySkillQuery {
	query := &SkillQuery{config: esq.config}
	for _, opt := range opts {
		opt(query)
	}
	esq.withSkillEdge = query
	return esq
}

// WithHiringJobEdge tells the query-builder to eager-load the nodes that are connected to
// the "hiring_job_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (esq *EntitySkillQuery) WithHiringJobEdge(opts ...func(*HiringJobQuery)) *EntitySkillQuery {
	query := &HiringJobQuery{config: esq.config}
	for _, opt := range opts {
		opt(query)
	}
	esq.withHiringJobEdge = query
	return esq
}

// WithCandidateEdge tells the query-builder to eager-load the nodes that are connected to
// the "candidate_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (esq *EntitySkillQuery) WithCandidateEdge(opts ...func(*CandidateQuery)) *EntitySkillQuery {
	query := &CandidateQuery{config: esq.config}
	for _, opt := range opts {
		opt(query)
	}
	esq.withCandidateEdge = query
	return esq
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
//	client.EntitySkill.Query().
//		GroupBy(entityskill.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (esq *EntitySkillQuery) GroupBy(field string, fields ...string) *EntitySkillGroupBy {
	grbuild := &EntitySkillGroupBy{config: esq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return esq.sqlQuery(ctx), nil
	}
	grbuild.label = entityskill.Label
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
//	client.EntitySkill.Query().
//		Select(entityskill.FieldCreatedAt).
//		Scan(ctx, &v)
func (esq *EntitySkillQuery) Select(fields ...string) *EntitySkillSelect {
	esq.fields = append(esq.fields, fields...)
	selbuild := &EntitySkillSelect{EntitySkillQuery: esq}
	selbuild.label = entityskill.Label
	selbuild.flds, selbuild.scan = &esq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a EntitySkillSelect configured with the given aggregations.
func (esq *EntitySkillQuery) Aggregate(fns ...AggregateFunc) *EntitySkillSelect {
	return esq.Select().Aggregate(fns...)
}

func (esq *EntitySkillQuery) prepareQuery(ctx context.Context) error {
	for _, f := range esq.fields {
		if !entityskill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if esq.path != nil {
		prev, err := esq.path(ctx)
		if err != nil {
			return err
		}
		esq.sql = prev
	}
	return nil
}

func (esq *EntitySkillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntitySkill, error) {
	var (
		nodes       = []*EntitySkill{}
		_spec       = esq.querySpec()
		loadedTypes = [3]bool{
			esq.withSkillEdge != nil,
			esq.withHiringJobEdge != nil,
			esq.withCandidateEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntitySkill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntitySkill{config: esq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(esq.modifiers) > 0 {
		_spec.Modifiers = esq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, esq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := esq.withSkillEdge; query != nil {
		if err := esq.loadSkillEdge(ctx, query, nodes, nil,
			func(n *EntitySkill, e *Skill) { n.Edges.SkillEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := esq.withHiringJobEdge; query != nil {
		if err := esq.loadHiringJobEdge(ctx, query, nodes, nil,
			func(n *EntitySkill, e *HiringJob) { n.Edges.HiringJobEdge = e }); err != nil {
			return nil, err
		}
	}
	if query := esq.withCandidateEdge; query != nil {
		if err := esq.loadCandidateEdge(ctx, query, nodes, nil,
			func(n *EntitySkill, e *Candidate) { n.Edges.CandidateEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range esq.loadTotal {
		if err := esq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (esq *EntitySkillQuery) loadSkillEdge(ctx context.Context, query *SkillQuery, nodes []*EntitySkill, init func(*EntitySkill), assign func(*EntitySkill, *Skill)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntitySkill)
	for i := range nodes {
		fk := nodes[i].SkillID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(skill.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "skill_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (esq *EntitySkillQuery) loadHiringJobEdge(ctx context.Context, query *HiringJobQuery, nodes []*EntitySkill, init func(*EntitySkill), assign func(*EntitySkill, *HiringJob)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntitySkill)
	for i := range nodes {
		fk := nodes[i].EntityID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(hiringjob.IDIn(ids...))
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
func (esq *EntitySkillQuery) loadCandidateEdge(ctx context.Context, query *CandidateQuery, nodes []*EntitySkill, init func(*EntitySkill), assign func(*EntitySkill, *Candidate)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EntitySkill)
	for i := range nodes {
		fk := nodes[i].EntityID
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
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (esq *EntitySkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esq.querySpec()
	if len(esq.modifiers) > 0 {
		_spec.Modifiers = esq.modifiers
	}
	_spec.Node.Columns = esq.fields
	if len(esq.fields) > 0 {
		_spec.Unique = esq.unique != nil && *esq.unique
	}
	return sqlgraph.CountNodes(ctx, esq.driver, _spec)
}

func (esq *EntitySkillQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := esq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (esq *EntitySkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entityskill.Table,
			Columns: entityskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entityskill.FieldID,
			},
		},
		From:   esq.sql,
		Unique: true,
	}
	if unique := esq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := esq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entityskill.FieldID)
		for i := range fields {
			if fields[i] != entityskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := esq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (esq *EntitySkillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(esq.driver.Dialect())
	t1 := builder.Table(entityskill.Table)
	columns := esq.fields
	if len(columns) == 0 {
		columns = entityskill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if esq.sql != nil {
		selector = esq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if esq.unique != nil && *esq.unique {
		selector.Distinct()
	}
	for _, p := range esq.predicates {
		p(selector)
	}
	for _, p := range esq.order {
		p(selector)
	}
	if offset := esq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntitySkillGroupBy is the group-by builder for EntitySkill entities.
type EntitySkillGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esgb *EntitySkillGroupBy) Aggregate(fns ...AggregateFunc) *EntitySkillGroupBy {
	esgb.fns = append(esgb.fns, fns...)
	return esgb
}

// Scan applies the group-by query and scans the result into the given value.
func (esgb *EntitySkillGroupBy) Scan(ctx context.Context, v any) error {
	query, err := esgb.path(ctx)
	if err != nil {
		return err
	}
	esgb.sql = query
	return esgb.sqlScan(ctx, v)
}

func (esgb *EntitySkillGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range esgb.fields {
		if !entityskill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := esgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (esgb *EntitySkillGroupBy) sqlQuery() *sql.Selector {
	selector := esgb.sql.Select()
	aggregation := make([]string, 0, len(esgb.fns))
	for _, fn := range esgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(esgb.fields)+len(esgb.fns))
		for _, f := range esgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(esgb.fields...)...)
}

// EntitySkillSelect is the builder for selecting fields of EntitySkill entities.
type EntitySkillSelect struct {
	*EntitySkillQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ess *EntitySkillSelect) Aggregate(fns ...AggregateFunc) *EntitySkillSelect {
	ess.fns = append(ess.fns, fns...)
	return ess
}

// Scan applies the selector query and scans the result into the given value.
func (ess *EntitySkillSelect) Scan(ctx context.Context, v any) error {
	if err := ess.prepareQuery(ctx); err != nil {
		return err
	}
	ess.sql = ess.EntitySkillQuery.sqlQuery(ctx)
	return ess.sqlScan(ctx, v)
}

func (ess *EntitySkillSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ess.fns))
	for _, fn := range ess.fns {
		aggregation = append(aggregation, fn(ess.sql))
	}
	switch n := len(*ess.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ess.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ess.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ess.sql.Query()
	if err := ess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

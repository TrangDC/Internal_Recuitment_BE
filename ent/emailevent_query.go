// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/emailevent"
	"trec/ent/emailtemplate"
	"trec/ent/outgoingemail"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailEventQuery is the builder for querying EmailEvent entities.
type EmailEventQuery struct {
	config
	limit                       *int
	offset                      *int
	unique                      *bool
	order                       []OrderFunc
	fields                      []string
	predicates                  []predicate.EmailEvent
	withTemplateEdges           *EmailTemplateQuery
	withOutgoingEmailEdges      *OutgoingEmailQuery
	modifiers                   []func(*sql.Selector)
	loadTotal                   []func(context.Context, []*EmailEvent) error
	withNamedTemplateEdges      map[string]*EmailTemplateQuery
	withNamedOutgoingEmailEdges map[string]*OutgoingEmailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailEventQuery builder.
func (eeq *EmailEventQuery) Where(ps ...predicate.EmailEvent) *EmailEventQuery {
	eeq.predicates = append(eeq.predicates, ps...)
	return eeq
}

// Limit adds a limit step to the query.
func (eeq *EmailEventQuery) Limit(limit int) *EmailEventQuery {
	eeq.limit = &limit
	return eeq
}

// Offset adds an offset step to the query.
func (eeq *EmailEventQuery) Offset(offset int) *EmailEventQuery {
	eeq.offset = &offset
	return eeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eeq *EmailEventQuery) Unique(unique bool) *EmailEventQuery {
	eeq.unique = &unique
	return eeq
}

// Order adds an order step to the query.
func (eeq *EmailEventQuery) Order(o ...OrderFunc) *EmailEventQuery {
	eeq.order = append(eeq.order, o...)
	return eeq
}

// QueryTemplateEdges chains the current query on the "template_edges" edge.
func (eeq *EmailEventQuery) QueryTemplateEdges() *EmailTemplateQuery {
	query := &EmailTemplateQuery{config: eeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emailevent.Table, emailevent.FieldID, selector),
			sqlgraph.To(emailtemplate.Table, emailtemplate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, emailevent.TemplateEdgesTable, emailevent.TemplateEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOutgoingEmailEdges chains the current query on the "outgoing_email_edges" edge.
func (eeq *EmailEventQuery) QueryOutgoingEmailEdges() *OutgoingEmailQuery {
	query := &OutgoingEmailQuery{config: eeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emailevent.Table, emailevent.FieldID, selector),
			sqlgraph.To(outgoingemail.Table, outgoingemail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, emailevent.OutgoingEmailEdgesTable, emailevent.OutgoingEmailEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmailEvent entity from the query.
// Returns a *NotFoundError when no EmailEvent was found.
func (eeq *EmailEventQuery) First(ctx context.Context) (*EmailEvent, error) {
	nodes, err := eeq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eeq *EmailEventQuery) FirstX(ctx context.Context) *EmailEvent {
	node, err := eeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailEvent ID from the query.
// Returns a *NotFoundError when no EmailEvent ID was found.
func (eeq *EmailEventQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eeq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eeq *EmailEventQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := eeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailEvent entity is found.
// Returns a *NotFoundError when no EmailEvent entities are found.
func (eeq *EmailEventQuery) Only(ctx context.Context) (*EmailEvent, error) {
	nodes, err := eeq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailevent.Label}
	default:
		return nil, &NotSingularError{emailevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eeq *EmailEventQuery) OnlyX(ctx context.Context) *EmailEvent {
	node, err := eeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailEvent ID in the query.
// Returns a *NotSingularError when more than one EmailEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (eeq *EmailEventQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eeq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailevent.Label}
	default:
		err = &NotSingularError{emailevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eeq *EmailEventQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := eeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailEvents.
func (eeq *EmailEventQuery) All(ctx context.Context) ([]*EmailEvent, error) {
	if err := eeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eeq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eeq *EmailEventQuery) AllX(ctx context.Context) []*EmailEvent {
	nodes, err := eeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailEvent IDs.
func (eeq *EmailEventQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := eeq.Select(emailevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eeq *EmailEventQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := eeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eeq *EmailEventQuery) Count(ctx context.Context) (int, error) {
	if err := eeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eeq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eeq *EmailEventQuery) CountX(ctx context.Context) int {
	count, err := eeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eeq *EmailEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := eeq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eeq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eeq *EmailEventQuery) ExistX(ctx context.Context) bool {
	exist, err := eeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eeq *EmailEventQuery) Clone() *EmailEventQuery {
	if eeq == nil {
		return nil
	}
	return &EmailEventQuery{
		config:                 eeq.config,
		limit:                  eeq.limit,
		offset:                 eeq.offset,
		order:                  append([]OrderFunc{}, eeq.order...),
		predicates:             append([]predicate.EmailEvent{}, eeq.predicates...),
		withTemplateEdges:      eeq.withTemplateEdges.Clone(),
		withOutgoingEmailEdges: eeq.withOutgoingEmailEdges.Clone(),
		// clone intermediate query.
		sql:    eeq.sql.Clone(),
		path:   eeq.path,
		unique: eeq.unique,
	}
}

// WithTemplateEdges tells the query-builder to eager-load the nodes that are connected to
// the "template_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (eeq *EmailEventQuery) WithTemplateEdges(opts ...func(*EmailTemplateQuery)) *EmailEventQuery {
	query := &EmailTemplateQuery{config: eeq.config}
	for _, opt := range opts {
		opt(query)
	}
	eeq.withTemplateEdges = query
	return eeq
}

// WithOutgoingEmailEdges tells the query-builder to eager-load the nodes that are connected to
// the "outgoing_email_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (eeq *EmailEventQuery) WithOutgoingEmailEdges(opts ...func(*OutgoingEmailQuery)) *EmailEventQuery {
	query := &OutgoingEmailQuery{config: eeq.config}
	for _, opt := range opts {
		opt(query)
	}
	eeq.withOutgoingEmailEdges = query
	return eeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Module emailevent.Module `json:"module,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailEvent.Query().
//		GroupBy(emailevent.FieldModule).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eeq *EmailEventQuery) GroupBy(field string, fields ...string) *EmailEventGroupBy {
	grbuild := &EmailEventGroupBy{config: eeq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eeq.sqlQuery(ctx), nil
	}
	grbuild.label = emailevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Module emailevent.Module `json:"module,omitempty"`
//	}
//
//	client.EmailEvent.Query().
//		Select(emailevent.FieldModule).
//		Scan(ctx, &v)
func (eeq *EmailEventQuery) Select(fields ...string) *EmailEventSelect {
	eeq.fields = append(eeq.fields, fields...)
	selbuild := &EmailEventSelect{EmailEventQuery: eeq}
	selbuild.label = emailevent.Label
	selbuild.flds, selbuild.scan = &eeq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a EmailEventSelect configured with the given aggregations.
func (eeq *EmailEventQuery) Aggregate(fns ...AggregateFunc) *EmailEventSelect {
	return eeq.Select().Aggregate(fns...)
}

func (eeq *EmailEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eeq.fields {
		if !emailevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eeq.path != nil {
		prev, err := eeq.path(ctx)
		if err != nil {
			return err
		}
		eeq.sql = prev
	}
	return nil
}

func (eeq *EmailEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailEvent, error) {
	var (
		nodes       = []*EmailEvent{}
		_spec       = eeq.querySpec()
		loadedTypes = [2]bool{
			eeq.withTemplateEdges != nil,
			eeq.withOutgoingEmailEdges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailEvent{config: eeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(eeq.modifiers) > 0 {
		_spec.Modifiers = eeq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eeq.withTemplateEdges; query != nil {
		if err := eeq.loadTemplateEdges(ctx, query, nodes,
			func(n *EmailEvent) { n.Edges.TemplateEdges = []*EmailTemplate{} },
			func(n *EmailEvent, e *EmailTemplate) { n.Edges.TemplateEdges = append(n.Edges.TemplateEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := eeq.withOutgoingEmailEdges; query != nil {
		if err := eeq.loadOutgoingEmailEdges(ctx, query, nodes,
			func(n *EmailEvent) { n.Edges.OutgoingEmailEdges = []*OutgoingEmail{} },
			func(n *EmailEvent, e *OutgoingEmail) {
				n.Edges.OutgoingEmailEdges = append(n.Edges.OutgoingEmailEdges, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range eeq.withNamedTemplateEdges {
		if err := eeq.loadTemplateEdges(ctx, query, nodes,
			func(n *EmailEvent) { n.appendNamedTemplateEdges(name) },
			func(n *EmailEvent, e *EmailTemplate) { n.appendNamedTemplateEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range eeq.withNamedOutgoingEmailEdges {
		if err := eeq.loadOutgoingEmailEdges(ctx, query, nodes,
			func(n *EmailEvent) { n.appendNamedOutgoingEmailEdges(name) },
			func(n *EmailEvent, e *OutgoingEmail) { n.appendNamedOutgoingEmailEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range eeq.loadTotal {
		if err := eeq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eeq *EmailEventQuery) loadTemplateEdges(ctx context.Context, query *EmailTemplateQuery, nodes []*EmailEvent, init func(*EmailEvent), assign func(*EmailEvent, *EmailTemplate)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*EmailEvent)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.InValues(emailevent.TemplateEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EventID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eeq *EmailEventQuery) loadOutgoingEmailEdges(ctx context.Context, query *OutgoingEmailQuery, nodes []*EmailEvent, init func(*EmailEvent), assign func(*EmailEvent, *OutgoingEmail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*EmailEvent)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.OutgoingEmail(func(s *sql.Selector) {
		s.Where(sql.InValues(emailevent.OutgoingEmailEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EventID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eeq *EmailEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eeq.querySpec()
	if len(eeq.modifiers) > 0 {
		_spec.Modifiers = eeq.modifiers
	}
	_spec.Node.Columns = eeq.fields
	if len(eeq.fields) > 0 {
		_spec.Unique = eeq.unique != nil && *eeq.unique
	}
	return sqlgraph.CountNodes(ctx, eeq.driver, _spec)
}

func (eeq *EmailEventQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eeq *EmailEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailevent.Table,
			Columns: emailevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailevent.FieldID,
			},
		},
		From:   eeq.sql,
		Unique: true,
	}
	if unique := eeq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eeq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailevent.FieldID)
		for i := range fields {
			if fields[i] != emailevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eeq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eeq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eeq *EmailEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eeq.driver.Dialect())
	t1 := builder.Table(emailevent.Table)
	columns := eeq.fields
	if len(columns) == 0 {
		columns = emailevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eeq.sql != nil {
		selector = eeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eeq.unique != nil && *eeq.unique {
		selector.Distinct()
	}
	for _, p := range eeq.predicates {
		p(selector)
	}
	for _, p := range eeq.order {
		p(selector)
	}
	if offset := eeq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eeq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTemplateEdges tells the query-builder to eager-load the nodes that are connected to the "template_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (eeq *EmailEventQuery) WithNamedTemplateEdges(name string, opts ...func(*EmailTemplateQuery)) *EmailEventQuery {
	query := &EmailTemplateQuery{config: eeq.config}
	for _, opt := range opts {
		opt(query)
	}
	if eeq.withNamedTemplateEdges == nil {
		eeq.withNamedTemplateEdges = make(map[string]*EmailTemplateQuery)
	}
	eeq.withNamedTemplateEdges[name] = query
	return eeq
}

// WithNamedOutgoingEmailEdges tells the query-builder to eager-load the nodes that are connected to the "outgoing_email_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (eeq *EmailEventQuery) WithNamedOutgoingEmailEdges(name string, opts ...func(*OutgoingEmailQuery)) *EmailEventQuery {
	query := &OutgoingEmailQuery{config: eeq.config}
	for _, opt := range opts {
		opt(query)
	}
	if eeq.withNamedOutgoingEmailEdges == nil {
		eeq.withNamedOutgoingEmailEdges = make(map[string]*OutgoingEmailQuery)
	}
	eeq.withNamedOutgoingEmailEdges[name] = query
	return eeq
}

// EmailEventGroupBy is the group-by builder for EmailEvent entities.
type EmailEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eegb *EmailEventGroupBy) Aggregate(fns ...AggregateFunc) *EmailEventGroupBy {
	eegb.fns = append(eegb.fns, fns...)
	return eegb
}

// Scan applies the group-by query and scans the result into the given value.
func (eegb *EmailEventGroupBy) Scan(ctx context.Context, v any) error {
	query, err := eegb.path(ctx)
	if err != nil {
		return err
	}
	eegb.sql = query
	return eegb.sqlScan(ctx, v)
}

func (eegb *EmailEventGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range eegb.fields {
		if !emailevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := eegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eegb *EmailEventGroupBy) sqlQuery() *sql.Selector {
	selector := eegb.sql.Select()
	aggregation := make([]string, 0, len(eegb.fns))
	for _, fn := range eegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(eegb.fields)+len(eegb.fns))
		for _, f := range eegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(eegb.fields...)...)
}

// EmailEventSelect is the builder for selecting fields of EmailEvent entities.
type EmailEventSelect struct {
	*EmailEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ees *EmailEventSelect) Aggregate(fns ...AggregateFunc) *EmailEventSelect {
	ees.fns = append(ees.fns, fns...)
	return ees
}

// Scan applies the selector query and scans the result into the given value.
func (ees *EmailEventSelect) Scan(ctx context.Context, v any) error {
	if err := ees.prepareQuery(ctx); err != nil {
		return err
	}
	ees.sql = ees.EmailEventQuery.sqlQuery(ctx)
	return ees.sqlScan(ctx, v)
}

func (ees *EmailEventSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ees.fns))
	for _, fn := range ees.fns {
		aggregation = append(aggregation, fn(ees.sql))
	}
	switch n := len(*ees.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ees.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ees.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ees.sql.Query()
	if err := ees.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

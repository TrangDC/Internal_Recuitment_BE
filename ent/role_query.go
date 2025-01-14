// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"trec/ent/emailroleattribute"
	"trec/ent/emailtemplate"
	"trec/ent/entitypermission"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/user"
	"trec/ent/userrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RoleQuery is the builder for querying Role entities.
type RoleQuery struct {
	config
	limit                        *int
	offset                       *int
	unique                       *bool
	order                        []OrderFunc
	fields                       []string
	predicates                   []predicate.Role
	withRolePermissionEdges      *EntityPermissionQuery
	withUserEdges                *UserQuery
	withEmailTemplateEdges       *EmailTemplateQuery
	withUserRoles                *UserRoleQuery
	withEmailTemplateRoles       *EmailRoleAttributeQuery
	modifiers                    []func(*sql.Selector)
	loadTotal                    []func(context.Context, []*Role) error
	withNamedRolePermissionEdges map[string]*EntityPermissionQuery
	withNamedUserEdges           map[string]*UserQuery
	withNamedEmailTemplateEdges  map[string]*EmailTemplateQuery
	withNamedUserRoles           map[string]*UserRoleQuery
	withNamedEmailTemplateRoles  map[string]*EmailRoleAttributeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleQuery builder.
func (rq *RoleQuery) Where(ps ...predicate.Role) *RoleQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RoleQuery) Limit(limit int) *RoleQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RoleQuery) Offset(offset int) *RoleQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoleQuery) Unique(unique bool) *RoleQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RoleQuery) Order(o ...OrderFunc) *RoleQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryRolePermissionEdges chains the current query on the "role_permission_edges" edge.
func (rq *RoleQuery) QueryRolePermissionEdges() *EntityPermissionQuery {
	query := &EntityPermissionQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(entitypermission.Table, entitypermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.RolePermissionEdgesTable, role.RolePermissionEdgesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserEdges chains the current query on the "user_edges" edge.
func (rq *RoleQuery) QueryUserEdges() *UserQuery {
	query := &UserQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.UserEdgesTable, role.UserEdgesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmailTemplateEdges chains the current query on the "email_template_edges" edge.
func (rq *RoleQuery) QueryEmailTemplateEdges() *EmailTemplateQuery {
	query := &EmailTemplateQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(emailtemplate.Table, emailtemplate.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.EmailTemplateEdgesTable, role.EmailTemplateEdgesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserRoles chains the current query on the "user_roles" edge.
func (rq *RoleQuery) QueryUserRoles() *UserRoleQuery {
	query := &UserRoleQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(userrole.Table, userrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, role.UserRolesTable, role.UserRolesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmailTemplateRoles chains the current query on the "email_template_roles" edge.
func (rq *RoleQuery) QueryEmailTemplateRoles() *EmailRoleAttributeQuery {
	query := &EmailRoleAttributeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, selector),
			sqlgraph.To(emailroleattribute.Table, emailroleattribute.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, role.EmailTemplateRolesTable, role.EmailTemplateRolesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Role entity from the query.
// Returns a *NotFoundError when no Role was found.
func (rq *RoleQuery) First(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{role.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoleQuery) FirstX(ctx context.Context) *Role {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Role ID from the query.
// Returns a *NotFoundError when no Role ID was found.
func (rq *RoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{role.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Role entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Role entity is found.
// Returns a *NotFoundError when no Role entities are found.
func (rq *RoleQuery) Only(ctx context.Context) (*Role, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{role.Label}
	default:
		return nil, &NotSingularError{role.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoleQuery) OnlyX(ctx context.Context) *Role {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Role ID in the query.
// Returns a *NotSingularError when more than one Role ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{role.Label}
	default:
		err = &NotSingularError{role.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Roles.
func (rq *RoleQuery) All(ctx context.Context) ([]*Role, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoleQuery) AllX(ctx context.Context) []*Role {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Role IDs.
func (rq *RoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rq.Select(role.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoleQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoleQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoleQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoleQuery) Clone() *RoleQuery {
	if rq == nil {
		return nil
	}
	return &RoleQuery{
		config:                  rq.config,
		limit:                   rq.limit,
		offset:                  rq.offset,
		order:                   append([]OrderFunc{}, rq.order...),
		predicates:              append([]predicate.Role{}, rq.predicates...),
		withRolePermissionEdges: rq.withRolePermissionEdges.Clone(),
		withUserEdges:           rq.withUserEdges.Clone(),
		withEmailTemplateEdges:  rq.withEmailTemplateEdges.Clone(),
		withUserRoles:           rq.withUserRoles.Clone(),
		withEmailTemplateRoles:  rq.withEmailTemplateRoles.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithRolePermissionEdges tells the query-builder to eager-load the nodes that are connected to
// the "role_permission_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithRolePermissionEdges(opts ...func(*EntityPermissionQuery)) *RoleQuery {
	query := &EntityPermissionQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRolePermissionEdges = query
	return rq
}

// WithUserEdges tells the query-builder to eager-load the nodes that are connected to
// the "user_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithUserEdges(opts ...func(*UserQuery)) *RoleQuery {
	query := &UserQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withUserEdges = query
	return rq
}

// WithEmailTemplateEdges tells the query-builder to eager-load the nodes that are connected to
// the "email_template_edges" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithEmailTemplateEdges(opts ...func(*EmailTemplateQuery)) *RoleQuery {
	query := &EmailTemplateQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withEmailTemplateEdges = query
	return rq
}

// WithUserRoles tells the query-builder to eager-load the nodes that are connected to
// the "user_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithUserRoles(opts ...func(*UserRoleQuery)) *RoleQuery {
	query := &UserRoleQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withUserRoles = query
	return rq
}

// WithEmailTemplateRoles tells the query-builder to eager-load the nodes that are connected to
// the "email_template_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithEmailTemplateRoles(opts ...func(*EmailRoleAttributeQuery)) *RoleQuery {
	query := &EmailRoleAttributeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withEmailTemplateRoles = query
	return rq
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
//	client.Role.Query().
//		GroupBy(role.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RoleQuery) GroupBy(field string, fields ...string) *RoleGroupBy {
	grbuild := &RoleGroupBy{config: rq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	grbuild.label = role.Label
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
//	client.Role.Query().
//		Select(role.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *RoleQuery) Select(fields ...string) *RoleSelect {
	rq.fields = append(rq.fields, fields...)
	selbuild := &RoleSelect{RoleQuery: rq}
	selbuild.label = role.Label
	selbuild.flds, selbuild.scan = &rq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a RoleSelect configured with the given aggregations.
func (rq *RoleQuery) Aggregate(fns ...AggregateFunc) *RoleSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !role.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Role, error) {
	var (
		nodes       = []*Role{}
		_spec       = rq.querySpec()
		loadedTypes = [5]bool{
			rq.withRolePermissionEdges != nil,
			rq.withUserEdges != nil,
			rq.withEmailTemplateEdges != nil,
			rq.withUserRoles != nil,
			rq.withEmailTemplateRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Role).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Role{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withRolePermissionEdges; query != nil {
		if err := rq.loadRolePermissionEdges(ctx, query, nodes,
			func(n *Role) { n.Edges.RolePermissionEdges = []*EntityPermission{} },
			func(n *Role, e *EntityPermission) {
				n.Edges.RolePermissionEdges = append(n.Edges.RolePermissionEdges, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := rq.withUserEdges; query != nil {
		if err := rq.loadUserEdges(ctx, query, nodes,
			func(n *Role) { n.Edges.UserEdges = []*User{} },
			func(n *Role, e *User) { n.Edges.UserEdges = append(n.Edges.UserEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withEmailTemplateEdges; query != nil {
		if err := rq.loadEmailTemplateEdges(ctx, query, nodes,
			func(n *Role) { n.Edges.EmailTemplateEdges = []*EmailTemplate{} },
			func(n *Role, e *EmailTemplate) { n.Edges.EmailTemplateEdges = append(n.Edges.EmailTemplateEdges, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withUserRoles; query != nil {
		if err := rq.loadUserRoles(ctx, query, nodes,
			func(n *Role) { n.Edges.UserRoles = []*UserRole{} },
			func(n *Role, e *UserRole) { n.Edges.UserRoles = append(n.Edges.UserRoles, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withEmailTemplateRoles; query != nil {
		if err := rq.loadEmailTemplateRoles(ctx, query, nodes,
			func(n *Role) { n.Edges.EmailTemplateRoles = []*EmailRoleAttribute{} },
			func(n *Role, e *EmailRoleAttribute) {
				n.Edges.EmailTemplateRoles = append(n.Edges.EmailTemplateRoles, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedRolePermissionEdges {
		if err := rq.loadRolePermissionEdges(ctx, query, nodes,
			func(n *Role) { n.appendNamedRolePermissionEdges(name) },
			func(n *Role, e *EntityPermission) { n.appendNamedRolePermissionEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedUserEdges {
		if err := rq.loadUserEdges(ctx, query, nodes,
			func(n *Role) { n.appendNamedUserEdges(name) },
			func(n *Role, e *User) { n.appendNamedUserEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedEmailTemplateEdges {
		if err := rq.loadEmailTemplateEdges(ctx, query, nodes,
			func(n *Role) { n.appendNamedEmailTemplateEdges(name) },
			func(n *Role, e *EmailTemplate) { n.appendNamedEmailTemplateEdges(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedUserRoles {
		if err := rq.loadUserRoles(ctx, query, nodes,
			func(n *Role) { n.appendNamedUserRoles(name) },
			func(n *Role, e *UserRole) { n.appendNamedUserRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rq.withNamedEmailTemplateRoles {
		if err := rq.loadEmailTemplateRoles(ctx, query, nodes,
			func(n *Role) { n.appendNamedEmailTemplateRoles(name) },
			func(n *Role, e *EmailRoleAttribute) { n.appendNamedEmailTemplateRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoleQuery) loadRolePermissionEdges(ctx context.Context, query *EntityPermissionQuery, nodes []*Role, init func(*Role), assign func(*Role, *EntityPermission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Role)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.EntityPermission(func(s *sql.Selector) {
		s.Where(sql.InValues(role.RolePermissionEdgesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EntityID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RoleQuery) loadUserEdges(ctx context.Context, query *UserQuery, nodes []*Role, init func(*Role), assign func(*Role, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Role)
	nids := make(map[uuid.UUID]map[*Role]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(role.UserEdgesTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(role.UserEdgesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(role.UserEdgesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(role.UserEdgesPrimaryKey[1]))
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
				nids[inValue] = map[*Role]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "user_edges" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RoleQuery) loadEmailTemplateEdges(ctx context.Context, query *EmailTemplateQuery, nodes []*Role, init func(*Role), assign func(*Role, *EmailTemplate)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Role)
	nids := make(map[uuid.UUID]map[*Role]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(role.EmailTemplateEdgesTable)
		s.Join(joinT).On(s.C(emailtemplate.FieldID), joinT.C(role.EmailTemplateEdgesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(role.EmailTemplateEdgesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(role.EmailTemplateEdgesPrimaryKey[0]))
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
				nids[inValue] = map[*Role]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "email_template_edges" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *RoleQuery) loadUserRoles(ctx context.Context, query *UserRoleQuery, nodes []*Role, init func(*Role), assign func(*Role, *UserRole)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Role)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.InValues(role.UserRolesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoleID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "role_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RoleQuery) loadEmailTemplateRoles(ctx context.Context, query *EmailRoleAttributeQuery, nodes []*Role, init func(*Role), assign func(*Role, *EmailRoleAttribute)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Role)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.EmailRoleAttribute(func(s *sql.Selector) {
		s.Where(sql.InValues(role.EmailTemplateRolesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoleID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "role_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoleQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rq *RoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: role.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for i := range fields {
			if fields[i] != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(role.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = role.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRolePermissionEdges tells the query-builder to eager-load the nodes that are connected to the "role_permission_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithNamedRolePermissionEdges(name string, opts ...func(*EntityPermissionQuery)) *RoleQuery {
	query := &EntityPermissionQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedRolePermissionEdges == nil {
		rq.withNamedRolePermissionEdges = make(map[string]*EntityPermissionQuery)
	}
	rq.withNamedRolePermissionEdges[name] = query
	return rq
}

// WithNamedUserEdges tells the query-builder to eager-load the nodes that are connected to the "user_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithNamedUserEdges(name string, opts ...func(*UserQuery)) *RoleQuery {
	query := &UserQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedUserEdges == nil {
		rq.withNamedUserEdges = make(map[string]*UserQuery)
	}
	rq.withNamedUserEdges[name] = query
	return rq
}

// WithNamedEmailTemplateEdges tells the query-builder to eager-load the nodes that are connected to the "email_template_edges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithNamedEmailTemplateEdges(name string, opts ...func(*EmailTemplateQuery)) *RoleQuery {
	query := &EmailTemplateQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedEmailTemplateEdges == nil {
		rq.withNamedEmailTemplateEdges = make(map[string]*EmailTemplateQuery)
	}
	rq.withNamedEmailTemplateEdges[name] = query
	return rq
}

// WithNamedUserRoles tells the query-builder to eager-load the nodes that are connected to the "user_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithNamedUserRoles(name string, opts ...func(*UserRoleQuery)) *RoleQuery {
	query := &UserRoleQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedUserRoles == nil {
		rq.withNamedUserRoles = make(map[string]*UserRoleQuery)
	}
	rq.withNamedUserRoles[name] = query
	return rq
}

// WithNamedEmailTemplateRoles tells the query-builder to eager-load the nodes that are connected to the "email_template_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rq *RoleQuery) WithNamedEmailTemplateRoles(name string, opts ...func(*EmailRoleAttributeQuery)) *RoleQuery {
	query := &EmailRoleAttributeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	if rq.withNamedEmailTemplateRoles == nil {
		rq.withNamedEmailTemplateRoles = make(map[string]*EmailRoleAttributeQuery)
	}
	rq.withNamedEmailTemplateRoles[name] = query
	return rq
}

// RoleGroupBy is the group-by builder for Role entities.
type RoleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoleGroupBy) Aggregate(fns ...AggregateFunc) *RoleGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RoleGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

func (rgb *RoleGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rgb.fields {
		if !role.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RoleGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RoleSelect is the builder for selecting fields of Role entities.
type RoleSelect struct {
	*RoleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoleSelect) Aggregate(fns ...AggregateFunc) *RoleSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoleSelect) Scan(ctx context.Context, v any) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RoleQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

func (rs *RoleSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(rs.sql))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

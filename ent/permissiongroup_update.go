// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/permission"
	"trec/ent/permissiongroup"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionGroupUpdate is the builder for updating PermissionGroup entities.
type PermissionGroupUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionGroupMutation
}

// Where appends a list predicates to the PermissionGroupUpdate builder.
func (pgu *PermissionGroupUpdate) Where(ps ...predicate.PermissionGroup) *PermissionGroupUpdate {
	pgu.mutation.Where(ps...)
	return pgu
}

// SetUpdatedAt sets the "updated_at" field.
func (pgu *PermissionGroupUpdate) SetUpdatedAt(t time.Time) *PermissionGroupUpdate {
	pgu.mutation.SetUpdatedAt(t)
	return pgu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pgu *PermissionGroupUpdate) SetNillableUpdatedAt(t *time.Time) *PermissionGroupUpdate {
	if t != nil {
		pgu.SetUpdatedAt(*t)
	}
	return pgu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pgu *PermissionGroupUpdate) ClearUpdatedAt() *PermissionGroupUpdate {
	pgu.mutation.ClearUpdatedAt()
	return pgu
}

// SetDeletedAt sets the "deleted_at" field.
func (pgu *PermissionGroupUpdate) SetDeletedAt(t time.Time) *PermissionGroupUpdate {
	pgu.mutation.SetDeletedAt(t)
	return pgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pgu *PermissionGroupUpdate) SetNillableDeletedAt(t *time.Time) *PermissionGroupUpdate {
	if t != nil {
		pgu.SetDeletedAt(*t)
	}
	return pgu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pgu *PermissionGroupUpdate) ClearDeletedAt() *PermissionGroupUpdate {
	pgu.mutation.ClearDeletedAt()
	return pgu
}

// SetTitle sets the "title" field.
func (pgu *PermissionGroupUpdate) SetTitle(s string) *PermissionGroupUpdate {
	pgu.mutation.SetTitle(s)
	return pgu
}

// SetParentID sets the "parent_id" field.
func (pgu *PermissionGroupUpdate) SetParentID(u uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.SetParentID(u)
	return pgu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pgu *PermissionGroupUpdate) SetNillableParentID(u *uuid.UUID) *PermissionGroupUpdate {
	if u != nil {
		pgu.SetParentID(*u)
	}
	return pgu
}

// ClearParentID clears the value of the "parent_id" field.
func (pgu *PermissionGroupUpdate) ClearParentID() *PermissionGroupUpdate {
	pgu.mutation.ClearParentID()
	return pgu
}

// SetGroupType sets the "group_type" field.
func (pgu *PermissionGroupUpdate) SetGroupType(pt permissiongroup.GroupType) *PermissionGroupUpdate {
	pgu.mutation.SetGroupType(pt)
	return pgu
}

// SetNillableGroupType sets the "group_type" field if the given value is not nil.
func (pgu *PermissionGroupUpdate) SetNillableGroupType(pt *permissiongroup.GroupType) *PermissionGroupUpdate {
	if pt != nil {
		pgu.SetGroupType(*pt)
	}
	return pgu
}

// SetOrderID sets the "order_id" field.
func (pgu *PermissionGroupUpdate) SetOrderID(i int) *PermissionGroupUpdate {
	pgu.mutation.ResetOrderID()
	pgu.mutation.SetOrderID(i)
	return pgu
}

// AddOrderID adds i to the "order_id" field.
func (pgu *PermissionGroupUpdate) AddOrderID(i int) *PermissionGroupUpdate {
	pgu.mutation.AddOrderID(i)
	return pgu
}

// SetGroupPermissionParentID sets the "group_permission_parent" edge to the PermissionGroup entity by ID.
func (pgu *PermissionGroupUpdate) SetGroupPermissionParentID(id uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.SetGroupPermissionParentID(id)
	return pgu
}

// SetNillableGroupPermissionParentID sets the "group_permission_parent" edge to the PermissionGroup entity by ID if the given value is not nil.
func (pgu *PermissionGroupUpdate) SetNillableGroupPermissionParentID(id *uuid.UUID) *PermissionGroupUpdate {
	if id != nil {
		pgu = pgu.SetGroupPermissionParentID(*id)
	}
	return pgu
}

// SetGroupPermissionParent sets the "group_permission_parent" edge to the PermissionGroup entity.
func (pgu *PermissionGroupUpdate) SetGroupPermissionParent(p *PermissionGroup) *PermissionGroupUpdate {
	return pgu.SetGroupPermissionParentID(p.ID)
}

// AddGroupPermissionChildIDs adds the "group_permission_children" edge to the PermissionGroup entity by IDs.
func (pgu *PermissionGroupUpdate) AddGroupPermissionChildIDs(ids ...uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.AddGroupPermissionChildIDs(ids...)
	return pgu
}

// AddGroupPermissionChildren adds the "group_permission_children" edges to the PermissionGroup entity.
func (pgu *PermissionGroupUpdate) AddGroupPermissionChildren(p ...*PermissionGroup) *PermissionGroupUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.AddGroupPermissionChildIDs(ids...)
}

// AddPermissionEdgeIDs adds the "permission_edges" edge to the Permission entity by IDs.
func (pgu *PermissionGroupUpdate) AddPermissionEdgeIDs(ids ...uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.AddPermissionEdgeIDs(ids...)
	return pgu
}

// AddPermissionEdges adds the "permission_edges" edges to the Permission entity.
func (pgu *PermissionGroupUpdate) AddPermissionEdges(p ...*Permission) *PermissionGroupUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.AddPermissionEdgeIDs(ids...)
}

// Mutation returns the PermissionGroupMutation object of the builder.
func (pgu *PermissionGroupUpdate) Mutation() *PermissionGroupMutation {
	return pgu.mutation
}

// ClearGroupPermissionParent clears the "group_permission_parent" edge to the PermissionGroup entity.
func (pgu *PermissionGroupUpdate) ClearGroupPermissionParent() *PermissionGroupUpdate {
	pgu.mutation.ClearGroupPermissionParent()
	return pgu
}

// ClearGroupPermissionChildren clears all "group_permission_children" edges to the PermissionGroup entity.
func (pgu *PermissionGroupUpdate) ClearGroupPermissionChildren() *PermissionGroupUpdate {
	pgu.mutation.ClearGroupPermissionChildren()
	return pgu
}

// RemoveGroupPermissionChildIDs removes the "group_permission_children" edge to PermissionGroup entities by IDs.
func (pgu *PermissionGroupUpdate) RemoveGroupPermissionChildIDs(ids ...uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.RemoveGroupPermissionChildIDs(ids...)
	return pgu
}

// RemoveGroupPermissionChildren removes "group_permission_children" edges to PermissionGroup entities.
func (pgu *PermissionGroupUpdate) RemoveGroupPermissionChildren(p ...*PermissionGroup) *PermissionGroupUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.RemoveGroupPermissionChildIDs(ids...)
}

// ClearPermissionEdges clears all "permission_edges" edges to the Permission entity.
func (pgu *PermissionGroupUpdate) ClearPermissionEdges() *PermissionGroupUpdate {
	pgu.mutation.ClearPermissionEdges()
	return pgu
}

// RemovePermissionEdgeIDs removes the "permission_edges" edge to Permission entities by IDs.
func (pgu *PermissionGroupUpdate) RemovePermissionEdgeIDs(ids ...uuid.UUID) *PermissionGroupUpdate {
	pgu.mutation.RemovePermissionEdgeIDs(ids...)
	return pgu
}

// RemovePermissionEdges removes "permission_edges" edges to Permission entities.
func (pgu *PermissionGroupUpdate) RemovePermissionEdges(p ...*Permission) *PermissionGroupUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pgu.RemovePermissionEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pgu *PermissionGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pgu.hooks) == 0 {
		if err = pgu.check(); err != nil {
			return 0, err
		}
		affected, err = pgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pgu.check(); err != nil {
				return 0, err
			}
			pgu.mutation = mutation
			affected, err = pgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pgu.hooks) - 1; i >= 0; i-- {
			if pgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pgu *PermissionGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := pgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pgu *PermissionGroupUpdate) Exec(ctx context.Context) error {
	_, err := pgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgu *PermissionGroupUpdate) ExecX(ctx context.Context) {
	if err := pgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pgu *PermissionGroupUpdate) check() error {
	if v, ok := pgu.mutation.GroupType(); ok {
		if err := permissiongroup.GroupTypeValidator(v); err != nil {
			return &ValidationError{Name: "group_type", err: fmt.Errorf(`ent: validator failed for field "PermissionGroup.group_type": %w`, err)}
		}
	}
	return nil
}

func (pgu *PermissionGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permissiongroup.Table,
			Columns: permissiongroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permissiongroup.FieldID,
			},
		},
	}
	if ps := pgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pgu.mutation.UpdatedAt(); ok {
		_spec.SetField(permissiongroup.FieldUpdatedAt, field.TypeTime, value)
	}
	if pgu.mutation.UpdatedAtCleared() {
		_spec.ClearField(permissiongroup.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pgu.mutation.DeletedAt(); ok {
		_spec.SetField(permissiongroup.FieldDeletedAt, field.TypeTime, value)
	}
	if pgu.mutation.DeletedAtCleared() {
		_spec.ClearField(permissiongroup.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pgu.mutation.Title(); ok {
		_spec.SetField(permissiongroup.FieldTitle, field.TypeString, value)
	}
	if value, ok := pgu.mutation.GroupType(); ok {
		_spec.SetField(permissiongroup.FieldGroupType, field.TypeEnum, value)
	}
	if value, ok := pgu.mutation.OrderID(); ok {
		_spec.SetField(permissiongroup.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := pgu.mutation.AddedOrderID(); ok {
		_spec.AddField(permissiongroup.FieldOrderID, field.TypeInt, value)
	}
	if pgu.mutation.GroupPermissionParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissiongroup.GroupPermissionParentTable,
			Columns: []string{permissiongroup.GroupPermissionParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.GroupPermissionParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissiongroup.GroupPermissionParentTable,
			Columns: []string{permissiongroup.GroupPermissionParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pgu.mutation.GroupPermissionChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.RemovedGroupPermissionChildrenIDs(); len(nodes) > 0 && !pgu.mutation.GroupPermissionChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.GroupPermissionChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pgu.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.RemovedPermissionEdgesIDs(); len(nodes) > 0 && !pgu.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pgu.mutation.PermissionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissiongroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PermissionGroupUpdateOne is the builder for updating a single PermissionGroup entity.
type PermissionGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionGroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pguo *PermissionGroupUpdateOne) SetUpdatedAt(t time.Time) *PermissionGroupUpdateOne {
	pguo.mutation.SetUpdatedAt(t)
	return pguo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pguo *PermissionGroupUpdateOne) SetNillableUpdatedAt(t *time.Time) *PermissionGroupUpdateOne {
	if t != nil {
		pguo.SetUpdatedAt(*t)
	}
	return pguo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pguo *PermissionGroupUpdateOne) ClearUpdatedAt() *PermissionGroupUpdateOne {
	pguo.mutation.ClearUpdatedAt()
	return pguo
}

// SetDeletedAt sets the "deleted_at" field.
func (pguo *PermissionGroupUpdateOne) SetDeletedAt(t time.Time) *PermissionGroupUpdateOne {
	pguo.mutation.SetDeletedAt(t)
	return pguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pguo *PermissionGroupUpdateOne) SetNillableDeletedAt(t *time.Time) *PermissionGroupUpdateOne {
	if t != nil {
		pguo.SetDeletedAt(*t)
	}
	return pguo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pguo *PermissionGroupUpdateOne) ClearDeletedAt() *PermissionGroupUpdateOne {
	pguo.mutation.ClearDeletedAt()
	return pguo
}

// SetTitle sets the "title" field.
func (pguo *PermissionGroupUpdateOne) SetTitle(s string) *PermissionGroupUpdateOne {
	pguo.mutation.SetTitle(s)
	return pguo
}

// SetParentID sets the "parent_id" field.
func (pguo *PermissionGroupUpdateOne) SetParentID(u uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.SetParentID(u)
	return pguo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pguo *PermissionGroupUpdateOne) SetNillableParentID(u *uuid.UUID) *PermissionGroupUpdateOne {
	if u != nil {
		pguo.SetParentID(*u)
	}
	return pguo
}

// ClearParentID clears the value of the "parent_id" field.
func (pguo *PermissionGroupUpdateOne) ClearParentID() *PermissionGroupUpdateOne {
	pguo.mutation.ClearParentID()
	return pguo
}

// SetGroupType sets the "group_type" field.
func (pguo *PermissionGroupUpdateOne) SetGroupType(pt permissiongroup.GroupType) *PermissionGroupUpdateOne {
	pguo.mutation.SetGroupType(pt)
	return pguo
}

// SetNillableGroupType sets the "group_type" field if the given value is not nil.
func (pguo *PermissionGroupUpdateOne) SetNillableGroupType(pt *permissiongroup.GroupType) *PermissionGroupUpdateOne {
	if pt != nil {
		pguo.SetGroupType(*pt)
	}
	return pguo
}

// SetOrderID sets the "order_id" field.
func (pguo *PermissionGroupUpdateOne) SetOrderID(i int) *PermissionGroupUpdateOne {
	pguo.mutation.ResetOrderID()
	pguo.mutation.SetOrderID(i)
	return pguo
}

// AddOrderID adds i to the "order_id" field.
func (pguo *PermissionGroupUpdateOne) AddOrderID(i int) *PermissionGroupUpdateOne {
	pguo.mutation.AddOrderID(i)
	return pguo
}

// SetGroupPermissionParentID sets the "group_permission_parent" edge to the PermissionGroup entity by ID.
func (pguo *PermissionGroupUpdateOne) SetGroupPermissionParentID(id uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.SetGroupPermissionParentID(id)
	return pguo
}

// SetNillableGroupPermissionParentID sets the "group_permission_parent" edge to the PermissionGroup entity by ID if the given value is not nil.
func (pguo *PermissionGroupUpdateOne) SetNillableGroupPermissionParentID(id *uuid.UUID) *PermissionGroupUpdateOne {
	if id != nil {
		pguo = pguo.SetGroupPermissionParentID(*id)
	}
	return pguo
}

// SetGroupPermissionParent sets the "group_permission_parent" edge to the PermissionGroup entity.
func (pguo *PermissionGroupUpdateOne) SetGroupPermissionParent(p *PermissionGroup) *PermissionGroupUpdateOne {
	return pguo.SetGroupPermissionParentID(p.ID)
}

// AddGroupPermissionChildIDs adds the "group_permission_children" edge to the PermissionGroup entity by IDs.
func (pguo *PermissionGroupUpdateOne) AddGroupPermissionChildIDs(ids ...uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.AddGroupPermissionChildIDs(ids...)
	return pguo
}

// AddGroupPermissionChildren adds the "group_permission_children" edges to the PermissionGroup entity.
func (pguo *PermissionGroupUpdateOne) AddGroupPermissionChildren(p ...*PermissionGroup) *PermissionGroupUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.AddGroupPermissionChildIDs(ids...)
}

// AddPermissionEdgeIDs adds the "permission_edges" edge to the Permission entity by IDs.
func (pguo *PermissionGroupUpdateOne) AddPermissionEdgeIDs(ids ...uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.AddPermissionEdgeIDs(ids...)
	return pguo
}

// AddPermissionEdges adds the "permission_edges" edges to the Permission entity.
func (pguo *PermissionGroupUpdateOne) AddPermissionEdges(p ...*Permission) *PermissionGroupUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.AddPermissionEdgeIDs(ids...)
}

// Mutation returns the PermissionGroupMutation object of the builder.
func (pguo *PermissionGroupUpdateOne) Mutation() *PermissionGroupMutation {
	return pguo.mutation
}

// ClearGroupPermissionParent clears the "group_permission_parent" edge to the PermissionGroup entity.
func (pguo *PermissionGroupUpdateOne) ClearGroupPermissionParent() *PermissionGroupUpdateOne {
	pguo.mutation.ClearGroupPermissionParent()
	return pguo
}

// ClearGroupPermissionChildren clears all "group_permission_children" edges to the PermissionGroup entity.
func (pguo *PermissionGroupUpdateOne) ClearGroupPermissionChildren() *PermissionGroupUpdateOne {
	pguo.mutation.ClearGroupPermissionChildren()
	return pguo
}

// RemoveGroupPermissionChildIDs removes the "group_permission_children" edge to PermissionGroup entities by IDs.
func (pguo *PermissionGroupUpdateOne) RemoveGroupPermissionChildIDs(ids ...uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.RemoveGroupPermissionChildIDs(ids...)
	return pguo
}

// RemoveGroupPermissionChildren removes "group_permission_children" edges to PermissionGroup entities.
func (pguo *PermissionGroupUpdateOne) RemoveGroupPermissionChildren(p ...*PermissionGroup) *PermissionGroupUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.RemoveGroupPermissionChildIDs(ids...)
}

// ClearPermissionEdges clears all "permission_edges" edges to the Permission entity.
func (pguo *PermissionGroupUpdateOne) ClearPermissionEdges() *PermissionGroupUpdateOne {
	pguo.mutation.ClearPermissionEdges()
	return pguo
}

// RemovePermissionEdgeIDs removes the "permission_edges" edge to Permission entities by IDs.
func (pguo *PermissionGroupUpdateOne) RemovePermissionEdgeIDs(ids ...uuid.UUID) *PermissionGroupUpdateOne {
	pguo.mutation.RemovePermissionEdgeIDs(ids...)
	return pguo
}

// RemovePermissionEdges removes "permission_edges" edges to Permission entities.
func (pguo *PermissionGroupUpdateOne) RemovePermissionEdges(p ...*Permission) *PermissionGroupUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pguo.RemovePermissionEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pguo *PermissionGroupUpdateOne) Select(field string, fields ...string) *PermissionGroupUpdateOne {
	pguo.fields = append([]string{field}, fields...)
	return pguo
}

// Save executes the query and returns the updated PermissionGroup entity.
func (pguo *PermissionGroupUpdateOne) Save(ctx context.Context) (*PermissionGroup, error) {
	var (
		err  error
		node *PermissionGroup
	)
	if len(pguo.hooks) == 0 {
		if err = pguo.check(); err != nil {
			return nil, err
		}
		node, err = pguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pguo.check(); err != nil {
				return nil, err
			}
			pguo.mutation = mutation
			node, err = pguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pguo.hooks) - 1; i >= 0; i-- {
			if pguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PermissionGroup)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PermissionGroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pguo *PermissionGroupUpdateOne) SaveX(ctx context.Context) *PermissionGroup {
	node, err := pguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pguo *PermissionGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := pguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pguo *PermissionGroupUpdateOne) ExecX(ctx context.Context) {
	if err := pguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pguo *PermissionGroupUpdateOne) check() error {
	if v, ok := pguo.mutation.GroupType(); ok {
		if err := permissiongroup.GroupTypeValidator(v); err != nil {
			return &ValidationError{Name: "group_type", err: fmt.Errorf(`ent: validator failed for field "PermissionGroup.group_type": %w`, err)}
		}
	}
	return nil
}

func (pguo *PermissionGroupUpdateOne) sqlSave(ctx context.Context) (_node *PermissionGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permissiongroup.Table,
			Columns: permissiongroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permissiongroup.FieldID,
			},
		},
	}
	id, ok := pguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PermissionGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissiongroup.FieldID)
		for _, f := range fields {
			if !permissiongroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permissiongroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pguo.mutation.UpdatedAt(); ok {
		_spec.SetField(permissiongroup.FieldUpdatedAt, field.TypeTime, value)
	}
	if pguo.mutation.UpdatedAtCleared() {
		_spec.ClearField(permissiongroup.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pguo.mutation.DeletedAt(); ok {
		_spec.SetField(permissiongroup.FieldDeletedAt, field.TypeTime, value)
	}
	if pguo.mutation.DeletedAtCleared() {
		_spec.ClearField(permissiongroup.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pguo.mutation.Title(); ok {
		_spec.SetField(permissiongroup.FieldTitle, field.TypeString, value)
	}
	if value, ok := pguo.mutation.GroupType(); ok {
		_spec.SetField(permissiongroup.FieldGroupType, field.TypeEnum, value)
	}
	if value, ok := pguo.mutation.OrderID(); ok {
		_spec.SetField(permissiongroup.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := pguo.mutation.AddedOrderID(); ok {
		_spec.AddField(permissiongroup.FieldOrderID, field.TypeInt, value)
	}
	if pguo.mutation.GroupPermissionParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissiongroup.GroupPermissionParentTable,
			Columns: []string{permissiongroup.GroupPermissionParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.GroupPermissionParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissiongroup.GroupPermissionParentTable,
			Columns: []string{permissiongroup.GroupPermissionParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pguo.mutation.GroupPermissionChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.RemovedGroupPermissionChildrenIDs(); len(nodes) > 0 && !pguo.mutation.GroupPermissionChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.GroupPermissionChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.GroupPermissionChildrenTable,
			Columns: []string{permissiongroup.GroupPermissionChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permissiongroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pguo.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.RemovedPermissionEdgesIDs(); len(nodes) > 0 && !pguo.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pguo.mutation.PermissionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permissiongroup.PermissionEdgesTable,
			Columns: []string{permissiongroup.PermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PermissionGroup{config: pguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissiongroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

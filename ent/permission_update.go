// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/entitypermission"
	"trec/ent/permission"
	"trec/ent/permissiongroup"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableUpdatedAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PermissionUpdate) ClearUpdatedAt() *PermissionUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PermissionUpdate) SetDeletedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDeletedAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *PermissionUpdate) ClearDeletedAt() *PermissionUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetTitle sets the "title" field.
func (pu *PermissionUpdate) SetTitle(s string) *PermissionUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetGroupID sets the "group_id" field.
func (pu *PermissionUpdate) SetGroupID(u uuid.UUID) *PermissionUpdate {
	pu.mutation.SetGroupID(u)
	return pu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableGroupID(u *uuid.UUID) *PermissionUpdate {
	if u != nil {
		pu.SetGroupID(*u)
	}
	return pu
}

// ClearGroupID clears the value of the "group_id" field.
func (pu *PermissionUpdate) ClearGroupID() *PermissionUpdate {
	pu.mutation.ClearGroupID()
	return pu
}

// SetForOwner sets the "for_owner" field.
func (pu *PermissionUpdate) SetForOwner(b bool) *PermissionUpdate {
	pu.mutation.SetForOwner(b)
	return pu
}

// SetNillableForOwner sets the "for_owner" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableForOwner(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetForOwner(*b)
	}
	return pu
}

// SetForTeam sets the "for_team" field.
func (pu *PermissionUpdate) SetForTeam(b bool) *PermissionUpdate {
	pu.mutation.SetForTeam(b)
	return pu
}

// SetNillableForTeam sets the "for_team" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableForTeam(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetForTeam(*b)
	}
	return pu
}

// SetForAll sets the "for_all" field.
func (pu *PermissionUpdate) SetForAll(b bool) *PermissionUpdate {
	pu.mutation.SetForAll(b)
	return pu
}

// SetNillableForAll sets the "for_all" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableForAll(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetForAll(*b)
	}
	return pu
}

// SetOperationName sets the "operation_name" field.
func (pu *PermissionUpdate) SetOperationName(s string) *PermissionUpdate {
	pu.mutation.SetOperationName(s)
	return pu
}

// SetNillableOperationName sets the "operation_name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableOperationName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetOperationName(*s)
	}
	return pu
}

// ClearOperationName clears the value of the "operation_name" field.
func (pu *PermissionUpdate) ClearOperationName() *PermissionUpdate {
	pu.mutation.ClearOperationName()
	return pu
}

// SetParentID sets the "parent_id" field.
func (pu *PermissionUpdate) SetParentID(u uuid.UUID) *PermissionUpdate {
	pu.mutation.SetParentID(u)
	return pu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableParentID(u *uuid.UUID) *PermissionUpdate {
	if u != nil {
		pu.SetParentID(*u)
	}
	return pu
}

// ClearParentID clears the value of the "parent_id" field.
func (pu *PermissionUpdate) ClearParentID() *PermissionUpdate {
	pu.mutation.ClearParentID()
	return pu
}

// SetOrderID sets the "order_id" field.
func (pu *PermissionUpdate) SetOrderID(i int) *PermissionUpdate {
	pu.mutation.ResetOrderID()
	pu.mutation.SetOrderID(i)
	return pu
}

// AddOrderID adds i to the "order_id" field.
func (pu *PermissionUpdate) AddOrderID(i int) *PermissionUpdate {
	pu.mutation.AddOrderID(i)
	return pu
}

// SetGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID.
func (pu *PermissionUpdate) SetGroupPermissionEdgeID(id uuid.UUID) *PermissionUpdate {
	pu.mutation.SetGroupPermissionEdgeID(id)
	return pu
}

// SetNillableGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID if the given value is not nil.
func (pu *PermissionUpdate) SetNillableGroupPermissionEdgeID(id *uuid.UUID) *PermissionUpdate {
	if id != nil {
		pu = pu.SetGroupPermissionEdgeID(*id)
	}
	return pu
}

// SetGroupPermissionEdge sets the "group_permission_edge" edge to the PermissionGroup entity.
func (pu *PermissionUpdate) SetGroupPermissionEdge(p *PermissionGroup) *PermissionUpdate {
	return pu.SetGroupPermissionEdgeID(p.ID)
}

// AddUserPermissionEdgeIDs adds the "user_permission_edge" edge to the EntityPermission entity by IDs.
func (pu *PermissionUpdate) AddUserPermissionEdgeIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.AddUserPermissionEdgeIDs(ids...)
	return pu
}

// AddUserPermissionEdge adds the "user_permission_edge" edges to the EntityPermission entity.
func (pu *PermissionUpdate) AddUserPermissionEdge(e ...*EntityPermission) *PermissionUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.AddUserPermissionEdgeIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearGroupPermissionEdge clears the "group_permission_edge" edge to the PermissionGroup entity.
func (pu *PermissionUpdate) ClearGroupPermissionEdge() *PermissionUpdate {
	pu.mutation.ClearGroupPermissionEdge()
	return pu
}

// ClearUserPermissionEdge clears all "user_permission_edge" edges to the EntityPermission entity.
func (pu *PermissionUpdate) ClearUserPermissionEdge() *PermissionUpdate {
	pu.mutation.ClearUserPermissionEdge()
	return pu
}

// RemoveUserPermissionEdgeIDs removes the "user_permission_edge" edge to EntityPermission entities by IDs.
func (pu *PermissionUpdate) RemoveUserPermissionEdgeIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.RemoveUserPermissionEdgeIDs(ids...)
	return pu
}

// RemoveUserPermissionEdge removes "user_permission_edge" edges to EntityPermission entities.
func (pu *PermissionUpdate) RemoveUserPermissionEdge(e ...*EntityPermission) *PermissionUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.RemoveUserPermissionEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permission.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(permission.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(permission.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(permission.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.ForOwner(); ok {
		_spec.SetField(permission.FieldForOwner, field.TypeBool, value)
	}
	if value, ok := pu.mutation.ForTeam(); ok {
		_spec.SetField(permission.FieldForTeam, field.TypeBool, value)
	}
	if value, ok := pu.mutation.ForAll(); ok {
		_spec.SetField(permission.FieldForAll, field.TypeBool, value)
	}
	if value, ok := pu.mutation.OperationName(); ok {
		_spec.SetField(permission.FieldOperationName, field.TypeString, value)
	}
	if pu.mutation.OperationNameCleared() {
		_spec.ClearField(permission.FieldOperationName, field.TypeString)
	}
	if value, ok := pu.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeUUID, value)
	}
	if pu.mutation.ParentIDCleared() {
		_spec.ClearField(permission.FieldParentID, field.TypeUUID)
	}
	if value, ok := pu.mutation.OrderID(); ok {
		_spec.SetField(permission.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedOrderID(); ok {
		_spec.AddField(permission.FieldOrderID, field.TypeInt, value)
	}
	if pu.mutation.GroupPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.GroupPermissionEdgeTable,
			Columns: []string{permission.GroupPermissionEdgeColumn},
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
	if nodes := pu.mutation.GroupPermissionEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.GroupPermissionEdgeTable,
			Columns: []string{permission.GroupPermissionEdgeColumn},
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
	if pu.mutation.UserPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUserPermissionEdgeIDs(); len(nodes) > 0 && !pu.mutation.UserPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserPermissionEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableUpdatedAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PermissionUpdateOne) ClearUpdatedAt() *PermissionUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PermissionUpdateOne) SetDeletedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDeletedAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *PermissionUpdateOne) ClearDeletedAt() *PermissionUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetTitle sets the "title" field.
func (puo *PermissionUpdateOne) SetTitle(s string) *PermissionUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetGroupID sets the "group_id" field.
func (puo *PermissionUpdateOne) SetGroupID(u uuid.UUID) *PermissionUpdateOne {
	puo.mutation.SetGroupID(u)
	return puo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableGroupID(u *uuid.UUID) *PermissionUpdateOne {
	if u != nil {
		puo.SetGroupID(*u)
	}
	return puo
}

// ClearGroupID clears the value of the "group_id" field.
func (puo *PermissionUpdateOne) ClearGroupID() *PermissionUpdateOne {
	puo.mutation.ClearGroupID()
	return puo
}

// SetForOwner sets the "for_owner" field.
func (puo *PermissionUpdateOne) SetForOwner(b bool) *PermissionUpdateOne {
	puo.mutation.SetForOwner(b)
	return puo
}

// SetNillableForOwner sets the "for_owner" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableForOwner(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetForOwner(*b)
	}
	return puo
}

// SetForTeam sets the "for_team" field.
func (puo *PermissionUpdateOne) SetForTeam(b bool) *PermissionUpdateOne {
	puo.mutation.SetForTeam(b)
	return puo
}

// SetNillableForTeam sets the "for_team" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableForTeam(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetForTeam(*b)
	}
	return puo
}

// SetForAll sets the "for_all" field.
func (puo *PermissionUpdateOne) SetForAll(b bool) *PermissionUpdateOne {
	puo.mutation.SetForAll(b)
	return puo
}

// SetNillableForAll sets the "for_all" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableForAll(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetForAll(*b)
	}
	return puo
}

// SetOperationName sets the "operation_name" field.
func (puo *PermissionUpdateOne) SetOperationName(s string) *PermissionUpdateOne {
	puo.mutation.SetOperationName(s)
	return puo
}

// SetNillableOperationName sets the "operation_name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableOperationName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetOperationName(*s)
	}
	return puo
}

// ClearOperationName clears the value of the "operation_name" field.
func (puo *PermissionUpdateOne) ClearOperationName() *PermissionUpdateOne {
	puo.mutation.ClearOperationName()
	return puo
}

// SetParentID sets the "parent_id" field.
func (puo *PermissionUpdateOne) SetParentID(u uuid.UUID) *PermissionUpdateOne {
	puo.mutation.SetParentID(u)
	return puo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableParentID(u *uuid.UUID) *PermissionUpdateOne {
	if u != nil {
		puo.SetParentID(*u)
	}
	return puo
}

// ClearParentID clears the value of the "parent_id" field.
func (puo *PermissionUpdateOne) ClearParentID() *PermissionUpdateOne {
	puo.mutation.ClearParentID()
	return puo
}

// SetOrderID sets the "order_id" field.
func (puo *PermissionUpdateOne) SetOrderID(i int) *PermissionUpdateOne {
	puo.mutation.ResetOrderID()
	puo.mutation.SetOrderID(i)
	return puo
}

// AddOrderID adds i to the "order_id" field.
func (puo *PermissionUpdateOne) AddOrderID(i int) *PermissionUpdateOne {
	puo.mutation.AddOrderID(i)
	return puo
}

// SetGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID.
func (puo *PermissionUpdateOne) SetGroupPermissionEdgeID(id uuid.UUID) *PermissionUpdateOne {
	puo.mutation.SetGroupPermissionEdgeID(id)
	return puo
}

// SetNillableGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableGroupPermissionEdgeID(id *uuid.UUID) *PermissionUpdateOne {
	if id != nil {
		puo = puo.SetGroupPermissionEdgeID(*id)
	}
	return puo
}

// SetGroupPermissionEdge sets the "group_permission_edge" edge to the PermissionGroup entity.
func (puo *PermissionUpdateOne) SetGroupPermissionEdge(p *PermissionGroup) *PermissionUpdateOne {
	return puo.SetGroupPermissionEdgeID(p.ID)
}

// AddUserPermissionEdgeIDs adds the "user_permission_edge" edge to the EntityPermission entity by IDs.
func (puo *PermissionUpdateOne) AddUserPermissionEdgeIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.AddUserPermissionEdgeIDs(ids...)
	return puo
}

// AddUserPermissionEdge adds the "user_permission_edge" edges to the EntityPermission entity.
func (puo *PermissionUpdateOne) AddUserPermissionEdge(e ...*EntityPermission) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.AddUserPermissionEdgeIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearGroupPermissionEdge clears the "group_permission_edge" edge to the PermissionGroup entity.
func (puo *PermissionUpdateOne) ClearGroupPermissionEdge() *PermissionUpdateOne {
	puo.mutation.ClearGroupPermissionEdge()
	return puo
}

// ClearUserPermissionEdge clears all "user_permission_edge" edges to the EntityPermission entity.
func (puo *PermissionUpdateOne) ClearUserPermissionEdge() *PermissionUpdateOne {
	puo.mutation.ClearUserPermissionEdge()
	return puo
}

// RemoveUserPermissionEdgeIDs removes the "user_permission_edge" edge to EntityPermission entities by IDs.
func (puo *PermissionUpdateOne) RemoveUserPermissionEdgeIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.RemoveUserPermissionEdgeIDs(ids...)
	return puo
}

// RemoveUserPermissionEdge removes "user_permission_edge" edges to EntityPermission entities.
func (puo *PermissionUpdateOne) RemoveUserPermissionEdge(e ...*EntityPermission) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.RemoveUserPermissionEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Permission)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PermissionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permission.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(permission.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(permission.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(permission.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.ForOwner(); ok {
		_spec.SetField(permission.FieldForOwner, field.TypeBool, value)
	}
	if value, ok := puo.mutation.ForTeam(); ok {
		_spec.SetField(permission.FieldForTeam, field.TypeBool, value)
	}
	if value, ok := puo.mutation.ForAll(); ok {
		_spec.SetField(permission.FieldForAll, field.TypeBool, value)
	}
	if value, ok := puo.mutation.OperationName(); ok {
		_spec.SetField(permission.FieldOperationName, field.TypeString, value)
	}
	if puo.mutation.OperationNameCleared() {
		_spec.ClearField(permission.FieldOperationName, field.TypeString)
	}
	if value, ok := puo.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeUUID, value)
	}
	if puo.mutation.ParentIDCleared() {
		_spec.ClearField(permission.FieldParentID, field.TypeUUID)
	}
	if value, ok := puo.mutation.OrderID(); ok {
		_spec.SetField(permission.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedOrderID(); ok {
		_spec.AddField(permission.FieldOrderID, field.TypeInt, value)
	}
	if puo.mutation.GroupPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.GroupPermissionEdgeTable,
			Columns: []string{permission.GroupPermissionEdgeColumn},
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
	if nodes := puo.mutation.GroupPermissionEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.GroupPermissionEdgeTable,
			Columns: []string{permission.GroupPermissionEdgeColumn},
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
	if puo.mutation.UserPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUserPermissionEdgeIDs(); len(nodes) > 0 && !puo.mutation.UserPermissionEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserPermissionEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.UserPermissionEdgeTable,
			Columns: []string{permission.UserPermissionEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

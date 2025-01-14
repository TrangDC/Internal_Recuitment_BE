// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/entitypermission"
	"trec/ent/permission"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntityPermissionUpdate is the builder for updating EntityPermission entities.
type EntityPermissionUpdate struct {
	config
	hooks    []Hook
	mutation *EntityPermissionMutation
}

// Where appends a list predicates to the EntityPermissionUpdate builder.
func (epu *EntityPermissionUpdate) Where(ps ...predicate.EntityPermission) *EntityPermissionUpdate {
	epu.mutation.Where(ps...)
	return epu
}

// SetEntityID sets the "entity_id" field.
func (epu *EntityPermissionUpdate) SetEntityID(u uuid.UUID) *EntityPermissionUpdate {
	epu.mutation.SetEntityID(u)
	return epu
}

// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableEntityID(u *uuid.UUID) *EntityPermissionUpdate {
	if u != nil {
		epu.SetEntityID(*u)
	}
	return epu
}

// ClearEntityID clears the value of the "entity_id" field.
func (epu *EntityPermissionUpdate) ClearEntityID() *EntityPermissionUpdate {
	epu.mutation.ClearEntityID()
	return epu
}

// SetPermissionID sets the "permission_id" field.
func (epu *EntityPermissionUpdate) SetPermissionID(u uuid.UUID) *EntityPermissionUpdate {
	epu.mutation.SetPermissionID(u)
	return epu
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillablePermissionID(u *uuid.UUID) *EntityPermissionUpdate {
	if u != nil {
		epu.SetPermissionID(*u)
	}
	return epu
}

// ClearPermissionID clears the value of the "permission_id" field.
func (epu *EntityPermissionUpdate) ClearPermissionID() *EntityPermissionUpdate {
	epu.mutation.ClearPermissionID()
	return epu
}

// SetForOwner sets the "for_owner" field.
func (epu *EntityPermissionUpdate) SetForOwner(b bool) *EntityPermissionUpdate {
	epu.mutation.SetForOwner(b)
	return epu
}

// SetNillableForOwner sets the "for_owner" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableForOwner(b *bool) *EntityPermissionUpdate {
	if b != nil {
		epu.SetForOwner(*b)
	}
	return epu
}

// SetForTeam sets the "for_team" field.
func (epu *EntityPermissionUpdate) SetForTeam(b bool) *EntityPermissionUpdate {
	epu.mutation.SetForTeam(b)
	return epu
}

// SetNillableForTeam sets the "for_team" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableForTeam(b *bool) *EntityPermissionUpdate {
	if b != nil {
		epu.SetForTeam(*b)
	}
	return epu
}

// SetForAll sets the "for_all" field.
func (epu *EntityPermissionUpdate) SetForAll(b bool) *EntityPermissionUpdate {
	epu.mutation.SetForAll(b)
	return epu
}

// SetNillableForAll sets the "for_all" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableForAll(b *bool) *EntityPermissionUpdate {
	if b != nil {
		epu.SetForAll(*b)
	}
	return epu
}

// SetEntityType sets the "entity_type" field.
func (epu *EntityPermissionUpdate) SetEntityType(et entitypermission.EntityType) *EntityPermissionUpdate {
	epu.mutation.SetEntityType(et)
	return epu
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableEntityType(et *entitypermission.EntityType) *EntityPermissionUpdate {
	if et != nil {
		epu.SetEntityType(*et)
	}
	return epu
}

// ClearEntityType clears the value of the "entity_type" field.
func (epu *EntityPermissionUpdate) ClearEntityType() *EntityPermissionUpdate {
	epu.mutation.ClearEntityType()
	return epu
}

// SetUpdatedAt sets the "updated_at" field.
func (epu *EntityPermissionUpdate) SetUpdatedAt(t time.Time) *EntityPermissionUpdate {
	epu.mutation.SetUpdatedAt(t)
	return epu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableUpdatedAt(t *time.Time) *EntityPermissionUpdate {
	if t != nil {
		epu.SetUpdatedAt(*t)
	}
	return epu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (epu *EntityPermissionUpdate) ClearUpdatedAt() *EntityPermissionUpdate {
	epu.mutation.ClearUpdatedAt()
	return epu
}

// SetPermissionEdgesID sets the "permission_edges" edge to the Permission entity by ID.
func (epu *EntityPermissionUpdate) SetPermissionEdgesID(id uuid.UUID) *EntityPermissionUpdate {
	epu.mutation.SetPermissionEdgesID(id)
	return epu
}

// SetNillablePermissionEdgesID sets the "permission_edges" edge to the Permission entity by ID if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillablePermissionEdgesID(id *uuid.UUID) *EntityPermissionUpdate {
	if id != nil {
		epu = epu.SetPermissionEdgesID(*id)
	}
	return epu
}

// SetPermissionEdges sets the "permission_edges" edge to the Permission entity.
func (epu *EntityPermissionUpdate) SetPermissionEdges(p *Permission) *EntityPermissionUpdate {
	return epu.SetPermissionEdgesID(p.ID)
}

// SetUserEdgeID sets the "user_edge" edge to the User entity by ID.
func (epu *EntityPermissionUpdate) SetUserEdgeID(id uuid.UUID) *EntityPermissionUpdate {
	epu.mutation.SetUserEdgeID(id)
	return epu
}

// SetNillableUserEdgeID sets the "user_edge" edge to the User entity by ID if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableUserEdgeID(id *uuid.UUID) *EntityPermissionUpdate {
	if id != nil {
		epu = epu.SetUserEdgeID(*id)
	}
	return epu
}

// SetUserEdge sets the "user_edge" edge to the User entity.
func (epu *EntityPermissionUpdate) SetUserEdge(u *User) *EntityPermissionUpdate {
	return epu.SetUserEdgeID(u.ID)
}

// SetRoleEdgeID sets the "role_edge" edge to the Role entity by ID.
func (epu *EntityPermissionUpdate) SetRoleEdgeID(id uuid.UUID) *EntityPermissionUpdate {
	epu.mutation.SetRoleEdgeID(id)
	return epu
}

// SetNillableRoleEdgeID sets the "role_edge" edge to the Role entity by ID if the given value is not nil.
func (epu *EntityPermissionUpdate) SetNillableRoleEdgeID(id *uuid.UUID) *EntityPermissionUpdate {
	if id != nil {
		epu = epu.SetRoleEdgeID(*id)
	}
	return epu
}

// SetRoleEdge sets the "role_edge" edge to the Role entity.
func (epu *EntityPermissionUpdate) SetRoleEdge(r *Role) *EntityPermissionUpdate {
	return epu.SetRoleEdgeID(r.ID)
}

// Mutation returns the EntityPermissionMutation object of the builder.
func (epu *EntityPermissionUpdate) Mutation() *EntityPermissionMutation {
	return epu.mutation
}

// ClearPermissionEdges clears the "permission_edges" edge to the Permission entity.
func (epu *EntityPermissionUpdate) ClearPermissionEdges() *EntityPermissionUpdate {
	epu.mutation.ClearPermissionEdges()
	return epu
}

// ClearUserEdge clears the "user_edge" edge to the User entity.
func (epu *EntityPermissionUpdate) ClearUserEdge() *EntityPermissionUpdate {
	epu.mutation.ClearUserEdge()
	return epu
}

// ClearRoleEdge clears the "role_edge" edge to the Role entity.
func (epu *EntityPermissionUpdate) ClearRoleEdge() *EntityPermissionUpdate {
	epu.mutation.ClearRoleEdge()
	return epu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (epu *EntityPermissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(epu.hooks) == 0 {
		if err = epu.check(); err != nil {
			return 0, err
		}
		affected, err = epu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = epu.check(); err != nil {
				return 0, err
			}
			epu.mutation = mutation
			affected, err = epu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(epu.hooks) - 1; i >= 0; i-- {
			if epu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = epu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (epu *EntityPermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := epu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (epu *EntityPermissionUpdate) Exec(ctx context.Context) error {
	_, err := epu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epu *EntityPermissionUpdate) ExecX(ctx context.Context) {
	if err := epu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epu *EntityPermissionUpdate) check() error {
	if v, ok := epu.mutation.EntityType(); ok {
		if err := entitypermission.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`ent: validator failed for field "EntityPermission.entity_type": %w`, err)}
		}
	}
	return nil
}

func (epu *EntityPermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitypermission.Table,
			Columns: entitypermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entitypermission.FieldID,
			},
		},
	}
	if ps := epu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epu.mutation.ForOwner(); ok {
		_spec.SetField(entitypermission.FieldForOwner, field.TypeBool, value)
	}
	if value, ok := epu.mutation.ForTeam(); ok {
		_spec.SetField(entitypermission.FieldForTeam, field.TypeBool, value)
	}
	if value, ok := epu.mutation.ForAll(); ok {
		_spec.SetField(entitypermission.FieldForAll, field.TypeBool, value)
	}
	if value, ok := epu.mutation.EntityType(); ok {
		_spec.SetField(entitypermission.FieldEntityType, field.TypeEnum, value)
	}
	if epu.mutation.EntityTypeCleared() {
		_spec.ClearField(entitypermission.FieldEntityType, field.TypeEnum)
	}
	if value, ok := epu.mutation.UpdatedAt(); ok {
		_spec.SetField(entitypermission.FieldUpdatedAt, field.TypeTime, value)
	}
	if epu.mutation.UpdatedAtCleared() {
		_spec.ClearField(entitypermission.FieldUpdatedAt, field.TypeTime)
	}
	if epu.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.PermissionEdgesTable,
			Columns: []string{entitypermission.PermissionEdgesColumn},
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
	if nodes := epu.mutation.PermissionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.PermissionEdgesTable,
			Columns: []string{entitypermission.PermissionEdgesColumn},
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
	if epu.mutation.UserEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.UserEdgeTable,
			Columns: []string{entitypermission.UserEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epu.mutation.UserEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.UserEdgeTable,
			Columns: []string{entitypermission.UserEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epu.mutation.RoleEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.RoleEdgeTable,
			Columns: []string{entitypermission.RoleEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epu.mutation.RoleEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.RoleEdgeTable,
			Columns: []string{entitypermission.RoleEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, epu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitypermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EntityPermissionUpdateOne is the builder for updating a single EntityPermission entity.
type EntityPermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityPermissionMutation
}

// SetEntityID sets the "entity_id" field.
func (epuo *EntityPermissionUpdateOne) SetEntityID(u uuid.UUID) *EntityPermissionUpdateOne {
	epuo.mutation.SetEntityID(u)
	return epuo
}

// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableEntityID(u *uuid.UUID) *EntityPermissionUpdateOne {
	if u != nil {
		epuo.SetEntityID(*u)
	}
	return epuo
}

// ClearEntityID clears the value of the "entity_id" field.
func (epuo *EntityPermissionUpdateOne) ClearEntityID() *EntityPermissionUpdateOne {
	epuo.mutation.ClearEntityID()
	return epuo
}

// SetPermissionID sets the "permission_id" field.
func (epuo *EntityPermissionUpdateOne) SetPermissionID(u uuid.UUID) *EntityPermissionUpdateOne {
	epuo.mutation.SetPermissionID(u)
	return epuo
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillablePermissionID(u *uuid.UUID) *EntityPermissionUpdateOne {
	if u != nil {
		epuo.SetPermissionID(*u)
	}
	return epuo
}

// ClearPermissionID clears the value of the "permission_id" field.
func (epuo *EntityPermissionUpdateOne) ClearPermissionID() *EntityPermissionUpdateOne {
	epuo.mutation.ClearPermissionID()
	return epuo
}

// SetForOwner sets the "for_owner" field.
func (epuo *EntityPermissionUpdateOne) SetForOwner(b bool) *EntityPermissionUpdateOne {
	epuo.mutation.SetForOwner(b)
	return epuo
}

// SetNillableForOwner sets the "for_owner" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableForOwner(b *bool) *EntityPermissionUpdateOne {
	if b != nil {
		epuo.SetForOwner(*b)
	}
	return epuo
}

// SetForTeam sets the "for_team" field.
func (epuo *EntityPermissionUpdateOne) SetForTeam(b bool) *EntityPermissionUpdateOne {
	epuo.mutation.SetForTeam(b)
	return epuo
}

// SetNillableForTeam sets the "for_team" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableForTeam(b *bool) *EntityPermissionUpdateOne {
	if b != nil {
		epuo.SetForTeam(*b)
	}
	return epuo
}

// SetForAll sets the "for_all" field.
func (epuo *EntityPermissionUpdateOne) SetForAll(b bool) *EntityPermissionUpdateOne {
	epuo.mutation.SetForAll(b)
	return epuo
}

// SetNillableForAll sets the "for_all" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableForAll(b *bool) *EntityPermissionUpdateOne {
	if b != nil {
		epuo.SetForAll(*b)
	}
	return epuo
}

// SetEntityType sets the "entity_type" field.
func (epuo *EntityPermissionUpdateOne) SetEntityType(et entitypermission.EntityType) *EntityPermissionUpdateOne {
	epuo.mutation.SetEntityType(et)
	return epuo
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableEntityType(et *entitypermission.EntityType) *EntityPermissionUpdateOne {
	if et != nil {
		epuo.SetEntityType(*et)
	}
	return epuo
}

// ClearEntityType clears the value of the "entity_type" field.
func (epuo *EntityPermissionUpdateOne) ClearEntityType() *EntityPermissionUpdateOne {
	epuo.mutation.ClearEntityType()
	return epuo
}

// SetUpdatedAt sets the "updated_at" field.
func (epuo *EntityPermissionUpdateOne) SetUpdatedAt(t time.Time) *EntityPermissionUpdateOne {
	epuo.mutation.SetUpdatedAt(t)
	return epuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableUpdatedAt(t *time.Time) *EntityPermissionUpdateOne {
	if t != nil {
		epuo.SetUpdatedAt(*t)
	}
	return epuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (epuo *EntityPermissionUpdateOne) ClearUpdatedAt() *EntityPermissionUpdateOne {
	epuo.mutation.ClearUpdatedAt()
	return epuo
}

// SetPermissionEdgesID sets the "permission_edges" edge to the Permission entity by ID.
func (epuo *EntityPermissionUpdateOne) SetPermissionEdgesID(id uuid.UUID) *EntityPermissionUpdateOne {
	epuo.mutation.SetPermissionEdgesID(id)
	return epuo
}

// SetNillablePermissionEdgesID sets the "permission_edges" edge to the Permission entity by ID if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillablePermissionEdgesID(id *uuid.UUID) *EntityPermissionUpdateOne {
	if id != nil {
		epuo = epuo.SetPermissionEdgesID(*id)
	}
	return epuo
}

// SetPermissionEdges sets the "permission_edges" edge to the Permission entity.
func (epuo *EntityPermissionUpdateOne) SetPermissionEdges(p *Permission) *EntityPermissionUpdateOne {
	return epuo.SetPermissionEdgesID(p.ID)
}

// SetUserEdgeID sets the "user_edge" edge to the User entity by ID.
func (epuo *EntityPermissionUpdateOne) SetUserEdgeID(id uuid.UUID) *EntityPermissionUpdateOne {
	epuo.mutation.SetUserEdgeID(id)
	return epuo
}

// SetNillableUserEdgeID sets the "user_edge" edge to the User entity by ID if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableUserEdgeID(id *uuid.UUID) *EntityPermissionUpdateOne {
	if id != nil {
		epuo = epuo.SetUserEdgeID(*id)
	}
	return epuo
}

// SetUserEdge sets the "user_edge" edge to the User entity.
func (epuo *EntityPermissionUpdateOne) SetUserEdge(u *User) *EntityPermissionUpdateOne {
	return epuo.SetUserEdgeID(u.ID)
}

// SetRoleEdgeID sets the "role_edge" edge to the Role entity by ID.
func (epuo *EntityPermissionUpdateOne) SetRoleEdgeID(id uuid.UUID) *EntityPermissionUpdateOne {
	epuo.mutation.SetRoleEdgeID(id)
	return epuo
}

// SetNillableRoleEdgeID sets the "role_edge" edge to the Role entity by ID if the given value is not nil.
func (epuo *EntityPermissionUpdateOne) SetNillableRoleEdgeID(id *uuid.UUID) *EntityPermissionUpdateOne {
	if id != nil {
		epuo = epuo.SetRoleEdgeID(*id)
	}
	return epuo
}

// SetRoleEdge sets the "role_edge" edge to the Role entity.
func (epuo *EntityPermissionUpdateOne) SetRoleEdge(r *Role) *EntityPermissionUpdateOne {
	return epuo.SetRoleEdgeID(r.ID)
}

// Mutation returns the EntityPermissionMutation object of the builder.
func (epuo *EntityPermissionUpdateOne) Mutation() *EntityPermissionMutation {
	return epuo.mutation
}

// ClearPermissionEdges clears the "permission_edges" edge to the Permission entity.
func (epuo *EntityPermissionUpdateOne) ClearPermissionEdges() *EntityPermissionUpdateOne {
	epuo.mutation.ClearPermissionEdges()
	return epuo
}

// ClearUserEdge clears the "user_edge" edge to the User entity.
func (epuo *EntityPermissionUpdateOne) ClearUserEdge() *EntityPermissionUpdateOne {
	epuo.mutation.ClearUserEdge()
	return epuo
}

// ClearRoleEdge clears the "role_edge" edge to the Role entity.
func (epuo *EntityPermissionUpdateOne) ClearRoleEdge() *EntityPermissionUpdateOne {
	epuo.mutation.ClearRoleEdge()
	return epuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (epuo *EntityPermissionUpdateOne) Select(field string, fields ...string) *EntityPermissionUpdateOne {
	epuo.fields = append([]string{field}, fields...)
	return epuo
}

// Save executes the query and returns the updated EntityPermission entity.
func (epuo *EntityPermissionUpdateOne) Save(ctx context.Context) (*EntityPermission, error) {
	var (
		err  error
		node *EntityPermission
	)
	if len(epuo.hooks) == 0 {
		if err = epuo.check(); err != nil {
			return nil, err
		}
		node, err = epuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = epuo.check(); err != nil {
				return nil, err
			}
			epuo.mutation = mutation
			node, err = epuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(epuo.hooks) - 1; i >= 0; i-- {
			if epuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = epuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, epuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EntityPermission)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EntityPermissionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (epuo *EntityPermissionUpdateOne) SaveX(ctx context.Context) *EntityPermission {
	node, err := epuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (epuo *EntityPermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := epuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epuo *EntityPermissionUpdateOne) ExecX(ctx context.Context) {
	if err := epuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epuo *EntityPermissionUpdateOne) check() error {
	if v, ok := epuo.mutation.EntityType(); ok {
		if err := entitypermission.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`ent: validator failed for field "EntityPermission.entity_type": %w`, err)}
		}
	}
	return nil
}

func (epuo *EntityPermissionUpdateOne) sqlSave(ctx context.Context) (_node *EntityPermission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitypermission.Table,
			Columns: entitypermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entitypermission.FieldID,
			},
		},
	}
	id, ok := epuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntityPermission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := epuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitypermission.FieldID)
		for _, f := range fields {
			if !entitypermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entitypermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := epuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := epuo.mutation.ForOwner(); ok {
		_spec.SetField(entitypermission.FieldForOwner, field.TypeBool, value)
	}
	if value, ok := epuo.mutation.ForTeam(); ok {
		_spec.SetField(entitypermission.FieldForTeam, field.TypeBool, value)
	}
	if value, ok := epuo.mutation.ForAll(); ok {
		_spec.SetField(entitypermission.FieldForAll, field.TypeBool, value)
	}
	if value, ok := epuo.mutation.EntityType(); ok {
		_spec.SetField(entitypermission.FieldEntityType, field.TypeEnum, value)
	}
	if epuo.mutation.EntityTypeCleared() {
		_spec.ClearField(entitypermission.FieldEntityType, field.TypeEnum)
	}
	if value, ok := epuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entitypermission.FieldUpdatedAt, field.TypeTime, value)
	}
	if epuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(entitypermission.FieldUpdatedAt, field.TypeTime)
	}
	if epuo.mutation.PermissionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.PermissionEdgesTable,
			Columns: []string{entitypermission.PermissionEdgesColumn},
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
	if nodes := epuo.mutation.PermissionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.PermissionEdgesTable,
			Columns: []string{entitypermission.PermissionEdgesColumn},
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
	if epuo.mutation.UserEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.UserEdgeTable,
			Columns: []string{entitypermission.UserEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epuo.mutation.UserEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.UserEdgeTable,
			Columns: []string{entitypermission.UserEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if epuo.mutation.RoleEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.RoleEdgeTable,
			Columns: []string{entitypermission.RoleEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := epuo.mutation.RoleEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entitypermission.RoleEdgeTable,
			Columns: []string{entitypermission.RoleEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntityPermission{config: epuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, epuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitypermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

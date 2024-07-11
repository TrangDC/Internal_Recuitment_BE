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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PermissionCreate) SetDeletedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDeletedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PermissionCreate) SetTitle(s string) *PermissionCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetGroupID sets the "group_id" field.
func (pc *PermissionCreate) SetGroupID(u uuid.UUID) *PermissionCreate {
	pc.mutation.SetGroupID(u)
	return pc
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableGroupID(u *uuid.UUID) *PermissionCreate {
	if u != nil {
		pc.SetGroupID(*u)
	}
	return pc
}

// SetForOwner sets the "for_owner" field.
func (pc *PermissionCreate) SetForOwner(b bool) *PermissionCreate {
	pc.mutation.SetForOwner(b)
	return pc
}

// SetNillableForOwner sets the "for_owner" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableForOwner(b *bool) *PermissionCreate {
	if b != nil {
		pc.SetForOwner(*b)
	}
	return pc
}

// SetForTeam sets the "for_team" field.
func (pc *PermissionCreate) SetForTeam(b bool) *PermissionCreate {
	pc.mutation.SetForTeam(b)
	return pc
}

// SetNillableForTeam sets the "for_team" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableForTeam(b *bool) *PermissionCreate {
	if b != nil {
		pc.SetForTeam(*b)
	}
	return pc
}

// SetForAll sets the "for_all" field.
func (pc *PermissionCreate) SetForAll(b bool) *PermissionCreate {
	pc.mutation.SetForAll(b)
	return pc
}

// SetNillableForAll sets the "for_all" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableForAll(b *bool) *PermissionCreate {
	if b != nil {
		pc.SetForAll(*b)
	}
	return pc
}

// SetOperationName sets the "operation_name" field.
func (pc *PermissionCreate) SetOperationName(s string) *PermissionCreate {
	pc.mutation.SetOperationName(s)
	return pc
}

// SetNillableOperationName sets the "operation_name" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableOperationName(s *string) *PermissionCreate {
	if s != nil {
		pc.SetOperationName(*s)
	}
	return pc
}

// SetParentID sets the "parent_id" field.
func (pc *PermissionCreate) SetParentID(u uuid.UUID) *PermissionCreate {
	pc.mutation.SetParentID(u)
	return pc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableParentID(u *uuid.UUID) *PermissionCreate {
	if u != nil {
		pc.SetParentID(*u)
	}
	return pc
}

// SetOrderID sets the "order_id" field.
func (pc *PermissionCreate) SetOrderID(i int) *PermissionCreate {
	pc.mutation.SetOrderID(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(u uuid.UUID) *PermissionCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID.
func (pc *PermissionCreate) SetGroupPermissionEdgeID(id uuid.UUID) *PermissionCreate {
	pc.mutation.SetGroupPermissionEdgeID(id)
	return pc
}

// SetNillableGroupPermissionEdgeID sets the "group_permission_edge" edge to the PermissionGroup entity by ID if the given value is not nil.
func (pc *PermissionCreate) SetNillableGroupPermissionEdgeID(id *uuid.UUID) *PermissionCreate {
	if id != nil {
		pc = pc.SetGroupPermissionEdgeID(*id)
	}
	return pc
}

// SetGroupPermissionEdge sets the "group_permission_edge" edge to the PermissionGroup entity.
func (pc *PermissionCreate) SetGroupPermissionEdge(p *PermissionGroup) *PermissionCreate {
	return pc.SetGroupPermissionEdgeID(p.ID)
}

// AddUserPermissionEdgeIDs adds the "user_permission_edge" edge to the EntityPermission entity by IDs.
func (pc *PermissionCreate) AddUserPermissionEdgeIDs(ids ...uuid.UUID) *PermissionCreate {
	pc.mutation.AddUserPermissionEdgeIDs(ids...)
	return pc
}

// AddUserPermissionEdge adds the "user_permission_edge" edges to the EntityPermission entity.
func (pc *PermissionCreate) AddUserPermissionEdge(e ...*EntityPermission) *PermissionCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddUserPermissionEdgeIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.ForOwner(); !ok {
		v := permission.DefaultForOwner
		pc.mutation.SetForOwner(v)
	}
	if _, ok := pc.mutation.ForTeam(); !ok {
		v := permission.DefaultForTeam
		pc.mutation.SetForTeam(v)
	}
	if _, ok := pc.mutation.ForAll(); !ok {
		v := permission.DefaultForAll
		pc.mutation.SetForAll(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Permission.created_at"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Permission.title"`)}
	}
	if _, ok := pc.mutation.ForOwner(); !ok {
		return &ValidationError{Name: "for_owner", err: errors.New(`ent: missing required field "Permission.for_owner"`)}
	}
	if _, ok := pc.mutation.ForTeam(); !ok {
		return &ValidationError{Name: "for_team", err: errors.New(`ent: missing required field "Permission.for_team"`)}
	}
	if _, ok := pc.mutation.ForAll(); !ok {
		return &ValidationError{Name: "for_all", err: errors.New(`ent: missing required field "Permission.for_all"`)}
	}
	if _, ok := pc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Permission.order_id"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: permission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permission.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(permission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(permission.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(permission.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.ForOwner(); ok {
		_spec.SetField(permission.FieldForOwner, field.TypeBool, value)
		_node.ForOwner = value
	}
	if value, ok := pc.mutation.ForTeam(); ok {
		_spec.SetField(permission.FieldForTeam, field.TypeBool, value)
		_node.ForTeam = value
	}
	if value, ok := pc.mutation.ForAll(); ok {
		_spec.SetField(permission.FieldForAll, field.TypeBool, value)
		_node.ForAll = value
	}
	if value, ok := pc.mutation.OperationName(); ok {
		_spec.SetField(permission.FieldOperationName, field.TypeString, value)
		_node.OperationName = value
	}
	if value, ok := pc.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeUUID, value)
		_node.ParentID = value
	}
	if value, ok := pc.mutation.OrderID(); ok {
		_spec.SetField(permission.FieldOrderID, field.TypeInt, value)
		_node.OrderID = value
	}
	if nodes := pc.mutation.GroupPermissionEdgeIDs(); len(nodes) > 0 {
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
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UserPermissionEdgeIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
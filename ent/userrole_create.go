// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/role"
	"trec/ent/user"
	"trec/ent/userrole"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (urc *UserRoleCreate) SetCreatedAt(t time.Time) *UserRoleCreate {
	urc.mutation.SetCreatedAt(t)
	return urc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableCreatedAt(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetCreatedAt(*t)
	}
	return urc
}

// SetUpdatedAt sets the "updated_at" field.
func (urc *UserRoleCreate) SetUpdatedAt(t time.Time) *UserRoleCreate {
	urc.mutation.SetUpdatedAt(t)
	return urc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableUpdatedAt(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetUpdatedAt(*t)
	}
	return urc
}

// SetDeletedAt sets the "deleted_at" field.
func (urc *UserRoleCreate) SetDeletedAt(t time.Time) *UserRoleCreate {
	urc.mutation.SetDeletedAt(t)
	return urc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableDeletedAt(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetDeletedAt(*t)
	}
	return urc
}

// SetRoleID sets the "role_id" field.
func (urc *UserRoleCreate) SetRoleID(u uuid.UUID) *UserRoleCreate {
	urc.mutation.SetRoleID(u)
	return urc
}

// SetUserID sets the "user_id" field.
func (urc *UserRoleCreate) SetUserID(u uuid.UUID) *UserRoleCreate {
	urc.mutation.SetUserID(u)
	return urc
}

// SetID sets the "id" field.
func (urc *UserRoleCreate) SetID(u uuid.UUID) *UserRoleCreate {
	urc.mutation.SetID(u)
	return urc
}

// SetUserEdgeID sets the "user_edge" edge to the User entity by ID.
func (urc *UserRoleCreate) SetUserEdgeID(id uuid.UUID) *UserRoleCreate {
	urc.mutation.SetUserEdgeID(id)
	return urc
}

// SetUserEdge sets the "user_edge" edge to the User entity.
func (urc *UserRoleCreate) SetUserEdge(u *User) *UserRoleCreate {
	return urc.SetUserEdgeID(u.ID)
}

// SetRoleEdgeID sets the "role_edge" edge to the Role entity by ID.
func (urc *UserRoleCreate) SetRoleEdgeID(id uuid.UUID) *UserRoleCreate {
	urc.mutation.SetRoleEdgeID(id)
	return urc
}

// SetRoleEdge sets the "role_edge" edge to the Role entity.
func (urc *UserRoleCreate) SetRoleEdge(r *Role) *UserRoleCreate {
	return urc.SetRoleEdgeID(r.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	var (
		err  error
		node *UserRole
	)
	urc.defaults()
	if len(urc.hooks) == 0 {
		if err = urc.check(); err != nil {
			return nil, err
		}
		node, err = urc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = urc.check(); err != nil {
				return nil, err
			}
			urc.mutation = mutation
			if node, err = urc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(urc.hooks) - 1; i >= 0; i-- {
			if urc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = urc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, urc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserRole)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserRoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRoleCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRoleCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urc *UserRoleCreate) defaults() {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		v := userrole.DefaultCreatedAt()
		urc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserRole.created_at"`)}
	}
	if _, ok := urc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "UserRole.role_id"`)}
	}
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserRole.user_id"`)}
	}
	if _, ok := urc.mutation.UserEdgeID(); !ok {
		return &ValidationError{Name: "user_edge", err: errors.New(`ent: missing required edge "UserRole.user_edge"`)}
	}
	if _, ok := urc.mutation.RoleEdgeID(); !ok {
		return &ValidationError{Name: "role_edge", err: errors.New(`ent: missing required edge "UserRole.role_edge"`)}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
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

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		}
	)
	if id, ok := urc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := urc.mutation.CreatedAt(); ok {
		_spec.SetField(userrole.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := urc.mutation.UpdatedAt(); ok {
		_spec.SetField(userrole.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := urc.mutation.DeletedAt(); ok {
		_spec.SetField(userrole.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := urc.mutation.UserEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserEdgeTable,
			Columns: []string{userrole.UserEdgeColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := urc.mutation.RoleEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleEdgeTable,
			Columns: []string{userrole.RoleEdgeColumn},
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
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserRoleCreateBulk is the builder for creating many UserRole entities in bulk.
type UserRoleCreateBulk struct {
	config
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}

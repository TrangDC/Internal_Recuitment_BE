// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/team"
	"trec/ent/teammanager"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TeamCreate) SetCreatedAt(t time.Time) *TeamCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TeamCreate) SetNillableCreatedAt(t *time.Time) *TeamCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TeamCreate) SetUpdatedAt(t time.Time) *TeamCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TeamCreate) SetNillableUpdatedAt(t *time.Time) *TeamCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TeamCreate) SetDeletedAt(t time.Time) *TeamCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TeamCreate) SetNillableDeletedAt(t *time.Time) *TeamCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetSlug sets the "slug" field.
func (tc *TeamCreate) SetSlug(s string) *TeamCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetName sets the "name" field.
func (tc *TeamCreate) SetName(s string) *TeamCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TeamCreate) SetID(u uuid.UUID) *TeamCreate {
	tc.mutation.SetID(u)
	return tc
}

// AddUserEdgeIDs adds the "user_edges" edge to the User entity by IDs.
func (tc *TeamCreate) AddUserEdgeIDs(ids ...uuid.UUID) *TeamCreate {
	tc.mutation.AddUserEdgeIDs(ids...)
	return tc
}

// AddUserEdges adds the "user_edges" edges to the User entity.
func (tc *TeamCreate) AddUserEdges(u ...*User) *TeamCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserEdgeIDs(ids...)
}

// AddUserTeamIDs adds the "user_teams" edge to the TeamManager entity by IDs.
func (tc *TeamCreate) AddUserTeamIDs(ids ...uuid.UUID) *TeamCreate {
	tc.mutation.AddUserTeamIDs(ids...)
	return tc
}

// AddUserTeams adds the "user_teams" edges to the TeamManager entity.
func (tc *TeamCreate) AddUserTeams(t ...*TeamManager) *TeamCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddUserTeamIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	var (
		err  error
		node *Team
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Team)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TeamMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeamCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeamCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TeamCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := team.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Team.created_at"`)}
	}
	if _, ok := tc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Team.slug"`)}
	}
	if v, ok := tc.mutation.Slug(); ok {
		if err := team.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Team.slug": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Team.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := team.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Team.name": %w`, err)}
		}
	}
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: team.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: team.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(team.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(team.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(team.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(team.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := tc.mutation.UserEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.UserEdgesTable,
			Columns: team.UserEdgesPrimaryKey,
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
		createE := &TeamManagerCreate{config: tc.config, mutation: newTeamManagerMutation(tc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UserTeamsTable,
			Columns: []string{team.UserTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teammanager.FieldID,
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

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	builders []*TeamCreate
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeamCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeamCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

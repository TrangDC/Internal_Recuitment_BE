// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/hiringteammanager"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringTeamUpdate is the builder for updating HiringTeam entities.
type HiringTeamUpdate struct {
	config
	hooks    []Hook
	mutation *HiringTeamMutation
}

// Where appends a list predicates to the HiringTeamUpdate builder.
func (htu *HiringTeamUpdate) Where(ps ...predicate.HiringTeam) *HiringTeamUpdate {
	htu.mutation.Where(ps...)
	return htu
}

// SetUpdatedAt sets the "updated_at" field.
func (htu *HiringTeamUpdate) SetUpdatedAt(t time.Time) *HiringTeamUpdate {
	htu.mutation.SetUpdatedAt(t)
	return htu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (htu *HiringTeamUpdate) SetNillableUpdatedAt(t *time.Time) *HiringTeamUpdate {
	if t != nil {
		htu.SetUpdatedAt(*t)
	}
	return htu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (htu *HiringTeamUpdate) ClearUpdatedAt() *HiringTeamUpdate {
	htu.mutation.ClearUpdatedAt()
	return htu
}

// SetDeletedAt sets the "deleted_at" field.
func (htu *HiringTeamUpdate) SetDeletedAt(t time.Time) *HiringTeamUpdate {
	htu.mutation.SetDeletedAt(t)
	return htu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (htu *HiringTeamUpdate) SetNillableDeletedAt(t *time.Time) *HiringTeamUpdate {
	if t != nil {
		htu.SetDeletedAt(*t)
	}
	return htu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (htu *HiringTeamUpdate) ClearDeletedAt() *HiringTeamUpdate {
	htu.mutation.ClearDeletedAt()
	return htu
}

// SetSlug sets the "slug" field.
func (htu *HiringTeamUpdate) SetSlug(s string) *HiringTeamUpdate {
	htu.mutation.SetSlug(s)
	return htu
}

// SetName sets the "name" field.
func (htu *HiringTeamUpdate) SetName(s string) *HiringTeamUpdate {
	htu.mutation.SetName(s)
	return htu
}

// AddUserEdgeIDs adds the "user_edges" edge to the User entity by IDs.
func (htu *HiringTeamUpdate) AddUserEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.AddUserEdgeIDs(ids...)
	return htu
}

// AddUserEdges adds the "user_edges" edges to the User entity.
func (htu *HiringTeamUpdate) AddUserEdges(u ...*User) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return htu.AddUserEdgeIDs(ids...)
}

// AddHiringTeamJobEdgeIDs adds the "hiring_team_job_edges" edge to the HiringJob entity by IDs.
func (htu *HiringTeamUpdate) AddHiringTeamJobEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.AddHiringTeamJobEdgeIDs(ids...)
	return htu
}

// AddHiringTeamJobEdges adds the "hiring_team_job_edges" edges to the HiringJob entity.
func (htu *HiringTeamUpdate) AddHiringTeamJobEdges(h ...*HiringJob) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htu.AddHiringTeamJobEdgeIDs(ids...)
}

// AddUserHiringTeamIDs adds the "user_hiring_teams" edge to the HiringTeamManager entity by IDs.
func (htu *HiringTeamUpdate) AddUserHiringTeamIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.AddUserHiringTeamIDs(ids...)
	return htu
}

// AddUserHiringTeams adds the "user_hiring_teams" edges to the HiringTeamManager entity.
func (htu *HiringTeamUpdate) AddUserHiringTeams(h ...*HiringTeamManager) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htu.AddUserHiringTeamIDs(ids...)
}

// Mutation returns the HiringTeamMutation object of the builder.
func (htu *HiringTeamUpdate) Mutation() *HiringTeamMutation {
	return htu.mutation
}

// ClearUserEdges clears all "user_edges" edges to the User entity.
func (htu *HiringTeamUpdate) ClearUserEdges() *HiringTeamUpdate {
	htu.mutation.ClearUserEdges()
	return htu
}

// RemoveUserEdgeIDs removes the "user_edges" edge to User entities by IDs.
func (htu *HiringTeamUpdate) RemoveUserEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.RemoveUserEdgeIDs(ids...)
	return htu
}

// RemoveUserEdges removes "user_edges" edges to User entities.
func (htu *HiringTeamUpdate) RemoveUserEdges(u ...*User) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return htu.RemoveUserEdgeIDs(ids...)
}

// ClearHiringTeamJobEdges clears all "hiring_team_job_edges" edges to the HiringJob entity.
func (htu *HiringTeamUpdate) ClearHiringTeamJobEdges() *HiringTeamUpdate {
	htu.mutation.ClearHiringTeamJobEdges()
	return htu
}

// RemoveHiringTeamJobEdgeIDs removes the "hiring_team_job_edges" edge to HiringJob entities by IDs.
func (htu *HiringTeamUpdate) RemoveHiringTeamJobEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.RemoveHiringTeamJobEdgeIDs(ids...)
	return htu
}

// RemoveHiringTeamJobEdges removes "hiring_team_job_edges" edges to HiringJob entities.
func (htu *HiringTeamUpdate) RemoveHiringTeamJobEdges(h ...*HiringJob) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htu.RemoveHiringTeamJobEdgeIDs(ids...)
}

// ClearUserHiringTeams clears all "user_hiring_teams" edges to the HiringTeamManager entity.
func (htu *HiringTeamUpdate) ClearUserHiringTeams() *HiringTeamUpdate {
	htu.mutation.ClearUserHiringTeams()
	return htu
}

// RemoveUserHiringTeamIDs removes the "user_hiring_teams" edge to HiringTeamManager entities by IDs.
func (htu *HiringTeamUpdate) RemoveUserHiringTeamIDs(ids ...uuid.UUID) *HiringTeamUpdate {
	htu.mutation.RemoveUserHiringTeamIDs(ids...)
	return htu
}

// RemoveUserHiringTeams removes "user_hiring_teams" edges to HiringTeamManager entities.
func (htu *HiringTeamUpdate) RemoveUserHiringTeams(h ...*HiringTeamManager) *HiringTeamUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htu.RemoveUserHiringTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (htu *HiringTeamUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(htu.hooks) == 0 {
		if err = htu.check(); err != nil {
			return 0, err
		}
		affected, err = htu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringTeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = htu.check(); err != nil {
				return 0, err
			}
			htu.mutation = mutation
			affected, err = htu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(htu.hooks) - 1; i >= 0; i-- {
			if htu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = htu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, htu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (htu *HiringTeamUpdate) SaveX(ctx context.Context) int {
	affected, err := htu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (htu *HiringTeamUpdate) Exec(ctx context.Context) error {
	_, err := htu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htu *HiringTeamUpdate) ExecX(ctx context.Context) {
	if err := htu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (htu *HiringTeamUpdate) check() error {
	if v, ok := htu.mutation.Slug(); ok {
		if err := hiringteam.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "HiringTeam.slug": %w`, err)}
		}
	}
	if v, ok := htu.mutation.Name(); ok {
		if err := hiringteam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HiringTeam.name": %w`, err)}
		}
	}
	return nil
}

func (htu *HiringTeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hiringteam.Table,
			Columns: hiringteam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringteam.FieldID,
			},
		},
	}
	if ps := htu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := htu.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringteam.FieldUpdatedAt, field.TypeTime, value)
	}
	if htu.mutation.UpdatedAtCleared() {
		_spec.ClearField(hiringteam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := htu.mutation.DeletedAt(); ok {
		_spec.SetField(hiringteam.FieldDeletedAt, field.TypeTime, value)
	}
	if htu.mutation.DeletedAtCleared() {
		_spec.ClearField(hiringteam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := htu.mutation.Slug(); ok {
		_spec.SetField(hiringteam.FieldSlug, field.TypeString, value)
	}
	if value, ok := htu.mutation.Name(); ok {
		_spec.SetField(hiringteam.FieldName, field.TypeString, value)
	}
	if htu.mutation.UserEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		createE := &HiringTeamManagerCreate{config: htu.config, mutation: newHiringTeamManagerMutation(htu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.RemovedUserEdgesIDs(); len(nodes) > 0 && !htu.mutation.UserEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
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
		createE := &HiringTeamManagerCreate{config: htu.config, mutation: newHiringTeamManagerMutation(htu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.UserEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
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
		createE := &HiringTeamManagerCreate{config: htu.config, mutation: newHiringTeamManagerMutation(htu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if htu.mutation.HiringTeamJobEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.RemovedHiringTeamJobEdgesIDs(); len(nodes) > 0 && !htu.mutation.HiringTeamJobEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.HiringTeamJobEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if htu.mutation.UserHiringTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.RemovedUserHiringTeamsIDs(); len(nodes) > 0 && !htu.mutation.UserHiringTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htu.mutation.UserHiringTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, htu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hiringteam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// HiringTeamUpdateOne is the builder for updating a single HiringTeam entity.
type HiringTeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HiringTeamMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (htuo *HiringTeamUpdateOne) SetUpdatedAt(t time.Time) *HiringTeamUpdateOne {
	htuo.mutation.SetUpdatedAt(t)
	return htuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (htuo *HiringTeamUpdateOne) SetNillableUpdatedAt(t *time.Time) *HiringTeamUpdateOne {
	if t != nil {
		htuo.SetUpdatedAt(*t)
	}
	return htuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (htuo *HiringTeamUpdateOne) ClearUpdatedAt() *HiringTeamUpdateOne {
	htuo.mutation.ClearUpdatedAt()
	return htuo
}

// SetDeletedAt sets the "deleted_at" field.
func (htuo *HiringTeamUpdateOne) SetDeletedAt(t time.Time) *HiringTeamUpdateOne {
	htuo.mutation.SetDeletedAt(t)
	return htuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (htuo *HiringTeamUpdateOne) SetNillableDeletedAt(t *time.Time) *HiringTeamUpdateOne {
	if t != nil {
		htuo.SetDeletedAt(*t)
	}
	return htuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (htuo *HiringTeamUpdateOne) ClearDeletedAt() *HiringTeamUpdateOne {
	htuo.mutation.ClearDeletedAt()
	return htuo
}

// SetSlug sets the "slug" field.
func (htuo *HiringTeamUpdateOne) SetSlug(s string) *HiringTeamUpdateOne {
	htuo.mutation.SetSlug(s)
	return htuo
}

// SetName sets the "name" field.
func (htuo *HiringTeamUpdateOne) SetName(s string) *HiringTeamUpdateOne {
	htuo.mutation.SetName(s)
	return htuo
}

// AddUserEdgeIDs adds the "user_edges" edge to the User entity by IDs.
func (htuo *HiringTeamUpdateOne) AddUserEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.AddUserEdgeIDs(ids...)
	return htuo
}

// AddUserEdges adds the "user_edges" edges to the User entity.
func (htuo *HiringTeamUpdateOne) AddUserEdges(u ...*User) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return htuo.AddUserEdgeIDs(ids...)
}

// AddHiringTeamJobEdgeIDs adds the "hiring_team_job_edges" edge to the HiringJob entity by IDs.
func (htuo *HiringTeamUpdateOne) AddHiringTeamJobEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.AddHiringTeamJobEdgeIDs(ids...)
	return htuo
}

// AddHiringTeamJobEdges adds the "hiring_team_job_edges" edges to the HiringJob entity.
func (htuo *HiringTeamUpdateOne) AddHiringTeamJobEdges(h ...*HiringJob) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htuo.AddHiringTeamJobEdgeIDs(ids...)
}

// AddUserHiringTeamIDs adds the "user_hiring_teams" edge to the HiringTeamManager entity by IDs.
func (htuo *HiringTeamUpdateOne) AddUserHiringTeamIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.AddUserHiringTeamIDs(ids...)
	return htuo
}

// AddUserHiringTeams adds the "user_hiring_teams" edges to the HiringTeamManager entity.
func (htuo *HiringTeamUpdateOne) AddUserHiringTeams(h ...*HiringTeamManager) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htuo.AddUserHiringTeamIDs(ids...)
}

// Mutation returns the HiringTeamMutation object of the builder.
func (htuo *HiringTeamUpdateOne) Mutation() *HiringTeamMutation {
	return htuo.mutation
}

// ClearUserEdges clears all "user_edges" edges to the User entity.
func (htuo *HiringTeamUpdateOne) ClearUserEdges() *HiringTeamUpdateOne {
	htuo.mutation.ClearUserEdges()
	return htuo
}

// RemoveUserEdgeIDs removes the "user_edges" edge to User entities by IDs.
func (htuo *HiringTeamUpdateOne) RemoveUserEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.RemoveUserEdgeIDs(ids...)
	return htuo
}

// RemoveUserEdges removes "user_edges" edges to User entities.
func (htuo *HiringTeamUpdateOne) RemoveUserEdges(u ...*User) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return htuo.RemoveUserEdgeIDs(ids...)
}

// ClearHiringTeamJobEdges clears all "hiring_team_job_edges" edges to the HiringJob entity.
func (htuo *HiringTeamUpdateOne) ClearHiringTeamJobEdges() *HiringTeamUpdateOne {
	htuo.mutation.ClearHiringTeamJobEdges()
	return htuo
}

// RemoveHiringTeamJobEdgeIDs removes the "hiring_team_job_edges" edge to HiringJob entities by IDs.
func (htuo *HiringTeamUpdateOne) RemoveHiringTeamJobEdgeIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.RemoveHiringTeamJobEdgeIDs(ids...)
	return htuo
}

// RemoveHiringTeamJobEdges removes "hiring_team_job_edges" edges to HiringJob entities.
func (htuo *HiringTeamUpdateOne) RemoveHiringTeamJobEdges(h ...*HiringJob) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htuo.RemoveHiringTeamJobEdgeIDs(ids...)
}

// ClearUserHiringTeams clears all "user_hiring_teams" edges to the HiringTeamManager entity.
func (htuo *HiringTeamUpdateOne) ClearUserHiringTeams() *HiringTeamUpdateOne {
	htuo.mutation.ClearUserHiringTeams()
	return htuo
}

// RemoveUserHiringTeamIDs removes the "user_hiring_teams" edge to HiringTeamManager entities by IDs.
func (htuo *HiringTeamUpdateOne) RemoveUserHiringTeamIDs(ids ...uuid.UUID) *HiringTeamUpdateOne {
	htuo.mutation.RemoveUserHiringTeamIDs(ids...)
	return htuo
}

// RemoveUserHiringTeams removes "user_hiring_teams" edges to HiringTeamManager entities.
func (htuo *HiringTeamUpdateOne) RemoveUserHiringTeams(h ...*HiringTeamManager) *HiringTeamUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return htuo.RemoveUserHiringTeamIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (htuo *HiringTeamUpdateOne) Select(field string, fields ...string) *HiringTeamUpdateOne {
	htuo.fields = append([]string{field}, fields...)
	return htuo
}

// Save executes the query and returns the updated HiringTeam entity.
func (htuo *HiringTeamUpdateOne) Save(ctx context.Context) (*HiringTeam, error) {
	var (
		err  error
		node *HiringTeam
	)
	if len(htuo.hooks) == 0 {
		if err = htuo.check(); err != nil {
			return nil, err
		}
		node, err = htuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringTeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = htuo.check(); err != nil {
				return nil, err
			}
			htuo.mutation = mutation
			node, err = htuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(htuo.hooks) - 1; i >= 0; i-- {
			if htuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = htuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, htuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HiringTeam)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HiringTeamMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (htuo *HiringTeamUpdateOne) SaveX(ctx context.Context) *HiringTeam {
	node, err := htuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (htuo *HiringTeamUpdateOne) Exec(ctx context.Context) error {
	_, err := htuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htuo *HiringTeamUpdateOne) ExecX(ctx context.Context) {
	if err := htuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (htuo *HiringTeamUpdateOne) check() error {
	if v, ok := htuo.mutation.Slug(); ok {
		if err := hiringteam.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "HiringTeam.slug": %w`, err)}
		}
	}
	if v, ok := htuo.mutation.Name(); ok {
		if err := hiringteam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HiringTeam.name": %w`, err)}
		}
	}
	return nil
}

func (htuo *HiringTeamUpdateOne) sqlSave(ctx context.Context) (_node *HiringTeam, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hiringteam.Table,
			Columns: hiringteam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringteam.FieldID,
			},
		},
	}
	id, ok := htuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HiringTeam.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := htuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hiringteam.FieldID)
		for _, f := range fields {
			if !hiringteam.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hiringteam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := htuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := htuo.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringteam.FieldUpdatedAt, field.TypeTime, value)
	}
	if htuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(hiringteam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := htuo.mutation.DeletedAt(); ok {
		_spec.SetField(hiringteam.FieldDeletedAt, field.TypeTime, value)
	}
	if htuo.mutation.DeletedAtCleared() {
		_spec.ClearField(hiringteam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := htuo.mutation.Slug(); ok {
		_spec.SetField(hiringteam.FieldSlug, field.TypeString, value)
	}
	if value, ok := htuo.mutation.Name(); ok {
		_spec.SetField(hiringteam.FieldName, field.TypeString, value)
	}
	if htuo.mutation.UserEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		createE := &HiringTeamManagerCreate{config: htuo.config, mutation: newHiringTeamManagerMutation(htuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.RemovedUserEdgesIDs(); len(nodes) > 0 && !htuo.mutation.UserEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
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
		createE := &HiringTeamManagerCreate{config: htuo.config, mutation: newHiringTeamManagerMutation(htuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.UserEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hiringteam.UserEdgesTable,
			Columns: hiringteam.UserEdgesPrimaryKey,
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
		createE := &HiringTeamManagerCreate{config: htuo.config, mutation: newHiringTeamManagerMutation(htuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if htuo.mutation.HiringTeamJobEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.RemovedHiringTeamJobEdgesIDs(); len(nodes) > 0 && !htuo.mutation.HiringTeamJobEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.HiringTeamJobEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringteam.HiringTeamJobEdgesTable,
			Columns: []string{hiringteam.HiringTeamJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if htuo.mutation.UserHiringTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.RemovedUserHiringTeamsIDs(); len(nodes) > 0 && !htuo.mutation.UserHiringTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := htuo.mutation.UserHiringTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   hiringteam.UserHiringTeamsTable,
			Columns: []string{hiringteam.UserHiringTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HiringTeam{config: htuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, htuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hiringteam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

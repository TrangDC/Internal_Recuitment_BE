// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/candidatejob"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/team"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringJobCreate is the builder for creating a HiringJob entity.
type HiringJobCreate struct {
	config
	mutation *HiringJobMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hjc *HiringJobCreate) SetCreatedAt(t time.Time) *HiringJobCreate {
	hjc.mutation.SetCreatedAt(t)
	return hjc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableCreatedAt(t *time.Time) *HiringJobCreate {
	if t != nil {
		hjc.SetCreatedAt(*t)
	}
	return hjc
}

// SetUpdatedAt sets the "updated_at" field.
func (hjc *HiringJobCreate) SetUpdatedAt(t time.Time) *HiringJobCreate {
	hjc.mutation.SetUpdatedAt(t)
	return hjc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableUpdatedAt(t *time.Time) *HiringJobCreate {
	if t != nil {
		hjc.SetUpdatedAt(*t)
	}
	return hjc
}

// SetDeletedAt sets the "deleted_at" field.
func (hjc *HiringJobCreate) SetDeletedAt(t time.Time) *HiringJobCreate {
	hjc.mutation.SetDeletedAt(t)
	return hjc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableDeletedAt(t *time.Time) *HiringJobCreate {
	if t != nil {
		hjc.SetDeletedAt(*t)
	}
	return hjc
}

// SetSlug sets the "slug" field.
func (hjc *HiringJobCreate) SetSlug(s string) *HiringJobCreate {
	hjc.mutation.SetSlug(s)
	return hjc
}

// SetName sets the "name" field.
func (hjc *HiringJobCreate) SetName(s string) *HiringJobCreate {
	hjc.mutation.SetName(s)
	return hjc
}

// SetDescription sets the "description" field.
func (hjc *HiringJobCreate) SetDescription(s string) *HiringJobCreate {
	hjc.mutation.SetDescription(s)
	return hjc
}

// SetAmount sets the "amount" field.
func (hjc *HiringJobCreate) SetAmount(i int) *HiringJobCreate {
	hjc.mutation.SetAmount(i)
	return hjc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableAmount(i *int) *HiringJobCreate {
	if i != nil {
		hjc.SetAmount(*i)
	}
	return hjc
}

// SetStatus sets the "status" field.
func (hjc *HiringJobCreate) SetStatus(h hiringjob.Status) *HiringJobCreate {
	hjc.mutation.SetStatus(h)
	return hjc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableStatus(h *hiringjob.Status) *HiringJobCreate {
	if h != nil {
		hjc.SetStatus(*h)
	}
	return hjc
}

// SetCreatedBy sets the "created_by" field.
func (hjc *HiringJobCreate) SetCreatedBy(u uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetCreatedBy(u)
	return hjc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableCreatedBy(u *uuid.UUID) *HiringJobCreate {
	if u != nil {
		hjc.SetCreatedBy(*u)
	}
	return hjc
}

// SetTeamID sets the "team_id" field.
func (hjc *HiringJobCreate) SetTeamID(u uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetTeamID(u)
	return hjc
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableTeamID(u *uuid.UUID) *HiringJobCreate {
	if u != nil {
		hjc.SetTeamID(*u)
	}
	return hjc
}

// SetLocation sets the "location" field.
func (hjc *HiringJobCreate) SetLocation(h hiringjob.Location) *HiringJobCreate {
	hjc.mutation.SetLocation(h)
	return hjc
}

// SetSalaryType sets the "salary_type" field.
func (hjc *HiringJobCreate) SetSalaryType(ht hiringjob.SalaryType) *HiringJobCreate {
	hjc.mutation.SetSalaryType(ht)
	return hjc
}

// SetSalaryFrom sets the "salary_from" field.
func (hjc *HiringJobCreate) SetSalaryFrom(i int) *HiringJobCreate {
	hjc.mutation.SetSalaryFrom(i)
	return hjc
}

// SetNillableSalaryFrom sets the "salary_from" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableSalaryFrom(i *int) *HiringJobCreate {
	if i != nil {
		hjc.SetSalaryFrom(*i)
	}
	return hjc
}

// SetSalaryTo sets the "salary_to" field.
func (hjc *HiringJobCreate) SetSalaryTo(i int) *HiringJobCreate {
	hjc.mutation.SetSalaryTo(i)
	return hjc
}

// SetNillableSalaryTo sets the "salary_to" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableSalaryTo(i *int) *HiringJobCreate {
	if i != nil {
		hjc.SetSalaryTo(*i)
	}
	return hjc
}

// SetCurrency sets the "currency" field.
func (hjc *HiringJobCreate) SetCurrency(h hiringjob.Currency) *HiringJobCreate {
	hjc.mutation.SetCurrency(h)
	return hjc
}

// SetLastApplyDate sets the "last_apply_date" field.
func (hjc *HiringJobCreate) SetLastApplyDate(t time.Time) *HiringJobCreate {
	hjc.mutation.SetLastApplyDate(t)
	return hjc
}

// SetNillableLastApplyDate sets the "last_apply_date" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableLastApplyDate(t *time.Time) *HiringJobCreate {
	if t != nil {
		hjc.SetLastApplyDate(*t)
	}
	return hjc
}

// SetPriority sets the "priority" field.
func (hjc *HiringJobCreate) SetPriority(i int) *HiringJobCreate {
	hjc.mutation.SetPriority(i)
	return hjc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillablePriority(i *int) *HiringJobCreate {
	if i != nil {
		hjc.SetPriority(*i)
	}
	return hjc
}

// SetHiringTeamID sets the "hiring_team_id" field.
func (hjc *HiringJobCreate) SetHiringTeamID(u uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetHiringTeamID(u)
	return hjc
}

// SetNillableHiringTeamID sets the "hiring_team_id" field if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableHiringTeamID(u *uuid.UUID) *HiringJobCreate {
	if u != nil {
		hjc.SetHiringTeamID(*u)
	}
	return hjc
}

// SetID sets the "id" field.
func (hjc *HiringJobCreate) SetID(u uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetID(u)
	return hjc
}

// SetOwnerEdgeID sets the "owner_edge" edge to the User entity by ID.
func (hjc *HiringJobCreate) SetOwnerEdgeID(id uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetOwnerEdgeID(id)
	return hjc
}

// SetNillableOwnerEdgeID sets the "owner_edge" edge to the User entity by ID if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableOwnerEdgeID(id *uuid.UUID) *HiringJobCreate {
	if id != nil {
		hjc = hjc.SetOwnerEdgeID(*id)
	}
	return hjc
}

// SetOwnerEdge sets the "owner_edge" edge to the User entity.
func (hjc *HiringJobCreate) SetOwnerEdge(u *User) *HiringJobCreate {
	return hjc.SetOwnerEdgeID(u.ID)
}

// SetTeamEdgeID sets the "team_edge" edge to the Team entity by ID.
func (hjc *HiringJobCreate) SetTeamEdgeID(id uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetTeamEdgeID(id)
	return hjc
}

// SetNillableTeamEdgeID sets the "team_edge" edge to the Team entity by ID if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableTeamEdgeID(id *uuid.UUID) *HiringJobCreate {
	if id != nil {
		hjc = hjc.SetTeamEdgeID(*id)
	}
	return hjc
}

// SetTeamEdge sets the "team_edge" edge to the Team entity.
func (hjc *HiringJobCreate) SetTeamEdge(t *Team) *HiringJobCreate {
	return hjc.SetTeamEdgeID(t.ID)
}

// AddCandidateJobEdgeIDs adds the "candidate_job_edges" edge to the CandidateJob entity by IDs.
func (hjc *HiringJobCreate) AddCandidateJobEdgeIDs(ids ...uuid.UUID) *HiringJobCreate {
	hjc.mutation.AddCandidateJobEdgeIDs(ids...)
	return hjc
}

// AddCandidateJobEdges adds the "candidate_job_edges" edges to the CandidateJob entity.
func (hjc *HiringJobCreate) AddCandidateJobEdges(c ...*CandidateJob) *HiringJobCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hjc.AddCandidateJobEdgeIDs(ids...)
}

// AddHiringJobSkillEdgeIDs adds the "hiring_job_skill_edges" edge to the EntitySkill entity by IDs.
func (hjc *HiringJobCreate) AddHiringJobSkillEdgeIDs(ids ...uuid.UUID) *HiringJobCreate {
	hjc.mutation.AddHiringJobSkillEdgeIDs(ids...)
	return hjc
}

// AddHiringJobSkillEdges adds the "hiring_job_skill_edges" edges to the EntitySkill entity.
func (hjc *HiringJobCreate) AddHiringJobSkillEdges(e ...*EntitySkill) *HiringJobCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return hjc.AddHiringJobSkillEdgeIDs(ids...)
}

// SetHiringTeamEdgeID sets the "hiring_team_edge" edge to the HiringTeam entity by ID.
func (hjc *HiringJobCreate) SetHiringTeamEdgeID(id uuid.UUID) *HiringJobCreate {
	hjc.mutation.SetHiringTeamEdgeID(id)
	return hjc
}

// SetNillableHiringTeamEdgeID sets the "hiring_team_edge" edge to the HiringTeam entity by ID if the given value is not nil.
func (hjc *HiringJobCreate) SetNillableHiringTeamEdgeID(id *uuid.UUID) *HiringJobCreate {
	if id != nil {
		hjc = hjc.SetHiringTeamEdgeID(*id)
	}
	return hjc
}

// SetHiringTeamEdge sets the "hiring_team_edge" edge to the HiringTeam entity.
func (hjc *HiringJobCreate) SetHiringTeamEdge(h *HiringTeam) *HiringJobCreate {
	return hjc.SetHiringTeamEdgeID(h.ID)
}

// Mutation returns the HiringJobMutation object of the builder.
func (hjc *HiringJobCreate) Mutation() *HiringJobMutation {
	return hjc.mutation
}

// Save creates the HiringJob in the database.
func (hjc *HiringJobCreate) Save(ctx context.Context) (*HiringJob, error) {
	var (
		err  error
		node *HiringJob
	)
	hjc.defaults()
	if len(hjc.hooks) == 0 {
		if err = hjc.check(); err != nil {
			return nil, err
		}
		node, err = hjc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hjc.check(); err != nil {
				return nil, err
			}
			hjc.mutation = mutation
			if node, err = hjc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hjc.hooks) - 1; i >= 0; i-- {
			if hjc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hjc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hjc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HiringJob)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HiringJobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hjc *HiringJobCreate) SaveX(ctx context.Context) *HiringJob {
	v, err := hjc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hjc *HiringJobCreate) Exec(ctx context.Context) error {
	_, err := hjc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjc *HiringJobCreate) ExecX(ctx context.Context) {
	if err := hjc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hjc *HiringJobCreate) defaults() {
	if _, ok := hjc.mutation.CreatedAt(); !ok {
		v := hiringjob.DefaultCreatedAt()
		hjc.mutation.SetCreatedAt(v)
	}
	if _, ok := hjc.mutation.Amount(); !ok {
		v := hiringjob.DefaultAmount
		hjc.mutation.SetAmount(v)
	}
	if _, ok := hjc.mutation.Status(); !ok {
		v := hiringjob.DefaultStatus
		hjc.mutation.SetStatus(v)
	}
	if _, ok := hjc.mutation.SalaryFrom(); !ok {
		v := hiringjob.DefaultSalaryFrom
		hjc.mutation.SetSalaryFrom(v)
	}
	if _, ok := hjc.mutation.SalaryTo(); !ok {
		v := hiringjob.DefaultSalaryTo
		hjc.mutation.SetSalaryTo(v)
	}
	if _, ok := hjc.mutation.Priority(); !ok {
		v := hiringjob.DefaultPriority
		hjc.mutation.SetPriority(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hjc *HiringJobCreate) check() error {
	if _, ok := hjc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HiringJob.created_at"`)}
	}
	if _, ok := hjc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "HiringJob.slug"`)}
	}
	if v, ok := hjc.mutation.Slug(); ok {
		if err := hiringjob.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "HiringJob.slug": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "HiringJob.name"`)}
	}
	if v, ok := hjc.mutation.Name(); ok {
		if err := hiringjob.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HiringJob.name": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "HiringJob.description"`)}
	}
	if v, ok := hjc.mutation.Description(); ok {
		if err := hiringjob.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "HiringJob.description": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "HiringJob.amount"`)}
	}
	if _, ok := hjc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "HiringJob.status"`)}
	}
	if v, ok := hjc.mutation.Status(); ok {
		if err := hiringjob.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HiringJob.status": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "HiringJob.location"`)}
	}
	if v, ok := hjc.mutation.Location(); ok {
		if err := hiringjob.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "HiringJob.location": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.SalaryType(); !ok {
		return &ValidationError{Name: "salary_type", err: errors.New(`ent: missing required field "HiringJob.salary_type"`)}
	}
	if v, ok := hjc.mutation.SalaryType(); ok {
		if err := hiringjob.SalaryTypeValidator(v); err != nil {
			return &ValidationError{Name: "salary_type", err: fmt.Errorf(`ent: validator failed for field "HiringJob.salary_type": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.SalaryFrom(); !ok {
		return &ValidationError{Name: "salary_from", err: errors.New(`ent: missing required field "HiringJob.salary_from"`)}
	}
	if _, ok := hjc.mutation.SalaryTo(); !ok {
		return &ValidationError{Name: "salary_to", err: errors.New(`ent: missing required field "HiringJob.salary_to"`)}
	}
	if _, ok := hjc.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "HiringJob.currency"`)}
	}
	if v, ok := hjc.mutation.Currency(); ok {
		if err := hiringjob.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "HiringJob.currency": %w`, err)}
		}
	}
	if _, ok := hjc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "HiringJob.priority"`)}
	}
	return nil
}

func (hjc *HiringJobCreate) sqlSave(ctx context.Context) (*HiringJob, error) {
	_node, _spec := hjc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hjc.driver, _spec); err != nil {
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

func (hjc *HiringJobCreate) createSpec() (*HiringJob, *sqlgraph.CreateSpec) {
	var (
		_node = &HiringJob{config: hjc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hiringjob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjob.FieldID,
			},
		}
	)
	if id, ok := hjc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hjc.mutation.CreatedAt(); ok {
		_spec.SetField(hiringjob.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hjc.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringjob.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hjc.mutation.DeletedAt(); ok {
		_spec.SetField(hiringjob.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := hjc.mutation.Slug(); ok {
		_spec.SetField(hiringjob.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := hjc.mutation.Name(); ok {
		_spec.SetField(hiringjob.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hjc.mutation.Description(); ok {
		_spec.SetField(hiringjob.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := hjc.mutation.Amount(); ok {
		_spec.SetField(hiringjob.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := hjc.mutation.Status(); ok {
		_spec.SetField(hiringjob.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := hjc.mutation.Location(); ok {
		_spec.SetField(hiringjob.FieldLocation, field.TypeEnum, value)
		_node.Location = value
	}
	if value, ok := hjc.mutation.SalaryType(); ok {
		_spec.SetField(hiringjob.FieldSalaryType, field.TypeEnum, value)
		_node.SalaryType = value
	}
	if value, ok := hjc.mutation.SalaryFrom(); ok {
		_spec.SetField(hiringjob.FieldSalaryFrom, field.TypeInt, value)
		_node.SalaryFrom = value
	}
	if value, ok := hjc.mutation.SalaryTo(); ok {
		_spec.SetField(hiringjob.FieldSalaryTo, field.TypeInt, value)
		_node.SalaryTo = value
	}
	if value, ok := hjc.mutation.Currency(); ok {
		_spec.SetField(hiringjob.FieldCurrency, field.TypeEnum, value)
		_node.Currency = value
	}
	if value, ok := hjc.mutation.LastApplyDate(); ok {
		_spec.SetField(hiringjob.FieldLastApplyDate, field.TypeTime, value)
		_node.LastApplyDate = value
	}
	if value, ok := hjc.mutation.Priority(); ok {
		_spec.SetField(hiringjob.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if nodes := hjc.mutation.OwnerEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hiringjob.OwnerEdgeTable,
			Columns: []string{hiringjob.OwnerEdgeColumn},
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
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hjc.mutation.TeamEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hiringjob.TeamEdgeTable,
			Columns: []string{hiringjob.TeamEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hjc.mutation.CandidateJobEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringjob.CandidateJobEdgesTable,
			Columns: []string{hiringjob.CandidateJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hjc.mutation.HiringJobSkillEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hiringjob.HiringJobSkillEdgesTable,
			Columns: []string{hiringjob.HiringJobSkillEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entityskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hjc.mutation.HiringTeamEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hiringjob.HiringTeamEdgeTable,
			Columns: []string{hiringjob.HiringTeamEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.HiringTeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HiringJobCreateBulk is the builder for creating many HiringJob entities in bulk.
type HiringJobCreateBulk struct {
	config
	builders []*HiringJobCreate
}

// Save creates the HiringJob entities in the database.
func (hjcb *HiringJobCreateBulk) Save(ctx context.Context) ([]*HiringJob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hjcb.builders))
	nodes := make([]*HiringJob, len(hjcb.builders))
	mutators := make([]Mutator, len(hjcb.builders))
	for i := range hjcb.builders {
		func(i int, root context.Context) {
			builder := hjcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HiringJobMutation)
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
					_, err = mutators[i+1].Mutate(root, hjcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hjcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hjcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hjcb *HiringJobCreateBulk) SaveX(ctx context.Context) []*HiringJob {
	v, err := hjcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hjcb *HiringJobCreateBulk) Exec(ctx context.Context) error {
	_, err := hjcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjcb *HiringJobCreateBulk) ExecX(ctx context.Context) {
	if err := hjcb.Exec(ctx); err != nil {
		panic(err)
	}
}

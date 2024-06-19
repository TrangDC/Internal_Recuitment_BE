// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidateinterview"
	"trec/ent/candidateinterviewer"
	"trec/ent/candidatejob"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateInterviewCreate is the builder for creating a CandidateInterview entity.
type CandidateInterviewCreate struct {
	config
	mutation *CandidateInterviewMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cic *CandidateInterviewCreate) SetCreatedAt(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetCreatedAt(t)
	return cic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCreatedAt(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetCreatedAt(*t)
	}
	return cic
}

// SetUpdatedAt sets the "updated_at" field.
func (cic *CandidateInterviewCreate) SetUpdatedAt(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetUpdatedAt(t)
	return cic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableUpdatedAt(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetUpdatedAt(*t)
	}
	return cic
}

// SetDeletedAt sets the "deleted_at" field.
func (cic *CandidateInterviewCreate) SetDeletedAt(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetDeletedAt(t)
	return cic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableDeletedAt(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetDeletedAt(*t)
	}
	return cic
}

// SetTitle sets the "title" field.
func (cic *CandidateInterviewCreate) SetTitle(s string) *CandidateInterviewCreate {
	cic.mutation.SetTitle(s)
	return cic
}

// SetCandidateJobStatus sets the "candidate_job_status" field.
func (cic *CandidateInterviewCreate) SetCandidateJobStatus(cjs candidateinterview.CandidateJobStatus) *CandidateInterviewCreate {
	cic.mutation.SetCandidateJobStatus(cjs)
	return cic
}

// SetNillableCandidateJobStatus sets the "candidate_job_status" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCandidateJobStatus(cjs *candidateinterview.CandidateJobStatus) *CandidateInterviewCreate {
	if cjs != nil {
		cic.SetCandidateJobStatus(*cjs)
	}
	return cic
}

// SetCandidateJobID sets the "candidate_job_id" field.
func (cic *CandidateInterviewCreate) SetCandidateJobID(u uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.SetCandidateJobID(u)
	return cic
}

// SetNillableCandidateJobID sets the "candidate_job_id" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCandidateJobID(u *uuid.UUID) *CandidateInterviewCreate {
	if u != nil {
		cic.SetCandidateJobID(*u)
	}
	return cic
}

// SetInterviewDate sets the "interview_date" field.
func (cic *CandidateInterviewCreate) SetInterviewDate(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetInterviewDate(t)
	return cic
}

// SetNillableInterviewDate sets the "interview_date" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableInterviewDate(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetInterviewDate(*t)
	}
	return cic
}

// SetStartFrom sets the "start_from" field.
func (cic *CandidateInterviewCreate) SetStartFrom(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetStartFrom(t)
	return cic
}

// SetNillableStartFrom sets the "start_from" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableStartFrom(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetStartFrom(*t)
	}
	return cic
}

// SetEndAt sets the "end_at" field.
func (cic *CandidateInterviewCreate) SetEndAt(t time.Time) *CandidateInterviewCreate {
	cic.mutation.SetEndAt(t)
	return cic
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableEndAt(t *time.Time) *CandidateInterviewCreate {
	if t != nil {
		cic.SetEndAt(*t)
	}
	return cic
}

// SetCreatedBy sets the "created_by" field.
func (cic *CandidateInterviewCreate) SetCreatedBy(u uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.SetCreatedBy(u)
	return cic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCreatedBy(u *uuid.UUID) *CandidateInterviewCreate {
	if u != nil {
		cic.SetCreatedBy(*u)
	}
	return cic
}

// SetDescription sets the "description" field.
func (cic *CandidateInterviewCreate) SetDescription(s string) *CandidateInterviewCreate {
	cic.mutation.SetDescription(s)
	return cic
}

// SetCandidateInterviewStatus sets the "candidate_interview_status" field.
func (cic *CandidateInterviewCreate) SetCandidateInterviewStatus(cis candidateinterview.CandidateInterviewStatus) *CandidateInterviewCreate {
	cic.mutation.SetCandidateInterviewStatus(cis)
	return cic
}

// SetNillableCandidateInterviewStatus sets the "candidate_interview_status" field if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCandidateInterviewStatus(cis *candidateinterview.CandidateInterviewStatus) *CandidateInterviewCreate {
	if cis != nil {
		cic.SetCandidateInterviewStatus(*cis)
	}
	return cic
}

// SetID sets the "id" field.
func (cic *CandidateInterviewCreate) SetID(u uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.SetID(u)
	return cic
}

// SetCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID.
func (cic *CandidateInterviewCreate) SetCandidateJobEdgeID(id uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.SetCandidateJobEdgeID(id)
	return cic
}

// SetNillableCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCandidateJobEdgeID(id *uuid.UUID) *CandidateInterviewCreate {
	if id != nil {
		cic = cic.SetCandidateJobEdgeID(*id)
	}
	return cic
}

// SetCandidateJobEdge sets the "candidate_job_edge" edge to the CandidateJob entity.
func (cic *CandidateInterviewCreate) SetCandidateJobEdge(c *CandidateJob) *CandidateInterviewCreate {
	return cic.SetCandidateJobEdgeID(c.ID)
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cic *CandidateInterviewCreate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.AddAttachmentEdgeIDs(ids...)
	return cic
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cic *CandidateInterviewCreate) AddAttachmentEdges(a ...*Attachment) *CandidateInterviewCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cic.AddAttachmentEdgeIDs(ids...)
}

// AddInterviewerEdgeIDs adds the "interviewer_edges" edge to the User entity by IDs.
func (cic *CandidateInterviewCreate) AddInterviewerEdgeIDs(ids ...uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.AddInterviewerEdgeIDs(ids...)
	return cic
}

// AddInterviewerEdges adds the "interviewer_edges" edges to the User entity.
func (cic *CandidateInterviewCreate) AddInterviewerEdges(u ...*User) *CandidateInterviewCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cic.AddInterviewerEdgeIDs(ids...)
}

// SetCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID.
func (cic *CandidateInterviewCreate) SetCreatedByEdgeID(id uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.SetCreatedByEdgeID(id)
	return cic
}

// SetNillableCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID if the given value is not nil.
func (cic *CandidateInterviewCreate) SetNillableCreatedByEdgeID(id *uuid.UUID) *CandidateInterviewCreate {
	if id != nil {
		cic = cic.SetCreatedByEdgeID(*id)
	}
	return cic
}

// SetCreatedByEdge sets the "created_by_edge" edge to the User entity.
func (cic *CandidateInterviewCreate) SetCreatedByEdge(u *User) *CandidateInterviewCreate {
	return cic.SetCreatedByEdgeID(u.ID)
}

// AddUserInterviewerIDs adds the "user_interviewers" edge to the CandidateInterviewer entity by IDs.
func (cic *CandidateInterviewCreate) AddUserInterviewerIDs(ids ...uuid.UUID) *CandidateInterviewCreate {
	cic.mutation.AddUserInterviewerIDs(ids...)
	return cic
}

// AddUserInterviewers adds the "user_interviewers" edges to the CandidateInterviewer entity.
func (cic *CandidateInterviewCreate) AddUserInterviewers(c ...*CandidateInterviewer) *CandidateInterviewCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cic.AddUserInterviewerIDs(ids...)
}

// Mutation returns the CandidateInterviewMutation object of the builder.
func (cic *CandidateInterviewCreate) Mutation() *CandidateInterviewMutation {
	return cic.mutation
}

// Save creates the CandidateInterview in the database.
func (cic *CandidateInterviewCreate) Save(ctx context.Context) (*CandidateInterview, error) {
	var (
		err  error
		node *CandidateInterview
	)
	cic.defaults()
	if len(cic.hooks) == 0 {
		if err = cic.check(); err != nil {
			return nil, err
		}
		node, err = cic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateInterviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cic.check(); err != nil {
				return nil, err
			}
			cic.mutation = mutation
			if node, err = cic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cic.hooks) - 1; i >= 0; i-- {
			if cic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateInterview)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateInterviewMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CandidateInterviewCreate) SaveX(ctx context.Context) *CandidateInterview {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *CandidateInterviewCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *CandidateInterviewCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *CandidateInterviewCreate) defaults() {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		v := candidateinterview.DefaultCreatedAt()
		cic.mutation.SetCreatedAt(v)
	}
	if _, ok := cic.mutation.CandidateJobStatus(); !ok {
		v := candidateinterview.DefaultCandidateJobStatus
		cic.mutation.SetCandidateJobStatus(v)
	}
	if _, ok := cic.mutation.CandidateInterviewStatus(); !ok {
		v := candidateinterview.DefaultCandidateInterviewStatus
		cic.mutation.SetCandidateInterviewStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *CandidateInterviewCreate) check() error {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CandidateInterview.created_at"`)}
	}
	if _, ok := cic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "CandidateInterview.title"`)}
	}
	if v, ok := cic.mutation.Title(); ok {
		if err := candidateinterview.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CandidateInterview.title": %w`, err)}
		}
	}
	if _, ok := cic.mutation.CandidateJobStatus(); !ok {
		return &ValidationError{Name: "candidate_job_status", err: errors.New(`ent: missing required field "CandidateInterview.candidate_job_status"`)}
	}
	if v, ok := cic.mutation.CandidateJobStatus(); ok {
		if err := candidateinterview.CandidateJobStatusValidator(v); err != nil {
			return &ValidationError{Name: "candidate_job_status", err: fmt.Errorf(`ent: validator failed for field "CandidateInterview.candidate_job_status": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "CandidateInterview.description"`)}
	}
	if _, ok := cic.mutation.CandidateInterviewStatus(); !ok {
		return &ValidationError{Name: "candidate_interview_status", err: errors.New(`ent: missing required field "CandidateInterview.candidate_interview_status"`)}
	}
	if v, ok := cic.mutation.CandidateInterviewStatus(); ok {
		if err := candidateinterview.CandidateInterviewStatusValidator(v); err != nil {
			return &ValidationError{Name: "candidate_interview_status", err: fmt.Errorf(`ent: validator failed for field "CandidateInterview.candidate_interview_status": %w`, err)}
		}
	}
	return nil
}

func (cic *CandidateInterviewCreate) sqlSave(ctx context.Context) (*CandidateInterview, error) {
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
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

func (cic *CandidateInterviewCreate) createSpec() (*CandidateInterview, *sqlgraph.CreateSpec) {
	var (
		_node = &CandidateInterview{config: cic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: candidateinterview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateinterview.FieldID,
			},
		}
	)
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cic.mutation.CreatedAt(); ok {
		_spec.SetField(candidateinterview.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cic.mutation.UpdatedAt(); ok {
		_spec.SetField(candidateinterview.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cic.mutation.DeletedAt(); ok {
		_spec.SetField(candidateinterview.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cic.mutation.Title(); ok {
		_spec.SetField(candidateinterview.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := cic.mutation.CandidateJobStatus(); ok {
		_spec.SetField(candidateinterview.FieldCandidateJobStatus, field.TypeEnum, value)
		_node.CandidateJobStatus = value
	}
	if value, ok := cic.mutation.InterviewDate(); ok {
		_spec.SetField(candidateinterview.FieldInterviewDate, field.TypeTime, value)
		_node.InterviewDate = value
	}
	if value, ok := cic.mutation.StartFrom(); ok {
		_spec.SetField(candidateinterview.FieldStartFrom, field.TypeTime, value)
		_node.StartFrom = value
	}
	if value, ok := cic.mutation.EndAt(); ok {
		_spec.SetField(candidateinterview.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := cic.mutation.Description(); ok {
		_spec.SetField(candidateinterview.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cic.mutation.CandidateInterviewStatus(); ok {
		_spec.SetField(candidateinterview.FieldCandidateInterviewStatus, field.TypeEnum, value)
		_node.CandidateInterviewStatus = value
	}
	if nodes := cic.mutation.CandidateJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateinterview.CandidateJobEdgeTable,
			Columns: []string{candidateinterview.CandidateJobEdgeColumn},
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
		_node.CandidateJobID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateinterview.AttachmentEdgesTable,
			Columns: []string{candidateinterview.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.InterviewerEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   candidateinterview.InterviewerEdgesTable,
			Columns: candidateinterview.InterviewerEdgesPrimaryKey,
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
		createE := &CandidateInterviewerCreate{config: cic.config, mutation: newCandidateInterviewerMutation(cic.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.CreatedByEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateinterview.CreatedByEdgeTable,
			Columns: []string{candidateinterview.CreatedByEdgeColumn},
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
	if nodes := cic.mutation.UserInterviewersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   candidateinterview.UserInterviewersTable,
			Columns: []string{candidateinterview.UserInterviewersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateinterviewer.FieldID,
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

// CandidateInterviewCreateBulk is the builder for creating many CandidateInterview entities in bulk.
type CandidateInterviewCreateBulk struct {
	config
	builders []*CandidateInterviewCreate
}

// Save creates the CandidateInterview entities in the database.
func (cicb *CandidateInterviewCreateBulk) Save(ctx context.Context) ([]*CandidateInterview, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*CandidateInterview, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CandidateInterviewMutation)
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
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *CandidateInterviewCreateBulk) SaveX(ctx context.Context) []*CandidateInterview {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *CandidateInterviewCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *CandidateInterviewCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}

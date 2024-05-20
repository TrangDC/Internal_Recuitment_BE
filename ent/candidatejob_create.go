// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/hiringjob"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateJobCreate is the builder for creating a CandidateJob entity.
type CandidateJobCreate struct {
	config
	mutation *CandidateJobMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cjc *CandidateJobCreate) SetCreatedAt(t time.Time) *CandidateJobCreate {
	cjc.mutation.SetCreatedAt(t)
	return cjc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableCreatedAt(t *time.Time) *CandidateJobCreate {
	if t != nil {
		cjc.SetCreatedAt(*t)
	}
	return cjc
}

// SetUpdatedAt sets the "updated_at" field.
func (cjc *CandidateJobCreate) SetUpdatedAt(t time.Time) *CandidateJobCreate {
	cjc.mutation.SetUpdatedAt(t)
	return cjc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableUpdatedAt(t *time.Time) *CandidateJobCreate {
	if t != nil {
		cjc.SetUpdatedAt(*t)
	}
	return cjc
}

// SetDeletedAt sets the "deleted_at" field.
func (cjc *CandidateJobCreate) SetDeletedAt(t time.Time) *CandidateJobCreate {
	cjc.mutation.SetDeletedAt(t)
	return cjc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableDeletedAt(t *time.Time) *CandidateJobCreate {
	if t != nil {
		cjc.SetDeletedAt(*t)
	}
	return cjc
}

// SetHiringJobID sets the "hiring_job_id" field.
func (cjc *CandidateJobCreate) SetHiringJobID(u uuid.UUID) *CandidateJobCreate {
	cjc.mutation.SetHiringJobID(u)
	return cjc
}

// SetNillableHiringJobID sets the "hiring_job_id" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableHiringJobID(u *uuid.UUID) *CandidateJobCreate {
	if u != nil {
		cjc.SetHiringJobID(*u)
	}
	return cjc
}

// SetCandidateID sets the "candidate_id" field.
func (cjc *CandidateJobCreate) SetCandidateID(u uuid.UUID) *CandidateJobCreate {
	cjc.mutation.SetCandidateID(u)
	return cjc
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableCandidateID(u *uuid.UUID) *CandidateJobCreate {
	if u != nil {
		cjc.SetCandidateID(*u)
	}
	return cjc
}

// SetStatus sets the "status" field.
func (cjc *CandidateJobCreate) SetStatus(c candidatejob.Status) *CandidateJobCreate {
	cjc.mutation.SetStatus(c)
	return cjc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableStatus(c *candidatejob.Status) *CandidateJobCreate {
	if c != nil {
		cjc.SetStatus(*c)
	}
	return cjc
}

// SetID sets the "id" field.
func (cjc *CandidateJobCreate) SetID(u uuid.UUID) *CandidateJobCreate {
	cjc.mutation.SetID(u)
	return cjc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableID(u *uuid.UUID) *CandidateJobCreate {
	if u != nil {
		cjc.SetID(*u)
	}
	return cjc
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cjc *CandidateJobCreate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobCreate {
	cjc.mutation.AddAttachmentEdgeIDs(ids...)
	return cjc
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cjc *CandidateJobCreate) AddAttachmentEdges(a ...*Attachment) *CandidateJobCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjc.AddAttachmentEdgeIDs(ids...)
}

// SetHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID.
func (cjc *CandidateJobCreate) SetHiringJobEdgeID(id uuid.UUID) *CandidateJobCreate {
	cjc.mutation.SetHiringJobEdgeID(id)
	return cjc
}

// SetNillableHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableHiringJobEdgeID(id *uuid.UUID) *CandidateJobCreate {
	if id != nil {
		cjc = cjc.SetHiringJobEdgeID(*id)
	}
	return cjc
}

// SetHiringJobEdge sets the "hiring_job_edge" edge to the HiringJob entity.
func (cjc *CandidateJobCreate) SetHiringJobEdge(h *HiringJob) *CandidateJobCreate {
	return cjc.SetHiringJobEdgeID(h.ID)
}

// AddCandidateJobFeedbackIDs adds the "candidate_job_feedback" edge to the CandidateJobFeedback entity by IDs.
func (cjc *CandidateJobCreate) AddCandidateJobFeedbackIDs(ids ...uuid.UUID) *CandidateJobCreate {
	cjc.mutation.AddCandidateJobFeedbackIDs(ids...)
	return cjc
}

// AddCandidateJobFeedback adds the "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (cjc *CandidateJobCreate) AddCandidateJobFeedback(c ...*CandidateJobFeedback) *CandidateJobCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjc.AddCandidateJobFeedbackIDs(ids...)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cjc *CandidateJobCreate) SetCandidateEdgeID(id uuid.UUID) *CandidateJobCreate {
	cjc.mutation.SetCandidateEdgeID(id)
	return cjc
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cjc *CandidateJobCreate) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateJobCreate {
	if id != nil {
		cjc = cjc.SetCandidateEdgeID(*id)
	}
	return cjc
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cjc *CandidateJobCreate) SetCandidateEdge(c *Candidate) *CandidateJobCreate {
	return cjc.SetCandidateEdgeID(c.ID)
}

// AddCandidateJobInterviewIDs adds the "candidate_job_interview" edge to the CandidateInterview entity by IDs.
func (cjc *CandidateJobCreate) AddCandidateJobInterviewIDs(ids ...uuid.UUID) *CandidateJobCreate {
	cjc.mutation.AddCandidateJobInterviewIDs(ids...)
	return cjc
}

// AddCandidateJobInterview adds the "candidate_job_interview" edges to the CandidateInterview entity.
func (cjc *CandidateJobCreate) AddCandidateJobInterview(c ...*CandidateInterview) *CandidateJobCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjc.AddCandidateJobInterviewIDs(ids...)
}

// Mutation returns the CandidateJobMutation object of the builder.
func (cjc *CandidateJobCreate) Mutation() *CandidateJobMutation {
	return cjc.mutation
}

// Save creates the CandidateJob in the database.
func (cjc *CandidateJobCreate) Save(ctx context.Context) (*CandidateJob, error) {
	var (
		err  error
		node *CandidateJob
	)
	cjc.defaults()
	if len(cjc.hooks) == 0 {
		if err = cjc.check(); err != nil {
			return nil, err
		}
		node, err = cjc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjc.check(); err != nil {
				return nil, err
			}
			cjc.mutation = mutation
			if node, err = cjc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cjc.hooks) - 1; i >= 0; i-- {
			if cjc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cjc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateJob)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateJobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cjc *CandidateJobCreate) SaveX(ctx context.Context) *CandidateJob {
	v, err := cjc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjc *CandidateJobCreate) Exec(ctx context.Context) error {
	_, err := cjc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjc *CandidateJobCreate) ExecX(ctx context.Context) {
	if err := cjc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cjc *CandidateJobCreate) defaults() {
	if _, ok := cjc.mutation.CreatedAt(); !ok {
		v := candidatejob.DefaultCreatedAt()
		cjc.mutation.SetCreatedAt(v)
	}
	if _, ok := cjc.mutation.Status(); !ok {
		v := candidatejob.DefaultStatus
		cjc.mutation.SetStatus(v)
	}
	if _, ok := cjc.mutation.ID(); !ok {
		v := candidatejob.DefaultID()
		cjc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjc *CandidateJobCreate) check() error {
	if _, ok := cjc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CandidateJob.created_at"`)}
	}
	if _, ok := cjc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "CandidateJob.status"`)}
	}
	if v, ok := cjc.mutation.Status(); ok {
		if err := candidatejob.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CandidateJob.status": %w`, err)}
		}
	}
	return nil
}

func (cjc *CandidateJobCreate) sqlSave(ctx context.Context) (*CandidateJob, error) {
	_node, _spec := cjc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cjc.driver, _spec); err != nil {
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

func (cjc *CandidateJobCreate) createSpec() (*CandidateJob, *sqlgraph.CreateSpec) {
	var (
		_node = &CandidateJob{config: cjc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: candidatejob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejob.FieldID,
			},
		}
	)
	if id, ok := cjc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cjc.mutation.CreatedAt(); ok {
		_spec.SetField(candidatejob.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cjc.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejob.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cjc.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejob.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cjc.mutation.Status(); ok {
		_spec.SetField(candidatejob.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := cjc.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejob.AttachmentEdgesTable,
			Columns: []string{candidatejob.AttachmentEdgesColumn},
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
	if nodes := cjc.mutation.HiringJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.HiringJobEdgeTable,
			Columns: []string{candidatejob.HiringJobEdgeColumn},
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
		_node.HiringJobID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cjc.mutation.CandidateJobFeedbackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejob.CandidateJobFeedbackTable,
			Columns: []string{candidatejob.CandidateJobFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejobfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cjc.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.CandidateEdgeTable,
			Columns: []string{candidatejob.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CandidateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cjc.mutation.CandidateJobInterviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejob.CandidateJobInterviewTable,
			Columns: []string{candidatejob.CandidateJobInterviewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateinterview.FieldID,
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

// CandidateJobCreateBulk is the builder for creating many CandidateJob entities in bulk.
type CandidateJobCreateBulk struct {
	config
	builders []*CandidateJobCreate
}

// Save creates the CandidateJob entities in the database.
func (cjcb *CandidateJobCreateBulk) Save(ctx context.Context) ([]*CandidateJob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cjcb.builders))
	nodes := make([]*CandidateJob, len(cjcb.builders))
	mutators := make([]Mutator, len(cjcb.builders))
	for i := range cjcb.builders {
		func(i int, root context.Context) {
			builder := cjcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CandidateJobMutation)
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
					_, err = mutators[i+1].Mutate(root, cjcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cjcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cjcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cjcb *CandidateJobCreateBulk) SaveX(ctx context.Context) []*CandidateJob {
	v, err := cjcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjcb *CandidateJobCreateBulk) Exec(ctx context.Context) error {
	_, err := cjcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjcb *CandidateJobCreateBulk) ExecX(ctx context.Context) {
	if err := cjcb.Exec(ctx); err != nil {
		panic(err)
	}
}

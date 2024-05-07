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
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateJobUpdate is the builder for updating CandidateJob entities.
type CandidateJobUpdate struct {
	config
	hooks    []Hook
	mutation *CandidateJobMutation
}

// Where appends a list predicates to the CandidateJobUpdate builder.
func (cju *CandidateJobUpdate) Where(ps ...predicate.CandidateJob) *CandidateJobUpdate {
	cju.mutation.Where(ps...)
	return cju
}

// SetUpdatedAt sets the "updated_at" field.
func (cju *CandidateJobUpdate) SetUpdatedAt(t time.Time) *CandidateJobUpdate {
	cju.mutation.SetUpdatedAt(t)
	return cju
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableUpdatedAt(t *time.Time) *CandidateJobUpdate {
	if t != nil {
		cju.SetUpdatedAt(*t)
	}
	return cju
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cju *CandidateJobUpdate) ClearUpdatedAt() *CandidateJobUpdate {
	cju.mutation.ClearUpdatedAt()
	return cju
}

// SetDeletedAt sets the "deleted_at" field.
func (cju *CandidateJobUpdate) SetDeletedAt(t time.Time) *CandidateJobUpdate {
	cju.mutation.SetDeletedAt(t)
	return cju
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableDeletedAt(t *time.Time) *CandidateJobUpdate {
	if t != nil {
		cju.SetDeletedAt(*t)
	}
	return cju
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cju *CandidateJobUpdate) ClearDeletedAt() *CandidateJobUpdate {
	cju.mutation.ClearDeletedAt()
	return cju
}

// SetHiringJobID sets the "hiring_job_id" field.
func (cju *CandidateJobUpdate) SetHiringJobID(u uuid.UUID) *CandidateJobUpdate {
	cju.mutation.SetHiringJobID(u)
	return cju
}

// SetNillableHiringJobID sets the "hiring_job_id" field if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableHiringJobID(u *uuid.UUID) *CandidateJobUpdate {
	if u != nil {
		cju.SetHiringJobID(*u)
	}
	return cju
}

// ClearHiringJobID clears the value of the "hiring_job_id" field.
func (cju *CandidateJobUpdate) ClearHiringJobID() *CandidateJobUpdate {
	cju.mutation.ClearHiringJobID()
	return cju
}

// SetCandidateID sets the "candidate_id" field.
func (cju *CandidateJobUpdate) SetCandidateID(u uuid.UUID) *CandidateJobUpdate {
	cju.mutation.SetCandidateID(u)
	return cju
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableCandidateID(u *uuid.UUID) *CandidateJobUpdate {
	if u != nil {
		cju.SetCandidateID(*u)
	}
	return cju
}

// ClearCandidateID clears the value of the "candidate_id" field.
func (cju *CandidateJobUpdate) ClearCandidateID() *CandidateJobUpdate {
	cju.mutation.ClearCandidateID()
	return cju
}

// SetStatus sets the "status" field.
func (cju *CandidateJobUpdate) SetStatus(c candidatejob.Status) *CandidateJobUpdate {
	cju.mutation.SetStatus(c)
	return cju
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableStatus(c *candidatejob.Status) *CandidateJobUpdate {
	if c != nil {
		cju.SetStatus(*c)
	}
	return cju
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cju *CandidateJobUpdate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.AddAttachmentEdgeIDs(ids...)
	return cju
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cju *CandidateJobUpdate) AddAttachmentEdges(a ...*Attachment) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cju.AddAttachmentEdgeIDs(ids...)
}

// SetHiringJob sets the "hiring_job" edge to the HiringJob entity.
func (cju *CandidateJobUpdate) SetHiringJob(h *HiringJob) *CandidateJobUpdate {
	return cju.SetHiringJobID(h.ID)
}

// AddCandidateJobFeedbackIDs adds the "candidate_job_feedback" edge to the CandidateJobFeedback entity by IDs.
func (cju *CandidateJobUpdate) AddCandidateJobFeedbackIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.AddCandidateJobFeedbackIDs(ids...)
	return cju
}

// AddCandidateJobFeedback adds the "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (cju *CandidateJobUpdate) AddCandidateJobFeedback(c ...*CandidateJobFeedback) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cju.AddCandidateJobFeedbackIDs(ids...)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cju *CandidateJobUpdate) SetCandidateEdgeID(id uuid.UUID) *CandidateJobUpdate {
	cju.mutation.SetCandidateEdgeID(id)
	return cju
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cju *CandidateJobUpdate) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateJobUpdate {
	if id != nil {
		cju = cju.SetCandidateEdgeID(*id)
	}
	return cju
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cju *CandidateJobUpdate) SetCandidateEdge(c *Candidate) *CandidateJobUpdate {
	return cju.SetCandidateEdgeID(c.ID)
}

// AddCandidateJobInterviewIDs adds the "candidate_job_interview" edge to the CandidateInterview entity by IDs.
func (cju *CandidateJobUpdate) AddCandidateJobInterviewIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.AddCandidateJobInterviewIDs(ids...)
	return cju
}

// AddCandidateJobInterview adds the "candidate_job_interview" edges to the CandidateInterview entity.
func (cju *CandidateJobUpdate) AddCandidateJobInterview(c ...*CandidateInterview) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cju.AddCandidateJobInterviewIDs(ids...)
}

// Mutation returns the CandidateJobMutation object of the builder.
func (cju *CandidateJobUpdate) Mutation() *CandidateJobMutation {
	return cju.mutation
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cju *CandidateJobUpdate) ClearAttachmentEdges() *CandidateJobUpdate {
	cju.mutation.ClearAttachmentEdges()
	return cju
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cju *CandidateJobUpdate) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cju
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cju *CandidateJobUpdate) RemoveAttachmentEdges(a ...*Attachment) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cju.RemoveAttachmentEdgeIDs(ids...)
}

// ClearHiringJob clears the "hiring_job" edge to the HiringJob entity.
func (cju *CandidateJobUpdate) ClearHiringJob() *CandidateJobUpdate {
	cju.mutation.ClearHiringJob()
	return cju
}

// ClearCandidateJobFeedback clears all "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (cju *CandidateJobUpdate) ClearCandidateJobFeedback() *CandidateJobUpdate {
	cju.mutation.ClearCandidateJobFeedback()
	return cju
}

// RemoveCandidateJobFeedbackIDs removes the "candidate_job_feedback" edge to CandidateJobFeedback entities by IDs.
func (cju *CandidateJobUpdate) RemoveCandidateJobFeedbackIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.RemoveCandidateJobFeedbackIDs(ids...)
	return cju
}

// RemoveCandidateJobFeedback removes "candidate_job_feedback" edges to CandidateJobFeedback entities.
func (cju *CandidateJobUpdate) RemoveCandidateJobFeedback(c ...*CandidateJobFeedback) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cju.RemoveCandidateJobFeedbackIDs(ids...)
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (cju *CandidateJobUpdate) ClearCandidateEdge() *CandidateJobUpdate {
	cju.mutation.ClearCandidateEdge()
	return cju
}

// ClearCandidateJobInterview clears all "candidate_job_interview" edges to the CandidateInterview entity.
func (cju *CandidateJobUpdate) ClearCandidateJobInterview() *CandidateJobUpdate {
	cju.mutation.ClearCandidateJobInterview()
	return cju
}

// RemoveCandidateJobInterviewIDs removes the "candidate_job_interview" edge to CandidateInterview entities by IDs.
func (cju *CandidateJobUpdate) RemoveCandidateJobInterviewIDs(ids ...uuid.UUID) *CandidateJobUpdate {
	cju.mutation.RemoveCandidateJobInterviewIDs(ids...)
	return cju
}

// RemoveCandidateJobInterview removes "candidate_job_interview" edges to CandidateInterview entities.
func (cju *CandidateJobUpdate) RemoveCandidateJobInterview(c ...*CandidateInterview) *CandidateJobUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cju.RemoveCandidateJobInterviewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cju *CandidateJobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cju.hooks) == 0 {
		if err = cju.check(); err != nil {
			return 0, err
		}
		affected, err = cju.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cju.check(); err != nil {
				return 0, err
			}
			cju.mutation = mutation
			affected, err = cju.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cju.hooks) - 1; i >= 0; i-- {
			if cju.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cju.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cju.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cju *CandidateJobUpdate) SaveX(ctx context.Context) int {
	affected, err := cju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cju *CandidateJobUpdate) Exec(ctx context.Context) error {
	_, err := cju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cju *CandidateJobUpdate) ExecX(ctx context.Context) {
	if err := cju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cju *CandidateJobUpdate) check() error {
	if v, ok := cju.mutation.Status(); ok {
		if err := candidatejob.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CandidateJob.status": %w`, err)}
		}
	}
	return nil
}

func (cju *CandidateJobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatejob.Table,
			Columns: candidatejob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejob.FieldID,
			},
		},
	}
	if ps := cju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cju.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejob.FieldUpdatedAt, field.TypeTime, value)
	}
	if cju.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidatejob.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cju.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejob.FieldDeletedAt, field.TypeTime, value)
	}
	if cju.mutation.DeletedAtCleared() {
		_spec.ClearField(candidatejob.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cju.mutation.Status(); ok {
		_spec.SetField(candidatejob.FieldStatus, field.TypeEnum, value)
	}
	if cju.mutation.AttachmentEdgesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cju.mutation.AttachmentEdgesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cju.mutation.HiringJobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.HiringJobTable,
			Columns: []string{candidatejob.HiringJobColumn},
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
	if nodes := cju.mutation.HiringJobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.HiringJobTable,
			Columns: []string{candidatejob.HiringJobColumn},
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
	if cju.mutation.CandidateJobFeedbackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.RemovedCandidateJobFeedbackIDs(); len(nodes) > 0 && !cju.mutation.CandidateJobFeedbackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.CandidateJobFeedbackIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cju.mutation.CandidateEdgeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cju.mutation.CandidateJobInterviewCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.RemovedCandidateJobInterviewIDs(); len(nodes) > 0 && !cju.mutation.CandidateJobInterviewCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cju.mutation.CandidateJobInterviewIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidatejob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CandidateJobUpdateOne is the builder for updating a single CandidateJob entity.
type CandidateJobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CandidateJobMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cjuo *CandidateJobUpdateOne) SetUpdatedAt(t time.Time) *CandidateJobUpdateOne {
	cjuo.mutation.SetUpdatedAt(t)
	return cjuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableUpdatedAt(t *time.Time) *CandidateJobUpdateOne {
	if t != nil {
		cjuo.SetUpdatedAt(*t)
	}
	return cjuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cjuo *CandidateJobUpdateOne) ClearUpdatedAt() *CandidateJobUpdateOne {
	cjuo.mutation.ClearUpdatedAt()
	return cjuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cjuo *CandidateJobUpdateOne) SetDeletedAt(t time.Time) *CandidateJobUpdateOne {
	cjuo.mutation.SetDeletedAt(t)
	return cjuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableDeletedAt(t *time.Time) *CandidateJobUpdateOne {
	if t != nil {
		cjuo.SetDeletedAt(*t)
	}
	return cjuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cjuo *CandidateJobUpdateOne) ClearDeletedAt() *CandidateJobUpdateOne {
	cjuo.mutation.ClearDeletedAt()
	return cjuo
}

// SetHiringJobID sets the "hiring_job_id" field.
func (cjuo *CandidateJobUpdateOne) SetHiringJobID(u uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.SetHiringJobID(u)
	return cjuo
}

// SetNillableHiringJobID sets the "hiring_job_id" field if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableHiringJobID(u *uuid.UUID) *CandidateJobUpdateOne {
	if u != nil {
		cjuo.SetHiringJobID(*u)
	}
	return cjuo
}

// ClearHiringJobID clears the value of the "hiring_job_id" field.
func (cjuo *CandidateJobUpdateOne) ClearHiringJobID() *CandidateJobUpdateOne {
	cjuo.mutation.ClearHiringJobID()
	return cjuo
}

// SetCandidateID sets the "candidate_id" field.
func (cjuo *CandidateJobUpdateOne) SetCandidateID(u uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.SetCandidateID(u)
	return cjuo
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableCandidateID(u *uuid.UUID) *CandidateJobUpdateOne {
	if u != nil {
		cjuo.SetCandidateID(*u)
	}
	return cjuo
}

// ClearCandidateID clears the value of the "candidate_id" field.
func (cjuo *CandidateJobUpdateOne) ClearCandidateID() *CandidateJobUpdateOne {
	cjuo.mutation.ClearCandidateID()
	return cjuo
}

// SetStatus sets the "status" field.
func (cjuo *CandidateJobUpdateOne) SetStatus(c candidatejob.Status) *CandidateJobUpdateOne {
	cjuo.mutation.SetStatus(c)
	return cjuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableStatus(c *candidatejob.Status) *CandidateJobUpdateOne {
	if c != nil {
		cjuo.SetStatus(*c)
	}
	return cjuo
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cjuo *CandidateJobUpdateOne) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.AddAttachmentEdgeIDs(ids...)
	return cjuo
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cjuo *CandidateJobUpdateOne) AddAttachmentEdges(a ...*Attachment) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjuo.AddAttachmentEdgeIDs(ids...)
}

// SetHiringJob sets the "hiring_job" edge to the HiringJob entity.
func (cjuo *CandidateJobUpdateOne) SetHiringJob(h *HiringJob) *CandidateJobUpdateOne {
	return cjuo.SetHiringJobID(h.ID)
}

// AddCandidateJobFeedbackIDs adds the "candidate_job_feedback" edge to the CandidateJobFeedback entity by IDs.
func (cjuo *CandidateJobUpdateOne) AddCandidateJobFeedbackIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.AddCandidateJobFeedbackIDs(ids...)
	return cjuo
}

// AddCandidateJobFeedback adds the "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (cjuo *CandidateJobUpdateOne) AddCandidateJobFeedback(c ...*CandidateJobFeedback) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjuo.AddCandidateJobFeedbackIDs(ids...)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cjuo *CandidateJobUpdateOne) SetCandidateEdgeID(id uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.SetCandidateEdgeID(id)
	return cjuo
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cjuo *CandidateJobUpdateOne) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateJobUpdateOne {
	if id != nil {
		cjuo = cjuo.SetCandidateEdgeID(*id)
	}
	return cjuo
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cjuo *CandidateJobUpdateOne) SetCandidateEdge(c *Candidate) *CandidateJobUpdateOne {
	return cjuo.SetCandidateEdgeID(c.ID)
}

// AddCandidateJobInterviewIDs adds the "candidate_job_interview" edge to the CandidateInterview entity by IDs.
func (cjuo *CandidateJobUpdateOne) AddCandidateJobInterviewIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.AddCandidateJobInterviewIDs(ids...)
	return cjuo
}

// AddCandidateJobInterview adds the "candidate_job_interview" edges to the CandidateInterview entity.
func (cjuo *CandidateJobUpdateOne) AddCandidateJobInterview(c ...*CandidateInterview) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjuo.AddCandidateJobInterviewIDs(ids...)
}

// Mutation returns the CandidateJobMutation object of the builder.
func (cjuo *CandidateJobUpdateOne) Mutation() *CandidateJobMutation {
	return cjuo.mutation
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cjuo *CandidateJobUpdateOne) ClearAttachmentEdges() *CandidateJobUpdateOne {
	cjuo.mutation.ClearAttachmentEdges()
	return cjuo
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cjuo *CandidateJobUpdateOne) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cjuo
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cjuo *CandidateJobUpdateOne) RemoveAttachmentEdges(a ...*Attachment) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjuo.RemoveAttachmentEdgeIDs(ids...)
}

// ClearHiringJob clears the "hiring_job" edge to the HiringJob entity.
func (cjuo *CandidateJobUpdateOne) ClearHiringJob() *CandidateJobUpdateOne {
	cjuo.mutation.ClearHiringJob()
	return cjuo
}

// ClearCandidateJobFeedback clears all "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (cjuo *CandidateJobUpdateOne) ClearCandidateJobFeedback() *CandidateJobUpdateOne {
	cjuo.mutation.ClearCandidateJobFeedback()
	return cjuo
}

// RemoveCandidateJobFeedbackIDs removes the "candidate_job_feedback" edge to CandidateJobFeedback entities by IDs.
func (cjuo *CandidateJobUpdateOne) RemoveCandidateJobFeedbackIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.RemoveCandidateJobFeedbackIDs(ids...)
	return cjuo
}

// RemoveCandidateJobFeedback removes "candidate_job_feedback" edges to CandidateJobFeedback entities.
func (cjuo *CandidateJobUpdateOne) RemoveCandidateJobFeedback(c ...*CandidateJobFeedback) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjuo.RemoveCandidateJobFeedbackIDs(ids...)
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (cjuo *CandidateJobUpdateOne) ClearCandidateEdge() *CandidateJobUpdateOne {
	cjuo.mutation.ClearCandidateEdge()
	return cjuo
}

// ClearCandidateJobInterview clears all "candidate_job_interview" edges to the CandidateInterview entity.
func (cjuo *CandidateJobUpdateOne) ClearCandidateJobInterview() *CandidateJobUpdateOne {
	cjuo.mutation.ClearCandidateJobInterview()
	return cjuo
}

// RemoveCandidateJobInterviewIDs removes the "candidate_job_interview" edge to CandidateInterview entities by IDs.
func (cjuo *CandidateJobUpdateOne) RemoveCandidateJobInterviewIDs(ids ...uuid.UUID) *CandidateJobUpdateOne {
	cjuo.mutation.RemoveCandidateJobInterviewIDs(ids...)
	return cjuo
}

// RemoveCandidateJobInterview removes "candidate_job_interview" edges to CandidateInterview entities.
func (cjuo *CandidateJobUpdateOne) RemoveCandidateJobInterview(c ...*CandidateInterview) *CandidateJobUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cjuo.RemoveCandidateJobInterviewIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cjuo *CandidateJobUpdateOne) Select(field string, fields ...string) *CandidateJobUpdateOne {
	cjuo.fields = append([]string{field}, fields...)
	return cjuo
}

// Save executes the query and returns the updated CandidateJob entity.
func (cjuo *CandidateJobUpdateOne) Save(ctx context.Context) (*CandidateJob, error) {
	var (
		err  error
		node *CandidateJob
	)
	if len(cjuo.hooks) == 0 {
		if err = cjuo.check(); err != nil {
			return nil, err
		}
		node, err = cjuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjuo.check(); err != nil {
				return nil, err
			}
			cjuo.mutation = mutation
			node, err = cjuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cjuo.hooks) - 1; i >= 0; i-- {
			if cjuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cjuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (cjuo *CandidateJobUpdateOne) SaveX(ctx context.Context) *CandidateJob {
	node, err := cjuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cjuo *CandidateJobUpdateOne) Exec(ctx context.Context) error {
	_, err := cjuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjuo *CandidateJobUpdateOne) ExecX(ctx context.Context) {
	if err := cjuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjuo *CandidateJobUpdateOne) check() error {
	if v, ok := cjuo.mutation.Status(); ok {
		if err := candidatejob.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CandidateJob.status": %w`, err)}
		}
	}
	return nil
}

func (cjuo *CandidateJobUpdateOne) sqlSave(ctx context.Context) (_node *CandidateJob, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatejob.Table,
			Columns: candidatejob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejob.FieldID,
			},
		},
	}
	id, ok := cjuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CandidateJob.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cjuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidatejob.FieldID)
		for _, f := range fields {
			if !candidatejob.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != candidatejob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cjuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cjuo.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejob.FieldUpdatedAt, field.TypeTime, value)
	}
	if cjuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidatejob.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cjuo.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejob.FieldDeletedAt, field.TypeTime, value)
	}
	if cjuo.mutation.DeletedAtCleared() {
		_spec.ClearField(candidatejob.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cjuo.mutation.Status(); ok {
		_spec.SetField(candidatejob.FieldStatus, field.TypeEnum, value)
	}
	if cjuo.mutation.AttachmentEdgesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cjuo.mutation.AttachmentEdgesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cjuo.mutation.HiringJobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.HiringJobTable,
			Columns: []string{candidatejob.HiringJobColumn},
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
	if nodes := cjuo.mutation.HiringJobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejob.HiringJobTable,
			Columns: []string{candidatejob.HiringJobColumn},
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
	if cjuo.mutation.CandidateJobFeedbackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.RemovedCandidateJobFeedbackIDs(); len(nodes) > 0 && !cjuo.mutation.CandidateJobFeedbackCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.CandidateJobFeedbackIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cjuo.mutation.CandidateEdgeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cjuo.mutation.CandidateJobInterviewCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.RemovedCandidateJobInterviewIDs(); len(nodes) > 0 && !cjuo.mutation.CandidateJobInterviewCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjuo.mutation.CandidateJobInterviewIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CandidateJob{config: cjuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cjuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidatejob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
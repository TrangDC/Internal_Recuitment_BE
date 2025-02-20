// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateJobFeedbackUpdate is the builder for updating CandidateJobFeedback entities.
type CandidateJobFeedbackUpdate struct {
	config
	hooks    []Hook
	mutation *CandidateJobFeedbackMutation
}

// Where appends a list predicates to the CandidateJobFeedbackUpdate builder.
func (cjfu *CandidateJobFeedbackUpdate) Where(ps ...predicate.CandidateJobFeedback) *CandidateJobFeedbackUpdate {
	cjfu.mutation.Where(ps...)
	return cjfu
}

// SetUpdatedAt sets the "updated_at" field.
func (cjfu *CandidateJobFeedbackUpdate) SetUpdatedAt(t time.Time) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetUpdatedAt(t)
	return cjfu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableUpdatedAt(t *time.Time) *CandidateJobFeedbackUpdate {
	if t != nil {
		cjfu.SetUpdatedAt(*t)
	}
	return cjfu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cjfu *CandidateJobFeedbackUpdate) ClearUpdatedAt() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearUpdatedAt()
	return cjfu
}

// SetDeletedAt sets the "deleted_at" field.
func (cjfu *CandidateJobFeedbackUpdate) SetDeletedAt(t time.Time) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetDeletedAt(t)
	return cjfu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableDeletedAt(t *time.Time) *CandidateJobFeedbackUpdate {
	if t != nil {
		cjfu.SetDeletedAt(*t)
	}
	return cjfu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cjfu *CandidateJobFeedbackUpdate) ClearDeletedAt() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearDeletedAt()
	return cjfu
}

// SetCandidateJobID sets the "candidate_job_id" field.
func (cjfu *CandidateJobFeedbackUpdate) SetCandidateJobID(u uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetCandidateJobID(u)
	return cjfu
}

// SetNillableCandidateJobID sets the "candidate_job_id" field if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableCandidateJobID(u *uuid.UUID) *CandidateJobFeedbackUpdate {
	if u != nil {
		cjfu.SetCandidateJobID(*u)
	}
	return cjfu
}

// ClearCandidateJobID clears the value of the "candidate_job_id" field.
func (cjfu *CandidateJobFeedbackUpdate) ClearCandidateJobID() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearCandidateJobID()
	return cjfu
}

// SetCreatedBy sets the "created_by" field.
func (cjfu *CandidateJobFeedbackUpdate) SetCreatedBy(u uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetCreatedBy(u)
	return cjfu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableCreatedBy(u *uuid.UUID) *CandidateJobFeedbackUpdate {
	if u != nil {
		cjfu.SetCreatedBy(*u)
	}
	return cjfu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cjfu *CandidateJobFeedbackUpdate) ClearCreatedBy() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearCreatedBy()
	return cjfu
}

// SetCandidateJobStatus sets the "candidate_job_status" field.
func (cjfu *CandidateJobFeedbackUpdate) SetCandidateJobStatus(cjs candidatejobfeedback.CandidateJobStatus) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetCandidateJobStatus(cjs)
	return cjfu
}

// SetNillableCandidateJobStatus sets the "candidate_job_status" field if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableCandidateJobStatus(cjs *candidatejobfeedback.CandidateJobStatus) *CandidateJobFeedbackUpdate {
	if cjs != nil {
		cjfu.SetCandidateJobStatus(*cjs)
	}
	return cjfu
}

// SetFeedback sets the "feedback" field.
func (cjfu *CandidateJobFeedbackUpdate) SetFeedback(s string) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetFeedback(s)
	return cjfu
}

// SetCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID.
func (cjfu *CandidateJobFeedbackUpdate) SetCreatedByEdgeID(id uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetCreatedByEdgeID(id)
	return cjfu
}

// SetNillableCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableCreatedByEdgeID(id *uuid.UUID) *CandidateJobFeedbackUpdate {
	if id != nil {
		cjfu = cjfu.SetCreatedByEdgeID(*id)
	}
	return cjfu
}

// SetCreatedByEdge sets the "created_by_edge" edge to the User entity.
func (cjfu *CandidateJobFeedbackUpdate) SetCreatedByEdge(u *User) *CandidateJobFeedbackUpdate {
	return cjfu.SetCreatedByEdgeID(u.ID)
}

// SetCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID.
func (cjfu *CandidateJobFeedbackUpdate) SetCandidateJobEdgeID(id uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.SetCandidateJobEdgeID(id)
	return cjfu
}

// SetNillableCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID if the given value is not nil.
func (cjfu *CandidateJobFeedbackUpdate) SetNillableCandidateJobEdgeID(id *uuid.UUID) *CandidateJobFeedbackUpdate {
	if id != nil {
		cjfu = cjfu.SetCandidateJobEdgeID(*id)
	}
	return cjfu
}

// SetCandidateJobEdge sets the "candidate_job_edge" edge to the CandidateJob entity.
func (cjfu *CandidateJobFeedbackUpdate) SetCandidateJobEdge(c *CandidateJob) *CandidateJobFeedbackUpdate {
	return cjfu.SetCandidateJobEdgeID(c.ID)
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cjfu *CandidateJobFeedbackUpdate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.AddAttachmentEdgeIDs(ids...)
	return cjfu
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cjfu *CandidateJobFeedbackUpdate) AddAttachmentEdges(a ...*Attachment) *CandidateJobFeedbackUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjfu.AddAttachmentEdgeIDs(ids...)
}

// Mutation returns the CandidateJobFeedbackMutation object of the builder.
func (cjfu *CandidateJobFeedbackUpdate) Mutation() *CandidateJobFeedbackMutation {
	return cjfu.mutation
}

// ClearCreatedByEdge clears the "created_by_edge" edge to the User entity.
func (cjfu *CandidateJobFeedbackUpdate) ClearCreatedByEdge() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearCreatedByEdge()
	return cjfu
}

// ClearCandidateJobEdge clears the "candidate_job_edge" edge to the CandidateJob entity.
func (cjfu *CandidateJobFeedbackUpdate) ClearCandidateJobEdge() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearCandidateJobEdge()
	return cjfu
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cjfu *CandidateJobFeedbackUpdate) ClearAttachmentEdges() *CandidateJobFeedbackUpdate {
	cjfu.mutation.ClearAttachmentEdges()
	return cjfu
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cjfu *CandidateJobFeedbackUpdate) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobFeedbackUpdate {
	cjfu.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cjfu
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cjfu *CandidateJobFeedbackUpdate) RemoveAttachmentEdges(a ...*Attachment) *CandidateJobFeedbackUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjfu.RemoveAttachmentEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cjfu *CandidateJobFeedbackUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cjfu.hooks) == 0 {
		if err = cjfu.check(); err != nil {
			return 0, err
		}
		affected, err = cjfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobFeedbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjfu.check(); err != nil {
				return 0, err
			}
			cjfu.mutation = mutation
			affected, err = cjfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cjfu.hooks) - 1; i >= 0; i-- {
			if cjfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cjfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cjfu *CandidateJobFeedbackUpdate) SaveX(ctx context.Context) int {
	affected, err := cjfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cjfu *CandidateJobFeedbackUpdate) Exec(ctx context.Context) error {
	_, err := cjfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjfu *CandidateJobFeedbackUpdate) ExecX(ctx context.Context) {
	if err := cjfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjfu *CandidateJobFeedbackUpdate) check() error {
	if v, ok := cjfu.mutation.CandidateJobStatus(); ok {
		if err := candidatejobfeedback.CandidateJobStatusValidator(v); err != nil {
			return &ValidationError{Name: "candidate_job_status", err: fmt.Errorf(`ent: validator failed for field "CandidateJobFeedback.candidate_job_status": %w`, err)}
		}
	}
	return nil
}

func (cjfu *CandidateJobFeedbackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatejobfeedback.Table,
			Columns: candidatejobfeedback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejobfeedback.FieldID,
			},
		},
	}
	if ps := cjfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cjfu.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejobfeedback.FieldUpdatedAt, field.TypeTime, value)
	}
	if cjfu.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidatejobfeedback.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cjfu.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejobfeedback.FieldDeletedAt, field.TypeTime, value)
	}
	if cjfu.mutation.DeletedAtCleared() {
		_spec.ClearField(candidatejobfeedback.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cjfu.mutation.CandidateJobStatus(); ok {
		_spec.SetField(candidatejobfeedback.FieldCandidateJobStatus, field.TypeEnum, value)
	}
	if value, ok := cjfu.mutation.Feedback(); ok {
		_spec.SetField(candidatejobfeedback.FieldFeedback, field.TypeString, value)
	}
	if cjfu.mutation.CreatedByEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CreatedByEdgeTable,
			Columns: []string{candidatejobfeedback.CreatedByEdgeColumn},
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
	if nodes := cjfu.mutation.CreatedByEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CreatedByEdgeTable,
			Columns: []string{candidatejobfeedback.CreatedByEdgeColumn},
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
	if cjfu.mutation.CandidateJobEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CandidateJobEdgeTable,
			Columns: []string{candidatejobfeedback.CandidateJobEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjfu.mutation.CandidateJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CandidateJobEdgeTable,
			Columns: []string{candidatejobfeedback.CandidateJobEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cjfu.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	if nodes := cjfu.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cjfu.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	if nodes := cjfu.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cjfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidatejobfeedback.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CandidateJobFeedbackUpdateOne is the builder for updating a single CandidateJobFeedback entity.
type CandidateJobFeedbackUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CandidateJobFeedbackMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetUpdatedAt(t time.Time) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetUpdatedAt(t)
	return cjfuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableUpdatedAt(t *time.Time) *CandidateJobFeedbackUpdateOne {
	if t != nil {
		cjfuo.SetUpdatedAt(*t)
	}
	return cjfuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearUpdatedAt() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearUpdatedAt()
	return cjfuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetDeletedAt(t time.Time) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetDeletedAt(t)
	return cjfuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableDeletedAt(t *time.Time) *CandidateJobFeedbackUpdateOne {
	if t != nil {
		cjfuo.SetDeletedAt(*t)
	}
	return cjfuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearDeletedAt() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearDeletedAt()
	return cjfuo
}

// SetCandidateJobID sets the "candidate_job_id" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCandidateJobID(u uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetCandidateJobID(u)
	return cjfuo
}

// SetNillableCandidateJobID sets the "candidate_job_id" field if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableCandidateJobID(u *uuid.UUID) *CandidateJobFeedbackUpdateOne {
	if u != nil {
		cjfuo.SetCandidateJobID(*u)
	}
	return cjfuo
}

// ClearCandidateJobID clears the value of the "candidate_job_id" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearCandidateJobID() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearCandidateJobID()
	return cjfuo
}

// SetCreatedBy sets the "created_by" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCreatedBy(u uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetCreatedBy(u)
	return cjfuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *CandidateJobFeedbackUpdateOne {
	if u != nil {
		cjfuo.SetCreatedBy(*u)
	}
	return cjfuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearCreatedBy() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearCreatedBy()
	return cjfuo
}

// SetCandidateJobStatus sets the "candidate_job_status" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCandidateJobStatus(cjs candidatejobfeedback.CandidateJobStatus) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetCandidateJobStatus(cjs)
	return cjfuo
}

// SetNillableCandidateJobStatus sets the "candidate_job_status" field if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableCandidateJobStatus(cjs *candidatejobfeedback.CandidateJobStatus) *CandidateJobFeedbackUpdateOne {
	if cjs != nil {
		cjfuo.SetCandidateJobStatus(*cjs)
	}
	return cjfuo
}

// SetFeedback sets the "feedback" field.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetFeedback(s string) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetFeedback(s)
	return cjfuo
}

// SetCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCreatedByEdgeID(id uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetCreatedByEdgeID(id)
	return cjfuo
}

// SetNillableCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableCreatedByEdgeID(id *uuid.UUID) *CandidateJobFeedbackUpdateOne {
	if id != nil {
		cjfuo = cjfuo.SetCreatedByEdgeID(*id)
	}
	return cjfuo
}

// SetCreatedByEdge sets the "created_by_edge" edge to the User entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCreatedByEdge(u *User) *CandidateJobFeedbackUpdateOne {
	return cjfuo.SetCreatedByEdgeID(u.ID)
}

// SetCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCandidateJobEdgeID(id uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.SetCandidateJobEdgeID(id)
	return cjfuo
}

// SetNillableCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID if the given value is not nil.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetNillableCandidateJobEdgeID(id *uuid.UUID) *CandidateJobFeedbackUpdateOne {
	if id != nil {
		cjfuo = cjfuo.SetCandidateJobEdgeID(*id)
	}
	return cjfuo
}

// SetCandidateJobEdge sets the "candidate_job_edge" edge to the CandidateJob entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) SetCandidateJobEdge(c *CandidateJob) *CandidateJobFeedbackUpdateOne {
	return cjfuo.SetCandidateJobEdgeID(c.ID)
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cjfuo *CandidateJobFeedbackUpdateOne) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.AddAttachmentEdgeIDs(ids...)
	return cjfuo
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) AddAttachmentEdges(a ...*Attachment) *CandidateJobFeedbackUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjfuo.AddAttachmentEdgeIDs(ids...)
}

// Mutation returns the CandidateJobFeedbackMutation object of the builder.
func (cjfuo *CandidateJobFeedbackUpdateOne) Mutation() *CandidateJobFeedbackMutation {
	return cjfuo.mutation
}

// ClearCreatedByEdge clears the "created_by_edge" edge to the User entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearCreatedByEdge() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearCreatedByEdge()
	return cjfuo
}

// ClearCandidateJobEdge clears the "candidate_job_edge" edge to the CandidateJob entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearCandidateJobEdge() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearCandidateJobEdge()
	return cjfuo
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) ClearAttachmentEdges() *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.ClearAttachmentEdges()
	return cjfuo
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cjfuo *CandidateJobFeedbackUpdateOne) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateJobFeedbackUpdateOne {
	cjfuo.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cjfuo
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cjfuo *CandidateJobFeedbackUpdateOne) RemoveAttachmentEdges(a ...*Attachment) *CandidateJobFeedbackUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cjfuo.RemoveAttachmentEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cjfuo *CandidateJobFeedbackUpdateOne) Select(field string, fields ...string) *CandidateJobFeedbackUpdateOne {
	cjfuo.fields = append([]string{field}, fields...)
	return cjfuo
}

// Save executes the query and returns the updated CandidateJobFeedback entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) Save(ctx context.Context) (*CandidateJobFeedback, error) {
	var (
		err  error
		node *CandidateJobFeedback
	)
	if len(cjfuo.hooks) == 0 {
		if err = cjfuo.check(); err != nil {
			return nil, err
		}
		node, err = cjfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobFeedbackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjfuo.check(); err != nil {
				return nil, err
			}
			cjfuo.mutation = mutation
			node, err = cjfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cjfuo.hooks) - 1; i >= 0; i-- {
			if cjfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cjfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateJobFeedback)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateJobFeedbackMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cjfuo *CandidateJobFeedbackUpdateOne) SaveX(ctx context.Context) *CandidateJobFeedback {
	node, err := cjfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cjfuo *CandidateJobFeedbackUpdateOne) Exec(ctx context.Context) error {
	_, err := cjfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjfuo *CandidateJobFeedbackUpdateOne) ExecX(ctx context.Context) {
	if err := cjfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjfuo *CandidateJobFeedbackUpdateOne) check() error {
	if v, ok := cjfuo.mutation.CandidateJobStatus(); ok {
		if err := candidatejobfeedback.CandidateJobStatusValidator(v); err != nil {
			return &ValidationError{Name: "candidate_job_status", err: fmt.Errorf(`ent: validator failed for field "CandidateJobFeedback.candidate_job_status": %w`, err)}
		}
	}
	return nil
}

func (cjfuo *CandidateJobFeedbackUpdateOne) sqlSave(ctx context.Context) (_node *CandidateJobFeedback, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidatejobfeedback.Table,
			Columns: candidatejobfeedback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejobfeedback.FieldID,
			},
		},
	}
	id, ok := cjfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CandidateJobFeedback.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cjfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidatejobfeedback.FieldID)
		for _, f := range fields {
			if !candidatejobfeedback.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != candidatejobfeedback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cjfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cjfuo.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejobfeedback.FieldUpdatedAt, field.TypeTime, value)
	}
	if cjfuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidatejobfeedback.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cjfuo.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejobfeedback.FieldDeletedAt, field.TypeTime, value)
	}
	if cjfuo.mutation.DeletedAtCleared() {
		_spec.ClearField(candidatejobfeedback.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cjfuo.mutation.CandidateJobStatus(); ok {
		_spec.SetField(candidatejobfeedback.FieldCandidateJobStatus, field.TypeEnum, value)
	}
	if value, ok := cjfuo.mutation.Feedback(); ok {
		_spec.SetField(candidatejobfeedback.FieldFeedback, field.TypeString, value)
	}
	if cjfuo.mutation.CreatedByEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CreatedByEdgeTable,
			Columns: []string{candidatejobfeedback.CreatedByEdgeColumn},
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
	if nodes := cjfuo.mutation.CreatedByEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CreatedByEdgeTable,
			Columns: []string{candidatejobfeedback.CreatedByEdgeColumn},
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
	if cjfuo.mutation.CandidateJobEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CandidateJobEdgeTable,
			Columns: []string{candidatejobfeedback.CandidateJobEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cjfuo.mutation.CandidateJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobfeedback.CandidateJobEdgeTable,
			Columns: []string{candidatejobfeedback.CandidateJobEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cjfuo.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	if nodes := cjfuo.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cjfuo.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	if nodes := cjfuo.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidatejobfeedback.AttachmentEdgesTable,
			Columns: []string{candidatejobfeedback.AttachmentEdgesColumn},
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
	_node = &CandidateJobFeedback{config: cjfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cjfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidatejobfeedback.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

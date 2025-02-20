// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateaward"
	"trec/ent/candidatecertificate"
	"trec/ent/candidateeducate"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatenote"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttachmentCreate) SetCreatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCreatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttachmentCreate) SetUpdatedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableUpdatedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AttachmentCreate) SetDeletedAt(t time.Time) *AttachmentCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableDeletedAt(t *time.Time) *AttachmentCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetDocumentID sets the "document_id" field.
func (ac *AttachmentCreate) SetDocumentID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetDocumentID(u)
	return ac
}

// SetDocumentName sets the "document_name" field.
func (ac *AttachmentCreate) SetDocumentName(s string) *AttachmentCreate {
	ac.mutation.SetDocumentName(s)
	return ac
}

// SetRelationType sets the "relation_type" field.
func (ac *AttachmentCreate) SetRelationType(at attachment.RelationType) *AttachmentCreate {
	ac.mutation.SetRelationType(at)
	return ac
}

// SetRelationID sets the "relation_id" field.
func (ac *AttachmentCreate) SetRelationID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetRelationID(u)
	return ac
}

// SetNillableRelationID sets the "relation_id" field if the given value is not nil.
func (ac *AttachmentCreate) SetNillableRelationID(u *uuid.UUID) *AttachmentCreate {
	if u != nil {
		ac.SetRelationID(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(u uuid.UUID) *AttachmentCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID.
func (ac *AttachmentCreate) SetCandidateJobEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateJobEdgeID(id)
	return ac
}

// SetNillableCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateJobEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateJobEdgeID(*id)
	}
	return ac
}

// SetCandidateJobEdge sets the "candidate_job_edge" edge to the CandidateJob entity.
func (ac *AttachmentCreate) SetCandidateJobEdge(c *CandidateJob) *AttachmentCreate {
	return ac.SetCandidateJobEdgeID(c.ID)
}

// SetCandidateJobFeedbackEdgeID sets the "candidate_job_feedback_edge" edge to the CandidateJobFeedback entity by ID.
func (ac *AttachmentCreate) SetCandidateJobFeedbackEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateJobFeedbackEdgeID(id)
	return ac
}

// SetNillableCandidateJobFeedbackEdgeID sets the "candidate_job_feedback_edge" edge to the CandidateJobFeedback entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateJobFeedbackEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateJobFeedbackEdgeID(*id)
	}
	return ac
}

// SetCandidateJobFeedbackEdge sets the "candidate_job_feedback_edge" edge to the CandidateJobFeedback entity.
func (ac *AttachmentCreate) SetCandidateJobFeedbackEdge(c *CandidateJobFeedback) *AttachmentCreate {
	return ac.SetCandidateJobFeedbackEdgeID(c.ID)
}

// SetCandidateInterviewEdgeID sets the "candidate_interview_edge" edge to the CandidateInterview entity by ID.
func (ac *AttachmentCreate) SetCandidateInterviewEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateInterviewEdgeID(id)
	return ac
}

// SetNillableCandidateInterviewEdgeID sets the "candidate_interview_edge" edge to the CandidateInterview entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateInterviewEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateInterviewEdgeID(*id)
	}
	return ac
}

// SetCandidateInterviewEdge sets the "candidate_interview_edge" edge to the CandidateInterview entity.
func (ac *AttachmentCreate) SetCandidateInterviewEdge(c *CandidateInterview) *AttachmentCreate {
	return ac.SetCandidateInterviewEdgeID(c.ID)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (ac *AttachmentCreate) SetCandidateEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateEdgeID(id)
	return ac
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateEdgeID(*id)
	}
	return ac
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (ac *AttachmentCreate) SetCandidateEdge(c *Candidate) *AttachmentCreate {
	return ac.SetCandidateEdgeID(c.ID)
}

// SetCandidateEducateEdgeID sets the "candidate_educate_edge" edge to the CandidateEducate entity by ID.
func (ac *AttachmentCreate) SetCandidateEducateEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateEducateEdgeID(id)
	return ac
}

// SetNillableCandidateEducateEdgeID sets the "candidate_educate_edge" edge to the CandidateEducate entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateEducateEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateEducateEdgeID(*id)
	}
	return ac
}

// SetCandidateEducateEdge sets the "candidate_educate_edge" edge to the CandidateEducate entity.
func (ac *AttachmentCreate) SetCandidateEducateEdge(c *CandidateEducate) *AttachmentCreate {
	return ac.SetCandidateEducateEdgeID(c.ID)
}

// SetCandidateAwardEdgeID sets the "candidate_award_edge" edge to the CandidateAward entity by ID.
func (ac *AttachmentCreate) SetCandidateAwardEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateAwardEdgeID(id)
	return ac
}

// SetNillableCandidateAwardEdgeID sets the "candidate_award_edge" edge to the CandidateAward entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateAwardEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateAwardEdgeID(*id)
	}
	return ac
}

// SetCandidateAwardEdge sets the "candidate_award_edge" edge to the CandidateAward entity.
func (ac *AttachmentCreate) SetCandidateAwardEdge(c *CandidateAward) *AttachmentCreate {
	return ac.SetCandidateAwardEdgeID(c.ID)
}

// SetCandidateCertificateEdgeID sets the "candidate_certificate_edge" edge to the CandidateCertificate entity by ID.
func (ac *AttachmentCreate) SetCandidateCertificateEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateCertificateEdgeID(id)
	return ac
}

// SetNillableCandidateCertificateEdgeID sets the "candidate_certificate_edge" edge to the CandidateCertificate entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateCertificateEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateCertificateEdgeID(*id)
	}
	return ac
}

// SetCandidateCertificateEdge sets the "candidate_certificate_edge" edge to the CandidateCertificate entity.
func (ac *AttachmentCreate) SetCandidateCertificateEdge(c *CandidateCertificate) *AttachmentCreate {
	return ac.SetCandidateCertificateEdgeID(c.ID)
}

// SetCandidateHistoryCallEdgeID sets the "candidate_history_call_edge" edge to the CandidateHistoryCall entity by ID.
func (ac *AttachmentCreate) SetCandidateHistoryCallEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateHistoryCallEdgeID(id)
	return ac
}

// SetNillableCandidateHistoryCallEdgeID sets the "candidate_history_call_edge" edge to the CandidateHistoryCall entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateHistoryCallEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateHistoryCallEdgeID(*id)
	}
	return ac
}

// SetCandidateHistoryCallEdge sets the "candidate_history_call_edge" edge to the CandidateHistoryCall entity.
func (ac *AttachmentCreate) SetCandidateHistoryCallEdge(c *CandidateHistoryCall) *AttachmentCreate {
	return ac.SetCandidateHistoryCallEdgeID(c.ID)
}

// SetCandidateNoteEdgeID sets the "candidate_note_edge" edge to the CandidateNote entity by ID.
func (ac *AttachmentCreate) SetCandidateNoteEdgeID(id uuid.UUID) *AttachmentCreate {
	ac.mutation.SetCandidateNoteEdgeID(id)
	return ac
}

// SetNillableCandidateNoteEdgeID sets the "candidate_note_edge" edge to the CandidateNote entity by ID if the given value is not nil.
func (ac *AttachmentCreate) SetNillableCandidateNoteEdgeID(id *uuid.UUID) *AttachmentCreate {
	if id != nil {
		ac = ac.SetCandidateNoteEdgeID(*id)
	}
	return ac
}

// SetCandidateNoteEdge sets the "candidate_note_edge" edge to the CandidateNote entity.
func (ac *AttachmentCreate) SetCandidateNoteEdge(c *CandidateNote) *AttachmentCreate {
	return ac.SetCandidateNoteEdgeID(c.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	var (
		err  error
		node *Attachment
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Attachment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AttachmentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttachmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttachmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttachmentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attachment.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Attachment.created_at"`)}
	}
	if _, ok := ac.mutation.DocumentID(); !ok {
		return &ValidationError{Name: "document_id", err: errors.New(`ent: missing required field "Attachment.document_id"`)}
	}
	if _, ok := ac.mutation.DocumentName(); !ok {
		return &ValidationError{Name: "document_name", err: errors.New(`ent: missing required field "Attachment.document_name"`)}
	}
	if v, ok := ac.mutation.DocumentName(); ok {
		if err := attachment.DocumentNameValidator(v); err != nil {
			return &ValidationError{Name: "document_name", err: fmt.Errorf(`ent: validator failed for field "Attachment.document_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.RelationType(); !ok {
		return &ValidationError{Name: "relation_type", err: errors.New(`ent: missing required field "Attachment.relation_type"`)}
	}
	if v, ok := ac.mutation.RelationType(); ok {
		if err := attachment.RelationTypeValidator(v); err != nil {
			return &ValidationError{Name: "relation_type", err: fmt.Errorf(`ent: validator failed for field "Attachment.relation_type": %w`, err)}
		}
	}
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(attachment.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.DocumentID(); ok {
		_spec.SetField(attachment.FieldDocumentID, field.TypeUUID, value)
		_node.DocumentID = value
	}
	if value, ok := ac.mutation.DocumentName(); ok {
		_spec.SetField(attachment.FieldDocumentName, field.TypeString, value)
		_node.DocumentName = value
	}
	if value, ok := ac.mutation.RelationType(); ok {
		_spec.SetField(attachment.FieldRelationType, field.TypeEnum, value)
		_node.RelationType = value
	}
	if nodes := ac.mutation.CandidateJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateJobEdgeTable,
			Columns: []string{attachment.CandidateJobEdgeColumn},
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
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateJobFeedbackEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateJobFeedbackEdgeTable,
			Columns: []string{attachment.CandidateJobFeedbackEdgeColumn},
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
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateInterviewEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateInterviewEdgeTable,
			Columns: []string{attachment.CandidateInterviewEdgeColumn},
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
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateEdgeTable,
			Columns: []string{attachment.CandidateEdgeColumn},
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
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateEducateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateEducateEdgeTable,
			Columns: []string{attachment.CandidateEducateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateeducate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateAwardEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateAwardEdgeTable,
			Columns: []string{attachment.CandidateAwardEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateaward.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateCertificateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateCertificateEdgeTable,
			Columns: []string{attachment.CandidateCertificateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatecertificate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateHistoryCallEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateHistoryCallEdgeTable,
			Columns: []string{attachment.CandidateHistoryCallEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatehistorycall.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CandidateNoteEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.CandidateNoteEdgeTable,
			Columns: []string{attachment.CandidateNoteEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatenote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	builders []*AttachmentCreate
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

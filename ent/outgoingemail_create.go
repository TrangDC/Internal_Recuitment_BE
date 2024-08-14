// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/candidate"
	"trec/ent/outgoingemail"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OutgoingEmailCreate is the builder for creating a OutgoingEmail entity.
type OutgoingEmailCreate struct {
	config
	mutation *OutgoingEmailMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oec *OutgoingEmailCreate) SetCreatedAt(t time.Time) *OutgoingEmailCreate {
	oec.mutation.SetCreatedAt(t)
	return oec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableCreatedAt(t *time.Time) *OutgoingEmailCreate {
	if t != nil {
		oec.SetCreatedAt(*t)
	}
	return oec
}

// SetUpdatedAt sets the "updated_at" field.
func (oec *OutgoingEmailCreate) SetUpdatedAt(t time.Time) *OutgoingEmailCreate {
	oec.mutation.SetUpdatedAt(t)
	return oec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableUpdatedAt(t *time.Time) *OutgoingEmailCreate {
	if t != nil {
		oec.SetUpdatedAt(*t)
	}
	return oec
}

// SetDeletedAt sets the "deleted_at" field.
func (oec *OutgoingEmailCreate) SetDeletedAt(t time.Time) *OutgoingEmailCreate {
	oec.mutation.SetDeletedAt(t)
	return oec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableDeletedAt(t *time.Time) *OutgoingEmailCreate {
	if t != nil {
		oec.SetDeletedAt(*t)
	}
	return oec
}

// SetTo sets the "to" field.
func (oec *OutgoingEmailCreate) SetTo(s []string) *OutgoingEmailCreate {
	oec.mutation.SetTo(s)
	return oec
}

// SetCc sets the "cc" field.
func (oec *OutgoingEmailCreate) SetCc(s []string) *OutgoingEmailCreate {
	oec.mutation.SetCc(s)
	return oec
}

// SetBcc sets the "bcc" field.
func (oec *OutgoingEmailCreate) SetBcc(s []string) *OutgoingEmailCreate {
	oec.mutation.SetBcc(s)
	return oec
}

// SetSubject sets the "subject" field.
func (oec *OutgoingEmailCreate) SetSubject(s string) *OutgoingEmailCreate {
	oec.mutation.SetSubject(s)
	return oec
}

// SetContent sets the "content" field.
func (oec *OutgoingEmailCreate) SetContent(s string) *OutgoingEmailCreate {
	oec.mutation.SetContent(s)
	return oec
}

// SetSignature sets the "signature" field.
func (oec *OutgoingEmailCreate) SetSignature(s string) *OutgoingEmailCreate {
	oec.mutation.SetSignature(s)
	return oec
}

// SetCandidateID sets the "candidate_id" field.
func (oec *OutgoingEmailCreate) SetCandidateID(u uuid.UUID) *OutgoingEmailCreate {
	oec.mutation.SetCandidateID(u)
	return oec
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableCandidateID(u *uuid.UUID) *OutgoingEmailCreate {
	if u != nil {
		oec.SetCandidateID(*u)
	}
	return oec
}

// SetRecipientType sets the "recipient_type" field.
func (oec *OutgoingEmailCreate) SetRecipientType(ot outgoingemail.RecipientType) *OutgoingEmailCreate {
	oec.mutation.SetRecipientType(ot)
	return oec
}

// SetEmailTemplateID sets the "email_template_id" field.
func (oec *OutgoingEmailCreate) SetEmailTemplateID(u uuid.UUID) *OutgoingEmailCreate {
	oec.mutation.SetEmailTemplateID(u)
	return oec
}

// SetNillableEmailTemplateID sets the "email_template_id" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableEmailTemplateID(u *uuid.UUID) *OutgoingEmailCreate {
	if u != nil {
		oec.SetEmailTemplateID(*u)
	}
	return oec
}

// SetStatus sets the "status" field.
func (oec *OutgoingEmailCreate) SetStatus(o outgoingemail.Status) *OutgoingEmailCreate {
	oec.mutation.SetStatus(o)
	return oec
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableStatus(o *outgoingemail.Status) *OutgoingEmailCreate {
	if o != nil {
		oec.SetStatus(*o)
	}
	return oec
}

// SetID sets the "id" field.
func (oec *OutgoingEmailCreate) SetID(u uuid.UUID) *OutgoingEmailCreate {
	oec.mutation.SetID(u)
	return oec
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (oec *OutgoingEmailCreate) SetCandidateEdgeID(id uuid.UUID) *OutgoingEmailCreate {
	oec.mutation.SetCandidateEdgeID(id)
	return oec
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (oec *OutgoingEmailCreate) SetNillableCandidateEdgeID(id *uuid.UUID) *OutgoingEmailCreate {
	if id != nil {
		oec = oec.SetCandidateEdgeID(*id)
	}
	return oec
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (oec *OutgoingEmailCreate) SetCandidateEdge(c *Candidate) *OutgoingEmailCreate {
	return oec.SetCandidateEdgeID(c.ID)
}

// Mutation returns the OutgoingEmailMutation object of the builder.
func (oec *OutgoingEmailCreate) Mutation() *OutgoingEmailMutation {
	return oec.mutation
}

// Save creates the OutgoingEmail in the database.
func (oec *OutgoingEmailCreate) Save(ctx context.Context) (*OutgoingEmail, error) {
	var (
		err  error
		node *OutgoingEmail
	)
	oec.defaults()
	if len(oec.hooks) == 0 {
		if err = oec.check(); err != nil {
			return nil, err
		}
		node, err = oec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutgoingEmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oec.check(); err != nil {
				return nil, err
			}
			oec.mutation = mutation
			if node, err = oec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oec.hooks) - 1; i >= 0; i-- {
			if oec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OutgoingEmail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OutgoingEmailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oec *OutgoingEmailCreate) SaveX(ctx context.Context) *OutgoingEmail {
	v, err := oec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oec *OutgoingEmailCreate) Exec(ctx context.Context) error {
	_, err := oec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oec *OutgoingEmailCreate) ExecX(ctx context.Context) {
	if err := oec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oec *OutgoingEmailCreate) defaults() {
	if _, ok := oec.mutation.CreatedAt(); !ok {
		v := outgoingemail.DefaultCreatedAt()
		oec.mutation.SetCreatedAt(v)
	}
	if _, ok := oec.mutation.Status(); !ok {
		v := outgoingemail.DefaultStatus
		oec.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oec *OutgoingEmailCreate) check() error {
	if _, ok := oec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OutgoingEmail.created_at"`)}
	}
	if _, ok := oec.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "OutgoingEmail.to"`)}
	}
	if _, ok := oec.mutation.Cc(); !ok {
		return &ValidationError{Name: "cc", err: errors.New(`ent: missing required field "OutgoingEmail.cc"`)}
	}
	if _, ok := oec.mutation.Bcc(); !ok {
		return &ValidationError{Name: "bcc", err: errors.New(`ent: missing required field "OutgoingEmail.bcc"`)}
	}
	if _, ok := oec.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "OutgoingEmail.subject"`)}
	}
	if v, ok := oec.mutation.Subject(); ok {
		if err := outgoingemail.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.subject": %w`, err)}
		}
	}
	if _, ok := oec.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "OutgoingEmail.content"`)}
	}
	if v, ok := oec.mutation.Content(); ok {
		if err := outgoingemail.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.content": %w`, err)}
		}
	}
	if _, ok := oec.mutation.Signature(); !ok {
		return &ValidationError{Name: "signature", err: errors.New(`ent: missing required field "OutgoingEmail.signature"`)}
	}
	if _, ok := oec.mutation.RecipientType(); !ok {
		return &ValidationError{Name: "recipient_type", err: errors.New(`ent: missing required field "OutgoingEmail.recipient_type"`)}
	}
	if v, ok := oec.mutation.RecipientType(); ok {
		if err := outgoingemail.RecipientTypeValidator(v); err != nil {
			return &ValidationError{Name: "recipient_type", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.recipient_type": %w`, err)}
		}
	}
	if _, ok := oec.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OutgoingEmail.status"`)}
	}
	if v, ok := oec.mutation.Status(); ok {
		if err := outgoingemail.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.status": %w`, err)}
		}
	}
	return nil
}

func (oec *OutgoingEmailCreate) sqlSave(ctx context.Context) (*OutgoingEmail, error) {
	_node, _spec := oec.createSpec()
	if err := sqlgraph.CreateNode(ctx, oec.driver, _spec); err != nil {
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

func (oec *OutgoingEmailCreate) createSpec() (*OutgoingEmail, *sqlgraph.CreateSpec) {
	var (
		_node = &OutgoingEmail{config: oec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outgoingemail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outgoingemail.FieldID,
			},
		}
	)
	if id, ok := oec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oec.mutation.CreatedAt(); ok {
		_spec.SetField(outgoingemail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oec.mutation.UpdatedAt(); ok {
		_spec.SetField(outgoingemail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oec.mutation.DeletedAt(); ok {
		_spec.SetField(outgoingemail.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oec.mutation.To(); ok {
		_spec.SetField(outgoingemail.FieldTo, field.TypeJSON, value)
		_node.To = value
	}
	if value, ok := oec.mutation.Cc(); ok {
		_spec.SetField(outgoingemail.FieldCc, field.TypeJSON, value)
		_node.Cc = value
	}
	if value, ok := oec.mutation.Bcc(); ok {
		_spec.SetField(outgoingemail.FieldBcc, field.TypeJSON, value)
		_node.Bcc = value
	}
	if value, ok := oec.mutation.Subject(); ok {
		_spec.SetField(outgoingemail.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := oec.mutation.Content(); ok {
		_spec.SetField(outgoingemail.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := oec.mutation.Signature(); ok {
		_spec.SetField(outgoingemail.FieldSignature, field.TypeString, value)
		_node.Signature = value
	}
	if value, ok := oec.mutation.RecipientType(); ok {
		_spec.SetField(outgoingemail.FieldRecipientType, field.TypeEnum, value)
		_node.RecipientType = value
	}
	if value, ok := oec.mutation.EmailTemplateID(); ok {
		_spec.SetField(outgoingemail.FieldEmailTemplateID, field.TypeUUID, value)
		_node.EmailTemplateID = value
	}
	if value, ok := oec.mutation.Status(); ok {
		_spec.SetField(outgoingemail.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := oec.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   outgoingemail.CandidateEdgeTable,
			Columns: []string{outgoingemail.CandidateEdgeColumn},
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
	return _node, _spec
}

// OutgoingEmailCreateBulk is the builder for creating many OutgoingEmail entities in bulk.
type OutgoingEmailCreateBulk struct {
	config
	builders []*OutgoingEmailCreate
}

// Save creates the OutgoingEmail entities in the database.
func (oecb *OutgoingEmailCreateBulk) Save(ctx context.Context) ([]*OutgoingEmail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oecb.builders))
	nodes := make([]*OutgoingEmail, len(oecb.builders))
	mutators := make([]Mutator, len(oecb.builders))
	for i := range oecb.builders {
		func(i int, root context.Context) {
			builder := oecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutgoingEmailMutation)
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
					_, err = mutators[i+1].Mutate(root, oecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oecb *OutgoingEmailCreateBulk) SaveX(ctx context.Context) []*OutgoingEmail {
	v, err := oecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oecb *OutgoingEmailCreateBulk) Exec(ctx context.Context) error {
	_, err := oecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oecb *OutgoingEmailCreateBulk) ExecX(ctx context.Context) {
	if err := oecb.Exec(ctx); err != nil {
		panic(err)
	}
}

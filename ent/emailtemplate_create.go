// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/emailevent"
	"trec/ent/emailroleattribute"
	"trec/ent/emailtemplate"
	"trec/ent/role"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailTemplateCreate is the builder for creating a EmailTemplate entity.
type EmailTemplateCreate struct {
	config
	mutation *EmailTemplateMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (etc *EmailTemplateCreate) SetCreatedAt(t time.Time) *EmailTemplateCreate {
	etc.mutation.SetCreatedAt(t)
	return etc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableCreatedAt(t *time.Time) *EmailTemplateCreate {
	if t != nil {
		etc.SetCreatedAt(*t)
	}
	return etc
}

// SetUpdatedAt sets the "updated_at" field.
func (etc *EmailTemplateCreate) SetUpdatedAt(t time.Time) *EmailTemplateCreate {
	etc.mutation.SetUpdatedAt(t)
	return etc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableUpdatedAt(t *time.Time) *EmailTemplateCreate {
	if t != nil {
		etc.SetUpdatedAt(*t)
	}
	return etc
}

// SetDeletedAt sets the "deleted_at" field.
func (etc *EmailTemplateCreate) SetDeletedAt(t time.Time) *EmailTemplateCreate {
	etc.mutation.SetDeletedAt(t)
	return etc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableDeletedAt(t *time.Time) *EmailTemplateCreate {
	if t != nil {
		etc.SetDeletedAt(*t)
	}
	return etc
}

// SetEvent sets the "event" field.
func (etc *EmailTemplateCreate) SetEvent(e emailtemplate.Event) *EmailTemplateCreate {
	etc.mutation.SetEvent(e)
	return etc
}

// SetNillableEvent sets the "event" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableEvent(e *emailtemplate.Event) *EmailTemplateCreate {
	if e != nil {
		etc.SetEvent(*e)
	}
	return etc
}

// SetSendTo sets the "send_to" field.
func (etc *EmailTemplateCreate) SetSendTo(s []string) *EmailTemplateCreate {
	etc.mutation.SetSendTo(s)
	return etc
}

// SetCc sets the "cc" field.
func (etc *EmailTemplateCreate) SetCc(s []string) *EmailTemplateCreate {
	etc.mutation.SetCc(s)
	return etc
}

// SetBcc sets the "bcc" field.
func (etc *EmailTemplateCreate) SetBcc(s []string) *EmailTemplateCreate {
	etc.mutation.SetBcc(s)
	return etc
}

// SetSubject sets the "subject" field.
func (etc *EmailTemplateCreate) SetSubject(s string) *EmailTemplateCreate {
	etc.mutation.SetSubject(s)
	return etc
}

// SetContent sets the "content" field.
func (etc *EmailTemplateCreate) SetContent(s string) *EmailTemplateCreate {
	etc.mutation.SetContent(s)
	return etc
}

// SetSignature sets the "signature" field.
func (etc *EmailTemplateCreate) SetSignature(s string) *EmailTemplateCreate {
	etc.mutation.SetSignature(s)
	return etc
}

// SetNillableSignature sets the "signature" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableSignature(s *string) *EmailTemplateCreate {
	if s != nil {
		etc.SetSignature(*s)
	}
	return etc
}

// SetStatus sets the "status" field.
func (etc *EmailTemplateCreate) SetStatus(e emailtemplate.Status) *EmailTemplateCreate {
	etc.mutation.SetStatus(e)
	return etc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (etc *EmailTemplateCreate) SetNillableStatus(e *emailtemplate.Status) *EmailTemplateCreate {
	if e != nil {
		etc.SetStatus(*e)
	}
	return etc
}

// SetEventID sets the "event_id" field.
func (etc *EmailTemplateCreate) SetEventID(u uuid.UUID) *EmailTemplateCreate {
	etc.mutation.SetEventID(u)
	return etc
}

// SetID sets the "id" field.
func (etc *EmailTemplateCreate) SetID(u uuid.UUID) *EmailTemplateCreate {
	etc.mutation.SetID(u)
	return etc
}

// AddRoleEdgeIDs adds the "role_edges" edge to the Role entity by IDs.
func (etc *EmailTemplateCreate) AddRoleEdgeIDs(ids ...uuid.UUID) *EmailTemplateCreate {
	etc.mutation.AddRoleEdgeIDs(ids...)
	return etc
}

// AddRoleEdges adds the "role_edges" edges to the Role entity.
func (etc *EmailTemplateCreate) AddRoleEdges(r ...*Role) *EmailTemplateCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return etc.AddRoleEdgeIDs(ids...)
}

// SetEventEdgeID sets the "event_edge" edge to the EmailEvent entity by ID.
func (etc *EmailTemplateCreate) SetEventEdgeID(id uuid.UUID) *EmailTemplateCreate {
	etc.mutation.SetEventEdgeID(id)
	return etc
}

// SetEventEdge sets the "event_edge" edge to the EmailEvent entity.
func (etc *EmailTemplateCreate) SetEventEdge(e *EmailEvent) *EmailTemplateCreate {
	return etc.SetEventEdgeID(e.ID)
}

// AddRoleEmailTemplateIDs adds the "role_email_templates" edge to the EmailRoleAttribute entity by IDs.
func (etc *EmailTemplateCreate) AddRoleEmailTemplateIDs(ids ...uuid.UUID) *EmailTemplateCreate {
	etc.mutation.AddRoleEmailTemplateIDs(ids...)
	return etc
}

// AddRoleEmailTemplates adds the "role_email_templates" edges to the EmailRoleAttribute entity.
func (etc *EmailTemplateCreate) AddRoleEmailTemplates(e ...*EmailRoleAttribute) *EmailTemplateCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddRoleEmailTemplateIDs(ids...)
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etc *EmailTemplateCreate) Mutation() *EmailTemplateMutation {
	return etc.mutation
}

// Save creates the EmailTemplate in the database.
func (etc *EmailTemplateCreate) Save(ctx context.Context) (*EmailTemplate, error) {
	var (
		err  error
		node *EmailTemplate
	)
	etc.defaults()
	if len(etc.hooks) == 0 {
		if err = etc.check(); err != nil {
			return nil, err
		}
		node, err = etc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etc.check(); err != nil {
				return nil, err
			}
			etc.mutation = mutation
			if node, err = etc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(etc.hooks) - 1; i >= 0; i-- {
			if etc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, etc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EmailTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EmailTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EmailTemplateCreate) SaveX(ctx context.Context) *EmailTemplate {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *EmailTemplateCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *EmailTemplateCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *EmailTemplateCreate) defaults() {
	if _, ok := etc.mutation.CreatedAt(); !ok {
		v := emailtemplate.DefaultCreatedAt()
		etc.mutation.SetCreatedAt(v)
	}
	if _, ok := etc.mutation.Status(); !ok {
		v := emailtemplate.DefaultStatus
		etc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EmailTemplateCreate) check() error {
	if _, ok := etc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailTemplate.created_at"`)}
	}
	if v, ok := etc.mutation.Event(); ok {
		if err := emailtemplate.EventValidator(v); err != nil {
			return &ValidationError{Name: "event", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.event": %w`, err)}
		}
	}
	if _, ok := etc.mutation.SendTo(); !ok {
		return &ValidationError{Name: "send_to", err: errors.New(`ent: missing required field "EmailTemplate.send_to"`)}
	}
	if _, ok := etc.mutation.Cc(); !ok {
		return &ValidationError{Name: "cc", err: errors.New(`ent: missing required field "EmailTemplate.cc"`)}
	}
	if _, ok := etc.mutation.Bcc(); !ok {
		return &ValidationError{Name: "bcc", err: errors.New(`ent: missing required field "EmailTemplate.bcc"`)}
	}
	if _, ok := etc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "EmailTemplate.subject"`)}
	}
	if v, ok := etc.mutation.Subject(); ok {
		if err := emailtemplate.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.subject": %w`, err)}
		}
	}
	if _, ok := etc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "EmailTemplate.content"`)}
	}
	if v, ok := etc.mutation.Content(); ok {
		if err := emailtemplate.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.content": %w`, err)}
		}
	}
	if _, ok := etc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "EmailTemplate.status"`)}
	}
	if v, ok := etc.mutation.Status(); ok {
		if err := emailtemplate.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "EmailTemplate.status": %w`, err)}
		}
	}
	if _, ok := etc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "EmailTemplate.event_id"`)}
	}
	if _, ok := etc.mutation.EventEdgeID(); !ok {
		return &ValidationError{Name: "event_edge", err: errors.New(`ent: missing required edge "EmailTemplate.event_edge"`)}
	}
	return nil
}

func (etc *EmailTemplateCreate) sqlSave(ctx context.Context) (*EmailTemplate, error) {
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
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

func (etc *EmailTemplateCreate) createSpec() (*EmailTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailTemplate{config: etc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: emailtemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailtemplate.FieldID,
			},
		}
	)
	if id, ok := etc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := etc.mutation.CreatedAt(); ok {
		_spec.SetField(emailtemplate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := etc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailtemplate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := etc.mutation.DeletedAt(); ok {
		_spec.SetField(emailtemplate.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := etc.mutation.Event(); ok {
		_spec.SetField(emailtemplate.FieldEvent, field.TypeEnum, value)
		_node.Event = value
	}
	if value, ok := etc.mutation.SendTo(); ok {
		_spec.SetField(emailtemplate.FieldSendTo, field.TypeJSON, value)
		_node.SendTo = value
	}
	if value, ok := etc.mutation.Cc(); ok {
		_spec.SetField(emailtemplate.FieldCc, field.TypeJSON, value)
		_node.Cc = value
	}
	if value, ok := etc.mutation.Bcc(); ok {
		_spec.SetField(emailtemplate.FieldBcc, field.TypeJSON, value)
		_node.Bcc = value
	}
	if value, ok := etc.mutation.Subject(); ok {
		_spec.SetField(emailtemplate.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := etc.mutation.Content(); ok {
		_spec.SetField(emailtemplate.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := etc.mutation.Signature(); ok {
		_spec.SetField(emailtemplate.FieldSignature, field.TypeString, value)
		_node.Signature = value
	}
	if value, ok := etc.mutation.Status(); ok {
		_spec.SetField(emailtemplate.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := etc.mutation.RoleEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   emailtemplate.RoleEdgesTable,
			Columns: emailtemplate.RoleEdgesPrimaryKey,
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
		createE := &EmailRoleAttributeCreate{config: etc.config, mutation: newEmailRoleAttributeMutation(etc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.EventEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailtemplate.EventEdgeTable,
			Columns: []string{emailtemplate.EventEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailevent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EventID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.RoleEmailTemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   emailtemplate.RoleEmailTemplatesTable,
			Columns: []string{emailtemplate.RoleEmailTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailroleattribute.FieldID,
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

// EmailTemplateCreateBulk is the builder for creating many EmailTemplate entities in bulk.
type EmailTemplateCreateBulk struct {
	config
	builders []*EmailTemplateCreate
}

// Save creates the EmailTemplate entities in the database.
func (etcb *EmailTemplateCreateBulk) Save(ctx context.Context) ([]*EmailTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EmailTemplate, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EmailTemplateCreateBulk) SaveX(ctx context.Context) []*EmailTemplate {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *EmailTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *EmailTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}

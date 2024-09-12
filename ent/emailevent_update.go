// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/emailevent"
	"trec/ent/emailtemplate"
	"trec/ent/outgoingemail"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailEventUpdate is the builder for updating EmailEvent entities.
type EmailEventUpdate struct {
	config
	hooks    []Hook
	mutation *EmailEventMutation
}

// Where appends a list predicates to the EmailEventUpdate builder.
func (eeu *EmailEventUpdate) Where(ps ...predicate.EmailEvent) *EmailEventUpdate {
	eeu.mutation.Where(ps...)
	return eeu
}

// SetModule sets the "module" field.
func (eeu *EmailEventUpdate) SetModule(e emailevent.Module) *EmailEventUpdate {
	eeu.mutation.SetModule(e)
	return eeu
}

// SetAction sets the "action" field.
func (eeu *EmailEventUpdate) SetAction(e emailevent.Action) *EmailEventUpdate {
	eeu.mutation.SetAction(e)
	return eeu
}

// SetName sets the "name" field.
func (eeu *EmailEventUpdate) SetName(s string) *EmailEventUpdate {
	eeu.mutation.SetName(s)
	return eeu
}

// SetUpdatedAt sets the "updated_at" field.
func (eeu *EmailEventUpdate) SetUpdatedAt(t time.Time) *EmailEventUpdate {
	eeu.mutation.SetUpdatedAt(t)
	return eeu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eeu *EmailEventUpdate) SetNillableUpdatedAt(t *time.Time) *EmailEventUpdate {
	if t != nil {
		eeu.SetUpdatedAt(*t)
	}
	return eeu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (eeu *EmailEventUpdate) ClearUpdatedAt() *EmailEventUpdate {
	eeu.mutation.ClearUpdatedAt()
	return eeu
}

// AddTemplateEdgeIDs adds the "template_edges" edge to the EmailTemplate entity by IDs.
func (eeu *EmailEventUpdate) AddTemplateEdgeIDs(ids ...uuid.UUID) *EmailEventUpdate {
	eeu.mutation.AddTemplateEdgeIDs(ids...)
	return eeu
}

// AddTemplateEdges adds the "template_edges" edges to the EmailTemplate entity.
func (eeu *EmailEventUpdate) AddTemplateEdges(e ...*EmailTemplate) *EmailEventUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eeu.AddTemplateEdgeIDs(ids...)
}

// AddOutgoingEmailEdgeIDs adds the "outgoing_email_edges" edge to the OutgoingEmail entity by IDs.
func (eeu *EmailEventUpdate) AddOutgoingEmailEdgeIDs(ids ...uuid.UUID) *EmailEventUpdate {
	eeu.mutation.AddOutgoingEmailEdgeIDs(ids...)
	return eeu
}

// AddOutgoingEmailEdges adds the "outgoing_email_edges" edges to the OutgoingEmail entity.
func (eeu *EmailEventUpdate) AddOutgoingEmailEdges(o ...*OutgoingEmail) *EmailEventUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eeu.AddOutgoingEmailEdgeIDs(ids...)
}

// Mutation returns the EmailEventMutation object of the builder.
func (eeu *EmailEventUpdate) Mutation() *EmailEventMutation {
	return eeu.mutation
}

// ClearTemplateEdges clears all "template_edges" edges to the EmailTemplate entity.
func (eeu *EmailEventUpdate) ClearTemplateEdges() *EmailEventUpdate {
	eeu.mutation.ClearTemplateEdges()
	return eeu
}

// RemoveTemplateEdgeIDs removes the "template_edges" edge to EmailTemplate entities by IDs.
func (eeu *EmailEventUpdate) RemoveTemplateEdgeIDs(ids ...uuid.UUID) *EmailEventUpdate {
	eeu.mutation.RemoveTemplateEdgeIDs(ids...)
	return eeu
}

// RemoveTemplateEdges removes "template_edges" edges to EmailTemplate entities.
func (eeu *EmailEventUpdate) RemoveTemplateEdges(e ...*EmailTemplate) *EmailEventUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eeu.RemoveTemplateEdgeIDs(ids...)
}

// ClearOutgoingEmailEdges clears all "outgoing_email_edges" edges to the OutgoingEmail entity.
func (eeu *EmailEventUpdate) ClearOutgoingEmailEdges() *EmailEventUpdate {
	eeu.mutation.ClearOutgoingEmailEdges()
	return eeu
}

// RemoveOutgoingEmailEdgeIDs removes the "outgoing_email_edges" edge to OutgoingEmail entities by IDs.
func (eeu *EmailEventUpdate) RemoveOutgoingEmailEdgeIDs(ids ...uuid.UUID) *EmailEventUpdate {
	eeu.mutation.RemoveOutgoingEmailEdgeIDs(ids...)
	return eeu
}

// RemoveOutgoingEmailEdges removes "outgoing_email_edges" edges to OutgoingEmail entities.
func (eeu *EmailEventUpdate) RemoveOutgoingEmailEdges(o ...*OutgoingEmail) *EmailEventUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eeu.RemoveOutgoingEmailEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eeu *EmailEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eeu.hooks) == 0 {
		if err = eeu.check(); err != nil {
			return 0, err
		}
		affected, err = eeu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eeu.check(); err != nil {
				return 0, err
			}
			eeu.mutation = mutation
			affected, err = eeu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eeu.hooks) - 1; i >= 0; i-- {
			if eeu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eeu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eeu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eeu *EmailEventUpdate) SaveX(ctx context.Context) int {
	affected, err := eeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eeu *EmailEventUpdate) Exec(ctx context.Context) error {
	_, err := eeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eeu *EmailEventUpdate) ExecX(ctx context.Context) {
	if err := eeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eeu *EmailEventUpdate) check() error {
	if v, ok := eeu.mutation.Module(); ok {
		if err := emailevent.ModuleValidator(v); err != nil {
			return &ValidationError{Name: "module", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.module": %w`, err)}
		}
	}
	if v, ok := eeu.mutation.Action(); ok {
		if err := emailevent.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.action": %w`, err)}
		}
	}
	return nil
}

func (eeu *EmailEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailevent.Table,
			Columns: emailevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailevent.FieldID,
			},
		},
	}
	if ps := eeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eeu.mutation.Module(); ok {
		_spec.SetField(emailevent.FieldModule, field.TypeEnum, value)
	}
	if value, ok := eeu.mutation.Action(); ok {
		_spec.SetField(emailevent.FieldAction, field.TypeEnum, value)
	}
	if value, ok := eeu.mutation.Name(); ok {
		_spec.SetField(emailevent.FieldName, field.TypeString, value)
	}
	if value, ok := eeu.mutation.UpdatedAt(); ok {
		_spec.SetField(emailevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if eeu.mutation.UpdatedAtCleared() {
		_spec.ClearField(emailevent.FieldUpdatedAt, field.TypeTime)
	}
	if eeu.mutation.TemplateEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.RemovedTemplateEdgesIDs(); len(nodes) > 0 && !eeu.mutation.TemplateEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.TemplateEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eeu.mutation.OutgoingEmailEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.RemovedOutgoingEmailEdgesIDs(); len(nodes) > 0 && !eeu.mutation.OutgoingEmailEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.OutgoingEmailEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EmailEventUpdateOne is the builder for updating a single EmailEvent entity.
type EmailEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailEventMutation
}

// SetModule sets the "module" field.
func (eeuo *EmailEventUpdateOne) SetModule(e emailevent.Module) *EmailEventUpdateOne {
	eeuo.mutation.SetModule(e)
	return eeuo
}

// SetAction sets the "action" field.
func (eeuo *EmailEventUpdateOne) SetAction(e emailevent.Action) *EmailEventUpdateOne {
	eeuo.mutation.SetAction(e)
	return eeuo
}

// SetName sets the "name" field.
func (eeuo *EmailEventUpdateOne) SetName(s string) *EmailEventUpdateOne {
	eeuo.mutation.SetName(s)
	return eeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (eeuo *EmailEventUpdateOne) SetUpdatedAt(t time.Time) *EmailEventUpdateOne {
	eeuo.mutation.SetUpdatedAt(t)
	return eeuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eeuo *EmailEventUpdateOne) SetNillableUpdatedAt(t *time.Time) *EmailEventUpdateOne {
	if t != nil {
		eeuo.SetUpdatedAt(*t)
	}
	return eeuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (eeuo *EmailEventUpdateOne) ClearUpdatedAt() *EmailEventUpdateOne {
	eeuo.mutation.ClearUpdatedAt()
	return eeuo
}

// AddTemplateEdgeIDs adds the "template_edges" edge to the EmailTemplate entity by IDs.
func (eeuo *EmailEventUpdateOne) AddTemplateEdgeIDs(ids ...uuid.UUID) *EmailEventUpdateOne {
	eeuo.mutation.AddTemplateEdgeIDs(ids...)
	return eeuo
}

// AddTemplateEdges adds the "template_edges" edges to the EmailTemplate entity.
func (eeuo *EmailEventUpdateOne) AddTemplateEdges(e ...*EmailTemplate) *EmailEventUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eeuo.AddTemplateEdgeIDs(ids...)
}

// AddOutgoingEmailEdgeIDs adds the "outgoing_email_edges" edge to the OutgoingEmail entity by IDs.
func (eeuo *EmailEventUpdateOne) AddOutgoingEmailEdgeIDs(ids ...uuid.UUID) *EmailEventUpdateOne {
	eeuo.mutation.AddOutgoingEmailEdgeIDs(ids...)
	return eeuo
}

// AddOutgoingEmailEdges adds the "outgoing_email_edges" edges to the OutgoingEmail entity.
func (eeuo *EmailEventUpdateOne) AddOutgoingEmailEdges(o ...*OutgoingEmail) *EmailEventUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eeuo.AddOutgoingEmailEdgeIDs(ids...)
}

// Mutation returns the EmailEventMutation object of the builder.
func (eeuo *EmailEventUpdateOne) Mutation() *EmailEventMutation {
	return eeuo.mutation
}

// ClearTemplateEdges clears all "template_edges" edges to the EmailTemplate entity.
func (eeuo *EmailEventUpdateOne) ClearTemplateEdges() *EmailEventUpdateOne {
	eeuo.mutation.ClearTemplateEdges()
	return eeuo
}

// RemoveTemplateEdgeIDs removes the "template_edges" edge to EmailTemplate entities by IDs.
func (eeuo *EmailEventUpdateOne) RemoveTemplateEdgeIDs(ids ...uuid.UUID) *EmailEventUpdateOne {
	eeuo.mutation.RemoveTemplateEdgeIDs(ids...)
	return eeuo
}

// RemoveTemplateEdges removes "template_edges" edges to EmailTemplate entities.
func (eeuo *EmailEventUpdateOne) RemoveTemplateEdges(e ...*EmailTemplate) *EmailEventUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eeuo.RemoveTemplateEdgeIDs(ids...)
}

// ClearOutgoingEmailEdges clears all "outgoing_email_edges" edges to the OutgoingEmail entity.
func (eeuo *EmailEventUpdateOne) ClearOutgoingEmailEdges() *EmailEventUpdateOne {
	eeuo.mutation.ClearOutgoingEmailEdges()
	return eeuo
}

// RemoveOutgoingEmailEdgeIDs removes the "outgoing_email_edges" edge to OutgoingEmail entities by IDs.
func (eeuo *EmailEventUpdateOne) RemoveOutgoingEmailEdgeIDs(ids ...uuid.UUID) *EmailEventUpdateOne {
	eeuo.mutation.RemoveOutgoingEmailEdgeIDs(ids...)
	return eeuo
}

// RemoveOutgoingEmailEdges removes "outgoing_email_edges" edges to OutgoingEmail entities.
func (eeuo *EmailEventUpdateOne) RemoveOutgoingEmailEdges(o ...*OutgoingEmail) *EmailEventUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eeuo.RemoveOutgoingEmailEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eeuo *EmailEventUpdateOne) Select(field string, fields ...string) *EmailEventUpdateOne {
	eeuo.fields = append([]string{field}, fields...)
	return eeuo
}

// Save executes the query and returns the updated EmailEvent entity.
func (eeuo *EmailEventUpdateOne) Save(ctx context.Context) (*EmailEvent, error) {
	var (
		err  error
		node *EmailEvent
	)
	if len(eeuo.hooks) == 0 {
		if err = eeuo.check(); err != nil {
			return nil, err
		}
		node, err = eeuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eeuo.check(); err != nil {
				return nil, err
			}
			eeuo.mutation = mutation
			node, err = eeuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eeuo.hooks) - 1; i >= 0; i-- {
			if eeuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eeuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eeuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EmailEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EmailEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eeuo *EmailEventUpdateOne) SaveX(ctx context.Context) *EmailEvent {
	node, err := eeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eeuo *EmailEventUpdateOne) Exec(ctx context.Context) error {
	_, err := eeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eeuo *EmailEventUpdateOne) ExecX(ctx context.Context) {
	if err := eeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eeuo *EmailEventUpdateOne) check() error {
	if v, ok := eeuo.mutation.Module(); ok {
		if err := emailevent.ModuleValidator(v); err != nil {
			return &ValidationError{Name: "module", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.module": %w`, err)}
		}
	}
	if v, ok := eeuo.mutation.Action(); ok {
		if err := emailevent.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.action": %w`, err)}
		}
	}
	return nil
}

func (eeuo *EmailEventUpdateOne) sqlSave(ctx context.Context) (_node *EmailEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailevent.Table,
			Columns: emailevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailevent.FieldID,
			},
		},
	}
	id, ok := eeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailevent.FieldID)
		for _, f := range fields {
			if !emailevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emailevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eeuo.mutation.Module(); ok {
		_spec.SetField(emailevent.FieldModule, field.TypeEnum, value)
	}
	if value, ok := eeuo.mutation.Action(); ok {
		_spec.SetField(emailevent.FieldAction, field.TypeEnum, value)
	}
	if value, ok := eeuo.mutation.Name(); ok {
		_spec.SetField(emailevent.FieldName, field.TypeString, value)
	}
	if value, ok := eeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(emailevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if eeuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(emailevent.FieldUpdatedAt, field.TypeTime)
	}
	if eeuo.mutation.TemplateEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.RemovedTemplateEdgesIDs(); len(nodes) > 0 && !eeuo.mutation.TemplateEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.TemplateEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.TemplateEdgesTable,
			Columns: []string{emailevent.TemplateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailtemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eeuo.mutation.OutgoingEmailEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.RemovedOutgoingEmailEdgesIDs(); len(nodes) > 0 && !eeuo.mutation.OutgoingEmailEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.OutgoingEmailEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   emailevent.OutgoingEmailEdgesTable,
			Columns: []string{emailevent.OutgoingEmailEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: outgoingemail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EmailEvent{config: eeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

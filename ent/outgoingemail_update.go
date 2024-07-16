// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/outgoingemail"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OutgoingEmailUpdate is the builder for updating OutgoingEmail entities.
type OutgoingEmailUpdate struct {
	config
	hooks    []Hook
	mutation *OutgoingEmailMutation
}

// Where appends a list predicates to the OutgoingEmailUpdate builder.
func (oeu *OutgoingEmailUpdate) Where(ps ...predicate.OutgoingEmail) *OutgoingEmailUpdate {
	oeu.mutation.Where(ps...)
	return oeu
}

// SetUpdatedAt sets the "updated_at" field.
func (oeu *OutgoingEmailUpdate) SetUpdatedAt(t time.Time) *OutgoingEmailUpdate {
	oeu.mutation.SetUpdatedAt(t)
	return oeu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oeu *OutgoingEmailUpdate) SetNillableUpdatedAt(t *time.Time) *OutgoingEmailUpdate {
	if t != nil {
		oeu.SetUpdatedAt(*t)
	}
	return oeu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oeu *OutgoingEmailUpdate) ClearUpdatedAt() *OutgoingEmailUpdate {
	oeu.mutation.ClearUpdatedAt()
	return oeu
}

// SetDeletedAt sets the "deleted_at" field.
func (oeu *OutgoingEmailUpdate) SetDeletedAt(t time.Time) *OutgoingEmailUpdate {
	oeu.mutation.SetDeletedAt(t)
	return oeu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oeu *OutgoingEmailUpdate) SetNillableDeletedAt(t *time.Time) *OutgoingEmailUpdate {
	if t != nil {
		oeu.SetDeletedAt(*t)
	}
	return oeu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (oeu *OutgoingEmailUpdate) ClearDeletedAt() *OutgoingEmailUpdate {
	oeu.mutation.ClearDeletedAt()
	return oeu
}

// SetTo sets the "to" field.
func (oeu *OutgoingEmailUpdate) SetTo(s []string) *OutgoingEmailUpdate {
	oeu.mutation.SetTo(s)
	return oeu
}

// AppendTo appends s to the "to" field.
func (oeu *OutgoingEmailUpdate) AppendTo(s []string) *OutgoingEmailUpdate {
	oeu.mutation.AppendTo(s)
	return oeu
}

// SetCc sets the "cc" field.
func (oeu *OutgoingEmailUpdate) SetCc(s []string) *OutgoingEmailUpdate {
	oeu.mutation.SetCc(s)
	return oeu
}

// AppendCc appends s to the "cc" field.
func (oeu *OutgoingEmailUpdate) AppendCc(s []string) *OutgoingEmailUpdate {
	oeu.mutation.AppendCc(s)
	return oeu
}

// SetBcc sets the "bcc" field.
func (oeu *OutgoingEmailUpdate) SetBcc(s []string) *OutgoingEmailUpdate {
	oeu.mutation.SetBcc(s)
	return oeu
}

// AppendBcc appends s to the "bcc" field.
func (oeu *OutgoingEmailUpdate) AppendBcc(s []string) *OutgoingEmailUpdate {
	oeu.mutation.AppendBcc(s)
	return oeu
}

// SetSubject sets the "subject" field.
func (oeu *OutgoingEmailUpdate) SetSubject(s string) *OutgoingEmailUpdate {
	oeu.mutation.SetSubject(s)
	return oeu
}

// SetContent sets the "content" field.
func (oeu *OutgoingEmailUpdate) SetContent(s string) *OutgoingEmailUpdate {
	oeu.mutation.SetContent(s)
	return oeu
}

// SetSignature sets the "signature" field.
func (oeu *OutgoingEmailUpdate) SetSignature(s string) *OutgoingEmailUpdate {
	oeu.mutation.SetSignature(s)
	return oeu
}

// SetEmailTemplateID sets the "email_template_id" field.
func (oeu *OutgoingEmailUpdate) SetEmailTemplateID(u uuid.UUID) *OutgoingEmailUpdate {
	oeu.mutation.SetEmailTemplateID(u)
	return oeu
}

// SetNillableEmailTemplateID sets the "email_template_id" field if the given value is not nil.
func (oeu *OutgoingEmailUpdate) SetNillableEmailTemplateID(u *uuid.UUID) *OutgoingEmailUpdate {
	if u != nil {
		oeu.SetEmailTemplateID(*u)
	}
	return oeu
}

// ClearEmailTemplateID clears the value of the "email_template_id" field.
func (oeu *OutgoingEmailUpdate) ClearEmailTemplateID() *OutgoingEmailUpdate {
	oeu.mutation.ClearEmailTemplateID()
	return oeu
}

// SetStatus sets the "status" field.
func (oeu *OutgoingEmailUpdate) SetStatus(o outgoingemail.Status) *OutgoingEmailUpdate {
	oeu.mutation.SetStatus(o)
	return oeu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oeu *OutgoingEmailUpdate) SetNillableStatus(o *outgoingemail.Status) *OutgoingEmailUpdate {
	if o != nil {
		oeu.SetStatus(*o)
	}
	return oeu
}

// Mutation returns the OutgoingEmailMutation object of the builder.
func (oeu *OutgoingEmailUpdate) Mutation() *OutgoingEmailMutation {
	return oeu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oeu *OutgoingEmailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oeu.hooks) == 0 {
		if err = oeu.check(); err != nil {
			return 0, err
		}
		affected, err = oeu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutgoingEmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oeu.check(); err != nil {
				return 0, err
			}
			oeu.mutation = mutation
			affected, err = oeu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oeu.hooks) - 1; i >= 0; i-- {
			if oeu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oeu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oeu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oeu *OutgoingEmailUpdate) SaveX(ctx context.Context) int {
	affected, err := oeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oeu *OutgoingEmailUpdate) Exec(ctx context.Context) error {
	_, err := oeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeu *OutgoingEmailUpdate) ExecX(ctx context.Context) {
	if err := oeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oeu *OutgoingEmailUpdate) check() error {
	if v, ok := oeu.mutation.Subject(); ok {
		if err := outgoingemail.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.subject": %w`, err)}
		}
	}
	if v, ok := oeu.mutation.Content(); ok {
		if err := outgoingemail.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.content": %w`, err)}
		}
	}
	if v, ok := oeu.mutation.Status(); ok {
		if err := outgoingemail.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.status": %w`, err)}
		}
	}
	return nil
}

func (oeu *OutgoingEmailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outgoingemail.Table,
			Columns: outgoingemail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outgoingemail.FieldID,
			},
		},
	}
	if ps := oeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeu.mutation.UpdatedAt(); ok {
		_spec.SetField(outgoingemail.FieldUpdatedAt, field.TypeTime, value)
	}
	if oeu.mutation.UpdatedAtCleared() {
		_spec.ClearField(outgoingemail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oeu.mutation.DeletedAt(); ok {
		_spec.SetField(outgoingemail.FieldDeletedAt, field.TypeTime, value)
	}
	if oeu.mutation.DeletedAtCleared() {
		_spec.ClearField(outgoingemail.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := oeu.mutation.To(); ok {
		_spec.SetField(outgoingemail.FieldTo, field.TypeJSON, value)
	}
	if value, ok := oeu.mutation.AppendedTo(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldTo, value)
		})
	}
	if value, ok := oeu.mutation.Cc(); ok {
		_spec.SetField(outgoingemail.FieldCc, field.TypeJSON, value)
	}
	if value, ok := oeu.mutation.AppendedCc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldCc, value)
		})
	}
	if value, ok := oeu.mutation.Bcc(); ok {
		_spec.SetField(outgoingemail.FieldBcc, field.TypeJSON, value)
	}
	if value, ok := oeu.mutation.AppendedBcc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldBcc, value)
		})
	}
	if value, ok := oeu.mutation.Subject(); ok {
		_spec.SetField(outgoingemail.FieldSubject, field.TypeString, value)
	}
	if value, ok := oeu.mutation.Content(); ok {
		_spec.SetField(outgoingemail.FieldContent, field.TypeString, value)
	}
	if value, ok := oeu.mutation.Signature(); ok {
		_spec.SetField(outgoingemail.FieldSignature, field.TypeString, value)
	}
	if value, ok := oeu.mutation.EmailTemplateID(); ok {
		_spec.SetField(outgoingemail.FieldEmailTemplateID, field.TypeUUID, value)
	}
	if oeu.mutation.EmailTemplateIDCleared() {
		_spec.ClearField(outgoingemail.FieldEmailTemplateID, field.TypeUUID)
	}
	if value, ok := oeu.mutation.Status(); ok {
		_spec.SetField(outgoingemail.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outgoingemail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OutgoingEmailUpdateOne is the builder for updating a single OutgoingEmail entity.
type OutgoingEmailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutgoingEmailMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oeuo *OutgoingEmailUpdateOne) SetUpdatedAt(t time.Time) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetUpdatedAt(t)
	return oeuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oeuo *OutgoingEmailUpdateOne) SetNillableUpdatedAt(t *time.Time) *OutgoingEmailUpdateOne {
	if t != nil {
		oeuo.SetUpdatedAt(*t)
	}
	return oeuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oeuo *OutgoingEmailUpdateOne) ClearUpdatedAt() *OutgoingEmailUpdateOne {
	oeuo.mutation.ClearUpdatedAt()
	return oeuo
}

// SetDeletedAt sets the "deleted_at" field.
func (oeuo *OutgoingEmailUpdateOne) SetDeletedAt(t time.Time) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetDeletedAt(t)
	return oeuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oeuo *OutgoingEmailUpdateOne) SetNillableDeletedAt(t *time.Time) *OutgoingEmailUpdateOne {
	if t != nil {
		oeuo.SetDeletedAt(*t)
	}
	return oeuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (oeuo *OutgoingEmailUpdateOne) ClearDeletedAt() *OutgoingEmailUpdateOne {
	oeuo.mutation.ClearDeletedAt()
	return oeuo
}

// SetTo sets the "to" field.
func (oeuo *OutgoingEmailUpdateOne) SetTo(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetTo(s)
	return oeuo
}

// AppendTo appends s to the "to" field.
func (oeuo *OutgoingEmailUpdateOne) AppendTo(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.AppendTo(s)
	return oeuo
}

// SetCc sets the "cc" field.
func (oeuo *OutgoingEmailUpdateOne) SetCc(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetCc(s)
	return oeuo
}

// AppendCc appends s to the "cc" field.
func (oeuo *OutgoingEmailUpdateOne) AppendCc(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.AppendCc(s)
	return oeuo
}

// SetBcc sets the "bcc" field.
func (oeuo *OutgoingEmailUpdateOne) SetBcc(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetBcc(s)
	return oeuo
}

// AppendBcc appends s to the "bcc" field.
func (oeuo *OutgoingEmailUpdateOne) AppendBcc(s []string) *OutgoingEmailUpdateOne {
	oeuo.mutation.AppendBcc(s)
	return oeuo
}

// SetSubject sets the "subject" field.
func (oeuo *OutgoingEmailUpdateOne) SetSubject(s string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetSubject(s)
	return oeuo
}

// SetContent sets the "content" field.
func (oeuo *OutgoingEmailUpdateOne) SetContent(s string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetContent(s)
	return oeuo
}

// SetSignature sets the "signature" field.
func (oeuo *OutgoingEmailUpdateOne) SetSignature(s string) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetSignature(s)
	return oeuo
}

// SetEmailTemplateID sets the "email_template_id" field.
func (oeuo *OutgoingEmailUpdateOne) SetEmailTemplateID(u uuid.UUID) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetEmailTemplateID(u)
	return oeuo
}

// SetNillableEmailTemplateID sets the "email_template_id" field if the given value is not nil.
func (oeuo *OutgoingEmailUpdateOne) SetNillableEmailTemplateID(u *uuid.UUID) *OutgoingEmailUpdateOne {
	if u != nil {
		oeuo.SetEmailTemplateID(*u)
	}
	return oeuo
}

// ClearEmailTemplateID clears the value of the "email_template_id" field.
func (oeuo *OutgoingEmailUpdateOne) ClearEmailTemplateID() *OutgoingEmailUpdateOne {
	oeuo.mutation.ClearEmailTemplateID()
	return oeuo
}

// SetStatus sets the "status" field.
func (oeuo *OutgoingEmailUpdateOne) SetStatus(o outgoingemail.Status) *OutgoingEmailUpdateOne {
	oeuo.mutation.SetStatus(o)
	return oeuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oeuo *OutgoingEmailUpdateOne) SetNillableStatus(o *outgoingemail.Status) *OutgoingEmailUpdateOne {
	if o != nil {
		oeuo.SetStatus(*o)
	}
	return oeuo
}

// Mutation returns the OutgoingEmailMutation object of the builder.
func (oeuo *OutgoingEmailUpdateOne) Mutation() *OutgoingEmailMutation {
	return oeuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oeuo *OutgoingEmailUpdateOne) Select(field string, fields ...string) *OutgoingEmailUpdateOne {
	oeuo.fields = append([]string{field}, fields...)
	return oeuo
}

// Save executes the query and returns the updated OutgoingEmail entity.
func (oeuo *OutgoingEmailUpdateOne) Save(ctx context.Context) (*OutgoingEmail, error) {
	var (
		err  error
		node *OutgoingEmail
	)
	if len(oeuo.hooks) == 0 {
		if err = oeuo.check(); err != nil {
			return nil, err
		}
		node, err = oeuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutgoingEmailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oeuo.check(); err != nil {
				return nil, err
			}
			oeuo.mutation = mutation
			node, err = oeuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oeuo.hooks) - 1; i >= 0; i-- {
			if oeuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oeuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oeuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (oeuo *OutgoingEmailUpdateOne) SaveX(ctx context.Context) *OutgoingEmail {
	node, err := oeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oeuo *OutgoingEmailUpdateOne) Exec(ctx context.Context) error {
	_, err := oeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oeuo *OutgoingEmailUpdateOne) ExecX(ctx context.Context) {
	if err := oeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oeuo *OutgoingEmailUpdateOne) check() error {
	if v, ok := oeuo.mutation.Subject(); ok {
		if err := outgoingemail.SubjectValidator(v); err != nil {
			return &ValidationError{Name: "subject", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.subject": %w`, err)}
		}
	}
	if v, ok := oeuo.mutation.Content(); ok {
		if err := outgoingemail.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.content": %w`, err)}
		}
	}
	if v, ok := oeuo.mutation.Status(); ok {
		if err := outgoingemail.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OutgoingEmail.status": %w`, err)}
		}
	}
	return nil
}

func (oeuo *OutgoingEmailUpdateOne) sqlSave(ctx context.Context) (_node *OutgoingEmail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outgoingemail.Table,
			Columns: outgoingemail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outgoingemail.FieldID,
			},
		},
	}
	id, ok := oeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OutgoingEmail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outgoingemail.FieldID)
		for _, f := range fields {
			if !outgoingemail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != outgoingemail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(outgoingemail.FieldUpdatedAt, field.TypeTime, value)
	}
	if oeuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(outgoingemail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oeuo.mutation.DeletedAt(); ok {
		_spec.SetField(outgoingemail.FieldDeletedAt, field.TypeTime, value)
	}
	if oeuo.mutation.DeletedAtCleared() {
		_spec.ClearField(outgoingemail.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := oeuo.mutation.To(); ok {
		_spec.SetField(outgoingemail.FieldTo, field.TypeJSON, value)
	}
	if value, ok := oeuo.mutation.AppendedTo(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldTo, value)
		})
	}
	if value, ok := oeuo.mutation.Cc(); ok {
		_spec.SetField(outgoingemail.FieldCc, field.TypeJSON, value)
	}
	if value, ok := oeuo.mutation.AppendedCc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldCc, value)
		})
	}
	if value, ok := oeuo.mutation.Bcc(); ok {
		_spec.SetField(outgoingemail.FieldBcc, field.TypeJSON, value)
	}
	if value, ok := oeuo.mutation.AppendedBcc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, outgoingemail.FieldBcc, value)
		})
	}
	if value, ok := oeuo.mutation.Subject(); ok {
		_spec.SetField(outgoingemail.FieldSubject, field.TypeString, value)
	}
	if value, ok := oeuo.mutation.Content(); ok {
		_spec.SetField(outgoingemail.FieldContent, field.TypeString, value)
	}
	if value, ok := oeuo.mutation.Signature(); ok {
		_spec.SetField(outgoingemail.FieldSignature, field.TypeString, value)
	}
	if value, ok := oeuo.mutation.EmailTemplateID(); ok {
		_spec.SetField(outgoingemail.FieldEmailTemplateID, field.TypeUUID, value)
	}
	if oeuo.mutation.EmailTemplateIDCleared() {
		_spec.ClearField(outgoingemail.FieldEmailTemplateID, field.TypeUUID)
	}
	if value, ok := oeuo.mutation.Status(); ok {
		_spec.SetField(outgoingemail.FieldStatus, field.TypeEnum, value)
	}
	_node = &OutgoingEmail{config: oeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outgoingemail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
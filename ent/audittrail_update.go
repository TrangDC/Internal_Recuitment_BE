// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/audittrail"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuditTrailUpdate is the builder for updating AuditTrail entities.
type AuditTrailUpdate struct {
	config
	hooks    []Hook
	mutation *AuditTrailMutation
}

// Where appends a list predicates to the AuditTrailUpdate builder.
func (atu *AuditTrailUpdate) Where(ps ...predicate.AuditTrail) *AuditTrailUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetCreatedBy sets the "created_by" field.
func (atu *AuditTrailUpdate) SetCreatedBy(u uuid.UUID) *AuditTrailUpdate {
	atu.mutation.SetCreatedBy(u)
	return atu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableCreatedBy(u *uuid.UUID) *AuditTrailUpdate {
	if u != nil {
		atu.SetCreatedBy(*u)
	}
	return atu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (atu *AuditTrailUpdate) ClearCreatedBy() *AuditTrailUpdate {
	atu.mutation.ClearCreatedBy()
	return atu
}

// SetRecordId sets the "recordId" field.
func (atu *AuditTrailUpdate) SetRecordId(u uuid.UUID) *AuditTrailUpdate {
	atu.mutation.SetRecordId(u)
	return atu
}

// SetModule sets the "module" field.
func (atu *AuditTrailUpdate) SetModule(a audittrail.Module) *AuditTrailUpdate {
	atu.mutation.SetModule(a)
	return atu
}

// SetActionType sets the "actionType" field.
func (atu *AuditTrailUpdate) SetActionType(at audittrail.ActionType) *AuditTrailUpdate {
	atu.mutation.SetActionType(at)
	return atu
}

// SetNillableActionType sets the "actionType" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableActionType(at *audittrail.ActionType) *AuditTrailUpdate {
	if at != nil {
		atu.SetActionType(*at)
	}
	return atu
}

// ClearActionType clears the value of the "actionType" field.
func (atu *AuditTrailUpdate) ClearActionType() *AuditTrailUpdate {
	atu.mutation.ClearActionType()
	return atu
}

// SetNote sets the "note" field.
func (atu *AuditTrailUpdate) SetNote(s string) *AuditTrailUpdate {
	atu.mutation.SetNote(s)
	return atu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableNote(s *string) *AuditTrailUpdate {
	if s != nil {
		atu.SetNote(*s)
	}
	return atu
}

// ClearNote clears the value of the "note" field.
func (atu *AuditTrailUpdate) ClearNote() *AuditTrailUpdate {
	atu.mutation.ClearNote()
	return atu
}

// SetRecordChanges sets the "record_changes" field.
func (atu *AuditTrailUpdate) SetRecordChanges(s string) *AuditTrailUpdate {
	atu.mutation.SetRecordChanges(s)
	return atu
}

// SetNillableRecordChanges sets the "record_changes" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableRecordChanges(s *string) *AuditTrailUpdate {
	if s != nil {
		atu.SetRecordChanges(*s)
	}
	return atu
}

// ClearRecordChanges clears the value of the "record_changes" field.
func (atu *AuditTrailUpdate) ClearRecordChanges() *AuditTrailUpdate {
	atu.mutation.ClearRecordChanges()
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *AuditTrailUpdate) SetUpdatedAt(t time.Time) *AuditTrailUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableUpdatedAt(t *time.Time) *AuditTrailUpdate {
	if t != nil {
		atu.SetUpdatedAt(*t)
	}
	return atu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atu *AuditTrailUpdate) ClearUpdatedAt() *AuditTrailUpdate {
	atu.mutation.ClearUpdatedAt()
	return atu
}

// SetDeletedAt sets the "deleted_at" field.
func (atu *AuditTrailUpdate) SetDeletedAt(t time.Time) *AuditTrailUpdate {
	atu.mutation.SetDeletedAt(t)
	return atu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableDeletedAt(t *time.Time) *AuditTrailUpdate {
	if t != nil {
		atu.SetDeletedAt(*t)
	}
	return atu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atu *AuditTrailUpdate) ClearDeletedAt() *AuditTrailUpdate {
	atu.mutation.ClearDeletedAt()
	return atu
}

// SetCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID.
func (atu *AuditTrailUpdate) SetCreatedByEdgeID(id uuid.UUID) *AuditTrailUpdate {
	atu.mutation.SetCreatedByEdgeID(id)
	return atu
}

// SetNillableCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID if the given value is not nil.
func (atu *AuditTrailUpdate) SetNillableCreatedByEdgeID(id *uuid.UUID) *AuditTrailUpdate {
	if id != nil {
		atu = atu.SetCreatedByEdgeID(*id)
	}
	return atu
}

// SetCreatedByEdge sets the "created_by_edge" edge to the User entity.
func (atu *AuditTrailUpdate) SetCreatedByEdge(u *User) *AuditTrailUpdate {
	return atu.SetCreatedByEdgeID(u.ID)
}

// Mutation returns the AuditTrailMutation object of the builder.
func (atu *AuditTrailUpdate) Mutation() *AuditTrailMutation {
	return atu.mutation
}

// ClearCreatedByEdge clears the "created_by_edge" edge to the User entity.
func (atu *AuditTrailUpdate) ClearCreatedByEdge() *AuditTrailUpdate {
	atu.mutation.ClearCreatedByEdge()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AuditTrailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuditTrailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AuditTrailUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AuditTrailUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AuditTrailUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AuditTrailUpdate) check() error {
	if v, ok := atu.mutation.Module(); ok {
		if err := audittrail.ModuleValidator(v); err != nil {
			return &ValidationError{Name: "module", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.module": %w`, err)}
		}
	}
	if v, ok := atu.mutation.ActionType(); ok {
		if err := audittrail.ActionTypeValidator(v); err != nil {
			return &ValidationError{Name: "actionType", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.actionType": %w`, err)}
		}
	}
	if v, ok := atu.mutation.Note(); ok {
		if err := audittrail.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.note": %w`, err)}
		}
	}
	return nil
}

func (atu *AuditTrailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audittrail.Table,
			Columns: audittrail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: audittrail.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.RecordId(); ok {
		_spec.SetField(audittrail.FieldRecordId, field.TypeUUID, value)
	}
	if value, ok := atu.mutation.Module(); ok {
		_spec.SetField(audittrail.FieldModule, field.TypeEnum, value)
	}
	if value, ok := atu.mutation.ActionType(); ok {
		_spec.SetField(audittrail.FieldActionType, field.TypeEnum, value)
	}
	if atu.mutation.ActionTypeCleared() {
		_spec.ClearField(audittrail.FieldActionType, field.TypeEnum)
	}
	if value, ok := atu.mutation.Note(); ok {
		_spec.SetField(audittrail.FieldNote, field.TypeString, value)
	}
	if atu.mutation.NoteCleared() {
		_spec.ClearField(audittrail.FieldNote, field.TypeString)
	}
	if value, ok := atu.mutation.RecordChanges(); ok {
		_spec.SetField(audittrail.FieldRecordChanges, field.TypeString, value)
	}
	if atu.mutation.RecordChangesCleared() {
		_spec.ClearField(audittrail.FieldRecordChanges, field.TypeString)
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(audittrail.FieldUpdatedAt, field.TypeTime, value)
	}
	if atu.mutation.UpdatedAtCleared() {
		_spec.ClearField(audittrail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.DeletedAt(); ok {
		_spec.SetField(audittrail.FieldDeletedAt, field.TypeTime, value)
	}
	if atu.mutation.DeletedAtCleared() {
		_spec.ClearField(audittrail.FieldDeletedAt, field.TypeTime)
	}
	if atu.mutation.CreatedByEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audittrail.CreatedByEdgeTable,
			Columns: []string{audittrail.CreatedByEdgeColumn},
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
	if nodes := atu.mutation.CreatedByEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audittrail.CreatedByEdgeTable,
			Columns: []string{audittrail.CreatedByEdgeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audittrail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AuditTrailUpdateOne is the builder for updating a single AuditTrail entity.
type AuditTrailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuditTrailMutation
}

// SetCreatedBy sets the "created_by" field.
func (atuo *AuditTrailUpdateOne) SetCreatedBy(u uuid.UUID) *AuditTrailUpdateOne {
	atuo.mutation.SetCreatedBy(u)
	return atuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *AuditTrailUpdateOne {
	if u != nil {
		atuo.SetCreatedBy(*u)
	}
	return atuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (atuo *AuditTrailUpdateOne) ClearCreatedBy() *AuditTrailUpdateOne {
	atuo.mutation.ClearCreatedBy()
	return atuo
}

// SetRecordId sets the "recordId" field.
func (atuo *AuditTrailUpdateOne) SetRecordId(u uuid.UUID) *AuditTrailUpdateOne {
	atuo.mutation.SetRecordId(u)
	return atuo
}

// SetModule sets the "module" field.
func (atuo *AuditTrailUpdateOne) SetModule(a audittrail.Module) *AuditTrailUpdateOne {
	atuo.mutation.SetModule(a)
	return atuo
}

// SetActionType sets the "actionType" field.
func (atuo *AuditTrailUpdateOne) SetActionType(at audittrail.ActionType) *AuditTrailUpdateOne {
	atuo.mutation.SetActionType(at)
	return atuo
}

// SetNillableActionType sets the "actionType" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableActionType(at *audittrail.ActionType) *AuditTrailUpdateOne {
	if at != nil {
		atuo.SetActionType(*at)
	}
	return atuo
}

// ClearActionType clears the value of the "actionType" field.
func (atuo *AuditTrailUpdateOne) ClearActionType() *AuditTrailUpdateOne {
	atuo.mutation.ClearActionType()
	return atuo
}

// SetNote sets the "note" field.
func (atuo *AuditTrailUpdateOne) SetNote(s string) *AuditTrailUpdateOne {
	atuo.mutation.SetNote(s)
	return atuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableNote(s *string) *AuditTrailUpdateOne {
	if s != nil {
		atuo.SetNote(*s)
	}
	return atuo
}

// ClearNote clears the value of the "note" field.
func (atuo *AuditTrailUpdateOne) ClearNote() *AuditTrailUpdateOne {
	atuo.mutation.ClearNote()
	return atuo
}

// SetRecordChanges sets the "record_changes" field.
func (atuo *AuditTrailUpdateOne) SetRecordChanges(s string) *AuditTrailUpdateOne {
	atuo.mutation.SetRecordChanges(s)
	return atuo
}

// SetNillableRecordChanges sets the "record_changes" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableRecordChanges(s *string) *AuditTrailUpdateOne {
	if s != nil {
		atuo.SetRecordChanges(*s)
	}
	return atuo
}

// ClearRecordChanges clears the value of the "record_changes" field.
func (atuo *AuditTrailUpdateOne) ClearRecordChanges() *AuditTrailUpdateOne {
	atuo.mutation.ClearRecordChanges()
	return atuo
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *AuditTrailUpdateOne) SetUpdatedAt(t time.Time) *AuditTrailUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableUpdatedAt(t *time.Time) *AuditTrailUpdateOne {
	if t != nil {
		atuo.SetUpdatedAt(*t)
	}
	return atuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atuo *AuditTrailUpdateOne) ClearUpdatedAt() *AuditTrailUpdateOne {
	atuo.mutation.ClearUpdatedAt()
	return atuo
}

// SetDeletedAt sets the "deleted_at" field.
func (atuo *AuditTrailUpdateOne) SetDeletedAt(t time.Time) *AuditTrailUpdateOne {
	atuo.mutation.SetDeletedAt(t)
	return atuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableDeletedAt(t *time.Time) *AuditTrailUpdateOne {
	if t != nil {
		atuo.SetDeletedAt(*t)
	}
	return atuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atuo *AuditTrailUpdateOne) ClearDeletedAt() *AuditTrailUpdateOne {
	atuo.mutation.ClearDeletedAt()
	return atuo
}

// SetCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID.
func (atuo *AuditTrailUpdateOne) SetCreatedByEdgeID(id uuid.UUID) *AuditTrailUpdateOne {
	atuo.mutation.SetCreatedByEdgeID(id)
	return atuo
}

// SetNillableCreatedByEdgeID sets the "created_by_edge" edge to the User entity by ID if the given value is not nil.
func (atuo *AuditTrailUpdateOne) SetNillableCreatedByEdgeID(id *uuid.UUID) *AuditTrailUpdateOne {
	if id != nil {
		atuo = atuo.SetCreatedByEdgeID(*id)
	}
	return atuo
}

// SetCreatedByEdge sets the "created_by_edge" edge to the User entity.
func (atuo *AuditTrailUpdateOne) SetCreatedByEdge(u *User) *AuditTrailUpdateOne {
	return atuo.SetCreatedByEdgeID(u.ID)
}

// Mutation returns the AuditTrailMutation object of the builder.
func (atuo *AuditTrailUpdateOne) Mutation() *AuditTrailMutation {
	return atuo.mutation
}

// ClearCreatedByEdge clears the "created_by_edge" edge to the User entity.
func (atuo *AuditTrailUpdateOne) ClearCreatedByEdge() *AuditTrailUpdateOne {
	atuo.mutation.ClearCreatedByEdge()
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AuditTrailUpdateOne) Select(field string, fields ...string) *AuditTrailUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AuditTrail entity.
func (atuo *AuditTrailUpdateOne) Save(ctx context.Context) (*AuditTrail, error) {
	var (
		err  error
		node *AuditTrail
	)
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuditTrailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AuditTrail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuditTrailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AuditTrailUpdateOne) SaveX(ctx context.Context) *AuditTrail {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AuditTrailUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AuditTrailUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AuditTrailUpdateOne) check() error {
	if v, ok := atuo.mutation.Module(); ok {
		if err := audittrail.ModuleValidator(v); err != nil {
			return &ValidationError{Name: "module", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.module": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.ActionType(); ok {
		if err := audittrail.ActionTypeValidator(v); err != nil {
			return &ValidationError{Name: "actionType", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.actionType": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.Note(); ok {
		if err := audittrail.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "AuditTrail.note": %w`, err)}
		}
	}
	return nil
}

func (atuo *AuditTrailUpdateOne) sqlSave(ctx context.Context) (_node *AuditTrail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audittrail.Table,
			Columns: audittrail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: audittrail.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuditTrail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audittrail.FieldID)
		for _, f := range fields {
			if !audittrail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != audittrail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.RecordId(); ok {
		_spec.SetField(audittrail.FieldRecordId, field.TypeUUID, value)
	}
	if value, ok := atuo.mutation.Module(); ok {
		_spec.SetField(audittrail.FieldModule, field.TypeEnum, value)
	}
	if value, ok := atuo.mutation.ActionType(); ok {
		_spec.SetField(audittrail.FieldActionType, field.TypeEnum, value)
	}
	if atuo.mutation.ActionTypeCleared() {
		_spec.ClearField(audittrail.FieldActionType, field.TypeEnum)
	}
	if value, ok := atuo.mutation.Note(); ok {
		_spec.SetField(audittrail.FieldNote, field.TypeString, value)
	}
	if atuo.mutation.NoteCleared() {
		_spec.ClearField(audittrail.FieldNote, field.TypeString)
	}
	if value, ok := atuo.mutation.RecordChanges(); ok {
		_spec.SetField(audittrail.FieldRecordChanges, field.TypeString, value)
	}
	if atuo.mutation.RecordChangesCleared() {
		_spec.ClearField(audittrail.FieldRecordChanges, field.TypeString)
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(audittrail.FieldUpdatedAt, field.TypeTime, value)
	}
	if atuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(audittrail.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.DeletedAt(); ok {
		_spec.SetField(audittrail.FieldDeletedAt, field.TypeTime, value)
	}
	if atuo.mutation.DeletedAtCleared() {
		_spec.ClearField(audittrail.FieldDeletedAt, field.TypeTime)
	}
	if atuo.mutation.CreatedByEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audittrail.CreatedByEdgeTable,
			Columns: []string{audittrail.CreatedByEdgeColumn},
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
	if nodes := atuo.mutation.CreatedByEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audittrail.CreatedByEdgeTable,
			Columns: []string{audittrail.CreatedByEdgeColumn},
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
	_node = &AuditTrail{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audittrail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"trec/ent/jobtitle"
	"trec/ent/predicate"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeJobTitle = "JobTitle"
)

// JobTitleMutation represents an operation that mutates the JobTitle nodes in the graph.
type JobTitleMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	code          *string
	name          *string
	description   *string
	specification *string
	created_at    *time.Time
	updated_at    *time.Time
	deleted_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*JobTitle, error)
	predicates    []predicate.JobTitle
}

var _ ent.Mutation = (*JobTitleMutation)(nil)

// jobtitleOption allows management of the mutation configuration using functional options.
type jobtitleOption func(*JobTitleMutation)

// newJobTitleMutation creates new mutation for the JobTitle entity.
func newJobTitleMutation(c config, op Op, opts ...jobtitleOption) *JobTitleMutation {
	m := &JobTitleMutation{
		config:        c,
		op:            op,
		typ:           TypeJobTitle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withJobTitleID sets the ID field of the mutation.
func withJobTitleID(id uuid.UUID) jobtitleOption {
	return func(m *JobTitleMutation) {
		var (
			err   error
			once  sync.Once
			value *JobTitle
		)
		m.oldValue = func(ctx context.Context) (*JobTitle, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().JobTitle.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withJobTitle sets the old JobTitle of the mutation.
func withJobTitle(node *JobTitle) jobtitleOption {
	return func(m *JobTitleMutation) {
		m.oldValue = func(context.Context) (*JobTitle, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m JobTitleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m JobTitleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of JobTitle entities.
func (m *JobTitleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *JobTitleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *JobTitleMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().JobTitle.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCode sets the "code" field.
func (m *JobTitleMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *JobTitleMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *JobTitleMutation) ResetCode() {
	m.code = nil
}

// SetName sets the "name" field.
func (m *JobTitleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *JobTitleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *JobTitleMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *JobTitleMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *JobTitleMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *JobTitleMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[jobtitle.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *JobTitleMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[jobtitle.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *JobTitleMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, jobtitle.FieldDescription)
}

// SetSpecification sets the "specification" field.
func (m *JobTitleMutation) SetSpecification(s string) {
	m.specification = &s
}

// Specification returns the value of the "specification" field in the mutation.
func (m *JobTitleMutation) Specification() (r string, exists bool) {
	v := m.specification
	if v == nil {
		return
	}
	return *v, true
}

// OldSpecification returns the old "specification" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldSpecification(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSpecification is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSpecification requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpecification: %w", err)
	}
	return oldValue.Specification, nil
}

// ClearSpecification clears the value of the "specification" field.
func (m *JobTitleMutation) ClearSpecification() {
	m.specification = nil
	m.clearedFields[jobtitle.FieldSpecification] = struct{}{}
}

// SpecificationCleared returns if the "specification" field was cleared in this mutation.
func (m *JobTitleMutation) SpecificationCleared() bool {
	_, ok := m.clearedFields[jobtitle.FieldSpecification]
	return ok
}

// ResetSpecification resets all changes to the "specification" field.
func (m *JobTitleMutation) ResetSpecification() {
	m.specification = nil
	delete(m.clearedFields, jobtitle.FieldSpecification)
}

// SetCreatedAt sets the "created_at" field.
func (m *JobTitleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *JobTitleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *JobTitleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *JobTitleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *JobTitleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *JobTitleMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[jobtitle.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *JobTitleMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[jobtitle.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *JobTitleMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, jobtitle.FieldUpdatedAt)
}

// SetDeletedAt sets the "deleted_at" field.
func (m *JobTitleMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *JobTitleMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the JobTitle entity.
// If the JobTitle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobTitleMutation) OldDeletedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (m *JobTitleMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[jobtitle.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the "deleted_at" field was cleared in this mutation.
func (m *JobTitleMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[jobtitle.FieldDeletedAt]
	return ok
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *JobTitleMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, jobtitle.FieldDeletedAt)
}

// Where appends a list predicates to the JobTitleMutation builder.
func (m *JobTitleMutation) Where(ps ...predicate.JobTitle) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *JobTitleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (JobTitle).
func (m *JobTitleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *JobTitleMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.code != nil {
		fields = append(fields, jobtitle.FieldCode)
	}
	if m.name != nil {
		fields = append(fields, jobtitle.FieldName)
	}
	if m.description != nil {
		fields = append(fields, jobtitle.FieldDescription)
	}
	if m.specification != nil {
		fields = append(fields, jobtitle.FieldSpecification)
	}
	if m.created_at != nil {
		fields = append(fields, jobtitle.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, jobtitle.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, jobtitle.FieldDeletedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *JobTitleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case jobtitle.FieldCode:
		return m.Code()
	case jobtitle.FieldName:
		return m.Name()
	case jobtitle.FieldDescription:
		return m.Description()
	case jobtitle.FieldSpecification:
		return m.Specification()
	case jobtitle.FieldCreatedAt:
		return m.CreatedAt()
	case jobtitle.FieldUpdatedAt:
		return m.UpdatedAt()
	case jobtitle.FieldDeletedAt:
		return m.DeletedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *JobTitleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case jobtitle.FieldCode:
		return m.OldCode(ctx)
	case jobtitle.FieldName:
		return m.OldName(ctx)
	case jobtitle.FieldDescription:
		return m.OldDescription(ctx)
	case jobtitle.FieldSpecification:
		return m.OldSpecification(ctx)
	case jobtitle.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case jobtitle.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case jobtitle.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	}
	return nil, fmt.Errorf("unknown JobTitle field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobTitleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case jobtitle.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case jobtitle.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case jobtitle.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case jobtitle.FieldSpecification:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpecification(v)
		return nil
	case jobtitle.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case jobtitle.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case jobtitle.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown JobTitle field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *JobTitleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *JobTitleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobTitleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown JobTitle numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *JobTitleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(jobtitle.FieldDescription) {
		fields = append(fields, jobtitle.FieldDescription)
	}
	if m.FieldCleared(jobtitle.FieldSpecification) {
		fields = append(fields, jobtitle.FieldSpecification)
	}
	if m.FieldCleared(jobtitle.FieldUpdatedAt) {
		fields = append(fields, jobtitle.FieldUpdatedAt)
	}
	if m.FieldCleared(jobtitle.FieldDeletedAt) {
		fields = append(fields, jobtitle.FieldDeletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *JobTitleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *JobTitleMutation) ClearField(name string) error {
	switch name {
	case jobtitle.FieldDescription:
		m.ClearDescription()
		return nil
	case jobtitle.FieldSpecification:
		m.ClearSpecification()
		return nil
	case jobtitle.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	case jobtitle.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown JobTitle nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *JobTitleMutation) ResetField(name string) error {
	switch name {
	case jobtitle.FieldCode:
		m.ResetCode()
		return nil
	case jobtitle.FieldName:
		m.ResetName()
		return nil
	case jobtitle.FieldDescription:
		m.ResetDescription()
		return nil
	case jobtitle.FieldSpecification:
		m.ResetSpecification()
		return nil
	case jobtitle.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case jobtitle.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case jobtitle.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown JobTitle field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *JobTitleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *JobTitleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *JobTitleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *JobTitleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *JobTitleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *JobTitleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *JobTitleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown JobTitle unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *JobTitleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown JobTitle edge %s", name)
}

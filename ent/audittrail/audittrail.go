// Code generated by ent, DO NOT EDIT.

package audittrail

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the audittrail type in the database.
	Label = "audit_trail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldRecordId holds the string denoting the recordid field in the database.
	FieldRecordId = "record_id"
	// FieldModule holds the string denoting the module field in the database.
	FieldModule = "module"
	// FieldActionType holds the string denoting the actiontype field in the database.
	FieldActionType = "action_type"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldRecordChanges holds the string denoting the record_changes field in the database.
	FieldRecordChanges = "record_changes"
	// EdgeUserEdge holds the string denoting the user_edge edge name in mutations.
	EdgeUserEdge = "user_edge"
	// Table holds the table name of the audittrail in the database.
	Table = "audit_trails"
	// UserEdgeTable is the table that holds the user_edge relation/edge.
	UserEdgeTable = "audit_trails"
	// UserEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserEdgeInverseTable = "users"
	// UserEdgeColumn is the table column denoting the user_edge relation/edge.
	UserEdgeColumn = "created_by"
)

// Columns holds all SQL columns for audittrail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCreatedBy,
	FieldRecordId,
	FieldModule,
	FieldActionType,
	FieldNote,
	FieldRecordChanges,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Module defines the type for the "module" enum field.
type Module string

// Module values.
const (
	ModuleTeams Module = "teams"
)

func (m Module) String() string {
	return string(m)
}

// ModuleValidator is a validator for the "module" field enum values. It is called by the builders before save.
func ModuleValidator(m Module) error {
	switch m {
	case ModuleTeams:
		return nil
	default:
		return fmt.Errorf("audittrail: invalid enum value for module field: %q", m)
	}
}

// ActionType defines the type for the "actionType" enum field.
type ActionType string

// ActionTypeCreate is the default value of the ActionType enum.
const DefaultActionType = ActionTypeCreate

// ActionType values.
const (
	ActionTypeCreate ActionType = "create"
	ActionTypeUpdate ActionType = "update"
	ActionTypeDelete ActionType = "delete"
)

func (at ActionType) String() string {
	return string(at)
}

// ActionTypeValidator is a validator for the "actionType" field enum values. It is called by the builders before save.
func ActionTypeValidator(at ActionType) error {
	switch at {
	case ActionTypeCreate, ActionTypeUpdate, ActionTypeDelete:
		return nil
	default:
		return fmt.Errorf("audittrail: invalid enum value for actionType field: %q", at)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Module) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Module) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Module(str)
	if err := ModuleValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Module", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e ActionType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *ActionType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = ActionType(str)
	if err := ActionTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

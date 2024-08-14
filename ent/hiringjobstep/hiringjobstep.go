// Code generated by ent, DO NOT EDIT.

package hiringjobstep

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the hiringjobstep type in the database.
	Label = "hiring_job_step"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHiringJobID holds the string denoting the hiring_job_id field in the database.
	FieldHiringJobID = "hiring_job_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeHiringJobEdge holds the string denoting the hiring_job_edge edge name in mutations.
	EdgeHiringJobEdge = "hiring_job_edge"
	// Table holds the table name of the hiringjobstep in the database.
	Table = "hiring_job_steps"
	// HiringJobEdgeTable is the table that holds the hiring_job_edge relation/edge.
	HiringJobEdgeTable = "hiring_job_steps"
	// HiringJobEdgeInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringJobEdgeInverseTable = "hiring_jobs"
	// HiringJobEdgeColumn is the table column denoting the hiring_job_edge relation/edge.
	HiringJobEdgeColumn = "hiring_job_id"
)

// Columns holds all SQL columns for hiringjobstep fields.
var Columns = []string{
	FieldID,
	FieldHiringJobID,
	FieldType,
	FieldCreatedAt,
	FieldUpdatedAt,
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
)

// Type defines the type for the "type" enum field.
type Type string

// TypeCreated is the default value of the Type enum.
const DefaultType = TypeCreated

// Type values.
const (
	TypeCreated Type = "created"
	TypeOpened  Type = "opened"
	TypeClosed  Type = "closed"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCreated, TypeOpened, TypeClosed:
		return nil
	default:
		return fmt.Errorf("hiringjobstep: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

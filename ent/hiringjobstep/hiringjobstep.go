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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeApprovalJob holds the string denoting the approval_job edge name in mutations.
	EdgeApprovalJob = "approval_job"
	// EdgeApprovalUser holds the string denoting the approval_user edge name in mutations.
	EdgeApprovalUser = "approval_user"
	// Table holds the table name of the hiringjobstep in the database.
	Table = "hiring_job_steps"
	// ApprovalJobTable is the table that holds the approval_job relation/edge.
	ApprovalJobTable = "hiring_job_steps"
	// ApprovalJobInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	ApprovalJobInverseTable = "hiring_jobs"
	// ApprovalJobColumn is the table column denoting the approval_job relation/edge.
	ApprovalJobColumn = "hiring_job_id"
	// ApprovalUserTable is the table that holds the approval_user relation/edge.
	ApprovalUserTable = "hiring_job_steps"
	// ApprovalUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ApprovalUserInverseTable = "users"
	// ApprovalUserColumn is the table column denoting the approval_user relation/edge.
	ApprovalUserColumn = "user_id"
)

// Columns holds all SQL columns for hiringjobstep fields.
var Columns = []string{
	FieldID,
	FieldHiringJobID,
	FieldUserID,
	FieldStatus,
	FieldOrderID,
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
	// OrderIDValidator is a validator for the "order_id" field. It is called by the builders before save.
	OrderIDValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusWaiting  Status = "waiting"
	StatusPending  Status = "pending"
	StatusAccepted Status = "accepted"
	StatusRejected Status = "rejected"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusWaiting, StatusPending, StatusAccepted, StatusRejected:
		return nil
	default:
		return fmt.Errorf("hiringjobstep: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

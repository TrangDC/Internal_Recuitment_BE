// Code generated by ent, DO NOT EDIT.

package candidate

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the candidate type in the database.
	Label = "candidate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldDob holds the string denoting the dob field in the database.
	FieldDob = "dob"
	// FieldIsBlacklist holds the string denoting the is_blacklist field in the database.
	FieldIsBlacklist = "is_blacklist"
	// FieldLastApplyDate holds the string denoting the last_apply_date field in the database.
	FieldLastApplyDate = "last_apply_date"
	// EdgeCandidateJobEdges holds the string denoting the candidate_job_edges edge name in mutations.
	EdgeCandidateJobEdges = "candidate_job_edges"
	// Table holds the table name of the candidate in the database.
	Table = "candidates"
	// CandidateJobEdgesTable is the table that holds the candidate_job_edges relation/edge.
	CandidateJobEdgesTable = "candidate_jobs"
	// CandidateJobEdgesInverseTable is the table name for the CandidateJob entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejob" package.
	CandidateJobEdgesInverseTable = "candidate_jobs"
	// CandidateJobEdgesColumn is the table column denoting the candidate_job_edges relation/edge.
	CandidateJobEdgesColumn = "candidate_id"
)

// Columns holds all SQL columns for candidate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldEmail,
	FieldPhone,
	FieldDob,
	FieldIsBlacklist,
	FieldLastApplyDate,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultIsBlacklist holds the default value on creation for the "is_blacklist" field.
	DefaultIsBlacklist bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

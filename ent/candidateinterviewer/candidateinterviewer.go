// Code generated by ent, DO NOT EDIT.

package candidateinterviewer

import (
	"time"
)

const (
	// Label holds the string label denoting the candidateinterviewer type in the database.
	Label = "candidate_interviewer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCandidateInterviewID holds the string denoting the candidate_interview_id field in the database.
	FieldCandidateInterviewID = "candidate_interview_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUserEdge holds the string denoting the user_edge edge name in mutations.
	EdgeUserEdge = "user_edge"
	// EdgeInterviewEdge holds the string denoting the interview_edge edge name in mutations.
	EdgeInterviewEdge = "interview_edge"
	// Table holds the table name of the candidateinterviewer in the database.
	Table = "candidate_interviewers"
	// UserEdgeTable is the table that holds the user_edge relation/edge.
	UserEdgeTable = "candidate_interviewers"
	// UserEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserEdgeInverseTable = "users"
	// UserEdgeColumn is the table column denoting the user_edge relation/edge.
	UserEdgeColumn = "user_id"
	// InterviewEdgeTable is the table that holds the interview_edge relation/edge.
	InterviewEdgeTable = "candidate_interviewers"
	// InterviewEdgeInverseTable is the table name for the CandidateInterview entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterview" package.
	InterviewEdgeInverseTable = "candidate_interviews"
	// InterviewEdgeColumn is the table column denoting the interview_edge relation/edge.
	InterviewEdgeColumn = "candidate_interview_id"
)

// Columns holds all SQL columns for candidateinterviewer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCandidateInterviewID,
	FieldUserID,
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

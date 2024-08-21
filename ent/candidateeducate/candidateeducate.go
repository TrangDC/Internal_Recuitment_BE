// Code generated by ent, DO NOT EDIT.

package candidateeducate

import (
	"time"
)

const (
	// Label holds the string label denoting the candidateeducate type in the database.
	Label = "candidate_educate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCandidateID holds the string denoting the candidate_id field in the database.
	FieldCandidateID = "candidate_id"
	// FieldSchoolName holds the string denoting the school_name field in the database.
	FieldSchoolName = "school_name"
	// FieldMajor holds the string denoting the major field in the database.
	FieldMajor = "major"
	// FieldGpa holds the string denoting the gpa field in the database.
	FieldGpa = "gpa"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldIsCurrent holds the string denoting the is_current field in the database.
	FieldIsCurrent = "is_current"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// Table holds the table name of the candidateeducate in the database.
	Table = "candidate_educates"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "relation_id"
	// CandidateEdgeTable is the table that holds the candidate_edge relation/edge.
	CandidateEdgeTable = "candidate_educates"
	// CandidateEdgeInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateEdgeInverseTable = "candidates"
	// CandidateEdgeColumn is the table column denoting the candidate_edge relation/edge.
	CandidateEdgeColumn = "candidate_id"
)

// Columns holds all SQL columns for candidateeducate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCandidateID,
	FieldSchoolName,
	FieldMajor,
	FieldGpa,
	FieldLocation,
	FieldDescription,
	FieldStartDate,
	FieldEndDate,
	FieldOrderID,
	FieldIsCurrent,
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
	// SchoolNameValidator is a validator for the "school_name" field. It is called by the builders before save.
	SchoolNameValidator func(string) error
	// DefaultIsCurrent holds the default value on creation for the "is_current" field.
	DefaultIsCurrent bool
)

// Code generated by ent, DO NOT EDIT.

package candidateexp

import (
	"time"
)

const (
	// Label holds the string label denoting the candidateexp type in the database.
	Label = "candidate_exp"
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
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// Table holds the table name of the candidateexp in the database.
	Table = "candidate_exps"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "candidate_exp_attachment_edges"
	// CandidateEdgeTable is the table that holds the candidate_edge relation/edge.
	CandidateEdgeTable = "candidate_exps"
	// CandidateEdgeInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateEdgeInverseTable = "candidates"
	// CandidateEdgeColumn is the table column denoting the candidate_edge relation/edge.
	CandidateEdgeColumn = "candidate_id"
)

// Columns holds all SQL columns for candidateexp fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCandidateID,
	FieldPosition,
	FieldCompany,
	FieldLocation,
	FieldDescription,
	FieldStartDate,
	FieldEndDate,
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
	// PositionValidator is a validator for the "position" field. It is called by the builders before save.
	PositionValidator func(string) error
	// CompanyValidator is a validator for the "company" field. It is called by the builders before save.
	CompanyValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

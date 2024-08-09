// Code generated by ent, DO NOT EDIT.

package candidatecertificate

import (
	"time"
)

const (
	// Label holds the string label denoting the candidatecertificate type in the database.
	Label = "candidate_certificate"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldAchievedDate holds the string denoting the achieved_date field in the database.
	FieldAchievedDate = "achieved_date"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// Table holds the table name of the candidatecertificate in the database.
	Table = "candidate_certificates"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "relation_id"
	// CandidateEdgeTable is the table that holds the candidate_edge relation/edge.
	CandidateEdgeTable = "candidate_certificates"
	// CandidateEdgeInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateEdgeInverseTable = "candidates"
	// CandidateEdgeColumn is the table column denoting the candidate_edge relation/edge.
	CandidateEdgeColumn = "candidate_id"
)

// Columns holds all SQL columns for candidatecertificate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCandidateID,
	FieldName,
	FieldScore,
	FieldAchievedDate,
	FieldOrderID,
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
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(string) error
)

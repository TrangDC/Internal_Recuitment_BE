// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the attachment type in the database.
	Label = "attachment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDocumentID holds the string denoting the document_id field in the database.
	FieldDocumentID = "document_id"
	// FieldDocumentName holds the string denoting the document_name field in the database.
	FieldDocumentName = "document_name"
	// FieldRelationType holds the string denoting the relation_type field in the database.
	FieldRelationType = "relation_type"
	// FieldRelationID holds the string denoting the relation_id field in the database.
	FieldRelationID = "relation_id"
	// EdgeCandidateJobEdge holds the string denoting the candidate_job_edge edge name in mutations.
	EdgeCandidateJobEdge = "candidate_job_edge"
	// EdgeCandidateJobFeedbackEdge holds the string denoting the candidate_job_feedback_edge edge name in mutations.
	EdgeCandidateJobFeedbackEdge = "candidate_job_feedback_edge"
	// EdgeCandidateInterviewEdge holds the string denoting the candidate_interview_edge edge name in mutations.
	EdgeCandidateInterviewEdge = "candidate_interview_edge"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// EdgeCandidateEducateEdge holds the string denoting the candidate_educate_edge edge name in mutations.
	EdgeCandidateEducateEdge = "candidate_educate_edge"
	// EdgeCandidateAwardEdge holds the string denoting the candidate_award_edge edge name in mutations.
	EdgeCandidateAwardEdge = "candidate_award_edge"
	// EdgeCandidateCertificateEdge holds the string denoting the candidate_certificate_edge edge name in mutations.
	EdgeCandidateCertificateEdge = "candidate_certificate_edge"
	// EdgeCandidateHistoryCallEdge holds the string denoting the candidate_history_call_edge edge name in mutations.
	EdgeCandidateHistoryCallEdge = "candidate_history_call_edge"
	// Table holds the table name of the attachment in the database.
	Table = "attachments"
	// CandidateJobEdgeTable is the table that holds the candidate_job_edge relation/edge.
	CandidateJobEdgeTable = "attachments"
	// CandidateJobEdgeInverseTable is the table name for the CandidateJob entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejob" package.
	CandidateJobEdgeInverseTable = "candidate_jobs"
	// CandidateJobEdgeColumn is the table column denoting the candidate_job_edge relation/edge.
	CandidateJobEdgeColumn = "relation_id"
	// CandidateJobFeedbackEdgeTable is the table that holds the candidate_job_feedback_edge relation/edge.
	CandidateJobFeedbackEdgeTable = "attachments"
	// CandidateJobFeedbackEdgeInverseTable is the table name for the CandidateJobFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejobfeedback" package.
	CandidateJobFeedbackEdgeInverseTable = "candidate_job_feedbacks"
	// CandidateJobFeedbackEdgeColumn is the table column denoting the candidate_job_feedback_edge relation/edge.
	CandidateJobFeedbackEdgeColumn = "relation_id"
	// CandidateInterviewEdgeTable is the table that holds the candidate_interview_edge relation/edge.
	CandidateInterviewEdgeTable = "attachments"
	// CandidateInterviewEdgeInverseTable is the table name for the CandidateInterview entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterview" package.
	CandidateInterviewEdgeInverseTable = "candidate_interviews"
	// CandidateInterviewEdgeColumn is the table column denoting the candidate_interview_edge relation/edge.
	CandidateInterviewEdgeColumn = "relation_id"
	// CandidateEdgeTable is the table that holds the candidate_edge relation/edge.
	CandidateEdgeTable = "attachments"
	// CandidateEdgeInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateEdgeInverseTable = "candidates"
	// CandidateEdgeColumn is the table column denoting the candidate_edge relation/edge.
	CandidateEdgeColumn = "relation_id"
	// CandidateEducateEdgeTable is the table that holds the candidate_educate_edge relation/edge.
	CandidateEducateEdgeTable = "attachments"
	// CandidateEducateEdgeInverseTable is the table name for the CandidateEducate entity.
	// It exists in this package in order to avoid circular dependency with the "candidateeducate" package.
	CandidateEducateEdgeInverseTable = "candidate_educates"
	// CandidateEducateEdgeColumn is the table column denoting the candidate_educate_edge relation/edge.
	CandidateEducateEdgeColumn = "relation_id"
	// CandidateAwardEdgeTable is the table that holds the candidate_award_edge relation/edge.
	CandidateAwardEdgeTable = "attachments"
	// CandidateAwardEdgeInverseTable is the table name for the CandidateAward entity.
	// It exists in this package in order to avoid circular dependency with the "candidateaward" package.
	CandidateAwardEdgeInverseTable = "candidate_awards"
	// CandidateAwardEdgeColumn is the table column denoting the candidate_award_edge relation/edge.
	CandidateAwardEdgeColumn = "relation_id"
	// CandidateCertificateEdgeTable is the table that holds the candidate_certificate_edge relation/edge.
	CandidateCertificateEdgeTable = "attachments"
	// CandidateCertificateEdgeInverseTable is the table name for the CandidateCertificate entity.
	// It exists in this package in order to avoid circular dependency with the "candidatecertificate" package.
	CandidateCertificateEdgeInverseTable = "candidate_certificates"
	// CandidateCertificateEdgeColumn is the table column denoting the candidate_certificate_edge relation/edge.
	CandidateCertificateEdgeColumn = "relation_id"
	// CandidateHistoryCallEdgeTable is the table that holds the candidate_history_call_edge relation/edge.
	CandidateHistoryCallEdgeTable = "attachments"
	// CandidateHistoryCallEdgeInverseTable is the table name for the CandidateHistoryCall entity.
	// It exists in this package in order to avoid circular dependency with the "candidatehistorycall" package.
	CandidateHistoryCallEdgeInverseTable = "candidate_history_calls"
	// CandidateHistoryCallEdgeColumn is the table column denoting the candidate_history_call_edge relation/edge.
	CandidateHistoryCallEdgeColumn = "relation_id"
)

// Columns holds all SQL columns for attachment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldDocumentID,
	FieldDocumentName,
	FieldRelationType,
	FieldRelationID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attachments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"candidate_exp_attachment_edges",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DocumentNameValidator is a validator for the "document_name" field. It is called by the builders before save.
	DocumentNameValidator func(string) error
)

// RelationType defines the type for the "relation_type" enum field.
type RelationType string

// RelationType values.
const (
	RelationTypeCandidateJobs         RelationType = "candidate_jobs"
	RelationTypeCandidateJobFeedbacks RelationType = "candidate_job_feedbacks"
	RelationTypeCandidates            RelationType = "candidates"
	RelationTypeCandidateEducates     RelationType = "candidate_educates"
	RelationTypeCandidateAwards       RelationType = "candidate_awards"
	RelationTypeCandidateCertificates RelationType = "candidate_certificates"
)

func (rt RelationType) String() string {
	return string(rt)
}

// RelationTypeValidator is a validator for the "relation_type" field enum values. It is called by the builders before save.
func RelationTypeValidator(rt RelationType) error {
	switch rt {
	case RelationTypeCandidateJobs, RelationTypeCandidateJobFeedbacks, RelationTypeCandidates, RelationTypeCandidateEducates, RelationTypeCandidateAwards, RelationTypeCandidateCertificates:
		return nil
	default:
		return fmt.Errorf("attachment: invalid enum value for relation_type field: %q", rt)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e RelationType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *RelationType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = RelationType(str)
	if err := RelationTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid RelationType", str)
	}
	return nil
}

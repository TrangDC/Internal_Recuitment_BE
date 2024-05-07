// Code generated by ent, DO NOT EDIT.

package candidatejob

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the candidatejob type in the database.
	Label = "candidate_job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldHiringJobID holds the string denoting the hiring_job_id field in the database.
	FieldHiringJobID = "hiring_job_id"
	// FieldCandidateID holds the string denoting the candidate_id field in the database.
	FieldCandidateID = "candidate_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeHiringJob holds the string denoting the hiring_job edge name in mutations.
	EdgeHiringJob = "hiring_job"
	// EdgeCandidateJobFeedback holds the string denoting the candidate_job_feedback edge name in mutations.
	EdgeCandidateJobFeedback = "candidate_job_feedback"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// EdgeCandidateJobInterview holds the string denoting the candidate_job_interview edge name in mutations.
	EdgeCandidateJobInterview = "candidate_job_interview"
	// Table holds the table name of the candidatejob in the database.
	Table = "candidate_jobs"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "relation_id"
	// HiringJobTable is the table that holds the hiring_job relation/edge.
	HiringJobTable = "candidate_jobs"
	// HiringJobInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringJobInverseTable = "hiring_jobs"
	// HiringJobColumn is the table column denoting the hiring_job relation/edge.
	HiringJobColumn = "hiring_job_id"
	// CandidateJobFeedbackTable is the table that holds the candidate_job_feedback relation/edge.
	CandidateJobFeedbackTable = "candidate_job_feedbacks"
	// CandidateJobFeedbackInverseTable is the table name for the CandidateJobFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejobfeedback" package.
	CandidateJobFeedbackInverseTable = "candidate_job_feedbacks"
	// CandidateJobFeedbackColumn is the table column denoting the candidate_job_feedback relation/edge.
	CandidateJobFeedbackColumn = "candidate_job_id"
	// CandidateEdgeTable is the table that holds the candidate_edge relation/edge.
	CandidateEdgeTable = "candidate_jobs"
	// CandidateEdgeInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateEdgeInverseTable = "candidates"
	// CandidateEdgeColumn is the table column denoting the candidate_edge relation/edge.
	CandidateEdgeColumn = "candidate_id"
	// CandidateJobInterviewTable is the table that holds the candidate_job_interview relation/edge.
	CandidateJobInterviewTable = "candidate_interviews"
	// CandidateJobInterviewInverseTable is the table name for the CandidateInterview entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterview" package.
	CandidateJobInterviewInverseTable = "candidate_interviews"
	// CandidateJobInterviewColumn is the table column denoting the candidate_job_interview relation/edge.
	CandidateJobInterviewColumn = "candidate_job_id"
)

// Columns holds all SQL columns for candidatejob fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHiringJobID,
	FieldCandidateID,
	FieldStatus,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusApplied is the default value of the Status enum.
const DefaultStatus = StatusApplied

// Status values.
const (
	StatusApplied      Status = "applied"
	StatusInterviewing Status = "interviewing"
	StatusOffering     Status = "offering"
	StatusHired        Status = "hired"
	StatusKiv          Status = "kiv"
	StatusOfferLost    Status = "offer_lost"
	StatusExStaff      Status = "ex_staff"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusApplied, StatusInterviewing, StatusOffering, StatusHired, StatusKiv, StatusOfferLost, StatusExStaff:
		return nil
	default:
		return fmt.Errorf("candidatejob: invalid enum value for status field: %q", s)
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
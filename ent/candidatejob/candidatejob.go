// Code generated by ent, DO NOT EDIT.

package candidatejob

import (
	"fmt"
	"io"
	"strconv"
	"time"
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
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldRecInChargeID holds the string denoting the rec_in_charge_id field in the database.
	FieldRecInChargeID = "rec_in_charge_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldFailedReason holds the string denoting the failed_reason field in the database.
	FieldFailedReason = "failed_reason"
	// FieldOnboardDate holds the string denoting the onboard_date field in the database.
	FieldOnboardDate = "onboard_date"
	// FieldOfferExpirationDate holds the string denoting the offer_expiration_date field in the database.
	FieldOfferExpirationDate = "offer_expiration_date"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeHiringJobEdge holds the string denoting the hiring_job_edge edge name in mutations.
	EdgeHiringJobEdge = "hiring_job_edge"
	// EdgeCandidateJobFeedback holds the string denoting the candidate_job_feedback edge name in mutations.
	EdgeCandidateJobFeedback = "candidate_job_feedback"
	// EdgeCandidateEdge holds the string denoting the candidate_edge edge name in mutations.
	EdgeCandidateEdge = "candidate_edge"
	// EdgeCandidateJobInterview holds the string denoting the candidate_job_interview edge name in mutations.
	EdgeCandidateJobInterview = "candidate_job_interview"
	// EdgeCreatedByEdge holds the string denoting the created_by_edge edge name in mutations.
	EdgeCreatedByEdge = "created_by_edge"
	// EdgeCandidateJobStep holds the string denoting the candidate_job_step edge name in mutations.
	EdgeCandidateJobStep = "candidate_job_step"
	// EdgeRecInChargeEdge holds the string denoting the rec_in_charge_edge edge name in mutations.
	EdgeRecInChargeEdge = "rec_in_charge_edge"
	// Table holds the table name of the candidatejob in the database.
	Table = "candidate_jobs"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "relation_id"
	// HiringJobEdgeTable is the table that holds the hiring_job_edge relation/edge.
	HiringJobEdgeTable = "candidate_jobs"
	// HiringJobEdgeInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringJobEdgeInverseTable = "hiring_jobs"
	// HiringJobEdgeColumn is the table column denoting the hiring_job_edge relation/edge.
	HiringJobEdgeColumn = "hiring_job_id"
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
	// CreatedByEdgeTable is the table that holds the created_by_edge relation/edge.
	CreatedByEdgeTable = "candidate_jobs"
	// CreatedByEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByEdgeInverseTable = "users"
	// CreatedByEdgeColumn is the table column denoting the created_by_edge relation/edge.
	CreatedByEdgeColumn = "created_by"
	// CandidateJobStepTable is the table that holds the candidate_job_step relation/edge.
	CandidateJobStepTable = "candidate_job_steps"
	// CandidateJobStepInverseTable is the table name for the CandidateJobStep entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejobstep" package.
	CandidateJobStepInverseTable = "candidate_job_steps"
	// CandidateJobStepColumn is the table column denoting the candidate_job_step relation/edge.
	CandidateJobStepColumn = "candidate_job_id"
	// RecInChargeEdgeTable is the table that holds the rec_in_charge_edge relation/edge.
	RecInChargeEdgeTable = "candidate_jobs"
	// RecInChargeEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecInChargeEdgeInverseTable = "users"
	// RecInChargeEdgeColumn is the table column denoting the rec_in_charge_edge relation/edge.
	RecInChargeEdgeColumn = "rec_in_charge_id"
)

// Columns holds all SQL columns for candidatejob fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHiringJobID,
	FieldCandidateID,
	FieldCreatedBy,
	FieldRecInChargeID,
	FieldStatus,
	FieldFailedReason,
	FieldOnboardDate,
	FieldOfferExpirationDate,
	FieldLevel,
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

// Status defines the type for the "status" enum field.
type Status string

// StatusApplied is the default value of the Status enum.
const DefaultStatus = StatusApplied

// Status values.
const (
	StatusApplied         Status = "applied"
	StatusInterviewing    Status = "interviewing"
	StatusOffering        Status = "offering"
	StatusHired           Status = "hired"
	StatusFailedCv        Status = "failed_cv"
	StatusFailedInterview Status = "failed_interview"
	StatusOfferLost       Status = "offer_lost"
	StatusExStaff         Status = "ex_staff"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusApplied, StatusInterviewing, StatusOffering, StatusHired, StatusFailedCv, StatusFailedInterview, StatusOfferLost, StatusExStaff:
		return nil
	default:
		return fmt.Errorf("candidatejob: invalid enum value for status field: %q", s)
	}
}

// Level defines the type for the "level" enum field.
type Level string

// Level values.
const (
	LevelIntern   Level = "intern"
	LevelFresher  Level = "fresher"
	LevelJunior   Level = "junior"
	LevelMiddle   Level = "middle"
	LevelSenior   Level = "senior"
	LevelManager  Level = "manager"
	LevelDirector Level = "director"
)

func (l Level) String() string {
	return string(l)
}

// LevelValidator is a validator for the "level" field enum values. It is called by the builders before save.
func LevelValidator(l Level) error {
	switch l {
	case LevelIntern, LevelFresher, LevelJunior, LevelMiddle, LevelSenior, LevelManager, LevelDirector:
		return nil
	default:
		return fmt.Errorf("candidatejob: invalid enum value for level field: %q", l)
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

// MarshalGQL implements graphql.Marshaler interface.
func (e Level) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Level) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Level(str)
	if err := LevelValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Level", str)
	}
	return nil
}

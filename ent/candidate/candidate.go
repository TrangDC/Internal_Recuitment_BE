// Code generated by ent, DO NOT EDIT.

package candidate

import (
	"fmt"
	"io"
	"strconv"
	"time"
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
	// FieldReferenceType holds the string denoting the reference_type field in the database.
	FieldReferenceType = "reference_type"
	// FieldReferenceValue holds the string denoting the reference_value field in the database.
	FieldReferenceValue = "reference_value"
	// FieldReferenceUID holds the string denoting the reference_uid field in the database.
	FieldReferenceUID = "reference_uid"
	// FieldRecruitTime holds the string denoting the recruit_time field in the database.
	FieldRecruitTime = "recruit_time"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// EdgeCandidateJobEdges holds the string denoting the candidate_job_edges edge name in mutations.
	EdgeCandidateJobEdges = "candidate_job_edges"
	// EdgeReferenceUserEdge holds the string denoting the reference_user_edge edge name in mutations.
	EdgeReferenceUserEdge = "reference_user_edge"
	// EdgeAttachmentEdges holds the string denoting the attachment_edges edge name in mutations.
	EdgeAttachmentEdges = "attachment_edges"
	// EdgeCandidateSkillEdges holds the string denoting the candidate_skill_edges edge name in mutations.
	EdgeCandidateSkillEdges = "candidate_skill_edges"
	// EdgeCandidateExpEdges holds the string denoting the candidate_exp_edges edge name in mutations.
	EdgeCandidateExpEdges = "candidate_exp_edges"
	// EdgeCandidateEducateEdges holds the string denoting the candidate_educate_edges edge name in mutations.
	EdgeCandidateEducateEdges = "candidate_educate_edges"
	// EdgeCandidateAwardEdges holds the string denoting the candidate_award_edges edge name in mutations.
	EdgeCandidateAwardEdges = "candidate_award_edges"
	// EdgeCandidateCertificateEdges holds the string denoting the candidate_certificate_edges edge name in mutations.
	EdgeCandidateCertificateEdges = "candidate_certificate_edges"
	// EdgeCandidateHistoryCallEdges holds the string denoting the candidate_history_call_edges edge name in mutations.
	EdgeCandidateHistoryCallEdges = "candidate_history_call_edges"
	// EdgeCandidateNoteEdges holds the string denoting the candidate_note_edges edge name in mutations.
	EdgeCandidateNoteEdges = "candidate_note_edges"
	// EdgeOutgoingEmailEdges holds the string denoting the outgoing_email_edges edge name in mutations.
	EdgeOutgoingEmailEdges = "outgoing_email_edges"
	// Table holds the table name of the candidate in the database.
	Table = "candidates"
	// CandidateJobEdgesTable is the table that holds the candidate_job_edges relation/edge.
	CandidateJobEdgesTable = "candidate_jobs"
	// CandidateJobEdgesInverseTable is the table name for the CandidateJob entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejob" package.
	CandidateJobEdgesInverseTable = "candidate_jobs"
	// CandidateJobEdgesColumn is the table column denoting the candidate_job_edges relation/edge.
	CandidateJobEdgesColumn = "candidate_id"
	// ReferenceUserEdgeTable is the table that holds the reference_user_edge relation/edge.
	ReferenceUserEdgeTable = "candidates"
	// ReferenceUserEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReferenceUserEdgeInverseTable = "users"
	// ReferenceUserEdgeColumn is the table column denoting the reference_user_edge relation/edge.
	ReferenceUserEdgeColumn = "reference_uid"
	// AttachmentEdgesTable is the table that holds the attachment_edges relation/edge.
	AttachmentEdgesTable = "attachments"
	// AttachmentEdgesInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentEdgesInverseTable = "attachments"
	// AttachmentEdgesColumn is the table column denoting the attachment_edges relation/edge.
	AttachmentEdgesColumn = "relation_id"
	// CandidateSkillEdgesTable is the table that holds the candidate_skill_edges relation/edge.
	CandidateSkillEdgesTable = "entity_skills"
	// CandidateSkillEdgesInverseTable is the table name for the EntitySkill entity.
	// It exists in this package in order to avoid circular dependency with the "entityskill" package.
	CandidateSkillEdgesInverseTable = "entity_skills"
	// CandidateSkillEdgesColumn is the table column denoting the candidate_skill_edges relation/edge.
	CandidateSkillEdgesColumn = "entity_id"
	// CandidateExpEdgesTable is the table that holds the candidate_exp_edges relation/edge.
	CandidateExpEdgesTable = "candidate_exps"
	// CandidateExpEdgesInverseTable is the table name for the CandidateExp entity.
	// It exists in this package in order to avoid circular dependency with the "candidateexp" package.
	CandidateExpEdgesInverseTable = "candidate_exps"
	// CandidateExpEdgesColumn is the table column denoting the candidate_exp_edges relation/edge.
	CandidateExpEdgesColumn = "candidate_id"
	// CandidateEducateEdgesTable is the table that holds the candidate_educate_edges relation/edge.
	CandidateEducateEdgesTable = "candidate_educates"
	// CandidateEducateEdgesInverseTable is the table name for the CandidateEducate entity.
	// It exists in this package in order to avoid circular dependency with the "candidateeducate" package.
	CandidateEducateEdgesInverseTable = "candidate_educates"
	// CandidateEducateEdgesColumn is the table column denoting the candidate_educate_edges relation/edge.
	CandidateEducateEdgesColumn = "candidate_id"
	// CandidateAwardEdgesTable is the table that holds the candidate_award_edges relation/edge.
	CandidateAwardEdgesTable = "candidate_awards"
	// CandidateAwardEdgesInverseTable is the table name for the CandidateAward entity.
	// It exists in this package in order to avoid circular dependency with the "candidateaward" package.
	CandidateAwardEdgesInverseTable = "candidate_awards"
	// CandidateAwardEdgesColumn is the table column denoting the candidate_award_edges relation/edge.
	CandidateAwardEdgesColumn = "candidate_id"
	// CandidateCertificateEdgesTable is the table that holds the candidate_certificate_edges relation/edge.
	CandidateCertificateEdgesTable = "candidate_certificates"
	// CandidateCertificateEdgesInverseTable is the table name for the CandidateCertificate entity.
	// It exists in this package in order to avoid circular dependency with the "candidatecertificate" package.
	CandidateCertificateEdgesInverseTable = "candidate_certificates"
	// CandidateCertificateEdgesColumn is the table column denoting the candidate_certificate_edges relation/edge.
	CandidateCertificateEdgesColumn = "candidate_id"
	// CandidateHistoryCallEdgesTable is the table that holds the candidate_history_call_edges relation/edge.
	CandidateHistoryCallEdgesTable = "candidate_history_calls"
	// CandidateHistoryCallEdgesInverseTable is the table name for the CandidateHistoryCall entity.
	// It exists in this package in order to avoid circular dependency with the "candidatehistorycall" package.
	CandidateHistoryCallEdgesInverseTable = "candidate_history_calls"
	// CandidateHistoryCallEdgesColumn is the table column denoting the candidate_history_call_edges relation/edge.
	CandidateHistoryCallEdgesColumn = "candidate_id"
	// CandidateNoteEdgesTable is the table that holds the candidate_note_edges relation/edge.
	CandidateNoteEdgesTable = "candidate_notes"
	// CandidateNoteEdgesInverseTable is the table name for the CandidateNote entity.
	// It exists in this package in order to avoid circular dependency with the "candidatenote" package.
	CandidateNoteEdgesInverseTable = "candidate_notes"
	// CandidateNoteEdgesColumn is the table column denoting the candidate_note_edges relation/edge.
	CandidateNoteEdgesColumn = "candidate_id"
	// OutgoingEmailEdgesTable is the table that holds the outgoing_email_edges relation/edge.
	OutgoingEmailEdgesTable = "outgoing_emails"
	// OutgoingEmailEdgesInverseTable is the table name for the OutgoingEmail entity.
	// It exists in this package in order to avoid circular dependency with the "outgoingemail" package.
	OutgoingEmailEdgesInverseTable = "outgoing_emails"
	// OutgoingEmailEdgesColumn is the table column denoting the outgoing_email_edges relation/edge.
	OutgoingEmailEdgesColumn = "candidate_id"
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
	FieldReferenceType,
	FieldReferenceValue,
	FieldReferenceUID,
	FieldRecruitTime,
	FieldDescription,
	FieldAvatar,
	FieldCountry,
	FieldAddress,
	FieldGender,
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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
)

// ReferenceType defines the type for the "reference_type" enum field.
type ReferenceType string

// ReferenceTypeEb is the default value of the ReferenceType enum.
const DefaultReferenceType = ReferenceTypeEb

// ReferenceType values.
const (
	ReferenceTypeEb             ReferenceType = "eb"
	ReferenceTypeRec            ReferenceType = "rec"
	ReferenceTypeHiringPlatform ReferenceType = "hiring_platform"
	ReferenceTypeReference      ReferenceType = "reference"
	ReferenceTypeHeadhunt       ReferenceType = "headhunt"
)

func (rt ReferenceType) String() string {
	return string(rt)
}

// ReferenceTypeValidator is a validator for the "reference_type" field enum values. It is called by the builders before save.
func ReferenceTypeValidator(rt ReferenceType) error {
	switch rt {
	case ReferenceTypeEb, ReferenceTypeRec, ReferenceTypeHiringPlatform, ReferenceTypeReference, ReferenceTypeHeadhunt:
		return nil
	default:
		return fmt.Errorf("candidate: invalid enum value for reference_type field: %q", rt)
	}
}

// Gender defines the type for the "gender" enum field.
type Gender string

// GenderOthers is the default value of the Gender enum.
const DefaultGender = GenderOthers

// Gender values.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOthers Gender = "others"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale, GenderOthers:
		return nil
	default:
		return fmt.Errorf("candidate: invalid enum value for gender field: %q", ge)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e ReferenceType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *ReferenceType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = ReferenceType(str)
	if err := ReferenceTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid ReferenceType", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Gender) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Gender) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Gender(str)
	if err := GenderValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
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
	// FieldWorkEmail holds the string denoting the work_email field in the database.
	FieldWorkEmail = "work_email"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldOid holds the string denoting the oid field in the database.
	FieldOid = "oid"
	// FieldRecTeamID holds the string denoting the rec_team_id field in the database.
	FieldRecTeamID = "rec_team_id"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldHiringTeamID holds the string denoting the hiring_team_id field in the database.
	FieldHiringTeamID = "hiring_team_id"
	// EdgeAuditEdge holds the string denoting the audit_edge edge name in mutations.
	EdgeAuditEdge = "audit_edge"
	// EdgeHiringOwner holds the string denoting the hiring_owner edge name in mutations.
	EdgeHiringOwner = "hiring_owner"
	// EdgeCandidateJobFeedback holds the string denoting the candidate_job_feedback edge name in mutations.
	EdgeCandidateJobFeedback = "candidate_job_feedback"
	// EdgeInterviewEdges holds the string denoting the interview_edges edge name in mutations.
	EdgeInterviewEdges = "interview_edges"
	// EdgeCandidateJobEdges holds the string denoting the candidate_job_edges edge name in mutations.
	EdgeCandidateJobEdges = "candidate_job_edges"
	// EdgeCandidateInterviewEdges holds the string denoting the candidate_interview_edges edge name in mutations.
	EdgeCandidateInterviewEdges = "candidate_interview_edges"
	// EdgeCandidateReferenceEdges holds the string denoting the candidate_reference_edges edge name in mutations.
	EdgeCandidateReferenceEdges = "candidate_reference_edges"
	// EdgeUserPermissionEdges holds the string denoting the user_permission_edges edge name in mutations.
	EdgeUserPermissionEdges = "user_permission_edges"
	// EdgeRoleEdges holds the string denoting the role_edges edge name in mutations.
	EdgeRoleEdges = "role_edges"
	// EdgeHiringTeamEdges holds the string denoting the hiring_team_edges edge name in mutations.
	EdgeHiringTeamEdges = "hiring_team_edges"
	// EdgeMemberOfHiringTeamEdges holds the string denoting the member_of_hiring_team_edges edge name in mutations.
	EdgeMemberOfHiringTeamEdges = "member_of_hiring_team_edges"
	// EdgeApproversHiringTeams holds the string denoting the approvers_hiring_teams edge name in mutations.
	EdgeApproversHiringTeams = "approvers_hiring_teams"
	// EdgeLeaderRecEdge holds the string denoting the leader_rec_edge edge name in mutations.
	EdgeLeaderRecEdge = "leader_rec_edge"
	// EdgeRecTeams holds the string denoting the rec_teams edge name in mutations.
	EdgeRecTeams = "rec_teams"
	// EdgeCandidateNoteEdges holds the string denoting the candidate_note_edges edge name in mutations.
	EdgeCandidateNoteEdges = "candidate_note_edges"
	// EdgeCandidateHistoryCallEdges holds the string denoting the candidate_history_call_edges edge name in mutations.
	EdgeCandidateHistoryCallEdges = "candidate_history_call_edges"
	// EdgeApprovalJobs holds the string denoting the approval_jobs edge name in mutations.
	EdgeApprovalJobs = "approval_jobs"
	// EdgeHiringJobRecEdges holds the string denoting the hiring_job_rec_edges edge name in mutations.
	EdgeHiringJobRecEdges = "hiring_job_rec_edges"
	// EdgeInterviewUsers holds the string denoting the interview_users edge name in mutations.
	EdgeInterviewUsers = "interview_users"
	// EdgeRoleUsers holds the string denoting the role_users edge name in mutations.
	EdgeRoleUsers = "role_users"
	// EdgeHiringTeamUsers holds the string denoting the hiring_team_users edge name in mutations.
	EdgeHiringTeamUsers = "hiring_team_users"
	// EdgeHiringTeamApprovers holds the string denoting the hiring_team_approvers edge name in mutations.
	EdgeHiringTeamApprovers = "hiring_team_approvers"
	// EdgeApprovalSteps holds the string denoting the approval_steps edge name in mutations.
	EdgeApprovalSteps = "approval_steps"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AuditEdgeTable is the table that holds the audit_edge relation/edge.
	AuditEdgeTable = "audit_trails"
	// AuditEdgeInverseTable is the table name for the AuditTrail entity.
	// It exists in this package in order to avoid circular dependency with the "audittrail" package.
	AuditEdgeInverseTable = "audit_trails"
	// AuditEdgeColumn is the table column denoting the audit_edge relation/edge.
	AuditEdgeColumn = "created_by"
	// HiringOwnerTable is the table that holds the hiring_owner relation/edge.
	HiringOwnerTable = "hiring_jobs"
	// HiringOwnerInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringOwnerInverseTable = "hiring_jobs"
	// HiringOwnerColumn is the table column denoting the hiring_owner relation/edge.
	HiringOwnerColumn = "created_by"
	// CandidateJobFeedbackTable is the table that holds the candidate_job_feedback relation/edge.
	CandidateJobFeedbackTable = "candidate_job_feedbacks"
	// CandidateJobFeedbackInverseTable is the table name for the CandidateJobFeedback entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejobfeedback" package.
	CandidateJobFeedbackInverseTable = "candidate_job_feedbacks"
	// CandidateJobFeedbackColumn is the table column denoting the candidate_job_feedback relation/edge.
	CandidateJobFeedbackColumn = "created_by"
	// InterviewEdgesTable is the table that holds the interview_edges relation/edge. The primary key declared below.
	InterviewEdgesTable = "candidate_interviewers"
	// InterviewEdgesInverseTable is the table name for the CandidateInterview entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterview" package.
	InterviewEdgesInverseTable = "candidate_interviews"
	// CandidateJobEdgesTable is the table that holds the candidate_job_edges relation/edge.
	CandidateJobEdgesTable = "candidate_jobs"
	// CandidateJobEdgesInverseTable is the table name for the CandidateJob entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejob" package.
	CandidateJobEdgesInverseTable = "candidate_jobs"
	// CandidateJobEdgesColumn is the table column denoting the candidate_job_edges relation/edge.
	CandidateJobEdgesColumn = "created_by"
	// CandidateInterviewEdgesTable is the table that holds the candidate_interview_edges relation/edge.
	CandidateInterviewEdgesTable = "candidate_interviews"
	// CandidateInterviewEdgesInverseTable is the table name for the CandidateInterview entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterview" package.
	CandidateInterviewEdgesInverseTable = "candidate_interviews"
	// CandidateInterviewEdgesColumn is the table column denoting the candidate_interview_edges relation/edge.
	CandidateInterviewEdgesColumn = "created_by"
	// CandidateReferenceEdgesTable is the table that holds the candidate_reference_edges relation/edge.
	CandidateReferenceEdgesTable = "candidates"
	// CandidateReferenceEdgesInverseTable is the table name for the Candidate entity.
	// It exists in this package in order to avoid circular dependency with the "candidate" package.
	CandidateReferenceEdgesInverseTable = "candidates"
	// CandidateReferenceEdgesColumn is the table column denoting the candidate_reference_edges relation/edge.
	CandidateReferenceEdgesColumn = "reference_uid"
	// UserPermissionEdgesTable is the table that holds the user_permission_edges relation/edge.
	UserPermissionEdgesTable = "entity_permissions"
	// UserPermissionEdgesInverseTable is the table name for the EntityPermission entity.
	// It exists in this package in order to avoid circular dependency with the "entitypermission" package.
	UserPermissionEdgesInverseTable = "entity_permissions"
	// UserPermissionEdgesColumn is the table column denoting the user_permission_edges relation/edge.
	UserPermissionEdgesColumn = "entity_id"
	// RoleEdgesTable is the table that holds the role_edges relation/edge. The primary key declared below.
	RoleEdgesTable = "user_roles"
	// RoleEdgesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleEdgesInverseTable = "roles"
	// HiringTeamEdgesTable is the table that holds the hiring_team_edges relation/edge. The primary key declared below.
	HiringTeamEdgesTable = "hiring_team_managers"
	// HiringTeamEdgesInverseTable is the table name for the HiringTeam entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteam" package.
	HiringTeamEdgesInverseTable = "hiring_teams"
	// MemberOfHiringTeamEdgesTable is the table that holds the member_of_hiring_team_edges relation/edge.
	MemberOfHiringTeamEdgesTable = "users"
	// MemberOfHiringTeamEdgesInverseTable is the table name for the HiringTeam entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteam" package.
	MemberOfHiringTeamEdgesInverseTable = "hiring_teams"
	// MemberOfHiringTeamEdgesColumn is the table column denoting the member_of_hiring_team_edges relation/edge.
	MemberOfHiringTeamEdgesColumn = "hiring_team_id"
	// ApproversHiringTeamsTable is the table that holds the approvers_hiring_teams relation/edge. The primary key declared below.
	ApproversHiringTeamsTable = "hiring_team_approvers"
	// ApproversHiringTeamsInverseTable is the table name for the HiringTeam entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteam" package.
	ApproversHiringTeamsInverseTable = "hiring_teams"
	// LeaderRecEdgeTable is the table that holds the leader_rec_edge relation/edge.
	LeaderRecEdgeTable = "rec_teams"
	// LeaderRecEdgeInverseTable is the table name for the RecTeam entity.
	// It exists in this package in order to avoid circular dependency with the "recteam" package.
	LeaderRecEdgeInverseTable = "rec_teams"
	// LeaderRecEdgeColumn is the table column denoting the leader_rec_edge relation/edge.
	LeaderRecEdgeColumn = "leader_id"
	// RecTeamsTable is the table that holds the rec_teams relation/edge.
	RecTeamsTable = "users"
	// RecTeamsInverseTable is the table name for the RecTeam entity.
	// It exists in this package in order to avoid circular dependency with the "recteam" package.
	RecTeamsInverseTable = "rec_teams"
	// RecTeamsColumn is the table column denoting the rec_teams relation/edge.
	RecTeamsColumn = "rec_team_id"
	// CandidateNoteEdgesTable is the table that holds the candidate_note_edges relation/edge.
	CandidateNoteEdgesTable = "candidate_notes"
	// CandidateNoteEdgesInverseTable is the table name for the CandidateNote entity.
	// It exists in this package in order to avoid circular dependency with the "candidatenote" package.
	CandidateNoteEdgesInverseTable = "candidate_notes"
	// CandidateNoteEdgesColumn is the table column denoting the candidate_note_edges relation/edge.
	CandidateNoteEdgesColumn = "created_by_id"
	// CandidateHistoryCallEdgesTable is the table that holds the candidate_history_call_edges relation/edge.
	CandidateHistoryCallEdgesTable = "candidate_history_calls"
	// CandidateHistoryCallEdgesInverseTable is the table name for the CandidateHistoryCall entity.
	// It exists in this package in order to avoid circular dependency with the "candidatehistorycall" package.
	CandidateHistoryCallEdgesInverseTable = "candidate_history_calls"
	// CandidateHistoryCallEdgesColumn is the table column denoting the candidate_history_call_edges relation/edge.
	CandidateHistoryCallEdgesColumn = "created_by_id"
	// ApprovalJobsTable is the table that holds the approval_jobs relation/edge. The primary key declared below.
	ApprovalJobsTable = "hiring_job_steps"
	// ApprovalJobsInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	ApprovalJobsInverseTable = "hiring_jobs"
	// HiringJobRecEdgesTable is the table that holds the hiring_job_rec_edges relation/edge.
	HiringJobRecEdgesTable = "hiring_jobs"
	// HiringJobRecEdgesInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringJobRecEdgesInverseTable = "hiring_jobs"
	// HiringJobRecEdgesColumn is the table column denoting the hiring_job_rec_edges relation/edge.
	HiringJobRecEdgesColumn = "rec_in_charge_id"
	// InterviewUsersTable is the table that holds the interview_users relation/edge.
	InterviewUsersTable = "candidate_interviewers"
	// InterviewUsersInverseTable is the table name for the CandidateInterviewer entity.
	// It exists in this package in order to avoid circular dependency with the "candidateinterviewer" package.
	InterviewUsersInverseTable = "candidate_interviewers"
	// InterviewUsersColumn is the table column denoting the interview_users relation/edge.
	InterviewUsersColumn = "user_id"
	// RoleUsersTable is the table that holds the role_users relation/edge.
	RoleUsersTable = "user_roles"
	// RoleUsersInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	RoleUsersInverseTable = "user_roles"
	// RoleUsersColumn is the table column denoting the role_users relation/edge.
	RoleUsersColumn = "user_id"
	// HiringTeamUsersTable is the table that holds the hiring_team_users relation/edge.
	HiringTeamUsersTable = "hiring_team_managers"
	// HiringTeamUsersInverseTable is the table name for the HiringTeamManager entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteammanager" package.
	HiringTeamUsersInverseTable = "hiring_team_managers"
	// HiringTeamUsersColumn is the table column denoting the hiring_team_users relation/edge.
	HiringTeamUsersColumn = "user_id"
	// HiringTeamApproversTable is the table that holds the hiring_team_approvers relation/edge.
	HiringTeamApproversTable = "hiring_team_approvers"
	// HiringTeamApproversInverseTable is the table name for the HiringTeamApprover entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteamapprover" package.
	HiringTeamApproversInverseTable = "hiring_team_approvers"
	// HiringTeamApproversColumn is the table column denoting the hiring_team_approvers relation/edge.
	HiringTeamApproversColumn = "user_id"
	// ApprovalStepsTable is the table that holds the approval_steps relation/edge.
	ApprovalStepsTable = "hiring_job_steps"
	// ApprovalStepsInverseTable is the table name for the HiringJobStep entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjobstep" package.
	ApprovalStepsInverseTable = "hiring_job_steps"
	// ApprovalStepsColumn is the table column denoting the approval_steps relation/edge.
	ApprovalStepsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldWorkEmail,
	FieldStatus,
	FieldOid,
	FieldRecTeamID,
	FieldLocation,
	FieldHiringTeamID,
}

var (
	// InterviewEdgesPrimaryKey and InterviewEdgesColumn2 are the table columns denoting the
	// primary key for the interview_edges relation (M2M).
	InterviewEdgesPrimaryKey = []string{"user_id", "candidate_interview_id"}
	// RoleEdgesPrimaryKey and RoleEdgesColumn2 are the table columns denoting the
	// primary key for the role_edges relation (M2M).
	RoleEdgesPrimaryKey = []string{"user_id", "role_id"}
	// HiringTeamEdgesPrimaryKey and HiringTeamEdgesColumn2 are the table columns denoting the
	// primary key for the hiring_team_edges relation (M2M).
	HiringTeamEdgesPrimaryKey = []string{"user_id", "hiring_team_id"}
	// ApproversHiringTeamsPrimaryKey and ApproversHiringTeamsColumn2 are the table columns denoting the
	// primary key for the approvers_hiring_teams relation (M2M).
	ApproversHiringTeamsPrimaryKey = []string{"hiring_team_id", "user_id"}
	// ApprovalJobsPrimaryKey and ApprovalJobsColumn2 are the table columns denoting the
	// primary key for the approval_jobs relation (M2M).
	ApprovalJobsPrimaryKey = []string{"user_id", "hiring_job_id"}
)

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
	// WorkEmailValidator is a validator for the "work_email" field. It is called by the builders before save.
	WorkEmailValidator func(string) error
	// OidValidator is a validator for the "oid" field. It is called by the builders before save.
	OidValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusInactive:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
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

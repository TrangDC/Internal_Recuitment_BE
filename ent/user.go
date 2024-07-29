// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/hiringteam"
	"trec/ent/recteam"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// WorkEmail holds the value of the "work_email" field.
	WorkEmail string `json:"work_email,omitempty"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// Oid holds the value of the "oid" field.
	Oid string `json:"oid,omitempty"`
	// RecTeamID holds the value of the "rec_team_id" field.
	RecTeamID uuid.UUID `json:"rec_team_id,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// HiringTeamID holds the value of the "hiring_team_id" field.
	HiringTeamID uuid.UUID `json:"hiring_team_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// AuditEdge holds the value of the audit_edge edge.
	AuditEdge []*AuditTrail `json:"audit_edge,omitempty"`
	// HiringOwner holds the value of the hiring_owner edge.
	HiringOwner []*HiringJob `json:"hiring_owner,omitempty"`
	// CandidateJobFeedback holds the value of the candidate_job_feedback edge.
	CandidateJobFeedback []*CandidateJobFeedback `json:"candidate_job_feedback,omitempty"`
	// InterviewEdges holds the value of the interview_edges edge.
	InterviewEdges []*CandidateInterview `json:"interview_edges,omitempty"`
	// CandidateJobEdges holds the value of the candidate_job_edges edge.
	CandidateJobEdges []*CandidateJob `json:"candidate_job_edges,omitempty"`
	// CandidateInterviewEdges holds the value of the candidate_interview_edges edge.
	CandidateInterviewEdges []*CandidateInterview `json:"candidate_interview_edges,omitempty"`
	// CandidateReferenceEdges holds the value of the candidate_reference_edges edge.
	CandidateReferenceEdges []*Candidate `json:"candidate_reference_edges,omitempty"`
	// UserPermissionEdges holds the value of the user_permission_edges edge.
	UserPermissionEdges []*EntityPermission `json:"user_permission_edges,omitempty"`
	// RoleEdges holds the value of the role_edges edge.
	RoleEdges []*Role `json:"role_edges,omitempty"`
	// HiringTeamEdges holds the value of the hiring_team_edges edge.
	HiringTeamEdges []*HiringTeam `json:"hiring_team_edges,omitempty"`
	// LedRecTeams holds the value of the led_rec_teams edge.
	LedRecTeams []*RecTeam `json:"led_rec_teams,omitempty"`
	// RecTeams holds the value of the rec_teams edge.
	RecTeams *RecTeam `json:"rec_teams,omitempty"`
	// MemberOfHiringTeamEdges holds the value of the member_of_hiring_team_edges edge.
	MemberOfHiringTeamEdges *HiringTeam `json:"member_of_hiring_team_edges,omitempty"`
	// InterviewUsers holds the value of the interview_users edge.
	InterviewUsers []*CandidateInterviewer `json:"interview_users,omitempty"`
	// RoleUsers holds the value of the role_users edge.
	RoleUsers []*UserRole `json:"role_users,omitempty"`
	// HiringTeamUsers holds the value of the hiring_team_users edge.
	HiringTeamUsers []*HiringTeamManager `json:"hiring_team_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [16]bool
	// totalCount holds the count of the edges above.
	totalCount [16]map[string]int

	namedAuditEdge               map[string][]*AuditTrail
	namedHiringOwner             map[string][]*HiringJob
	namedCandidateJobFeedback    map[string][]*CandidateJobFeedback
	namedInterviewEdges          map[string][]*CandidateInterview
	namedCandidateJobEdges       map[string][]*CandidateJob
	namedCandidateInterviewEdges map[string][]*CandidateInterview
	namedCandidateReferenceEdges map[string][]*Candidate
	namedUserPermissionEdges     map[string][]*EntityPermission
	namedRoleEdges               map[string][]*Role
	namedHiringTeamEdges         map[string][]*HiringTeam
	namedLedRecTeams             map[string][]*RecTeam
	namedInterviewUsers          map[string][]*CandidateInterviewer
	namedRoleUsers               map[string][]*UserRole
	namedHiringTeamUsers         map[string][]*HiringTeamManager
}

// AuditEdgeOrErr returns the AuditEdge value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AuditEdgeOrErr() ([]*AuditTrail, error) {
	if e.loadedTypes[0] {
		return e.AuditEdge, nil
	}
	return nil, &NotLoadedError{edge: "audit_edge"}
}

// HiringOwnerOrErr returns the HiringOwner value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HiringOwnerOrErr() ([]*HiringJob, error) {
	if e.loadedTypes[1] {
		return e.HiringOwner, nil
	}
	return nil, &NotLoadedError{edge: "hiring_owner"}
}

// CandidateJobFeedbackOrErr returns the CandidateJobFeedback value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CandidateJobFeedbackOrErr() ([]*CandidateJobFeedback, error) {
	if e.loadedTypes[2] {
		return e.CandidateJobFeedback, nil
	}
	return nil, &NotLoadedError{edge: "candidate_job_feedback"}
}

// InterviewEdgesOrErr returns the InterviewEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InterviewEdgesOrErr() ([]*CandidateInterview, error) {
	if e.loadedTypes[3] {
		return e.InterviewEdges, nil
	}
	return nil, &NotLoadedError{edge: "interview_edges"}
}

// CandidateJobEdgesOrErr returns the CandidateJobEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CandidateJobEdgesOrErr() ([]*CandidateJob, error) {
	if e.loadedTypes[4] {
		return e.CandidateJobEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_job_edges"}
}

// CandidateInterviewEdgesOrErr returns the CandidateInterviewEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CandidateInterviewEdgesOrErr() ([]*CandidateInterview, error) {
	if e.loadedTypes[5] {
		return e.CandidateInterviewEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_interview_edges"}
}

// CandidateReferenceEdgesOrErr returns the CandidateReferenceEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CandidateReferenceEdgesOrErr() ([]*Candidate, error) {
	if e.loadedTypes[6] {
		return e.CandidateReferenceEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_reference_edges"}
}

// UserPermissionEdgesOrErr returns the UserPermissionEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserPermissionEdgesOrErr() ([]*EntityPermission, error) {
	if e.loadedTypes[7] {
		return e.UserPermissionEdges, nil
	}
	return nil, &NotLoadedError{edge: "user_permission_edges"}
}

// RoleEdgesOrErr returns the RoleEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoleEdgesOrErr() ([]*Role, error) {
	if e.loadedTypes[8] {
		return e.RoleEdges, nil
	}
	return nil, &NotLoadedError{edge: "role_edges"}
}

// HiringTeamEdgesOrErr returns the HiringTeamEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HiringTeamEdgesOrErr() ([]*HiringTeam, error) {
	if e.loadedTypes[9] {
		return e.HiringTeamEdges, nil
	}
	return nil, &NotLoadedError{edge: "hiring_team_edges"}
}

// LedRecTeamsOrErr returns the LedRecTeams value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LedRecTeamsOrErr() ([]*RecTeam, error) {
	if e.loadedTypes[10] {
		return e.LedRecTeams, nil
	}
	return nil, &NotLoadedError{edge: "led_rec_teams"}
}

// RecTeamsOrErr returns the RecTeams value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RecTeamsOrErr() (*RecTeam, error) {
	if e.loadedTypes[11] {
		if e.RecTeams == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: recteam.Label}
		}
		return e.RecTeams, nil
	}
	return nil, &NotLoadedError{edge: "rec_teams"}
}

// MemberOfHiringTeamEdgesOrErr returns the MemberOfHiringTeamEdges value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) MemberOfHiringTeamEdgesOrErr() (*HiringTeam, error) {
	if e.loadedTypes[12] {
		if e.MemberOfHiringTeamEdges == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hiringteam.Label}
		}
		return e.MemberOfHiringTeamEdges, nil
	}
	return nil, &NotLoadedError{edge: "member_of_hiring_team_edges"}
}

// InterviewUsersOrErr returns the InterviewUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) InterviewUsersOrErr() ([]*CandidateInterviewer, error) {
	if e.loadedTypes[13] {
		return e.InterviewUsers, nil
	}
	return nil, &NotLoadedError{edge: "interview_users"}
}

// RoleUsersOrErr returns the RoleUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoleUsersOrErr() ([]*UserRole, error) {
	if e.loadedTypes[14] {
		return e.RoleUsers, nil
	}
	return nil, &NotLoadedError{edge: "role_users"}
}

// HiringTeamUsersOrErr returns the HiringTeamUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) HiringTeamUsersOrErr() ([]*HiringTeamManager, error) {
	if e.loadedTypes[15] {
		return e.HiringTeamUsers, nil
	}
	return nil, &NotLoadedError{edge: "hiring_team_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldName, user.FieldWorkEmail, user.FieldStatus, user.FieldOid, user.FieldLocation:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID, user.FieldRecTeamID, user.FieldHiringTeamID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldWorkEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field work_email", values[i])
			} else if value.Valid {
				u.WorkEmail = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldOid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				u.Oid = value.String
			}
		case user.FieldRecTeamID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field rec_team_id", values[i])
			} else if value != nil {
				u.RecTeamID = *value
			}
		case user.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				u.Location = value.String
			}
		case user.FieldHiringTeamID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field hiring_team_id", values[i])
			} else if value != nil {
				u.HiringTeamID = *value
			}
		}
	}
	return nil
}

// QueryAuditEdge queries the "audit_edge" edge of the User entity.
func (u *User) QueryAuditEdge() *AuditTrailQuery {
	return (&UserClient{config: u.config}).QueryAuditEdge(u)
}

// QueryHiringOwner queries the "hiring_owner" edge of the User entity.
func (u *User) QueryHiringOwner() *HiringJobQuery {
	return (&UserClient{config: u.config}).QueryHiringOwner(u)
}

// QueryCandidateJobFeedback queries the "candidate_job_feedback" edge of the User entity.
func (u *User) QueryCandidateJobFeedback() *CandidateJobFeedbackQuery {
	return (&UserClient{config: u.config}).QueryCandidateJobFeedback(u)
}

// QueryInterviewEdges queries the "interview_edges" edge of the User entity.
func (u *User) QueryInterviewEdges() *CandidateInterviewQuery {
	return (&UserClient{config: u.config}).QueryInterviewEdges(u)
}

// QueryCandidateJobEdges queries the "candidate_job_edges" edge of the User entity.
func (u *User) QueryCandidateJobEdges() *CandidateJobQuery {
	return (&UserClient{config: u.config}).QueryCandidateJobEdges(u)
}

// QueryCandidateInterviewEdges queries the "candidate_interview_edges" edge of the User entity.
func (u *User) QueryCandidateInterviewEdges() *CandidateInterviewQuery {
	return (&UserClient{config: u.config}).QueryCandidateInterviewEdges(u)
}

// QueryCandidateReferenceEdges queries the "candidate_reference_edges" edge of the User entity.
func (u *User) QueryCandidateReferenceEdges() *CandidateQuery {
	return (&UserClient{config: u.config}).QueryCandidateReferenceEdges(u)
}

// QueryUserPermissionEdges queries the "user_permission_edges" edge of the User entity.
func (u *User) QueryUserPermissionEdges() *EntityPermissionQuery {
	return (&UserClient{config: u.config}).QueryUserPermissionEdges(u)
}

// QueryRoleEdges queries the "role_edges" edge of the User entity.
func (u *User) QueryRoleEdges() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRoleEdges(u)
}

// QueryHiringTeamEdges queries the "hiring_team_edges" edge of the User entity.
func (u *User) QueryHiringTeamEdges() *HiringTeamQuery {
	return (&UserClient{config: u.config}).QueryHiringTeamEdges(u)
}

// QueryLedRecTeams queries the "led_rec_teams" edge of the User entity.
func (u *User) QueryLedRecTeams() *RecTeamQuery {
	return (&UserClient{config: u.config}).QueryLedRecTeams(u)
}

// QueryRecTeams queries the "rec_teams" edge of the User entity.
func (u *User) QueryRecTeams() *RecTeamQuery {
	return (&UserClient{config: u.config}).QueryRecTeams(u)
}

// QueryMemberOfHiringTeamEdges queries the "member_of_hiring_team_edges" edge of the User entity.
func (u *User) QueryMemberOfHiringTeamEdges() *HiringTeamQuery {
	return (&UserClient{config: u.config}).QueryMemberOfHiringTeamEdges(u)
}

// QueryInterviewUsers queries the "interview_users" edge of the User entity.
func (u *User) QueryInterviewUsers() *CandidateInterviewerQuery {
	return (&UserClient{config: u.config}).QueryInterviewUsers(u)
}

// QueryRoleUsers queries the "role_users" edge of the User entity.
func (u *User) QueryRoleUsers() *UserRoleQuery {
	return (&UserClient{config: u.config}).QueryRoleUsers(u)
}

// QueryHiringTeamUsers queries the "hiring_team_users" edge of the User entity.
func (u *User) QueryHiringTeamUsers() *HiringTeamManagerQuery {
	return (&UserClient{config: u.config}).QueryHiringTeamUsers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("work_email=")
	builder.WriteString(u.WorkEmail)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("oid=")
	builder.WriteString(u.Oid)
	builder.WriteString(", ")
	builder.WriteString("rec_team_id=")
	builder.WriteString(fmt.Sprintf("%v", u.RecTeamID))
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(u.Location)
	builder.WriteString(", ")
	builder.WriteString("hiring_team_id=")
	builder.WriteString(fmt.Sprintf("%v", u.HiringTeamID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAuditEdge returns the AuditEdge named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAuditEdge(name string) ([]*AuditTrail, error) {
	if u.Edges.namedAuditEdge == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAuditEdge[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAuditEdge(name string, edges ...*AuditTrail) {
	if u.Edges.namedAuditEdge == nil {
		u.Edges.namedAuditEdge = make(map[string][]*AuditTrail)
	}
	if len(edges) == 0 {
		u.Edges.namedAuditEdge[name] = []*AuditTrail{}
	} else {
		u.Edges.namedAuditEdge[name] = append(u.Edges.namedAuditEdge[name], edges...)
	}
}

// NamedHiringOwner returns the HiringOwner named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedHiringOwner(name string) ([]*HiringJob, error) {
	if u.Edges.namedHiringOwner == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedHiringOwner[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedHiringOwner(name string, edges ...*HiringJob) {
	if u.Edges.namedHiringOwner == nil {
		u.Edges.namedHiringOwner = make(map[string][]*HiringJob)
	}
	if len(edges) == 0 {
		u.Edges.namedHiringOwner[name] = []*HiringJob{}
	} else {
		u.Edges.namedHiringOwner[name] = append(u.Edges.namedHiringOwner[name], edges...)
	}
}

// NamedCandidateJobFeedback returns the CandidateJobFeedback named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCandidateJobFeedback(name string) ([]*CandidateJobFeedback, error) {
	if u.Edges.namedCandidateJobFeedback == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCandidateJobFeedback[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCandidateJobFeedback(name string, edges ...*CandidateJobFeedback) {
	if u.Edges.namedCandidateJobFeedback == nil {
		u.Edges.namedCandidateJobFeedback = make(map[string][]*CandidateJobFeedback)
	}
	if len(edges) == 0 {
		u.Edges.namedCandidateJobFeedback[name] = []*CandidateJobFeedback{}
	} else {
		u.Edges.namedCandidateJobFeedback[name] = append(u.Edges.namedCandidateJobFeedback[name], edges...)
	}
}

// NamedInterviewEdges returns the InterviewEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedInterviewEdges(name string) ([]*CandidateInterview, error) {
	if u.Edges.namedInterviewEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedInterviewEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedInterviewEdges(name string, edges ...*CandidateInterview) {
	if u.Edges.namedInterviewEdges == nil {
		u.Edges.namedInterviewEdges = make(map[string][]*CandidateInterview)
	}
	if len(edges) == 0 {
		u.Edges.namedInterviewEdges[name] = []*CandidateInterview{}
	} else {
		u.Edges.namedInterviewEdges[name] = append(u.Edges.namedInterviewEdges[name], edges...)
	}
}

// NamedCandidateJobEdges returns the CandidateJobEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCandidateJobEdges(name string) ([]*CandidateJob, error) {
	if u.Edges.namedCandidateJobEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCandidateJobEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCandidateJobEdges(name string, edges ...*CandidateJob) {
	if u.Edges.namedCandidateJobEdges == nil {
		u.Edges.namedCandidateJobEdges = make(map[string][]*CandidateJob)
	}
	if len(edges) == 0 {
		u.Edges.namedCandidateJobEdges[name] = []*CandidateJob{}
	} else {
		u.Edges.namedCandidateJobEdges[name] = append(u.Edges.namedCandidateJobEdges[name], edges...)
	}
}

// NamedCandidateInterviewEdges returns the CandidateInterviewEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCandidateInterviewEdges(name string) ([]*CandidateInterview, error) {
	if u.Edges.namedCandidateInterviewEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCandidateInterviewEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCandidateInterviewEdges(name string, edges ...*CandidateInterview) {
	if u.Edges.namedCandidateInterviewEdges == nil {
		u.Edges.namedCandidateInterviewEdges = make(map[string][]*CandidateInterview)
	}
	if len(edges) == 0 {
		u.Edges.namedCandidateInterviewEdges[name] = []*CandidateInterview{}
	} else {
		u.Edges.namedCandidateInterviewEdges[name] = append(u.Edges.namedCandidateInterviewEdges[name], edges...)
	}
}

// NamedCandidateReferenceEdges returns the CandidateReferenceEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedCandidateReferenceEdges(name string) ([]*Candidate, error) {
	if u.Edges.namedCandidateReferenceEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedCandidateReferenceEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedCandidateReferenceEdges(name string, edges ...*Candidate) {
	if u.Edges.namedCandidateReferenceEdges == nil {
		u.Edges.namedCandidateReferenceEdges = make(map[string][]*Candidate)
	}
	if len(edges) == 0 {
		u.Edges.namedCandidateReferenceEdges[name] = []*Candidate{}
	} else {
		u.Edges.namedCandidateReferenceEdges[name] = append(u.Edges.namedCandidateReferenceEdges[name], edges...)
	}
}

// NamedUserPermissionEdges returns the UserPermissionEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedUserPermissionEdges(name string) ([]*EntityPermission, error) {
	if u.Edges.namedUserPermissionEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedUserPermissionEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedUserPermissionEdges(name string, edges ...*EntityPermission) {
	if u.Edges.namedUserPermissionEdges == nil {
		u.Edges.namedUserPermissionEdges = make(map[string][]*EntityPermission)
	}
	if len(edges) == 0 {
		u.Edges.namedUserPermissionEdges[name] = []*EntityPermission{}
	} else {
		u.Edges.namedUserPermissionEdges[name] = append(u.Edges.namedUserPermissionEdges[name], edges...)
	}
}

// NamedRoleEdges returns the RoleEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRoleEdges(name string) ([]*Role, error) {
	if u.Edges.namedRoleEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRoleEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRoleEdges(name string, edges ...*Role) {
	if u.Edges.namedRoleEdges == nil {
		u.Edges.namedRoleEdges = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		u.Edges.namedRoleEdges[name] = []*Role{}
	} else {
		u.Edges.namedRoleEdges[name] = append(u.Edges.namedRoleEdges[name], edges...)
	}
}

// NamedHiringTeamEdges returns the HiringTeamEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedHiringTeamEdges(name string) ([]*HiringTeam, error) {
	if u.Edges.namedHiringTeamEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedHiringTeamEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedHiringTeamEdges(name string, edges ...*HiringTeam) {
	if u.Edges.namedHiringTeamEdges == nil {
		u.Edges.namedHiringTeamEdges = make(map[string][]*HiringTeam)
	}
	if len(edges) == 0 {
		u.Edges.namedHiringTeamEdges[name] = []*HiringTeam{}
	} else {
		u.Edges.namedHiringTeamEdges[name] = append(u.Edges.namedHiringTeamEdges[name], edges...)
	}
}

// NamedLedRecTeams returns the LedRecTeams named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedLedRecTeams(name string) ([]*RecTeam, error) {
	if u.Edges.namedLedRecTeams == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedLedRecTeams[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedLedRecTeams(name string, edges ...*RecTeam) {
	if u.Edges.namedLedRecTeams == nil {
		u.Edges.namedLedRecTeams = make(map[string][]*RecTeam)
	}
	if len(edges) == 0 {
		u.Edges.namedLedRecTeams[name] = []*RecTeam{}
	} else {
		u.Edges.namedLedRecTeams[name] = append(u.Edges.namedLedRecTeams[name], edges...)
	}
}

// NamedInterviewUsers returns the InterviewUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedInterviewUsers(name string) ([]*CandidateInterviewer, error) {
	if u.Edges.namedInterviewUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedInterviewUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedInterviewUsers(name string, edges ...*CandidateInterviewer) {
	if u.Edges.namedInterviewUsers == nil {
		u.Edges.namedInterviewUsers = make(map[string][]*CandidateInterviewer)
	}
	if len(edges) == 0 {
		u.Edges.namedInterviewUsers[name] = []*CandidateInterviewer{}
	} else {
		u.Edges.namedInterviewUsers[name] = append(u.Edges.namedInterviewUsers[name], edges...)
	}
}

// NamedRoleUsers returns the RoleUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRoleUsers(name string) ([]*UserRole, error) {
	if u.Edges.namedRoleUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRoleUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRoleUsers(name string, edges ...*UserRole) {
	if u.Edges.namedRoleUsers == nil {
		u.Edges.namedRoleUsers = make(map[string][]*UserRole)
	}
	if len(edges) == 0 {
		u.Edges.namedRoleUsers[name] = []*UserRole{}
	} else {
		u.Edges.namedRoleUsers[name] = append(u.Edges.namedRoleUsers[name], edges...)
	}
}

// NamedHiringTeamUsers returns the HiringTeamUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedHiringTeamUsers(name string) ([]*HiringTeamManager, error) {
	if u.Edges.namedHiringTeamUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedHiringTeamUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedHiringTeamUsers(name string, edges ...*HiringTeamManager) {
	if u.Edges.namedHiringTeamUsers == nil {
		u.Edges.namedHiringTeamUsers = make(map[string][]*HiringTeamManager)
	}
	if len(edges) == 0 {
		u.Edges.namedHiringTeamUsers[name] = []*HiringTeamManager{}
	} else {
		u.Edges.namedHiringTeamUsers[name] = append(u.Edges.namedHiringTeamUsers[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}

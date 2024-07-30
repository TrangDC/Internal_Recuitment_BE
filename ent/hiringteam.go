// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/hiringteam"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// HiringTeam is the model entity for the HiringTeam schema.
type HiringTeam struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HiringTeamQuery when eager-loading is set.
	Edges HiringTeamEdges `json:"edges"`
}

// HiringTeamEdges holds the relations/edges for other nodes in the graph.
type HiringTeamEdges struct {
	// The uniqueness of the user is enforced on the edge schema
	UserEdges []*User `json:"user_edges,omitempty"`
	// HiringTeamJobEdges holds the value of the hiring_team_job_edges edge.
	HiringTeamJobEdges []*HiringJob `json:"hiring_team_job_edges,omitempty"`
	// HiringMemberEdges holds the value of the hiring_member_edges edge.
	HiringMemberEdges []*User `json:"hiring_member_edges,omitempty"`
	// ApproversUsers holds the value of the approvers_users edge.
	ApproversUsers []*User `json:"approvers_users,omitempty"`
	// UserHiringTeams holds the value of the user_hiring_teams edge.
	UserHiringTeams []*HiringTeamManager `json:"user_hiring_teams,omitempty"`
	// HiringTeamApprovers holds the value of the hiring_team_approvers edge.
	HiringTeamApprovers []*HiringTeamApprover `json:"hiring_team_approvers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedUserEdges           map[string][]*User
	namedHiringTeamJobEdges  map[string][]*HiringJob
	namedHiringMemberEdges   map[string][]*User
	namedApproversUsers      map[string][]*User
	namedUserHiringTeams     map[string][]*HiringTeamManager
	namedHiringTeamApprovers map[string][]*HiringTeamApprover
}

// UserEdgesOrErr returns the UserEdges value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) UserEdgesOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.UserEdges, nil
	}
	return nil, &NotLoadedError{edge: "user_edges"}
}

// HiringTeamJobEdgesOrErr returns the HiringTeamJobEdges value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) HiringTeamJobEdgesOrErr() ([]*HiringJob, error) {
	if e.loadedTypes[1] {
		return e.HiringTeamJobEdges, nil
	}
	return nil, &NotLoadedError{edge: "hiring_team_job_edges"}
}

// HiringMemberEdgesOrErr returns the HiringMemberEdges value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) HiringMemberEdgesOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.HiringMemberEdges, nil
	}
	return nil, &NotLoadedError{edge: "hiring_member_edges"}
}

// ApproversUsersOrErr returns the ApproversUsers value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) ApproversUsersOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.ApproversUsers, nil
	}
	return nil, &NotLoadedError{edge: "approvers_users"}
}

// UserHiringTeamsOrErr returns the UserHiringTeams value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) UserHiringTeamsOrErr() ([]*HiringTeamManager, error) {
	if e.loadedTypes[4] {
		return e.UserHiringTeams, nil
	}
	return nil, &NotLoadedError{edge: "user_hiring_teams"}
}

// HiringTeamApproversOrErr returns the HiringTeamApprovers value or an error if the edge
// was not loaded in eager-loading.
func (e HiringTeamEdges) HiringTeamApproversOrErr() ([]*HiringTeamApprover, error) {
	if e.loadedTypes[5] {
		return e.HiringTeamApprovers, nil
	}
	return nil, &NotLoadedError{edge: "hiring_team_approvers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HiringTeam) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hiringteam.FieldSlug, hiringteam.FieldName, hiringteam.FieldDescription:
			values[i] = new(sql.NullString)
		case hiringteam.FieldCreatedAt, hiringteam.FieldUpdatedAt, hiringteam.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case hiringteam.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type HiringTeam", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HiringTeam fields.
func (ht *HiringTeam) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hiringteam.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ht.ID = *value
			}
		case hiringteam.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ht.CreatedAt = value.Time
			}
		case hiringteam.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ht.UpdatedAt = value.Time
			}
		case hiringteam.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ht.DeletedAt = value.Time
			}
		case hiringteam.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				ht.Slug = value.String
			}
		case hiringteam.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ht.Name = value.String
			}
		case hiringteam.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ht.Description = value.String
			}
		}
	}
	return nil
}

// QueryUserEdges queries the "user_edges" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryUserEdges() *UserQuery {
	return (&HiringTeamClient{config: ht.config}).QueryUserEdges(ht)
}

// QueryHiringTeamJobEdges queries the "hiring_team_job_edges" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryHiringTeamJobEdges() *HiringJobQuery {
	return (&HiringTeamClient{config: ht.config}).QueryHiringTeamJobEdges(ht)
}

// QueryHiringMemberEdges queries the "hiring_member_edges" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryHiringMemberEdges() *UserQuery {
	return (&HiringTeamClient{config: ht.config}).QueryHiringMemberEdges(ht)
}

// QueryApproversUsers queries the "approvers_users" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryApproversUsers() *UserQuery {
	return (&HiringTeamClient{config: ht.config}).QueryApproversUsers(ht)
}

// QueryUserHiringTeams queries the "user_hiring_teams" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryUserHiringTeams() *HiringTeamManagerQuery {
	return (&HiringTeamClient{config: ht.config}).QueryUserHiringTeams(ht)
}

// QueryHiringTeamApprovers queries the "hiring_team_approvers" edge of the HiringTeam entity.
func (ht *HiringTeam) QueryHiringTeamApprovers() *HiringTeamApproverQuery {
	return (&HiringTeamClient{config: ht.config}).QueryHiringTeamApprovers(ht)
}

// Update returns a builder for updating this HiringTeam.
// Note that you need to call HiringTeam.Unwrap() before calling this method if this HiringTeam
// was returned from a transaction, and the transaction was committed or rolled back.
func (ht *HiringTeam) Update() *HiringTeamUpdateOne {
	return (&HiringTeamClient{config: ht.config}).UpdateOne(ht)
}

// Unwrap unwraps the HiringTeam entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ht *HiringTeam) Unwrap() *HiringTeam {
	_tx, ok := ht.config.driver.(*txDriver)
	if !ok {
		panic("ent: HiringTeam is not a transactional entity")
	}
	ht.config.driver = _tx.drv
	return ht
}

// String implements the fmt.Stringer.
func (ht *HiringTeam) String() string {
	var builder strings.Builder
	builder.WriteString("HiringTeam(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ht.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ht.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ht.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ht.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(ht.Slug)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ht.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ht.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUserEdges returns the UserEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedUserEdges(name string) ([]*User, error) {
	if ht.Edges.namedUserEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedUserEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedUserEdges(name string, edges ...*User) {
	if ht.Edges.namedUserEdges == nil {
		ht.Edges.namedUserEdges = make(map[string][]*User)
	}
	if len(edges) == 0 {
		ht.Edges.namedUserEdges[name] = []*User{}
	} else {
		ht.Edges.namedUserEdges[name] = append(ht.Edges.namedUserEdges[name], edges...)
	}
}

// NamedHiringTeamJobEdges returns the HiringTeamJobEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedHiringTeamJobEdges(name string) ([]*HiringJob, error) {
	if ht.Edges.namedHiringTeamJobEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedHiringTeamJobEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedHiringTeamJobEdges(name string, edges ...*HiringJob) {
	if ht.Edges.namedHiringTeamJobEdges == nil {
		ht.Edges.namedHiringTeamJobEdges = make(map[string][]*HiringJob)
	}
	if len(edges) == 0 {
		ht.Edges.namedHiringTeamJobEdges[name] = []*HiringJob{}
	} else {
		ht.Edges.namedHiringTeamJobEdges[name] = append(ht.Edges.namedHiringTeamJobEdges[name], edges...)
	}
}

// NamedHiringMemberEdges returns the HiringMemberEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedHiringMemberEdges(name string) ([]*User, error) {
	if ht.Edges.namedHiringMemberEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedHiringMemberEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedHiringMemberEdges(name string, edges ...*User) {
	if ht.Edges.namedHiringMemberEdges == nil {
		ht.Edges.namedHiringMemberEdges = make(map[string][]*User)
	}
	if len(edges) == 0 {
		ht.Edges.namedHiringMemberEdges[name] = []*User{}
	} else {
		ht.Edges.namedHiringMemberEdges[name] = append(ht.Edges.namedHiringMemberEdges[name], edges...)
	}
}

// NamedApproversUsers returns the ApproversUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedApproversUsers(name string) ([]*User, error) {
	if ht.Edges.namedApproversUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedApproversUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedApproversUsers(name string, edges ...*User) {
	if ht.Edges.namedApproversUsers == nil {
		ht.Edges.namedApproversUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		ht.Edges.namedApproversUsers[name] = []*User{}
	} else {
		ht.Edges.namedApproversUsers[name] = append(ht.Edges.namedApproversUsers[name], edges...)
	}
}

// NamedUserHiringTeams returns the UserHiringTeams named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedUserHiringTeams(name string) ([]*HiringTeamManager, error) {
	if ht.Edges.namedUserHiringTeams == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedUserHiringTeams[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedUserHiringTeams(name string, edges ...*HiringTeamManager) {
	if ht.Edges.namedUserHiringTeams == nil {
		ht.Edges.namedUserHiringTeams = make(map[string][]*HiringTeamManager)
	}
	if len(edges) == 0 {
		ht.Edges.namedUserHiringTeams[name] = []*HiringTeamManager{}
	} else {
		ht.Edges.namedUserHiringTeams[name] = append(ht.Edges.namedUserHiringTeams[name], edges...)
	}
}

// NamedHiringTeamApprovers returns the HiringTeamApprovers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ht *HiringTeam) NamedHiringTeamApprovers(name string) ([]*HiringTeamApprover, error) {
	if ht.Edges.namedHiringTeamApprovers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ht.Edges.namedHiringTeamApprovers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ht *HiringTeam) appendNamedHiringTeamApprovers(name string, edges ...*HiringTeamApprover) {
	if ht.Edges.namedHiringTeamApprovers == nil {
		ht.Edges.namedHiringTeamApprovers = make(map[string][]*HiringTeamApprover)
	}
	if len(edges) == 0 {
		ht.Edges.namedHiringTeamApprovers[name] = []*HiringTeamApprover{}
	} else {
		ht.Edges.namedHiringTeamApprovers[name] = append(ht.Edges.namedHiringTeamApprovers[name], edges...)
	}
}

// HiringTeams is a parsable slice of HiringTeam.
type HiringTeams []*HiringTeam

func (ht HiringTeams) config(cfg config) {
	for _i := range ht {
		ht[_i].config = cfg
	}
}

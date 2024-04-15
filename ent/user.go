// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
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
	// Oid holds the value of the "oid" field.
	Oid string `json:"oid,omitempty"`
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
	// TeamEdges holds the value of the team_edges edge.
	TeamEdges []*Team `json:"team_edges,omitempty"`
	// TeamUsers holds the value of the team_users edge.
	TeamUsers []*TeamManager `json:"team_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedAuditEdge   map[string][]*AuditTrail
	namedHiringOwner map[string][]*HiringJob
	namedTeamEdges   map[string][]*Team
	namedTeamUsers   map[string][]*TeamManager
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

// TeamEdgesOrErr returns the TeamEdges value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TeamEdgesOrErr() ([]*Team, error) {
	if e.loadedTypes[2] {
		return e.TeamEdges, nil
	}
	return nil, &NotLoadedError{edge: "team_edges"}
}

// TeamUsersOrErr returns the TeamUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TeamUsersOrErr() ([]*TeamManager, error) {
	if e.loadedTypes[3] {
		return e.TeamUsers, nil
	}
	return nil, &NotLoadedError{edge: "team_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldName, user.FieldWorkEmail, user.FieldOid:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
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
		case user.FieldOid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				u.Oid = value.String
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

// QueryTeamEdges queries the "team_edges" edge of the User entity.
func (u *User) QueryTeamEdges() *TeamQuery {
	return (&UserClient{config: u.config}).QueryTeamEdges(u)
}

// QueryTeamUsers queries the "team_users" edge of the User entity.
func (u *User) QueryTeamUsers() *TeamManagerQuery {
	return (&UserClient{config: u.config}).QueryTeamUsers(u)
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
	builder.WriteString("oid=")
	builder.WriteString(u.Oid)
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

// NamedTeamEdges returns the TeamEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedTeamEdges(name string) ([]*Team, error) {
	if u.Edges.namedTeamEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedTeamEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedTeamEdges(name string, edges ...*Team) {
	if u.Edges.namedTeamEdges == nil {
		u.Edges.namedTeamEdges = make(map[string][]*Team)
	}
	if len(edges) == 0 {
		u.Edges.namedTeamEdges[name] = []*Team{}
	} else {
		u.Edges.namedTeamEdges[name] = append(u.Edges.namedTeamEdges[name], edges...)
	}
}

// NamedTeamUsers returns the TeamUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedTeamUsers(name string) ([]*TeamManager, error) {
	if u.Edges.namedTeamUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedTeamUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedTeamUsers(name string, edges ...*TeamManager) {
	if u.Edges.namedTeamUsers == nil {
		u.Edges.namedTeamUsers = make(map[string][]*TeamManager)
	}
	if len(edges) == 0 {
		u.Edges.namedTeamUsers[name] = []*TeamManager{}
	} else {
		u.Edges.namedTeamUsers[name] = append(u.Edges.namedTeamUsers[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}

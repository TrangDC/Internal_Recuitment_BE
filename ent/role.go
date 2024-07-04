// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/role"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Role is the model entity for the Role schema.
type Role struct {
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// RolePermissionEdges holds the value of the role_permission_edges edge.
	RolePermissionEdges []*EntityPermission `json:"role_permission_edges,omitempty"`
	// UserEdges holds the value of the user_edges edge.
	UserEdges []*User `json:"user_edges,omitempty"`
	// EmailTemplateEdges holds the value of the email_template_edges edge.
	EmailTemplateEdges []*EmailTemplate `json:"email_template_edges,omitempty"`
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole `json:"user_roles,omitempty"`
	// EmailTemplateRoles holds the value of the email_template_roles edge.
	EmailTemplateRoles []*EmailRoleAttribute `json:"email_template_roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedRolePermissionEdges map[string][]*EntityPermission
	namedUserEdges           map[string][]*User
	namedEmailTemplateEdges  map[string][]*EmailTemplate
	namedUserRoles           map[string][]*UserRole
	namedEmailTemplateRoles  map[string][]*EmailRoleAttribute
}

// RolePermissionEdgesOrErr returns the RolePermissionEdges value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RolePermissionEdgesOrErr() ([]*EntityPermission, error) {
	if e.loadedTypes[0] {
		return e.RolePermissionEdges, nil
	}
	return nil, &NotLoadedError{edge: "role_permission_edges"}
}

// UserEdgesOrErr returns the UserEdges value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserEdgesOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.UserEdges, nil
	}
	return nil, &NotLoadedError{edge: "user_edges"}
}

// EmailTemplateEdgesOrErr returns the EmailTemplateEdges value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) EmailTemplateEdgesOrErr() ([]*EmailTemplate, error) {
	if e.loadedTypes[2] {
		return e.EmailTemplateEdges, nil
	}
	return nil, &NotLoadedError{edge: "email_template_edges"}
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[3] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// EmailTemplateRolesOrErr returns the EmailTemplateRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) EmailTemplateRolesOrErr() ([]*EmailRoleAttribute, error) {
	if e.loadedTypes[4] {
		return e.EmailTemplateRoles, nil
	}
	return nil, &NotLoadedError{edge: "email_template_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldName, role.FieldDescription:
			values[i] = new(sql.NullString)
		case role.FieldCreatedAt, role.FieldUpdatedAt, role.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case role.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case role.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case role.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case role.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = value.Time
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		}
	}
	return nil
}

// QueryRolePermissionEdges queries the "role_permission_edges" edge of the Role entity.
func (r *Role) QueryRolePermissionEdges() *EntityPermissionQuery {
	return (&RoleClient{config: r.config}).QueryRolePermissionEdges(r)
}

// QueryUserEdges queries the "user_edges" edge of the Role entity.
func (r *Role) QueryUserEdges() *UserQuery {
	return (&RoleClient{config: r.config}).QueryUserEdges(r)
}

// QueryEmailTemplateEdges queries the "email_template_edges" edge of the Role entity.
func (r *Role) QueryEmailTemplateEdges() *EmailTemplateQuery {
	return (&RoleClient{config: r.config}).QueryEmailTemplateEdges(r)
}

// QueryUserRoles queries the "user_roles" edge of the Role entity.
func (r *Role) QueryUserRoles() *UserRoleQuery {
	return (&RoleClient{config: r.config}).QueryUserRoles(r)
}

// QueryEmailTemplateRoles queries the "email_template_roles" edge of the Role entity.
func (r *Role) QueryEmailTemplateRoles() *EmailRoleAttributeQuery {
	return (&RoleClient{config: r.config}).QueryEmailTemplateRoles(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(r.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedRolePermissionEdges returns the RolePermissionEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Role) NamedRolePermissionEdges(name string) ([]*EntityPermission, error) {
	if r.Edges.namedRolePermissionEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedRolePermissionEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Role) appendNamedRolePermissionEdges(name string, edges ...*EntityPermission) {
	if r.Edges.namedRolePermissionEdges == nil {
		r.Edges.namedRolePermissionEdges = make(map[string][]*EntityPermission)
	}
	if len(edges) == 0 {
		r.Edges.namedRolePermissionEdges[name] = []*EntityPermission{}
	} else {
		r.Edges.namedRolePermissionEdges[name] = append(r.Edges.namedRolePermissionEdges[name], edges...)
	}
}

// NamedUserEdges returns the UserEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Role) NamedUserEdges(name string) ([]*User, error) {
	if r.Edges.namedUserEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedUserEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Role) appendNamedUserEdges(name string, edges ...*User) {
	if r.Edges.namedUserEdges == nil {
		r.Edges.namedUserEdges = make(map[string][]*User)
	}
	if len(edges) == 0 {
		r.Edges.namedUserEdges[name] = []*User{}
	} else {
		r.Edges.namedUserEdges[name] = append(r.Edges.namedUserEdges[name], edges...)
	}
}

// NamedEmailTemplateEdges returns the EmailTemplateEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Role) NamedEmailTemplateEdges(name string) ([]*EmailTemplate, error) {
	if r.Edges.namedEmailTemplateEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedEmailTemplateEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Role) appendNamedEmailTemplateEdges(name string, edges ...*EmailTemplate) {
	if r.Edges.namedEmailTemplateEdges == nil {
		r.Edges.namedEmailTemplateEdges = make(map[string][]*EmailTemplate)
	}
	if len(edges) == 0 {
		r.Edges.namedEmailTemplateEdges[name] = []*EmailTemplate{}
	} else {
		r.Edges.namedEmailTemplateEdges[name] = append(r.Edges.namedEmailTemplateEdges[name], edges...)
	}
}

// NamedUserRoles returns the UserRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Role) NamedUserRoles(name string) ([]*UserRole, error) {
	if r.Edges.namedUserRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedUserRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Role) appendNamedUserRoles(name string, edges ...*UserRole) {
	if r.Edges.namedUserRoles == nil {
		r.Edges.namedUserRoles = make(map[string][]*UserRole)
	}
	if len(edges) == 0 {
		r.Edges.namedUserRoles[name] = []*UserRole{}
	} else {
		r.Edges.namedUserRoles[name] = append(r.Edges.namedUserRoles[name], edges...)
	}
}

// NamedEmailTemplateRoles returns the EmailTemplateRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Role) NamedEmailTemplateRoles(name string) ([]*EmailRoleAttribute, error) {
	if r.Edges.namedEmailTemplateRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedEmailTemplateRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Role) appendNamedEmailTemplateRoles(name string, edges ...*EmailRoleAttribute) {
	if r.Edges.namedEmailTemplateRoles == nil {
		r.Edges.namedEmailTemplateRoles = make(map[string][]*EmailRoleAttribute)
	}
	if len(edges) == 0 {
		r.Edges.namedEmailTemplateRoles[name] = []*EmailRoleAttribute{}
	} else {
		r.Edges.namedEmailTemplateRoles[name] = append(r.Edges.namedEmailTemplateRoles[name], edges...)
	}
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

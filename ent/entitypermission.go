// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/entitypermission"
	"trec/ent/permission"
	"trec/ent/role"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// EntityPermission is the model entity for the EntityPermission schema.
type EntityPermission struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// EntityID holds the value of the "entity_id" field.
	EntityID uuid.UUID `json:"entity_id,omitempty"`
	// PermissionID holds the value of the "permission_id" field.
	PermissionID uuid.UUID `json:"permission_id,omitempty"`
	// ForOwner holds the value of the "for_owner" field.
	ForOwner bool `json:"for_owner,omitempty"`
	// ForTeam holds the value of the "for_team" field.
	ForTeam bool `json:"for_team,omitempty"`
	// ForAll holds the value of the "for_all" field.
	ForAll bool `json:"for_all,omitempty"`
	// EntityType holds the value of the "entity_type" field.
	EntityType entitypermission.EntityType `json:"entity_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntityPermissionQuery when eager-loading is set.
	Edges EntityPermissionEdges `json:"edges"`
}

// EntityPermissionEdges holds the relations/edges for other nodes in the graph.
type EntityPermissionEdges struct {
	// PermissionEdges holds the value of the permission_edges edge.
	PermissionEdges *Permission `json:"permission_edges,omitempty"`
	// UserEdge holds the value of the user_edge edge.
	UserEdge *User `json:"user_edge,omitempty"`
	// RoleEdge holds the value of the role_edge edge.
	RoleEdge *Role `json:"role_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// PermissionEdgesOrErr returns the PermissionEdges value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntityPermissionEdges) PermissionEdgesOrErr() (*Permission, error) {
	if e.loadedTypes[0] {
		if e.PermissionEdges == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: permission.Label}
		}
		return e.PermissionEdges, nil
	}
	return nil, &NotLoadedError{edge: "permission_edges"}
}

// UserEdgeOrErr returns the UserEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntityPermissionEdges) UserEdgeOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.UserEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserEdge, nil
	}
	return nil, &NotLoadedError{edge: "user_edge"}
}

// RoleEdgeOrErr returns the RoleEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntityPermissionEdges) RoleEdgeOrErr() (*Role, error) {
	if e.loadedTypes[2] {
		if e.RoleEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.RoleEdge, nil
	}
	return nil, &NotLoadedError{edge: "role_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntityPermission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entitypermission.FieldForOwner, entitypermission.FieldForTeam, entitypermission.FieldForAll:
			values[i] = new(sql.NullBool)
		case entitypermission.FieldEntityType:
			values[i] = new(sql.NullString)
		case entitypermission.FieldCreatedAt, entitypermission.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case entitypermission.FieldID, entitypermission.FieldEntityID, entitypermission.FieldPermissionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntityPermission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntityPermission fields.
func (ep *EntityPermission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entitypermission.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ep.ID = *value
			}
		case entitypermission.FieldEntityID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field entity_id", values[i])
			} else if value != nil {
				ep.EntityID = *value
			}
		case entitypermission.FieldPermissionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field permission_id", values[i])
			} else if value != nil {
				ep.PermissionID = *value
			}
		case entitypermission.FieldForOwner:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_owner", values[i])
			} else if value.Valid {
				ep.ForOwner = value.Bool
			}
		case entitypermission.FieldForTeam:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_team", values[i])
			} else if value.Valid {
				ep.ForTeam = value.Bool
			}
		case entitypermission.FieldForAll:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field for_all", values[i])
			} else if value.Valid {
				ep.ForAll = value.Bool
			}
		case entitypermission.FieldEntityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_type", values[i])
			} else if value.Valid {
				ep.EntityType = entitypermission.EntityType(value.String)
			}
		case entitypermission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ep.CreatedAt = value.Time
			}
		case entitypermission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ep.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPermissionEdges queries the "permission_edges" edge of the EntityPermission entity.
func (ep *EntityPermission) QueryPermissionEdges() *PermissionQuery {
	return (&EntityPermissionClient{config: ep.config}).QueryPermissionEdges(ep)
}

// QueryUserEdge queries the "user_edge" edge of the EntityPermission entity.
func (ep *EntityPermission) QueryUserEdge() *UserQuery {
	return (&EntityPermissionClient{config: ep.config}).QueryUserEdge(ep)
}

// QueryRoleEdge queries the "role_edge" edge of the EntityPermission entity.
func (ep *EntityPermission) QueryRoleEdge() *RoleQuery {
	return (&EntityPermissionClient{config: ep.config}).QueryRoleEdge(ep)
}

// Update returns a builder for updating this EntityPermission.
// Note that you need to call EntityPermission.Unwrap() before calling this method if this EntityPermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EntityPermission) Update() *EntityPermissionUpdateOne {
	return (&EntityPermissionClient{config: ep.config}).UpdateOne(ep)
}

// Unwrap unwraps the EntityPermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *EntityPermission) Unwrap() *EntityPermission {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntityPermission is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EntityPermission) String() string {
	var builder strings.Builder
	builder.WriteString("EntityPermission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("entity_id=")
	builder.WriteString(fmt.Sprintf("%v", ep.EntityID))
	builder.WriteString(", ")
	builder.WriteString("permission_id=")
	builder.WriteString(fmt.Sprintf("%v", ep.PermissionID))
	builder.WriteString(", ")
	builder.WriteString("for_owner=")
	builder.WriteString(fmt.Sprintf("%v", ep.ForOwner))
	builder.WriteString(", ")
	builder.WriteString("for_team=")
	builder.WriteString(fmt.Sprintf("%v", ep.ForTeam))
	builder.WriteString(", ")
	builder.WriteString("for_all=")
	builder.WriteString(fmt.Sprintf("%v", ep.ForAll))
	builder.WriteString(", ")
	builder.WriteString("entity_type=")
	builder.WriteString(fmt.Sprintf("%v", ep.EntityType))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ep.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ep.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EntityPermissions is a parsable slice of EntityPermission.
type EntityPermissions []*EntityPermission

func (ep EntityPermissions) config(cfg config) {
	for _i := range ep {
		ep[_i].config = cfg
	}
}
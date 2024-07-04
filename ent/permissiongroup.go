// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/permissiongroup"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PermissionGroup is the model entity for the PermissionGroup schema.
type PermissionGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID uuid.UUID `json:"parent_id,omitempty"`
	// GroupType holds the value of the "group_type" field.
	GroupType permissiongroup.GroupType `json:"group_type,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int `json:"order_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionGroupQuery when eager-loading is set.
	Edges PermissionGroupEdges `json:"edges"`
}

// PermissionGroupEdges holds the relations/edges for other nodes in the graph.
type PermissionGroupEdges struct {
	// GroupPermissionParent holds the value of the group_permission_parent edge.
	GroupPermissionParent *PermissionGroup `json:"group_permission_parent,omitempty"`
	// GroupPermissionChildren holds the value of the group_permission_children edge.
	GroupPermissionChildren []*PermissionGroup `json:"group_permission_children,omitempty"`
	// PermissionEdges holds the value of the permission_edges edge.
	PermissionEdges []*Permission `json:"permission_edges,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedGroupPermissionChildren map[string][]*PermissionGroup
	namedPermissionEdges         map[string][]*Permission
}

// GroupPermissionParentOrErr returns the GroupPermissionParent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionGroupEdges) GroupPermissionParentOrErr() (*PermissionGroup, error) {
	if e.loadedTypes[0] {
		if e.GroupPermissionParent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: permissiongroup.Label}
		}
		return e.GroupPermissionParent, nil
	}
	return nil, &NotLoadedError{edge: "group_permission_parent"}
}

// GroupPermissionChildrenOrErr returns the GroupPermissionChildren value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionGroupEdges) GroupPermissionChildrenOrErr() ([]*PermissionGroup, error) {
	if e.loadedTypes[1] {
		return e.GroupPermissionChildren, nil
	}
	return nil, &NotLoadedError{edge: "group_permission_children"}
}

// PermissionEdgesOrErr returns the PermissionEdges value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionGroupEdges) PermissionEdgesOrErr() ([]*Permission, error) {
	if e.loadedTypes[2] {
		return e.PermissionEdges, nil
	}
	return nil, &NotLoadedError{edge: "permission_edges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissiongroup.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case permissiongroup.FieldTitle, permissiongroup.FieldGroupType:
			values[i] = new(sql.NullString)
		case permissiongroup.FieldCreatedAt, permissiongroup.FieldUpdatedAt, permissiongroup.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case permissiongroup.FieldID, permissiongroup.FieldParentID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PermissionGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PermissionGroup fields.
func (pg *PermissionGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permissiongroup.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pg.ID = *value
			}
		case permissiongroup.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pg.CreatedAt = value.Time
			}
		case permissiongroup.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pg.UpdatedAt = value.Time
			}
		case permissiongroup.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pg.DeletedAt = value.Time
			}
		case permissiongroup.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pg.Title = value.String
			}
		case permissiongroup.FieldParentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value != nil {
				pg.ParentID = *value
			}
		case permissiongroup.FieldGroupType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_type", values[i])
			} else if value.Valid {
				pg.GroupType = permissiongroup.GroupType(value.String)
			}
		case permissiongroup.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				pg.OrderID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryGroupPermissionParent queries the "group_permission_parent" edge of the PermissionGroup entity.
func (pg *PermissionGroup) QueryGroupPermissionParent() *PermissionGroupQuery {
	return (&PermissionGroupClient{config: pg.config}).QueryGroupPermissionParent(pg)
}

// QueryGroupPermissionChildren queries the "group_permission_children" edge of the PermissionGroup entity.
func (pg *PermissionGroup) QueryGroupPermissionChildren() *PermissionGroupQuery {
	return (&PermissionGroupClient{config: pg.config}).QueryGroupPermissionChildren(pg)
}

// QueryPermissionEdges queries the "permission_edges" edge of the PermissionGroup entity.
func (pg *PermissionGroup) QueryPermissionEdges() *PermissionQuery {
	return (&PermissionGroupClient{config: pg.config}).QueryPermissionEdges(pg)
}

// Update returns a builder for updating this PermissionGroup.
// Note that you need to call PermissionGroup.Unwrap() before calling this method if this PermissionGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (pg *PermissionGroup) Update() *PermissionGroupUpdateOne {
	return (&PermissionGroupClient{config: pg.config}).UpdateOne(pg)
}

// Unwrap unwraps the PermissionGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pg *PermissionGroup) Unwrap() *PermissionGroup {
	_tx, ok := pg.config.driver.(*txDriver)
	if !ok {
		panic("ent: PermissionGroup is not a transactional entity")
	}
	pg.config.driver = _tx.drv
	return pg
}

// String implements the fmt.Stringer.
func (pg *PermissionGroup) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pg.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pg.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pg.Title)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", pg.ParentID))
	builder.WriteString(", ")
	builder.WriteString("group_type=")
	builder.WriteString(fmt.Sprintf("%v", pg.GroupType))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", pg.OrderID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedGroupPermissionChildren returns the GroupPermissionChildren named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pg *PermissionGroup) NamedGroupPermissionChildren(name string) ([]*PermissionGroup, error) {
	if pg.Edges.namedGroupPermissionChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pg.Edges.namedGroupPermissionChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pg *PermissionGroup) appendNamedGroupPermissionChildren(name string, edges ...*PermissionGroup) {
	if pg.Edges.namedGroupPermissionChildren == nil {
		pg.Edges.namedGroupPermissionChildren = make(map[string][]*PermissionGroup)
	}
	if len(edges) == 0 {
		pg.Edges.namedGroupPermissionChildren[name] = []*PermissionGroup{}
	} else {
		pg.Edges.namedGroupPermissionChildren[name] = append(pg.Edges.namedGroupPermissionChildren[name], edges...)
	}
}

// NamedPermissionEdges returns the PermissionEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pg *PermissionGroup) NamedPermissionEdges(name string) ([]*Permission, error) {
	if pg.Edges.namedPermissionEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pg.Edges.namedPermissionEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pg *PermissionGroup) appendNamedPermissionEdges(name string, edges ...*Permission) {
	if pg.Edges.namedPermissionEdges == nil {
		pg.Edges.namedPermissionEdges = make(map[string][]*Permission)
	}
	if len(edges) == 0 {
		pg.Edges.namedPermissionEdges[name] = []*Permission{}
	} else {
		pg.Edges.namedPermissionEdges[name] = append(pg.Edges.namedPermissionEdges[name], edges...)
	}
}

// PermissionGroups is a parsable slice of PermissionGroup.
type PermissionGroups []*PermissionGroup

func (pg PermissionGroups) config(cfg config) {
	for _i := range pg {
		pg[_i].config = cfg
	}
}

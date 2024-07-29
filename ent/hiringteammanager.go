// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/hiringteam"
	"trec/ent/hiringteammanager"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// HiringTeamManager is the model entity for the HiringTeamManager schema.
type HiringTeamManager struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// HiringTeamID holds the value of the "hiring_team_id" field.
	HiringTeamID uuid.UUID `json:"hiring_team_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HiringTeamManagerQuery when eager-loading is set.
	Edges HiringTeamManagerEdges `json:"edges"`
}

// HiringTeamManagerEdges holds the relations/edges for other nodes in the graph.
type HiringTeamManagerEdges struct {
	// UserEdge holds the value of the user_edge edge.
	UserEdge *User `json:"user_edge,omitempty"`
	// HiringTeamEdge holds the value of the hiring_team_edge edge.
	HiringTeamEdge *HiringTeam `json:"hiring_team_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserEdgeOrErr returns the UserEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HiringTeamManagerEdges) UserEdgeOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserEdge, nil
	}
	return nil, &NotLoadedError{edge: "user_edge"}
}

// HiringTeamEdgeOrErr returns the HiringTeamEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HiringTeamManagerEdges) HiringTeamEdgeOrErr() (*HiringTeam, error) {
	if e.loadedTypes[1] {
		if e.HiringTeamEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hiringteam.Label}
		}
		return e.HiringTeamEdge, nil
	}
	return nil, &NotLoadedError{edge: "hiring_team_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HiringTeamManager) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hiringteammanager.FieldCreatedAt, hiringteammanager.FieldUpdatedAt, hiringteammanager.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case hiringteammanager.FieldID, hiringteammanager.FieldHiringTeamID, hiringteammanager.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type HiringTeamManager", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HiringTeamManager fields.
func (htm *HiringTeamManager) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hiringteammanager.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				htm.ID = *value
			}
		case hiringteammanager.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				htm.CreatedAt = value.Time
			}
		case hiringteammanager.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				htm.UpdatedAt = value.Time
			}
		case hiringteammanager.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				htm.DeletedAt = value.Time
			}
		case hiringteammanager.FieldHiringTeamID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field hiring_team_id", values[i])
			} else if value != nil {
				htm.HiringTeamID = *value
			}
		case hiringteammanager.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				htm.UserID = *value
			}
		}
	}
	return nil
}

// QueryUserEdge queries the "user_edge" edge of the HiringTeamManager entity.
func (htm *HiringTeamManager) QueryUserEdge() *UserQuery {
	return (&HiringTeamManagerClient{config: htm.config}).QueryUserEdge(htm)
}

// QueryHiringTeamEdge queries the "hiring_team_edge" edge of the HiringTeamManager entity.
func (htm *HiringTeamManager) QueryHiringTeamEdge() *HiringTeamQuery {
	return (&HiringTeamManagerClient{config: htm.config}).QueryHiringTeamEdge(htm)
}

// Update returns a builder for updating this HiringTeamManager.
// Note that you need to call HiringTeamManager.Unwrap() before calling this method if this HiringTeamManager
// was returned from a transaction, and the transaction was committed or rolled back.
func (htm *HiringTeamManager) Update() *HiringTeamManagerUpdateOne {
	return (&HiringTeamManagerClient{config: htm.config}).UpdateOne(htm)
}

// Unwrap unwraps the HiringTeamManager entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (htm *HiringTeamManager) Unwrap() *HiringTeamManager {
	_tx, ok := htm.config.driver.(*txDriver)
	if !ok {
		panic("ent: HiringTeamManager is not a transactional entity")
	}
	htm.config.driver = _tx.drv
	return htm
}

// String implements the fmt.Stringer.
func (htm *HiringTeamManager) String() string {
	var builder strings.Builder
	builder.WriteString("HiringTeamManager(")
	builder.WriteString(fmt.Sprintf("id=%v, ", htm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(htm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(htm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(htm.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("hiring_team_id=")
	builder.WriteString(fmt.Sprintf("%v", htm.HiringTeamID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", htm.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// HiringTeamManagers is a parsable slice of HiringTeamManager.
type HiringTeamManagers []*HiringTeamManager

func (htm HiringTeamManagers) config(cfg config) {
	for _i := range htm {
		htm[_i].config = cfg
	}
}
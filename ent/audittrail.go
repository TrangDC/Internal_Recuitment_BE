// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/audittrail"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AuditTrail is the model entity for the AuditTrail schema.
type AuditTrail struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// RecordId holds the value of the "recordId" field.
	RecordId uuid.UUID `json:"recordId,omitempty"`
	// Module holds the value of the "module" field.
	Module audittrail.Module `json:"module,omitempty"`
	// ActionType holds the value of the "actionType" field.
	ActionType audittrail.ActionType `json:"actionType,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// RecordChanges holds the value of the "record_changes" field.
	RecordChanges string `json:"record_changes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuditTrailQuery when eager-loading is set.
	Edges AuditTrailEdges `json:"edges"`
}

// AuditTrailEdges holds the relations/edges for other nodes in the graph.
type AuditTrailEdges struct {
	// UserEdge holds the value of the user_edge edge.
	UserEdge *User `json:"user_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// UserEdgeOrErr returns the UserEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuditTrailEdges) UserEdgeOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserEdge, nil
	}
	return nil, &NotLoadedError{edge: "user_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuditTrail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case audittrail.FieldModule, audittrail.FieldActionType, audittrail.FieldNote, audittrail.FieldRecordChanges:
			values[i] = new(sql.NullString)
		case audittrail.FieldCreatedAt, audittrail.FieldUpdatedAt, audittrail.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case audittrail.FieldID, audittrail.FieldCreatedBy, audittrail.FieldRecordId:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuditTrail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuditTrail fields.
func (at *AuditTrail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audittrail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case audittrail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = value.Time
			}
		case audittrail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				at.UpdatedAt = value.Time
			}
		case audittrail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				at.DeletedAt = value.Time
			}
		case audittrail.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				at.CreatedBy = *value
			}
		case audittrail.FieldRecordId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field recordId", values[i])
			} else if value != nil {
				at.RecordId = *value
			}
		case audittrail.FieldModule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module", values[i])
			} else if value.Valid {
				at.Module = audittrail.Module(value.String)
			}
		case audittrail.FieldActionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actionType", values[i])
			} else if value.Valid {
				at.ActionType = audittrail.ActionType(value.String)
			}
		case audittrail.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				at.Note = value.String
			}
		case audittrail.FieldRecordChanges:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record_changes", values[i])
			} else if value.Valid {
				at.RecordChanges = value.String
			}
		}
	}
	return nil
}

// QueryUserEdge queries the "user_edge" edge of the AuditTrail entity.
func (at *AuditTrail) QueryUserEdge() *UserQuery {
	return (&AuditTrailClient{config: at.config}).QueryUserEdge(at)
}

// Update returns a builder for updating this AuditTrail.
// Note that you need to call AuditTrail.Unwrap() before calling this method if this AuditTrail
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AuditTrail) Update() *AuditTrailUpdateOne {
	return (&AuditTrailClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the AuditTrail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AuditTrail) Unwrap() *AuditTrail {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuditTrail is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AuditTrail) String() string {
	var builder strings.Builder
	builder.WriteString("AuditTrail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("created_at=")
	builder.WriteString(at.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(at.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(at.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", at.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("recordId=")
	builder.WriteString(fmt.Sprintf("%v", at.RecordId))
	builder.WriteString(", ")
	builder.WriteString("module=")
	builder.WriteString(fmt.Sprintf("%v", at.Module))
	builder.WriteString(", ")
	builder.WriteString("actionType=")
	builder.WriteString(fmt.Sprintf("%v", at.ActionType))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(at.Note)
	builder.WriteString(", ")
	builder.WriteString("record_changes=")
	builder.WriteString(at.RecordChanges)
	builder.WriteByte(')')
	return builder.String()
}

// AuditTrails is a parsable slice of AuditTrail.
type AuditTrails []*AuditTrail

func (at AuditTrails) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}

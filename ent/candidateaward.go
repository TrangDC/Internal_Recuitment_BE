// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/candidate"
	"trec/ent/candidateaward"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CandidateAward is the model entity for the CandidateAward schema.
type CandidateAward struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// CandidateID holds the value of the "candidate_id" field.
	CandidateID uuid.UUID `json:"candidate_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AchievedDate holds the value of the "achieved_date" field.
	AchievedDate time.Time `json:"achieved_date,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int `json:"order_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CandidateAwardQuery when eager-loading is set.
	Edges CandidateAwardEdges `json:"edges"`
}

// CandidateAwardEdges holds the relations/edges for other nodes in the graph.
type CandidateAwardEdges struct {
	// AttachmentEdges holds the value of the attachment_edges edge.
	AttachmentEdges []*Attachment `json:"attachment_edges,omitempty"`
	// CandidateEdge holds the value of the candidate_edge edge.
	CandidateEdge *Candidate `json:"candidate_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedAttachmentEdges map[string][]*Attachment
}

// AttachmentEdgesOrErr returns the AttachmentEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateAwardEdges) AttachmentEdgesOrErr() ([]*Attachment, error) {
	if e.loadedTypes[0] {
		return e.AttachmentEdges, nil
	}
	return nil, &NotLoadedError{edge: "attachment_edges"}
}

// CandidateEdgeOrErr returns the CandidateEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CandidateAwardEdges) CandidateEdgeOrErr() (*Candidate, error) {
	if e.loadedTypes[1] {
		if e.CandidateEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: candidate.Label}
		}
		return e.CandidateEdge, nil
	}
	return nil, &NotLoadedError{edge: "candidate_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CandidateAward) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case candidateaward.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case candidateaward.FieldName:
			values[i] = new(sql.NullString)
		case candidateaward.FieldCreatedAt, candidateaward.FieldUpdatedAt, candidateaward.FieldDeletedAt, candidateaward.FieldAchievedDate:
			values[i] = new(sql.NullTime)
		case candidateaward.FieldID, candidateaward.FieldCandidateID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CandidateAward", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CandidateAward fields.
func (ca *CandidateAward) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case candidateaward.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ca.ID = *value
			}
		case candidateaward.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ca.CreatedAt = value.Time
			}
		case candidateaward.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ca.UpdatedAt = value.Time
			}
		case candidateaward.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ca.DeletedAt = value.Time
			}
		case candidateaward.FieldCandidateID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field candidate_id", values[i])
			} else if value != nil {
				ca.CandidateID = *value
			}
		case candidateaward.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ca.Name = value.String
			}
		case candidateaward.FieldAchievedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field achieved_date", values[i])
			} else if value.Valid {
				ca.AchievedDate = value.Time
			}
		case candidateaward.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				ca.OrderID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAttachmentEdges queries the "attachment_edges" edge of the CandidateAward entity.
func (ca *CandidateAward) QueryAttachmentEdges() *AttachmentQuery {
	return (&CandidateAwardClient{config: ca.config}).QueryAttachmentEdges(ca)
}

// QueryCandidateEdge queries the "candidate_edge" edge of the CandidateAward entity.
func (ca *CandidateAward) QueryCandidateEdge() *CandidateQuery {
	return (&CandidateAwardClient{config: ca.config}).QueryCandidateEdge(ca)
}

// Update returns a builder for updating this CandidateAward.
// Note that you need to call CandidateAward.Unwrap() before calling this method if this CandidateAward
// was returned from a transaction, and the transaction was committed or rolled back.
func (ca *CandidateAward) Update() *CandidateAwardUpdateOne {
	return (&CandidateAwardClient{config: ca.config}).UpdateOne(ca)
}

// Unwrap unwraps the CandidateAward entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ca *CandidateAward) Unwrap() *CandidateAward {
	_tx, ok := ca.config.driver.(*txDriver)
	if !ok {
		panic("ent: CandidateAward is not a transactional entity")
	}
	ca.config.driver = _tx.drv
	return ca
}

// String implements the fmt.Stringer.
func (ca *CandidateAward) String() string {
	var builder strings.Builder
	builder.WriteString("CandidateAward(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ca.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ca.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ca.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ca.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("candidate_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.CandidateID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ca.Name)
	builder.WriteString(", ")
	builder.WriteString("achieved_date=")
	builder.WriteString(ca.AchievedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", ca.OrderID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAttachmentEdges returns the AttachmentEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ca *CandidateAward) NamedAttachmentEdges(name string) ([]*Attachment, error) {
	if ca.Edges.namedAttachmentEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ca.Edges.namedAttachmentEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ca *CandidateAward) appendNamedAttachmentEdges(name string, edges ...*Attachment) {
	if ca.Edges.namedAttachmentEdges == nil {
		ca.Edges.namedAttachmentEdges = make(map[string][]*Attachment)
	}
	if len(edges) == 0 {
		ca.Edges.namedAttachmentEdges[name] = []*Attachment{}
	} else {
		ca.Edges.namedAttachmentEdges[name] = append(ca.Edges.namedAttachmentEdges[name], edges...)
	}
}

// CandidateAwards is a parsable slice of CandidateAward.
type CandidateAwards []*CandidateAward

func (ca CandidateAwards) config(cfg config) {
	for _i := range ca {
		ca[_i].config = cfg
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"trec/ent/candidate"
	"trec/ent/outgoingemail"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// OutgoingEmail is the model entity for the OutgoingEmail schema.
type OutgoingEmail struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// To holds the value of the "to" field.
	To []string `json:"to,omitempty"`
	// Cc holds the value of the "cc" field.
	Cc []string `json:"cc,omitempty"`
	// Bcc holds the value of the "bcc" field.
	Bcc []string `json:"bcc,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature string `json:"signature,omitempty"`
	// CandidateID holds the value of the "candidate_id" field.
	CandidateID uuid.UUID `json:"candidate_id,omitempty"`
	// RecipientType holds the value of the "recipient_type" field.
	RecipientType outgoingemail.RecipientType `json:"recipient_type,omitempty"`
	// EmailTemplateID holds the value of the "email_template_id" field.
	EmailTemplateID uuid.UUID `json:"email_template_id,omitempty"`
	// Status holds the value of the "status" field.
	Status outgoingemail.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutgoingEmailQuery when eager-loading is set.
	Edges OutgoingEmailEdges `json:"edges"`
}

// OutgoingEmailEdges holds the relations/edges for other nodes in the graph.
type OutgoingEmailEdges struct {
	// CandidateEdge holds the value of the candidate_edge edge.
	CandidateEdge *Candidate `json:"candidate_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// CandidateEdgeOrErr returns the CandidateEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutgoingEmailEdges) CandidateEdgeOrErr() (*Candidate, error) {
	if e.loadedTypes[0] {
		if e.CandidateEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: candidate.Label}
		}
		return e.CandidateEdge, nil
	}
	return nil, &NotLoadedError{edge: "candidate_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutgoingEmail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case outgoingemail.FieldTo, outgoingemail.FieldCc, outgoingemail.FieldBcc:
			values[i] = new([]byte)
		case outgoingemail.FieldSubject, outgoingemail.FieldContent, outgoingemail.FieldSignature, outgoingemail.FieldRecipientType, outgoingemail.FieldStatus:
			values[i] = new(sql.NullString)
		case outgoingemail.FieldCreatedAt, outgoingemail.FieldUpdatedAt, outgoingemail.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case outgoingemail.FieldID, outgoingemail.FieldCandidateID, outgoingemail.FieldEmailTemplateID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OutgoingEmail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutgoingEmail fields.
func (oe *OutgoingEmail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outgoingemail.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				oe.ID = *value
			}
		case outgoingemail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oe.CreatedAt = value.Time
			}
		case outgoingemail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oe.UpdatedAt = value.Time
			}
		case outgoingemail.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oe.DeletedAt = value.Time
			}
		case outgoingemail.FieldTo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oe.To); err != nil {
					return fmt.Errorf("unmarshal field to: %w", err)
				}
			}
		case outgoingemail.FieldCc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oe.Cc); err != nil {
					return fmt.Errorf("unmarshal field cc: %w", err)
				}
			}
		case outgoingemail.FieldBcc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field bcc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oe.Bcc); err != nil {
					return fmt.Errorf("unmarshal field bcc: %w", err)
				}
			}
		case outgoingemail.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				oe.Subject = value.String
			}
		case outgoingemail.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				oe.Content = value.String
			}
		case outgoingemail.FieldSignature:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value.Valid {
				oe.Signature = value.String
			}
		case outgoingemail.FieldCandidateID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field candidate_id", values[i])
			} else if value != nil {
				oe.CandidateID = *value
			}
		case outgoingemail.FieldRecipientType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient_type", values[i])
			} else if value.Valid {
				oe.RecipientType = outgoingemail.RecipientType(value.String)
			}
		case outgoingemail.FieldEmailTemplateID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field email_template_id", values[i])
			} else if value != nil {
				oe.EmailTemplateID = *value
			}
		case outgoingemail.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				oe.Status = outgoingemail.Status(value.String)
			}
		}
	}
	return nil
}

// QueryCandidateEdge queries the "candidate_edge" edge of the OutgoingEmail entity.
func (oe *OutgoingEmail) QueryCandidateEdge() *CandidateQuery {
	return (&OutgoingEmailClient{config: oe.config}).QueryCandidateEdge(oe)
}

// Update returns a builder for updating this OutgoingEmail.
// Note that you need to call OutgoingEmail.Unwrap() before calling this method if this OutgoingEmail
// was returned from a transaction, and the transaction was committed or rolled back.
func (oe *OutgoingEmail) Update() *OutgoingEmailUpdateOne {
	return (&OutgoingEmailClient{config: oe.config}).UpdateOne(oe)
}

// Unwrap unwraps the OutgoingEmail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oe *OutgoingEmail) Unwrap() *OutgoingEmail {
	_tx, ok := oe.config.driver.(*txDriver)
	if !ok {
		panic("ent: OutgoingEmail is not a transactional entity")
	}
	oe.config.driver = _tx.drv
	return oe
}

// String implements the fmt.Stringer.
func (oe *OutgoingEmail) String() string {
	var builder strings.Builder
	builder.WriteString("OutgoingEmail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oe.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(fmt.Sprintf("%v", oe.To))
	builder.WriteString(", ")
	builder.WriteString("cc=")
	builder.WriteString(fmt.Sprintf("%v", oe.Cc))
	builder.WriteString(", ")
	builder.WriteString("bcc=")
	builder.WriteString(fmt.Sprintf("%v", oe.Bcc))
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(oe.Subject)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(oe.Content)
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(oe.Signature)
	builder.WriteString(", ")
	builder.WriteString("candidate_id=")
	builder.WriteString(fmt.Sprintf("%v", oe.CandidateID))
	builder.WriteString(", ")
	builder.WriteString("recipient_type=")
	builder.WriteString(fmt.Sprintf("%v", oe.RecipientType))
	builder.WriteString(", ")
	builder.WriteString("email_template_id=")
	builder.WriteString(fmt.Sprintf("%v", oe.EmailTemplateID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", oe.Status))
	builder.WriteByte(')')
	return builder.String()
}

// OutgoingEmails is a parsable slice of OutgoingEmail.
type OutgoingEmails []*OutgoingEmail

func (oe OutgoingEmails) config(cfg config) {
	for _i := range oe {
		oe[_i].config = cfg
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/candidateinterview"
	"trec/ent/candidateinterviewer"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CandidateInterviewer is the model entity for the CandidateInterviewer schema.
type CandidateInterviewer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// CandidateInterviewID holds the value of the "candidate_interview_id" field.
	CandidateInterviewID uuid.UUID `json:"candidate_interview_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CandidateInterviewerQuery when eager-loading is set.
	Edges CandidateInterviewerEdges `json:"edges"`
}

// CandidateInterviewerEdges holds the relations/edges for other nodes in the graph.
type CandidateInterviewerEdges struct {
	// UserEdge holds the value of the user_edge edge.
	UserEdge *User `json:"user_edge,omitempty"`
	// InterviewEdge holds the value of the interview_edge edge.
	InterviewEdge *CandidateInterview `json:"interview_edge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserEdgeOrErr returns the UserEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CandidateInterviewerEdges) UserEdgeOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserEdge, nil
	}
	return nil, &NotLoadedError{edge: "user_edge"}
}

// InterviewEdgeOrErr returns the InterviewEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CandidateInterviewerEdges) InterviewEdgeOrErr() (*CandidateInterview, error) {
	if e.loadedTypes[1] {
		if e.InterviewEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: candidateinterview.Label}
		}
		return e.InterviewEdge, nil
	}
	return nil, &NotLoadedError{edge: "interview_edge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CandidateInterviewer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case candidateinterviewer.FieldCreatedAt, candidateinterviewer.FieldUpdatedAt, candidateinterviewer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case candidateinterviewer.FieldID, candidateinterviewer.FieldCandidateInterviewID, candidateinterviewer.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CandidateInterviewer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CandidateInterviewer fields.
func (ci *CandidateInterviewer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case candidateinterviewer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case candidateinterviewer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ci.CreatedAt = value.Time
			}
		case candidateinterviewer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ci.UpdatedAt = value.Time
			}
		case candidateinterviewer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ci.DeletedAt = value.Time
			}
		case candidateinterviewer.FieldCandidateInterviewID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field candidate_interview_id", values[i])
			} else if value != nil {
				ci.CandidateInterviewID = *value
			}
		case candidateinterviewer.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ci.UserID = *value
			}
		}
	}
	return nil
}

// QueryUserEdge queries the "user_edge" edge of the CandidateInterviewer entity.
func (ci *CandidateInterviewer) QueryUserEdge() *UserQuery {
	return (&CandidateInterviewerClient{config: ci.config}).QueryUserEdge(ci)
}

// QueryInterviewEdge queries the "interview_edge" edge of the CandidateInterviewer entity.
func (ci *CandidateInterviewer) QueryInterviewEdge() *CandidateInterviewQuery {
	return (&CandidateInterviewerClient{config: ci.config}).QueryInterviewEdge(ci)
}

// Update returns a builder for updating this CandidateInterviewer.
// Note that you need to call CandidateInterviewer.Unwrap() before calling this method if this CandidateInterviewer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CandidateInterviewer) Update() *CandidateInterviewerUpdateOne {
	return (&CandidateInterviewerClient{config: ci.config}).UpdateOne(ci)
}

// Unwrap unwraps the CandidateInterviewer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CandidateInterviewer) Unwrap() *CandidateInterviewer {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CandidateInterviewer is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CandidateInterviewer) String() string {
	var builder strings.Builder
	builder.WriteString("CandidateInterviewer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ci.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ci.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ci.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("candidate_interview_id=")
	builder.WriteString(fmt.Sprintf("%v", ci.CandidateInterviewID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ci.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// CandidateInterviewers is a parsable slice of CandidateInterviewer.
type CandidateInterviewers []*CandidateInterviewer

func (ci CandidateInterviewers) config(cfg config) {
	for _i := range ci {
		ci[_i].config = cfg
	}
}

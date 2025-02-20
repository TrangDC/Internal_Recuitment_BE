// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/jobposition"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// JobPosition is the model entity for the JobPosition schema.
type JobPosition struct {
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
	// The values are being populated by the JobPositionQuery when eager-loading is set.
	Edges JobPositionEdges `json:"edges"`
}

// JobPositionEdges holds the relations/edges for other nodes in the graph.
type JobPositionEdges struct {
	// HiringJobPositionEdges holds the value of the hiring_job_position_edges edge.
	HiringJobPositionEdges []*HiringJob `json:"hiring_job_position_edges,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedHiringJobPositionEdges map[string][]*HiringJob
}

// HiringJobPositionEdgesOrErr returns the HiringJobPositionEdges value or an error if the edge
// was not loaded in eager-loading.
func (e JobPositionEdges) HiringJobPositionEdgesOrErr() ([]*HiringJob, error) {
	if e.loadedTypes[0] {
		return e.HiringJobPositionEdges, nil
	}
	return nil, &NotLoadedError{edge: "hiring_job_position_edges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobPosition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobposition.FieldName, jobposition.FieldDescription:
			values[i] = new(sql.NullString)
		case jobposition.FieldCreatedAt, jobposition.FieldUpdatedAt, jobposition.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case jobposition.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type JobPosition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobPosition fields.
func (jp *JobPosition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobposition.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				jp.ID = *value
			}
		case jobposition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				jp.CreatedAt = value.Time
			}
		case jobposition.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				jp.UpdatedAt = value.Time
			}
		case jobposition.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				jp.DeletedAt = value.Time
			}
		case jobposition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				jp.Name = value.String
			}
		case jobposition.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				jp.Description = value.String
			}
		}
	}
	return nil
}

// QueryHiringJobPositionEdges queries the "hiring_job_position_edges" edge of the JobPosition entity.
func (jp *JobPosition) QueryHiringJobPositionEdges() *HiringJobQuery {
	return (&JobPositionClient{config: jp.config}).QueryHiringJobPositionEdges(jp)
}

// Update returns a builder for updating this JobPosition.
// Note that you need to call JobPosition.Unwrap() before calling this method if this JobPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (jp *JobPosition) Update() *JobPositionUpdateOne {
	return (&JobPositionClient{config: jp.config}).UpdateOne(jp)
}

// Unwrap unwraps the JobPosition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jp *JobPosition) Unwrap() *JobPosition {
	_tx, ok := jp.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobPosition is not a transactional entity")
	}
	jp.config.driver = _tx.drv
	return jp
}

// String implements the fmt.Stringer.
func (jp *JobPosition) String() string {
	var builder strings.Builder
	builder.WriteString("JobPosition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(jp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(jp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(jp.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(jp.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(jp.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedHiringJobPositionEdges returns the HiringJobPositionEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (jp *JobPosition) NamedHiringJobPositionEdges(name string) ([]*HiringJob, error) {
	if jp.Edges.namedHiringJobPositionEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := jp.Edges.namedHiringJobPositionEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (jp *JobPosition) appendNamedHiringJobPositionEdges(name string, edges ...*HiringJob) {
	if jp.Edges.namedHiringJobPositionEdges == nil {
		jp.Edges.namedHiringJobPositionEdges = make(map[string][]*HiringJob)
	}
	if len(edges) == 0 {
		jp.Edges.namedHiringJobPositionEdges[name] = []*HiringJob{}
	} else {
		jp.Edges.namedHiringJobPositionEdges[name] = append(jp.Edges.namedHiringJobPositionEdges[name], edges...)
	}
}

// JobPositions is a parsable slice of JobPosition.
type JobPositions []*JobPosition

func (jp JobPositions) config(cfg config) {
	for _i := range jp {
		jp[_i].config = cfg
	}
}

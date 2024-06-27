// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/skill"
	"trec/ent/skilltype"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Skill is the model entity for the Skill schema.
type Skill struct {
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
	// SkillTypeID holds the value of the "skill_type_id" field.
	SkillTypeID uuid.UUID `json:"skill_type_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkillQuery when eager-loading is set.
	Edges SkillEdges `json:"edges"`
}

// SkillEdges holds the relations/edges for other nodes in the graph.
type SkillEdges struct {
	// SkillTypeEdge holds the value of the skill_type_edge edge.
	SkillTypeEdge *SkillType `json:"skill_type_edge,omitempty"`
	// EntitySkillEdges holds the value of the entity_skill_edges edge.
	EntitySkillEdges []*EntitySkill `json:"entity_skill_edges,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedEntitySkillEdges map[string][]*EntitySkill
}

// SkillTypeEdgeOrErr returns the SkillTypeEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkillEdges) SkillTypeEdgeOrErr() (*SkillType, error) {
	if e.loadedTypes[0] {
		if e.SkillTypeEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: skilltype.Label}
		}
		return e.SkillTypeEdge, nil
	}
	return nil, &NotLoadedError{edge: "skill_type_edge"}
}

// EntitySkillEdgesOrErr returns the EntitySkillEdges value or an error if the edge
// was not loaded in eager-loading.
func (e SkillEdges) EntitySkillEdgesOrErr() ([]*EntitySkill, error) {
	if e.loadedTypes[1] {
		return e.EntitySkillEdges, nil
	}
	return nil, &NotLoadedError{edge: "entity_skill_edges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Skill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case skill.FieldName, skill.FieldDescription:
			values[i] = new(sql.NullString)
		case skill.FieldCreatedAt, skill.FieldUpdatedAt, skill.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case skill.FieldID, skill.FieldSkillTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Skill", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Skill fields.
func (s *Skill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skill.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case skill.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case skill.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case skill.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case skill.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case skill.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case skill.FieldSkillTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field skill_type_id", values[i])
			} else if value != nil {
				s.SkillTypeID = *value
			}
		}
	}
	return nil
}

// QuerySkillTypeEdge queries the "skill_type_edge" edge of the Skill entity.
func (s *Skill) QuerySkillTypeEdge() *SkillTypeQuery {
	return (&SkillClient{config: s.config}).QuerySkillTypeEdge(s)
}

// QueryEntitySkillEdges queries the "entity_skill_edges" edge of the Skill entity.
func (s *Skill) QueryEntitySkillEdges() *EntitySkillQuery {
	return (&SkillClient{config: s.config}).QueryEntitySkillEdges(s)
}

// Update returns a builder for updating this Skill.
// Note that you need to call Skill.Unwrap() before calling this method if this Skill
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Skill) Update() *SkillUpdateOne {
	return (&SkillClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Skill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Skill) Unwrap() *Skill {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Skill is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Skill) String() string {
	var builder strings.Builder
	builder.WriteString("Skill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("skill_type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SkillTypeID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedEntitySkillEdges returns the EntitySkillEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Skill) NamedEntitySkillEdges(name string) ([]*EntitySkill, error) {
	if s.Edges.namedEntitySkillEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedEntitySkillEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Skill) appendNamedEntitySkillEdges(name string, edges ...*EntitySkill) {
	if s.Edges.namedEntitySkillEdges == nil {
		s.Edges.namedEntitySkillEdges = make(map[string][]*EntitySkill)
	}
	if len(edges) == 0 {
		s.Edges.namedEntitySkillEdges[name] = []*EntitySkill{}
	} else {
		s.Edges.namedEntitySkillEdges[name] = append(s.Edges.namedEntitySkillEdges[name], edges...)
	}
}

// Skills is a parsable slice of Skill.
type Skills []*Skill

func (s Skills) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}

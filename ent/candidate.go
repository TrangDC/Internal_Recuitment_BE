// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"trec/ent/candidate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Candidate is the model entity for the Candidate schema.
type Candidate struct {
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
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob time.Time `json:"dob,omitempty"`
	// IsBlacklist holds the value of the "is_blacklist" field.
	IsBlacklist bool `json:"is_blacklist,omitempty"`
	// LastApplyDate holds the value of the "last_apply_date" field.
	LastApplyDate time.Time `json:"last_apply_date,omitempty"`
	// ReferenceType holds the value of the "reference_type" field.
	ReferenceType candidate.ReferenceType `json:"reference_type,omitempty"`
	// ReferenceValue holds the value of the "reference_value" field.
	ReferenceValue string `json:"reference_value,omitempty"`
	// ReferenceUID holds the value of the "reference_uid" field.
	ReferenceUID uuid.UUID `json:"reference_uid,omitempty"`
	// RecruitTime holds the value of the "recruit_time" field.
	RecruitTime time.Time `json:"recruit_time,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar uuid.UUID `json:"avatar,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CandidateQuery when eager-loading is set.
	Edges CandidateEdges `json:"edges"`
}

// CandidateEdges holds the relations/edges for other nodes in the graph.
type CandidateEdges struct {
	// CandidateJobEdges holds the value of the candidate_job_edges edge.
	CandidateJobEdges []*CandidateJob `json:"candidate_job_edges,omitempty"`
	// ReferenceUserEdge holds the value of the reference_user_edge edge.
	ReferenceUserEdge *User `json:"reference_user_edge,omitempty"`
	// AttachmentEdges holds the value of the attachment_edges edge.
	AttachmentEdges []*Attachment `json:"attachment_edges,omitempty"`
	// CandidateSkillEdges holds the value of the candidate_skill_edges edge.
	CandidateSkillEdges []*EntitySkill `json:"candidate_skill_edges,omitempty"`
	// CandidateExpEdges holds the value of the candidate_exp_edges edge.
	CandidateExpEdges []*CandidateExp `json:"candidate_exp_edges,omitempty"`
	// CandidateEducateEdges holds the value of the candidate_educate_edges edge.
	CandidateEducateEdges []*CandidateEducate `json:"candidate_educate_edges,omitempty"`
	// CandidateAwardEdges holds the value of the candidate_award_edges edge.
	CandidateAwardEdges []*CandidateAward `json:"candidate_award_edges,omitempty"`
	// CandidateCertificateEdges holds the value of the candidate_certificate_edges edge.
	CandidateCertificateEdges []*CandidateCertificate `json:"candidate_certificate_edges,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedCandidateJobEdges         map[string][]*CandidateJob
	namedAttachmentEdges           map[string][]*Attachment
	namedCandidateSkillEdges       map[string][]*EntitySkill
	namedCandidateExpEdges         map[string][]*CandidateExp
	namedCandidateEducateEdges     map[string][]*CandidateEducate
	namedCandidateAwardEdges       map[string][]*CandidateAward
	namedCandidateCertificateEdges map[string][]*CandidateCertificate
}

// CandidateJobEdgesOrErr returns the CandidateJobEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateJobEdgesOrErr() ([]*CandidateJob, error) {
	if e.loadedTypes[0] {
		return e.CandidateJobEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_job_edges"}
}

// ReferenceUserEdgeOrErr returns the ReferenceUserEdge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CandidateEdges) ReferenceUserEdgeOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.ReferenceUserEdge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ReferenceUserEdge, nil
	}
	return nil, &NotLoadedError{edge: "reference_user_edge"}
}

// AttachmentEdgesOrErr returns the AttachmentEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) AttachmentEdgesOrErr() ([]*Attachment, error) {
	if e.loadedTypes[2] {
		return e.AttachmentEdges, nil
	}
	return nil, &NotLoadedError{edge: "attachment_edges"}
}

// CandidateSkillEdgesOrErr returns the CandidateSkillEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateSkillEdgesOrErr() ([]*EntitySkill, error) {
	if e.loadedTypes[3] {
		return e.CandidateSkillEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_skill_edges"}
}

// CandidateExpEdgesOrErr returns the CandidateExpEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateExpEdgesOrErr() ([]*CandidateExp, error) {
	if e.loadedTypes[4] {
		return e.CandidateExpEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_exp_edges"}
}

// CandidateEducateEdgesOrErr returns the CandidateEducateEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateEducateEdgesOrErr() ([]*CandidateEducate, error) {
	if e.loadedTypes[5] {
		return e.CandidateEducateEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_educate_edges"}
}

// CandidateAwardEdgesOrErr returns the CandidateAwardEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateAwardEdgesOrErr() ([]*CandidateAward, error) {
	if e.loadedTypes[6] {
		return e.CandidateAwardEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_award_edges"}
}

// CandidateCertificateEdgesOrErr returns the CandidateCertificateEdges value or an error if the edge
// was not loaded in eager-loading.
func (e CandidateEdges) CandidateCertificateEdgesOrErr() ([]*CandidateCertificate, error) {
	if e.loadedTypes[7] {
		return e.CandidateCertificateEdges, nil
	}
	return nil, &NotLoadedError{edge: "candidate_certificate_edges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Candidate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case candidate.FieldIsBlacklist:
			values[i] = new(sql.NullBool)
		case candidate.FieldName, candidate.FieldEmail, candidate.FieldPhone, candidate.FieldReferenceType, candidate.FieldReferenceValue, candidate.FieldDescription, candidate.FieldCountry, candidate.FieldAddress:
			values[i] = new(sql.NullString)
		case candidate.FieldCreatedAt, candidate.FieldUpdatedAt, candidate.FieldDeletedAt, candidate.FieldDob, candidate.FieldLastApplyDate, candidate.FieldRecruitTime:
			values[i] = new(sql.NullTime)
		case candidate.FieldID, candidate.FieldReferenceUID, candidate.FieldAvatar:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Candidate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Candidate fields.
func (c *Candidate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case candidate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case candidate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case candidate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case candidate.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case candidate.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case candidate.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case candidate.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case candidate.FieldDob:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				c.Dob = value.Time
			}
		case candidate.FieldIsBlacklist:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_blacklist", values[i])
			} else if value.Valid {
				c.IsBlacklist = value.Bool
			}
		case candidate.FieldLastApplyDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_apply_date", values[i])
			} else if value.Valid {
				c.LastApplyDate = value.Time
			}
		case candidate.FieldReferenceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_type", values[i])
			} else if value.Valid {
				c.ReferenceType = candidate.ReferenceType(value.String)
			}
		case candidate.FieldReferenceValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_value", values[i])
			} else if value.Valid {
				c.ReferenceValue = value.String
			}
		case candidate.FieldReferenceUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field reference_uid", values[i])
			} else if value != nil {
				c.ReferenceUID = *value
			}
		case candidate.FieldRecruitTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field recruit_time", values[i])
			} else if value.Valid {
				c.RecruitTime = value.Time
			}
		case candidate.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case candidate.FieldAvatar:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value != nil {
				c.Avatar = *value
			}
		case candidate.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				c.Country = value.String
			}
		case candidate.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		}
	}
	return nil
}

// QueryCandidateJobEdges queries the "candidate_job_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateJobEdges() *CandidateJobQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateJobEdges(c)
}

// QueryReferenceUserEdge queries the "reference_user_edge" edge of the Candidate entity.
func (c *Candidate) QueryReferenceUserEdge() *UserQuery {
	return (&CandidateClient{config: c.config}).QueryReferenceUserEdge(c)
}

// QueryAttachmentEdges queries the "attachment_edges" edge of the Candidate entity.
func (c *Candidate) QueryAttachmentEdges() *AttachmentQuery {
	return (&CandidateClient{config: c.config}).QueryAttachmentEdges(c)
}

// QueryCandidateSkillEdges queries the "candidate_skill_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateSkillEdges() *EntitySkillQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateSkillEdges(c)
}

// QueryCandidateExpEdges queries the "candidate_exp_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateExpEdges() *CandidateExpQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateExpEdges(c)
}

// QueryCandidateEducateEdges queries the "candidate_educate_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateEducateEdges() *CandidateEducateQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateEducateEdges(c)
}

// QueryCandidateAwardEdges queries the "candidate_award_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateAwardEdges() *CandidateAwardQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateAwardEdges(c)
}

// QueryCandidateCertificateEdges queries the "candidate_certificate_edges" edge of the Candidate entity.
func (c *Candidate) QueryCandidateCertificateEdges() *CandidateCertificateQuery {
	return (&CandidateClient{config: c.config}).QueryCandidateCertificateEdges(c)
}

// Update returns a builder for updating this Candidate.
// Note that you need to call Candidate.Unwrap() before calling this method if this Candidate
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Candidate) Update() *CandidateUpdateOne {
	return (&CandidateClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Candidate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Candidate) Unwrap() *Candidate {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Candidate is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Candidate) String() string {
	var builder strings.Builder
	builder.WriteString("Candidate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("dob=")
	builder.WriteString(c.Dob.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_blacklist=")
	builder.WriteString(fmt.Sprintf("%v", c.IsBlacklist))
	builder.WriteString(", ")
	builder.WriteString("last_apply_date=")
	builder.WriteString(c.LastApplyDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("reference_type=")
	builder.WriteString(fmt.Sprintf("%v", c.ReferenceType))
	builder.WriteString(", ")
	builder.WriteString("reference_value=")
	builder.WriteString(c.ReferenceValue)
	builder.WriteString(", ")
	builder.WriteString("reference_uid=")
	builder.WriteString(fmt.Sprintf("%v", c.ReferenceUID))
	builder.WriteString(", ")
	builder.WriteString("recruit_time=")
	builder.WriteString(c.RecruitTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(fmt.Sprintf("%v", c.Avatar))
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(c.Country)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCandidateJobEdges returns the CandidateJobEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateJobEdges(name string) ([]*CandidateJob, error) {
	if c.Edges.namedCandidateJobEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateJobEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateJobEdges(name string, edges ...*CandidateJob) {
	if c.Edges.namedCandidateJobEdges == nil {
		c.Edges.namedCandidateJobEdges = make(map[string][]*CandidateJob)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateJobEdges[name] = []*CandidateJob{}
	} else {
		c.Edges.namedCandidateJobEdges[name] = append(c.Edges.namedCandidateJobEdges[name], edges...)
	}
}

// NamedAttachmentEdges returns the AttachmentEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedAttachmentEdges(name string) ([]*Attachment, error) {
	if c.Edges.namedAttachmentEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAttachmentEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedAttachmentEdges(name string, edges ...*Attachment) {
	if c.Edges.namedAttachmentEdges == nil {
		c.Edges.namedAttachmentEdges = make(map[string][]*Attachment)
	}
	if len(edges) == 0 {
		c.Edges.namedAttachmentEdges[name] = []*Attachment{}
	} else {
		c.Edges.namedAttachmentEdges[name] = append(c.Edges.namedAttachmentEdges[name], edges...)
	}
}

// NamedCandidateSkillEdges returns the CandidateSkillEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateSkillEdges(name string) ([]*EntitySkill, error) {
	if c.Edges.namedCandidateSkillEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateSkillEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateSkillEdges(name string, edges ...*EntitySkill) {
	if c.Edges.namedCandidateSkillEdges == nil {
		c.Edges.namedCandidateSkillEdges = make(map[string][]*EntitySkill)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateSkillEdges[name] = []*EntitySkill{}
	} else {
		c.Edges.namedCandidateSkillEdges[name] = append(c.Edges.namedCandidateSkillEdges[name], edges...)
	}
}

// NamedCandidateExpEdges returns the CandidateExpEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateExpEdges(name string) ([]*CandidateExp, error) {
	if c.Edges.namedCandidateExpEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateExpEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateExpEdges(name string, edges ...*CandidateExp) {
	if c.Edges.namedCandidateExpEdges == nil {
		c.Edges.namedCandidateExpEdges = make(map[string][]*CandidateExp)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateExpEdges[name] = []*CandidateExp{}
	} else {
		c.Edges.namedCandidateExpEdges[name] = append(c.Edges.namedCandidateExpEdges[name], edges...)
	}
}

// NamedCandidateEducateEdges returns the CandidateEducateEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateEducateEdges(name string) ([]*CandidateEducate, error) {
	if c.Edges.namedCandidateEducateEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateEducateEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateEducateEdges(name string, edges ...*CandidateEducate) {
	if c.Edges.namedCandidateEducateEdges == nil {
		c.Edges.namedCandidateEducateEdges = make(map[string][]*CandidateEducate)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateEducateEdges[name] = []*CandidateEducate{}
	} else {
		c.Edges.namedCandidateEducateEdges[name] = append(c.Edges.namedCandidateEducateEdges[name], edges...)
	}
}

// NamedCandidateAwardEdges returns the CandidateAwardEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateAwardEdges(name string) ([]*CandidateAward, error) {
	if c.Edges.namedCandidateAwardEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateAwardEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateAwardEdges(name string, edges ...*CandidateAward) {
	if c.Edges.namedCandidateAwardEdges == nil {
		c.Edges.namedCandidateAwardEdges = make(map[string][]*CandidateAward)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateAwardEdges[name] = []*CandidateAward{}
	} else {
		c.Edges.namedCandidateAwardEdges[name] = append(c.Edges.namedCandidateAwardEdges[name], edges...)
	}
}

// NamedCandidateCertificateEdges returns the CandidateCertificateEdges named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Candidate) NamedCandidateCertificateEdges(name string) ([]*CandidateCertificate, error) {
	if c.Edges.namedCandidateCertificateEdges == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCandidateCertificateEdges[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Candidate) appendNamedCandidateCertificateEdges(name string, edges ...*CandidateCertificate) {
	if c.Edges.namedCandidateCertificateEdges == nil {
		c.Edges.namedCandidateCertificateEdges = make(map[string][]*CandidateCertificate)
	}
	if len(edges) == 0 {
		c.Edges.namedCandidateCertificateEdges[name] = []*CandidateCertificate{}
	} else {
		c.Edges.namedCandidateCertificateEdges[name] = append(c.Edges.namedCandidateCertificateEdges[name], edges...)
	}
}

// Candidates is a parsable slice of Candidate.
type Candidates []*Candidate

func (c Candidates) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

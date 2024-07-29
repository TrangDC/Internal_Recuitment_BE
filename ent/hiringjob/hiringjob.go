// Code generated by ent, DO NOT EDIT.

package hiringjob

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the hiringjob type in the database.
	Label = "hiring_job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldSalaryType holds the string denoting the salary_type field in the database.
	FieldSalaryType = "salary_type"
	// FieldSalaryFrom holds the string denoting the salary_from field in the database.
	FieldSalaryFrom = "salary_from"
	// FieldSalaryTo holds the string denoting the salary_to field in the database.
	FieldSalaryTo = "salary_to"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldLastApplyDate holds the string denoting the last_apply_date field in the database.
	FieldLastApplyDate = "last_apply_date"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldHiringTeamID holds the string denoting the hiring_team_id field in the database.
	FieldHiringTeamID = "hiring_team_id"
	// EdgeOwnerEdge holds the string denoting the owner_edge edge name in mutations.
	EdgeOwnerEdge = "owner_edge"
	// EdgeCandidateJobEdges holds the string denoting the candidate_job_edges edge name in mutations.
	EdgeCandidateJobEdges = "candidate_job_edges"
	// EdgeHiringJobSkillEdges holds the string denoting the hiring_job_skill_edges edge name in mutations.
	EdgeHiringJobSkillEdges = "hiring_job_skill_edges"
	// EdgeHiringTeamEdge holds the string denoting the hiring_team_edge edge name in mutations.
	EdgeHiringTeamEdge = "hiring_team_edge"
	// Table holds the table name of the hiringjob in the database.
	Table = "hiring_jobs"
	// OwnerEdgeTable is the table that holds the owner_edge relation/edge.
	OwnerEdgeTable = "hiring_jobs"
	// OwnerEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerEdgeInverseTable = "users"
	// OwnerEdgeColumn is the table column denoting the owner_edge relation/edge.
	OwnerEdgeColumn = "created_by"
	// CandidateJobEdgesTable is the table that holds the candidate_job_edges relation/edge.
	CandidateJobEdgesTable = "candidate_jobs"
	// CandidateJobEdgesInverseTable is the table name for the CandidateJob entity.
	// It exists in this package in order to avoid circular dependency with the "candidatejob" package.
	CandidateJobEdgesInverseTable = "candidate_jobs"
	// CandidateJobEdgesColumn is the table column denoting the candidate_job_edges relation/edge.
	CandidateJobEdgesColumn = "hiring_job_id"
	// HiringJobSkillEdgesTable is the table that holds the hiring_job_skill_edges relation/edge.
	HiringJobSkillEdgesTable = "entity_skills"
	// HiringJobSkillEdgesInverseTable is the table name for the EntitySkill entity.
	// It exists in this package in order to avoid circular dependency with the "entityskill" package.
	HiringJobSkillEdgesInverseTable = "entity_skills"
	// HiringJobSkillEdgesColumn is the table column denoting the hiring_job_skill_edges relation/edge.
	HiringJobSkillEdgesColumn = "entity_id"
	// HiringTeamEdgeTable is the table that holds the hiring_team_edge relation/edge.
	HiringTeamEdgeTable = "hiring_jobs"
	// HiringTeamEdgeInverseTable is the table name for the HiringTeam entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteam" package.
	HiringTeamEdgeInverseTable = "hiring_teams"
	// HiringTeamEdgeColumn is the table column denoting the hiring_team_edge relation/edge.
	HiringTeamEdgeColumn = "hiring_team_id"
)

// Columns holds all SQL columns for hiringjob fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSlug,
	FieldName,
	FieldDescription,
	FieldAmount,
	FieldStatus,
	FieldCreatedBy,
	FieldLocation,
	FieldSalaryType,
	FieldSalaryFrom,
	FieldSalaryTo,
	FieldCurrency,
	FieldLastApplyDate,
	FieldPriority,
	FieldHiringTeamID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int
	// DefaultSalaryFrom holds the default value on creation for the "salary_from" field.
	DefaultSalaryFrom int
	// DefaultSalaryTo holds the default value on creation for the "salary_to" field.
	DefaultSalaryTo int
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
)

// Status defines the type for the "status" enum field.
type Status string

// StatusOpened is the default value of the Status enum.
const DefaultStatus = StatusOpened

// Status values.
const (
	StatusDraft  Status = "draft"
	StatusOpened Status = "opened"
	StatusClosed Status = "closed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDraft, StatusOpened, StatusClosed:
		return nil
	default:
		return fmt.Errorf("hiringjob: invalid enum value for status field: %q", s)
	}
}

// Location defines the type for the "location" enum field.
type Location string

// Location values.
const (
	LocationHaNoi     Location = "ha_noi"
	LocationHoChiMinh Location = "ho_chi_minh"
	LocationDaNang    Location = "da_nang"
	LocationJapan     Location = "japan"
	LocationSingapore Location = "singapore"
)

func (l Location) String() string {
	return string(l)
}

// LocationValidator is a validator for the "location" field enum values. It is called by the builders before save.
func LocationValidator(l Location) error {
	switch l {
	case LocationHaNoi, LocationHoChiMinh, LocationDaNang, LocationJapan, LocationSingapore:
		return nil
	default:
		return fmt.Errorf("hiringjob: invalid enum value for location field: %q", l)
	}
}

// SalaryType defines the type for the "salary_type" enum field.
type SalaryType string

// SalaryType values.
const (
	SalaryTypeRange     SalaryType = "range"
	SalaryTypeUpTo      SalaryType = "up_to"
	SalaryTypeNegotiate SalaryType = "negotiate"
	SalaryTypeMinimum   SalaryType = "minimum"
)

func (st SalaryType) String() string {
	return string(st)
}

// SalaryTypeValidator is a validator for the "salary_type" field enum values. It is called by the builders before save.
func SalaryTypeValidator(st SalaryType) error {
	switch st {
	case SalaryTypeRange, SalaryTypeUpTo, SalaryTypeNegotiate, SalaryTypeMinimum:
		return nil
	default:
		return fmt.Errorf("hiringjob: invalid enum value for salary_type field: %q", st)
	}
}

// Currency defines the type for the "currency" enum field.
type Currency string

// Currency values.
const (
	CurrencyVnd Currency = "vnd"
	CurrencyUsd Currency = "usd"
	CurrencyJpy Currency = "jpy"
)

func (c Currency) String() string {
	return string(c)
}

// CurrencyValidator is a validator for the "currency" field enum values. It is called by the builders before save.
func CurrencyValidator(c Currency) error {
	switch c {
	case CurrencyVnd, CurrencyUsd, CurrencyJpy:
		return nil
	default:
		return fmt.Errorf("hiringjob: invalid enum value for currency field: %q", c)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Location) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Location) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Location(str)
	if err := LocationValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Location", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e SalaryType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *SalaryType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = SalaryType(str)
	if err := SalaryTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid SalaryType", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Currency) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Currency) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Currency(str)
	if err := CurrencyValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Currency", str)
	}
	return nil
}

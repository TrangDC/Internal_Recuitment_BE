// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorkEmail holds the string denoting the work_email field in the database.
	FieldWorkEmail = "work_email"
	// FieldOid holds the string denoting the oid field in the database.
	FieldOid = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeAuditEdge holds the string denoting the audit_edge edge name in mutations.
	EdgeAuditEdge = "audit_edge"
	// EdgeTeamEdges holds the string denoting the team_edges edge name in mutations.
	EdgeTeamEdges = "team_edges"
	// EdgeTeamUsers holds the string denoting the team_users edge name in mutations.
	EdgeTeamUsers = "team_users"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AuditEdgeTable is the table that holds the audit_edge relation/edge.
	AuditEdgeTable = "audit_trails"
	// AuditEdgeInverseTable is the table name for the AuditTrail entity.
	// It exists in this package in order to avoid circular dependency with the "audittrail" package.
	AuditEdgeInverseTable = "audit_trails"
	// AuditEdgeColumn is the table column denoting the audit_edge relation/edge.
	AuditEdgeColumn = "created_by"
	// TeamEdgesTable is the table that holds the team_edges relation/edge. The primary key declared below.
	TeamEdgesTable = "team_managers"
	// TeamEdgesInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamEdgesInverseTable = "teams"
	// TeamUsersTable is the table that holds the team_users relation/edge.
	TeamUsersTable = "team_managers"
	// TeamUsersInverseTable is the table name for the TeamManager entity.
	// It exists in this package in order to avoid circular dependency with the "teammanager" package.
	TeamUsersInverseTable = "team_managers"
	// TeamUsersColumn is the table column denoting the team_users relation/edge.
	TeamUsersColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWorkEmail,
	FieldOid,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	// TeamEdgesPrimaryKey and TeamEdgesColumn2 are the table columns denoting the
	// primary key for the team_edges relation (M2M).
	TeamEdgesPrimaryKey = []string{"user_id", "team_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// WorkEmailValidator is a validator for the "work_email" field. It is called by the builders before save.
	WorkEmailValidator func(string) error
	// OidValidator is a validator for the "oid" field. It is called by the builders before save.
	OidValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

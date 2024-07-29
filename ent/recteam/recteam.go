// Code generated by ent, DO NOT EDIT.

package recteam

import (
	"time"
)

const (
	// Label holds the string label denoting the recteam type in the database.
	Label = "rec_team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLeaderID holds the string denoting the leader_id field in the database.
	FieldLeaderID = "leader_id"
	// EdgeRecLeaderEdge holds the string denoting the rec_leader_edge edge name in mutations.
	EdgeRecLeaderEdge = "rec_leader_edge"
	// EdgeRecMemberEdges holds the string denoting the rec_member_edges edge name in mutations.
	EdgeRecMemberEdges = "rec_member_edges"
	// Table holds the table name of the recteam in the database.
	Table = "rec_teams"
	// RecLeaderEdgeTable is the table that holds the rec_leader_edge relation/edge.
	RecLeaderEdgeTable = "rec_teams"
	// RecLeaderEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecLeaderEdgeInverseTable = "users"
	// RecLeaderEdgeColumn is the table column denoting the rec_leader_edge relation/edge.
	RecLeaderEdgeColumn = "leader_id"
	// RecMemberEdgesTable is the table that holds the rec_member_edges relation/edge.
	RecMemberEdgesTable = "users"
	// RecMemberEdgesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RecMemberEdgesInverseTable = "users"
	// RecMemberEdgesColumn is the table column denoting the rec_member_edges relation/edge.
	RecMemberEdgesColumn = "rec_team_id"
)

// Columns holds all SQL columns for recteam fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDescription,
	FieldLeaderID,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

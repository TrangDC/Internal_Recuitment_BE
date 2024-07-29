// Code generated by ent, DO NOT EDIT.

package hiringteammanager

import (
	"time"
)

const (
	// Label holds the string label denoting the hiringteammanager type in the database.
	Label = "hiring_team_manager"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldHiringTeamID holds the string denoting the hiring_team_id field in the database.
	FieldHiringTeamID = "hiring_team_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUserEdge holds the string denoting the user_edge edge name in mutations.
	EdgeUserEdge = "user_edge"
	// EdgeHiringTeamEdge holds the string denoting the hiring_team_edge edge name in mutations.
	EdgeHiringTeamEdge = "hiring_team_edge"
	// Table holds the table name of the hiringteammanager in the database.
	Table = "hiring_team_managers"
	// UserEdgeTable is the table that holds the user_edge relation/edge.
	UserEdgeTable = "hiring_team_managers"
	// UserEdgeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserEdgeInverseTable = "users"
	// UserEdgeColumn is the table column denoting the user_edge relation/edge.
	UserEdgeColumn = "user_id"
	// HiringTeamEdgeTable is the table that holds the hiring_team_edge relation/edge.
	HiringTeamEdgeTable = "hiring_team_managers"
	// HiringTeamEdgeInverseTable is the table name for the HiringTeam entity.
	// It exists in this package in order to avoid circular dependency with the "hiringteam" package.
	HiringTeamEdgeInverseTable = "hiring_teams"
	// HiringTeamEdgeColumn is the table column denoting the hiring_team_edge relation/edge.
	HiringTeamEdgeColumn = "hiring_team_id"
)

// Columns holds all SQL columns for hiringteammanager fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHiringTeamID,
	FieldUserID,
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
)

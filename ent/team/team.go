// Code generated by ent, DO NOT EDIT.

package team

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
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
	// EdgeUserEdges holds the string denoting the user_edges edge name in mutations.
	EdgeUserEdges = "user_edges"
	// EdgeHiringTeam holds the string denoting the hiring_team edge name in mutations.
	EdgeHiringTeam = "hiring_team"
	// EdgeUserTeams holds the string denoting the user_teams edge name in mutations.
	EdgeUserTeams = "user_teams"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// UserEdgesTable is the table that holds the user_edges relation/edge. The primary key declared below.
	UserEdgesTable = "team_managers"
	// UserEdgesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserEdgesInverseTable = "users"
	// HiringTeamTable is the table that holds the hiring_team relation/edge.
	HiringTeamTable = "hiring_jobs"
	// HiringTeamInverseTable is the table name for the HiringJob entity.
	// It exists in this package in order to avoid circular dependency with the "hiringjob" package.
	HiringTeamInverseTable = "hiring_jobs"
	// HiringTeamColumn is the table column denoting the hiring_team relation/edge.
	HiringTeamColumn = "team_id"
	// UserTeamsTable is the table that holds the user_teams relation/edge.
	UserTeamsTable = "team_managers"
	// UserTeamsInverseTable is the table name for the TeamManager entity.
	// It exists in this package in order to avoid circular dependency with the "teammanager" package.
	UserTeamsInverseTable = "team_managers"
	// UserTeamsColumn is the table column denoting the user_teams relation/edge.
	UserTeamsColumn = "team_id"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSlug,
	FieldName,
}

var (
	// UserEdgesPrimaryKey and UserEdgesColumn2 are the table columns denoting the
	// primary key for the user_edges relation (M2M).
	UserEdgesPrimaryKey = []string{"user_id", "team_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

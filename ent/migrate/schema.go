// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuditTrailsColumns holds the columns for the "audit_trails" table.
	AuditTrailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "record_id", Type: field.TypeUUID},
		{Name: "module", Type: field.TypeEnum, Enums: []string{"teams"}},
		{Name: "action_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"create", "update", "delete"}, Default: "create"},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "record_changes", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
	}
	// AuditTrailsTable holds the schema information for the "audit_trails" table.
	AuditTrailsTable = &schema.Table{
		Name:       "audit_trails",
		Columns:    AuditTrailsColumns,
		PrimaryKey: []*schema.Column{AuditTrailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "audit_trails_users_audit_edge",
				Columns:    []*schema.Column{AuditTrailsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// TeamManagersColumns holds the columns for the "team_managers" table.
	TeamManagersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "team_id", Type: field.TypeUUID},
	}
	// TeamManagersTable holds the schema information for the "team_managers" table.
	TeamManagersTable = &schema.Table{
		Name:       "team_managers",
		Columns:    TeamManagersColumns,
		PrimaryKey: []*schema.Column{TeamManagersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "team_managers_users_user_edge",
				Columns:    []*schema.Column{TeamManagersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "team_managers_teams_team_edge",
				Columns:    []*schema.Column{TeamManagersColumns[5]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "teammanager_user_id_team_id",
				Unique:  true,
				Columns: []*schema.Column{TeamManagersColumns[4], TeamManagersColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "work_email", Type: field.TypeString, Size: 255},
		{Name: "oid", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuditTrailsTable,
		TeamsTable,
		TeamManagersTable,
		UsersTable,
	}
)

func init() {
	AuditTrailsTable.ForeignKeys[0].RefTable = UsersTable
	TeamManagersTable.ForeignKeys[0].RefTable = UsersTable
	TeamManagersTable.ForeignKeys[1].RefTable = TeamsTable
}

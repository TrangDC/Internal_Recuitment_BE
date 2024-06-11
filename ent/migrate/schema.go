// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "document_id", Type: field.TypeUUID, Unique: true},
		{Name: "document_name", Type: field.TypeString, Size: 255},
		{Name: "relation_type", Type: field.TypeEnum, Enums: []string{"candidate_jobs", "candidate_job_feedbacks"}},
		{Name: "relation_id", Type: field.TypeUUID, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_candidate_interviews_attachment_edges",
				Columns:    []*schema.Column{AttachmentsColumns[7]},
				RefColumns: []*schema.Column{CandidateInterviewsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "attachments_candidate_jobs_attachment_edges",
				Columns:    []*schema.Column{AttachmentsColumns[7]},
				RefColumns: []*schema.Column{CandidateJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "attachments_candidate_job_feedbacks_attachment_edges",
				Columns:    []*schema.Column{AttachmentsColumns[7]},
				RefColumns: []*schema.Column{CandidateJobFeedbacksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AuditTrailsColumns holds the columns for the "audit_trails" table.
	AuditTrailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "record_id", Type: field.TypeUUID},
		{Name: "module", Type: field.TypeEnum, Enums: []string{"teams", "hiring_jobs", "candidates"}},
		{Name: "action_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"create", "update", "delete"}, Default: "create"},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "record_changes", Type: field.TypeString, Nullable: true, Size: 2147483647},
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
	// CandidatesColumns holds the columns for the "candidates" table.
	CandidatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "phone", Type: field.TypeString, Size: 255},
		{Name: "dob", Type: field.TypeTime, Nullable: true},
		{Name: "is_blacklist", Type: field.TypeBool, Default: false},
		{Name: "last_apply_date", Type: field.TypeTime, Nullable: true},
	}
	// CandidatesTable holds the schema information for the "candidates" table.
	CandidatesTable = &schema.Table{
		Name:       "candidates",
		Columns:    CandidatesColumns,
		PrimaryKey: []*schema.Column{CandidatesColumns[0]},
	}
	// CandidateInterviewsColumns holds the columns for the "candidate_interviews" table.
	CandidateInterviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "candidate_job_status", Type: field.TypeEnum, Enums: []string{"applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff"}, Default: "applied"},
		{Name: "interview_date", Type: field.TypeTime, Nullable: true},
		{Name: "start_from", Type: field.TypeTime, Nullable: true},
		{Name: "end_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "candidate_job_id", Type: field.TypeUUID, Nullable: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
	}
	// CandidateInterviewsTable holds the schema information for the "candidate_interviews" table.
	CandidateInterviewsTable = &schema.Table{
		Name:       "candidate_interviews",
		Columns:    CandidateInterviewsColumns,
		PrimaryKey: []*schema.Column{CandidateInterviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidate_interviews_candidate_jobs_candidate_job_interview",
				Columns:    []*schema.Column{CandidateInterviewsColumns[10]},
				RefColumns: []*schema.Column{CandidateJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_interviews_users_candidate_interview_edges",
				Columns:    []*schema.Column{CandidateInterviewsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CandidateInterviewersColumns holds the columns for the "candidate_interviewers" table.
	CandidateInterviewersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "candidate_interview_id", Type: field.TypeUUID},
	}
	// CandidateInterviewersTable holds the schema information for the "candidate_interviewers" table.
	CandidateInterviewersTable = &schema.Table{
		Name:       "candidate_interviewers",
		Columns:    CandidateInterviewersColumns,
		PrimaryKey: []*schema.Column{CandidateInterviewersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidate_interviewers_users_user_edge",
				Columns:    []*schema.Column{CandidateInterviewersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "candidate_interviewers_candidate_interviews_interview_edge",
				Columns:    []*schema.Column{CandidateInterviewersColumns[5]},
				RefColumns: []*schema.Column{CandidateInterviewsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "candidateinterviewer_user_id_candidate_interview_id",
				Unique:  true,
				Columns: []*schema.Column{CandidateInterviewersColumns[4], CandidateInterviewersColumns[5]},
			},
		},
	}
	// CandidateJobsColumns holds the columns for the "candidate_jobs" table.
	CandidateJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff"}, Default: "applied"},
		{Name: "failed_reason", Type: field.TypeJSON, Nullable: true},
		{Name: "candidate_id", Type: field.TypeUUID, Nullable: true},
		{Name: "hiring_job_id", Type: field.TypeUUID, Nullable: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
	}
	// CandidateJobsTable holds the schema information for the "candidate_jobs" table.
	CandidateJobsTable = &schema.Table{
		Name:       "candidate_jobs",
		Columns:    CandidateJobsColumns,
		PrimaryKey: []*schema.Column{CandidateJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidate_jobs_candidates_candidate_job_edges",
				Columns:    []*schema.Column{CandidateJobsColumns[6]},
				RefColumns: []*schema.Column{CandidatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_jobs_hiring_jobs_candidate_job_edges",
				Columns:    []*schema.Column{CandidateJobsColumns[7]},
				RefColumns: []*schema.Column{HiringJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_jobs_users_candidate_job_edges",
				Columns:    []*schema.Column{CandidateJobsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CandidateJobFeedbacksColumns holds the columns for the "candidate_job_feedbacks" table.
	CandidateJobFeedbacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "candidate_job_status", Type: field.TypeEnum, Enums: []string{"applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff"}, Default: "applied"},
		{Name: "feedback", Type: field.TypeString, Size: 2147483647},
		{Name: "candidate_job_id", Type: field.TypeUUID, Nullable: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
	}
	// CandidateJobFeedbacksTable holds the schema information for the "candidate_job_feedbacks" table.
	CandidateJobFeedbacksTable = &schema.Table{
		Name:       "candidate_job_feedbacks",
		Columns:    CandidateJobFeedbacksColumns,
		PrimaryKey: []*schema.Column{CandidateJobFeedbacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidate_job_feedbacks_candidate_jobs_candidate_job_feedback",
				Columns:    []*schema.Column{CandidateJobFeedbacksColumns[6]},
				RefColumns: []*schema.Column{CandidateJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_job_feedbacks_users_candidate_job_feedback",
				Columns:    []*schema.Column{CandidateJobFeedbacksColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CandidateJobStepsColumns holds the columns for the "candidate_job_steps" table.
	CandidateJobStepsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "candidate_job_status", Type: field.TypeEnum, Enums: []string{"applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff"}, Default: "applied"},
		{Name: "candidate_job_id", Type: field.TypeUUID, Nullable: true},
	}
	// CandidateJobStepsTable holds the schema information for the "candidate_job_steps" table.
	CandidateJobStepsTable = &schema.Table{
		Name:       "candidate_job_steps",
		Columns:    CandidateJobStepsColumns,
		PrimaryKey: []*schema.Column{CandidateJobStepsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidate_job_steps_candidate_jobs_candidate_job_step",
				Columns:    []*schema.Column{CandidateJobStepsColumns[5]},
				RefColumns: []*schema.Column{CandidateJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HiringJobsColumns holds the columns for the "hiring_jobs" table.
	HiringJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "amount", Type: field.TypeInt, Default: 0},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "opened", "closed"}, Default: "opened"},
		{Name: "location", Type: field.TypeEnum, Enums: []string{"ha_noi", "ho_chi_minh", "da_nang", "japan"}},
		{Name: "salary_type", Type: field.TypeEnum, Enums: []string{"range", "up_to", "negotiate", "minimum"}},
		{Name: "salary_from", Type: field.TypeInt, Default: 0},
		{Name: "salary_to", Type: field.TypeInt, Default: 0},
		{Name: "currency", Type: field.TypeEnum, Enums: []string{"vnd", "usd", "jpy"}},
		{Name: "last_apply_date", Type: field.TypeTime, Nullable: true},
		{Name: "priority", Type: field.TypeInt, Default: 4},
		{Name: "team_id", Type: field.TypeUUID, Nullable: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
	}
	// HiringJobsTable holds the schema information for the "hiring_jobs" table.
	HiringJobsTable = &schema.Table{
		Name:       "hiring_jobs",
		Columns:    HiringJobsColumns,
		PrimaryKey: []*schema.Column{HiringJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hiring_jobs_teams_team_job_edges",
				Columns:    []*schema.Column{HiringJobsColumns[16]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "hiring_jobs_users_hiring_owner",
				Columns:    []*schema.Column{HiringJobsColumns[17]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "name", Type: field.TypeString, Size: 255},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// TeamManagersColumns holds the columns for the "team_managers" table.
	TeamManagersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
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
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "work_email", Type: field.TypeString, Size: 255},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "inactive"}, Default: "active"},
		{Name: "oid", Type: field.TypeString, Unique: true, Size: 255},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentsTable,
		AuditTrailsTable,
		CandidatesTable,
		CandidateInterviewsTable,
		CandidateInterviewersTable,
		CandidateJobsTable,
		CandidateJobFeedbacksTable,
		CandidateJobStepsTable,
		HiringJobsTable,
		SkillsTable,
		TeamsTable,
		TeamManagersTable,
		UsersTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = CandidateInterviewsTable
	AttachmentsTable.ForeignKeys[1].RefTable = CandidateJobsTable
	AttachmentsTable.ForeignKeys[2].RefTable = CandidateJobFeedbacksTable
	AuditTrailsTable.ForeignKeys[0].RefTable = UsersTable
	CandidateInterviewsTable.ForeignKeys[0].RefTable = CandidateJobsTable
	CandidateInterviewsTable.ForeignKeys[1].RefTable = UsersTable
	CandidateInterviewersTable.ForeignKeys[0].RefTable = UsersTable
	CandidateInterviewersTable.ForeignKeys[1].RefTable = CandidateInterviewsTable
	CandidateJobsTable.ForeignKeys[0].RefTable = CandidatesTable
	CandidateJobsTable.ForeignKeys[1].RefTable = HiringJobsTable
	CandidateJobsTable.ForeignKeys[2].RefTable = UsersTable
	CandidateJobFeedbacksTable.ForeignKeys[0].RefTable = CandidateJobsTable
	CandidateJobFeedbacksTable.ForeignKeys[1].RefTable = UsersTable
	CandidateJobStepsTable.ForeignKeys[0].RefTable = CandidateJobsTable
	HiringJobsTable.ForeignKeys[0].RefTable = TeamsTable
	HiringJobsTable.ForeignKeys[1].RefTable = UsersTable
	TeamManagersTable.ForeignKeys[0].RefTable = UsersTable
	TeamManagersTable.ForeignKeys[1].RefTable = TeamsTable
}

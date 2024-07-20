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
		{Name: "relation_type", Type: field.TypeEnum, Enums: []string{"candidate_jobs", "candidate_job_feedbacks", "candidates"}},
		{Name: "relation_id", Type: field.TypeUUID, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_candidates_attachment_edges",
				Columns:    []*schema.Column{AttachmentsColumns[7]},
				RefColumns: []*schema.Column{CandidatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
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
		{Name: "module", Type: field.TypeEnum, Enums: []string{"teams", "hiring_jobs", "candidates", "skills", "users", "skill_types", "roles", "email_templates"}},
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
		{Name: "reference_type", Type: field.TypeEnum, Enums: []string{"eb", "rec", "hiring_platform", "reference", "headhunt"}, Default: "eb"},
		{Name: "reference_value", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "recruit_time", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "reference_uid", Type: field.TypeUUID, Nullable: true},
	}
	// CandidatesTable holds the schema information for the "candidates" table.
	CandidatesTable = &schema.Table{
		Name:       "candidates",
		Columns:    CandidatesColumns,
		PrimaryKey: []*schema.Column{CandidatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidates_users_candidate_reference_edges",
				Columns:    []*schema.Column{CandidatesColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
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
		{Name: "status", Type: field.TypeEnum, Enums: []string{"invited_to_interview", "interviewing", "done", "cancelled"}, Default: "invited_to_interview"},
		{Name: "location", Type: field.TypeString, Size: 512},
		{Name: "meeting_link", Type: field.TypeString, Size: 2147483647},
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
				Columns:    []*schema.Column{CandidateInterviewsColumns[13]},
				RefColumns: []*schema.Column{CandidateJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_interviews_users_candidate_interview_edges",
				Columns:    []*schema.Column{CandidateInterviewsColumns[14]},
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
		{Name: "onboard_date", Type: field.TypeTime, Nullable: true},
		{Name: "offer_expiration_date", Type: field.TypeTime, Nullable: true},
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
				Columns:    []*schema.Column{CandidateJobsColumns[8]},
				RefColumns: []*schema.Column{CandidatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_jobs_hiring_jobs_candidate_job_edges",
				Columns:    []*schema.Column{CandidateJobsColumns[9]},
				RefColumns: []*schema.Column{HiringJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "candidate_jobs_users_candidate_job_edges",
				Columns:    []*schema.Column{CandidateJobsColumns[10]},
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
	// EmailRoleAttributesColumns holds the columns for the "email_role_attributes" table.
	EmailRoleAttributesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email_template_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// EmailRoleAttributesTable holds the schema information for the "email_role_attributes" table.
	EmailRoleAttributesTable = &schema.Table{
		Name:       "email_role_attributes",
		Columns:    EmailRoleAttributesColumns,
		PrimaryKey: []*schema.Column{EmailRoleAttributesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "email_role_attributes_teams_email_template_edge",
				Columns:    []*schema.Column{EmailRoleAttributesColumns[4]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "email_role_attributes_roles_role_edge",
				Columns:    []*schema.Column{EmailRoleAttributesColumns[5]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "emailroleattribute_role_id_email_template_id",
				Unique:  true,
				Columns: []*schema.Column{EmailRoleAttributesColumns[5], EmailRoleAttributesColumns[4]},
			},
		},
	}
	// EmailTemplatesColumns holds the columns for the "email_templates" table.
	EmailTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "event", Type: field.TypeEnum, Enums: []string{"candidate_applied_to_kiv", "candidate_interviewing_to_kiv", "candidate_interviewing_to_offering", "created_interview", "updating_interview", "cancel_interview"}},
		{Name: "send_to", Type: field.TypeJSON},
		{Name: "cc", Type: field.TypeJSON},
		{Name: "bcc", Type: field.TypeJSON},
		{Name: "subject", Type: field.TypeString, Size: 255},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "signature", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "inactive"}, Default: "active"},
	}
	// EmailTemplatesTable holds the schema information for the "email_templates" table.
	EmailTemplatesTable = &schema.Table{
		Name:       "email_templates",
		Columns:    EmailTemplatesColumns,
		PrimaryKey: []*schema.Column{EmailTemplatesColumns[0]},
	}
	// EntityPermissionsColumns holds the columns for the "entity_permissions" table.
	EntityPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "for_owner", Type: field.TypeBool, Default: false},
		{Name: "for_team", Type: field.TypeBool, Default: false},
		{Name: "for_all", Type: field.TypeBool, Default: false},
		{Name: "entity_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"user", "role"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "permission_id", Type: field.TypeUUID, Nullable: true},
		{Name: "entity_id", Type: field.TypeUUID, Nullable: true},
	}
	// EntityPermissionsTable holds the schema information for the "entity_permissions" table.
	EntityPermissionsTable = &schema.Table{
		Name:       "entity_permissions",
		Columns:    EntityPermissionsColumns,
		PrimaryKey: []*schema.Column{EntityPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entity_permissions_permissions_user_permission_edge",
				Columns:    []*schema.Column{EntityPermissionsColumns[7]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "entity_permissions_roles_role_permission_edges",
				Columns:    []*schema.Column{EntityPermissionsColumns[8]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "entity_permissions_users_user_permission_edges",
				Columns:    []*schema.Column{EntityPermissionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EntitySkillsColumns holds the columns for the "entity_skills" table.
	EntitySkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "entity_type", Type: field.TypeEnum, Enums: []string{"candidate", "hiring_job"}},
		{Name: "order_id", Type: field.TypeInt, Nullable: true},
		{Name: "entity_id", Type: field.TypeUUID, Nullable: true},
		{Name: "skill_id", Type: field.TypeUUID, Nullable: true},
	}
	// EntitySkillsTable holds the schema information for the "entity_skills" table.
	EntitySkillsTable = &schema.Table{
		Name:       "entity_skills",
		Columns:    EntitySkillsColumns,
		PrimaryKey: []*schema.Column{EntitySkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entity_skills_candidates_candidate_skill_edges",
				Columns:    []*schema.Column{EntitySkillsColumns[6]},
				RefColumns: []*schema.Column{CandidatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "entity_skills_hiring_jobs_hiring_job_skill_edges",
				Columns:    []*schema.Column{EntitySkillsColumns[6]},
				RefColumns: []*schema.Column{HiringJobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "entity_skills_skills_entity_skill_edges",
				Columns:    []*schema.Column{EntitySkillsColumns[7]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
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
		{Name: "location", Type: field.TypeEnum, Enums: []string{"ha_noi", "ho_chi_minh", "da_nang", "japan", "singapore"}},
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
	// OutgoingEmailsColumns holds the columns for the "outgoing_emails" table.
	OutgoingEmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "to", Type: field.TypeJSON},
		{Name: "cc", Type: field.TypeJSON},
		{Name: "bcc", Type: field.TypeJSON},
		{Name: "subject", Type: field.TypeString, Size: 2147483647},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "signature", Type: field.TypeString, Size: 2147483647},
		{Name: "email_template_id", Type: field.TypeUUID, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "sent", "failed"}, Default: "pending"},
	}
	// OutgoingEmailsTable holds the schema information for the "outgoing_emails" table.
	OutgoingEmailsTable = &schema.Table{
		Name:       "outgoing_emails",
		Columns:    OutgoingEmailsColumns,
		PrimaryKey: []*schema.Column{OutgoingEmailsColumns[0]},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "for_owner", Type: field.TypeBool, Default: false},
		{Name: "for_team", Type: field.TypeBool, Default: false},
		{Name: "for_all", Type: field.TypeBool, Default: false},
		{Name: "operation_name", Type: field.TypeString, Nullable: true},
		{Name: "parent_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeUUID, Nullable: true},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permissions_permission_groups_permission_edges",
				Columns:    []*schema.Column{PermissionsColumns[11]},
				RefColumns: []*schema.Column{PermissionGroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PermissionGroupsColumns holds the columns for the "permission_groups" table.
	PermissionGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "group_type", Type: field.TypeEnum, Enums: []string{"function", "system"}, Default: "function"},
		{Name: "order_id", Type: field.TypeInt},
		{Name: "parent_id", Type: field.TypeUUID, Nullable: true},
	}
	// PermissionGroupsTable holds the schema information for the "permission_groups" table.
	PermissionGroupsTable = &schema.Table{
		Name:       "permission_groups",
		Columns:    PermissionGroupsColumns,
		PrimaryKey: []*schema.Column{PermissionGroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permission_groups_permission_groups_group_permission_children",
				Columns:    []*schema.Column{PermissionGroupsColumns[7]},
				RefColumns: []*schema.Column{PermissionGroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "skill_type_id", Type: field.TypeUUID, Nullable: true},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "skills_skill_types_skill_edges",
				Columns:    []*schema.Column{SkillsColumns[6]},
				RefColumns: []*schema.Column{SkillTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillTypesColumns holds the columns for the "skill_types" table.
	SkillTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 255},
	}
	// SkillTypesTable holds the schema information for the "skill_types" table.
	SkillTypesTable = &schema.Table{
		Name:       "skill_types",
		Columns:    SkillTypesColumns,
		PrimaryKey: []*schema.Column{SkillTypesColumns[0]},
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
		{Name: "team_id", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_teams_member_edges",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_users_user_edge",
				Columns:    []*schema.Column{UserRolesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_roles_roles_role_edge",
				Columns:    []*schema.Column{UserRolesColumns[5]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userrole_user_id_role_id",
				Unique:  true,
				Columns: []*schema.Column{UserRolesColumns[4], UserRolesColumns[5]},
			},
		},
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
		EmailRoleAttributesTable,
		EmailTemplatesTable,
		EntityPermissionsTable,
		EntitySkillsTable,
		HiringJobsTable,
		OutgoingEmailsTable,
		PermissionsTable,
		PermissionGroupsTable,
		RolesTable,
		SkillsTable,
		SkillTypesTable,
		TeamsTable,
		TeamManagersTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = CandidatesTable
	AttachmentsTable.ForeignKeys[1].RefTable = CandidateInterviewsTable
	AttachmentsTable.ForeignKeys[2].RefTable = CandidateJobsTable
	AttachmentsTable.ForeignKeys[3].RefTable = CandidateJobFeedbacksTable
	AuditTrailsTable.ForeignKeys[0].RefTable = UsersTable
	CandidatesTable.ForeignKeys[0].RefTable = UsersTable
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
	EmailRoleAttributesTable.ForeignKeys[0].RefTable = TeamsTable
	EmailRoleAttributesTable.ForeignKeys[1].RefTable = RolesTable
	EntityPermissionsTable.ForeignKeys[0].RefTable = PermissionsTable
	EntityPermissionsTable.ForeignKeys[1].RefTable = RolesTable
	EntityPermissionsTable.ForeignKeys[2].RefTable = UsersTable
	EntitySkillsTable.ForeignKeys[0].RefTable = CandidatesTable
	EntitySkillsTable.ForeignKeys[1].RefTable = HiringJobsTable
	EntitySkillsTable.ForeignKeys[2].RefTable = SkillsTable
	HiringJobsTable.ForeignKeys[0].RefTable = TeamsTable
	HiringJobsTable.ForeignKeys[1].RefTable = UsersTable
	PermissionsTable.ForeignKeys[0].RefTable = PermissionGroupsTable
	PermissionGroupsTable.ForeignKeys[0].RefTable = PermissionGroupsTable
	SkillsTable.ForeignKeys[0].RefTable = SkillTypesTable
	TeamManagersTable.ForeignKeys[0].RefTable = UsersTable
	TeamManagersTable.ForeignKeys[1].RefTable = TeamsTable
	UsersTable.ForeignKeys[0].RefTable = TeamsTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}

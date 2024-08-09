package repository

import (
	"context"
	"fmt"
	"trec/ent"

	"github.com/pkg/errors"
)

// Repository is a registry of all repositories
type Repository interface {
	User() UserRepository
	JobPosition() JobPositionRepository
	HiringJob() HiringJobRepository
	AuditTrail() AuditTrailRepository
	Candidate() CandidateRepository
	CandidateJob() CandidateJobRepository
	Attachment() AttachmentRepository
	CandidateJobFeedback() CandidateJobFeedbackRepository
	CandidateInterview() CandidateInterviewRepository
	ImportData() ImportDataRepository
	CandidateJobStep() CandidateJobStepRepository
	CandidateInterviewer() CandidateInterviewerRepository
	Skill() SkillRepository
	EntitySkill() EntitySkillRepository
	SkillType() SkillTypeRepository
	Role() RoleRepository
	EntityPermission() EntityPermissionRepository
	PermissionGroup() PermissionGroupRepository
	EmailTemplate() EmailTemplateRepository
	OutgoingEmail() OutgoingEmailRepository
	Report() ReportRepository
	HiringTeam() HiringTeamRepository
	HiringTeamApprover() HiringTeamApproverRepository
	RecTeam() RecTeamRepository
	CandidateExp() CandidateExpRepository
	CandidateEducate() CandidateEducateRepository
	CandidateAward() CandidateAwardRepository
	CandidateCertificate() CandidateCertificateRepository
	CandidateHistoryCall() CandidateHistoryCallRepository
	CandidateNote() CandidateNoteRepository

	// DoInTx executes the given function in a transaction.
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error
}

// RepoImpl is implementation of Repository
type RepoImpl struct {
	entClient            *ent.Client
	entTx                *ent.Tx
	user                 UserRepository
	jobPosition          JobPositionRepository
	hiringJob            HiringJobRepository
	auditTrail           AuditTrailRepository
	candidate            CandidateRepository
	candidateJob         CandidateJobRepository
	attachment           AttachmentRepository
	candidateJobFeedback CandidateJobFeedbackRepository
	candidateInterview   CandidateInterviewRepository
	importData           ImportDataRepository
	candidateJobStep     CandidateJobStepRepository
	candidateInterviewer CandidateInterviewerRepository
	skill                SkillRepository
	entitySkill          EntitySkillRepository
	skillType            SkillTypeRepository
	role                 RoleRepository
	entityPermission     EntityPermissionRepository
	permissionGroup      PermissionGroupRepository
	emailTemplate        EmailTemplateRepository
	outgoingEmail        OutgoingEmailRepository
	report               ReportRepository
	hiringTeam           HiringTeamRepository
	hiringTeamApprover   HiringTeamApproverRepository
	recTeam              RecTeamRepository
	candidateExp         CandidateExpRepository
	candidateEducate     CandidateEducateRepository
	candidateAward       CandidateAwardRepository
	candidateCertificate CandidateCertificateRepository
	candidateHistoryCall CandidateHistoryCallRepository
	candidateNote        CandidateNoteRepository
}

// NewRepository creates new repository registry
func NewRepository(entClient *ent.Client) Repository {
	return &RepoImpl{
		entClient:            entClient,
		user:                 NewUserRepository(entClient),
		jobPosition:          NewJobPositionRepository(entClient),
		hiringJob:            NewHiringJobRepository(entClient),
		auditTrail:           NewAuditTrailRepository(entClient),
		candidate:            NewCandidateRepository(entClient),
		candidateJob:         NewCandidateJobRepository(entClient),
		attachment:           NewAttachmentRepository(entClient),
		candidateJobFeedback: NewCandidateJobFeedbackRepository(entClient),
		candidateInterview:   NewCandidateInterviewRepository(entClient),
		importData:           NewImportDataRepository(entClient),
		candidateJobStep:     NewCandidateJobStepRepository(entClient),
		candidateInterviewer: NewCandidateInterviewerRepository(entClient),
		skill:                NewSkillRepository(entClient),
		entitySkill:          NewEntitySkillRepository(entClient),
		skillType:            NewSkillTypeRepository(entClient),
		role:                 NewRoleRepository(entClient),
		entityPermission:     NewEntityPermissionRepository(entClient),
		permissionGroup:      NewPermissionGroupRepository(entClient),
		emailTemplate:        NewEmailTemplateRepository(entClient),
		outgoingEmail:        NewOutgoingEmailRepository(entClient),
		report:               NewReportRepository(entClient),
		hiringTeam:           NewHiringTeamRepository(entClient),
		hiringTeamApprover:   NewHiringTeamApproverRepository(entClient),
		recTeam:              NewRecTeamRepository(entClient),
		candidateExp:         NewCandidateExpRepository(entClient),
		candidateEducate:     NewCandidateEducateRepository(entClient),
		candidateAward:       NewCandidateAwardRepository(entClient),
		candidateCertificate: NewCandidateCertificateRepository(entClient),
		candidateHistoryCall: NewCandidateHistoryCallRepository(entClient),
		candidateNote:        NewCandidateNoteRepository(entClient),
	}
}

func (r *RepoImpl) User() UserRepository {
	return r.user
}

func (r *RepoImpl) JobPosition() JobPositionRepository {
	return r.jobPosition
}

func (r *RepoImpl) HiringJob() HiringJobRepository {
	return r.hiringJob
}

func (r *RepoImpl) AuditTrail() AuditTrailRepository {
	return r.auditTrail
}

func (r *RepoImpl) Candidate() CandidateRepository {
	return r.candidate
}

func (r *RepoImpl) CandidateJob() CandidateJobRepository {
	return r.candidateJob
}

func (r *RepoImpl) Attachment() AttachmentRepository {
	return r.attachment
}

func (r *RepoImpl) CandidateJobFeedback() CandidateJobFeedbackRepository {
	return r.candidateJobFeedback
}

func (r *RepoImpl) CandidateInterview() CandidateInterviewRepository {
	return r.candidateInterview
}

func (r *RepoImpl) ImportData() ImportDataRepository {
	return r.importData
}

func (r *RepoImpl) CandidateJobStep() CandidateJobStepRepository {
	return r.candidateJobStep
}

func (r *RepoImpl) CandidateInterviewer() CandidateInterviewerRepository {
	return r.candidateInterviewer
}

func (r *RepoImpl) Skill() SkillRepository {
	return r.skill
}

func (r *RepoImpl) EntitySkill() EntitySkillRepository {
	return r.entitySkill
}

func (r *RepoImpl) SkillType() SkillTypeRepository {
	return r.skillType
}

func (r *RepoImpl) Role() RoleRepository {
	return r.role
}

func (r *RepoImpl) EntityPermission() EntityPermissionRepository {
	return r.entityPermission
}

func (r *RepoImpl) PermissionGroup() PermissionGroupRepository {
	return r.permissionGroup
}

func (r *RepoImpl) EmailTemplate() EmailTemplateRepository {
	return r.emailTemplate
}

func (r *RepoImpl) OutgoingEmail() OutgoingEmailRepository {
	return r.outgoingEmail
}

func (r *RepoImpl) Report() ReportRepository {
	return r.report
}

func (r *RepoImpl) HiringTeam() HiringTeamRepository {
	return r.hiringTeam
}

func (r *RepoImpl) HiringTeamApprover() HiringTeamApproverRepository {
	return r.hiringTeamApprover
}

func (r *RepoImpl) RecTeam() RecTeamRepository {
	return r.recTeam
}

func (r *RepoImpl) CandidateExp() CandidateExpRepository {
	return r.candidateExp
}

func (r *RepoImpl) CandidateEducate() CandidateEducateRepository {
	return r.candidateEducate
}

func (r *RepoImpl) CandidateAward() CandidateAwardRepository {
	return r.candidateAward
}

func (r *RepoImpl) CandidateCertificate() CandidateCertificateRepository {
	return r.candidateCertificate
}

func (r *RepoImpl) CandidateHistoryCall() CandidateHistoryCallRepository {
	return r.candidateHistoryCall
}

func (r *RepoImpl) CandidateNote() CandidateNoteRepository {
	return r.candidateNote
}

// DoInTx executes the given function in a transaction.
func (r *RepoImpl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error {
	if r.entTx != nil {
		return errors.WithStack(errors.New("invalid tx state, no nested tx allowed"))
	}

	tx, err := r.entClient.Tx(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	committed := false

	defer func() {
		if committed {
			return
		}
		// rollback if not committed
		_ = tx.Rollback()
	}()

	impl := &RepoImpl{
		entTx:                tx,
		user:                 NewUserRepository(tx.Client()),
		jobPosition:          NewJobPositionRepository(tx.Client()),
		hiringJob:            NewHiringJobRepository(tx.Client()),
		auditTrail:           NewAuditTrailRepository(tx.Client()),
		candidate:            NewCandidateRepository(tx.Client()),
		candidateJob:         NewCandidateJobRepository(tx.Client()),
		attachment:           NewAttachmentRepository(tx.Client()),
		candidateJobFeedback: NewCandidateJobFeedbackRepository(tx.Client()),
		candidateInterview:   NewCandidateInterviewRepository(tx.Client()),
		candidateJobStep:     NewCandidateJobStepRepository(tx.Client()),
		candidateInterviewer: NewCandidateInterviewerRepository(tx.Client()),
		skill:                NewSkillRepository(tx.Client()),
		entitySkill:          NewEntitySkillRepository(tx.Client()),
		skillType:            NewSkillTypeRepository(tx.Client()),
		role:                 NewRoleRepository(tx.Client()),
		entityPermission:     NewEntityPermissionRepository(tx.Client()),
		emailTemplate:        NewEmailTemplateRepository(tx.Client()),
		outgoingEmail:        NewOutgoingEmailRepository(tx.Client()),
		hiringTeam:           NewHiringTeamRepository(tx.Client()),
		hiringTeamApprover:   NewHiringTeamApproverRepository(tx.Client()),
		recTeam:              NewRecTeamRepository(tx.Client()),
		candidateExp:         NewCandidateExpRepository(tx.Client()),
		candidateEducate:     NewCandidateEducateRepository(tx.Client()),
		candidateAward:       NewCandidateAwardRepository(tx.Client()),
		candidateCertificate: NewCandidateCertificateRepository(tx.Client()),
		candidateHistoryCall: NewCandidateHistoryCallRepository(tx.Client()),
		candidateNote:        NewCandidateNoteRepository(tx.Client()),
	}

	if err := txFunc(ctx, impl); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(fmt.Errorf("failed to commit tx: %s", err.Error()))
	}

	committed = true
	return nil
}

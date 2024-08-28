package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

// ID is the resolver for the id field.
func (r *hiringJobResolver) ID(ctx context.Context, obj *ent.HiringJob) (string, error) {
	return obj.ID.String(), nil
}

// Location is the resolver for the location field.
func (r *hiringJobResolver) Location(ctx context.Context, obj *ent.HiringJob) (ent.LocationEnum, error) {
	return ent.LocationEnum(obj.Location), nil
}

// SalaryType is the resolver for the salary_type field.
func (r *hiringJobResolver) SalaryType(ctx context.Context, obj *ent.HiringJob) (ent.SalaryTypeEnum, error) {
	return ent.SalaryTypeEnum(obj.SalaryType), nil
}

// Currency is the resolver for the currency field.
func (r *hiringJobResolver) Currency(ctx context.Context, obj *ent.HiringJob) (ent.CurrencyEnum, error) {
	return ent.CurrencyEnum(obj.Currency), nil
}

// JobPosition is the resolver for the job_position field.
func (r *hiringJobResolver) JobPosition(ctx context.Context, obj *ent.HiringJob) (*ent.JobPosition, error) {
	return obj.Edges.JobPositionEdge, nil
}

// JobPositionID is the resolver for the job_position_id field.
func (r *hiringJobResolver) JobPositionID(ctx context.Context, obj *ent.HiringJob) (*string, error) {
	jobPosition := ""
	if obj.JobPositionID != uuid.Nil {
		jobPosition = obj.JobPositionID.String()
	}
	return &jobPosition, nil
}

// HiringTeam is the resolver for the hiring_team field.
func (r *hiringJobResolver) HiringTeam(ctx context.Context, obj *ent.HiringJob) (*ent.HiringTeam, error) {
	return obj.Edges.HiringTeamEdge, nil
}

// RecTeam is the resolver for the rec_team field.
func (r *hiringJobResolver) RecTeam(ctx context.Context, obj *ent.HiringJob) (*ent.RecTeam, error) {
	return obj.Edges.RecTeamEdge, nil
}

// RecInCharge is the resolver for the rec_in_charge field.
func (r *hiringJobResolver) RecInCharge(ctx context.Context, obj *ent.HiringJob) (*ent.User, error) {
	return obj.Edges.RecInChargeEdge, nil
}

// User is the resolver for the user field.
func (r *hiringJobResolver) User(ctx context.Context, obj *ent.HiringJob) (*ent.User, error) {
	return obj.Edges.OwnerEdge, nil
}

// Status is the resolver for the status field.
func (r *hiringJobResolver) Status(ctx context.Context, obj *ent.HiringJob) (ent.HiringJobStatus, error) {
	return ent.HiringJobStatus(obj.Status), nil
}

// TotalCandidatesRecruited is the resolver for the total_candidates_recruited field.
func (r *hiringJobResolver) TotalCandidatesRecruited(ctx context.Context, obj *ent.HiringJob) (int, error) {
	return len(lo.Filter(obj.Edges.CandidateJobEdges, func(item *ent.CandidateJob, index int) bool {
		return item.Status == candidatejob.StatusHired || item.Status == candidatejob.StatusExStaff
	})), nil
}

// IsAbleToClose is the resolver for the is_able_to_close field.
func (r *hiringJobResolver) IsAbleToClose(ctx context.Context, obj *ent.HiringJob) (bool, error) {
	return obj.Status == hiringjob.StatusOpened, nil
}

// EntitySkillTypes is the resolver for the entity_skill_types field.
func (r *hiringJobResolver) EntitySkillTypes(ctx context.Context, obj *ent.HiringJob) ([]*ent.EntitySkillType, error) {
	return r.serviceRegistry.HiringJob().GroupSkillType(obj.Edges.HiringJobSkillEdges), nil
}

// Steps is the resolver for the steps field.
func (r *hiringJobResolver) Steps(ctx context.Context, obj *ent.HiringJob) ([]*ent.HiringJobStep, error) {
	return obj.Edges.ApprovalStepsOrErr()
}

// Level is the resolver for the level field.
func (r *hiringJobResolver) Level(ctx context.Context, obj *ent.HiringJob) (ent.HiringJobLevel, error) {
	return ent.HiringJobLevel(obj.Level), nil
}

// IsAbleToCancel is the resolver for the is_able_to_cancel field.
func (r *hiringJobResolver) IsAbleToCancel(ctx context.Context, obj *ent.HiringJob) (bool, error) {
	switch obj.Status {
	case hiringjob.StatusPendingApprovals:
		return true, nil
	case hiringjob.StatusOpened:
		canCancel := lo.NoneBy(
			obj.Edges.CandidateJobEdges,
			func(item *ent.CandidateJob) bool {
				return item.Status == candidatejob.StatusHired || item.Status == candidatejob.StatusExStaff
			},
		)
		return canCancel, nil
	default:
		return false, nil
	}
}

// HiringJob returns graphql1.HiringJobResolver implementation.
func (r *Resolver) HiringJob() graphql1.HiringJobResolver { return &hiringJobResolver{r} }

type hiringJobResolver struct{ *Resolver }

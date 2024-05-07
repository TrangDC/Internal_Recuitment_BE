package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// GetPreRequest is the resolver for the GetPreRequest field.
func (r *queryResolver) GetPreRequest(ctx context.Context) (string, error) {
	return "", nil
}

// GetTeam is the resolver for the GetTeam field.
func (r *queryResolver) GetTeam(ctx context.Context, id string) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().GetTeam(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllTeams is the resolver for the GetAllTeams field.
func (r *queryResolver) GetAllTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.TeamFilter, freeWord *ent.TeamFreeWord, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error) {
	result, err := r.serviceRegistry.Team().GetTeams(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelectionUsers is the resolver for the SelectionUsers field.
func (r *queryResolver) SelectionUsers(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error) {
	result, err := r.serviceRegistry.User().Selections(ctx, pagination, filter, freeWord, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHiringJob is the resolver for the GetHiringJob field.
func (r *queryResolver) GetHiringJob(ctx context.Context, id string) (*ent.HiringJobResponse, error) {
	result, err := r.serviceRegistry.HiringJob().GetHiringJob(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllHiringJobs is the resolver for the GetAllHiringJobs field.
func (r *queryResolver) GetAllHiringJobs(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringJobFilter, freeWord *ent.HiringJobFreeWord, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error) {
	result, err := r.serviceRegistry.HiringJob().GetHiringJobs(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAuditTrail is the resolver for the GetAuditTrail field.
func (r *queryResolver) GetAuditTrail(ctx context.Context, id string) (*ent.AuditTrailResponse, error) {
	result, err := r.serviceRegistry.AuditTrail().GetAuditTrail(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllAuditTrails is the resolver for the GetAllAuditTrails field.
func (r *queryResolver) GetAllAuditTrails(ctx context.Context, pagination *ent.PaginationInput, filter *ent.AuditTrailFilter, freeWord *ent.AuditTrailFreeWord, orderBy *ent.AuditTrailOrder) (*ent.AuditTrailResponseGetAll, error) {
	result, err := r.serviceRegistry.AuditTrail().GetAuditTrails(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCandidate is the resolver for the GetCandidate field.
func (r *queryResolver) GetCandidate(ctx context.Context, id string) (*ent.CandidateResponse, error) {
	result, err := r.serviceRegistry.Candidate().GetCandidate(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllCandidates is the resolver for the GetAllCandidates field.
func (r *queryResolver) GetAllCandidates(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateFilter, freeWord *ent.CandidateFreeWord, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error) {
	result, err := r.serviceRegistry.Candidate().GetCandidates(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCandidateJob is the resolver for the GetCandidateJob field.
func (r *queryResolver) GetCandidateJob(ctx context.Context, id string) (*ent.CandidateJobResponse, error) {
	result, err := r.serviceRegistry.CandidateJob().GetCandidateJob(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllCandidateJobs is the resolver for the GetAllCandidateJobs field.
func (r *queryResolver) GetAllCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateJobFilter, freeWord *ent.CandidateJobFreeWord, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	result, err := r.serviceRegistry.CandidateJob().GetCandidateJobs(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCandidateJobFeedback is the resolver for the GetCandidateJobFeedback field.
func (r *queryResolver) GetCandidateJobFeedback(ctx context.Context, id string) (*ent.CandidateJobFeedbackResponse, error) {
	panic(fmt.Errorf("not implemented: GetCandidateJobFeedback - GetCandidateJobFeedback"))
}

// GetAllCandidateJobFeedbacks is the resolver for the GetAllCandidateJobFeedbacks field.
func (r *queryResolver) GetAllCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateJobFeedbackFilter, freeWord *ent.CandidateJobFeedbackFreeWord, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error) {
	panic(fmt.Errorf("not implemented: GetAllCandidateJobFeedbacks - GetAllCandidateJobFeedbacks"))
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
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

// Team is the resolver for the team field.
func (r *hiringJobResolver) Team(ctx context.Context, obj *ent.HiringJob) (*ent.Team, error) {
	return obj.Edges.TeamEdge, nil
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
	return len(obj.Edges.CandidateJobEdges), nil
}

// HiringJob returns graphql1.HiringJobResolver implementation.
func (r *Resolver) HiringJob() graphql1.HiringJobResolver { return &hiringJobResolver{r} }

type hiringJobResolver struct{ *Resolver }

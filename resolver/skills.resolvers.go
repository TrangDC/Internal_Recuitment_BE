package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *skillResolver) ID(ctx context.Context, obj *ent.Skill) (string, error) {
	return obj.ID.String(), nil
}

// SkillType is the resolver for the skill_type field.
func (r *skillResolver) SkillType(ctx context.Context, obj *ent.Skill) (*ent.SkillType, error) {
	return obj.Edges.SkillTypeEdge, nil
}

// Skill returns graphql1.SkillResolver implementation.
func (r *Resolver) Skill() graphql1.SkillResolver { return &skillResolver{r} }

type skillResolver struct{ *Resolver }

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *skillTypeResolver) ID(ctx context.Context, obj *ent.SkillType) (string, error) {
	return obj.ID.String(), nil
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *skillTypeResolver) IsAbleToDelete(ctx context.Context, obj *ent.SkillType) (bool, error) {
	if len(obj.Edges.SkillEdges) > 0 {
		return false, nil
	}
	return true, nil
}

// Skills is the resolver for the skills field.
func (r *skillTypeResolver) Skills(ctx context.Context, obj *ent.SkillType) ([]*ent.Skill, error) {
	return obj.Edges.SkillEdges, nil
}

// SkillType returns graphql1.SkillTypeResolver implementation.
func (r *Resolver) SkillType() graphql1.SkillTypeResolver { return &skillTypeResolver{r} }

type skillTypeResolver struct{ *Resolver }

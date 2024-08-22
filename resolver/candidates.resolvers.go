package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	"trec/ent/candidatejob"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

// ID is the resolver for the id field.
func (r *candidateResolver) ID(ctx context.Context, obj *ent.Candidate) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *candidateResolver) Status(ctx context.Context, obj *ent.Candidate) (ent.CandidateStatusEnum, error) {
	if obj.Edges.CandidateJobEdges == nil || len(obj.Edges.CandidateJobEdges) == 0 {
		return ent.CandidateStatusEnumNew, nil
	}
	processingCdJobs := lo.Filter(obj.Edges.CandidateJobEdges, func(entity *ent.CandidateJob, _ int) bool {
		return ent.CandidateJobStatusOpen.IsValid(ent.CandidateJobStatusOpen(entity.Status))
	})
	if len(processingCdJobs) == 0 {
		return ent.CandidateStatusEnum(obj.Edges.CandidateJobEdges[0].Status), nil
	}
	return ent.CandidateStatusEnum(processingCdJobs[0].Status), nil
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *candidateResolver) IsAbleToDelete(ctx context.Context, obj *ent.Candidate) (bool, error) {
	candidateJobStatusOpen := lo.Map(ent.AllCandidateJobStatusOpen, func(entity ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(entity.String())
	})
	var result bool
	for _, entity := range obj.Edges.CandidateJobEdges {
		if lo.Contains(candidateJobStatusOpen, entity.Status) {
			result = true
			break
		}
	}
	return result, nil
}

// HiringJobTitle is the resolver for the hiring_job_title field.
func (r *candidateResolver) HiringJobTitle(ctx context.Context, obj *ent.Candidate) (string, error) {
	candidateJobStatusOpen := lo.Map(ent.AllCandidateJobStatusOpen, func(entity ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(entity.String())
	})
	var result string
	for _, entity := range obj.Edges.CandidateJobEdges {
		if lo.Contains(candidateJobStatusOpen, entity.Status) {
			result = entity.Edges.HiringJobEdge.Name
			break
		}
	}
	return result, nil
}

// ReferenceType is the resolver for the reference_type field.
func (r *candidateResolver) ReferenceType(ctx context.Context, obj *ent.Candidate) (ent.CandidateReferenceType, error) {
	return ent.CandidateReferenceType(obj.ReferenceType), nil
}

// ReferenceUID is the resolver for the reference_uid field.
func (r *candidateResolver) ReferenceUID(ctx context.Context, obj *ent.Candidate) (string, error) {
	return obj.ReferenceUID.String(), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateResolver) Attachments(ctx context.Context, obj *ent.Candidate) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// EntitySkillTypes is the resolver for the entity_skill_types field.
func (r *candidateResolver) EntitySkillTypes(ctx context.Context, obj *ent.Candidate) ([]*ent.EntitySkillType, error) {
	return r.serviceRegistry.Candidate().GroupSkillType(obj.Edges.CandidateSkillEdges), nil
}

// ReferenceUser is the resolver for the reference_user field.
func (r *candidateResolver) ReferenceUser(ctx context.Context, obj *ent.Candidate) (*ent.User, error) {
	return obj.Edges.ReferenceUserEdge, nil
}

// CandidateExp is the resolver for the candidate_exp field.
func (r *candidateResolver) CandidateExp(ctx context.Context, obj *ent.Candidate) ([]*ent.CandidateExp, error) {
	return obj.Edges.CandidateExpEdges, nil
}

// CandidateEducate is the resolver for the candidate_educate field.
func (r *candidateResolver) CandidateEducate(ctx context.Context, obj *ent.Candidate) ([]*ent.CandidateEducate, error) {
	return obj.Edges.CandidateEducateEdges, nil
}

// CandidateAward is the resolver for the candidate_award field.
func (r *candidateResolver) CandidateAward(ctx context.Context, obj *ent.Candidate) ([]*ent.CandidateAward, error) {
	return obj.Edges.CandidateAwardEdges, nil
}

// CandidateCertificate is the resolver for the candidate_certificate field.
func (r *candidateResolver) CandidateCertificate(ctx context.Context, obj *ent.Candidate) ([]*ent.CandidateCertificate, error) {
	return obj.Edges.CandidateCertificateEdges, nil
}

// Avatar is the resolver for the avatar field.
func (r *candidateResolver) Avatar(ctx context.Context, obj *ent.Candidate) (string, error) {
	if obj.Avatar == uuid.Nil {
		return "", nil
	}
	return obj.Avatar.String(), nil
}

// Gender is the resolver for the gender field.
func (r *candidateResolver) Gender(ctx context.Context, obj *ent.Candidate) (ent.CandidateGenderEnum, error) {
	return ent.CandidateGenderEnum(obj.Gender), nil
}

// Candidate returns graphql1.CandidateResolver implementation.
func (r *Resolver) Candidate() graphql1.CandidateResolver { return &candidateResolver{r} }

type candidateResolver struct{ *Resolver }

package repository

import (
	"context"
	"time"
	"trec/ent"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateInterviewerRepository interface {
	CreateBulkCandidateInterview(ctx context.Context, memberIds []uuid.UUID, candidateInterviews []*ent.CandidateInterview) (
		[]*ent.CandidateInterviewer, error)
}

type candidateInterviewerRepoImpl struct {
	client *ent.Client
}

func NewCandidateInterviewerRepository(client *ent.Client) CandidateInterviewerRepository {
	return &candidateInterviewerRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *candidateInterviewerRepoImpl) BuildCreate() *ent.CandidateInterviewerCreate {
	return rps.client.CandidateInterviewer.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *candidateInterviewerRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateInterviewerCreate) ([]*ent.CandidateInterviewer, error) {
	return rps.client.CandidateInterviewer.CreateBulk(input...).Save(ctx)
}

func (rps *candidateInterviewerRepoImpl) CreateBulkCandidateInterview(ctx context.Context, memberIds []uuid.UUID, candidateInterviews []*ent.CandidateInterview) (
	[]*ent.CandidateInterviewer, error) {
	var createInterviewers []*ent.CandidateInterviewerCreate
	for _, record := range candidateInterviews {
		createBulkThings := lo.Map(memberIds, func(item uuid.UUID, index int) *ent.CandidateInterviewerCreate {
			return rps.BuildCreate().SetCandidateInterviewID(record.ID).SetUserID(item).SetID(uuid.New())
		})
		createInterviewers = append(createInterviewers, createBulkThings...)
	}
	return rps.BuildBulkCreate(ctx, createInterviewers)
}

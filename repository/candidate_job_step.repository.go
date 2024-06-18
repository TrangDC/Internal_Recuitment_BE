package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"

	"github.com/google/uuid"
)

type CandidateJobStepRepository interface {
	// mutation
	CreateCandidateJobStep(ctx context.Context, status candidatejob.Status, candidateJobId uuid.UUID) error
}

type candidateJobStepRepoImpl struct {
	client *ent.Client
}

func NewCandidateJobStepRepository(client *ent.Client) CandidateJobStepRepository {
	return &candidateJobStepRepoImpl{
		client: client,
	}
}

// Base function
func (rps candidateJobStepRepoImpl) BuildCreate() *ent.CandidateJobStepCreate {
	return rps.client.CandidateJobStep.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

// mutation
func (rps candidateJobStepRepoImpl) CreateCandidateJobStep(ctx context.Context, status candidatejob.Status, candidateJobId uuid.UUID) error {
	_, err := rps.BuildCreate().SetCandidateJobStatus(candidatejobstep.CandidateJobStatus(status.String())).SetCandidateJobID(candidateJobId).Save(ctx)
	return err
}

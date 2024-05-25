package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidatejob"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateRepository interface {
	// mutation
	CreateCandidate(ctx context.Context, input *ent.NewCandidateInput) (*ent.Candidate, error)
	UpdateCandidate(ctx context.Context, record *ent.Candidate, input *ent.UpdateCandidateInput) (*ent.Candidate, error)
	DeleteCandidate(ctx context.Context, record *ent.Candidate) error
	BuildBulkCreate(ctx context.Context, input []*ent.NewCandidateInput) ([]*ent.Candidate, error)
	SetBlackListCandidate(ctx context.Context, record *ent.Candidate, isBlackList bool) (*ent.Candidate, error)
	// query
	GetCandidate(ctx context.Context, candidateId uuid.UUID) (*ent.Candidate, error)
	BuildQuery() *ent.CandidateQuery
	BuildCount(ctx context.Context, query *ent.CandidateQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateQuery) ([]*ent.Candidate, error)
	// common function
	ValidEmail(ctx context.Context, candidateId uuid.UUID, email string) error
}

type candidateRepoImpl struct {
	client *ent.Client
}

func NewCandidateRepository(client *ent.Client) CandidateRepository {
	return &candidateRepoImpl{
		client: client,
	}
}

var defaultTime = time.Time{}

// Base function
func (rps candidateRepoImpl) BuildCreate() *ent.CandidateCreate {
	return rps.client.Candidate.Create().SetUpdatedAt(time.Now())
}

func (rps candidateRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.NewCandidateInput) ([]*ent.Candidate, error) {
	var createBulk []*ent.CandidateCreate
	for _, v := range input {
		create := rps.BuildCreate().
			SetName(strings.TrimSpace(v.Name)).
			SetEmail(strings.TrimSpace(v.Email)).
			SetPhone(strings.TrimSpace(v.Phone))
		if v.Dob != &defaultTime {
			create.SetDob(*v.Dob)
		}
		createBulk = append(createBulk, create)
	}
	return rps.client.Candidate.CreateBulk(createBulk...).Save(ctx)
}

func (rps candidateRepoImpl) BuildUpdate() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateRepoImpl) BuildDelete() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps candidateRepoImpl) BuildQuery() *ent.CandidateQuery {
	candidateJobStatusOpen := lo.Map(ent.AllCandidateJobStatusOpen, func(entity ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(entity.String())
	})
	return rps.client.Candidate.Query().Where(candidate.DeletedAtIsNil()).WithCandidateJobEdges(
		func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.StatusIn(candidateJobStatusOpen...), candidatejob.DeletedAtIsNil()).Limit(1).WithHiringJobEdge()
		},
	)
}

func (rps candidateRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateQuery) (*ent.Candidate, error) {
	return query.First(ctx)
}

func (rps candidateRepoImpl) BuildList(ctx context.Context, query *ent.CandidateQuery) ([]*ent.Candidate, error) {
	return query.All(ctx)
}

func (rps candidateRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateQuery) (int, error) {
	return query.Count(ctx)
}

func (rps candidateRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps candidateRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.Candidate) *ent.CandidateUpdateOne {
	return model.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateUpdateOne) (*ent.Candidate, error) {
	return update.Save(ctx)
}

// mutation
func (rps candidateRepoImpl) CreateCandidate(ctx context.Context, input *ent.NewCandidateInput) (*ent.Candidate, error) {
	create := rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetEmail(strings.TrimSpace(input.Email)).
		SetPhone(strings.TrimSpace(input.Phone))
	if input.Dob != nil && *input.Dob != defaultTime {
		create.SetDob(*input.Dob)
	}
	return create.Save(ctx)
}

func (rps candidateRepoImpl) UpdateCandidate(ctx context.Context, record *ent.Candidate, input *ent.UpdateCandidateInput) (*ent.Candidate, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetEmail(strings.TrimSpace(input.Email)).
		SetPhone(strings.TrimSpace(input.Phone))
	if input.Dob != nil && *input.Dob != defaultTime {
		update.SetDob(*input.Dob)
	}
	return update.Save(ctx)
}

func (rps candidateRepoImpl) DeleteCandidate(ctx context.Context, record *ent.Candidate) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

func (rps candidateRepoImpl) SetBlackListCandidate(ctx context.Context, record *ent.Candidate, isBlackList bool) (*ent.Candidate, error) {
	return rps.BuildUpdateOne(ctx, record).SetIsBlacklist(isBlackList).Save(ctx)
}

// query
func (rps candidateRepoImpl) GetCandidate(ctx context.Context, candidateId uuid.UUID) (*ent.Candidate, error) {
	return rps.BuildQuery().Where(candidate.IDEQ(candidateId)).First(ctx)
}

// common function
func (rps candidateRepoImpl) ValidEmail(ctx context.Context, candidateId uuid.UUID, email string) error {
	query := rps.BuildQuery().Where(candidate.EmailEqualFold(strings.TrimSpace(strings.ToLower(email))))
	if candidateId != uuid.Nil {
		query = query.Where(candidate.IDNEQ(candidateId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if isExist {
		return fmt.Errorf("model.candidates.validation.email_exist")
	}
	return err
}

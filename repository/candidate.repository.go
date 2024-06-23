package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidatejob"

	"github.com/google/uuid"
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
	ValidEmail(ctx context.Context, candidateId uuid.UUID, email string) (error, error)
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
	return rps.client.Candidate.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
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
	return rps.client.Candidate.Query().Where(candidate.DeletedAtIsNil()).WithCandidateJobEdges(
		func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.DeletedAtIsNil()).Order(ent.Desc(candidatejob.FieldUpdatedAt)).WithHiringJobEdge()
		},
	).WithReferenceUserEdge().WithAttachmentEdges(
		func(query *ent.AttachmentQuery) {
			query.Where(attachment.DeletedAtIsNil())
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
		SetPhone(strings.TrimSpace(input.Phone)).
		SetCountry(strings.TrimSpace(input.Country)).
		SetReferenceType(candidate.ReferenceType(input.ReferenceType)).
		SetReferenceValue(strings.TrimSpace(input.ReferenceValue)).
		SetDescription(strings.TrimSpace(input.Description))
	if input.Dob != nil && !input.Dob.IsZero() {
		create.SetDob(*input.Dob)
	}
	if input.ReferenceUID != "" {
		create.SetReferenceUID(uuid.MustParse(input.ReferenceUID))
	}
	if input.RecruitTime != nil && !input.RecruitTime.IsZero() {
		create.SetRecruitTime(*input.RecruitTime)
	}
	return create.Save(ctx)
}

func (rps candidateRepoImpl) UpdateCandidate(ctx context.Context, record *ent.Candidate, input *ent.UpdateCandidateInput) (*ent.Candidate, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetEmail(strings.TrimSpace(input.Email)).
		SetPhone(strings.TrimSpace(input.Phone)).
		SetCountry(strings.TrimSpace(input.Country)).
		SetReferenceType(candidate.ReferenceType(input.ReferenceType)).
		SetReferenceValue(strings.TrimSpace(input.ReferenceValue)).
		SetDescription(strings.TrimSpace(input.Description))
	if input.Dob != nil && !input.Dob.IsZero() {
		update.SetDob(*input.Dob)
	} else {
		update.ClearDob()
	}
	if input.ReferenceUID != "" {
		update.SetReferenceUID(uuid.MustParse(input.ReferenceUID))
	} else {
		update.ClearReferenceUID()
	}
	if input.RecruitTime != nil && !input.RecruitTime.IsZero() {
		update.SetRecruitTime(*input.RecruitTime)
	} else {
		update.ClearRecruitTime()
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
func (rps candidateRepoImpl) ValidEmail(ctx context.Context, candidateId uuid.UUID, email string) (error, error) {
	query := rps.BuildQuery().Where(candidate.EmailEqualFold(strings.TrimSpace(strings.ToLower(email))))
	if candidateId != uuid.Nil {
		query = query.Where(candidate.IDNEQ(candidateId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.candidates.validation.email_exist"), nil
	}
	return nil, nil
}

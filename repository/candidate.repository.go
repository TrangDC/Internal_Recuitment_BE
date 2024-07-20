package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatejobstep"
	"trec/ent/entityskill"
	"trec/ent/skill"
	"trec/ent/skilltype"

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
	DeleteRelationCandidate(ctx context.Context, candidateId uuid.UUID) error
	// query
	GetCandidate(ctx context.Context, candidateId uuid.UUID) (*ent.Candidate, error)
	BuildQuery() *ent.CandidateQuery
	BuildBaseQuery() *ent.CandidateQuery
	BuildCount(ctx context.Context, query *ent.CandidateQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateQuery) ([]*ent.Candidate, error)
	BuildGet(ctx context.Context, query *ent.CandidateQuery) (*ent.Candidate, error)
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
	).WithCandidateSkillEdges(
		func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).WithSkillEdge(
				func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(
						func(stq *ent.SkillTypeQuery) {
							stq.Where(skilltype.DeletedAtIsNil())
						},
					)
				},
			)
		},
	)
}

func (rps candidateRepoImpl) BuildBaseQuery() *ent.CandidateQuery {
	return rps.client.Candidate.Query().Where(candidate.DeletedAtIsNil())
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

func (rps candidateRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.Candidate) *ent.CandidateUpdateOne {
	return record.Update().SetUpdatedAt(time.Now().UTC())
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
	_, err := rps.BuildUpdateOne(ctx, record).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	candidateJobIds := lo.Map(record.Edges.CandidateJobEdges, func(v *ent.CandidateJob, index int) uuid.UUID {
		return v.ID
	})
	_, err = rps.client.CandidateJob.Update().Where(candidatejob.CandidateID(record.ID)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobFeedback.Update().Where(candidatejobfeedback.CandidateJobIDIn(candidateJobIds...)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobStep.Update().Where(candidatejobstep.CandidateJobID(record.ID)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	attachmentRelationIds := candidateJobIds
	attachmentRelationIds = append(attachmentRelationIds, record.ID)
	_, err = rps.client.Attachment.Update().Where(attachment.RelationIDIn(attachmentRelationIds...)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateInterview.Update().Where(candidateinterview.CandidateJobIDIn(candidateJobIds...)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.EntitySkill.Update().Where(entityskill.EntityIDEQ(record.ID)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (rps candidateRepoImpl) SetBlackListCandidate(ctx context.Context, record *ent.Candidate, isBlackList bool) (*ent.Candidate, error) {
	return rps.BuildUpdateOne(ctx, record).SetIsBlacklist(isBlackList).Save(ctx)
}

func (rps candidateRepoImpl) DeleteRelationCandidate(ctx context.Context, candidateId uuid.UUID) error {
	_, err := rps.client.CandidateJob.Update().Where(candidatejob.CandidateID(candidateId)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobFeedback.Update().Where(candidatejobfeedback.HasCandidateJobEdgeWith(
		candidatejob.CandidateID(candidateId),
	)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobStep.Delete().Where(candidatejobstep.HasCandidateJobEdgeWith(
		candidatejob.CandidateID(candidateId),
	)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.Attachment.Update().Where(attachment.RelationID(candidateId)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateInterview.Update().Where(candidateinterview.HasCandidateJobEdgeWith(
		candidatejob.CandidateID(candidateId),
	)).
		SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.EntitySkill.Delete().Where(entityskill.EntityIDEQ(candidateId)).Exec(ctx)
	return err
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

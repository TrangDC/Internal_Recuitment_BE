package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/hiringteam"
	"trec/ent/hiringteamapprover"
	"trec/ent/hiringteammanager"
	"trec/ent/user"
	"trec/internal/util"

	"github.com/google/uuid"
)

type HiringTeamRepository interface {
	CreateHiringTeam(ctx context.Context, input ent.NewHiringTeamInput, memberIds []uuid.UUID) (*ent.HiringTeam, error)
	UpdateHiringTeam(ctx context.Context, record *ent.HiringTeam, input ent.UpdateHiringTeamInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.HiringTeam, error)
	DeleteHiringTeam(ctx context.Context, record *ent.HiringTeam, memberIds []uuid.UUID) (*ent.HiringTeam, error)
	DeleteRelationHiringTeam(ctx context.Context, hiringTeamID uuid.UUID) error

	// query
	GetHiringTeam(ctx context.Context, id uuid.UUID) (*ent.HiringTeam, error)
	BuildQuery() *ent.HiringTeamQuery
	BuildBaseQuery() *ent.HiringTeamQuery
	BuildCount(ctx context.Context, query *ent.HiringTeamQuery) (int, error)
	BuildList(ctx context.Context, query *ent.HiringTeamQuery) ([]*ent.HiringTeam, error)
	BuildGetOne(ctx context.Context, query *ent.HiringTeamQuery) (*ent.HiringTeam, error)
	IsManagerOfHiringTeam(ctx context.Context, userID uuid.UUID) (bool, error)

	// common function
	ValidInput(ctx context.Context, hiringTeamID uuid.UUID, name string, memberIds []uuid.UUID, approverCount int) (error, error)
}

type hiringTeamRepoImpl struct {
	client *ent.Client
}

func NewHiringTeamRepository(client *ent.Client) HiringTeamRepository {
	return &hiringTeamRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *hiringTeamRepoImpl) BuildCreate() *ent.HiringTeamCreate {
	return rps.client.HiringTeam.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *hiringTeamRepoImpl) BuildUpdate() *ent.HiringTeamUpdate {
	return rps.client.HiringTeam.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *hiringTeamRepoImpl) BuildDelete() *ent.HiringTeamUpdate {
	return rps.client.HiringTeam.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps *hiringTeamRepoImpl) BuildQuery() *ent.HiringTeamQuery {
	return rps.client.HiringTeam.Query().Where(hiringteam.DeletedAtIsNil()).
		WithUserEdges(func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		}).
		WithHiringTeamJobEdges(func(query *ent.HiringJobQuery) {
			query.Where(hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusPendingApprovals)).
				Order(ent.Desc(hiringjob.FieldLastApplyDate)).
				WithApprovalSteps(func(query *ent.HiringJobStepQuery) {
					query.WithApprovalUser().Order(ent.Asc(hiringjobstep.FieldOrderID))
				})
		}).
		WithHiringMemberEdges(func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		}).
		WithHiringTeamApprovers(func(query *ent.HiringTeamApproverQuery) {
			query.WithUserEdge().Order(ent.Asc(hiringteamapprover.FieldOrderID))
		})
}

func (rps *hiringTeamRepoImpl) BuildGetOne(ctx context.Context, query *ent.HiringTeamQuery) (*ent.HiringTeam, error) {
	return query.First(ctx)
}

func (rps *hiringTeamRepoImpl) BuildBaseQuery() *ent.HiringTeamQuery {
	return rps.client.HiringTeam.Query().Where(hiringteam.DeletedAtIsNil())
}

func (rps *hiringTeamRepoImpl) BuildGet(ctx context.Context, query *ent.HiringTeamQuery) (*ent.HiringTeam, error) {
	return query.First(ctx)
}

func (rps *hiringTeamRepoImpl) BuildList(ctx context.Context, query *ent.HiringTeamQuery) ([]*ent.HiringTeam, error) {
	return query.All(ctx)
}

func (rps *hiringTeamRepoImpl) BuildCount(ctx context.Context, query *ent.HiringTeamQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *hiringTeamRepoImpl) BuildExist(ctx context.Context, query *ent.HiringTeamQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *hiringTeamRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.HiringTeam) *ent.HiringTeamUpdateOne {
	return record.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *hiringTeamRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.HiringTeamUpdateOne) (*ent.HiringTeam, error) {
	return update.Save(ctx)
}

// mutation
func (rps *hiringTeamRepoImpl) CreateHiringTeam(ctx context.Context, input ent.NewHiringTeamInput, memberIds []uuid.UUID) (*ent.HiringTeam, error) {
	create := rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).SetSlug(util.SlugGeneration(input.Name)).
		SetDescription(input.Description).
		AddUserEdgeIDs(memberIds...)
	return create.Save(ctx)
}

func (rps *hiringTeamRepoImpl) UpdateHiringTeam(ctx context.Context, record *ent.HiringTeam, input ent.UpdateHiringTeamInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.HiringTeam, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).SetSlug(util.SlugGeneration(input.Name)).
		SetDescription(input.Description).
		AddUserEdgeIDs(newMemberIds...).RemoveUserEdgeIDs(removeMemberIds...)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *hiringTeamRepoImpl) DeleteHiringTeam(ctx context.Context, record *ent.HiringTeam, memberIds []uuid.UUID) (*ent.HiringTeam, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC()).
		RemoveUserEdgeIDs(memberIds...)
	return update.Save(ctx)
}

func (rps *hiringTeamRepoImpl) DeleteRelationHiringTeam(ctx context.Context, hiringTeamID uuid.UUID) error {
	_, err := rps.client.HiringJob.Update().Where(hiringjob.HiringTeamID(hiringTeamID)).SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.EntitySkill.Delete().Where(entityskill.HasHiringJobEdgeWith(hiringjob.HiringTeamID(hiringTeamID))).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJob.Update().Where(candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(hiringTeamID))).SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).ClearCandidateJobStep().Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobStep.Delete().Where(candidatejobstep.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(hiringTeamID)))).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateInterview.Update().Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(hiringTeamID)))).SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.HiringTeamApprover.Delete().Where(hiringteamapprover.HiringTeamIDEQ(hiringTeamID)).Exec(ctx)
	return err
}

// query
func (rps *hiringTeamRepoImpl) GetHiringTeam(ctx context.Context, id uuid.UUID) (*ent.HiringTeam, error) {
	query := rps.BuildQuery().Where(hiringteam.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

func (rps *hiringTeamRepoImpl) IsManagerOfHiringTeam(ctx context.Context, userID uuid.UUID) (bool, error) {
	query := rps.BuildQuery().
		Where(hiringteam.HasUserEdgesWith(user.IDEQ(userID)))
	return rps.BuildExist(ctx, query)
}

// common function
func (rps *hiringTeamRepoImpl) ValidInput(ctx context.Context, hiringTeamID uuid.UUID, name string, memberIds []uuid.UUID, approverCount int) (error, error) {
	query := rps.BuildQuery()
	if hiringTeamID != uuid.Nil {
		query = query.Where(hiringteam.IDNEQ(hiringTeamID))
	}
	nameExist, err := rps.BuildExist(ctx, query.Clone().Where(hiringteam.NameEqualFold(strings.TrimSpace(name))))
	if err != nil {
		return nil, err
	}
	if nameExist {
		return fmt.Errorf("model.hiring_teams.validation.name_exist"), nil
	}
	userInAnotherTeam, err := rps.BuildExist(ctx, query.Clone().Where(hiringteam.HasUserEdgesWith(user.IDIn(memberIds...)), hiringteam.HasUserHiringTeamsWith(hiringteammanager.DeletedAtIsNil())))
	if err != nil {
		return nil, err
	}
	if userInAnotherTeam {
		return fmt.Errorf("model.hiring_teams.validation.user_in_another_hiring_team"), nil
	}
	if approverCount == 0 {
		return fmt.Errorf("model.hiring_teams.validation.empty_approver_list"), nil
	}
	return nil, nil
}

// Path: repository/hiring_team.repository.go

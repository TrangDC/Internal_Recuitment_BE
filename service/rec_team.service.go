package service

import (
	"context"
	"net/http"
	"trec/dto"
	"trec/ent"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type RecTeamService interface {
	// mutation
	CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput, note string) (*ent.RecTeamResponse, error)
	DeleteRecTeam(ctx context.Context, id uuid.UUID, note string) error
}

type recTeamSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewRecTeamService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) RecTeamService {
	return &recTeamSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

func (s *recTeamSvcImpl) CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput, note string) (*ent.RecTeamResponse, error) {
	var result *ent.RecTeam
	errString, err := s.repoRegistry.RecTeam().ValidInput(ctx, uuid.Nil, input.Name, uuid.MustParse(input.LeaderID))
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = s.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.RecTeam().CreateRecTeam(ctx, input)
		return err
	})
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	results, _ := s.repoRegistry.RecTeam().GetRecTeam(ctx, result.ID)
	return &ent.RecTeamResponse{
		Data: results,
	}, nil
}

func (s *recTeamSvcImpl) DeleteRecTeam(ctx context.Context, id uuid.UUID, note string) error {
	recTeam, err := s.repoRegistry.RecTeam().GetRecTeam(ctx, id)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(recTeam.Edges.RecMemberEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = s.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.RecTeam().DeleteRecTeam(ctx, recTeam, memberIds)
		return err
	})
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}
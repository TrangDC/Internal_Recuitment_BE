package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/team"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateJobFeedbackService interface {
	// mutation
	CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error)
	UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error)
	DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID, note string) error
	// query
	GetCandidateJobFeedback(ctx context.Context, id uuid.UUID) (*ent.CandidateJobFeedbackResponse, error)
	GetCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFeedbackFreeWord, filter *ent.CandidateJobFeedbackFilter, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error)
}
type candidateJobFeedbackSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	dtoRegistry   dto.Dto
	logger        *zap.Logger
}

func NewCandidateJobFeedbackService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateJobFeedbackService {
	return &candidateJobFeedbackSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		dtoRegistry:   dtoRegistry,
		logger:        logger,
	}
}

func (svc *candidateJobFeedbackSvcImpl) CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error) {
	var candidateJobFeedback *ent.CandidateJobFeedback
	status, errString, err := svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, uuid.MustParse(input.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, uuid.MustParse(input.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().CreateCandidateJobFeedback(ctx, input, status)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJobFeedback.ID, attachment.RelationTypeCandidateJobFeedbacks, repoRegistry)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, candidateJobFeedback.ID)
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	createdById := payload.UserID
	if record.CreatedBy != createdById {
		return nil, util.WrapGQLError(ctx, "model.candidate_job_feedbacks.validation.editor_not_is_owner", http.StatusBadGateway, util.ErrorFlagValidateFail)
	}
	_, errString, err := svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		_, err = repoRegistry.CandidateJobFeedback().UpdateCandidateJobFeedback(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, record.ID, input.Attachments, record.Edges.AttachmentEdges, attachment.RelationTypeCandidateJobFeedbacks)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.Edges.CandidateJobEdge.HiringJobID)).WithTeamEdge(
		func(query *ent.TeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, _ := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if !svc.validPermissionDelete(payload, record, hiringJob.Edges.TeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	_, errString, err := svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateJobFeedback().DeleteCandidateJobFeedback(ctx, record)
		if err != nil {
			return err
		}
		err = svc.attachmentSvc.RemoveAttachment(ctx, record.ID, repoRegistry)
		return err
	})
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobFeedbackSvcImpl) GetCandidateJobFeedback(ctx context.Context, id uuid.UUID) (*ent.CandidateJobFeedbackResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.CandidateJobFeedback().BuildQuery().Where(candidatejobfeedback.IDEQ(id))
	svc.validPermissionGet(payload, query)
	result, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) GetCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFeedbackFreeWord, filter *ent.CandidateJobFeedbackFilter, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateJobFeedbackResponseGetAll
	var edges []*ent.CandidateJobFeedbackEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJobFeedback().BuildQuery().Where(candidatejobfeedback.CandidateJobID(uuid.MustParse(filter.CandidateJobID)))
	svc.validPermissionGet(payload, query)
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJobFeedback().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidatejobfeedback.FieldCreatedAt)
	if orderBy != nil {
		order = ent.Desc(strings.ToLower(orderBy.Field.String()))
		if orderBy.Direction == ent.OrderDirectionAsc {
			order = ent.Asc(strings.ToLower(orderBy.Field.String()))
		}
	}
	query = query.Order(order)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	candidateJobFeedbacks, err := svc.repoRegistry.CandidateJobFeedback().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateJobFeedbacks, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) *ent.CandidateJobFeedbackEdge {
		return &ent.CandidateJobFeedbackEdge{
			Node: candidateJobFeedback,
			Cursor: ent.Cursor{
				Value: candidateJobFeedback.ID.String(),
			},
		}
	})
	result = &ent.CandidateJobFeedbackResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

// common function
func (svc *candidateJobFeedbackSvcImpl) freeWord(candidateJobFeedbackQuery *ent.CandidateJobFeedbackQuery, input *ent.CandidateJobFeedbackFreeWord) {
	predicate := []predicate.CandidateJobFeedback{}
	if input != nil {
		if input.Feedback != nil {
			predicate = append(predicate, candidatejobfeedback.FeedbackContainsFold(strings.TrimSpace(*input.Feedback)))
		}
		if input.UserName != nil {
			predicate = append(predicate, candidatejobfeedback.HasCreatedByEdgeWith(
				user.NameContainsFold(strings.TrimSpace(*input.UserName)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobFeedbackQuery.Where(candidatejobfeedback.Or(predicate...))
	}
}

func (svc *candidateJobFeedbackSvcImpl) filter(candidateJobFeedbackQuery *ent.CandidateJobFeedbackQuery, input *ent.CandidateJobFeedbackFilter) {
	if input.CreatedBy != nil {
		candidateJobFeedbackQuery.Where(candidatejobfeedback.CreatedByEQ(uuid.MustParse(*input.CreatedBy)))
	}
}

// permission
func (svc candidateJobFeedbackSvcImpl) validPermissionDelete(payload *middleware.Payload, record *ent.CandidateJobFeedback, teamRecord *ent.Team) bool {
	memberIds := lo.Map(teamRecord.Edges.MemberEdges, func(item *ent.User, index int) uuid.UUID {
		return item.ID
	})
	managerIds := lo.Map(teamRecord.Edges.UserEdges, func(item *ent.User, index int) uuid.UUID {
		return item.ID
	})
	if payload.ForTeam {
		if lo.Contains(memberIds, payload.UserID) || lo.Contains(managerIds, payload.UserID) {
			return true
		}
	}
	if payload.ForAll {
		return true
	}
	if payload.ForOwner && record.CreatedBy == payload.UserID {
		return true
	}
	return false
}

func (svc candidateJobFeedbackSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.CandidateJobFeedbackQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForTeam {
		query.Where(candidatejobfeedback.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HasTeamEdgeWith(
			team.Or(team.HasUserEdgesWith(user.IDEQ(payload.UserID)), team.HasMemberEdgesWith(user.IDEQ(payload.UserID))),
		))))
	}
	if payload.ForOwner {
		query.Where(candidatejobfeedback.CreatedByEQ(payload.UserID))
	}
}

// Path: service/candidate_job_feedback.service.go

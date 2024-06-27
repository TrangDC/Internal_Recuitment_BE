package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/team"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateJobService interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, note string) (*ent.CandidateJobResponse, error)
	UpdateCandidateJobStatus(ctx context.Context, input ent.UpdateCandidateJobStatus, id uuid.UUID, note string) (*ent.CandidateJobResponse, error)
	DeleteCandidateJob(ctx context.Context, id uuid.UUID, note string) error
	UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID, note string) (*ent.CandidateJobResponse, error)

	// query
	GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error)
	GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error)
	GetCandidateStatus(ctx context.Context, id uuid.UUID) ent.CandidateStatusEnum
	GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput,
		filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (
		*ent.CandidateJobGroupByStatusResponse, error)
	GetCandidateJobGroupByInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateJobGroupByInterviewResponse, error)
}

type candidateJobSvcImpl struct {
	attachmentSvc       AttachmentService
	candidateJobStepSvc CandidateJobStepService
	repoRegistry        repository.Repository
	dtoRegistry         dto.Dto
	logger              *zap.Logger
}

func NewCandidateJobService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateJobService {
	return &candidateJobSvcImpl{
		attachmentSvc:       NewAttachmentService(repoRegistry, logger),
		candidateJobStepSvc: NewCandidateJobStepService(repoRegistry, logger),
		repoRegistry:        repoRegistry,
		dtoRegistry:         dtoRegistry,
		logger:              logger,
	}
}

func (svc *candidateJobSvcImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, note string) (*ent.CandidateJobResponse, error) {
	var candidateJob *ent.CandidateJob
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, uuid.MustParse(input.CandidateID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJob().ValidStatus(ctx, uuid.MustParse(input.CandidateID), uuid.Nil, &input.Status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJob, err = repoRegistry.CandidateJob().CreateCandidateJob(ctx, input)
		if err != nil {
			return err
		}
		_, err := svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		if err != nil {
			return err
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidateJob.Status, candidateJob.ID, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, candidateJob.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobStatus(ctx context.Context, input ent.UpdateCandidateJobStatus, id uuid.UUID, note string) (*ent.CandidateJobResponse, error) {
	record, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if record.Edges.HiringJobEdge.Status == hiringjob.StatusClosed {
		return nil, util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, record.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJob().ValidStatus(ctx, record.CandidateID, id, &input.Status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err := repoRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, result.Status, result.ID, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID, note string) (*ent.CandidateJobResponse, error) {
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusClosed {
		return nil, util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, candidateJob.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = svc.attachmentSvc.RemoveAttachment(ctx, candidateJob.ID, repoRegistry)
		if err != nil {
			return err
		}
		_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailUpdate(candidateJob, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error) {
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) DeleteCandidateJob(ctx context.Context, id uuid.UUID, note string) error {
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusClosed && candidateJob.Status != candidatejob.StatusApplied {
		return util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	// if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusOpened && !ent.CandidateJobStatusEnded.IsValid(ent.CandidateJobStatusEnded(candidateJob.Status)) {
	// 	return util.WrapGQLError(ctx, "model.candidate_job.validation.status_is_invalid_to_delete", http.StatusBadRequest, util.ErrorFlagValidateFail)
	// }
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, candidateJob.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.CandidateJob().DeleteCandidateJob(ctx, candidateJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailCreate(candidateJob)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, candidateJob.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	var result *ent.CandidateJobResponseGetAll
	var edges []*ent.CandidateJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.CandidateIDEQ(uuid.MustParse(filter.CandidateID)))
	svc.filter(ctx, query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidatejob.FieldCreatedAt)
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
	candidateJobs, err := svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateJobs, func(candidateJob *ent.CandidateJob, index int) *ent.CandidateJobEdge {
		return &ent.CandidateJobEdge{
			Node: candidateJob,
			Cursor: ent.Cursor{
				Value: candidateJob.ID.String(),
			},
		}
	})
	result = &ent.CandidateJobResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc candidateJobSvcImpl) GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput,
	filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (
	*ent.CandidateJobGroupByStatusResponse, error) {
	var result *ent.CandidateJobGroupByStatusResponse
	var edges *ent.CandidateJobGroupByStatus
	var page int
	var perPage int
	var err error
	var candidateJobs []*ent.CandidateJob
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(
		candidatejob.HasCandidateEdgeWith(
			candidate.DeletedAtIsNil(), candidate.IsBlacklist(false),
		), candidatejob.HasHiringJobEdgeWith(
			hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusOpened),
		))
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	svc.customFilter(query, filter)
	svc.customFreeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if orderBy != nil {
		if ent.CandidateJobOrderByAdditionalField.IsValid(ent.CandidateJobOrderByAdditionalField(orderBy.Field.String())) {
			candidateJobs, err = svc.getListByAdditionalOrder(ctx, query, page, perPage, orderBy)
			if err != nil {
				return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
		} else {
			candidateJobs, err = svc.getListByNormalOrder(ctx, query, page, perPage, orderBy)
			if err != nil {
				return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
		}
	} else {
		candidateJobs, err = svc.repoRegistry.CandidateJob().BuildList(ctx, query)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		// Split slice by page and perPage
		if page != 0 && perPage != 0 {
			if len(candidateJobs) == 0 {
				candidateJobs = nil
			} else {
				start := (page - 1) * perPage
				end := start + perPage
				if start > len(candidateJobs) {
					candidateJobs = nil
				}
				if start <= len(candidateJobs) && end > len(candidateJobs) {
					candidateJobs = candidateJobs[start:]
				}
				candidateJobs = candidateJobs[start:end]
			}
		}
	}
	edges = &ent.CandidateJobGroupByStatus{
		Hired: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusHired
		}),
		Kiv: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusKiv
		}),
		OfferLost: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusOfferLost
		}),
		Offering: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusOffering
		}),
		ExStaff: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusExStaff
		}),
		Applied: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusApplied
		}),
		Interviewing: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusInterviewing
		}),
	}
	result = &ent.CandidateJobGroupByStatusResponse{
		Data: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobGroupByInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateJobGroupByInterviewResponse, error) {
	var edges *ent.CandidateJobGroupByInterview
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.IDEQ(id)).WithCandidateJobInterview(
		func(query *ent.CandidateInterviewQuery) {
			query.Where(candidateinterview.DeletedAtIsNil()).WithCreatedByEdge().WithInterviewerEdges().WithCandidateJobEdge().
				Order(ent.Desc(candidateinterview.FieldCreatedAt))
		},
	).WithCandidateJobFeedback(
		func(query *ent.CandidateJobFeedbackQuery) {
			query.Where(candidatejobfeedback.DeletedAtIsNil()).WithAttachmentEdges(
				func(query *ent.AttachmentQuery) {
					query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobFeedbacks))
				},
			).WithCreatedByEdge().Order(ent.Desc(candidatejobfeedback.FieldCreatedAt))
		},
	)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = &ent.CandidateJobGroupByInterview{
		Applied: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusApplied
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusApplied
			}),
		},
		Interviewing: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusInterviewing
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusInterviewing
			}),
		},
		Offering: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusOffering
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusOffering
			}),
		},
		Hired: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusHired
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusHired
			}),
		},
		OfferLost: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusOfferLost
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusOfferLost
			}),
		},
		Kiv: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusKiv
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusKiv
			}),
		},
		ExStaff: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusExStaff
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusExStaff
			}),
		},
	}
	result := &ent.CandidateJobGroupByInterviewResponse{
		Data: edges,
	}
	return result, nil
}

// resolver
func (svc *candidateJobSvcImpl) GetCandidateStatus(ctx context.Context, id uuid.UUID) ent.CandidateStatusEnum {
	var candidateJobs []*ent.CandidateJob
	var err error
	openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(s)
	})
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.CandidateIDEQ(id)).Order(ent.Desc(candidatejob.FieldUpdatedAt)).Limit(1)
	candidateJobs, err = svc.repoRegistry.CandidateJob().BuildList(ctx, query.Clone().Where(
		candidatejob.StatusIn(openStatus...),
	))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return ent.CandidateStatusEnumNew
	}
	if len(candidateJobs) == 0 {
		candidateJobs, err = svc.repoRegistry.CandidateJob().BuildList(ctx, query.Clone().Where(
			candidatejob.StatusNotIn(openStatus...),
		))
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return ent.CandidateStatusEnumNew
		}
		if len(candidateJobs) == 0 {
			return ent.CandidateStatusEnumNew
		}
	}
	return ent.CandidateStatusEnum(candidateJobs[0].Status)
}

// common function

func (svc candidateJobSvcImpl) getListByNormalOrder(ctx context.Context, query *ent.CandidateJobQuery, page int, perPage int, orderBy *ent.CandidateJobByOrder) ([]*ent.CandidateJob, error) {
	order := ent.Desc(strings.ToLower(orderBy.Field.String()))
	if orderBy.Direction == ent.OrderDirectionAsc {
		order = ent.Asc(strings.ToLower(orderBy.Field.String()))
	}
	query = query.Order(order)
	if page != 0 && perPage != 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	candidateJobs, err := svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	return candidateJobs, nil
}

func (svc candidateJobSvcImpl) getListByAdditionalOrder(ctx context.Context, query *ent.CandidateJobQuery, page int, perPage int, orderBy *ent.CandidateJobByOrder) ([]*ent.CandidateJob, error) {
	candidateJobs, err := svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	switch orderBy.Field {
	case ent.CandidateJobOrderByFieldPriority:
		sort.Slice(candidateJobs, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return candidateJobs[i].Edges.HiringJobEdge.Priority < candidateJobs[j].Edges.HiringJobEdge.Priority
			} else {
				return candidateJobs[i].Edges.HiringJobEdge.Priority > candidateJobs[j].Edges.HiringJobEdge.Priority
			}
		})
	case ent.CandidateJobOrderByFieldCreatedAt:
		sort.Slice(candidateJobs, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return candidateJobs[i].CreatedAt.Before(candidateJobs[j].CreatedAt)
			} else {
				return candidateJobs[i].CreatedAt.After(candidateJobs[j].CreatedAt)
			}
		})
	}
	// Split slice by page and perPage
	if page != 0 && perPage != 0 {
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(candidateJobs) {
			return nil, nil
		}
		if start <= len(candidateJobs) && end > len(candidateJobs) {
			return candidateJobs[start:], nil
		}
		candidateJobs = candidateJobs[start:end]
	}
	return candidateJobs, nil
}

func (svc *candidateJobSvcImpl) freeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobFreeWord) {
	var predicate []predicate.CandidateJob
	if input != nil {
		if input.Team != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.HasTeamEdgeWith(
					team.NameEqualFold(strings.TrimSpace(*input.Team)),
				),
			))
		}
		if input.HiringJob != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.NameEqualFold(strings.TrimSpace(*input.HiringJob)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predicate...))
	}
}

func (svc *candidateJobSvcImpl) filter(ctx context.Context, candidateJobQuery *ent.CandidateJobQuery, input ent.CandidateJobFilter) {
	if input.Status != nil {
		candidateJobQuery.Where(candidatejob.StatusEQ(candidatejob.Status(*input.Status)))
	}
	if input.TeamID != nil {
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasTeamEdgeWith(
				team.IDEQ(uuid.MustParse(*input.TeamID)),
			),
		))
	}
	if input.HiringJobID != nil {
		candidateJobQuery.Where(candidatejob.HiringJobIDEQ(uuid.MustParse(*input.HiringJobID)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
	if input.FailedReason != nil && len(input.FailedReason) != 0 {
		candidateJobIds := []uuid.UUID{}
		queryString := "SELECT id FROM candidate_jobs WHERE "
		for i, reason := range input.FailedReason {
			queryString += "failed_reason @> '[\"" + reason.String() + "\"]'::jsonb"
			if i != len(input.FailedReason)-1 {
				queryString += " AND "
			}
		}
		queryString += ";"
		rows, _ := candidateJobQuery.QueryContext(ctx, queryString)
		if rows != nil {
			for rows.Next() {
				var id uuid.UUID
				rows.Scan(&id)
				candidateJobIds = append(candidateJobIds, id)
			}
			candidateJobQuery.Where(candidatejob.IDIn(candidateJobIds...))
		} else {
			candidateJobQuery.Where(candidatejob.IDEQ(uuid.Nil))
		}
	}
}

func (svc *candidateJobSvcImpl) customFreeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobGroupByStatusFreeWord) {
	var predicate []predicate.CandidateJob
	if input != nil {
		if input.JobTitle != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.NameContainsFold(strings.TrimSpace(*input.JobTitle)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predicate...))
	}
}

func (svc *candidateJobSvcImpl) customFilter(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobGroupByStatusFilter) {
	if input == nil {
		return
	}
	if input.HiringJobID != nil {
		hiringJobIds := lo.Map(input.HiringJobID, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HiringJobIDIn(hiringJobIds...))
	}
	if input.TeamID != nil {
		teamIds := lo.Map(input.HiringJobID, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.TeamIDIn(teamIds...)))
	}
	if input.Priority != nil {
		priority := lo.Map(input.Priority, func(id int, index int) int {
			return id
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.PriorityIn(priority...)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
}

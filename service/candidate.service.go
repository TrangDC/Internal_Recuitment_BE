package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/entityskill"
	"trec/ent/predicate"
	"trec/ent/skill"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateService interface {
	// mutation
	CreateCandidate(ctx context.Context, input *ent.NewCandidateInput, note string) (*ent.CandidateResponse, error)
	UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID, note string) (*ent.CandidateResponse, error)
	DeleteCandidate(ctx context.Context, id uuid.UUID, note string) error
	SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool, note string) error
	// query
	GetCandidate(ctx context.Context, id uuid.UUID) (*ent.CandidateResponse, error)
	GetCandidates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord, filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord, filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateSelectionResponseGetAll, error)

	//resolved
	GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType
}

type candidateSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	dtoRegistry   dto.Dto
	logger        *zap.Logger
}

func NewCandidateService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateService {
	return &candidateSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		dtoRegistry:   dtoRegistry,
		logger:        logger,
	}
}

func (svc *candidateSvcImpl) CreateCandidate(ctx context.Context, input *ent.NewCandidateInput, note string) (*ent.CandidateResponse, error) {
	var record *ent.Candidate
	errString, err := svc.repoRegistry.Candidate().ValidEmail(ctx, uuid.Nil, input.Email)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if input.ReferenceType == ent.CandidateReferenceTypeRec && input.ReferenceUID == "" {
		return nil, util.WrapGQLError(ctx, "model.candidates.validation.reference_uid_required", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.Candidate().CreateCandidate(ctx, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, record.ID, input.Attachments, nil, attachment.RelationTypeCandidates)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().CreateAndUpdateEntitySkill(ctx, record.ID, input.EntitySkillRecords, nil, entityskill.EntityTypeCandidate)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.Candidate().GetCandidate(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.Candidate().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) DeleteCandidate(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if len(record.Edges.CandidateJobEdges) > 0 {
		openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
			return candidatejob.Status(s)
		})
		candidateJobProcessing := lo.Filter(record.Edges.CandidateJobEdges, func(candidateJob *ent.CandidateJob, i int) bool {
			return lo.Contains(openStatus, candidateJob.Status)
		})
		if len(candidateJobProcessing) > 0 {
			return util.WrapGQLError(ctx, "model.candidates.validation.candidate_job_exist", http.StatusBadRequest, util.ErrorFlagValidateFail)
		}
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.Candidate().DeleteCandidate(ctx, record)
		if err != nil {
			return err
		}
		err = svc.attachmentSvc.RemoveAttachment(ctx, record.ID, repoRegistry)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().DeleteAllEntitySkill(ctx, record.ID)
		return err
	})
	jsonString, err := svc.dtoRegistry.Candidate().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateSvcImpl) UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID, note string) (*ent.CandidateResponse, error) {
	var result *ent.Candidate
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.Candidate().ValidEmail(ctx, id, input.Email)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if input.ReferenceType == ent.CandidateReferenceTypeRec && input.ReferenceUID == "" {
		return nil, util.WrapGQLError(ctx, "model.candidates.validation.reference_uid_required", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.Candidate().UpdateCandidate(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, record.ID, input.Attachments, record.Edges.AttachmentEdges, attachment.RelationTypeCandidates)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().CreateAndUpdateEntitySkill(ctx, record.ID, input.EntitySkillRecords, record.Edges.CandidateSkillEdges, entityskill.EntityTypeCandidate)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.Candidate().GetCandidate(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.Candidate().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool, note string) error {
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Candidate().SetBlackListCandidate(ctx, record, isBlackList)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.Candidate().GetCandidate(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.Candidate().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateSvcImpl) GetCandidate(ctx context.Context, id uuid.UUID) (*ent.CandidateResponse, error) {
	result, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) GetCandidates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord,
	filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error) {
	var result *ent.CandidateResponseGetAll
	var edges []*ent.CandidateEdge
	var page int
	var perPage int
	query := svc.repoRegistry.Candidate().BuildQuery()
	candidates, count, page, perPage, err := svc.getCandidates(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(candidates, func(candidate *ent.Candidate, index int) *ent.CandidateEdge {
		return &ent.CandidateEdge{
			Node: candidate,
			Cursor: ent.Cursor{
				Value: candidate.ID.String(),
			},
		}
	})
	result = &ent.CandidateResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *candidateSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord,
	filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateSelectionResponseGetAll, error) {
	var result *ent.CandidateSelectionResponseGetAll
	var edges []*ent.CandidateSelectionEdge
	var page int
	var perPage int
	query := svc.repoRegistry.Candidate().BuildBaseQuery()
	candidates, count, page, perPage, err := svc.getCandidates(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(candidates, func(candidate *ent.Candidate, index int) *ent.CandidateSelectionEdge {
		return &ent.CandidateSelectionEdge{
			Node: &ent.CandidateSelection{
				ID:   candidate.ID.String(),
				Name: candidate.Name,
			},
			Cursor: ent.Cursor{
				Value: candidate.ID.String(),
			},
		}
	})
	result = &ent.CandidateSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc candidateSvcImpl) GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType {
	return svc.dtoRegistry.EntitySkill().GroupSkillType(input)
}

func (svc *candidateSvcImpl) getCandidates(ctx context.Context, query *ent.CandidateQuery, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord,
	filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) ([]*ent.Candidate, int, int, int, error) {
	var page int
	var perPage int
	svc.filter(ctx, query, filter)
	svc.freeWord(query, freeWord)
	if filter != nil && filter.JobID != nil {
		if filter.IsAbleToInterview != nil && *filter.IsAbleToInterview {
			query = query.Where(candidate.HasCandidateJobEdgesWith(
				candidatejob.HiringJobIDEQ(uuid.MustParse(*filter.JobID)),
				candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing),
			))
		} else {
			query = query.Where(candidate.HasCandidateJobEdgesWith(candidatejob.HiringJobIDEQ(uuid.MustParse(*filter.JobID))))
		}
	}
	count, err := svc.repoRegistry.Candidate().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidate.FieldCreatedAt)
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
	candidates, err := svc.repoRegistry.Candidate().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return candidates, count, page, perPage, nil
}

// common function
func (svc *candidateSvcImpl) freeWord(candidateQuery *ent.CandidateQuery, input *ent.CandidateFreeWord) {
	predicate := []predicate.Candidate{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, candidate.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
		if input.Email != nil {
			predicate = append(predicate, candidate.EmailContainsFold(strings.TrimSpace(*input.Email)))
		}
		if input.Phone != nil {
			predicate = append(predicate, candidate.PhoneContainsFold(strings.TrimSpace(*input.Phone)))
		}
	}
	if len(predicate) > 0 {
		candidateQuery.Where(candidate.Or(predicate...))
	}
}

func (svc *candidateSvcImpl) filter(ctx context.Context, candidateQuery *ent.CandidateQuery, input *ent.CandidateFilter) {
	if input != nil {
		if input.Name != nil {
			candidateQuery.Where(candidate.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.Email != nil {
			candidateQuery.Where(candidate.EmailEqualFold(strings.TrimSpace(*input.Email)))
		}
		if input.Phone != nil {
			candidateQuery.Where(candidate.PhoneEqualFold(strings.TrimSpace(*input.Phone)))
		}
		if input.DobFromDate != nil && input.DobToDate != nil {
			candidateQuery.Where(candidate.DobGTE(*input.DobFromDate), candidate.DobLTE(*input.DobToDate))
		}
		if input.IsBlackList != nil {
			candidateQuery.Where(candidate.IsBlacklist(*input.IsBlackList))
		}
		if input.FromDate != nil && input.ToDate != nil {
			candidateQuery.Where(candidate.CreatedAtGTE(*input.FromDate), candidate.CreatedAtLTE(*input.ToDate))
		}
		if input.IsAbleToInterview != nil && *input.IsAbleToInterview {
			candidateQuery.Where(candidate.HasCandidateJobEdgesWith(
				candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing),
			))
		}
		if input.ReferenceUID != nil {
			memberIds := lo.Map(input.ReferenceUID, func(uid string, index int) uuid.UUID {
				return uuid.MustParse(uid)
			})
			candidateQuery.Where(candidate.ReferenceUIDIn(memberIds...))
		}
		if input.RecruitTimeFromDate != nil && input.RecruitTimeToDate != nil {
			candidateQuery.Where(candidate.RecruitTimeGTE(*input.RecruitTimeFromDate), candidate.RecruitTimeLTE(*input.RecruitTimeToDate))
		}
		if input.Status != nil {
			if ent.CandidateStatusEnum(input.Status.String()) == ent.CandidateStatusEnumNew {
				candidates, _ := svc.repoRegistry.Candidate().BuildList(ctx,
					svc.repoRegistry.Candidate().BuildQuery())
				candidateWithoutJobs := lo.Filter(candidates, func(entity *ent.Candidate, i int) bool {
					return len(entity.Edges.CandidateJobEdges) == 0
				})
				candidateQuery.Where(candidate.IDIn(lo.Map(candidateWithoutJobs, func(entity *ent.Candidate, i int) uuid.UUID {
					return entity.ID
				})...))
			} else {
				if ent.CandidateJobStatusOpen.IsValid(ent.CandidateJobStatusOpen(input.Status.String())) {
					candidateQuery.Where(candidate.HasCandidateJobEdgesWith(
						candidatejob.StatusEQ(candidatejob.Status(*input.Status)), candidatejob.DeletedAtIsNil(),
					))
				} else {
					candidateStatusOpen := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
						return candidatejob.Status(s)
					})
					candidates, _ := svc.repoRegistry.Candidate().BuildList(ctx,
						svc.repoRegistry.Candidate().BuildQuery().Where(candidate.HasCandidateJobEdgesWith(
							candidatejob.DeletedAtIsNil(),
						)))
					candidates = lo.Filter(candidates, func(entity *ent.Candidate, i int) bool {
						return len(lo.Filter(entity.Edges.CandidateJobEdges, func(candidateJob *ent.CandidateJob, j int) bool {
							return lo.Contains(candidateStatusOpen, candidateJob.Status)
						})) == 0
					})
					candidateJobEQStatus := lo.Filter(candidates, func(entity *ent.Candidate, i int) bool {
						return entity.Edges.CandidateJobEdges[0].Status.String() == input.Status.String()
					})
					candidateQuery.Where(candidate.IDIn(lo.Map(candidateJobEQStatus, func(entity *ent.Candidate, i int) uuid.UUID {
						return entity.ID
					})...))
				}
			}
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
			rows, _ := candidateQuery.QueryContext(ctx, queryString)
			if rows != nil {
				for rows.Next() {
					var id uuid.UUID
					rows.Scan(&id)
					candidateJobIds = append(candidateJobIds, id)
				}
				candidateQuery.Where(candidate.HasCandidateJobEdgesWith(
					candidatejob.IDIn(candidateJobIds...),
				))
			} else {
				candidateQuery.Where(candidate.IDIn(uuid.Nil))
			}
		}
		if input.SkillTypeIds != nil {
			ids := lo.Map(input.SkillTypeIds, func(skillType string, index int) uuid.UUID {
				return uuid.MustParse(skillType)
			})
			candidateQuery.Where(candidate.HasCandidateSkillEdgesWith(
				entityskill.HasSkillEdgeWith(skill.SkillTypeIDIn(ids...)),
			))
		}
		if input.SkillIds != nil {
			ids := lo.Map(input.SkillIds, func(skillType string, index int) uuid.UUID {
				return uuid.MustParse(skillType)
			})
			candidateQuery.Where(candidate.HasCandidateSkillEdgesWith(
				entityskill.HasSkillEdgeWith(skill.IDIn(ids...)),
			))
		}
		if input.ReferenceType != nil {
			referenceTypes := lo.Map(input.ReferenceType, func(referenceType ent.CandidateReferenceType, index int) candidate.ReferenceType {
				return candidate.ReferenceType(referenceType)
			})
			candidateQuery.Where(candidate.ReferenceTypeIn(referenceTypes...))
		}
	}
}

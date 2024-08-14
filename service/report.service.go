package service

import (
	"context"
	"net/http"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/internal/util"
	"trec/repository"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ReportService interface {
	ReportCandidateLCC(ctx context.Context) (*ent.ReportCandidateLCCResponse, error)
	ReportCandidateColumnChart(ctx context.Context, filter ent.ReportFilter) (*ent.ReportCandidateColumnChartResponse, error)
	ReportApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationResponse, error)
	ReportApplicationReportTable(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationReportTableResponse, error)
	ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error)
	ReportCandidateConversionRateTable(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error)
}

type reportSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewReportService(repoRegistry repository.Repository, logger *zap.Logger) ReportService {
	return &reportSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc reportSvcImpl) ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error) {
	result, err := svc.repoRegistry.Report().CandidateJobConversion(ctx, nil, uuid.Nil, "")
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.ReportCandidateConversionRateChartResponse{
		Data: result}, nil
}

func (svc reportSvcImpl) ReportCandidateConversionRateTable(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error) {
	var (
		page    int
		perPage int
		results []*ent.CandidateConversionRateReportEdge
	)
	query := svc.repoRegistry.HiringTeam().BuildBaseQuery().
		WithHiringTeamJobEdges(func(hrjQ *ent.HiringJobQuery) {
			hrjQ.Where(hiringjob.DeletedAtIsNil()).
				WithCandidateJobEdges(func(cjQ *ent.CandidateJobQuery) {
					cjQ.Where(candidatejob.DeletedAtIsNil())
				})
		})
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	count, err := svc.repoRegistry.HiringTeam().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	orderFunc := ent.Desc(hiringteam.FieldCreatedAt)
	if orderBy != nil {
		orderField := hiringteam.FieldCreatedAt
		if orderBy.Field == ent.ReportOrderByFieldHiringTeamName {
			orderField = hiringteam.FieldName
		}
		orderFunc = ent.Desc(orderField)
		if orderBy.Direction == ent.OrderDirectionAsc {
			orderFunc = ent.Asc(orderField)
		}
	}
	query = query.Order(orderFunc)
	if perPage != 0 && page != 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	hiringTeams, err := svc.repoRegistry.HiringTeam().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	for _, team := range hiringTeams {
		result := &ent.CandidateConversionRateReport{
			ID:             team.ID.String(),
			HiringTeamName: team.Name,
			Applied:        0,
			Interviewing:   0,
			Offering:       0,
			Hired:          0,
		}
		if len(team.Edges.HiringTeamJobEdges) > 0 {
			candidateJobIds := lo.Flatten(lo.Map(lo.Map(team.Edges.HiringTeamJobEdges, func(hrj *ent.HiringJob, _ int) *ent.HiringJob {
				return hrj
			}), func(hrj *ent.HiringJob, _ int) []uuid.UUID {
				return lo.Map(hrj.Edges.CandidateJobEdges, func(cj *ent.CandidateJob, _ int) uuid.UUID {
					return cj.ID
				})
			}))
			result, err = svc.repoRegistry.Report().CandidateJobConversion(ctx, candidateJobIds, team.ID, team.Name)
			if err != nil {
				svc.logger.Error(err.Error(), zap.Error(err))
				return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
		}
		results = append(results, &ent.CandidateConversionRateReportEdge{
			Node: result,
		})
	}
	return &ent.ReportCandidateConversionRateTableResponse{
		Edges: results,
		Pagination: &ent.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   count,
		}}, nil
}

func (svc reportSvcImpl) ReportApplicationReportTable(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationReportTableResponse, error) {
	processingResult, err := svc.getApplicationProcessing(ctx, filter)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	failKivResult, err := svc.repoRegistry.Report().GetApplicationFail(ctx, filter, "kiv")
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	failOfferLostResult, err := svc.repoRegistry.Report().GetApplicationFail(ctx, filter, candidatejob.StatusOfferLost)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.ReportApplicationReportTableResponse{
		Data: &ent.ApplicationReportTable{
			Processing:  &processingResult,
			Kiv:         &failKivResult,
			OfferedLost: &failOfferLostResult,
		}}, nil
}

func (svc reportSvcImpl) ReportCandidateLCC(ctx context.Context) (*ent.ReportCandidateLCCResponse, error) {
	result := &ent.ReportCandidateLCCResponse{}
	var totalCandidate int
	var totalBlackListCandidate int
	var totalNonBlackListCandidate int
	var recruitment ent.ReportRecruitment
	queryString := "select count(*) as total_records, SUM(case when is_blacklist = true then 1 else 0 end ) as blacklist_true_count " +
		" from candidates where deleted_at is null;"
	rows, err := svc.repoRegistry.Candidate().BuildBaseQuery().QueryContext(ctx, queryString)
	if err != nil {
		return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	defer rows.Close()
	if rows != nil {
		rows.Next()
		if err := rows.Scan(&totalCandidate, &totalBlackListCandidate); err != nil {
			return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		totalNonBlackListCandidate = totalCandidate - totalBlackListCandidate
	}
	recruitment, err = svc.repoRegistry.Report().ReportRecruitment(ctx, carbon.Carbon{}, carbon.Carbon{})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.ReportCandidateLCCResponse{
		Data: &ent.ReportCandidateLcc{
			Total:        totalCandidate,
			BlackList:    totalBlackListCandidate,
			NonBlackList: totalNonBlackListCandidate,
			Recruitment:  &recruitment,
		}}, nil
}

func (svc reportSvcImpl) ReportCandidateColumnChart(ctx context.Context, filter ent.ReportFilter) (*ent.ReportCandidateColumnChartResponse, error) {
	result := ent.ReportCandidateColumnChartEdge{
		Node: []*ent.ReportRecruitment{},
	}
	fromDate, toDate, err := svc.repoRegistry.Report().ValidTimeSelect(filter)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	for i := 0; i < 12; i++ {
		filterFormDate := fromDate
		filterEndDate := fromDate
		switch filter.FilterPeriod {
		case ent.ReportFilterPeriodWeek:
			filterFormDate = fromDate.AddWeeks(i)
			filterEndDate = filterFormDate.EndOfWeek()
		case ent.ReportFilterPeriodMonth:
			filterFormDate = fromDate.AddMonths(i)
			filterEndDate = filterFormDate.EndOfMonth()
		case ent.ReportFilterPeriodQuarter:
			filterFormDate = fromDate.AddQuarters(i)
			filterEndDate = filterFormDate.EndOfQuarter()
		case ent.ReportFilterPeriodYear:
			filterFormDate = fromDate.AddYears(i)
			filterEndDate = filterFormDate.EndOfYear()
		}
		if filterFormDate.Gt(toDate) {
			break
		}
		entity, err := svc.repoRegistry.Report().ReportRecruitment(ctx, filterFormDate, filterEndDate)
		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result.Node = append(result.Node, &entity)
	}
	return &ent.ReportCandidateColumnChartResponse{
		Edges: &result,
	}, err
}

func (svc reportSvcImpl) ReportApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationResponse, error) {
	result := &ent.ReportApplicationEdge{
		Node: []*ent.ReportApplication{},
	}
	fromDate, toDate, err := svc.repoRegistry.Report().ValidTimeSelect(filter)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	for i := 0; i < 12; i++ {
		filterFormDate := fromDate
		filterEndDate := fromDate
		switch filter.FilterPeriod {
		case ent.ReportFilterPeriodWeek:
			filterFormDate = fromDate.AddWeeks(i)
			filterEndDate = filterFormDate.EndOfWeek()
		case ent.ReportFilterPeriodMonth:
			filterFormDate = fromDate.AddMonths(i)
			filterEndDate = filterFormDate.EndOfMonth()
		case ent.ReportFilterPeriodQuarter:
			filterFormDate = fromDate.AddQuarters(i)
			filterEndDate = filterFormDate.EndOfQuarter()
		case ent.ReportFilterPeriodYear:
			filterFormDate = fromDate.AddYears(i)
			filterEndDate = filterFormDate.EndOfYear()
		}
		if filterFormDate.Gt(toDate) {
			break
		}
		entity, err := svc.repoRegistry.Report().ReportApplication(ctx, filterFormDate, filterEndDate)
		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result.Node = append(result.Node, &entity)
	}
	return &ent.ReportApplicationResponse{
		Edges: result,
	}, nil
}

// common
func (svc reportSvcImpl) getApplicationProcessing(ctx context.Context, filter ent.ReportFilter) (ent.ApplicationReportProcessing, error) {
	result := ent.ApplicationReportProcessing{}
	processingCandidateJobIDs, err := svc.repoRegistry.CandidateJob().BuildIDList(
		ctx,
		svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(
			candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing),
			candidatejob.CreatedAtGTE(filter.FromDate), candidatejob.CreatedAtLTE(filter.ToDate),
		),
	)
	if err != nil {
		return result, err
	}
	statusCount, err := svc.repoRegistry.CandidateInterview().BuildStatusCountByCdJobID(ctx, processingCandidateJobIDs)
	if err != nil {
		return result, err
	}
	for _, count := range statusCount {
		switch count.Status {
		case candidateinterview.StatusInvitedToInterview:
			result.InviteToInterview = count.Count
		case candidateinterview.StatusInterviewing:
			result.Interviewing = count.Count
		case candidateinterview.StatusDone:
			result.Done = count.Count
		case candidateinterview.StatusCancelled:
			result.Cancelled = count.Count
		}
	}
	return result, nil
}

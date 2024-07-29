package service

import (
	"context"
	"fmt"
	"net/http"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/team"
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
	var page int
	var perPage int
	var results []*ent.CandidateConversionRateReportEdge
	query := svc.repoRegistry.HiringTeam().BuildBaseQuery().WithHiringTeamJobEdges(
		func(hrjQ *ent.HiringJobQuery) {
			hrjQ.Where(hiringjob.DeletedAtIsNil()).WithCandidateJobEdges(
				func(cjQ *ent.CandidateJobQuery) {
					cjQ.Where(candidatejob.DeletedAtIsNil())
				},
			)
		},
	)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	count, err := svc.repoRegistry.HiringTeam().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if orderBy == nil {
		query = query.Order(ent.Desc(team.FieldCreatedAt))
	} else {
		fieldName := "created_at"
		switch orderBy.Field {
		case ent.ReportOrderByFieldTeamCreatedAt:
			fieldName = "created_at"
		case ent.ReportOrderByFieldTeamName:
			fieldName = "name"
		}
		if orderBy.Direction == ent.OrderDirectionAsc {
			query = query.Order(ent.Asc(fieldName))
		} else {
			query = query.Order(ent.Desc(fieldName))
		}
	}
	if perPage != 0 && page != 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	hiringTeams, err := svc.repoRegistry.HiringTeam().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	for _, team := range hiringTeams {
		candidateJobIds := lo.Flatten(lo.Map(lo.Map(team.Edges.HiringTeamJobEdges, func(hrj *ent.HiringJob, index int) *ent.HiringJob {
			return hrj
		}), func(hrj *ent.HiringJob, index int) []uuid.UUID {
			return lo.Map(hrj.Edges.CandidateJobEdges, func(cj *ent.CandidateJob, index int) uuid.UUID {
				return cj.ID
			})
		}))
		result, err := svc.repoRegistry.Report().CandidateJobConversion(ctx, candidateJobIds, team.ID, team.Name)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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
	failKivResult, err := svc.repoRegistry.Report().GetApplicationFail(ctx, filter, candidatejob.StatusKiv)
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
	queryString := "select cdi.status as status, count(*) as count from candidate_interviews as cdi " +
		"where cdi.candidate_job_status IN ('%s', '%s') and cdi.deleted_at is null" +
		" and cdi.created_at between '%s' and '%s' group by cdi.status;"
	queryString = fmt.Sprintf(queryString, candidateinterview.CandidateJobStatusApplied, candidateinterview.CandidateJobStatusInterviewing, filter.FromDate.Format("2006-01-02 15:04:05"), filter.ToDate.Format("2006-01-02 15:04:05"))
	rows, err := svc.repoRegistry.CandidateInterview().BuildBaseQuery().QueryContext(ctx, queryString)
	if err != nil {
		return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			var status string
			var count int
			if err := rows.Scan(&status, &count); err != nil {
				return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
			switch status {
			case candidateinterview.StatusInvitedToInterview.String():
				result.InviteToInterview = count
			case candidateinterview.StatusInterviewing.String():
				result.Interviewing = count
			case candidateinterview.StatusDone.String():
				result.Done = count
			case candidateinterview.StatusCancelled.String():
				result.Cancelled = count
			}
		}
		if err := rows.Err(); err != nil {
			return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	return result, nil
}

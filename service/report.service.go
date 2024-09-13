package service

import (
	"context"
	"net/http"
	"time"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/jobposition"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/repository"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ReportService interface {
	// Candidate report
	ReportCandidateLCC(ctx context.Context) (*ent.ReportCandidateLCCResponse, error)
	ReportCandidateColumnChart(ctx context.Context, filter ent.ReportFilter) (*ent.ReportCandidateColumnChartResponse, error)
	// Application report
	ReportApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationResponse, error)
	ReportProcessingApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportProcessingApplicationResponse, error)
	ReportFailedApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportFailedApplicationResponse, error)
	ReportHiredApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportHiredApplicationResponse, error)
	// Candidate conversion rate report
	ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error)
	ReportCdConvRateByHiringTeam(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error)
	ReportCdConvRateByJobPosition(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error)
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
	result, err := svc.repoRegistry.Report().CandidateJobConversion(ctx, uuid.Nil, uuid.Nil)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.ReportCandidateConversionRateChartResponse{
		Data: result}, nil
}

func (svc reportSvcImpl) ReportCdConvRateByHiringTeam(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error) {
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
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
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
		result, err := svc.repoRegistry.Report().CandidateJobConversion(ctx, team.ID, uuid.Nil)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result.HiringTeamName = team.Name
		results = append(results, &ent.CandidateConversionRateReportEdge{Node: result})
	}
	return &ent.ReportCandidateConversionRateTableResponse{
		Edges: results,
		Pagination: &ent.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   count,
		}}, nil
}

func (svc *reportSvcImpl) ReportCdConvRateByJobPosition(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error) {
	results := make([]*ent.CandidateConversionRateReportEdge, 0)
	query := svc.repoRegistry.JobPosition().BuildBaseQuery().
		WithHiringJobPositionEdges(func(query *ent.HiringJobQuery) {
			query.Where(hiringjob.DeletedAtIsNil()).
				WithCandidateJobEdges(func(query *ent.CandidateJobQuery) {
					query.Where(candidatejob.DeletedAtIsNil())
				})
		})
	count, err := svc.repoRegistry.JobPosition().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	orderFunc := ent.Desc(jobposition.FieldCreatedAt)
	if orderBy != nil {
		orderField := jobposition.FieldCreatedAt
		if orderBy.Field == ent.ReportOrderByFieldJobPositionName {
			orderField = jobposition.FieldName
		}
		orderFunc = ent.Desc(orderField)
		if orderBy.Direction == ent.OrderDirectionAsc {
			orderFunc = ent.Asc(orderField)
		}
	}
	query = query.Order(orderFunc)
	page, perPage := 0, 0
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	if perPage > 0 && page > 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	jobPositions, err := svc.repoRegistry.JobPosition().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}

	for _, position := range jobPositions {
		result, err := svc.repoRegistry.Report().CandidateJobConversion(ctx, uuid.Nil, position.ID)
		if err != nil {
			svc.logger.Error(err.Error())
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result.JobPositionName = position.Name
		results = append(results, &ent.CandidateConversionRateReportEdge{Node: result})
	}
	return &ent.ReportCandidateConversionRateTableResponse{
		Edges: results,
		Pagination: &ent.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   count,
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
		entity, err := svc.repoRegistry.Report().ReportApplication(ctx, filterFormDate, filterEndDate, filter.HiringTeamID)
		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result.Node = append(result.Node, &entity)
	}
	return &ent.ReportApplicationResponse{
		Edges: result,
	}, nil
}

func (svc *reportSvcImpl) ReportProcessingApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportProcessingApplicationResponse, error) {
	result := make([]*ent.ReportProcessingApplicationEdge, 0)
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
		entity, err := svc.getApplicationProcessing(ctx, filterFormDate, filterEndDate, filter.HiringTeamID)
		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		result = append(result, &ent.ReportProcessingApplicationEdge{Node: entity})
	}
	return &ent.ReportProcessingApplicationResponse{Edges: result}, nil
}

func (svc *reportSvcImpl) ReportFailedApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportFailedApplicationResponse, error) {
	result := &ent.ReportFailedApplicationResponse{
		Data: &ent.ReportFailedApplication{
			FailedCv:        &ent.ApplicationReportFailReason{},
			FailedInterview: &ent.ApplicationReportFailReason{},
			OfferLost:       &ent.ApplicationReportFailReason{},
		},
	}
	var err error
	result.Data.FailedCv, err = svc.repoRegistry.Report().GetApplicationFail(ctx, filter, candidatejob.StatusFailedCv)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result.Data.FailedInterview, err = svc.repoRegistry.Report().GetApplicationFail(ctx, filter, candidatejob.StatusFailedInterview)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result.Data.OfferLost, err = svc.repoRegistry.Report().GetApplicationFail(ctx, filter, candidatejob.StatusOfferLost)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return result, nil
}

func (svc *reportSvcImpl) ReportHiredApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportHiredApplicationResponse, error) {
	jobPositions, err := svc.repoRegistry.JobPosition().BuildList(
		ctx,
		svc.repoRegistry.JobPosition().BuildBaseQuery().
			WithHiringJobPositionEdges(func(query *ent.HiringJobQuery) {
				query.Where(hiringjob.DeletedAtIsNil())
				if filter.HiringTeamID != nil {
					query.Where(hiringjob.HiringTeamID(uuid.MustParse(*filter.HiringTeamID)))
				}
				query.WithCandidateJobEdges(func(query *ent.CandidateJobQuery) {
					query.Where(
						candidatejob.DeletedAtIsNil(), candidatejob.StatusIn(candidatejob.StatusHired, candidatejob.StatusExStaff),
						candidatejob.HasCandidateJobStepWith(
							candidatejobstep.CandidateJobStatusEQ(candidatejobstep.CandidateJobStatusHired),
							candidatejobstep.CreatedAtGTE(filter.FromDate), candidatejobstep.CreatedAtLTE(filter.ToDate),
						),
					)
				})
			}).
			Order(ent.Asc(jobposition.FieldName)),
	)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges := make([]*ent.ReportHiredApplicationEdge, 0)
	for _, pos := range jobPositions {
		node := &ent.ReportHiredApplication{
			JobPositionName: pos.Name,
		}
		for _, job := range pos.Edges.HiringJobPositionEdges {
			switch job.Level {
			case hiringjob.LevelIntern:
				node.Intern += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelFresher:
				node.Fresher += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelJunior:
				node.Junior += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelMiddle:
				node.Middle += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelSenior:
				node.Senior += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelManager:
				node.Manager += len(job.Edges.CandidateJobEdges)
			case hiringjob.LevelDirector:
				node.Director += len(job.Edges.CandidateJobEdges)
			}
		}
		edges = append(edges, &ent.ReportHiredApplicationEdge{Node: node})
	}
	return &ent.ReportHiredApplicationResponse{Edges: edges}, nil
}

// common
func (svc reportSvcImpl) getApplicationProcessing(ctx context.Context, fromDate, toDate carbon.Carbon, hiringTeamIDStr *string) (*ent.ReportProcessingApplication, error) {
	utc, _ := time.LoadLocation(carbon.UTC)
	stdFromDate := fromDate.SetLocation(utc).StdTime()
	stdToDate := toDate.SetLocation(utc).StdTime()
	result := &ent.ReportProcessingApplication{
		FromDate:        stdFromDate,
		ToDate:          stdToDate,
		ActualInterview: 0,
		Cancel:          0,
	}
	predicates := []predicate.CandidateJob{
		candidatejob.StatusEQ(candidatejob.StatusInterviewing),
		candidatejob.CreatedAtGTE(stdFromDate), candidatejob.CreatedAtLTE(stdToDate),
	}
	if hiringTeamIDStr != nil {
		predicates = append(predicates, candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(uuid.MustParse(*hiringTeamIDStr))))
	}
	processingCandidateJobIDs, err := svc.repoRegistry.CandidateJob().BuildIDList(ctx, svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(predicates...))
	if err != nil {
		return result, err
	}
	statusCount, err := svc.repoRegistry.CandidateInterview().BuildStatusCountByCdJobID(ctx, processingCandidateJobIDs)
	if err != nil {
		return result, err
	}
	for _, count := range statusCount {
		switch count.Status {
		case candidateinterview.StatusInvitedToInterview, candidateinterview.StatusInterviewing, candidateinterview.StatusDone:
			result.ActualInterview += count.Count
		case candidateinterview.StatusCancelled:
			result.Cancel += count.Count
		}
	}
	return result, nil
}

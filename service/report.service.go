package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/team"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ReportService interface {
	GetCandidateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateReportResponse, error)
	GetRecruitmentReport(ctx context.Context, filter ent.ReportFilter) (*ent.RecruitmentReportResponse, error)
	GetCandidateConversionRateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateConversionRateReportResponse, error)
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

func (svc *reportSvcImpl) GetCandidateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateReportResponse, error) {
	candidates, err := svc.repoRegistry.Candidate().BuildList(ctx, svc.repoRegistry.Candidate().BuildBaseQuery())
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	initCandidateNumByRef := func() []*ent.ReportNumberByType {
		return []*ent.ReportNumberByType{
			{Type: candidate.ReferenceTypeEb.String(), Number: 0},
			{Type: candidate.ReferenceTypeRec.String(), Number: 0},
			{Type: candidate.ReferenceTypeHiringPlatform.String(), Number: 0},
			{Type: candidate.ReferenceTypeReference.String(), Number: 0},
			{Type: candidate.ReferenceTypeHeadhunt.String(), Number: 0},
		}
	}
	candidateNumByRef := lo.SliceToMap(
		initCandidateNumByRef(),
		func(input *ent.ReportNumberByType) (candidate.ReferenceType, *ent.ReportNumberByType) {
			return candidate.ReferenceType(input.Type), input
		},
	)
	candidateStatsByTime := &ent.ReportStatsByTime{
		Total:              0,
		NumberByType:       initCandidateNumByRef(),
		StatsPerTimePeriod: svc.createReportStatsByFilter(filter, initCandidateNumByRef),
	}

	blacklistNum := 0
	for _, candidateRec := range candidates {
		if candidateRec.IsBlacklist {
			blacklistNum++
		}
		candidateNumByRef[candidateRec.ReferenceType].Number++
		_, statsPerPeriodIndex, exists := lo.FindIndexOf(
			candidateStatsByTime.StatsPerTimePeriod,
			func(stats *ent.ReportStatsPerTimePeriod) bool {
				return stats.FromDate.Compare(candidateRec.RecruitTime) <= 0 && stats.ToDate.Compare(candidateRec.RecruitTime) >= 0
			},
		)
		if exists {
			candidateStatsByTime.Total++
			candidateStatsByTime.StatsPerTimePeriod[statsPerPeriodIndex].Total++

			_, numPerTypeIndex, _ := lo.FindIndexOf(
				candidateStatsByTime.NumberByType,
				func(numByType *ent.ReportNumberByType) bool {
					return numByType.Type == candidateRec.ReferenceType.String()
				},
			)
			candidateStatsByTime.NumberByType[numPerTypeIndex].Number++

			_, numPerTypePerPeriodIndex, _ := lo.FindIndexOf(
				candidateStatsByTime.StatsPerTimePeriod[statsPerPeriodIndex].NumberByType,
				func(numByType *ent.ReportNumberByType) bool {
					return numByType.Type == candidateRec.ReferenceType.String()
				},
			)
			candidateStatsByTime.StatsPerTimePeriod[statsPerPeriodIndex].NumberByType[numPerTypePerPeriodIndex].Number++
		}
	}

	result := &ent.CandidateReportResponse{
		Data: &ent.CandidateReport{
			Total:           len(candidates),
			ActiveNumber:    len(candidates) - blacklistNum,
			BlacklistNumber: blacklistNum,
			NumberByRefType: lo.Values(candidateNumByRef),
			StatsByTime:     candidateStatsByTime,
		},
	}

	return result, nil
}

func (svc *reportSvcImpl) GetRecruitmentReport(ctx context.Context, filter ent.ReportFilter) (*ent.RecruitmentReportResponse, error) {
	applicants, err := svc.repoRegistry.CandidateJob().BuildList(
		ctx,
		svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(
			candidatejob.StatusNotIn(candidatejob.StatusExStaff, candidatejob.StatusOffering),
			candidatejob.CreatedAtGTE(filter.FromDate),
			candidatejob.CreatedAtLTE(filter.ToDate),
		),
	)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}

	initApplicantsNumByStatus := func() []*ent.ReportNumberByType {
		return []*ent.ReportNumberByType{
			{Type: candidatejob.StatusApplied.String(), Number: 0},
			{Type: candidatejob.StatusInterviewing.String(), Number: 0},
			{Type: candidatejob.StatusHired.String(), Number: 0},
			{Type: candidatejob.StatusKiv.String(), Number: 0},
			{Type: candidatejob.StatusOfferLost.String(), Number: 0},
		}
	}
	result := &ent.RecruitmentReportResponse{
		Data: &ent.ReportStatsByTime{
			Total:              len(applicants),
			NumberByType:       initApplicantsNumByStatus(),
			StatsPerTimePeriod: svc.createReportStatsByFilter(filter, initApplicantsNumByStatus),
		},
	}

	for _, applicant := range applicants {
		_, statsByPeriodIndex, exists := lo.FindIndexOf(
			result.Data.StatsPerTimePeriod,
			func(stats *ent.ReportStatsPerTimePeriod) bool {
				return stats.FromDate.Compare(applicant.CreatedAt) <= 0 && stats.ToDate.Compare(applicant.CreatedAt) >= 0
			},
		)
		if exists {
			result.Data.StatsPerTimePeriod[statsByPeriodIndex].Total++
			_, numByTypeIndex, _ := lo.FindIndexOf(
				result.Data.NumberByType,
				func(numByType *ent.ReportNumberByType) bool {
					return numByType.Type == applicant.Status.String()
				},
			)
			result.Data.NumberByType[numByTypeIndex].Number++
			_, numByTypePerPeriodIndex, _ := lo.FindIndexOf(
				result.Data.StatsPerTimePeriod[statsByPeriodIndex].NumberByType,
				func(numByType *ent.ReportNumberByType) bool {
					return numByType.Type == applicant.Status.String()
				},
			)
			result.Data.StatsPerTimePeriod[statsByPeriodIndex].NumberByType[numByTypePerPeriodIndex].Number++
		}
	}

	return result, nil
}

func (svc *reportSvcImpl) GetCandidateConversionRateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateConversionRateReportResponse, error) {
	query := svc.repoRegistry.CandidateJob().BuildBaseQuery().
		Where(candidatejob.StatusIn(
			candidatejob.StatusApplied,
			candidatejob.StatusInterviewing,
			candidatejob.StatusOffering,
			candidatejob.StatusHired,
		))
	if filter.FilterPeriod != ent.ReportFilterPeriodAll {
		query.Where(
			candidatejob.CreatedAtGTE(filter.FromDate),
			candidatejob.CreatedAtLTE(filter.ToDate),
		)
	}
	applicants, err := svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}

	applicantsNumByStatus := map[candidatejob.Status]*ent.ReportNumberByType{
		candidatejob.StatusApplied:      {Type: candidatejob.StatusApplied.String(), Number: 0},
		candidatejob.StatusHired:        {Type: candidatejob.StatusHired.String(), Number: 0},
		candidatejob.StatusOffering:     {Type: candidatejob.StatusOffering.String(), Number: 0},
		candidatejob.StatusInterviewing: {Type: candidatejob.StatusInterviewing.String(), Number: 0},
	}
	for _, applicant := range applicants {
		applicantsNumByStatus[applicant.Status].Number++
	}

	result := &ent.CandidateConversionRateReportResponse{
		Data: lo.Values(applicantsNumByStatus),
	}
	return result, nil
}

func (svc *reportSvcImpl) createReportStatsByFilter(filter ent.ReportFilter, createStatsNumberByType func() []*ent.ReportNumberByType) []*ent.ReportStatsPerTimePeriod {
	result := make([]*ent.ReportStatsPerTimePeriod, 0)
	start := filter.FromDate

	// Helper function to calculate the next period date
	getNextPeriodDate := func(start time.Time, years, months, days int) (time.Time, time.Time) {
		temp := start.AddDate(years, months, days)
		toDate := temp.AddDate(0, 0, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		if toDate.After(filter.ToDate) {
			toDate = filter.ToDate
		}
		return temp, toDate
	}

	for start.Before(filter.ToDate) {
		var temp, toDate time.Time

		switch filter.FilterPeriod {
		case ent.ReportFilterPeriodYear:
			temp, toDate = getNextPeriodDate(start, 1, 0, 0)
		case ent.ReportFilterPeriodQuarter:
			temp, toDate = getNextPeriodDate(start, 0, 3, 0)
		case ent.ReportFilterPeriodMonth:
			temp, toDate = getNextPeriodDate(start, 0, 1, 0)
		case ent.ReportFilterPeriodWeek:
			temp, toDate = getNextPeriodDate(start, 0, 0, 7)
		default:
			// Handle unexpected filter period
			result = append(result, &ent.ReportStatsPerTimePeriod{
				FromDate:     time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				ToDate:       time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC),
				Total:        0,
				NumberByType: createStatsNumberByType(),
			})
			return result
		}

		result = append(result, &ent.ReportStatsPerTimePeriod{
			FromDate:     start,
			ToDate:       toDate,
			Total:        0,
			NumberByType: createStatsNumberByType(),
		})
		start = temp
	}

	return result
}

func (svc reportSvcImpl) ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error) {
	result, err := svc.candidateJobConversion(ctx, nil, uuid.Nil, "")
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.ReportCandidateConversionRateChartResponse{
		Data: result}, nil
}

func (svc reportSvcImpl) candidateJobConversion(ctx context.Context, candidateJobIds []uuid.UUID, teamId uuid.UUID, teamName string) (*ent.CandidateConversionRateReport, error) {
	var applied int
	var interviewing int
	var offering int
	var hired int
	queryString := `
	select
		cjs.candidate_job_status,
		COUNT(*) as count
	from
		candidate_job_steps cjs
	join candidate_jobs cj on
		cjs.candidate_job_id = cj.id
	where
		cj.deleted_at is null %s and cjs.candidate_job_status in ('applied', 'interviewing', 'offering', 'hired')
	group by
		cjs.candidate_job_status ;`
	if len(candidateJobIds) > 0 {
		queryString = fmt.Sprintf(queryString, fmt.Sprintf("and cj.id in ('%s')", strings.Join(lo.Map(candidateJobIds, func(id uuid.UUID, index int) string { return id.String() }), "','")))
	} else {
		queryString = fmt.Sprintf(queryString, "")
	}
	rows, err := svc.repoRegistry.CandidateJobStep().BuildQuery().QueryContext(ctx, queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			var status string
			var count int
			if err := rows.Scan(&status, &count); err != nil {
				return nil, err
			}
			switch status {
			case "applied":
				applied = count
			case "interviewing":
				interviewing = count
			case "offering":
				offering = count
			case "hired":
				hired = count
			}
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	resultId := ""
	if teamId != uuid.Nil {
		resultId = teamId.String()
	}
	return &ent.CandidateConversionRateReport{
		ID:           resultId,
		TeamName:     teamName,
		Applied:      applied,
		Interviewing: interviewing,
		Offering:     offering,
		Hired:        hired,
	}, nil
}

func (svc reportSvcImpl) ReportCandidateConversionRateTable(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error) {
	var page int
	var perPage int
	var results []*ent.CandidateConversionRateReportEdge
	query := svc.repoRegistry.Team().BuildBaseQuery().Where().WithTeamJobEdges(
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
	count, err := svc.repoRegistry.Team().BuildCount(ctx, query)
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
	teams, err := svc.repoRegistry.Team().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	for _, team := range teams {
		candidateJobIds := lo.Flatten(lo.Map(lo.Map(team.Edges.TeamJobEdges, func(hrj *ent.HiringJob, index int) *ent.HiringJob {
			return hrj
		}), func(hrj *ent.HiringJob, index int) []uuid.UUID {
			return lo.Map(hrj.Edges.CandidateJobEdges, func(cj *ent.CandidateJob, index int) uuid.UUID {
				return cj.ID
			})
		}))
		result, err := svc.candidateJobConversion(ctx, candidateJobIds, team.ID, team.Name)
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

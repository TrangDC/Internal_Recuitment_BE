package service

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/repository"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ReportService interface {
	GetCandidateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateReportResponse, error)
	GetRecruitmentReport(ctx context.Context, filter ent.ReportFilter) (*ent.RecruitmentReportResponse, error)
	GetCandidateConversionRateReport(ctx context.Context, filter ent.ReportFilter) (*ent.CandidateConversionRateReportResponse, error)
	ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error)
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
		cj.deleted_at is null and cjs.candidate_job_status in ('applied', 'interviewing', 'offering', 'hired')
	group by
		cjs.candidate_job_status ;`
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
	return &ent.ReportCandidateConversionRateChartResponse{
		Data: &ent.CandidateConversionRateReport{
			Applied:      applied,
			Interviewing: interviewing,
			Offering:     offering,
			Hired:        hired,
		}}, nil
}

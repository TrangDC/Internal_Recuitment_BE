package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/internal/util"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ReportRepository interface {
	CandidateJobConversion(ctx context.Context, candidateJobIds []uuid.UUID, teamId uuid.UUID, teamName string) (*ent.CandidateConversionRateReport, error)
	GetApplicationFail(ctx context.Context, filter ent.ReportFilter, status candidatejob.Status) (ent.ApplicationReportFailReason, error)
	ReportRecruitment(ctx context.Context, fromDate, toDate carbon.Carbon) (ent.ReportRecruitment, error)
	ReportApplication(ctx context.Context, fromDate, toDate carbon.Carbon) (ent.ReportApplication, error)
	ValidTimeSelect(filter ent.ReportFilter) (carbon.Carbon, carbon.Carbon, error)
}

type reportRepoImpl struct {
	client *ent.Client
}

func NewReportRepository(client *ent.Client) ReportRepository {
	return &reportRepoImpl{
		client: client,
	}
}

func (rps reportRepoImpl) CandidateJobConversion(ctx context.Context, candidateJobIds []uuid.UUID, teamId uuid.UUID, teamName string) (*ent.CandidateConversionRateReport, error) {
	var applied int
	var interviewing int
	var offering int
	var hired int
	queryString := "select cjs.candidate_job_status, count(*) as count from candidate_job_steps cjs join candidate_jobs cj on cjs.candidate_job_id = cj.id " +
		"where cj.deleted_at is null %s and cjs.candidate_job_status in ('applied', 'interviewing', 'offering', 'hired') group by cjs.candidate_job_status ;"
	if len(candidateJobIds) > 0 {
		queryString = fmt.Sprintf(queryString, fmt.Sprintf("and cj.id in ('%s')", strings.Join(lo.Map(candidateJobIds, func(id uuid.UUID, index int) string { return id.String() }), "','")))
	} else {
		queryString = fmt.Sprintf(queryString, "")
	}
	rows, err := rps.client.CandidateJobStep.Query().QueryContext(ctx, queryString)
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
			case candidatejob.StatusApplied.String():
				applied = count
			case candidatejob.StatusInterviewing.String():
				interviewing = count
			case candidatejob.StatusOffering.String():
				offering = count
			case candidatejob.StatusHired.String():
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

func (rps reportRepoImpl) GetApplicationFail(ctx context.Context, filter ent.ReportFilter, status candidatejob.Status) (ent.ApplicationReportFailReason, error) {
	result := ent.ApplicationReportFailReason{}
	queryString := "select value as failed_reason, count(*) as count from candidate_jobs as cdj, jsonb_array_elements_text(failed_reason) as value " +
		"where cdj.status = '%s' and cdj.deleted_at is null and cdj.created_at between '%s' and '%s' group by value order by count desc;"
	queryString = fmt.Sprintf(
		queryString, status, filter.FromDate.Format("2006-01-02 15:04:05"), filter.ToDate.Format("2006-01-02 15:04:05"))
	rows, err := rps.client.CandidateJob.Query().QueryContext(ctx, queryString)
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
			case ent.CandidateJobFailedReasonPoorProfessionalism.String():
				result.PoorProfessionalism = count
			case ent.CandidateJobFailedReasonPoorFitAndEngagement.String():
				result.PoorFitAndEngagement = count
			case ent.CandidateJobFailedReasonOverExpectations.String():
				result.OverExpectations = count
			case ent.CandidateJobFailedReasonOverQualification.String():
				result.OverQualification = count
			case ent.CandidateJobFailedReasonLanguageDeficiency.String():
				result.LanguageDeficiency = count
			case ent.CandidateJobFailedReasonWeakTechnicalSkills.String():
				result.WeakTechnicalSkills = count
			case ent.CandidateJobFailedReasonPoorInterpersonalSkills.String():
				result.PoorInterpersonalSkills = count
			case ent.CandidateJobFailedReasonPoorProblemSolvingSkills.String():
				result.PoorProblemSolvingSkills = count
			case ent.CandidateJobFailedReasonPoorManagementSkills.String():
				result.CandidateWithdrawal = count
			case ent.CandidateJobFailedReasonCandidateWithdrawal.String():
				result.Others = count
			}
		}
		if err := rows.Err(); err != nil {
			return result, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	return result, nil
}

func (rps reportRepoImpl) ReportRecruitment(ctx context.Context, fromDate, toDate carbon.Carbon) (ent.ReportRecruitment, error) {
	var result ent.ReportRecruitment
	queryString := "select reference_type, count(*) as count from candidates where deleted_at is null " +
		"group by reference_type order by reference_type; "
	if !fromDate.IsZero() || !toDate.IsZero() {
		utc, _ := time.LoadLocation(carbon.UTC)
		fromDateStd := fromDate.SetLocation(utc).StdTime()
		toDateStd := toDate.SetLocation(utc).StdTime()
		queryString = "select reference_type, count(*) as count from candidates where deleted_at is null " +
			"and recruit_time between '%s' and '%s' group by reference_type order by reference_type;"
		queryString = fmt.Sprintf(queryString, fromDate.SetLocation(utc).ToDateTimeString(), toDate.SetLocation(utc).ToDateTimeString())
		result.FromDate = &fromDateStd
		result.ToDate = &toDateStd
	}
	rows, err := rps.client.Candidate.Query().QueryContext(ctx, queryString)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			var typeName any
			var count int
			if err := rows.Scan(&typeName, &count); err != nil {
				return result, err
			}
			switch typeName {
			case candidate.ReferenceTypeEb.String():
				result.Eb = count
			case candidate.ReferenceTypeRec.String():
				result.Rec = count
			case candidate.ReferenceTypeHiringPlatform.String():
				result.HiringPlatform = count
			case candidate.ReferenceTypeReference.String():
				result.Reference = count
			case candidate.ReferenceTypeHeadhunt.String():
				result.Headhunt = count
			}
		}
	}
	return result, nil
}

func (rps reportRepoImpl) ReportApplication(ctx context.Context, fromDate, toDate carbon.Carbon) (ent.ReportApplication, error) {
	var result ent.ReportApplication
	utc, _ := time.LoadLocation(carbon.UTC)
	fromDateStd := fromDate.SetLocation(utc).StdTime()
	toDateStd := toDate.SetLocation(utc).StdTime()
	queryString := "select status, count(*) from candidate_jobs " +
		"where deleted_at is null and created_at between '%s' and '%s' group by status;"
	queryString = fmt.Sprintf(queryString, fromDate.SetLocation(utc).ToDateTimeString(), toDate.SetLocation(utc).ToDateTimeString())
	result.FromDate = fromDateStd
	result.ToDate = toDateStd
	rows, err := rps.client.CandidateJob.Query().QueryContext(ctx, queryString)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			var status string
			var count int
			if err := rows.Scan(&status, &count); err != nil {
				return result, err
			}
			switch status {
			case candidatejob.StatusApplied.String():
				result.Applied = count
			case candidatejob.StatusInterviewing.String():
				result.Interviewing = count
			case candidatejob.StatusOffering.String():
				result.Offering = count
			case candidatejob.StatusHired.String():
				result.Hired = count
			case candidatejob.StatusKiv.String():
				result.Kiv = count
			case candidatejob.StatusOfferLost.String():
				result.OfferLost = count
			case candidatejob.StatusOffering.String():
				result.Offering = count
			case candidatejob.StatusExStaff.String():
				result.ExStaff = count
			}
		}
		if err := rows.Err(); err != nil {
			return result, err
		}
	}
	return result, nil
}

func (rps reportRepoImpl) ValidTimeSelect(filter ent.ReportFilter) (carbon.Carbon, carbon.Carbon, error) {
	fromDate := carbon.Parse(filter.FromDate.String())
	toDate := carbon.Parse(filter.ToDate.String())
	if filter.FromDate.IsZero() || filter.ToDate.IsZero() {
		return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.time_range_required")
	}
	if filter.FromDate.After(filter.ToDate) {
		return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.invalid_time_range")
	}
	switch filter.FilterPeriod {
	case ent.ReportFilterPeriodWeek:
		fromDate = fromDate.StartOfWeek()
		toDate = toDate.EndOfWeek()
		if fromDate.AddWeeks(12).EndOfWeek().Lte(toDate) {
			return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.out_of_range_weeks")
		}
	case ent.ReportFilterPeriodMonth:
		fromDate = fromDate.StartOfMonth()
		toDate = toDate.EndOfMonth()
		if fromDate.AddMonths(12).EndOfMonth().Lte(toDate) {
			return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.out_of_range_months")
		}
	case ent.ReportFilterPeriodQuarter:
		fromDate = fromDate.StartOfQuarter()
		toDate = toDate.EndOfQuarter()
		if fromDate.AddQuarters(12).EndOfQuarter().Lte(toDate) {
			return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.out_of_range_quarters")
		}
	case ent.ReportFilterPeriodYear:
		fromDate = fromDate.StartOfYear()
		toDate = toDate.EndOfYear()
		if fromDate.AddYears(12).EndOfYear().Lte(toDate) {
			return carbon.Carbon{}, carbon.Carbon{}, fmt.Errorf("model.reports.validation.out_of_range_years")
		}
	}
	return fromDate, toDate, nil
}

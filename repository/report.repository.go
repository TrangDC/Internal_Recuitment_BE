package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/models"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
)

type ReportRepository interface {
	CandidateJobConversion(ctx context.Context, hiringTeamID, jobPositionID uuid.UUID) (*ent.CandidateConversionRateReport, error)
	GetApplicationFail(ctx context.Context, filter ent.ReportFilter, status candidatejob.Status) (*ent.ApplicationReportFailReason, error)
	ReportRecruitment(ctx context.Context, fromDate, toDate carbon.Carbon) (ent.ReportRecruitment, error)
	ReportApplication(ctx context.Context, fromDate, toDate carbon.Carbon, hiringTeamIDStr *string) (ent.ReportApplication, error)
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

func (rps reportRepoImpl) CandidateJobConversion(ctx context.Context, hiringTeamID, jobPositionID uuid.UUID) (*ent.CandidateConversionRateReport, error) {
	result := &ent.CandidateConversionRateReport{}
	cdPredicates := []predicate.CandidateJob{candidatejob.DeletedAtIsNil()}
	switch {
	case hiringTeamID != uuid.Nil:
		result.ID = hiringTeamID.String()
		cdPredicates = append(cdPredicates, candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(hiringTeamID)))
	case jobPositionID != uuid.Nil:
		result.ID = jobPositionID.String()
		cdPredicates = append(cdPredicates, candidatejob.HasHiringJobEdgeWith(hiringjob.JobPositionID(jobPositionID)))
	}
	countsByStatus := make([]models.CdJobStepCountByStatus, 0)
	err := rps.client.CandidateJobStep.Query().
		Select(candidatejobstep.FieldCandidateJobStatus).
		Where(
			candidatejobstep.CandidateJobStatusIn(
				candidatejobstep.CandidateJobStatusApplied,
				candidatejobstep.CandidateJobStatusInterviewing,
				candidatejobstep.CandidateJobStatusOffering,
				candidatejobstep.CandidateJobStatusHired,
			),
			candidatejobstep.HasCandidateJobEdgeWith(cdPredicates...),
		).
		GroupBy(candidatejobstep.FieldCandidateJobStatus).Aggregate(ent.Count()).
		Scan(ctx, &countsByStatus)
	if err != nil {
		return nil, err
	}
	for _, v := range countsByStatus {
		switch v.Status {
		case candidatejobstep.CandidateJobStatusApplied:
			result.Applied = v.Count
		case candidatejobstep.CandidateJobStatusInterviewing:
			result.Interviewing = v.Count
		case candidatejobstep.CandidateJobStatusOffering:
			result.Offering = v.Count
		case candidatejobstep.CandidateJobStatusHired:
			result.Hired = v.Count
		}
	}
	return result, nil
}

func (rps reportRepoImpl) GetApplicationFail(ctx context.Context, filter ent.ReportFilter, status candidatejob.Status) (*ent.ApplicationReportFailReason, error) {
	result := &ent.ApplicationReportFailReason{}
	where := "WHERE cdj.status = '%s' AND cdj.deleted_at IS NULL AND cdj.created_at BETWEEN '%s' AND '%s' "
	if filter.HiringTeamID != nil {
		where += fmt.Sprintf(" AND cdj.hiring_job_id IN (SELECT id FROM hiring_jobs WHERE hiring_team_id = '%s') ", *filter.HiringTeamID)
	}
	queryString := "SELECT value AS failed_reason, count(*) AS count " +
		"FROM candidate_jobs AS cdj, jsonb_array_elements_text(failed_reason) AS value " +
		where +
		"GROUP BY value ORDER BY count DESC;"
	queryString = fmt.Sprintf(
		queryString,
		status, filter.FromDate.Format("2006-01-02 15:04:05"), filter.ToDate.Format("2006-01-02 15:04:05"),
	)
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
				result.PoorManagementSkills = count
			case ent.CandidateJobFailedReasonCandidateWithdrawal.String():
				result.CandidateWithdrawal = count
			case ent.CandidateJobFailedReasonOthers.String():
				result.Others = count
			}
		}
		if err := rows.Err(); err != nil {
			return result, err
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

func (rps reportRepoImpl) ReportApplication(ctx context.Context, fromDate, toDate carbon.Carbon, hiringTeamIDStr *string) (ent.ReportApplication, error) {
	utc, _ := time.LoadLocation(carbon.UTC)
	fromDateStd := fromDate.SetLocation(utc).StdTime()
	toDateStd := toDate.SetLocation(utc).StdTime()
	result := ent.ReportApplication{
		FromDate: fromDateStd,
		ToDate:   toDateStd,
	}
	cdJobPredicates := []predicate.CandidateJob{
		candidatejob.DeletedAtIsNil(),
		candidatejob.CreatedAtGTE(fromDateStd), candidatejob.CreatedAtLTE(toDateStd),
	}
	if hiringTeamIDStr != nil {
		cdJobPredicates = append(cdJobPredicates, candidatejob.HasHiringJobEdgeWith(hiringjob.HiringTeamID(uuid.MustParse(*hiringTeamIDStr))))
	}
	countsByStatus := make([]models.CandidateJobCountByStatus, 0)
	err := rps.client.CandidateJob.Query().
		Select(candidatejob.FieldStatus).
		Where(cdJobPredicates...).
		GroupBy(candidatejob.FieldStatus).
		Aggregate(ent.Count()).
		Scan(ctx, &countsByStatus)
	if err != nil {
		return result, err
	}
	for _, v := range countsByStatus {
		switch v.Status {
		case candidatejob.StatusApplied:
			result.Applied = v.Count
		case candidatejob.StatusInterviewing:
			result.Interviewing = v.Count
		case candidatejob.StatusOffering:
			result.Offering = v.Count
		case candidatejob.StatusHired:
			result.Hired = v.Count
		case candidatejob.StatusFailedCv:
			result.FailedCv = v.Count
		case candidatejob.StatusFailedInterview:
			result.FailedInterview = v.Count
		case candidatejob.StatusOfferLost:
			result.OfferLost = v.Count
		case candidatejob.StatusExStaff:
			result.ExStaff = v.Count
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

package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"trec/ent"
	"trec/internal/util"
	"trec/repository"

	"github.com/99designs/gqlgen/graphql"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ImportDataService interface {
	ImportCandidate(ctx context.Context, data graphql.Upload) error
}
type importDataSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewImportDataService(repoRegistry repository.Repository, logger *zap.Logger) ImportDataService {
	return &importDataSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *importDataSvcImpl) ImportCandidate(ctx context.Context, data graphql.Upload) error {
	oldCandidates, err := svc.repoRegistry.Candidate().BuildList(ctx,
		svc.repoRegistry.Candidate().BuildQuery())
	if err != nil {
		return util.WrapGQLError(ctx, "excel.import.error", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateEmails := lo.Map(oldCandidates, func(candidate *ent.Candidate, index int) string {
		return candidate.Email
	})
	rows, err := svc.repoRegistry.ImportData().ReadXlsxFile(data)
	if err != nil {
		return err
	}
	rows = rows[4:]
	newCandidateEmails := lo.Map(rows, func(row []string, index int) string {
		return strings.TrimSpace(row[2])
	})
	for _, v := range rows {
		if len(v) != 5 {
			return util.WrapGQLError(ctx, "excel.import.candidates.invalid_format", http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	if len(lo.FindDuplicates(newCandidateEmails)) != 0 {
		return util.WrapGQLError(ctx, "excel.import.candidates.dublicate_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[1] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_name", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[2] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[3] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_dob", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[4] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.phone_number", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Intersect(candidateEmails, newCandidateEmails)) > 0 {
		return util.WrapGQLError(ctx, "excel.import.candidates.dublicate_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateInputs := lo.Map(rows, func(row []string, index int) *ent.NewCandidateInput {
		dob, err := svc.convertStringDate(row[3])
		if err != nil {
			return nil
		}
		return &ent.NewCandidateInput{
			Name:  row[1],
			Email: row[2],
			Dob:   dob,
			Phone: row[4],
		}
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.Candidate().BuildBulkCreate(ctx, candidateInputs)
		return err
	})
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc importDataSvcImpl) convertStringDate(input string) (time.Time, error) {
	result, err := time.Parse("02/01/06", input)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("02-01-06", input)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("01-02-06", input)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("02-01-2006", input)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("02/01/2006", input)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("06", input)
	if err == nil {
		return result, nil
	}
	dateEpoch, err := strconv.Atoi(input)
	if err != nil {
		return time.Time{}, fmt.Errorf("fail to convert date time")
	}
	excelEpoch := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	result = excelEpoch.AddDate(0, 0, dateEpoch-1)
	return result, nil
}

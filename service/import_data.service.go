package service

import (
	"context"
	"net/http"
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
	newCandidateEmails := lo.Map(rows, func(row []string, index int) string {
		return strings.TrimSpace(row[2])
	})
	if len(rows[0]) != 5 {
		return util.WrapGQLError(ctx, "excel.import.candidates.invalid_format", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	rows = rows[1:]
	if len(lo.FindDuplicates(newCandidateEmails)) != 0 {
		return util.WrapGQLError(ctx, "excel.import.candidates.duplicate_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[1] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_name", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[2] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[3] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.phone_number", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Intersect(candidateEmails, newCandidateEmails)) > 0 {
		return util.WrapGQLError(ctx, "model.candidates.validation.email_exist", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	var candidateInputs []*ent.NewCandidateInput
	for _, row := range rows {
		if len(row) < 4 {
			return util.WrapGQLError(ctx, "excel.import.candidates.invalid_format", http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
		newInput := &ent.NewCandidateInput{
			Name:  row[1],
			Email: row[2],
			Phone: row[3],
			Dob:   &util.DefaultTime,
		}
		if len(row) == 5 {
			dob, err := time.Parse("01-02-06", row[4])
			if err != nil {
				svc.logger.Error("error parsing dob", zap.Error(err))
				return util.WrapGQLError(ctx, "excel.import.candidates.dob_invalid_format", http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
			newInput.Dob = &dob
		}
		candidateInputs = append(candidateInputs, newInput)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.Candidate().BuildBulkCreate(ctx, candidateInputs)
		return err
	})
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

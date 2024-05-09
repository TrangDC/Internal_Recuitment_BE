package service

import (
	"context"
	"net/http"
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
		return row[3]
	})
	if len(lo.FindDuplicates(newCandidateEmails)) != 0 {
		return util.WrapGQLError(ctx, "excel.import.candidates.dublicate_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[2] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_name", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[3] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[4] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.missing_dob", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Map(rows, func(row []string, index int) string { return row[5] })) != len(rows) {
		return util.WrapGQLError(ctx, "excel.import.candidates.phone_number", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if len(lo.Intersect(candidateEmails, newCandidateEmails)) > 0 {
		return util.WrapGQLError(ctx, "excel.import.candidates.dublicate_email", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	createBulk := lo.Map(rows, func(row []string, index int) *ent.NewCandidateInput {
		dob, err := time.Parse(time.Time{}.Format("2006-01-02"), row[4])
		if err != nil {
			return nil
		}
		return &ent.NewCandidateInput{
			Name:  row[2],
			Email: row[3],
			Dob:   dob,
			Phone: row[5],
		}
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.Candidate().BuildBulkCreate(ctx, createBulk)
		return err
	})
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

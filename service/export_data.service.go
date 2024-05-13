package service

import (
	"context"
	"encoding/base64"
	"net/http"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"trec/ent"

	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type ExportDataService interface {
	ExportSampleCandidates(ctx context.Context, lang ent.I18nLanguage) (*ent.Base64Response, error)
}
type exportDataSvcImpl struct {
	repoRegistry repository.Repository
	i18n         models.I18n
	logger       *zap.Logger
}

func NewExportDataService(repoRegistry repository.Repository, i18n models.I18n, logger *zap.Logger) ExportDataService {
	return &exportDataSvcImpl{
		repoRegistry: repoRegistry,
		i18n:         i18n,
		logger:       logger,
	}
}

func (svc *exportDataSvcImpl) ExportSampleCandidates(ctx context.Context, lang ent.I18nLanguage) (*ent.Base64Response, error) {
	var i18nObject models.I18nObject
	switch lang {
	case ent.I18nLanguageEn:
		i18nObject = svc.i18n.En
	case ent.I18nLanguageVi:
		i18nObject = svc.i18n.Vi
	}
	file := svc.candidateHeader(i18nObject)
	excelBytes, err := file.WriteToBuffer()
	if err != nil {
		svc.logger.Error("error while convert excel to base64", zap.Error(err))
		return nil, util.WrapGQLError(ctx, "excel.convert.error", http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	excelBase64 := base64.StdEncoding.EncodeToString(excelBytes.Bytes())
	return &ent.Base64Response{
		Data: excelBase64,
	}, nil
}

func (svc exportDataSvcImpl) candidateHeader(i18nObj models.I18nObject) *excelize.File {
	file := excelize.NewFile()
	sheetName := file.GetSheetName(0)
	mergeCells(file, sheetName, "A2", "C2")
	setCellValue(file, sheetName, "A2", i18nObj.Model.Candidates.ModelName)
	setCellValue(file, sheetName, "A4", i18nObj.Excel.Id)
	setCellValue(file, sheetName, "B4", i18nObj.Model.Candidates.Name)
	setCellValue(file, sheetName, "C4", i18nObj.Model.Candidates.Email)
	setCellValue(file, sheetName, "D4", i18nObj.Model.Candidates.Dob)
	setCellValue(file, sheetName, "E4", i18nObj.Model.Candidates.Phone)
	return file
}

func setCellValue(f *excelize.File, sheet, cell, value string) {
	_ = f.SetCellValue(sheet, cell, value)
}

func mergeCells(f *excelize.File, sheet, cellB string, cellE string) {
	_ = f.MergeCell(sheet, cellB, cellE)
}

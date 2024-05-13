package repository

import (
	"fmt"
	"os"
	"trec/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/xuri/excelize/v2"
)

type ImportDataRepository interface {
	ReadXlsxFile(data graphql.Upload) ([][]string, error)
}

type importDataRepoImpl struct {
	client *ent.Client
}

func NewImportDataRepository(client *ent.Client) ImportDataRepository {
	return &importDataRepoImpl{
		client: client,
	}
}

func (rps importDataRepoImpl) ReadXlsxFile(data graphql.Upload) ([][]string, error) {
	destinationPath := "../" + data.Filename // Use an absolute path within the container
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return nil, fmt.Errorf("error creating destination file: %w", err)
	}
	defer destinationFile.Close()

	xlsx, err := excelize.OpenReader(data.File)
	if err != nil {
		return nil, fmt.Errorf("error opening Excel file: %w", err)
	}
	sheetMap := xlsx.GetSheetMap()
	if len(sheetMap) == 0 {
		return nil, fmt.Errorf("file import is not valid format")
	}
	rows, _ := xlsx.GetRows(sheetMap[1])
	return rows, nil
}

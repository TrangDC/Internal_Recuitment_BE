package repository

import (
	"context"
	"trec/ent"
)

type PreRepository interface {
	PreFunction(ctx context.Context) (string, error)
}

type jobTitleRepoImpl struct {
	client *ent.Client
}

func NewPreRepository(client *ent.Client) PreRepository {
	return &jobTitleRepoImpl{
		client: client,
	}
}

func (rps jobTitleRepoImpl) PreFunction(ctx context.Context) (string, error) {
	return "Success - Repository", nil
}

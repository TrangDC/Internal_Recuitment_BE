package repository

import (
	"trec/ent"
)

type HiringTeamApproverRepository interface {
}

type hiringTeamApproverRepoImpl struct {
	client *ent.Client
}

func NewHiringTeamApproverRepository(client *ent.Client) HiringTeamApproverRepository {
	return &hiringTeamApproverRepoImpl{
		client: client,
	}
}

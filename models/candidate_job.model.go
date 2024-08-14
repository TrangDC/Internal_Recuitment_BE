package models

import (
	"time"
	"trec/ent"
	"trec/ent/candidatejob"

	"github.com/google/uuid"
)

type CandidateJobStep struct {
	Index  int
	Status candidatejob.Status
}

var CandidateJobSteps = []CandidateJobStep{
	{Index: 0, Status: candidatejob.StatusApplied},
	{Index: 1, Status: candidatejob.StatusInterviewing},
	{Index: 2, Status: candidatejob.StatusOffering},
	{Index: 3, Status: candidatejob.StatusHired},
	{Index: 4, Status: candidatejob.StatusOfferLost},
	{Index: 5, Status: candidatejob.StatusFailedCv},
	{Index: 6, Status: candidatejob.StatusFailedInterview},
	{Index: 7, Status: candidatejob.StatusExStaff},
}

type CandidateJobValidInput struct {
	Status         ent.CandidateJobStatus `json:"status"`
	OnboardDate    *time.Time             `json:"onboard_date"`
	OfferExpDate   *time.Time             `json:"offer_expiration_date"`
	CandidateId    uuid.UUID              `json:"candidate_id"`
	CandidateJobId uuid.UUID
	FailedReason   []ent.CandidateJobFailedReason `json:"failed_reason"`
}

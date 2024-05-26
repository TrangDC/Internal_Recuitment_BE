package models

import "trec/ent/candidatejob"

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
	{Index: 5, Status: candidatejob.StatusKiv},
	{Index: 6, Status: candidatejob.StatusExStaff},
}

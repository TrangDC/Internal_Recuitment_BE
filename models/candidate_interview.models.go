package models

import (
	"time"

	"github.com/google/uuid"
)

type CandidateInterviewFilter struct {
	InterviewDate *time.Time `json:"interview_date"`
	StartFrom     *time.Time `json:"start_from"`
	EndAt         *time.Time `json:"end_at"`
	Interviewer   []string   `json:"interviewer"`
	FromDate      *time.Time `json:"from_date"`
	ToDate        *time.Time `json:"to_date"`
	TeamId        *string    `json:"team_id"`
	HiringJobId   *string    `json:"hiring_job_id"`
}

type CandidateInterviewInputValidate struct {
	Title          string     `json:"title"`
	CandidateJobId uuid.UUID  `json:"candidate_job_id"`
	StartFrom      *time.Time `json:"start_from"`
	EndAt          *time.Time `json:"end_at"`
	InterviewDate  *time.Time `json:"interview_date"`
}
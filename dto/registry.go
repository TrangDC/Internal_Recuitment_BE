package dto

import (
	"reflect"
	"sort"
)

// Dto is the interface for all dto.
type Dto interface {
	Candidate() CandidateDto
	CandidateJob() CandidateJobDto
	CandidateInterview() CandidateInterviewDto
	CandidateJobFeedback() CandidateJobFeedbackDto
}

// dtoImpl is the implementation of Dto.
type dtoImpl struct {
	candidateDto            CandidateDto
	candidateJobDto         CandidateJobDto
	candidateInterviewDto   CandidateInterviewDto
	candidateJobFeedbackDto CandidateJobFeedbackDto
}

// NewDto creates a new Dto.
func NewDto() Dto {
	return &dtoImpl{
		candidateDto:            NewCandidateDto(),
		candidateJobDto:         NewCandidateJobDto(),
		candidateInterviewDto:   NewCandidateInterviewDto(),
		candidateJobFeedbackDto: NewCandidateJobFeedbackDto(),
	}
}

var (
	CandidateI18n            = "model.candidates.model_name"
	CandidateJobI18n         = "model.candidate_jobs.model_name"
	CandidateInterviewI18n   = "model.candidate_interviews.model_name"
	CandidateJobFeedbackI18n = "model.candidate_job_feedbacks.model_name"
)

func (i dtoImpl) Candidate() CandidateDto {
	return i.candidateDto
}

func (i dtoImpl) CandidateJob() CandidateJobDto {
	return i.candidateJobDto
}

func (i dtoImpl) CandidateInterview() CandidateInterviewDto {
	return i.candidateInterviewDto
}

func (i dtoImpl) CandidateJobFeedback() CandidateJobFeedbackDto {
	return i.candidateJobFeedbackDto
}

func CompareArray(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sortedArr1 := make([]string, len(arr1))
	sortedArr2 := make([]string, len(arr2))
	copy(sortedArr1, arr1)
	copy(sortedArr2, arr2)
	sort.Strings(sortedArr1)
	sort.Strings(sortedArr2)
	return reflect.DeepEqual(sortedArr1, sortedArr2)
}

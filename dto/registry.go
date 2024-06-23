package dto

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// Dto is the interface for all dto.
type Dto interface {
	Azure() AzureDto
	Candidate() CandidateDto
	CandidateJob() CandidateJobDto
	CandidateInterview() CandidateInterviewDto
	CandidateJobFeedback() CandidateJobFeedbackDto
	HiringJob() HiringJobDto
	Skill() SkillDto
	Team() TeamDto
	User() UserDto
}

// dtoImpl is the implementation of Dto.
type dtoImpl struct {
	azureDto                AzureDto
	candidateDto            CandidateDto
	candidateJobDto         CandidateJobDto
	candidateInterviewDto   CandidateInterviewDto
	candidateJobFeedbackDto CandidateJobFeedbackDto
	hiringJobDto            HiringJobDto
	skillDto                SkillDto
	teamDto                 TeamDto
	userDto                 UserDto
}

// NewDto creates a new Dto.
func NewDto() Dto {
	return &dtoImpl{
		azureDto:                NewAzureDto(),
		candidateDto:            NewCandidateDto(),
		candidateJobDto:         NewCandidateJobDto(),
		candidateInterviewDto:   NewCandidateInterviewDto(),
		candidateJobFeedbackDto: NewCandidateJobFeedbackDto(),
		hiringJobDto:            NewHiringJobDto(),
		skillDto:                NewSkillDto(),
		teamDto:                 NewTeamDto(),
		userDto:                 NewUserDto(),
	}
}

var (
	CandidateI18n            = "model.candidates.model_name"
	CandidateJobI18n         = "model.candidate_jobs.model_name"
	CandidateInterviewI18n   = "model.candidate_interviews.model_name"
	CandidateJobFeedbackI18n = "model.candidate_job_feedbacks.model_name"
	HiringJobI18n            = "model.hiring_jobs.model_name"
	SkillI18n                = "model.skills.model_name"
	TeamI18n                 = "model.teams.model_name"
	UserI18n                 = "model.users.model_name"
)

func (i dtoImpl) Azure() AzureDto {
	return i.azureDto
}

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

func (i dtoImpl) HiringJob() HiringJobDto {
	return i.hiringJobDto
}

func (i dtoImpl) Skill() SkillDto {
	return i.skillDto
}

func (i dtoImpl) Team() TeamDto {
	return i.teamDto
}

func (i dtoImpl) User() UserDto {
	return i.userDto
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

func IsRecordEdited(createdAt, updatedAt time.Time) bool {
	createdAtStart := createdAt.Add(-1 * time.Second)
	createdAtEnd := createdAt.Add(1 * time.Second)
	if createdAtStart.Before(updatedAt) && createdAtEnd.After(updatedAt) {
		return false
	}
	return true
}

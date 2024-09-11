package dto

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// Dto is the interface for all dto.
type Dto interface {
	Azure() AzureDto
	Candidate() CandidateDto
	CandidateJob() CandidateJobDto
	CandidateInterview() CandidateInterviewDto
	CandidateJobFeedback() CandidateJobFeedbackDto
	JobPosition() JobPositionDto
	HiringJob() HiringJobDto
	Skill() SkillDto
	EntitySkill() EntitySkillDto
	SkillType() SkillTypeDto
	HiringTeam() HiringTeamDto
	User() UserDto
	Role() RoleDto
	EmailTemplate() EmailTemplateDto
	RecTeam() RecTeamDto
	CandidateHistoryCall() CandidateHistoryCallDto
	CandidateNote() CandidateNoteDto
}

// dtoImpl is the implementation of Dto.
type dtoImpl struct {
	azureDto                AzureDto
	candidateDto            CandidateDto
	candidateJobDto         CandidateJobDto
	candidateInterviewDto   CandidateInterviewDto
	candidateJobFeedbackDto CandidateJobFeedbackDto
	jobPositionDto          JobPositionDto
	hiringJobDto            HiringJobDto
	skillDto                SkillDto
	entitySkillDto          EntitySkillDto
	skillTypeDto            SkillTypeDto
	hiringTeamDto           HiringTeamDto
	userDto                 UserDto
	roleDto                 RoleDto
	emailTemplateDto        EmailTemplateDto
	recTeamDto              RecTeamDto
	candidateHistoryCallDto CandidateHistoryCallDto
	candidateNoteDto        CandidateNoteDto
}

// NewDto creates a new Dto.
func NewDto() Dto {
	return &dtoImpl{
		azureDto:                NewAzureDto(),
		candidateDto:            NewCandidateDto(),
		candidateJobDto:         NewCandidateJobDto(),
		candidateInterviewDto:   NewCandidateInterviewDto(),
		candidateJobFeedbackDto: NewCandidateJobFeedbackDto(),
		jobPositionDto:          NewJobPositionDto(),
		hiringJobDto:            NewHiringJobDto(),
		skillDto:                NewSkillDto(),
		entitySkillDto:          NewEntitySkillDto(),
		skillTypeDto:            NewSkillTypeDto(),
		hiringTeamDto:           NewHiringTeamDto(),
		userDto:                 NewUserDto(),
		roleDto:                 NewRoleDto(),
		emailTemplateDto:        NewEmailTemplateDto(),
		recTeamDto:              NewRecTeamDto(),
		candidateHistoryCallDto: NewCandidateHistoryCallDto(),
		candidateNoteDto:        NewCandidateNoteDto(),
	}
}

const (
	CandidateI18n            = "model.candidates.model_name"
	CandidateJobI18n         = "model.candidate_jobs.model_name"
	CandidateInterviewI18n   = "model.candidate_interviews.model_name"
	CandidateJobFeedbackI18n = "model.candidate_job_feedbacks.model_name"
	JobPositionI18n          = "model.job_positions.model_name"
	HiringJobI18n            = "model.hiring_jobs.model_name"
	SkillI18n                = "model.skills.model_name"
	SkillTypeI18n            = "model.skill_types.model_name"
	HiringTeamI18n           = "model.hiring_teams.model_name"
	UserI18n                 = "model.users.model_name"
	RoleI18n                 = "model.roles.model_name"
	EmailTemplateI18n        = "model.email_templates.model_name"
	RecTeamI18n              = "model.rec_teams.model_name"
	EntityPermissionI18n     = "model.entity_permissions.model_name"
	CandidateHistoryCallI18n = "model.candidate_history_calls.model_name"
	CandidateNoteI18n        = "model.candidate_notes.model_name"
	CandidateAwardI18n       = "model.candidate_awards.model_name"
	CandidateExpI18n         = "model.candidate_exps.model_name"
	CandidateEducateI18n     = "model.candidate_educates.model_name"
	CandidateCertificateI18n = "model.candidate_certificates.model_name"
	HiringJobStepI18n        = "model.hiring_job_steps.model_name"
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

func (i dtoImpl) JobPosition() JobPositionDto {
	return i.jobPositionDto
}

func (i dtoImpl) HiringJob() HiringJobDto {
	return i.hiringJobDto
}

func (i dtoImpl) Skill() SkillDto {
	return i.skillDto
}

func (i dtoImpl) HiringTeam() HiringTeamDto {
	return i.hiringTeamDto
}

func (i dtoImpl) User() UserDto {
	return i.userDto
}

func (i dtoImpl) EntitySkill() EntitySkillDto {
	return i.entitySkillDto
}

func (i dtoImpl) SkillType() SkillTypeDto {
	return i.skillTypeDto
}

func (i dtoImpl) Role() RoleDto {
	return i.roleDto
}

func (i dtoImpl) EmailTemplate() EmailTemplateDto {
	return i.emailTemplateDto
}

func (i dtoImpl) RecTeam() RecTeamDto {
	return i.recTeamDto
}

func (i dtoImpl) CandidateHistoryCall() CandidateHistoryCallDto {
	return i.candidateHistoryCallDto
}

func (i dtoImpl) CandidateNote() CandidateNoteDto {
	return i.candidateNoteDto
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

func FormatCurrency(number int) string {
	numStr := strconv.Itoa(number)
	length := len(numStr)
	var result []string
	for i := length; i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{numStr[start:i]}, result...)
	}
	return strings.Join(result, ",")
}

func ConvertTimeZone(input time.Time, location string) (time.Time, string) {
	inputTime := carbon.Parse(input.Format("2006-01-02 15:04:05"))
	timeZone := ""
	locationString := carbon.HoChiMinh
	switch location {
	case "Singapore":
		locationString = carbon.Singapore
	case "Thailand":
		locationString = carbon.Bangkok
	case "India":
		locationString = carbon.Kolkata
	case "Japan":
		locationString = carbon.Tokyo
	case "China":
		locationString = carbon.Shanghai
	case "Australia":
		locationString = carbon.Sydney
	case "United States":
		locationString = carbon.NewYork
	case "United Kingdom":
		locationString = carbon.London
	case "Germany":
		locationString = carbon.Berlin
	case "France":
		locationString = carbon.Paris
	}
	timeLocation, err := time.LoadLocation(locationString)
	if err != nil {
		log.Println("Error when load location", err)
	}
	timeWithLocation := input.In(timeLocation).Format("2006-01-02 15:04:05")
	result := carbon.Parse(timeWithLocation)
	numebrOfTz := inputTime.DiffInHours(result)
	if numebrOfTz < 0 {
		timeZone = "-" + fmt.Sprint(numebrOfTz)
	} else {
		timeZone = "+" + fmt.Sprint(numebrOfTz)
	}
	return result.StdTime(), timeZone
}

func FindCUDArray(arr1, arr2 []interface{}) (creIds, updIds, delIds []uuid.UUID) {
	arr1Ids := lo.Map(arr1, func(entity interface{}, _ int) uuid.UUID {
		value := reflect.ValueOf(interface{}(entity)).Elem()
		return value.Field(1).Interface().(uuid.UUID)
	})
	arr2Ids := lo.Map(arr2, func(entity interface{}, _ int) uuid.UUID {
		value := reflect.ValueOf(interface{}(entity)).Elem()
		return value.Field(1).Interface().(uuid.UUID)
	})
	delIds, credIds := lo.Difference(arr1Ids, arr2Ids)
	updIds = lo.Intersect(arr1Ids, arr2Ids)
	return credIds, updIds, delIds
}

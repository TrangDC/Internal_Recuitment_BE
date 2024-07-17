package service

import (
	"context"
	"fmt"
	"strings"
	"trec/config"
	"trec/dto"
	"trec/ent"
	"trec/internal/servicebus"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type EmailService interface {
	GenerateEmail(ctx context.Context, users []*ent.User, record *ent.EmailTemplate,
		groupModule models.GroupModule) []models.MessageInput
	SentEmail(ctx context.Context, input []models.MessageInput) error
}

type emailSvcImpl struct {
	repoRegistry     repository.Repository
	serviceBusClient servicebus.ServiceBus
	dtoRegistry      dto.Dto
	logger           *zap.Logger
	configs          *config.Configurations
}

func NewEmailService(repoRegistry repository.Repository, serviceBusClient servicebus.ServiceBus, dtoRegistry dto.Dto,
	logger *zap.Logger, configs *config.Configurations) EmailService {
	return &emailSvcImpl{
		repoRegistry:     repoRegistry,
		serviceBusClient: serviceBusClient,
		dtoRegistry:      dtoRegistry,
		logger:           logger,
		configs:          configs,
	}
}

func (svc *emailSvcImpl) GenerateEmail(ctx context.Context, users []*ent.User, template *ent.EmailTemplate,
	groupModule models.GroupModule) []models.MessageInput {
	var messages []models.MessageInput
	sendToIDs, filteredUsers := svc.getSentTo(users, template, groupModule)
	keywords := svc.mappingKeyword(groupModule)
	sendToIDs = lo.Uniq(sendToIDs)
	for _, userID := range sendToIDs {
		user, _ := lo.Find(filteredUsers, func(u *ent.User) bool {
			return u.ID == userID
		})
		subject := strings.ReplaceAll(template.Subject, "{{ gl:receiver_name }}", user.Name)
		content := strings.ReplaceAll(template.Content, "{{ gl:receiver_name }}", user.Name)
		for key, value := range keywords {
			subject = strings.ReplaceAll(subject, key, value)
			content = strings.ReplaceAll(content, key, value)
		}
		message := models.MessageInput{
			ID:        template.ID,
			To:        []string{user.WorkEmail},
			Cc:        template.Cc,
			Bcc:       template.Bcc,
			Subject:   subject,
			Content:   content,
			Signature: template.Signature,
		}
		messages = append(messages, message)
	}
	return messages
}

func (svc *emailSvcImpl) getSentTo(users []*ent.User, record *ent.EmailTemplate, groupModule models.GroupModule) ([]uuid.UUID, []*ent.User) {
	var sendToIds []uuid.UUID
	sentTo := lo.Map(record.SendTo, func(entity string, index int) ent.EmailTemplateSendTo {
		if ent.EmailTemplateApplicationEventEnum.IsValid(ent.EmailTemplateApplicationEventEnum(record.Event.String())) &&
			!ent.EmailTemplateApplicationSendToEnum.IsValid(ent.EmailTemplateApplicationSendToEnum(entity)) {
			return ""
		}
		return ent.EmailTemplateSendTo(entity)
	})
	roleMemberIds := lo.Flatten(lo.Map(record.Edges.RoleEdges, func(entity *ent.Role, index int) []uuid.UUID {
		return lo.Map(entity.Edges.UserEdges, func(entity *ent.User, index int) uuid.UUID {
			return entity.ID
		})
	}))
	sendToIds = append(sendToIds, roleMemberIds...)
	for _, value := range sentTo {
		switch value {
		case "":
			continue
		case ent.EmailTemplateSendToInterviewer:
			interviewerIds := lo.Map(groupModule.Interview.Edges.InterviewerEdges, func(entity *ent.User, index int) uuid.UUID {
				return entity.ID
			})
			sendToIds = append(sendToIds, interviewerIds...)
		case ent.EmailTemplateSendToJobRequest:
			requestId := groupModule.HiringJob.CreatedBy
			sendToIds = append(sendToIds, requestId)
		case ent.EmailTemplateSendToTeamManager:
			managerIds := lo.Map(groupModule.Team.Edges.UserEdges, func(entity *ent.User, index int) uuid.UUID {
				return entity.ID
			})
			sendToIds = append(sendToIds, managerIds...)
		case ent.EmailTemplateSendToTeamMember:
			memberIds := lo.Map(groupModule.Team.Edges.MemberEdges, func(entity *ent.User, index int) uuid.UUID {
				return entity.ID
			})
			sendToIds = append(sendToIds, memberIds...)
		case ent.EmailTemplateSendToCandidate:
			sendToIds = append(sendToIds, groupModule.Candidate.ID)
			users = append(users, &ent.User{
				ID:        groupModule.Candidate.ID,
				Name:      groupModule.Candidate.Name,
				WorkEmail: groupModule.Candidate.Email,
			})
		}
	}
	return lo.Uniq(sendToIds), users
}

func (svc *emailSvcImpl) mappingKeyword(groupModule models.GroupModule) map[string]string {
	var keywords = models.AllEmailTPKeyword
	if groupModule.Team != nil {
		managerNames := lo.Map(groupModule.Team.Edges.UserEdges, func(entity *ent.User, index int) string {
			return entity.Name
		})
		keywords["{{ tm:name }}"] = groupModule.Team.Name
		keywords["{{ tm:manager_name }}"] = strings.Join(managerNames, ", ")
		keywords["{{ lk:team }}"] = fmt.Sprintf("%s/dashboard/team-detail/%s", svc.configs.App.AppUrl, groupModule.Team.ID)
	}
	if groupModule.HiringJob != nil {
		skillNames := lo.Map(groupModule.HiringJob.Edges.HiringJobSkillEdges, func(entity *ent.EntitySkill, index int) string {
			return entity.Edges.SkillEdge.Name
		})
		keywords["{{ hrjb:name }}"] = groupModule.HiringJob.Name
		keywords["{{ hrjb:skill_name }}"] = strings.Join(skillNames, ", ")
		keywords["{{ hrjb:location }}"] = svc.dtoRegistry.HiringJob().MappingLocation(groupModule.HiringJob.Location) // enum
		keywords["{{ hrjb:requester }}"] = groupModule.HiringJob.Edges.OwnerEdge.Name
		keywords["{{ hrjb:staff_required }}"] = fmt.Sprint(groupModule.HiringJob.Amount)
		keywords["{{ hrjb:status }}"] = svc.dtoRegistry.HiringJob().MappingStatus(groupModule.HiringJob.Status)       // enum
		keywords["{{ hrjb:priority }}"] = svc.dtoRegistry.HiringJob().MappingPriority(groupModule.HiringJob.Priority) // enum
		keywords["{{ hrjb:salary }}"] = svc.dtoRegistry.HiringJob().MappingSalary(groupModule.HiringJob)              // string by type of salary_type
		keywords["{{ hrjb:description }}"] = groupModule.HiringJob.Description
		keywords["{{ lk:job }}"] = fmt.Sprintf("%s/dashboard/job-detail/%s", svc.configs.App.AppUrl, groupModule.HiringJob.ID)
	}
	if groupModule.Candidate != nil {
		var referenceUserName = ""
		skillNames := lo.Map(groupModule.Candidate.Edges.CandidateSkillEdges, func(entity *ent.EntitySkill, index int) string {
			return entity.Edges.SkillEdge.Name
		})
		referenceUser := groupModule.Candidate.Edges.ReferenceUserEdge
		if referenceUser != nil {
			referenceUserName = referenceUser.Name
		}
		if !groupModule.Candidate.RecruitTime.IsZero() {
			keywords["{{ cd:recruit_date }}"] = groupModule.Candidate.RecruitTime.Format("02-01-2006")
		}
		if !groupModule.Candidate.Dob.IsZero() {
			keywords["{{ cd:dob }}"] = groupModule.Candidate.Dob.Format("02-01-2006")
		}
		source := svc.dtoRegistry.Candidate().MappingReferenceType(groupModule.Candidate.ReferenceType, groupModule.Candidate.ReferenceValue)
		keywords["{{ cd:name }}"] = groupModule.Candidate.Name
		keywords["{{ cd:email }}"] = groupModule.Candidate.Email
		keywords["{{ cd:phone }}"] = groupModule.Candidate.Phone
		keywords["{{ cd:recruiter }}"] = referenceUserName
		keywords["{{ cd:source }}"] = source // enum
		keywords["{{ cd:skill_name }}"] = strings.Join(skillNames, ", ")
		keywords["{{ lk:candidate }}"] = fmt.Sprintf("%s/dashboard/candidate-detail/%s", svc.configs.App.AppUrl, groupModule.Candidate.ID)
	}
	if groupModule.CandidateJob != nil {
		keywords["{{ cdjb:status }}"] = svc.dtoRegistry.CandidateJob().MappingStatus(groupModule.CandidateJob.Status) // enum
		keywords["{{ cdjb:applied_date }}"] = groupModule.CandidateJob.CreatedAt.Format("02-01-2006")
		keywords["{{ lk:candidate_job_application }}"] = fmt.Sprintf("%s/dashboard/job-application-detail/%s", svc.configs.App.AppUrl, groupModule.CandidateJob.ID)
	}
	if groupModule.Interview != nil {
		interviewers := lo.Map(groupModule.Interview.Edges.InterviewerEdges, func(entity *ent.User, index int) string {
			return entity.Name
		})
		inteviewStartTime := groupModule.Interview.StartFrom.Format("15:04")
		inteviewEndTime := groupModule.Interview.EndAt.Format("15:04")
		keywords["{{ intv:title }}"] = groupModule.Interview.Title
		keywords["{{ intv:interviewer_name }}"] = strings.Join(interviewers, ", ")
		keywords["{{ intv:date }}"] = groupModule.Interview.InterviewDate.Format("02-01-2006")
		keywords["{{ intv:time }}"] = fmt.Sprintf("%s - %s (UTC+0)", inteviewStartTime, inteviewEndTime)
		keywords["{{ lk:interview }}"] = fmt.Sprintf("%sdashboard/calendars?interview_id=%s&is_open_detail=true", svc.configs.App.AppUrl, groupModule.Interview.ID) // connect with FE
		// keywords["{{ intv:location }}"] = groupModule.Interview.Location // enum
	}
	return keywords
}

func (svc *emailSvcImpl) SentEmail(ctx context.Context, input []models.MessageInput) error {
	for _, msg := range input {
		msg.QueueName = "trec-email-event-trigger"
		err := svc.serviceBusClient.SendEmailTriggerMessage(ctx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

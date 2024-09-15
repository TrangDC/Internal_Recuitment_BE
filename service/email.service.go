package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/config"
	"trec/dto"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/outgoingemail"
	"trec/internal/servicebus"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type EmailService interface {
	GenerateEmail(ctx context.Context, users []*ent.User, record *ent.EmailTemplate, groupModule models.GroupModule) []models.MessageInput
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

func (svc *emailSvcImpl) GenerateEmail(
	ctx context.Context, users []*ent.User, template *ent.EmailTemplate,
	groupModule models.GroupModule,
) []models.MessageInput {
	var messages []models.MessageInput
	sendTos, filteredUsers := svc.getSentTo(users, template, groupModule)
	sendTos = lo.Uniq(sendTos)
	for _, sentTo := range sendTos {
		var user *ent.User
		user, _ = lo.Find(filteredUsers, func(u *ent.User) bool {
			return u.ID == sentTo.ID
		})
		keywords := svc.mappingKeyword(groupModule, user.ID)
		subject := strings.ReplaceAll(template.Subject, "{{ gl:receiver_name }}", user.Name)
		content := strings.ReplaceAll(template.Content, "{{ gl:receiver_name }}", user.Name)
		signature := strings.ReplaceAll(template.Signature, "{{ gl:receiver_name }}", user.Name)
		if groupModule.Interview != nil {
			interviewDate, timeZone := dto.ConvertTimeZone(groupModule.Interview.InterviewDate, user.Location)
			interviewStartTime, _ := dto.ConvertTimeZone(groupModule.Interview.StartFrom, user.Location)
			interviewEndTime, _ := dto.ConvertTimeZone(groupModule.Interview.EndAt, user.Location)
			keywords["{{ intv:date }}"] = interviewDate.Format("02-01-2006")
			keywords["{{ intv:time }}"] = fmt.Sprintf("%s - %s (UTC%s)", interviewStartTime.Format("15:04"), interviewEndTime.Format("15:04"), timeZone)
		}
		if !groupModule.CandidateJob.OnboardDate.IsZero() && !groupModule.CandidateJob.OfferExpirationDate.IsZero() {
			offerEpxDate, _ := dto.ConvertTimeZone(groupModule.CandidateJob.OfferExpirationDate, user.Location)
			onboardDate, _ := dto.ConvertTimeZone(groupModule.CandidateJob.OnboardDate, user.Location)
			keywords["{{ cdjb:onboard_date }}"] = onboardDate.Format("02-01-2006")
			keywords["{{ cdjb:offer_expiration_date }}"] = offerEpxDate.Format("02-01-2006")
		}
		for key, value := range keywords {
			subject = strings.ReplaceAll(subject, key, value)
			content = strings.ReplaceAll(content, key, value)
			signature = strings.ReplaceAll(signature, key, value)
		}
		message := models.MessageInput{
			ID:            template.ID,
			To:            []string{user.WorkEmail},
			Cc:            template.Cc,
			Bcc:           template.Bcc,
			Subject:       subject,
			Content:       content,
			Signature:     signature,
			RecipientType: sentTo.RecipientType,
			EventID:       template.EventID,
		}
		messages = append(messages, message)
	}
	return messages
}

func (svc *emailSvcImpl) getSentTo(users []*ent.User, record *ent.EmailTemplate, groupModule models.GroupModule) ([]models.SentTo, []*ent.User) {
	var sendTos []models.SentTo
	sentTo := lo.Map(record.SendTo, func(entity string, index int) ent.EmailTemplateSendTo {
		if ent.EmailTemplateApplicationEventEnum.IsValid(ent.EmailTemplateApplicationEventEnum(record.Event.String())) &&
			!ent.EmailTemplateApplicationSendToEnum.IsValid(ent.EmailTemplateApplicationSendToEnum(entity)) {
			return ""
		}
		return ent.EmailTemplateSendTo(entity)
	})
	roleMemberSentTos := lo.Flatten(lo.Map(record.Edges.RoleEdges, func(entity *ent.Role, index int) []models.SentTo {
		return lo.Map(entity.Edges.UserEdges, func(entity *ent.User, index int) models.SentTo {
			return models.SentTo{
				ID:            entity.ID,
				RecipientType: outgoingemail.RecipientTypeRole,
			}
		})
	}))
	sendTos = append(sendTos, roleMemberSentTos...)
	for _, value := range sentTo {
		switch value {
		case "":
			continue
		case ent.EmailTemplateSendToInterviewer:
			interviewerSentTos := lo.Map(groupModule.Interview.Edges.InterviewerEdges, func(entity *ent.User, index int) models.SentTo {
				return models.SentTo{
					ID:            entity.ID,
					RecipientType: outgoingemail.RecipientTypeInterviewer,
				}
			})
			sendTos = append(sendTos, interviewerSentTos...)
		case ent.EmailTemplateSendToJobRequest:
			requestId := groupModule.HiringJob.CreatedBy
			sendTos = append(sendTos, models.SentTo{
				ID:            requestId,
				RecipientType: outgoingemail.RecipientTypeJobRequest,
			})
		case ent.EmailTemplateSendToJobRecInCharge:
			sendTos = append(sendTos, models.SentTo{
				ID:            groupModule.HiringJob.RecInChargeID,
				RecipientType: outgoingemail.RecipientTypeJobRecInCharge,
			})
		case ent.EmailTemplateSendToCdJobRecInCharge:
			sendTos = append(sendTos, models.SentTo{
				ID:            groupModule.CandidateJob.RecInChargeID,
				RecipientType: outgoingemail.RecipientTypeCdJobRecInCharge,
			})
		case ent.EmailTemplateSendToHiringTeamManager:
			managerSentTos := lo.Map(groupModule.HiringTeam.Edges.UserEdges, func(entity *ent.User, index int) models.SentTo {
				return models.SentTo{
					ID:            entity.ID,
					RecipientType: outgoingemail.RecipientTypeHiringTeamManager,
				}
			})
			sendTos = append(sendTos, managerSentTos...)
		case ent.EmailTemplateSendToHiringApprover:
			approverSentTos := lo.Map(groupModule.HiringTeam.Edges.ApproversUsers, func(entity *ent.User, _ int) models.SentTo {
				return models.SentTo{
					ID:            entity.ID,
					RecipientType: outgoingemail.RecipientTypeHiringApprover,
				}
			})
			sendTos = append(sendTos, approverSentTos...)
		case ent.EmailTemplateSendToHiringTeamMember:
			memberSentTos := lo.Map(groupModule.HiringTeam.Edges.HiringMemberEdges, func(entity *ent.User, index int) models.SentTo {
				return models.SentTo{
					ID:            entity.ID,
					RecipientType: outgoingemail.RecipientTypeHiringTeamMember,
				}
			})
			sendTos = append(sendTos, memberSentTos...)
		case ent.EmailTemplateSendToRecLeader:
			sendTos = append(sendTos, models.SentTo{
				ID:            groupModule.RecTeam.LeaderID,
				RecipientType: outgoingemail.RecipientTypeRecLeader,
			})
		case ent.EmailTemplateSendToRecMember:
			recMemberSentTos := lo.Map(groupModule.RecTeam.Edges.RecMemberEdges, func(entity *ent.User, _ int) models.SentTo {
				return models.SentTo{
					ID:            entity.ID,
					RecipientType: outgoingemail.RecipientTypeRecMember,
				}
			})
			sendTos = append(sendTos, recMemberSentTos...)
		case ent.EmailTemplateSendToCandidate:
			sendTos = append(sendTos, models.SentTo{
				ID:            groupModule.Candidate.ID,
				RecipientType: outgoingemail.RecipientTypeCandidate,
			})
			users = append(users, &ent.User{
				ID:        groupModule.Candidate.ID,
				Name:      groupModule.Candidate.Name,
				WorkEmail: groupModule.Candidate.Email,
				Location:  groupModule.Candidate.Country,
			})
		}
	}
	return lo.Uniq(sendTos), users
}

func (svc *emailSvcImpl) mappingKeyword(groupModule models.GroupModule, userID uuid.UUID) map[string]string {
	var keywords = models.AllEmailTPKeyword
	if groupModule.RecTeam != nil {
		keywords["{{ rec:name }}"] = groupModule.RecTeam.Name
		keywords["{{ rec:leader }}"] = groupModule.RecTeam.Edges.RecLeaderEdge.Name
		keywords["{{ lk:rec_team }}"] = fmt.Sprintf(
			"%s/rec-dashboard/team-detail/%s",
			svc.configs.App.AppUrl, groupModule.RecTeam.ID,
		)
	}
	if groupModule.HiringTeam != nil {
		managerNames := lo.Map(groupModule.HiringTeam.Edges.UserEdges, func(entity *ent.User, index int) string {
			return entity.Name
		})
		keywords["{{ hrtm:name }}"] = groupModule.HiringTeam.Name
		keywords["{{ hrtm:manager_name }}"] = strings.Join(managerNames, ", ")
		keywords["{{ lk:hiring_team }}"] = fmt.Sprintf("%s/dashboard/team-detail/%s", svc.configs.App.AppUrl, groupModule.HiringTeam.ID)
		keywords["{{ hrtm:approvers }}"] = strings.Join(
			lo.Map(groupModule.HiringTeam.Edges.ApproversUsers, func(entity *ent.User, _ int) string { return entity.Name }),
			", ",
		)
	}
	if groupModule.JobPosition != nil {
		keywords["{{ jbpos:name }}"] = groupModule.JobPosition.Name
	}
	if groupModule.HiringJob != nil {
		skillNames := lo.Map(groupModule.HiringJob.Edges.HiringJobSkillEdges, func(entity *ent.EntitySkill, index int) string {
			return entity.Edges.SkillEdge.Name
		})
		keywords["{{ hrjb:name }}"] = groupModule.HiringJob.Name
		keywords["{{ hrjb:rec_in_charge }}"] = groupModule.HiringJob.Edges.RecInChargeEdge.Name
		keywords["{{ hrjb:skill_name }}"] = strings.Join(skillNames, ", ")
		keywords["{{ hrjb:level }}"] = groupModule.HiringJob.Level.String()
		keywords["{{ hrjb:location }}"] = svc.dtoRegistry.HiringJob().MappingLocation(groupModule.HiringJob.Location) // enum
		keywords["{{ hrjb:requester }}"] = groupModule.HiringJob.Edges.OwnerEdge.Name
		keywords["{{ hrjb:staff_required }}"] = fmt.Sprint(groupModule.HiringJob.Amount)
		keywords["{{ hrjb:status }}"] = svc.dtoRegistry.HiringJob().MappingStatus(groupModule.HiringJob.Status)       // enum
		keywords["{{ hrjb:priority }}"] = svc.dtoRegistry.HiringJob().MappingPriority(groupModule.HiringJob.Priority) // enum
		keywords["{{ hrjb:salary }}"] = svc.dtoRegistry.HiringJob().MappingSalary(groupModule.HiringJob)              // string by type of salary_type
		keywords["{{ hrjb:description }}"] = groupModule.HiringJob.Description
		linkFmt := "%s/dashboard/job-detail/%s"
		if groupModule.HiringJob.Status != hiringjob.StatusPendingApprovals {
			linkFmt = "%s/dashboard/job-overview/%s"
		}
		keywords["{{ lk:job }}"] = fmt.Sprintf(linkFmt, svc.configs.App.AppUrl, groupModule.HiringJob.ID)
		keywords["{{ hrjb:resolution_time }}"] = func(openedAt, closedAt time.Time) string {
			// TODO: Handle stacked resolution time when reopen
			if openedAt.IsZero() {
				return ""
			}
			if closedAt.IsZero() {
				return time.Now().UTC().Sub(openedAt).String()
			}
			return closedAt.Sub(openedAt).String()
		}(groupModule.HiringJob.OpenedAt, groupModule.HiringJob.ClosedAt)
		hiringJobAuditTrail, exists := lo.Find(groupModule.AuditTrails, func(item *ent.AuditTrail) bool {
			return item.RecordId == groupModule.HiringJob.ID
		})
		if !exists {
			keywords["{{ hrjb:audit_trail }}"] = "N/A"
		}
		var err error
		keywords["{{ hrjb:audit_trail }}"], err = svc.dtoRegistry.EmailTemplate().FormatAuditTrail4Email(hiringJobAuditTrail)
		if err != nil {
			svc.logger.Error(err.Error())
		}
		keywords["{{lk:approve}}"] = fmt.Sprintf(
			"%s/dashboard/job-detail/%s?approval_status=%s&confirm_modal=open",
			svc.configs.App.AppUrl, groupModule.HiringJob.ID,
			hiringjobstep.StatusAccepted.String(),
		)
		keywords["{{lk:reject}}"] = fmt.Sprintf(
			"%s/dashboard/job-detail/%s?approval_status=%s&confirm_modal=open",
			svc.configs.App.AppUrl, groupModule.HiringJob.ID,
			hiringjobstep.StatusRejected.String(),
		)
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
		keywords["{{ cd:address }}"] = groupModule.Candidate.Address
		keywords["{{ cd:recruiter }}"] = referenceUserName
		keywords["{{ cd:source }}"] = source // enum
		keywords["{{ cd:skill_name }}"] = strings.Join(skillNames, ", ")
		keywords["{{ lk:candidate }}"] = fmt.Sprintf("%s/dashboard/candidate-detail/%s", svc.configs.App.AppUrl, groupModule.Candidate.ID)
	}
	if groupModule.CandidateJob != nil {
		keywords["{{ cdjb:status }}"] = svc.dtoRegistry.CandidateJob().MappingStatus(groupModule.CandidateJob.Status) // enum
		keywords["{{ cdjb:applied_date }}"] = groupModule.CandidateJob.CreatedAt.Format("02-01-2006")
		keywords["{{ lk:candidate_job_application }}"] = fmt.Sprintf("%s/dashboard/job-application-detail/%s", svc.configs.App.AppUrl, groupModule.CandidateJob.ID)
		keywords["{{ cdjb:rec_in_charge }}"] = groupModule.CandidateJob.Edges.RecInChargeEdge.Name
	}
	if groupModule.Interview != nil {
		interviewers := lo.Map(groupModule.Interview.Edges.InterviewerEdges, func(entity *ent.User, index int) string {
			return entity.Name
		})
		keywords["{{ intv:title }}"] = groupModule.Interview.Title
		keywords["{{ intv:interviewer_name }}"] = strings.Join(interviewers, ", ")
		keywords["{{ lk:interview }}"] = fmt.Sprintf("%s/dashboard/calendars?interview_id=%s&is_open_detail=true", svc.configs.App.AppUrl, groupModule.Interview.ID) // connect with FE
		keywords["{{ intv:location }}"] = svc.dtoRegistry.CandidateInterview().MappingLocation(groupModule.Interview.Location)                                       // enum
		keywords["{{ intv:meeting_link }}"] = groupModule.Interview.MeetingLink
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

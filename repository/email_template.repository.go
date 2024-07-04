package repository

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/emailtemplate"
	"trec/ent/role"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type EmailTemplateRepository interface {
	CreateEmailTemplate(ctx context.Context, input ent.NewEmailTemplateInput, roleIds []uuid.UUID, sendTo []string) (*ent.EmailTemplate, error)
	UpdateEmailTemplate(ctx context.Context, record *ent.EmailTemplate, input ent.UpdateEmailTemplateInput, newRoleIds, removeRoleIds []uuid.UUID, sendTo []string) (*ent.EmailTemplate, error)
	UpdateEmailTemplateStatus(ctx context.Context, record *ent.EmailTemplate, input ent.UpdateEmailTemplateStatusInput) error
	DeleteEmailTemplate(ctx context.Context, record *ent.EmailTemplate, roleIds []uuid.UUID) (*ent.EmailTemplate, error)

	// query
	GetEmailTemplate(ctx context.Context, id uuid.UUID) (*ent.EmailTemplate, error)
	BuildQuery() *ent.EmailTemplateQuery
	BuildBaseQuery() *ent.EmailTemplateQuery
	BuildCount(ctx context.Context, query *ent.EmailTemplateQuery) (int, error)
	BuildList(ctx context.Context, query *ent.EmailTemplateQuery) ([]*ent.EmailTemplate, error)
	BuildGetOne(ctx context.Context, query *ent.EmailTemplateQuery) (*ent.EmailTemplate, error)

	// common function
	ValidKeyword(input string, keywordArray []string) error
}

type emailtemplateRepoImpl struct {
	client *ent.Client
}

func NewEmailTemplateRepository(client *ent.Client) EmailTemplateRepository {
	return &emailtemplateRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *emailtemplateRepoImpl) BuildCreate() *ent.EmailTemplateCreate {
	return rps.client.EmailTemplate.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *emailtemplateRepoImpl) BuildUpdate() *ent.EmailTemplateUpdate {
	return rps.client.EmailTemplate.Update().SetUpdatedAt(time.Now())
}

func (rps *emailtemplateRepoImpl) BuildDelete() *ent.EmailTemplateUpdate {
	return rps.client.EmailTemplate.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *emailtemplateRepoImpl) BuildQuery() *ent.EmailTemplateQuery {
	return rps.client.EmailTemplate.Query().Where(emailtemplate.DeletedAtIsNil()).WithRoleEdges(
		func(rq *ent.RoleQuery) {
			rq.Where(role.DeletedAtIsNil())
		},
	)
}

func (rps *emailtemplateRepoImpl) BuildGetOne(ctx context.Context, query *ent.EmailTemplateQuery) (*ent.EmailTemplate, error) {
	return query.First(ctx)
}

func (rps *emailtemplateRepoImpl) BuildBaseQuery() *ent.EmailTemplateQuery {
	return rps.client.EmailTemplate.Query().Where(emailtemplate.DeletedAtIsNil())
}

func (rps *emailtemplateRepoImpl) BuildGet(ctx context.Context, query *ent.EmailTemplateQuery) (*ent.EmailTemplate, error) {
	return query.First(ctx)
}

func (rps *emailtemplateRepoImpl) BuildList(ctx context.Context, query *ent.EmailTemplateQuery) ([]*ent.EmailTemplate, error) {
	return query.All(ctx)
}

func (rps *emailtemplateRepoImpl) BuildCount(ctx context.Context, query *ent.EmailTemplateQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *emailtemplateRepoImpl) BuildExist(ctx context.Context, query *ent.EmailTemplateQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *emailtemplateRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.EmailTemplate) *ent.EmailTemplateUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

func (rps *emailtemplateRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.EmailTemplateUpdateOne) (*ent.EmailTemplate, error) {
	return update.Save(ctx)
}

// mutation
func (rps *emailtemplateRepoImpl) CreateEmailTemplate(ctx context.Context, input ent.NewEmailTemplateInput, roleIds []uuid.UUID, sendTo []string) (*ent.EmailTemplate, error) {
	cc := lo.Map(input.Cc, func(s string, index int) string {
		return s
	})
	bcc := lo.Map(input.Bcc, func(s string, index int) string {
		return s
	})
	create := rps.BuildCreate().
		SetEvent(emailtemplate.Event(input.Event)).
		SetSendTo(sendTo).
		SetSubject(input.Subject).
		SetContent(input.Content).
		SetSignature(input.Signature).
		AddRoleEdgeIDs(roleIds...).
		SetBcc(bcc).
		SetCc(cc)
	return create.Save(ctx)
}

func (rps *emailtemplateRepoImpl) UpdateEmailTemplate(ctx context.Context, record *ent.EmailTemplate, input ent.UpdateEmailTemplateInput, newRoleIds []uuid.UUID, removeRoleIds []uuid.UUID, sendTo []string) (*ent.EmailTemplate, error) {
	cc := lo.Map(input.Cc, func(s string, index int) string {
		return s
	})
	bcc := lo.Map(input.Bcc, func(s string, index int) string {
		return s
	})
	update := rps.BuildUpdateOne(ctx, record).
		SetEvent(emailtemplate.Event(input.Event)).
		SetSendTo(sendTo).
		SetSubject(input.Subject).
		SetContent(input.Content).
		SetSignature(input.Signature).
		AddRoleEdgeIDs(newRoleIds...).RemoveRoleEdgeIDs(removeRoleIds...).
		SetBcc(bcc).
		SetCc(cc)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps emailtemplateRepoImpl) UpdateEmailTemplateStatus(ctx context.Context, record *ent.EmailTemplate, input ent.UpdateEmailTemplateStatusInput) error {
	_, err := record.Update().SetStatus(emailtemplate.Status(*input.Status)).SetUpdatedAt(time.Now().UTC()).Save(ctx)
	return err
}

func (rps *emailtemplateRepoImpl) DeleteEmailTemplate(ctx context.Context, record *ent.EmailTemplate, roleIds []uuid.UUID) (*ent.EmailTemplate, error) {
	update := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveRoleEdgeIDs(roleIds...)
	return update.Save(ctx)
}

// query
func (rps *emailtemplateRepoImpl) GetEmailTemplate(ctx context.Context, id uuid.UUID) (*ent.EmailTemplate, error) {
	query := rps.BuildQuery().Where(emailtemplate.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps emailtemplateRepoImpl) ValidKeyword(input string, keywordArray []string) error {
	re := regexp.MustCompile(`{{\s*([^}]+?)\s*}}`)
	inputMatches := re.FindAllStringSubmatch(input, -1)
	for _, match := range inputMatches {
		fmt.Println("========>", match[1])
		if lo.Contains(keywordArray, match[1]) {
			return nil
		} else {
			prefix := strings.Split(match[1], ":")
			if models.EmailTpErrorString[prefix[0]] != "" {
				return fmt.Errorf(models.EmailTpErrorString[prefix[0]])
			} else {
				return fmt.Errorf("model.email_template.validation.keyword_not_found")
			}
		}
	}
	return nil
}

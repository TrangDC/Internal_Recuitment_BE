package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/outgoingemail"
	"trec/models"

	"github.com/google/uuid"
)

type OutgoingEmailRepository interface {
	CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput) ([]*ent.OutgoingEmail, error)
	CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error)
}

type outgoingEmailRepoImpl struct {
	client *ent.Client
}

func NewOutgoingEmailRepository(client *ent.Client) OutgoingEmailRepository {
	return &outgoingEmailRepoImpl{
		client: client,
	}
}

// Base function
func (rps *outgoingEmailRepoImpl) BuildCreate() *ent.OutgoingEmailCreate {
	return rps.client.OutgoingEmail.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *outgoingEmailRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.OutgoingEmailCreate) ([]*ent.OutgoingEmail, error) {
	return rps.client.OutgoingEmail.CreateBulk(input...).Save(ctx)
}

func (rps *outgoingEmailRepoImpl) BuildUpdate() *ent.OutgoingEmailUpdate {
	return rps.client.OutgoingEmail.Update().SetUpdatedAt(time.Now())
}

func (rps *outgoingEmailRepoImpl) BuildDelete() *ent.OutgoingEmailUpdate {
	return rps.client.OutgoingEmail.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *outgoingEmailRepoImpl) BuildQuery() *ent.OutgoingEmailQuery {
	return rps.client.OutgoingEmail.Query().Where(outgoingemail.DeletedAtIsNil())
}

func (rps *outgoingEmailRepoImpl) BuildGetOne(ctx context.Context, query *ent.OutgoingEmailQuery) (*ent.OutgoingEmail, error) {
	return query.First(ctx)
}

func (rps outgoingEmailRepoImpl) BuildBaseQuery() *ent.OutgoingEmailQuery {
	return rps.client.OutgoingEmail.Query().Where(outgoingemail.DeletedAtIsNil())
}

func (rps *outgoingEmailRepoImpl) BuildGet(ctx context.Context, query *ent.OutgoingEmailQuery) (*ent.OutgoingEmail, error) {
	return query.First(ctx)
}

func (rps *outgoingEmailRepoImpl) BuildList(ctx context.Context, query *ent.OutgoingEmailQuery) ([]*ent.OutgoingEmail, error) {
	return query.All(ctx)
}

func (rps *outgoingEmailRepoImpl) BuildCount(ctx context.Context, query *ent.OutgoingEmailQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *outgoingEmailRepoImpl) BuildExist(ctx context.Context, query *ent.OutgoingEmailQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *outgoingEmailRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.OutgoingEmail) *ent.OutgoingEmailUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

func (rps *outgoingEmailRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.OutgoingEmailUpdateOne) (*ent.OutgoingEmail, error) {
	return update.Save(ctx)
}

func (rps *outgoingEmailRepoImpl) CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput) ([]*ent.OutgoingEmail, error) {
	results := []*ent.OutgoingEmail{}
	createBulk := []*ent.OutgoingEmailCreate{}
	for _, v := range input {
		create := rps.BuildCreate().
			SetTo(v.To).
			SetCc(v.Cc).
			SetBcc(v.Cc).
			SetSubject(v.Subject).
			SetContent(v.Content).
			SetSignature(v.Signature).
			SetEmailTemplateID(v.ID)
		createBulk = append(createBulk, create)
	}
	for _, v := range createBulk {
		record, _ := v.Save(ctx)
		results = append(results, record)
	}
	return results, nil
}

func (rps *outgoingEmailRepoImpl) CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error) {
	query := rps.BuildQuery().Where(outgoingemail.IDEQ(uuid.MustParse(input.ID)))
	record, err := rps.BuildGetOne(ctx, query)
	if err != nil {
		return nil, err
	}
	update := rps.BuildUpdateOne(ctx, record)
	if input.IsSuccess {
		update.SetStatus("sent")
	} else {
		update.SetStatus("failed")
	}
	record, err = rps.BuildSaveUpdateOne(ctx, update)
	if err != nil {
		return nil, err
	}
	return record, nil
}

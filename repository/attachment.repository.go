package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/attachment"

	"github.com/google/uuid"
)

type AttachmentRepository interface {
	// mutation
	CreateAttachment(ctx context.Context, input []*ent.NewAttachmentInput, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error)
	RemoveAttachment(ctx context.Context, relationId uuid.UUID) error
	// query
	GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error)
	GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error)

	// common
	BuildQuery() *ent.AttachmentQuery
	BuildCount(ctx context.Context, query *ent.AttachmentQuery) (int, error)
	BuildList(ctx context.Context, query *ent.AttachmentQuery) ([]*ent.Attachment, error)
	BuildGet(ctx context.Context, query *ent.AttachmentQuery) (*ent.Attachment, error)
}

type attachmentRepoImpl struct {
	client *ent.Client
}

func NewAttachmentRepository(client *ent.Client) AttachmentRepository {
	return &attachmentRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *attachmentRepoImpl) BuildCreate() *ent.AttachmentCreate {
	return rps.client.Attachment.Create().SetUpdatedAt(currentTime).SetCreatedAt(currentTime)
}

func (rps *attachmentRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.AttachmentCreate) ([]*ent.Attachment, error) {
	return rps.client.Attachment.CreateBulk(input...).Save(ctx)
}

func (rps *attachmentRepoImpl) BuildUpdate() *ent.AttachmentUpdate {
	return rps.client.Attachment.Update().SetUpdatedAt(time.Now())
}

func (rps *attachmentRepoImpl) BuildDelete() *ent.AttachmentUpdate {
	return rps.client.Attachment.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *attachmentRepoImpl) BuildQuery() *ent.AttachmentQuery {
	return rps.client.Attachment.Query().Where(attachment.DeletedAtIsNil())
}

func (rps *attachmentRepoImpl) BuildGet(ctx context.Context, query *ent.AttachmentQuery) (*ent.Attachment, error) {
	return query.First(ctx)
}

func (rps *attachmentRepoImpl) BuildList(ctx context.Context, query *ent.AttachmentQuery) ([]*ent.Attachment, error) {
	return query.All(ctx)
}

func (rps *attachmentRepoImpl) BuildCount(ctx context.Context, query *ent.AttachmentQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *attachmentRepoImpl) BuildExist(ctx context.Context, query *ent.AttachmentQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *attachmentRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.Attachment) *ent.AttachmentUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *attachmentRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.AttachmentUpdateOne) (*ent.Attachment, error) {
	return update.Save(ctx)
}

func (rps *attachmentRepoImpl) CreateAttachment(ctx context.Context, input []*ent.NewAttachmentInput, relationId uuid.UUID, relationType attachment.RelationType) (
	[]*ent.Attachment, error) {
	var attachmentCreateBulk []*ent.AttachmentCreate
	for _, attachmentInput := range input {
		attachmentCreateBulk = append(attachmentCreateBulk,
			rps.BuildCreate().SetDocumentID(uuid.MustParse(attachmentInput.DocumentID)).
				SetRelationID(relationId).SetDocumentName(attachmentInput.DocumentName).
				SetRelationType(relationType))
	}
	attachments, err := rps.BuildBulkCreate(ctx, attachmentCreateBulk)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (rps *attachmentRepoImpl) RemoveAttachment(ctx context.Context, relationId uuid.UUID) error {
	return rps.BuildDelete().Where(attachment.RelationIDEQ(relationId)).Exec(ctx)
}

func (rps *attachmentRepoImpl) GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error) {
	return rps.BuildQuery().Where(attachment.ID(attachmentId)).First(ctx)
}

func (rps *attachmentRepoImpl) GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error) {
	return rps.BuildQuery().Where(attachment.RelationIDEQ(relationId), attachment.RelationTypeEQ(relationType)).All(ctx)
}

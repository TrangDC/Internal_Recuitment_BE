package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/attachment"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type AttachmentRepository interface {
	// mutation
	CreateAttachment(ctx context.Context, input []*ent.NewAttachmentInput, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error)
	RemoveAttachment(ctx context.Context, relationId uuid.UUID) error
	RemoveBulkAttachment(ctx context.Context, ids []uuid.UUID) error
	CreateBulkAttachment(ctx context.Context, attachments []*ent.Attachment) error
	// query
	GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error)
	GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error)

	// common
	BuildQuery() *ent.AttachmentQuery
	BuildCount(ctx context.Context, query *ent.AttachmentQuery) (int, error)
	BuildList(ctx context.Context, query *ent.AttachmentQuery) ([]*ent.Attachment, error)
	BuildGet(ctx context.Context, query *ent.AttachmentQuery) (*ent.Attachment, error)

	CreateAndUpdateAttachment(ctx context.Context, relationId uuid.UUID, input []*ent.NewAttachmentInput,
		attachmentRecords []*ent.Attachment, relationType attachment.RelationType) error
	DeleteAllAttachment(ctx context.Context, relationId uuid.UUID) error
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
	return rps.client.Attachment.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *attachmentRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.AttachmentCreate) ([]*ent.Attachment, error) {
	return rps.client.Attachment.CreateBulk(input...).Save(ctx)
}

func (rps *attachmentRepoImpl) BuildUpdate() *ent.AttachmentUpdate {
	return rps.client.Attachment.Update().SetUpdatedAt(time.Now())
}

func (rps *attachmentRepoImpl) BuildDelete() *ent.AttachmentDelete {
	return rps.client.Attachment.Delete()
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

func (rps *attachmentRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.Attachment) *ent.AttachmentUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
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
	_, err := rps.BuildDelete().Where(attachment.RelationIDEQ(relationId)).Exec(ctx)
	return err
}

func (rps *attachmentRepoImpl) RemoveBulkAttachment(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.BuildDelete().Where(attachment.IDIn(ids...)).Exec(ctx)
	return err
}

func (rps attachmentRepoImpl) CreateBulkAttachment(ctx context.Context, attachments []*ent.Attachment) error {
	var createBulk []*ent.AttachmentCreate
	for _, record := range attachments {
		createBulk = append(createBulk,
			rps.client.Attachment.Create().
				SetRelationID(record.RelationID).
				SetRelationType(record.RelationType).
				SetDocumentID(record.DocumentID).
				SetDocumentName(record.DocumentName).
				SetCreatedAt(time.Now().UTC()).
				SetUpdatedAt(time.Now().UTC()),
		)
	}
	_, err := rps.client.Attachment.CreateBulk(createBulk...).Save(ctx)
	return err
}

func (rps *attachmentRepoImpl) GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error) {
	return rps.BuildQuery().Where(attachment.ID(attachmentId)).First(ctx)
}

func (rps *attachmentRepoImpl) GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error) {
	return rps.BuildQuery().Where(attachment.RelationIDEQ(relationId), attachment.RelationTypeEQ(relationType)).All(ctx)
}

func (rps attachmentRepoImpl) createBulkAttachment(ctx context.Context, input []*ent.NewAttachmentInput,
	relationId uuid.UUID, relationType attachment.RelationType) error {
	var recordCreate []*ent.AttachmentCreate
	for _, entity := range input {
		recordCreate = append(recordCreate,
			rps.client.Attachment.Create().
				SetRelationID(relationId).
				SetRelationType(relationType).
				SetDocumentID(uuid.MustParse(entity.DocumentID)).
				SetDocumentName(entity.DocumentName).
				SetCreatedAt(time.Now().UTC()).
				SetUpdatedAt(time.Now().UTC()),
		)
	}
	_, err := rps.client.Attachment.CreateBulk(recordCreate...).Save(ctx)
	return err
}

func (rps attachmentRepoImpl) DeleteAttachment(ctx context.Context, attachmentIds []uuid.UUID) error {
	_, err := rps.client.Attachment.Update().Where(attachment.IDIn(attachmentIds...)).
		SetDeletedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps attachmentRepoImpl) DeleteAllAttachment(ctx context.Context, relationId uuid.UUID) error {
	_, err := rps.client.Attachment.Update().Where(attachment.RelationIDEQ(relationId)).
		SetDeletedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps attachmentRepoImpl) CreateAndUpdateAttachment(ctx context.Context, relationId uuid.UUID, input []*ent.NewAttachmentInput,
	attachmentRecords []*ent.Attachment, relationType attachment.RelationType) error {
	var newInput []*ent.NewAttachmentInput
	var deletedIds []uuid.UUID
	for _, entity := range input {
		if entity.ID == nil || *entity.ID == "" {
			newInput = append(newInput, entity)
		}
	}
	if len(attachmentRecords) > 0 {
		for _, entity := range attachmentRecords {
			_, exist := lo.Find(input, func(record *ent.NewAttachmentInput) bool {
				return record.ID != nil && *record.ID != "" && *record.ID == entity.ID.String()
			})
			if !exist {
				deletedIds = append(deletedIds, entity.ID)
			}
		}
	}
	if len(newInput) > 0 {
		err := rps.createBulkAttachment(ctx, newInput, relationId, relationType)
		if err != nil {
			return err
		}
	}
	if len(deletedIds) > 0 {
		err := rps.DeleteAttachment(ctx, deletedIds)
		if err != nil {
			return err
		}
	}
	return nil
}

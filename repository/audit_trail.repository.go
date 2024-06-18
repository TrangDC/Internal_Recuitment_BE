package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/audittrail"

	"github.com/google/uuid"
)

type AuditTrailRepository interface {
	AuditTrailMutation(ctx context.Context, recordId uuid.UUID, module audittrail.Module, recordChange string, mutationType audittrail.ActionType, note string) error
	//query
	GetAuditTrail(ctx context.Context, auditTrailId uuid.UUID) (*ent.AuditTrail, error)
	BuildQuery() *ent.AuditTrailQuery
	BuildCount(ctx context.Context, query *ent.AuditTrailQuery) (int, error)
	BuildList(ctx context.Context, query *ent.AuditTrailQuery) ([]*ent.AuditTrail, error)
}

type auditTrailRepoImpl struct {
	client *ent.Client
}

func NewAuditTrailRepository(client *ent.Client) AuditTrailRepository {
	return &auditTrailRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *auditTrailRepoImpl) BuildCreate() *ent.AuditTrailCreate {
	return rps.client.AuditTrail.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *auditTrailRepoImpl) BuildUpdate() *ent.AuditTrailUpdate {
	return rps.client.AuditTrail.Update().SetUpdatedAt(time.Now())
}

func (rps *auditTrailRepoImpl) BuildDelete() *ent.AuditTrailUpdate {
	return rps.client.AuditTrail.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *auditTrailRepoImpl) BuildQuery() *ent.AuditTrailQuery {
	return rps.client.AuditTrail.Query().Where(audittrail.DeletedAtIsNil()).WithUserEdge()
}

func (rps *auditTrailRepoImpl) BuildGet(ctx context.Context, query *ent.AuditTrailQuery) (*ent.AuditTrail, error) {
	return query.First(ctx)
}

func (rps *auditTrailRepoImpl) BuildList(ctx context.Context, query *ent.AuditTrailQuery) ([]*ent.AuditTrail, error) {
	return query.All(ctx)
}

func (rps *auditTrailRepoImpl) BuildCount(ctx context.Context, query *ent.AuditTrailQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *auditTrailRepoImpl) BuildExist(ctx context.Context, query *ent.AuditTrailQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *auditTrailRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.AuditTrail) *ent.AuditTrailUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *auditTrailRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.AuditTrailUpdateOne) (*ent.AuditTrail, error) {
	return update.Save(ctx)
}

// query
func (rps *auditTrailRepoImpl) GetAuditTrail(ctx context.Context, auditTrailId uuid.UUID) (*ent.AuditTrail, error) {
	return rps.BuildGet(ctx, rps.BuildQuery().Where(audittrail.ID(auditTrailId)))
}

// common function
func (rps *auditTrailRepoImpl) AuditTrailMutation(ctx context.Context, recordId uuid.UUID, module audittrail.Module, recordChange string, mutationType audittrail.ActionType, note string) error {
	createdById := ctx.Value("user_id").(uuid.UUID)
	_, err := rps.BuildCreate().SetRecordId(recordId).
		SetModule(module).
		SetActionType(mutationType).
		SetRecordChanges(recordChange).
		SetCreatedBy(createdById).
		SetNote(note).Save(ctx)
	return err
}

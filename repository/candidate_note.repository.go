package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidatenote"
	"trec/middleware"

	"github.com/google/uuid"
)

type CandidateNoteRepository interface {
	CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNote, error)
	UpdateCandidateNote(ctx context.Context, candidateNote *ent.CandidateNote, input ent.UpdateCandidateNoteInput) error
	DeleteCandidateNote(ctx context.Context, candidateNote *ent.CandidateNote) error
	GetCandidateNote(ctx context.Context, candidateNoteID uuid.UUID) (*ent.CandidateNote, error)
}

type candidateNoteRepoImpl struct {
	entClient *ent.Client
}

func NewCandidateNoteRepository(entClient *ent.Client) CandidateNoteRepository {
	return &candidateNoteRepoImpl{
		entClient: entClient,
	}
}

// mutation
func (rps *candidateNoteRepoImpl) CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNote, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	createdByID := payload.UserID
	return rps.entClient.CandidateNote.Create().
		SetCandidateID(uuid.MustParse(input.CandidateID)).
		SetCreatedByID(createdByID).
		SetName(strings.TrimSpace(input.Name)).SetDescription(strings.TrimSpace(input.Description)).
		Save(ctx)
}

func (rps *candidateNoteRepoImpl) UpdateCandidateNote(ctx context.Context, candidateNote *ent.CandidateNote, input ent.UpdateCandidateNoteInput) error {
	_, err := candidateNote.Update().
		SetName(strings.TrimSpace(input.Name)).SetDescription(strings.TrimSpace(input.Description)).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps *candidateNoteRepoImpl) DeleteCandidateNote(ctx context.Context, candidateNote *ent.CandidateNote) error {
	currentTime := time.Now().UTC()
	_, err := candidateNote.Update().
		SetDeletedAt(currentTime).SetUpdatedAt(currentTime).
		Save(ctx)
	return err
}

// query
func (rps *candidateNoteRepoImpl) BuildBaseQuery() *ent.CandidateNoteQuery {
	return rps.entClient.CandidateNote.Query().Where(candidatenote.DeletedAtIsNil())
}

func (rps *candidateNoteRepoImpl) BuildQuery() *ent.CandidateNoteQuery {
	return rps.BuildBaseQuery().
		WithCandidateEdge().
		WithCreatedByEdge().
		WithAttachmentEdges()
}

func (rps *candidateNoteRepoImpl) GetCandidateNote(ctx context.Context, candidateNoteID uuid.UUID) (*ent.CandidateNote, error) {
	return rps.BuildQuery().Where(candidatenote.ID(candidateNoteID)).First(ctx)
}

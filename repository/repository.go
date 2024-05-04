package repository

import (
	"context"
	"fmt"
	"trec/ent"

	"github.com/pkg/errors"
)

// Repository is a registry of all repositories
type Repository interface {
	User() UserRepository
	Team() TeamRepository
	HiringJob() HiringJobRepository
	AuditTrail() AuditTrailRepository

	// DoInTx executes the given function in a transaction.
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error
}

// RepoImpl is implementation of Repository
type RepoImpl struct {
	entClient  *ent.Client
	entTx      *ent.Tx
	user       UserRepository
	team       TeamRepository
	hiringJob  HiringJobRepository
	auditTrail AuditTrailRepository
}

// NewRepository creates new repository registry
func NewRepository(entClient *ent.Client) Repository {
	return &RepoImpl{
		entClient:  entClient,
		user:       NewUserRepository(entClient),
		team:       NewTeamRepository(entClient),
		hiringJob:  NewHiringJobRepository(entClient),
		auditTrail: NewAuditTrailRepository(entClient),
	}
}

func (r *RepoImpl) User() UserRepository {
	return r.user
}

func (r *RepoImpl) Team() TeamRepository {
	return r.team
}

func (r *RepoImpl) HiringJob() HiringJobRepository {
	return r.hiringJob
}

func (r *RepoImpl) AuditTrail() AuditTrailRepository {
	return r.auditTrail
}

// DoInTx executes the given function in a transaction.
func (r *RepoImpl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error {
	if r.entTx != nil {
		return errors.WithStack(errors.New("invalid tx state, no nested tx allowed"))
	}

	tx, err := r.entClient.Tx(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	commited := false

	defer func() {
		if commited {
			return
		}
		// rollback if not commited
		_ = tx.Rollback()
	}()

	impl := &RepoImpl{
		entTx:      tx,
		user:       NewUserRepository(tx.Client()),
		team:       NewTeamRepository(tx.Client()),
		hiringJob:  NewHiringJobRepository(tx.Client()),
		auditTrail: NewAuditTrailRepository(tx.Client()),
	}

	if err := txFunc(ctx, impl); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(fmt.Errorf("failed to commit tx: %s", err.Error()))
	}

	commited = true
	return nil
}

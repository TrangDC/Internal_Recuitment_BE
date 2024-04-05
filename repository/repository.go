package repository

import (
	"context"
	"fmt"
	"trec/ent"

	"github.com/pkg/errors"
)

// Repository is a registry of all repositories
type Repository interface {
	JobTitle() PreRepository

	// DoInTx executes the given function in a transaction.
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error
}

// RepoImpl is implementation of Repository
type RepoImpl struct {
	// BankAccountRepository
	entClient *ent.Client
	entTx     *ent.Tx
	jobTitle  PreRepository
}

// NewRepository creates new repository registry
func NewRepository(entClient *ent.Client) Repository {
	return &RepoImpl{
		entClient: entClient,
		jobTitle:  NewPreRepository(entClient),
	}
}

func (r *RepoImpl) JobTitle() PreRepository {
	return r.jobTitle
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
		entTx:    tx,
		jobTitle: NewPreRepository(tx.Client()),
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

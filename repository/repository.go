package repository

import (
	"context"
	"fmt"
	"intern_traning/ent"

	"github.com/pkg/errors"
)

// Repository is a registry of all repositories
type Repository interface {
	User() UserRepository
	News() NewsRepository

	// DoInTx executes the given function in a transaction.
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error
}

// RepoImpl is implementation of Repository
type RepoImpl struct {
	// BankAccountRepository
	entClient *ent.Client
	entTx     *ent.Tx
	user      UserRepository
	news      NewsRepository
}

// NewRepository creates new repository registry
func NewRepository(entClient *ent.Client) Repository {
	return &RepoImpl{
		entClient: entClient,
		user:      NewUserRepository(entClient),
		news:      NewNewsRepository(entClient),
	}
}

func (r *RepoImpl) User() UserRepository {
	return r.user
}

func (r *RepoImpl) News() NewsRepository {
	return r.news
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
		entTx: tx,
		user:  NewUserRepository(tx.Client()),
		news:  NewNewsRepository(tx.Client()),
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

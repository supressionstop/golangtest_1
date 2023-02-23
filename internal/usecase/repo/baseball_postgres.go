package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"softpro6/internal/usecase"
	"softpro6/pkg/postgres"
	"time"
)

type BaseballRepo struct {
	pg *postgres.Postgres
}

func NewBaseballRepo(pg *postgres.Postgres) *BaseballRepo {
	return &BaseballRepo{pg: pg}
}

func (r BaseballRepo) Store(ctx context.Context, sport usecase.Sport) error {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("BaseballRepo - Postgres - Store - r.pg.Pool.BeginTx - unable to start transaction: %v", err)
	}
	defer func() {
		err = r.finishTransaction(ctx, err, tx)
	}()

	sql, args, err := r.pg.Builder.
		Insert("baseball").
		Columns("rate", "created_at", "provider").
		Values(sport.Rate().Value(), sport.CreatedAt(), sport.Rate().Provider()).
		ToSql()
	if err != nil {
		return fmt.Errorf("BaseballRepo - Postgres - Store - r.pg.Builder: %w", err)
	}

	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("BaseballRepo - Postgres - Store - r.pg.Pool.Exec: %w", err)
	}

	return nil
}

func (r BaseballRepo) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("%s: %s", rollbackErr, err)
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			return fmt.Errorf("failed to commit tx: %s", err)
		}

		return nil
	}
}

func (r BaseballRepo) GetRecent(ctx context.Context) (usecase.Sport, error) {
	return nil, nil
}
func (r BaseballRepo) IsSynced(ctx context.Context, after time.Time) (bool, error) {
	return false, nil
}

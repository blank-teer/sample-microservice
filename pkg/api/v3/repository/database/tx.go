package database

import (
	"context"
	"database/sql"
	"fmt"

	"order-manager/pkg/api/v3/repository"
)

type ctxKeyTx struct{}

type tx struct {
	*sql.DB
}

func NewTx(db *sql.DB) repository.Tx {
	return tx{
		DB: db,
	}
}

func (t tx) Do(ctx context.Context, fn func(context.Context) error) (err error) {
	tx, errBeginTx := t.BeginTx(ctx, nil)
	if errBeginTx != nil {
		err = fmt.Errorf("database.tx.Do.BeginTx: failed: %w", errBeginTx)
		return
	}

	defer func() {
		if err == nil {
			return
		}

		if errRollback := tx.Rollback(); errRollback != nil {
			err = fmt.Errorf("database.tx.Do.Rollback: failed: %w: %s", errRollback, err)
		}
	}()

	ctx = context.WithValue(ctx, ctxKeyTx{}, tx)

	if errFn := fn(ctx); errFn != nil {
		err = fmt.Errorf("database.tx.Do.fn: failed: %w", errFn)
		return
	}

	if errCommit := tx.Commit(); errCommit != nil {
		err = fmt.Errorf("database.tx.Do.Commit: failed: %w", errCommit)
		return
	}

	return
}

func getTxFrom(ctx context.Context) Driver {
	if tx, ok := ctx.Value(ctxKeyTx{}).(Driver); ok {
		return tx
	}
	return nil
}

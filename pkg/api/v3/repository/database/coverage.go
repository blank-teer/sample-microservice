package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/repository"
	"order-manager/pkg/cfg"
)

type repoCoverage struct {
	db  *sql.DB
	cfg cfg.Postgres
}

func NewRepoCoverage(db *sql.DB, cfg cfg.Postgres) repository.Coverage {
	return repoCoverage{
		db:  db,
		cfg: cfg,
	}
}

func (r repoCoverage) Create(ctx context.Context, emissionID uuid.UUID, coverage models.Coverage) error {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    INSERT INTO "coverage" (emission_id, 
                            order_id, 
                            amount) 
    VALUES ($1, $2, $3)`

	var ex Executor = r.db
	if tx := getTxFrom(ctx); tx != nil {
		ex = tx
	}

	_, err := ex.ExecContext(queryCtx, query,
		emissionID,
		coverage.OrderID,
		coverage.Amount,
	)
	if err != nil {
		return err
	}
	return nil
}

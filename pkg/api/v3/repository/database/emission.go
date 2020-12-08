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

type repoEmission struct {
	db  *sql.DB
	cfg cfg.Postgres
}

func NewRepoEmission(db *sql.DB, cfg cfg.Postgres) repository.Emission {
	return repoEmission{
		db:  db,
		cfg: cfg,
	}
}

func (r repoEmission) Create(ctx context.Context, e models.Emission) error {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    INSERT INTO "emission" (id, 
                            type,
                            gold_unit,    
                            gold_amount,  
                            base_currency,
                            base_currency_amount,  
                            rate_base_to_unit,         
                            reference,    
                            details)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	var ex Executor = r.db
	if tx := getTxFrom(ctx); tx != nil {
		ex = tx
	}

	_, err := ex.ExecContext(queryCtx, query,
		e.ID,
		e.Type,
		e.GoldUnit,
		e.GoldAmount,
		e.BaseCurrency,
		e.BaseCurrencyAmount,
		e.RateBaseToUnit,
		e.Reference,
		e.Details,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r repoEmission) Get(ctx context.Context, emissionID uuid.UUID) (models.Emission, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT num,          
           id,           
           type,         
           gold_unit,    
           gold_amount,  
           base_currency,
           base_currency_amount,  
           rate_base_to_unit,         
           reference,    
           details,      
           created_at,   
           updated_at
      FROM emission 
     WHERE id = $1`

	row := r.db.QueryRowContext(queryCtx, query, emissionID)

	var e models.Emission
	if err := row.Scan(
		&e.Num,
		&e.ID,
		&e.Type,
		&e.GoldUnit,
		&e.GoldAmount,
		&e.BaseCurrency,
		&e.BaseCurrencyAmount,
		&e.RateBaseToUnit,
		&e.Reference,
		&e.Details,
		&e.CreatedAt,
		&e.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Emission{}, nil
		}
		return models.Emission{}, err
	}
	return e, nil
}

func (r repoEmission) GetSummary(ctx context.Context) ([]models.EmissionSummary, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT gold_unit                             AS unit, 
           coalesce(sum(gold_amount),0)          AS gold, 
           coalesce(sum(base_currency_amount),0) AS fiat 
      FROM emission
     WHERE type = $1
  GROUP BY gold_unit`

	rows, err := r.db.QueryContext(queryCtx, query, models.EmissionTypeDeposit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		s  models.EmissionSummary
		ss = make([]models.EmissionSummary, 0)
	)

	for rows.Next() {
		err := rows.Scan(
			&s.Unit,
			&s.Gold,
			&s.Fiat,
		)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}

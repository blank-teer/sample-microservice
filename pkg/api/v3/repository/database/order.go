package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/errs"
	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/repository"
	"order-manager/pkg/api/v3/repository/database/qbuilder"
	"order-manager/pkg/cfg"
)

type repoOrder struct {
	logger log.Logger
	db     *sql.DB
	cfg    cfg.Postgres
}

func NewRepoOrder(lg log.Logger, db *sql.DB, cfg cfg.Postgres) repository.Order {
	return repoOrder{
		logger: lg,
		db:     db,
		cfg:    cfg,
	}
}

func (r repoOrder) Create(ctx context.Context, o models.Order) error {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    INSERT INTO "orders" (id,
                          type,
                          provider,
                          user_id,
                          wallet_address,
                          base_currency,
                          base_currency_amount,
                          base_currency_price,
                          pay_currency,
                          pay_currency_amount,
                          pay_currency_price,
                          rate_pay_to_base,
                          coins,
                          fee)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	var ex Executor = r.db
	if tx := getTxFrom(ctx); tx != nil {
		ex = tx
	}

	_, err := ex.ExecContext(queryCtx, query,
		o.ID,
		o.Type,
		o.Provider,
		o.UserID,
		o.WalletAddress,
		o.BaseCurrency,
		o.BaseCurrencyAmount,
		o.BaseCurrencyPrice,
		o.PayCurrency,
		o.PayCurrencyAmount,
		o.PayCurrencyPrice,
		o.RatePayToBase,
		o.Coins,
		o.Fee,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r repoOrder) Exists(ctx context.Context, orderID uuid.UUID) error {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT exists(SELECT 1
                  FROM orders
                  WHERE id = $1)`

	row := r.db.QueryRowContext(queryCtx, query, orderID)

	var e bool
	if err := row.Scan(&e); err != nil {
		return err
	}
	if !e {
		return errs.ErrNotFound
	}
	return nil
}

func (r repoOrder) Get(ctx context.Context, orderID uuid.UUID) (models.Order, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT id,
           type,
           status,
           provider,
           user_id,
           wallet_address,
           base_currency,
           base_currency_amount,
           base_currency_price,
           pay_currency,
           pay_currency_amount,
           pay_currency_price,
           rate_pay_to_base,
           coins,
           fee,
           created_at,
           updated_at
      FROM "orders"
     WHERE id = $1`

	row := r.db.QueryRowContext(queryCtx, query, orderID)

	var o models.Order
	if err := row.Scan(
		&o.ID,
		&o.Type,
		&o.Status,
		&o.Provider,
		&o.UserID,
		&o.WalletAddress,
		&o.BaseCurrency,
		&o.BaseCurrencyAmount,
		&o.BaseCurrencyPrice,
		&o.PayCurrency,
		&o.PayCurrencyAmount,
		&o.PayCurrencyPrice,
		&o.RatePayToBase,
		&o.Coins,
		&o.Fee,
		&o.CreatedAt,
		&o.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, errs.ErrNotFound
		}
		return models.Order{}, err
	}
	return o, nil
}

func (r repoOrder) GetList(ctx context.Context, ff []repository.Filter, cc []repository.Condition) ([]models.Order, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT o.id,
           o.type,
           o.status,
           o.provider,
           o.user_id,
           o.wallet_address,
           o.base_currency,
           o.base_currency_amount,
           o.base_currency_price,
           o.pay_currency,
           o.pay_currency_amount,
           o.pay_currency_price,
           o.rate_pay_to_base,
           o.coins,
           o.fee,
           o.created_at,
           o.updated_at
      FROM "orders" AS o`

	if len(cc) > 0 {
		for _, c := range cc {
			if _, ok := c.(qbuilder.ConditionOrdersCoverage); ok {
				query += `
    LEFT JOIN (SELECT order_id, amount
               FROM coverage
               GROUP BY order_id, amount) AS c
           ON c.order_id = o.id`

				switch c {
				case qbuilder.ConditionOrdersCoverageNot:
					f := qbuilder.NewFilter("OrderID", qbuilder.TableCoverage.As("c"), qbuilder.ConditionIs, nil)
					ff = append(ff, f)
				case qbuilder.ConditionOrdersCoverageFully:
					f := qbuilder.NewFilter("Amount", qbuilder.TableCoverage.As("c"), qbuilder.ConditionEquals, qbuilder.Target{FieldName: "BaseCurrencyAmount", Table: qbuilder.TableOrders.As("o")})
					ff = append(ff, f)
				case qbuilder.ConditionOrdersCoveragePartially:
					f := qbuilder.NewFilter("Amount", qbuilder.TableCoverage.As("c"), qbuilder.ConditionLessThan, qbuilder.Target{FieldName: "BaseCurrencyAmount", Table: qbuilder.TableOrders.As("o")})
					ff = append(ff, f)
				}
			}
		}
	}

	qb := qbuilder.NewQBuilder(r.logger, query, ff...)
	query = qb.Build()

	r.logger.Info(query)

	rows, err := r.db.QueryContext(queryCtx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		o  models.Order
		oo = make([]models.Order, 0)
	)

	for rows.Next() {
		err := rows.Scan(
			&o.ID,
			&o.Type,
			&o.Status,
			&o.Provider,
			&o.UserID,
			&o.WalletAddress,
			&o.BaseCurrency,
			&o.BaseCurrencyAmount,
			&o.BaseCurrencyPrice,
			&o.PayCurrency,
			&o.PayCurrencyAmount,
			&o.PayCurrencyPrice,
			&o.RatePayToBase,
			&o.Coins,
			&o.Fee,
			&o.CreatedAt,
			&o.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		oo = append(oo, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return oo, nil
}

func (r repoOrder) GetSummary(ctx context.Context) ([]models.OrderSummary, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    SELECT coalesce(sum(o.base_currency_amount - coalesce(c.amount, 0)),0) AS fiat,
	coalesce(sum(o.coins),0)                                        AS coins
      FROM orders AS o
 LEFT JOIN (SELECT order_id, amount
            FROM coverage
            GROUP BY order_id, amount) AS c 
        ON c.order_id = o.id
     WHERE o.type = $1
       AND o.status IN ($2, $3)
       AND (c.order_id IS NULL OR c.amount < o.base_currency_amount)`

	rows, err := r.db.QueryContext(queryCtx, query,
		models.OrderTypeDeposit,
		models.OrderStatusPaid,
		models.OrderStatusIssued,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		s  models.OrderSummary
		ss = make([]models.OrderSummary, 0)
	)

	for rows.Next() {
		err := rows.Scan(
			&s.Fiat,
			&s.Coins,
		)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func (r repoOrder) UpdateStatus(ctx context.Context, orderID uuid.UUID, s models.OrderStatus) error {
	queryCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.cfg.QueryTimeout))
	defer cancel()

	query := `
    UPDATE "orders"
       SET status = $1,
           updated_at = $2
     WHERE id = $3`

	var ex Executor = r.db
	if tx := getTxFrom(ctx); tx != nil {
		ex = tx
	}

	_, err := ex.ExecContext(queryCtx, query,
		s,
		time.Now().Unix(),
		orderID,
	)
	if err != nil {
		return fmt.Errorf("order id '%s': failed to set '%s' status", orderID, err)
	}
	return nil
}

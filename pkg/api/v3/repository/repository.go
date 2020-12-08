package repository

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"order-manager/pkg/api/v3/models"
)

type Tx interface {
	Do(ctx context.Context, fn func(context.Context) error) (err error)
}

type Filter interface {
	Apply(*strings.Builder) error
}

type Condition interface {
	Validate() error
}

type Order interface {
	Create(ctx context.Context, o models.Order) error
	Exists(ctx context.Context, orderID uuid.UUID) error
	Get(ctx context.Context, orderID uuid.UUID) (models.Order, error)
	GetList(ctx context.Context, filters []Filter, conditions []Condition) ([]models.Order, error)
	GetSummary(ctx context.Context) ([]models.OrderSummary, error)
	UpdateStatus(ctx context.Context, orderID uuid.UUID, status models.OrderStatus) error
}

type Emission interface {
	Create(ctx context.Context, e models.Emission) error
	Get(ctx context.Context, emissionID uuid.UUID) (models.Emission, error)
	GetSummary(ctx context.Context) ([]models.EmissionSummary, error)
}

type Coverage interface {
	Create(ctx context.Context, emissionID uuid.UUID, coverage models.Coverage) error
}

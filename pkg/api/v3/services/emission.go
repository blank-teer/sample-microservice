package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/scc/go-common/log"
	txmSender "github.com/scc/tx-manager/pkg/api/v3/services/sender"

	"order-manager/pkg/api/v3/errs"
	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/repository"
)

type Emission interface {
	Create(ctx context.Context, e models.Emission) (emissionID uuid.UUID, err error)
	Get(ctx context.Context, emissionID string) (models.Emission, error)
	GetSummary(ctx context.Context) ([]models.EmissionSummary, error)
}

type serviceEmission struct {
	logger       log.Logger
	repoEmission repository.Emission
	repoOrder    repository.Order
	repoCoverage repository.Coverage
	repoTx       repository.Tx
	sender       *txmSender.Service
}

func NewEmission(l log.Logger, rE repository.Emission, rO repository.Order, rC repository.Coverage, rTx repository.Tx) Emission {
	return serviceEmission{
		logger:       l,
		repoEmission: rE,
		repoOrder:    rO,
		repoCoverage: rC,
		repoTx:       rTx,
	}
}

func (s serviceEmission) Create(ctx context.Context, e models.Emission) (uuid.UUID, error) {
	e.ID = uuid.New()
	if err := e.Validate(); err != nil {
		return uuid.UUID{}, err
	}

	err := s.repoTx.Do(ctx, func(ctx context.Context) error {

		if err := s.repoEmission.Create(ctx, e); err != nil {
			return err
		}

		for _, c := range e.Coverage {
			if err := s.repoCoverage.Create(ctx, e.ID, c); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return e.ID, nil
}

func (s serviceEmission) Get(ctx context.Context, emissionID string) (models.Emission, error) {
	eid, err := uuid.Parse(emissionID)
	if err != nil {
		return models.Emission{}, fmt.Errorf("%w: emission id: %s: %s", errs.ErrNotValid, emissionID, err)
	}
	return s.repoEmission.Get(ctx, eid)
}

func (s serviceEmission) GetSummary(ctx context.Context) ([]models.EmissionSummary, error) {
	return s.repoEmission.GetSummary(ctx)
}

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	gcBroker "github.com/scc/go-common/broker"
	"github.com/scc/go-common/log"
	txmModels "github.com/scc/tx-manager/pkg/api/v3/models"
	txmSender "github.com/scc/tx-manager/pkg/api/v3/services/psender"

	"order-manager/pkg/api/v3/delivery/broker"
	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers"
	"order-manager/pkg/api/v3/errs"
	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/repository"
	"order-manager/pkg/api/v3/repository/database/qbuilder"
)

type Order interface {
	Create(ctx context.Context, o models.Order) (oderID uuid.UUID, err error)
	Done(ctx context.Context, orderID uuid.UUID) error
	Get(ctx context.Context, orderID string) (models.Order, error)
	GetList(ctx context.Context, by FilterBy) ([]models.Order, error)
	GetSummary(ctx context.Context) (models.OrderSummary, error)
	UpdateStatus(ctx context.Context, orderID uuid.UUID, st models.OrderStatus) error
}

type serviceOrder struct {
	logger log.Logger
	repo   repository.Order
	sender *txmSender.Service
	broker broker.Broker
}

func NewOrder(l log.Logger, r repository.Order, s *txmSender.Service, b broker.Broker) Order {
	return serviceOrder{
		logger: l,
		repo:   r,
		sender: s,
		broker: b,
	}
}

func (s serviceOrder) Create(ctx context.Context, o models.Order) (uuid.UUID, error) {
	o.ID = uuid.New()

	if err := o.Validate(); err != nil {
		return uuid.UUID{}, err
	}

	if err := s.repo.Create(ctx, o); err != nil {
		return uuid.UUID{}, err
	}
	return o.ID, nil
}

func (s serviceOrder) Done(ctx context.Context, orderID uuid.UUID) error {
	if err := s.repo.UpdateStatus(ctx, orderID, models.OrderStatusPaid); err != nil {
		return err
	}

	o, err := s.repo.Get(ctx, orderID)
	if err != nil {
		return err
	}

	m := txmModels.CreateIssue{
		WalletAddress: o.WalletAddress,
		Asset:         models.AssetVC,
		Amount:        uint64(o.Coins),
	}

	done := make(chan struct{})
	sanitize := func(sub gcBroker.Subscription) {
		select {
		case <-done:
			sub.Unsubscribe()
		case <-time.After(1 * time.Minute):
			sub.Unsubscribe()
		}
	}

	callback := func(_ string, data []byte) {
		var r gcBroker.Response
		if err := helpers.ParsePayload(bytes.NewBuffer(data), &r); err != nil {
			s.logger.Errorf("services.Done.ParsePayload: %s", err)
			_ = s.repo.UpdateStatus(ctx, orderID, models.OrderStatusFailed)
			return
		}

		rStr, _ := json.MarshalIndent(r, "", "\t")
		s.logger.Infof("services.Done.CreateIssue: order id '%s': tx performing result: %s\n", o.ID, string(rStr))

		if r.Code != gcBroker.StatusOk {
			_ = s.repo.UpdateStatus(ctx, orderID, models.OrderStatusFailed)
		}

		done <- struct{}{}
	}

	replyTopic := uuid.New().String()
	sub, err := s.broker.Subscribe(replyTopic, callback)
	if err != nil {
		return err
	}
	go sanitize(sub)

	r, err := s.sender.CreateIssue(ctx, uuid.New().String(), gcBroker.Info{ReplyTo: replyTopic}, m)

	rStr, _ := json.MarshalIndent(r, "", "\t")
	s.logger.Infof("services.Done.CreateIssue: order id '%s': tx pushing result: %s\n", o.ID, string(rStr))

	if (r != nil && r.Code != 200) || err != nil {
		if errUpdSt := s.repo.UpdateStatus(ctx, orderID, models.OrderStatusFailed); errUpdSt != nil {
			return fmt.Errorf("%s: %s", err, errUpdSt)
		}
		return err
	}

	return s.repo.UpdateStatus(ctx, orderID, models.OrderStatusIssued)
}

func (s serviceOrder) Get(ctx context.Context, orderID string) (models.Order, error) {
	oid, err := uuid.Parse(orderID)
	if err != nil {
		return models.Order{}, fmt.Errorf("%w: order id: %s: %s", errs.ErrNotValid, orderID, err)
	}

	return s.repo.Get(ctx, oid)
}

type FilterBy struct {
	UserID       string
	OrderType    string
	CoverageRate string
}

func (s serviceOrder) GetList(ctx context.Context, by FilterBy) ([]models.Order, error) {
	var ff []repository.Filter
	var cc []repository.Condition

	if uid, err := strconv.Atoi(by.UserID); err == nil {
		f := qbuilder.NewFilter("UserID", qbuilder.TableOrders.As("o"), qbuilder.ConditionEquals, uid)
		ff = append(ff, f)
	}

	if err := models.OrderType(by.OrderType).Validate(); err == nil {
		f := qbuilder.NewFilter("Type", qbuilder.TableOrders.As("o"), qbuilder.ConditionEquals, by.OrderType)
		ff = append(ff, f)
	}

	if err := qbuilder.ConditionOrdersCoverage(by.CoverageRate).Validate(); err == nil {
		cc = append(cc, qbuilder.ConditionOrdersCoverage(by.CoverageRate))
	}

	return s.repo.GetList(ctx, ff, cc)
}

func (s serviceOrder) GetSummary(ctx context.Context) (models.OrderSummary, error) {
	summary, err := s.repo.GetSummary(ctx)
	if err != nil {
		return models.OrderSummary{}, err
	}

	return summary[0], nil
}

func (s serviceOrder) UpdateStatus(ctx context.Context, orderID uuid.UUID, st models.OrderStatus) error {
	if err := st.Validate(); err != nil {
		return err
	}

	if err := s.repo.Exists(ctx, orderID); err != nil {
		return err
	}

	return s.repo.UpdateStatus(ctx, orderID, st)
}

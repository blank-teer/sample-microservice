package listeners

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"
	gcBroker "github.com/scc/go-common/broker"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/delivery/broker"
	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers"
	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/services"
)

type listenerOrders struct {
	logger  log.Logger
	broker  broker.Broker
	service services.Order
}

func NewOrders(l log.Logger, b broker.Broker, s services.Order) broker.Listener {
	return listenerOrders{
		logger:  l,
		broker:  b,
		service: s,
	}
}

func (l listenerOrders) Listen() error {
	if _, err := l.broker.Subscribe(gcBroker.SubjectOrdDepositVC, l.Create); err != nil {
		return fmt.Errorf("listeners.Listen.Subscribe: %s: %w", gcBroker.SubjectOrdDepositVC, err)
	}
	if _, err := l.broker.Subscribe(gcBroker.SubjectOrdDepositVCDone, l.Done); err != nil {
		return fmt.Errorf("listeners.Listen.Subscribe: %s: %w", gcBroker.SubjectOrdDepositVCDone, err)
	}
	if _, err := l.broker.Subscribe(gcBroker.SubjectOrdSummaryFiat, l.GetSummary); err != nil {
		return fmt.Errorf("listeners.Listen.Subscribe: %s: %w", "order.manager.vc.summary", err)
	}
	return nil
}

type ResponsePayloadCreate struct {
	OrderID uuid.UUID `json:"orderID"`
	Status  int       `json:"status"`
}

func (l listenerOrders) Create(replySubject string, data []byte) {
	ctx := context.Background()

	var o models.Order
	if err := helpers.ParsePayload(bytes.NewBuffer(data), &o); err != nil {
		l.logger.Errorf("listeners.Create.Unmarshal: %s", err)
		l.broker.RespondError(replySubject, err)
		return
	}

	orderID, err := l.service.Create(ctx, o)
	if err != nil {
		l.logger.Errorf("listeners.service.Create: %s", err)
		l.broker.RespondError(replySubject, err)
		return
	}

	rPld := ResponsePayloadCreate{
		OrderID: orderID,
		Status:  200,
	}

	l.logger.Infof("listeners.Create: done: %+v", rPld)
	l.broker.RespondSuccess(replySubject, "", rPld)
}

type RequestDoneOrder struct {
	OrderID uuid.UUID `json:"id"`
}

func (l listenerOrders) Done(_ string, data []byte) {
	ctx := context.Background()

	var p RequestDoneOrder
	if err := helpers.ParsePayload(bytes.NewBuffer(data), &p); err != nil {
		l.logger.Errorf("listeners.Done.ParsePayload: %s", err)
		return
	}

	if err := l.service.Done(ctx, p.OrderID); err != nil {
		l.logger.Errorf("listeners.service.Done: %s", err)
		return
	}

	l.logger.Infof("listeners.Done: done: order id '%s': async issue requested", p.OrderID)
}

func (l listenerOrders) GetSummary(replySubject string, _ []byte) {
	ctx := context.Background()

	summary, err := l.service.GetSummary(ctx)
	if err != nil {
		l.broker.RespondError(replySubject, err)
		return
	}

	l.broker.RespondSuccess(replySubject, "", summary)
}

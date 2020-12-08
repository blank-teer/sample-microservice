package listeners

import (
	"context"
	"fmt"

	gcBroker "github.com/scc/go-common/broker"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/delivery/broker"
	"order-manager/pkg/api/v3/services"
)

type listenerEmission struct {
	logger  log.Logger
	broker  broker.Broker
	service services.Emission
}

func NewEmission(l log.Logger, b broker.Broker, s services.Emission) broker.Listener {
	return listenerEmission{
		logger:  l,
		broker:  b,
		service: s,
	}
}

func (l listenerEmission) Listen() error {
	if _, err := l.broker.Subscribe(gcBroker.SubjectOrdSummaryGold, l.GetSummary); err != nil {
		return fmt.Errorf("listeners.Listen.Subscribe: %s: %w", "order.manager.gold.summary", err)
	}
	return nil
}

func (l listenerEmission) GetSummary(replySubject string, _ []byte) {
	ctx := context.Background()

	summary, err := l.service.GetSummary(ctx)
	if err != nil {
		l.broker.RespondError(replySubject, err)
		return
	}

	l.broker.RespondSuccess(replySubject, "", summary)
}

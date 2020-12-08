package listeners

import (
	"fmt"

	gcBroker "github.com/scc/go-common/broker"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/delivery/broker"
)

type listenerMeta struct {
	logger log.Logger
	broker broker.Broker
}

func NewMeta(l log.Logger, b broker.Broker) broker.Listener {
	return listenerMeta{
		logger: l,
		broker: b,
	}
}

func (l listenerMeta) Listen() error {
	if _, err := l.broker.Subscribe(gcBroker.SubjectOrdServicePing, l.Ping); err != nil {
		return fmt.Errorf("listeners.Listen.Subscribe: %s: %w", gcBroker.SubjectOrdServicePing, err)
	}
	return nil
}

func (l listenerMeta) Ping(replySubject string, _ []byte) {
	l.broker.RespondSuccess(replySubject, "", "pong")
}

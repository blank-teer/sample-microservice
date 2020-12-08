package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	gcBroker "github.com/scc/go-common/broker"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/errs"
	"order-manager/pkg/cfg"
)

type Broker struct {
	gcBroker.Broker
	logger    log.Logger
	config    cfg.MessageBroker
	listeners []Listener
}

func New(c cfg.MessageBroker, l log.Logger) (Broker, error) {
	b, err := gcBroker.RetryOpen(c.ConnectionAddr, c.Name, c.Pass, c.ServiceName, c.MaxRetries, c.RetryDelayMs, l)
	if err != nil {
		return Broker{}, fmt.Errorf("failed to open conn: %s", err)
	}
	return Broker{
		Broker: b,
		logger: l,
		config: c,
	}, nil
}

func (b *Broker) AddListeners(ll ...Listener) {
	b.listeners = append(b.listeners, ll...)
}

var ErrGeneric = errors.New("something went wrong, see app logs (order-manager) for details")

func (b Broker) RespondError(subj string, err error) {
	resp := gcBroker.Response{
		Code:    resolveCode(err),
		Message: err.Error(),
	}

	data, err := json.Marshal(resp)
	if err != nil {
		b.logger.Errorf("broker.RespondError.Marshal: %s", err)
		return
	}

	if err := b.Publish(subj, data); err != nil {
		b.logger.Errorf("broker.RespondError.Publish: %s", err)
		return
	}

	b.logger.Infof("broker.RespondError.Publish: done: %+v", resp)
}

func (b Broker) RespondSuccess(subj, msg string, pld interface{}) {
	bts, err := json.Marshal(pld)
	if err != nil {
		b.logger.Errorf("broker.RespondSuccess.Marshal: %s", err)
		b.RespondError(subj, err)
		return
	}

	resp := gcBroker.Response{
		Code:    gcBroker.StatusOk,
		Message: msg,
		Payload: bts,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		b.logger.Errorf("broker.RespondSuccess.Marshal: %s", err)
		b.RespondError(subj, err)
		return
	}

	if err := b.Publish(subj, data); err != nil {
		b.logger.Errorf("broker.RespondSuccess.Publish: %s", err)
		b.RespondError(subj, err)
		return
	}

	b.logger.Infof("broker.RespondSuccess.Publish: done: %+v", pld)
}

func resolveCode(err error) int {
	if errors.Is(err, errs.ErrNotFound) {
		return gcBroker.StatusNotFound
	}
	if errors.Is(err, errs.ErrNotValid) {
		return gcBroker.StatusBadRequest
	}
	return gcBroker.StatusInternalServiceError
}

// Run starts a Broker
func (b Broker) Run() error {
	for _, l := range b.listeners {
		if err := l.Listen(); err != nil {
			b.logger.Errorf("broker.Run.Listen: %s", err)
			return err
		}
	}
	return nil
}

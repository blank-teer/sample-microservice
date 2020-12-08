package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/scc/go-common/log"
	"github.com/scc/go-common/storage"
	"github.com/scc/tx-manager/pkg/api/v3/services/psender"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"order-manager/pkg/api/v3/delivery/broker"
	"order-manager/pkg/api/v3/delivery/broker/listeners"
	"order-manager/pkg/api/v3/delivery/rest"
	"order-manager/pkg/api/v3/delivery/rest/handlers"
	"order-manager/pkg/api/v3/repository"
	"order-manager/pkg/api/v3/repository/database"
	"order-manager/pkg/api/v3/services"
	"order-manager/pkg/cfg"
)

func NewConfig() (*cfg.Cfg, error) {
	cfgPath := flag.String("c", cfg.FilePath, "configuration file")
	flag.Parse()

	c, err := cfg.Init(*cfgPath)
	if err != nil {
		fmt.Printf("failed to init cfg: %s", err)
		return nil, err
	}

	return c, err
}

func NewLogger(cfg *cfg.Cfg) log.Logger {
	logOut := log.Output(os.Stdout)

	logOptions := []log.Option{
		logOut,
		log.Tags(map[string]interface{}{"service": serviceName}),
		log.Level(cfg.Logger.DebugLevel),
		log.Formatter(log.Format(cfg.Logger.LogFormat), false, "2006-01-02 15:04:05"),
	}

	if cfg.Logger.IncludeCallerMethod {
		logOptions = append(logOptions, log.WithCallerReporting())
	}

	l := log.New(logOptions...)
	l.Infof("starting %s service", serviceName)

	return l
}

func NewAPI(c *cfg.Cfg, l log.Logger, hMeta handlers.Meta, hOrder handlers.Order, hEmission handlers.Emission) rest.API {
	return rest.New(c.APIServer, l, hMeta, hOrder, hEmission)
}

func NewBroker(c *cfg.Cfg, l log.Logger) (broker.Broker, error) {
	return broker.New(c.MessageBroker, l)
}

func NewPostgres(c *cfg.Cfg, l log.Logger) (*sql.DB, error) {
	db, err := storage.RetryOpen("postgres", c.Storage.Postgres.ConnectionString, c.Storage.Postgres.MaxRetries, c.Storage.Postgres.RetryDelay, l)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to the database")
	}
	return db.DB(), nil
}

func NewRepoOrder(db *sql.DB, l log.Logger, c *cfg.Cfg) repository.Order {
	return database.NewRepoOrder(l, db, c.Storage.Postgres)
}

func NewRepoEmission(db *sql.DB, l log.Logger, c *cfg.Cfg) repository.Emission {
	return database.NewRepoEmission(db, c.Storage.Postgres)
}

func NewRepoCoverage(db *sql.DB, l log.Logger, c *cfg.Cfg) repository.Coverage {
	return database.NewRepoCoverage(db, c.Storage.Postgres)
}

func NewRepoTx(db *sql.DB) repository.Tx {
	return database.NewTx(db)
}

func NewTxmSender(b broker.Broker) *psender.Service {
	return psender.NewService(b)
}

func NewServiceOrder(
	rOrder repository.Order,
	txms *psender.Service,
	b broker.Broker,
	l log.Logger,
) services.Order {
	return services.NewOrder(l, rOrder, txms, b)
}

func NewServiceEmission(
	rEmission repository.Emission,
	rOrder repository.Order,
	rCoverage repository.Coverage,
	rTx repository.Tx,
	l log.Logger,
) services.Emission {
	return services.NewEmission(l, rEmission, rOrder, rCoverage, rTx)
}

func NewHandlerMeta(l log.Logger) handlers.Meta {
	return handlers.NewMeta(l)
}

func NewHandlerOrder(sOrder services.Order, l log.Logger) handlers.Order {
	return handlers.NewOrders(l, sOrder)
}

func NewHandlerEmission(sEmission services.Emission, l log.Logger) handlers.Emission {
	return handlers.NewEmission(l, sEmission)
}

func NewBrokerListeners(b broker.Broker, l log.Logger, sOrder services.Order, sEmission services.Emission) []broker.Listener {
	return []broker.Listener{
		listeners.NewMeta(l, b),
		listeners.NewOrders(l, b, sOrder),
		listeners.NewEmission(l, b, sEmission),
	}
}

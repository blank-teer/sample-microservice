package main

import (
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

const (
	serviceName = "order-manager"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewAPI,
			NewBroker,
			NewPostgres,
			NewRepoOrder,
			NewRepoEmission,
			NewRepoCoverage,
			NewRepoTx,
			NewTxmSender,
			NewServiceOrder,
			NewServiceEmission,
			NewHandlerMeta,
			NewHandlerOrder,
			NewHandlerEmission,
			NewBrokerListeners,
		),

		fx.Invoke(
			Migrate,
			RunBroker,
			RunAPI,
			Shutdown,
		),
	)

	app.Run()
}

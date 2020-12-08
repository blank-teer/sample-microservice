package rest

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"order-manager/pkg/api/v3/delivery/rest/handlers"
	"order-manager/pkg/cfg"

	"github.com/scc/go-common/log"
)

type API struct {
	logger log.Logger
	config cfg.Server

	hMeta     handlers.Meta
	hOrder    handlers.Order
	hEmission handlers.Emission
}

func New(
	c cfg.Server,
	l log.Logger,
	hStatus handlers.Meta,
	hOrder handlers.Order,
	hEmission handlers.Emission,
) API {
	return API{
		config:    c,
		logger:    l,
		hMeta:     hStatus,
		hOrder:    hOrder,
		hEmission: hEmission,
	}
}

func (a API) Run() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	a.registerRoutes(a.config.HTTP, r)

	if err := http.ListenAndServe(a.config.HTTP.ListenAddr, r); err != nil {
		a.logger.Errorf("rest.Run: %s", err)
		return err
	}
	return nil
}

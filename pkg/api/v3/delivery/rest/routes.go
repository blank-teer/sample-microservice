package rest

import (
	"github.com/go-chi/chi"
	"github.com/scc/go-common/auth"

	"order-manager/pkg/cfg"
)

// @title Order-Manager REST API
// @version 1.0
// @description Endpoints documentation.

// @contact.name Some Cool Company
// @license.name All rights reserve.

// @BasePath /v3/ord

// @securityDefinitions.basic BasicAuth
func (a API) registerRoutes(cfg cfg.HTTP, r chi.Router) {
	r.Route("/v3/ord", func(r chi.Router) {
		r.Get("/status", a.hMeta.Status)

		r.Route("/user", func(r chi.Router) {
			r.Use(auth.Use(auth.CheckAuthorization(cfg.AuthRequired)))
			r.Get("/orders", a.hOrder.GetByUserID)
		})

		r.Route("/adm", func(r chi.Router) {
			r.Use(auth.Use(auth.CheckSign(cfg.AuthRequired)))
			r.Route("/orders", func(r chi.Router) {
				r.Get("/", a.hOrder.GetAll)
				r.Get("/{orderID}", a.hOrder.Get)
			})
			r.Route("/emission", func(r chi.Router) {
				r.Post("/", a.hEmission.Create)
			})
		})
	})
}

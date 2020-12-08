package handlers

import (
	"net/http"

	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers/render"
)

type Meta struct {
	logger log.Logger
}

func NewMeta(l log.Logger) Meta {
	return Meta{
		logger: l,
	}
}

// @Summary Get service status
// @Tags Meta
// @Produce json
// @Success 200 {object} render.BodySuccess
// @Router /status [get]
func (h Meta) Status(w http.ResponseWriter, _ *http.Request) {
	render.OK(w, nil)
}

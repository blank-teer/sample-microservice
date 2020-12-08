package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/scc/go-common/log"

	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers"
	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers/render"
	"order-manager/pkg/api/v3/models"
	"order-manager/pkg/api/v3/services"
)

const (
	paramEmissionID = "emissionID"
)

type Emission struct {
	logger  log.Logger
	service services.Emission
}

func NewEmission(l log.Logger, s services.Emission) Emission {
	return Emission{
		logger:  l,
		service: s,
	}
}

type ResponseDataCreate struct {
	EmissionID uuid.UUID `json:"emission_id" swaggertype:"string" format:"uuid"`
}

// @Summary Create
// @Tags Emission
// @Accept json
// @Produce json
// @Param Authorization header string true "Token retrieved from admin-api by /auth endpoint."
// @Param emission body models.Emission true "Emission content"
// @Success 200 {object} render.BodySuccess{data=ResponseDataCreate}
// @Failure 400 {object} render.BodyError
// @Failure 500 {object} render.BodyError
// @Router /adm/emission [post]
func (h Emission) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var e models.Emission
	if err := helpers.ParsePayload(r.Body, &e); err != nil {
		render.Error(w, err)
		return
	}

	emissionID, err := h.service.Create(ctx, e)
	if err != nil {
		render.Error(w, err)
		return
	}

	rData := ResponseDataCreate{EmissionID: emissionID}
	render.Created(w, rData)
}

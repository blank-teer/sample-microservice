package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scc/go-common/log"
	"github.com/scc/go-common/types/currency"

	"order-manager/pkg/api/v3/delivery/rest/handlers/helpers/render"
	"order-manager/pkg/api/v3/services"
)

const (
	pathParamOrderID = "orderID"
)

const (
	queryParamOrderType = "type"
	queryParamCoverage  = "coverage"
)

const (
	headerXAuthUser = "X-Auth-User"
)

type Order struct {
	logger  log.Logger
	service services.Order
}

func NewOrders(l log.Logger, s services.Order) Order {
	return Order{
		logger:  l,
		service: s,
	}
}

type RequestCreateAndDone struct {
	WalletAddress string        `json:"wallet_address"`
	CoinsAmount   currency.Coin `json:"amount"`
}

// @Summary Get all
// @Tags Orders
// @Produce json
// @Param Authorization header string true "Token retrieved from admin-api by /auth endpoint."
// @Param type query string false "Filter by order type" Enums(deposit,withdraw)
// @Param coverage query string false "Filter by coverage" Enums(not,fully,partially)
// @Success 200 {object} render.BodySuccess{data=[]models.Order}
// @Failure 400 {object} render.BodyError
// @Failure 500 {object} render.BodyError
// @Router /adm/orders [get]
func (h Order) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	by := services.FilterBy{
		OrderType:    r.URL.Query().Get(queryParamOrderType),
		CoverageRate: r.URL.Query().Get(queryParamCoverage),
	}

	oo, err := h.service.GetList(ctx, by)
	if err != nil {
		render.Error(w, err)
		return
	}

	render.OK(w, oo)
}

// @Summary Get by id
// @Tags Orders
// @Produce json
// @Param Authorization header string true "Token retrieved from admin-api by /auth endpoint."
// @Param orderID path string true "Order ID (in UUID format)"
// @Success 200 {object} render.BodySuccess{data=models.Order}
// @Failure 400 {object} render.BodyError
// @Failure 404 {object} render.BodyError
// @Failure 500 {object} render.BodyError
// @Router /adm/orders/{orderID} [get]
func (h Order) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderID := chi.URLParam(r, pathParamOrderID)

	o, err := h.service.Get(ctx, orderID)
	if err != nil {
		render.Error(w, err)
		return
	}

	render.OK(w, o)
}

// @Summary Get by user id
// @Tags Orders
// @Produce json
// @Param Authorization header string true "Token retrieved from user-api by /login endpoint."
// @Success 200 {object} render.BodySuccess{data=[]models.Order}
// @Failure 400 {object} render.BodyError
// @Failure 500 {object} render.BodyError
// @Router /user/orders [get]
func (h Order) GetByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	by := services.FilterBy{
		UserID: r.Header.Get(headerXAuthUser),
	}

	o, err := h.service.GetList(ctx, by)
	if err != nil {
		render.Error(w, err)
		return
	}

	render.OK(w, o)
}

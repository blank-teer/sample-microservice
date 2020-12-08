package models

import (
	"github.com/google/uuid"
	"github.com/scc/go-common/types/currency"
)

type Coverage struct {
	EmissionID uuid.UUID     `validate:"required" db:"emission_id" json:"emission_id" swaggertype:"string" format:"uuid"`
	OrderID    uuid.UUID     `validate:"required" db:"order_id" json:"order_id" swaggertype:"string" format:"uuid"`
	Amount     currency.Fiat `validate:"required" db:"amount" json:"amount" swaggertype:"integer" example:"60"`
}

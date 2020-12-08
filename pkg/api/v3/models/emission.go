package models

import (
	"encoding/json"
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/scc/go-common/types/currency"

	"order-manager/pkg/api/v3/errs"
)

type Emission struct {
	// Num is a serial number of the emission in the database
	Num int64 `db:"num" json:"num" example:"3"`

	ID uuid.UUID `db:"id" json:"id" swaggertype:"string" format:"uuid"`

	// Coverage is an array of pairs 'order - coverage amount' that says how much the order is covered by this emission
	Coverage []Coverage `validate:"required" db:"-" json:"coverage"`

	Type       EmissionType `validate:"required" db:"type" json:"type" enums:"deposit,withdraw"`
	GoldUnit   string       `validate:"required" db:"gold_unit" json:"gold_unit" example:"grams"`
	GoldAmount int64        `validate:"required" db:"gold_amount" json:"gold_amount" example:"30"`

	// BaseCurrencyAmount is summary fiat amount covered by this emission
	BaseCurrencyAmount currency.Fiat `validate:"required" db:"base_currency_amount" json:"base_currency_amount" swaggertype:"integer" example:"100"`

	BaseCurrency   string  `validate:"required" db:"base_currency" json:"base_currency" example:"EUR"`
	RateBaseToUnit float64 `validate:"required" db:"rate_base_to_unit" json:"rate_base_to_unit" example:"12.5"`

	// Reference is any additional (unnecessary) information about emission
	Reference string `db:"reference" json:"reference" example:"any text"`

	// Details is detailed information about emission in raw json format
	Details json.RawMessage `validate:"required" db:"details" json:"details" swaggertype:"string" format:"json" example:"{'field1':'value1', 'field2':'value2'}"`

	CreatedAt int64 `db:"created_at" json:"created_at" example:"1598884059"`
	UpdatedAt int64 `db:"updated_at" json:"updated_at" example:"1598884069"`
}

func (e Emission) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.ID, is.UUIDv4),
		v.Field(&e.Coverage, v.Required),
		v.Field(&e.Type, v.Required),
		v.Field(&e.GoldUnit, v.Required),
		v.Field(&e.GoldAmount, v.Required),
		v.Field(&e.BaseCurrency, v.Required),
		v.Field(&e.BaseCurrencyAmount, v.Required),
		v.Field(&e.RateBaseToUnit, v.Required),
		v.Field(&e.Details, v.Required),
	)
}

type EmissionSummary struct {
	Unit string        `json:"unit" db:"unit"`
	Gold int64         `json:"gold" db:"gold_amount"`
	Fiat currency.Fiat `json:"fiat" db:"base_currency_amount"`
}

type EmissionType string

const (
	EmissionTypeDeposit  EmissionType = "deposit"
	EmissionTypeWithdraw EmissionType = "withdraw"
)

var emissionTypes = map[EmissionType]struct{}{
	EmissionTypeDeposit:  {},
	EmissionTypeWithdraw: {},
}

func (t EmissionType) Validate() error {
	if _, ok := emissionTypes[t]; !ok {
		return fmt.Errorf("%w: emission type: %s", errs.ErrNotValid, t)
	}
	return nil
}

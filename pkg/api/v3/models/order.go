package models

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/scc/go-common/types/currency"
	userapi "github.com/scc/go-common/types/user_api"

	"order-manager/pkg/api/v3/errs"
)

type Order struct {
	ID     uuid.UUID   `db:"id" json:"id" swaggertype:"string" format:"uuid"`
	Type   OrderType   `validate:"required" db:"type" json:"type" enums:"deposit,withdraw"`
	Status OrderStatus `db:"status" json:"status" enums:"new,paid,issued,failed"`

	// Provider is a name of the service initiated the request to create order
	Provider string `validate:"required" db:"provider" json:"provider" example:"acq"`

	// UserID is salted user id which encrypts real user id persisted in database and used by backend as integer value
	UserID userapi.UserId `validate:"required" db:"user_id" json:"user_id" swaggertype:"string" example:"KJO3XGN63QAAG"`

	// WalletAddress is a user's cosmos wallet address from which the order was paid
	WalletAddress string `validate:"required" db:"wallet_address" json:"wallet_address" example:"cosmos1nrxyjtv3vl2p4hqfx6mdpn7yrakzh9dwxnyd8a"`

	// BaseCurrency is a currency we convert every order currency to
	BaseCurrency string `validate:"required" db:"base_currency" json:"base_currency" example:"EUR"`

	BaseCurrencyAmount currency.Fiat  `validate:"required" db:"base_currency_amount" json:"base_currency_amount" swaggertype:"integer" example:"100"`
	BaseCurrencyPrice  currency.Price `validate:"required" db:"base_currency_price" json:"base_currency_price" swaggertype:"number" format:"double" example:"100.5"`

	// PayCurrency is a currency of the user's order
	PayCurrency       string         `validate:"required" db:"pay_currency" json:"pay_currency" example:"UAH"`
	PayCurrencyAmount currency.Fiat  `validate:"required" db:"pay_currency_amount" json:"pay_currency_amount" swaggertype:"integer" example:"200"`
	PayCurrencyPrice  currency.Price `validate:"required" db:"pay_currency_price" json:"pay_currency_price" swaggertype:"number" format:"double" example:"200.5"`
	RatePayToBase     currency.Price `validate:"required" db:"rate_pay_to_base" json:"rate_pay_to_base" swaggertype:"number" format:"double" example:"1.5"`
	Coins             currency.Coin  `validate:"required" db:"coins" json:"coins" swaggertype:"integer" example:"200"`
	Fee               currency.Fiat  `validate:"required" db:"fee" json:"fee" swaggertype:"integer" example:"50"`
	CreatedAt         int64          `db:"created_at" json:"created_at" example:"1598884059"`
	UpdatedAt         int64          `db:"updated_at" json:"updated_at" example:"1598884069"`
}

func (o Order) Validate() error {
	return v.ValidateStruct(&o,
		v.Field(&o.ID, is.UUIDv4),
		v.Field(&o.Type, v.Required),
		v.Field(&o.Provider, v.Required),
		v.Field(&o.UserID, v.Required),
		v.Field(&o.WalletAddress, v.Required),
		v.Field(&o.BaseCurrency, v.Required),
		v.Field(&o.BaseCurrencyAmount, v.Required),
		v.Field(&o.BaseCurrencyPrice, v.Required),
		v.Field(&o.PayCurrency, v.Required),
		v.Field(&o.PayCurrencyAmount, v.Required),
		v.Field(&o.PayCurrencyPrice, v.Required),
		v.Field(&o.RatePayToBase, v.Required),
		v.Field(&o.Coins, v.Required),
		v.Field(&o.Fee, v.Required),
	)
}

type OrderSummary struct {
	Fiat  currency.Fiat `json:"fiat" db:"base_currency_amount"`
	Coins currency.Coin `json:"coins" db:"coins"`
}

type OrderType string

const (
	OrderTypeDeposit  OrderType = "deposit"
	OrderTypeWithdraw OrderType = "withdraw"
)

var orderTypes = map[OrderType]struct{}{
	OrderTypeDeposit:  {},
	OrderTypeWithdraw: {},
}

func (t OrderType) Validate() error {
	if _, ok := orderTypes[t]; !ok {
		return fmt.Errorf("%w: order type: %s", errs.ErrNotValid, t)
	}
	return nil
}

type OrderStatus string

const (
	OrderStatusNew    OrderStatus = "new"    // after order creation
	OrderStatusPaid   OrderStatus = "paid"   // after user payment for order
	OrderStatusIssued OrderStatus = "issued" // after success payment and success issue
	OrderStatusFailed OrderStatus = "failed" // after any failure
)

var orderStatuses = map[OrderStatus]struct{}{
	OrderStatusNew:    {},
	OrderStatusPaid:   {},
	OrderStatusIssued: {},
	OrderStatusFailed: {},
}

func (s OrderStatus) Validate() error {
	if _, ok := orderStatuses[s]; !ok {
		return fmt.Errorf("%w: order status: %s", errs.ErrNotValid, s)
	}
	return nil
}

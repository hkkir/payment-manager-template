package model

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
Abstract each storage item as "record" independent from the
context of each storage service provider.

E.g.
In Postgres, we may call each record as: row
In MongoDB, we may call each record as: document etc..

In general, we call each storage item as "record".
*/
type PaymentRecord struct {
	id                    int
	provider_id           int
	market_country_id     int
	currency_id           int
	intent_id             string
	provider_resource_id  string
	payment_method_id     int
	custom_transaction_id string
	client_id             string
	client_resource_id    string
	state                 int
	amount                decimal.Decimal
	payer_id              string
	payee_id              string
	custom_props          string
	create_time           time.Time
	update_time           time.Time
}

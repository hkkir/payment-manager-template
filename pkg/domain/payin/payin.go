package payin

/*
DDD suggests that every implementation decisions should be made around the domain.

So, /pkg/domain/{domain_name}.go is the place where we define the domain.

Meaning that here we define:

- Required domain models
- Domain errors
- Domain usecases
- Required services (e.g. storage, external integrations, etc.)
- Validators

"Organize by responsibility" is a rule of thumb in Go world: https://rakyll.org/style-packages/

However, gradually a domain definiton might grow. If this happens, I advise to flatten all these
definitions to their own folder/file organization with the exception

In the beginning we do exceptions for "error.go" and "validation.go". They deserve their own files
because when we look into a domain folder (e.g. /pkg/domain/payin), we can clearly see that they
are part of the usecases of the domain without cluttering the CORE domain definition (model, usecase,
storage, external services etc.)

If a domain grows too much from number of definitions perspective, that might be a sign that we may
need to redefine our domains or need to split the domain into multiple subdomains.
*/

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	Id                  int
	Provider            string // ConvPay_proto.PaymentProvider
	MarketCountry       string // marketplace_type.Type
	Currency            string // currency_type.Type
	ProviderResurceId   string
	PaymentMethod       string // ConvPay_proto.PaymentMethodType
	CustomTransactionId string
	ClientId            string
	ClientResourceId    string
	State               string // ConvPay_proto.PaymentState
	Amount              decimal.Decimal
	PayerId             string
	PayeeId             string
	CustomProps         string // This shall be unmarshalled into its own models depending on business logic
	CreateTime          time.Time
	UpdateTime          time.Time
}

type PayinUsecase interface {
	CreatePayment(ctx context.Context, payment *Payment) (createdPayment *Payment, err error)
	CapturePayment(ctx context.Context, id string) error
	GetPaymentById(ctx context.Context, id string) (*Payment, error)
	CancelPayment(ctx context.Context, id string) error
}

type ProviderService interface {
	Create(ctx context.Context, payment *Payment) (err error)
	Get(ctx context.Context, paymentId string) (payment *Payment, err error)
	Refund(ctx context.Context, paymentId string) (refundId string, err error)
	RefundPartial(ctx context.Context, paymentId string, amount decimal.Decimal) (refundId string, err error)
	Capture(ctx context.Context, paymentId string) (err error)
	Cancel(ctx context.Context, paymentId string) (err error)
}

type StorageService interface {
	CreatePayment(ctx context.Context, payment *Payment) (createdPayment *Payment, err error)
	UpdatePayment(ctx context.Context, payment *Payment) (updatedPayment *Payment, err error)
	GetPayment(ctx context.Context, id string) (payment *Payment, err error)
	QueryPayments(ctx context.Context, query string, args ...interface{}) (payments []*Payment, err error)
}

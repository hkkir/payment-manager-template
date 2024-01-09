package inai

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
)

type inaiService struct{}

func New() payin.ProviderService {
	fmt.Println("New Inai Service")
	return &inaiService{}
}

func (s *inaiService) Create(ctx context.Context, payment *payin.Payment) (err error) {
	fmt.Println("Inai Service Create Order")
	return
}

func (s *inaiService) Get(ctx context.Context, paymentId string) (payment *payin.Payment, err error) {
	fmt.Println("Inai Service Get Order")
	return
}

func (s *inaiService) Refund(ctx context.Context, paymentId string) (refundId string, err error) {
	fmt.Println("Inai Service Refund Order")
	return
}

func (s *inaiService) RefundPartial(ctx context.Context, paymentId string, amount decimal.Decimal) (refundId string, err error) {
	fmt.Println("Inai Service Refund Partial Order")
	return
}

func (s *inaiService) Capture(ctx context.Context, paymentId string) (err error) {
	fmt.Println("Inai Service Capture Order")
	return
}

func (s *inaiService) Cancel(ctx context.Context, paymentId string) (err error) {
	fmt.Println("Inai Service Cancel Order")
	return
}

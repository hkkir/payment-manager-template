package payin

import (
	"context"
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
)

type PayInPaymentHandler struct {
	payinUseCase payin.PayinUsecase
}

func New(payinUseCase payin.PayinUsecase) *PayInPaymentHandler {
	fmt.Println("New PayInPaymentHandler")
	return &PayInPaymentHandler{
		payinUseCase: payinUseCase,
	}
}

func (handler *PayInPaymentHandler) CreatePaymentAsync(ctx context.Context) {
	fmt.Println("PayInPaymentHandler Create Payment Async")
	payment := &payin.Payment{}
	handler.payinUseCase.CreatePayment(ctx, payment)
}

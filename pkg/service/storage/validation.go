package storage

import (
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/service/storage/model"
)

type validator struct{}

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) Validate(paymentRecord *model.PaymentRecord) error {
	fmt.Println("Validating PaymentRecord")
	return nil
}

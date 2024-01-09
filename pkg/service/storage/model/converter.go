package model

import (
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
)

type payInRecordConverter struct{}

func NewPayInRecordConverter() *payInRecordConverter {
	return &payInRecordConverter{}
}

func (c *payInRecordConverter) ToStorageModel(payment *payin.Payment) (*PaymentRecord, error) {
	fmt.Println("Converting to storage model")
	return nil, nil
}

func (c *payInRecordConverter) ToDomainModel(record *PaymentRecord) (*payin.Payment, error) {
	fmt.Println("Converting to domain model")
	return nil, nil
}

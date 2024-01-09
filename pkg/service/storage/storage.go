package storage

import (
	"context"
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
	"github.com/hkkir/payment-manager-template/pkg/service/storage/model"
)

type storage struct{}

func New() payin.StorageService {
	// integrate postgresql implementation here
	fmt.Println("New StorageService")
	return &storage{}
}

func (s *storage) CreatePayment(ctx context.Context, payment *payin.Payment) (createdPayment *payin.Payment, err error) {
	// Convert business model to storage model
	storageRecordConverter := model.NewPayInRecordConverter()
	paymentRecord, _ := storageRecordConverter.ToStorageModel(payment)

	// Validate the payment record
	validator := NewValidator()
	validator.Validate(paymentRecord)

	payment, _ = storageRecordConverter.ToDomainModel(paymentRecord)

	fmt.Println("Storage Return Create Payment")
	return payment, err
}

func (s *storage) UpdatePayment(ctx context.Context, payment *payin.Payment) (createdPayment *payin.Payment, err error) {
	fmt.Println("Storage Return Update Payment")
	return
}

func (s *storage) GetPayment(ctx context.Context, id string) (payment *payin.Payment, err error) {
	fmt.Println("Storage Return Get Payment")
	return
}

func (s *storage) QueryPayments(ctx context.Context, query string, args ...interface{}) (payments []*payin.Payment, err error) {
	fmt.Println("Storage Return Query Payments")
	return
}

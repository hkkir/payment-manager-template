package payin

import (
	"context"
	"fmt"
)

type payInUsecase struct {
	storageService  StorageService
	providerService ProviderService
	validator       *validator
}

func NewUsecase(
	storageService StorageService,
	providerService ProviderService,
	validator *validator,
) PayinUsecase {
	return &payInUsecase{
		storageService:  storageService,
		providerService: providerService,
		validator:       validator,
	}
}

func (u *payInUsecase) CreatePayment(ctx context.Context, payment *Payment) (createdPayment *Payment, err error) {
	fmt.Println("Usecase Create Payment Start")

	// Validate the payment
	u.validator.Validate(payment)

	// Call the provider service to create the payment
	u.providerService.Create(ctx, payment)

	// Call the storage service to create the payment
	createdPayment, _ = u.storageService.CreatePayment(ctx, payment)

	fmt.Println("Usecase Create Payment End")

	return createdPayment, err
}

func (u *payInUsecase) CapturePayment(ctx context.Context, id string) (err error) {
	err = ErrImplementMe
	return
}

func (u *payInUsecase) GetPaymentById(ctx context.Context, id string) (payment *Payment, err error) {
	err = ErrImplementMe
	return
}

func (u *payInUsecase) CancelPayment(ctx context.Context, id string) (err error) {
	err = ErrImplementMe
	return
}

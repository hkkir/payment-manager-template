package app

/*
We define a default application layer.

This layer is responsible for initializing all the handlers and injecting the
required dependencies for each handler.
*/

import (
	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
	payin_handler "github.com/hkkir/payment-manager-template/pkg/handler/payin"
	"github.com/hkkir/payment-manager-template/pkg/service"
)

type paymentManagerApp struct {
	*payin_handler.PayInPaymentHandler
}

// Do not stutter when naming the method.
// E.g. NewPaymentManagerApp() is not good.
func New() *paymentManagerApp {
	app := newPaymentManagerApp()
	return app
}

func newPaymentManagerApp() *paymentManagerApp {
	// For demo purpose: this info would come from gRPC request
	provider := "inai"

	// Initialize the service layer
	service := *service.New(provider)

	// Get the required services from the service layer
	storageService := service.GetStorageService()
	providerService := service.GetProviderService()

	payinValidator := payin.NewValidator()

	// Each usecase initializes with the required services to perform the business logic
	payInUsecase := payin.NewUsecase(storageService, providerService, payinValidator)

	// Each handler initializes with its own usecase
	payinPaymentHandler := payin_handler.New(payInUsecase)

	return &paymentManagerApp{
		PayInPaymentHandler: payinPaymentHandler,
	}
}

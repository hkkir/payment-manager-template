package main

/*
This top-level package is used to bootstrap the application mimicking
top level server instantiating the payment manager app.
*/

import (
	"context"
	"fmt"

	default_app "github.com/hkkir/payment-manager-template/pkg/app"
)

func main() {
	fmt.Println("Main Start")

	paymentManagerApp := default_app.New()

	paymentManagerApp.PayInPaymentHandler.CreatePaymentAsync(context.Background())

	fmt.Println("Main End")
}

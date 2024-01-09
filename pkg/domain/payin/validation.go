package payin

import "fmt"

type validator struct{}

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) Validate(payment *Payment) error {
	fmt.Println("Validating Payment")
	return nil
}

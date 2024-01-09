package provider

import (
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
	"github.com/hkkir/payment-manager-template/pkg/service/provider/inai"
)

func GetProviderService(provider_id string) (payin.ProviderService, error) {
	fmt.Println("Factory Get Provider Service")
	switch provider_id {
	case "inai":
		fmt.Println("Factory Get Inai Service")
		return inai.New(), nil
	default:
		return nil, ErrProviderNotFound
	}
}

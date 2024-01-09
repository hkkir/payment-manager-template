package service

/*
This package is responsible for aggregating all services in a single struct 
which would be initialized in the application layer.

Aggregation of all services manifests that "this application requires all 
these services to function."
*/

import (
	"fmt"

	"github.com/hkkir/payment-manager-template/pkg/domain/payin"
	"github.com/hkkir/payment-manager-template/pkg/service/provider"
	"github.com/hkkir/payment-manager-template/pkg/service/storage"
)

// Aggregate all services as a single struct
type service struct {
	StorageService  payin.StorageService
	ProviderService payin.ProviderService
}

func New(provider_id string) *service {
	fmt.Println("New Service")
	service := newService(provider_id)
	return service
}

func newService(provider_id string) *service {
	storageService := storage.New()
	providerService, _ := provider.GetProviderService(provider_id)
	return &service{
		StorageService:  storageService,
		ProviderService: providerService,
	}
}

func (s *service) GetStorageService() payin.StorageService {
	fmt.Println("Get Storage Service")
	return s.StorageService
}

func (s *service) GetProviderService() payin.ProviderService {
	fmt.Println("Get Provider Service")
	return s.ProviderService
}

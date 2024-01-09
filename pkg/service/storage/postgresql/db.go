package postgresql

import (
	"context"

	"github.com/hkkir/payment-manager-template/pkg/service/storage/model"
)

type PostgresPayInStorageService struct{}

func (p PostgresPayInStorageService) Create(ctx context.Context, record model.PaymentRecord) (err error) {
	return err
}

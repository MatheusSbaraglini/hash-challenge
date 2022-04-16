package product

import (
	"context"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

type ServiceMock struct {
	AddProductsFn func(ctx context.Context, products *domain.ProductInput) (*domain.Checkout, error)
}

func (sm *ServiceMock) AddProducts(ctx context.Context, products *domain.ProductInput) (*domain.Checkout, error) {
	return sm.AddProductsFn(ctx, products)
}

type StorageMock struct {
	FindOneFn      func(ctx context.Context, ID int) (*domain.Product, error)
	FindAllGiftsFn func(ctx context.Context) ([]*domain.Product, error)
}

func (sm *StorageMock) FindOne(ctx context.Context, ID int) (*domain.Product, error) {
	return sm.FindOneFn(ctx, ID)
}

func (sm *StorageMock) FindAllGifts(ctx context.Context) ([]*domain.Product, error) {
	return sm.FindAllGiftsFn(ctx)
}

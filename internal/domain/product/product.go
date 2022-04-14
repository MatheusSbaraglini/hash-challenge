package product

import (
	"context"
	"log"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

type service struct {
	log *log.Logger
}

func NewCheckoutService(log *log.Logger) domain.CheckoutService {
	return &service{
		log: log,
	}
}

func (s *service) AddProducts(ctx context.Context, products *domain.ProductInput) (*domain.Checkout, error) {
	checkout := &domain.Checkout{
		Products: make([]*domain.ProductCheckout, 0),
	}

	for _, inputProduct := range products.Products {
		checkoutProduct := &domain.ProductCheckout{
			ID: inputProduct.ID,
			Quantity: inputProduct.Quantity,
		}
		// get product data by id

		// check if it is gift

		// calc discount

		checkout.Products = append(checkout.Products, checkoutProduct)
	}

	// check if it is black friday and add a gift

	checkout.CalculateAmounts()

	return checkout, nil
}

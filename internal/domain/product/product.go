package product

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

type service struct {
	discountClient   domain.DiscountClient
	productStorage   domain.ProductStorage
	funcNow          func() time.Time
	blackFridayStart time.Time
	blackfridayEnd   time.Time
	log              *log.Logger
}

func NewCheckoutService(discountClient domain.DiscountClient, productStorage domain.ProductStorage, funcNow func() time.Time, blackFridayStart time.Time, blackfridayEnd time.Time, log *log.Logger) domain.CheckoutService {
	return &service{
		discountClient:   discountClient,
		productStorage:   productStorage,
		funcNow:          funcNow,
		blackFridayStart: blackFridayStart,
		blackfridayEnd:   blackfridayEnd,
		log:              log,
	}
}

func (s *service) AddProducts(ctx context.Context, products *domain.ProductInput) (*domain.Checkout, error) {
	checkout := &domain.Checkout{
		Products: make([]*domain.ProductCheckout, 0),
	}

	for _, inputProduct := range products.Products {
		// do not calc products with invalid quantity
		if inputProduct.Quantity <= 0 {
			continue
		}

		product, err := s.productStorage.FindOne(ctx, inputProduct.ID)
		if err != nil {
			return nil, err
		}

		if product.IsGift {
			continue
		}

		checkoutProduct := &domain.ProductCheckout{
			ID:          inputProduct.ID,
			Quantity:    inputProduct.Quantity,
			UnitAmount:  product.Amount,
			TotalAmount: product.Amount * inputProduct.Quantity,
			IsGift:      false,
		}

		discountPercentage, err := s.discountClient.GetDiscountProductPercentage(ctx, inputProduct.ID)
		if err != nil {
			s.log.Printf("discount client could not return the product percentage: %v", err)
		}

		checkoutProduct.Discount = int((float32(checkoutProduct.TotalAmount) * (discountPercentage / 100)))

		checkout.Products = append(checkout.Products, checkoutProduct)
	}

	if s.isBlackFriday() {
		gift, err := s.getProductGift(ctx)
		if err != nil {
			s.log.Printf("could not retrieve a gift: %v", err)
		}

		if gift != nil {
			checkout.Products = append(checkout.Products, gift)
		}
	}

	checkout.CalculateAmounts()

	return checkout, nil
}

func (s *service) isBlackFriday() bool {
	now := s.funcNow()

	if now.After(s.blackFridayStart) && now.Before(s.blackfridayEnd) {
		return true
	}

	return false
}

func (s *service) getProductGift(ctx context.Context) (*domain.ProductCheckout, error) {
	gifts, err := s.productStorage.FindAllGifts(ctx)
	if err != nil {
		return nil, err
	}

	if len(gifts) > 1 {
		return nil, fmt.Errorf("could not determine one gift, there is many of them")
	}

	return &domain.ProductCheckout{
		ID:          gifts[0].ID,
		Quantity:    1,
		UnitAmount:  0,
		TotalAmount: 0,
		Discount:    0,
		IsGift:      true,
	}, nil
}

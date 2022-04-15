package domain_test

import (
	"testing"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCheckout_CalculateAmounts(t *testing.T) {
	checkout := &domain.Checkout{
		Products: []*domain.ProductCheckout{
			{
				ID:          1,
				Quantity:    2,
				UnitAmount:  10000,
				TotalAmount: 20000,
				Discount:    500,
				IsGift:      false,
			},
			{
				ID:          3,
				Quantity:    1,
				UnitAmount:  0,
				TotalAmount: 0,
				Discount:    0,
				IsGift:      true,
			},
		},
	}

	assert.Zero(t, checkout.TotalAmount)
	assert.Zero(t, checkout.TotalAmountWithDiscount)
	assert.Zero(t, checkout.TotalDiscount)

	checkout.CalculateAmounts()

	assert.Equal(t, 20000, checkout.TotalAmount)
	assert.Equal(t, 19500, checkout.TotalAmountWithDiscount)
	assert.Equal(t, 500, checkout.TotalDiscount)
}

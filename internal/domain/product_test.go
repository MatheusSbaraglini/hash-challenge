package domain_test

import (
	"testing"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProductInput_Validate(t *testing.T) {
	t.Run("should validate successfully", func(t *testing.T) {
		input := &domain.ProductInput{
			Products: []struct {
				ID       int `json:"id"`
				Quantity int `json:"quantity"`
			}{
				{
					ID:       1,
					Quantity: 1,
				},
				{
					ID:       2,
					Quantity: 50,
				},
			},
		}

		assert.NoError(t, input.Validate())
	})

	t.Run("should return error when add a product with quantity 0", func(t *testing.T) {
		input := &domain.ProductInput{
			Products: []struct {
				ID       int `json:"id"`
				Quantity int `json:"quantity"`
			}{
				{
					ID:       1,
					Quantity: 1,
				},
				{
					ID:       2,
					Quantity: 0,
				},
			},
		}

		assert.EqualError(t, input.Validate(), "quantity should be greater than 0 for product ID: 2")
	})
}

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

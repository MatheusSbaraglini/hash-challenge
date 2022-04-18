package domain

import "context"

type DiscountClient interface {
	GetDiscountProductPercentage(ctx context.Context, productID int) (float32, error)
}

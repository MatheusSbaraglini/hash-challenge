package discount

import "context"

type ClientMock struct {
	GetDiscountProductPercentageFn func(ctx context.Context, productID int) (float32, error)
}

func (cm *ClientMock) GetDiscountProductPercentage(ctx context.Context, productID int) (float32, error) {
	return cm.GetDiscountProductPercentageFn(ctx, productID)
}

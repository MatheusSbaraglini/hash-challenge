package product_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"github.com/matheussbaraglini/hash-challenge/internal/domain/product"
	"github.com/stretchr/testify/assert"
)

const (
	dateLayout = "02/01/2006 15:04:05"
)

func TestCheckoutService_AddProducts(t *testing.T) {
	testLog := log.New(os.Stderr, "", log.Ldate)

	t.Run("should add products successfully", func(t *testing.T) {
		storageMock := &product.StorageMock{
			FindOneFn: func(ctx context.Context, ID int) (*domain.Product, error) {
				products := make(map[int]*domain.Product)

				for i := 1; i <= 3; i++ {
					isGift := i == 3

					products[i] = &domain.Product{
						ID:          i,
						Title:       fmt.Sprintf("title %d", i),
						Description: fmt.Sprintf("desc %d", i),
						Amount:      i * 500,
						IsGift:      isGift,
					}
				}

				return products[ID], nil
			},
			FindAllGiftsFn: func(ctx context.Context) ([]*domain.Product, error) {
				gifts := []*domain.Product{
					{
						ID:          3,
						Title:       "title 3",
						Description: "desc 3",
						Amount:      1500,
						IsGift:      true,
					},
				}

				return gifts, nil
			},
		}

		funcNow := func() time.Time {
			return time.Date(2022, time.April, 15, 9, 30, 0, 0, time.Local)
		}

		blackFridayStart, err := time.ParseInLocation(dateLayout, "15/04/2022 09:00:00", time.Local)
		assert.NoError(t, err)

		blackFridayEnd, err := time.ParseInLocation(dateLayout, "15/04/2022 10:00:00", time.Local)
		assert.NoError(t, err)

		service := product.NewCheckoutService(storageMock, funcNow, blackFridayStart, blackFridayEnd, testLog)

		input := &domain.ProductInput{
			Products: []struct {
				ID       int "json:\"id\""
				Quantity int "json:\"quantity\""
			}{
				{
					ID:       1,
					Quantity: 5,
				},
				{
					ID:       2,
					Quantity: 3,
				},
			},
		}
		checkout, err := service.AddProducts(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, checkout)

		assert.Equal(t, 5500, checkout.TotalAmount)
		assert.Equal(t, 5500, checkout.TotalAmountWithDiscount)
		assert.Equal(t, 0, checkout.TotalDiscount)

		expected := []*domain.ProductCheckout{
			{
				ID:          1,
				Quantity:    5,
				UnitAmount:  500,
				TotalAmount: 2500,
				Discount:    0,
				IsGift:      false,
			},
			{
				ID:          2,
				Quantity:    3,
				UnitAmount:  1000,
				TotalAmount: 3000,
				Discount:    0,
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
		}

		assert.Equal(t, expected, checkout.Products)
	})

	t.Run("should not calc gift product on input", func(t *testing.T) {
		storageMock := &product.StorageMock{
			FindOneFn: func(ctx context.Context, ID int) (*domain.Product, error) {
				products := make(map[int]*domain.Product)

				for i := 1; i <= 3; i++ {
					isGift := i == 3

					products[i] = &domain.Product{
						ID:          i,
						Title:       fmt.Sprintf("title %d", i),
						Description: fmt.Sprintf("desc %d", i),
						Amount:      i * 500,
						IsGift:      isGift,
					}
				}

				return products[ID], nil
			},
			FindAllGiftsFn: func(ctx context.Context) ([]*domain.Product, error) {
				gifts := []*domain.Product{
					{
						ID:          3,
						Title:       "title 3",
						Description: "desc 3",
						Amount:      1500,
						IsGift:      true,
					},
				}

				return gifts, nil
			},
		}

		funcNow := func() time.Time {
			return time.Date(2022, time.April, 15, 9, 30, 0, 0, time.Local)
		}
		blackFridayStart, err := time.ParseInLocation(dateLayout, "15/04/2022 09:00:00", time.Local)
		assert.NoError(t, err)

		blackFridayEnd, err := time.ParseInLocation(dateLayout, "15/04/2022 10:00:00", time.Local)
		assert.NoError(t, err)

		service := product.NewCheckoutService(storageMock, funcNow, blackFridayStart, blackFridayEnd, testLog)

		input := &domain.ProductInput{
			Products: []struct {
				ID       int "json:\"id\""
				Quantity int "json:\"quantity\""
			}{
				{
					ID:       1,
					Quantity: 5,
				},
				{
					ID:       2,
					Quantity: 3,
				},
				{
					ID:       3, // gift
					Quantity: 2,
				},
			},
		}
		checkout, err := service.AddProducts(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, checkout)

		assert.Equal(t, 5500, checkout.TotalAmount)
		assert.Equal(t, 5500, checkout.TotalAmountWithDiscount)
		assert.Equal(t, 0, checkout.TotalDiscount)

		expected := []*domain.ProductCheckout{
			{
				ID:          1,
				Quantity:    5,
				UnitAmount:  500,
				TotalAmount: 2500,
				Discount:    0,
				IsGift:      false,
			},
			{
				ID:          2,
				Quantity:    3,
				UnitAmount:  1000,
				TotalAmount: 3000,
				Discount:    0,
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
		}

		assert.Equal(t, expected, checkout.Products)
	})

	t.Run("should not add a gift out of black friday", func(t *testing.T) {
		storageMock := &product.StorageMock{
			FindOneFn: func(ctx context.Context, ID int) (*domain.Product, error) {
				products := make(map[int]*domain.Product)

				for i := 1; i <= 3; i++ {
					isGift := i == 3

					products[i] = &domain.Product{
						ID:          i,
						Title:       fmt.Sprintf("title %d", i),
						Description: fmt.Sprintf("desc %d", i),
						Amount:      i * 500,
						IsGift:      isGift,
					}
				}

				return products[ID], nil
			},
			FindAllGiftsFn: func(ctx context.Context) ([]*domain.Product, error) {
				gifts := []*domain.Product{
					{
						ID:          3,
						Title:       "title 3",
						Description: "desc 3",
						Amount:      1500,
						IsGift:      true,
					},
				}

				return gifts, nil
			},
		}

		funcNow := func() time.Time {
			return time.Date(2022, time.April, 15, 11, 30, 0, 0, time.Local)
		}
		blackFridayStart, err := time.ParseInLocation(dateLayout, "15/04/2022 09:00:00", time.Local)
		assert.NoError(t, err)

		blackFridayEnd, err := time.ParseInLocation(dateLayout, "15/04/2022 10:00:00", time.Local)
		assert.NoError(t, err)

		service := product.NewCheckoutService(storageMock, funcNow, blackFridayStart, blackFridayEnd, testLog)

		input := &domain.ProductInput{
			Products: []struct {
				ID       int "json:\"id\""
				Quantity int "json:\"quantity\""
			}{
				{
					ID:       1,
					Quantity: 5,
				},
				{
					ID:       2,
					Quantity: 3,
				},
			},
		}
		checkout, err := service.AddProducts(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, checkout)

		assert.Equal(t, 5500, checkout.TotalAmount)
		assert.Equal(t, 5500, checkout.TotalAmountWithDiscount)
		assert.Equal(t, 0, checkout.TotalDiscount)

		expected := []*domain.ProductCheckout{
			{
				ID:          1,
				Quantity:    5,
				UnitAmount:  500,
				TotalAmount: 2500,
				Discount:    0,
				IsGift:      false,
			},
			{
				ID:          2,
				Quantity:    3,
				UnitAmount:  1000,
				TotalAmount: 3000,
				Discount:    0,
				IsGift:      false,
			},
		}

		assert.Equal(t, expected, checkout.Products)
	})

	t.Run("should not add a gift when it's not possible determine one", func(t *testing.T) {
		storageMock := &product.StorageMock{
			FindOneFn: func(ctx context.Context, ID int) (*domain.Product, error) {
				products := make(map[int]*domain.Product)

				for i := 1; i <= 3; i++ {
					isGift := i == 3

					products[i] = &domain.Product{
						ID:          i,
						Title:       fmt.Sprintf("title %d", i),
						Description: fmt.Sprintf("desc %d", i),
						Amount:      i * 500,
						IsGift:      isGift,
					}
				}

				return products[ID], nil
			},
			FindAllGiftsFn: func(ctx context.Context) ([]*domain.Product, error) {
				gifts := []*domain.Product{
					{
						ID:          3,
						Title:       "title 3",
						Description: "desc 3",
						Amount:      1500,
						IsGift:      true,
					},
					{
						ID:          4,
						Title:       "title 4",
						Description: "desc 4",
						Amount:      1500,
						IsGift:      true,
					},
				}

				return gifts, nil
			},
		}

		funcNow := func() time.Time {
			return time.Date(2022, time.April, 15, 9, 30, 0, 0, time.Local)
		}
		blackFridayStart, err := time.ParseInLocation(dateLayout, "15/04/2022 09:00:00", time.Local)
		assert.NoError(t, err)

		blackFridayEnd, err := time.ParseInLocation(dateLayout, "15/04/2022 10:00:00", time.Local)
		assert.NoError(t, err)

		service := product.NewCheckoutService(storageMock, funcNow, blackFridayStart, blackFridayEnd, testLog)

		input := &domain.ProductInput{
			Products: []struct {
				ID       int "json:\"id\""
				Quantity int "json:\"quantity\""
			}{
				{
					ID:       1,
					Quantity: 5,
				},
				{
					ID:       2,
					Quantity: 3,
				},
			},
		}
		checkout, err := service.AddProducts(context.Background(), input)
		assert.NoError(t, err)
		assert.NotNil(t, checkout)

		assert.Equal(t, 5500, checkout.TotalAmount)
		assert.Equal(t, 5500, checkout.TotalAmountWithDiscount)
		assert.Equal(t, 0, checkout.TotalDiscount)

		expected := []*domain.ProductCheckout{
			{
				ID:          1,
				Quantity:    5,
				UnitAmount:  500,
				TotalAmount: 2500,
				Discount:    0,
				IsGift:      false,
			},
			{
				ID:          2,
				Quantity:    3,
				UnitAmount:  1000,
				TotalAmount: 3000,
				Discount:    0,
				IsGift:      false,
			},
		}

		assert.Equal(t, expected, checkout.Products)
	})
}

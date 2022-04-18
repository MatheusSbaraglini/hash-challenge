package client

import (
	"context"

	"github.com/matheussbaraglini/hash-challenge/discount"
	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"google.golang.org/grpc"
)

type discountClient struct {
	baseURL string
}

func NewDiscountClient(baseURL string) domain.DiscountClient {
	return &discountClient{
		baseURL: baseURL,
	}
}

func (dc *discountClient) getDiscountClient() (discount.DiscountClient, error) {
	conn, err := grpc.Dial(dc.baseURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return discount.NewDiscountClient(conn), nil
}

func (dc *discountClient) GetDiscountProductPercentage(ctx context.Context, productID int) (float32, error) {
	request := &discount.GetDiscountRequest{
		ProductID: int32(productID),
	}

	discountClient, err := dc.getDiscountClient()
	if err != nil {
		return 0, err
	}

	response, err := discountClient.GetDiscount(ctx, request)
	if err != nil {
		return 0, err
	}

	return response.Percentage, nil
}

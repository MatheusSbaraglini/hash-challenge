package http_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"github.com/matheussbaraglini/hash-challenge/internal/domain/product"
	internalHTTP "github.com/matheussbaraglini/hash-challenge/internal/infrastructure/server/http"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	t.Run("should validate successfully", func(t *testing.T) {
		serviceMock := &product.ServiceMock{
			AddProductsFn: func(ctx context.Context, products *domain.ProductInput) (*domain.Checkout, error) {
				assert.Len(t, products.Products, 1)

				return &domain.Checkout{
					TotalAmount: 20000,
					TotalAmountWithDiscount: 20000,
					TotalDiscount: 0,
					Products: []*domain.ProductCheckout{
						{
							ID: 1,
							Quantity: 2,
							UnitAmount: 10000,
							TotalAmount: 20000,
							Discount: 0,
							IsGift: false,
						},
						{
							ID: 6,
							Quantity: 1,
							UnitAmount: 0,
							TotalAmount: 0,
							Discount: 0,
							IsGift: true,
						},
					},
				}, nil
			},
		}

		server := httptest.NewServer(internalHTTP.NewHandler(serviceMock))
		defer server.Close()

		URL, _ := url.Parse(server.URL)

		body := strings.NewReader(`
		{
			"products": [
				{
					"id": 1,
					"quantity": 2
				}
			]
		}`)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/checkout", URL), body)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)

		bodyBytes, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, bodyBytes)

		expected := `
		{
			"total_amount": 20000,
			"total_amount_with_discount": 20000,
			"total_discount": 0,
			"products": [
				{
					"id": 1,
					"quantity": 2,
					"unit_amount": 10000,
					"total_amount": 20000,
					"discount": 0,
					"is_gift": false
				},
				{
					"id": 6,
					"quantity": 1,
					"unit_amount": 0,
					"total_amount": 0,
					"discount": 0,
					"is_gift": true
				}
			]
		}`

		assert.JSONEq(t, expected, string(bodyBytes))
	})
}

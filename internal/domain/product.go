package domain

import (
	"fmt"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

type ProductInput struct {
	Products []struct {
		ID       int `json:"id"`
		Quantity int `json:"quantity"`
	} `json:"products"`
}

func (pi *ProductInput) Validate() error {
	for _, prd := range pi.Products {
		if prd.Quantity == 0 {
			return fmt.Errorf("quantity should be greater than 0 for product ID: %d", prd.ID)
		}
	}

	return nil
}

type ProductCheckout struct {
	ID          int  `json:"id"`
	Quantity    int  `json:"quantity"`
	UnitAmount  int  `json:"unit_amount"`
	TotalAmount int  `json:"total_amount"`
	Discount    int  `json:"discount"`
	IsGift      bool `json:"is_gift"`
}

type Checkout struct {
	TotalAmount             int                `json:"total_amount"`
	TotalAmountWithDiscount int                `json:"total_amount_with_discount"`
	TotalDiscount           int                `json:"total_discount"`
	Products                []*ProductCheckout `json:"products"`
}

func (c *Checkout) CalculateAmounts() {
	totalAmount := 0
	totalDiscount := 0

	for _, prdCheckout := range c.Products {
		totalAmount += prdCheckout.TotalAmount
		totalDiscount += prdCheckout.Discount
	}

	c.TotalAmount = totalAmount
	c.TotalAmountWithDiscount = totalAmount - totalDiscount
	c.TotalDiscount = totalDiscount

}

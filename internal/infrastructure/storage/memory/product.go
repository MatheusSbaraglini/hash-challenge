package memory

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

type productStorage struct {
	products []*domain.Product
}

func NewMemoryProductStorage(file *os.File) (domain.ProductStorage, error) {
	productsByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read products file: %v", err)
	}

	var products []*domain.Product
	if err := json.Unmarshal(productsByte, &products); err != nil {
		return nil, fmt.Errorf("failed to read products file: %v", err)
	}

	return NewMemoryProductStorageWithProducts(products), nil
}

func NewMemoryProductStorageWithProducts(products []*domain.Product) domain.ProductStorage {
	return &productStorage{
		products: products,
	}
}

func (ps *productStorage) FindOne(ctx context.Context, ID int) (*domain.Product, error) {
	for _, prd := range ps.products {
		if prd.ID == ID {
			return prd, nil
		}
	}

	return nil, fmt.Errorf(`product Id "%d" not found`, ID)
}

func (ps *productStorage) FindAllGifts(ctx context.Context) ([]*domain.Product, error) {
	allGifts := make([]*domain.Product, 0)

	for _, prd := range ps.products {
		if prd.IsGift {
			allGifts = append(allGifts, prd)
		}
	}

	return allGifts, nil
}

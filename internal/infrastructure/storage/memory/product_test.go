package memory_test

import (
	"context"
	"testing"

	"github.com/matheussbaraglini/hash-challenge/internal/domain"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/storage/memory"
	"github.com/stretchr/testify/assert"
)

var (
	products = []*domain.Product{
		{
			ID:          1,
			Title:       "product 1",
			Description: "description of product 1",
			Amount:      15157,
			IsGift:      false,
		},
		{
			ID:          2,
			Title:       "product 2",
			Description: "description of product 2",
			Amount:      93811,
			IsGift:      false,
		},
		{
			ID:          3,
			Title:       "product 3",
			Description: "description of product 3",
			Amount:      60356,
			IsGift:      true,
		},
		{
			ID:          4,
			Title:       "product 4",
			Description: "description of product 4",
			Amount:      56230,
			IsGift:      true,
		},
	}
)

func TestProductStorage_FindOne(t *testing.T) {
	t.Run("should find a product by ID successfully", func(t *testing.T) {
		memoryStorage := memory.NewMemoryProductStorageWithProducts(products)

		found, err := memoryStorage.FindOne(context.Background(), 2)
		assert.NoError(t, err)

		assert.Equal(t, 2, found.ID)
		assert.Equal(t, "product 2", found.Title)
		assert.Equal(t, "description of product 2", found.Description)
		assert.Equal(t, 93811, found.Amount)
		assert.False(t, found.IsGift)

		found, err = memoryStorage.FindOne(context.Background(), 4)
		assert.NoError(t, err)

		assert.Equal(t, 4, found.ID)
		assert.Equal(t, "product 4", found.Title)
		assert.Equal(t, "description of product 4", found.Description)
		assert.Equal(t, 56230, found.Amount)
		assert.True(t, found.IsGift)
	})

	t.Run("should return not found error given a inexistent id", func(t *testing.T) {
		memoryStorage := memory.NewMemoryProductStorageWithProducts(products)

		found, err := memoryStorage.FindOne(context.Background(), 10)
		assert.Nil(t, found)
		assert.Error(t, err)
		assert.EqualError(t, err, `product Id "10" not found`)
	})
}

func TestProductStorage_FindAllGifts(t *testing.T) {
	t.Run("should retrieve all gifts", func(t *testing.T) {
		memoryStorage := memory.NewMemoryProductStorageWithProducts(products)

		gifts, err := memoryStorage.FindAllGifts(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, gifts)

		assert.Len(t, gifts, 2)

		expected := []*domain.Product{
			{
				ID:          3,
				Title:       "product 3",
				Description: "description of product 3",
				Amount:      60356,
				IsGift:      true,
			},
			{
				ID:          4,
				Title:       "product 4",
				Description: "description of product 4",
				Amount:      56230,
				IsGift:      true,
			},
		}

		assert.Equal(t, expected, gifts)
	})
}
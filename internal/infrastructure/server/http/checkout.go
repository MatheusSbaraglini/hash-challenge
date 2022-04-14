package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

func (h *handler) checkout(c *gin.Context) {
	products := &domain.ProductInput{}
	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// add products

	c.JSON(http.StatusOK, products) // temp
}

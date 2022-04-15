package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheussbaraglini/hash-challenge/internal/domain"
)

type handler struct {
	checkoutService domain.CheckoutService
}

func NewHandler(checkoutService domain.CheckoutService) http.Handler {
	handler := &handler{
		checkoutService: checkoutService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/checkout", handler.checkout)

	return router
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() http.Handler {
	handler := &handler{}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/checkout", handler.checkout)

	return router
}

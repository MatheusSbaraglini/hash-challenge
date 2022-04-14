package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) checkout(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

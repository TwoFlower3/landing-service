package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainPage of service
func MainPage(c *gin.Context) {
	c.String(http.StatusOK, "interview service")
}

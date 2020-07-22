package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/twoflower3/interview-service/pkg/handlers"
)

// API ...
func API(router gin.IRouter) {
	router.POST("/send", handlers.Send)
}

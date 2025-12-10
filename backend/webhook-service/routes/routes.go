package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulkale/webhook-service/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Webhook Service is Running"})
	})

	r.POST("/webhook", controllers.HandleStripeWebhook)

	return r
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulkale/payment-service/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Payment Service is Running"})
	})

	// Fix: Removed the paymentGroup. The route should be at the root path
	// since the API Gateway strips the prefix.
	r.POST("/create-checkout-session", controllers.CreateCheckoutSession)

	return r
}

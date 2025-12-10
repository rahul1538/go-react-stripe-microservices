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

	paymentGroup := r.Group("/payments")
	{
		// We renamed this route to fit the new method
		paymentGroup.POST("/create-checkout-session", controllers.CreateCheckoutSession)
	}

	return r
}

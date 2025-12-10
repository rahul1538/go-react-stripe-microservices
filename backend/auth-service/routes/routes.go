package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulkale/auth-service/controllers" // Import the controllers
)

func SetupRoutes(jwtSecret string) *gin.Engine {
	r := gin.Default()

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Auth Service is running!"})
	})

	// Auth Routes
	// Auth Routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login) // <--- I removed the "//" so it is active now!
	}

	return r
}

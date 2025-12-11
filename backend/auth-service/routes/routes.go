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

	// Auth Routes - NO GROUP NEEDED
	// The API Gateway has stripped the "/auth" prefix, so the Auth Service should
	// listen for the root paths: /register and /login.

	// Fix: Changed from auth.POST("/register", ...) to r.POST("/register", ...)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Removed the redundant 'auth := r.Group("/auth")' block entirely.

	return r
}

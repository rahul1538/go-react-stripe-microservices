package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/gin-contrib/cors" // <--- Import this
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	// 2. Add CORS Middleware (Crucial for React connection)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow React Frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Gateway is Running!"})
	})

	// 3. Define Service URLs from .env
	authURL := os.Getenv("AUTH_SERVICE_URL")
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
	webhookURL := os.Getenv("WEBHOOK_SERVICE_URL")

	// 4. Setup Reverse Proxy Routes
	r.Any("/auth/*proxyPath", proxyRequest(authURL))
	r.Any("/payments/*proxyPath", proxyRequest(paymentURL))
	r.Any("/webhook/*proxyPath", proxyRequest(webhookURL))

	// 5. Start Gateway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API Gateway running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run gateway:", err)
	}
}

// proxyRequest sends the request to the target service
func proxyRequest(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Gateway is Running!"})
	})

	// 3. Define Service URLs from environment variables
	authURL := os.Getenv("AUTH_SERVICE_URL")
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
	webhookURL := os.Getenv("WEBHOOK_SERVICE_URL")

	if authURL == "" || paymentURL == "" || webhookURL == "" {
		log.Fatal("One or more service URLs (AUTH, PAYMENT, WEBHOOK) are not set in environment variables.")
	}

	// 4. Setup Reverse Proxy Routes
	r.Any("/auth/*proxyPath", proxyRequest(authURL, "/auth"))
	r.Any("/payments/*proxyPath", proxyRequest(paymentURL, "/payments"))
	r.Any("/webhook/*proxyPath", proxyRequest(webhookURL, "/webhook"))

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

// proxyRequest sets up the proxy and handles the path stripping logic
func proxyRequest(target string, prefixToStrip string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			log.Printf("Error parsing target URL %s: %v", target, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL configuration"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {

			// --- CRITICAL PATH STRIPPING LOGIC ---
			originalPath := req.URL.Path

			// Strip the prefix (e.g., remove "/auth" to get "/register")
			req.URL.Path = strings.TrimPrefix(originalPath, prefixToStrip)

			// --- STANDARD PROXY FORWARDING ---
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.RawQuery = c.Request.URL.RawQuery
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

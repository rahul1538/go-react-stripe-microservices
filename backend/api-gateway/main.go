package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings" // <--- Import added for string manipulation (TrimPrefix)
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
	// NOTE: You should restrict this to your Vercel URL in production, 
    // but leaving it open for initial testing.
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173", "https://go-react-stripe-microservices.vercel.app"}, 
        // Using "*" for maximum compatibility during debugging
        AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Gateway is Running!"})
	})

	// 3. Define Service URLs from .env
	authURL := os.Getenv("AUTH_SERVICE_URL")
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
	webhookURL := os.Getenv("WEBHOOK_SERVICE_URL")

	// Check if critical URLs are set (This check helps debugging)
    if authURL == "" || paymentURL == "" || webhookURL == "" {
        log.Fatal("One or more service URLs (AUTH, PAYMENT, WEBHOOK) are not set in environment variables.")
    }

	// 4. Setup Reverse Proxy Routes
	// CRITICAL CHANGE: Pass the prefix to strip (e.g., "/auth")
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
// It takes the target URL and the prefix that needs to be stripped from the path.
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
            // 1. Get the original requested path (e.g., /auth/register)
			originalPath := req.URL.Path
            
            // 2. Strip the prefix (e.g., remove "/auth" to get "/register")
            // This is necessary because the Auth Service is only listening at /register, not /auth/register
			req.URL.Path = strings.TrimPrefix(originalPath, prefixToStrip)
            
            // 3. Handle Gin's parameter capture (proxyPath) cleanup
            // The Gin router adds a leading slash, which is why TrimPrefix is better than strings.Replace
            
            // --- STANDARD PROXY FORWARDING ---
            
            // Set headers and host for the target service
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
            
            // Ensure query parameters are preserved (e.g., ?id=123)
            req.URL.RawQuery = c.Request.URL.RawQuery 
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
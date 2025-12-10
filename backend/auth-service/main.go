package main

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Recommended: To load .env files locally

	"github.com/rahulkale/auth-service/config"
	"github.com/rahulkale/auth-service/routes"
	// "github.com/rahulkale/auth-service/utils" // Uncomment this only if you have functions in utils to run
)

func main() {
	// 1. Load Environment Variables (Optional but recommended for local dev)
	// This helps read the .env file if it's not automatically loaded
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env variables")
	}

	// 2. Load JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// For local testing, you might want a fallback, but strictly:
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// 3. Connect to Database (MongoDB or Postgres)
	// Ensure your config package has a ConnectDB function
	err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// 4. Utils Initialization
	// Only uncomment this if you have defined InitUtils() in your utils folder!
	// utils.InitUtils()

	// 5. Setup Routes
	// Ensure routes.SetupRoutes accepts the secret or handle it inside the package
	r := routes.SetupRoutes(jwtSecret)

	// 6. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083" // Default fallback
	}

	log.Printf("Auth-service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

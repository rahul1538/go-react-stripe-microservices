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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env variables")
	}

	// 2. Load JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// 3. Connect to Database (MongoDB or Postgres)
	err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// 4. Utils Initialization
	// Only uncomment this if you have defined InitUtils() in your utils folder!
	// utils.InitUtils()

	// 5. Setup Routes
	r := routes.SetupRoutes(jwtSecret)

	// 6. Start Server
	// FIX: Ensure it runs on the port specified by the environment variable (Render requirement)
	port := os.Getenv("PORT")
	if port == "" {
		// Use a common default for local dev if PORT is not set
		port = "8080"
	}

	// Log the actual port being used
	log.Printf("Auth-service running on port %s", port)

	// Run on the derived port
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

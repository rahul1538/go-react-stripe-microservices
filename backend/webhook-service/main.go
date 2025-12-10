package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rahulkale/webhook-service/config"
	"github.com/rahulkale/webhook-service/routes"
)

func main() {
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2. Connect to Database
	// This is CRITICAL. Without this, we can't update the payment status to "Success"
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 3. Setup Routes
	// This connects the URL "/webhook" to your Controller logic
	r := routes.SetupRoutes()

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Webhook Service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

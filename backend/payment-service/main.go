package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rahulkale/payment-service/config" // Import config
	"github.com/rahulkale/payment-service/routes"
)

func main() {
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2. Connect to Database <--- ADD THIS BLOCK
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 3. Setup Routes
	r := routes.SetupRoutes()

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Payment Service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

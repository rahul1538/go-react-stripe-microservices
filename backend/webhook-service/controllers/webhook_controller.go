package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulkale/webhook-service/config"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleStripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	// 1. Verify the signature
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		fmt.Printf("⚠️  Webhook signature verification failed. %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signature"})
		return
	}

	// 2. Handle "payment_intent.succeeded"
	if event.Type == "payment_intent.succeeded" {
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Printf("Error parsing webhook JSON: %v\n", err)
			c.Status(http.StatusBadRequest)
			return
		}

		fmt.Printf("💰 Payment for %d succeeded! ID: %s\n", paymentIntent.Amount, paymentIntent.ID)

		// 3. Update Database
		collection := config.DB.Database("payment_db").Collection("transactions")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		filter := bson.M{"stripe_id": paymentIntent.ID}
		update := bson.M{"$set": bson.M{"status": "succeeded", "updated_at": time.Now()}}

		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Printf("Error updating database: %v\n", err)
		} else {
			fmt.Println("✅ Database updated successfully.")
		}
	}

	c.Status(http.StatusOK)
}

package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulkale/payment-service/config"
	"github.com/rahulkale/payment-service/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session" // Import Session package
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

// CreateCheckoutSession generates a Stripe URL for the user to pay
func CreateCheckoutSession(c *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Create the Checkout Session
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(req.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Microservices Invoice"),
					},
					UnitAmount: stripe.Int64(req.Amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		// Redirect back to React Dashboard after payment
		SuccessURL: stripe.String("http://localhost:5173/dashboard?status=success"),
		CancelURL:  stripe.String("http://localhost:5173/dashboard?status=canceled"),
	}

	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Save "Pending" transaction to MongoDB
	paymentRecord := models.Payment{
		ID:        primitive.NewObjectID(),
		Amount:    req.Amount,
		Currency:  req.Currency,
		StripeID:  s.ID, // Save Session ID to track it later
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	collection := config.DB.Database("payment_db").Collection("transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.InsertOne(ctx, paymentRecord)

	// 3. Send the URL back to Frontend
	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Amount    int64              `bson:"amount" json:"amount"` // Stored in cents
	Currency  string             `bson:"currency" json:"currency"`
	StripeID  string             `bson:"stripe_id" json:"stripe_id"` // The "pi_..." ID from Stripe
	Status    string             `bson:"status" json:"status"`       // e.g., "pending", "succeeded"
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

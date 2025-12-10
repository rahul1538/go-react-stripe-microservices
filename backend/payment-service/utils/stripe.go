package utils

import (
	"os"

	stripe "github.com/stripe/stripe-go/v74"
)

func InitStripe() {
	key := os.Getenv("STRIPE_KEY")
	if key == "" {
		panic("STRIPE_KEY is required")
	}
	stripe.Key = key
}

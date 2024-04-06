package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	stripeWebhook "github.com/stripe/stripe-go/v76/webhook"
)

// This is your Stripe CLI webhook secret for testing your endpoint locally.
// const endpointSecret = "whsec_x3beJhuHM4C6G4SvVQSMgnv9tIQZ31lH"
const endpointSecret = "whsec_9cc7537a2aadc5c017ca45fcb8e0300a45aeeecbba031cf34cc43e05454bfdfb"

func StripeWebhookHandler(c *gin.Context) {

	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	payload, err := c.GetRawData()
	if err != nil {
		fmt.Printf("⚠️  Error parsing request body: %v.\n", err)
		c.JSON(400, dto.ErrorResponse{Error: "Error parsing request body"})
		return
	}
	event, err := stripeWebhook.ConstructEventWithOptions(payload, c.Request.Header.Get("Stripe-Signature"),
		endpointSecret, stripeWebhook.ConstructEventOptions{
			IgnoreAPIVersionMismatch: true,
		})
	if err != nil {
		fmt.Printf("⚠️  Webhook signature verification failed.\n")
		c.JSON(400, dto.ErrorResponse{Error: "Webhook signature verification failed."})
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "customer.subscription.created":
		err := handleSubscriptionCreated(event)
		if err != nil {
			fmt.Printf("⚠️  Error handling subscription created event: %v.\n", err)
			c.JSON(500, dto.ErrorResponse{Error: "Error handling subscription created event"})
			return
		}
		break
	default:
		fmt.Printf("⚠️  Unhandled event type: %s.\n", event.Type)
	}

	fmt.Printf("✅  Successfully parsed event of type: %s.\n", event.Type)
	c.JSON(200, nil)
}

func handleSubscriptionCreated(event stripe.Event) error {
	// Handle the subscription updated event
	var subscription stripe.Subscription
	err := json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		return err
	}
	userID, exists := subscription.Metadata["user_id"]
	if !exists {
		fmt.Printf("⚠️  Subscription does not have user_id metadata.")
	}

	fmt.Printf("Subscription created for user: %s.\n", userID)
	return nil
}

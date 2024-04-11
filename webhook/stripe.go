package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	stripeWebhook "github.com/stripe/stripe-go/v76/webhook"
	"os"
)

// This is your Stripe CLI webhook secret for testing your endpoint locally.
var endpointSecret string

func init() {
	endpointSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
}

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
	stripeCustomerId := subscription.Customer.ID
	stripeSubscriptionId := subscription.ID
	fmt.Printf("🔔  Customer %s subscribed to plan %s.\n", stripeCustomerId, stripeSubscriptionId)
	return nil
}

package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	stripeWebhook "github.com/stripe/stripe-go/v76/webhook"
)

func init() {
	stripe.Key = "sk_test_51MtUANAZeDyeb7mKt6yl1sljMOxWwg7Evyp3Pz7PMqlkxgekFvhe01Fm8ubPDukZpKVskIQRgnllmSa4mRHmB3HY00hK1AVsRr"
}

func StripeWebhookHandler(c *gin.Context) {

	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	endpointSecret := "whsec_x3beJhuHM4C6G4SvVQSMgnv9tIQZ31lH"

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
	case "checkout.session.completed":
		// Then define and call a function to handle the event checkout.session.completed
	case "customer.subscription.deleted":
		// Then define and call a function to handle the event customer.subscription.deleted
	case "customer.subscription.updated":
		// Then define and call a function to handle the event customer.subscription.updated
		// ... handle other event types
	default:
		fmt.Printf("⚠️  Unhandled event type: %s.\n", event.Type)
	}

	fmt.Printf("✅  Successfully parsed event of type: %s.\n", event.Type)
	c.JSON(200, nil)
}

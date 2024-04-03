package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func ListSubscriptions(c *gin.Context) {

	// TODO: Fetch the subscriptions from stripe
	params := &stripe.SubscriptionListParams{}
	params.Limit = stripe.Int64(5)
	params.Status = stripe.String("all")
}

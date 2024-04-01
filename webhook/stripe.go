package webhook

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

func init() {
	stripe.Key = "sk_test_51MtUANAZeDyeb7mKt6yl1sljMOxWwg7Evyp3Pz7PMqlkxgekFvhe01Fm8ubPDukZpKVskIQRgnllmSa4mRHmB3HY00hK1AVsRr"
}

func StripeWebhookHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

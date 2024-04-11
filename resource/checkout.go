package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/common"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentlink"
)

const price = "price_1P25CWAZeDyeb7mK0DCiCk69"

// GetPaymentLinkHandler creates a payment link
// @Summary Get a payment link
// @Description Get a payment link
// @Tags Payment
// @Accept json
// @Produce json
// @Success 201 {object} dto.GetPaymentLinkRequest
// @Failure 400 {object} dto.ErrorResponse
// @Router /p/link [get]
func GetPaymentLinkHandler(c *gin.Context) {

	var userInfo common.UserInfo
	err := common.GetUserInfoByContext(c, &userInfo)
	if err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String(price),
				Quantity: stripe.Int64(1),
				AdjustableQuantity: &stripe.PaymentLinkLineItemAdjustableQuantityParams{
					Enabled: stripe.Bool(true),
					Minimum: stripe.Int64(1),
					Maximum: stripe.Int64(10),
				},
			},
		},
		SubscriptionData: &stripe.PaymentLinkSubscriptionDataParams{
			Metadata: map[string]string{
				"email": userInfo.Email,
			},
		},
	}

	result, err := paymentlink.New(params)
	if err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(201, dto.GetPaymentLinkRequest{
		URL: result.URL,
	})
	return
}

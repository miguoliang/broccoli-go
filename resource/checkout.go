package resource

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/common"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentlink"
)

const price = "price_1P0kpBAZeDyeb7mKNxJgIUSN"

const userPoolId = "us-east-1_Qbzi9lvVB"

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

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(userInfo.Email),
	}
	stripeCustomer, err := customer.New(customerParams)
	if err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(500, dto.ErrorResponse{Error: err.Error()})
		return
	}

	userPoolId := userPoolId
	provider := cognitoidentityprovider.NewFromConfig(cfg)
	_, err = provider.AdminUpdateUserAttributes(context.TODO(), &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: &userPoolId,
		Username:   &userInfo.UserId,
		ClientMetadata: map[string]string{
			"StripeCustomerId": stripeCustomer.ID,
		},
	})

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

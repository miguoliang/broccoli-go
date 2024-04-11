package resource

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/common"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
)

func ListSubscriptionsHandler(c *gin.Context) {

	var userInfo common.UserInfo
	err := common.GetUserInfoByContext(c, &userInfo)
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
	user, err := provider.AdminGetUser(context.TODO(), &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &userPoolId,
		Username:   &userInfo.Username,
	})
	if err != nil {
		c.JSON(500, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var stripeCustomerId string
	for _, attr := range user.UserAttributes {
		if *attr.Name == "custom:stripeCustomerId" {
			stripeCustomerId = *attr.Value
			break
		}
	}

	if stripeCustomerId == "" {
		c.JSON(400, dto.ErrorResponse{Error: "No stripe customer id"})
		return
	}

	params := stripe.SubscriptionListParams{
		Customer: stripe.String(stripeCustomerId),
		Status:   stripe.String("all"),
	}
	result := subscription.List(&params)
	c.JSON(200, result.SubscriptionList())
}

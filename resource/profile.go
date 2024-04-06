package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/common"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
)

func ListSubscriptionsHandler(c *gin.Context) {

	userID, err := common.GetUserInfoByContext(c)
	if err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	params := &stripe.SubscriptionListParams{}
}

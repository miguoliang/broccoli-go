package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/common"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
)

// ListSubscriptionsHandler list subscriptions
// @Summary List subscriptions
// @Description List subscriptions
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []dto.Subscription
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /profile/subscriptions [get]
func ListSubscriptionsHandler(c *gin.Context) {

	var userInfo common.UserInfo
	err := common.GetUserInfoByContext(c, &userInfo)
	if err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	params := stripe.SubscriptionSearchParams{
		SearchParams: stripe.SearchParams{
			Query: "status:'active' AND metadata['email']:'" + userInfo.Email + "'",
		},
	}
	result := subscription.Search(&params)
	subs := make([]dto.Subscription, 0)
	for result.Next() {
		sub := result.Subscription()
		subs = append(subs, dto.Subscription{
			ID:            sub.ID,
			Interval:      string(sub.Items.Data[0].Plan.Interval),
			IntervalCount: int(sub.Items.Data[0].Plan.IntervalCount),
		})
	}

	c.JSON(200, subs)
}

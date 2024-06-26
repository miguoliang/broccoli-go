package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/common"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/miguoliang/broccoli-go/internal/persistence"
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

// GetUsageHandler get usage
// @Summary Get usage
// @Description Get usage
// @Tags profile
// @Accept json
// @Produce json
// @Param start_date query string true "Start date"
// @Param end_date query string true "End date"
// @Success 200 {object} []persistence.Usage
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /profile/usages [get]
func GetUsageHandler(c *gin.Context) {

	var request dto.GetUsageRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	var usages []persistence.Usage
	db := persistence.GetDatabaseConnection()
	if result := db.Find(&usages).
		Where("`date` > ?", request.StartDate).
		Where("`date` < ?", request.EndDate); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(200, usages)
}

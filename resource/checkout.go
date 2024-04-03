package resource

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentlink"
	"strings"
)

const price = "price_1P0kpBAZeDyeb7mKNxJgIUSN"

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

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.JSON(400, dto.ErrorResponse{Error: "no jwt token"})
		return
	}
	if gin.IsDebugging() {
		fmt.Println("jwtToken: ", jwtToken)
	}

	userID, exists := getUserId(jwtToken)
	if exists != nil {
		c.JSON(400, dto.ErrorResponse{Error: exists.Error()})
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
		Metadata: map[string]string{
			"user_id": userID,
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

func getUserId(jwtToken string) (string, error) {

	token := jwtToken[7:]
	pieces := strings.Split(token, ".")
	if len(pieces) != 3 {
		return "", fmt.Errorf("invalid token")
	}

	payload := pieces[1]
	jsonStr, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return "", err
	}

	var userData map[string]interface{}
	err = json.Unmarshal(jsonStr, &userData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		// Handle parsing errors (optional)
		return "", err
	}

	userID, exists := userData["sub"]
	if exists == false {
		return "", fmt.Errorf("no user id in token")
	}

	return userID.(string), nil
}

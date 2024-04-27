package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/webhook"
)

// SetupRouter sets up the routes for the API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Group("/webhook").
		POST("/stripe", webhook.StripeWebhookHandler)

	api := r.Group("/api")

	api.Group("/vertex").
		GET("/:id", FindVertexByIdHandler).
		GET("", SearchVerticesHandler).
		POST("", CreateVertexHandler).
		DELETE("/:id", DeleteVertexByIdHandler).
		POST("/:id/property", CreateVertexPropertyHandler)

	api.Group("/edge").
		POST("", CreateEdgeHandler).
		GET("", SearchEdgesHandler)

	api.Group("/p").
		GET("/link", GetPaymentLinkHandler)

	api.Group("/profile").
		GET("/subscriptions", ListSubscriptionsHandler)

	return r
}

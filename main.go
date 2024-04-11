package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/resource"
	"github.com/miguoliang/broccoli-go/webhook"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v76"
	"os"
)

var ginLambda *ginadapter.GinLambda

func main() {

	println(viper.GetString("gin.word"))

	// Set up the router
	r := setupRouter()

	if os.Getenv("LAMBDA_RUNTIME_DIR") != "" {
		// Running on AWS Lambda
		lambda.Start(Handler)
	} else {
		// Running on local
		err := r.Run()
		if err != nil {
			panic(err)
		}
	}
}

func init() {

	// Set configuration file paths based on Gin mode
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Look for configuration files in the current directory

	viper.AutomaticEnv() // Enable automatic environment variable parsing
	err := viper.ReadInConfig()
	if err != nil {
		panic("No configuration file loaded - using defaults")
	}

	stripe.Key = viper.GetString("stripe.secret_key")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

// @title Broccoli API
// @description This is the API for Broccoli
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @schemes http
// @schemes https
// @contact.name Guoliang Mi
// @contact.email boymgl@qq.com
// @contact.url https://miguoliang.com
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Group("/webhook").
		POST("/stripe", webhook.StripeWebhookHandler)

	api := r.Group("/api")

	api.Group("/vertex").
		GET("/:id", resource.FindVertexByIdHandler).
		GET("", resource.SearchVerticesHandler).
		POST("", resource.CreateVertexHandler).
		DELETE("/:id", resource.DeleteVertexByIdHandler).
		POST("/:id/property", resource.CreateVertexPropertyHandler)

	api.Group("/edge").
		POST("", resource.CreateEdgeHandler).
		GET("", resource.SearchEdgesHandler)

	api.Group("/p").
		GET("/link", resource.GetPaymentLinkHandler)

	api.Group("/profile").
		GET("/subscriptions", resource.ListSubscriptionsHandler)

	ginLambda = ginadapter.New(r)

	return r
}

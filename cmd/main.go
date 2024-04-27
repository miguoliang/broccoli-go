package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/resource"
	"github.com/stripe/stripe-go/v76"
	"os"
)

var ginLambda *ginadapter.GinLambda

func init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	mode := os.Getenv("GIN_MODE")
	if mode != "" {
		gin.SetMode(mode)
	}
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
func main() {

	// Set up the router
	r := resource.SetupRouter()

	ginLambda = ginadapter.New(r)

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

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

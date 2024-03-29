package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiepd/cognito-go"
	"github.com/miguoliang/broccoli-go/resource"
)

func main() {

	// Set up cognito
	c, _ := cognito.NewCognitoClient("ap-southeast-2", "cognito-app", "xxx")

	// Set up the router
	r := setupRouter(c)

	// Run the server
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func setupRouter(c cognito.Client) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	api.Group("/vertex", c.Authorize).
		GET("/:id", resource.FindVertexByIdHandler).
		GET("", resource.SearchVerticesHandler).
		POST("", resource.CreateVertexHandler).
		DELETE("/:id", resource.DeleteVertexByIdHandler).
		POST("/:id/property", resource.CreateVertexPropertyHandler)

	api.Group("/edge", c.Authorize).
		POST("", resource.CreateEdgeHandler).
		GET("", resource.SearchEdgesHandler)
	return r
}

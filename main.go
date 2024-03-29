package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/middleware"
	"github.com/miguoliang/broccoli-go/resource"
)

func main() {

	// Set up the router
	r := setupRouter(middleware.CheckJWT("user"))

	// Run the server
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func setupRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api", handlerFunc)

	api.Group("/vertex").
		GET("/:id", resource.FindVertexByIdHandler).
		GET("", resource.SearchVerticesHandler).
		POST("", resource.CreateVertexHandler).
		DELETE("/:id", resource.DeleteVertexByIdHandler).
		POST("/:id/property", resource.CreateVertexPropertyHandler)

	api.Group("/edge").
		POST("", resource.CreateEdgeHandler).
		GET("", resource.SearchEdgesHandler)
	return r
}

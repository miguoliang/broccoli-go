package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/miguoliang/broccoli-go/persistence"
)

func CreateVertexPropertyHandler(c *gin.Context) {
	var createVertexPropertyRequest dto.CreateVertexPropertyRequest
	if err := c.ShouldBindJSON(&createVertexPropertyRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var vertex persistence.Vertex
	var id = c.Param("id")
	db := persistence.GetDatabaseConnection()
	if result := db.First(&vertex, id); result.Error != nil {
		c.JSON(404, dto.ErrorResponse{Error: "vertex not found"})
		return
	}
	var vertexProperty persistence.VertexProperty
	vertexProperty.VertexID = vertex.ID
	vertexProperty.Key = createVertexPropertyRequest.Key
	vertexProperty.Value = createVertexPropertyRequest.Value
	if result := db.Create(&vertexProperty); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, dto.CreateVertexPropertyResponse{ID: vertexProperty.ID})
}

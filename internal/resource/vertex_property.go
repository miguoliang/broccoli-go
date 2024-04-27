package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/miguoliang/broccoli-go/internal/persistence"
)

// CreateVertexPropertyHandler
// @Summary Create vertex property
// @Description Create vertex property
// @ID create-vertex-property
// @Produce json
// @Param id path int true "Vertex ID"
// @Param request body dto.CreateVertexPropertyRequest true "Request body"
// @Success 201 {object} dto.CreateVertexPropertyResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /vertex/{id}/property [post]
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
	c.JSON(201, dto.CreatedResponse{ID: vertexProperty.ID})
}

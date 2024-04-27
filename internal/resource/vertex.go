package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/internal/dto"
	"github.com/miguoliang/broccoli-go/internal/persistence"
	"strings"
)

// FindVertexByIdHandler
// @Summary Find vertex by ID
// @Description Find vertex by ID
// @ID find-vertex-by-id
// @Produce json
// @Param id path int true "Vertex ID"
// @Success 200 {object} persistence.Vertex
// @Failure 404 {object} dto.ErrorResponse
// @Router /vertex/{id} [get]
func FindVertexByIdHandler(c *gin.Context) {
	id := c.Param("id")
	var vertex persistence.Vertex
	db := persistence.GetDatabaseConnection()
	if result := db.Preload("Properties").First(&vertex, id); result.Error != nil {
		c.JSON(404, dto.ErrorResponse{Error: "vertex not found"})
		return
	}
	c.JSON(200, vertex)
}

// SearchVerticesHandler
// @Summary Search vertices
// @Description Search vertices
// @ID search-vertices
// @Produce json
// @Param q query string true "Search query"
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {object} dto.SearchVerticesResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /vertex [get]
func SearchVerticesHandler(c *gin.Context) {
	var searchVerticesRequest dto.SearchVerticesRequest
	if err := c.ShouldBindQuery(&searchVerticesRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var vertices []persistence.Vertex
	limit := searchVerticesRequest.Size
	if limit == 0 {
		limit = 10
	}
	offset := (searchVerticesRequest.Page - 1) * limit
	db := persistence.GetDatabaseConnection()
	if result := db.Offset(offset).Limit(limit).Where("name LIKE ?", "%"+searchVerticesRequest.Q+"%").Find(&vertices); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	var total int64
	if result := db.Model(&persistence.Vertex{}).Where("name LIKE ?", "%"+searchVerticesRequest.Q+"%").Count(&total); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(200, dto.SearchVerticesResponse{
		Total: total,
		Page:  searchVerticesRequest.Page,
		Size:  searchVerticesRequest.Size,
		Data:  vertices,
	})
}

// CreateVertexHandler
// @Summary Create vertex
// @Description Create vertex
// @ID create-vertex
// @Accept json
// @Produce json
// @Param request body dto.CreateVertexRequest true "Create Vertex Request"
// @Success 201 {object} dto.CreateVertexResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /vertex [post]
func CreateVertexHandler(c *gin.Context) {
	var createVertexRequest dto.CreateVertexRequest
	if err := c.ShouldBindJSON(&createVertexRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var vertex persistence.Vertex
	vertex.Name = createVertexRequest.Name
	vertex.Type = createVertexRequest.Type
	properties := make([]persistence.VertexProperty, 0)
	for k, v := range createVertexRequest.Properties {
		properties = append(properties, persistence.VertexProperty{
			Key:   k,
			Value: v,
		})
	}
	vertex.Properties = properties
	db := persistence.GetDatabaseConnection()
	if result := db.Create(&vertex); result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			c.JSON(409, dto.ErrorResponse{Error: "name and type must be unique"})
			return
		}
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, dto.CreatedResponse{ID: vertex.ID})
}

// DeleteVertexByIdHandler
// @Summary Delete vertex by ID
// @Description Delete vertex by ID
// @ID delete-vertex-by-id
// @Produce json
// @Param id path int true "Vertex ID"
// @Success 204
// @Failure 500 {object} dto.ErrorResponse
// @Router /vertex/{id} [delete]
func DeleteVertexByIdHandler(c *gin.Context) {
	id := c.Param("id")

	// Check if vertex exists
	var vertex persistence.Vertex
	if result := persistence.GetDatabaseConnection().First(&vertex, id); result.Error != nil {
		c.JSON(404, dto.ErrorResponse{Error: "vertex not found"})
		return
	}

	// Delete vertex
	db := persistence.GetDatabaseConnection()
	if result := db.Delete(&persistence.Vertex{}, id); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(204, nil)
}

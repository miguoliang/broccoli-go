package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/miguoliang/broccoli-go/persistence"
)

// FindVertexByIdHandler
// @Summary Find vertex by ID
// @Description Find vertex by ID
// @ID find-vertex-by-id
// @Produce json
// @Param id path int true "Vertex ID"
// @Success 200 {object} Vertex
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
	c.JSON(200, dto.PageResponse[persistence.Vertex]{
		Total: total,
		Page:  searchVerticesRequest.Page,
		Size:  searchVerticesRequest.Size,
		Data:  vertices,
	})
}

func CreateVertexHandler(c *gin.Context) {
	var createVertexRequest dto.CreateVertexRequest
	if err := c.ShouldBindJSON(&createVertexRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var vertex persistence.Vertex
	vertex.Name = createVertexRequest.Name
	vertex.Type = createVertexRequest.Type
	db := persistence.GetDatabaseConnection()
	if result := db.Create(&vertex); result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: vertices.name, vertices.type" {
			c.JSON(409, dto.ErrorResponse{Error: "name and type must be unique"})
			return
		}
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, dto.CreateVertexResponse{ID: vertex.ID})
}

func DeleteVertexByIdHandler(c *gin.Context) {
	id := c.Param("id")
	db := persistence.GetDatabaseConnection()
	if result := db.Delete(&persistence.Vertex{}, id); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(204, nil)
}

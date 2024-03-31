package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/miguoliang/broccoli-go/persistence"
	"strings"
)

// CreateEdgeHandler
// @Summary Create edge
// @Description Create edge
// @ID create-edge
// @Produce json
// @Param request body dto.CreateEdgeRequest true "Request body"
// @Success 201 {object} dto.CreateEdgeResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /edge [post]
func CreateEdgeHandler(c *gin.Context) {
	var createEdgeRequest dto.CreateEdgeRequest
	if err := c.ShouldBindJSON(&createEdgeRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var edge persistence.Edge
	edge.From = createEdgeRequest.From
	edge.To = createEdgeRequest.To
	edge.Type = createEdgeRequest.Type
	db := persistence.GetDatabaseConnection()
	if result := db.Create(&edge); result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			c.JSON(409, dto.ErrorResponse{Error: "from, to and type must be unique"})
			return
		}
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, dto.CreateEdgeResponse{ID: edge.ID})
}

// SearchEdgesHandler
// @Summary Search edges
// @Description Search edges
// @ID search-edges
// @Produce json
// @Param type query string false "Type"
// @Param from query string false "From"
// @Param to query string false "To"
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} dto.PageResponse[persistence.Edge]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /edge [get]
func SearchEdgesHandler(c *gin.Context) {
	var searchEdgesRequest dto.SearchEdgesRequest
	if err := c.ShouldBindQuery(&searchEdgesRequest); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}
	var edges []persistence.Edge
	limit := searchEdgesRequest.Size
	if limit == 0 {
		limit = 10
	}
	offset := (searchEdgesRequest.Page - 1) * limit
	db := persistence.GetDatabaseConnection()
	if result := db.Offset(offset).
		Where("`type` in ?", searchEdgesRequest.Type).
		Where("`from` in ?", searchEdgesRequest.From).
		Where("`to` in ?", searchEdgesRequest.To).
		Limit(limit).
		Find(&edges); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}

	var total int64
	if result := db.Model(&persistence.Edge{}).
		Where("`type` in ?", searchEdgesRequest.Type).
		Where("`from` in ?", searchEdgesRequest.From).
		Where("`to` in ?", searchEdgesRequest.To).
		Count(&total); result.Error != nil {
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(200, dto.PageResponse[persistence.Edge]{
		Total: total,
		Page:  searchEdgesRequest.Page,
		Size:  searchEdgesRequest.Size,
		Data:  edges,
	})
}

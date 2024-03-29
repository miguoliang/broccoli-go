package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/broccoli-go/dto"
	"github.com/miguoliang/broccoli-go/persistence"
)

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
		if result.Error.Error() == "UNIQUE constraint failed: edges.from, edges.to, edges.type" {
			c.JSON(409, dto.ErrorResponse{Error: "from, to and type must be unique"})
			return
		}
		c.JSON(500, dto.ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, dto.CreateEdgeResponse{ID: edge.ID})
}

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

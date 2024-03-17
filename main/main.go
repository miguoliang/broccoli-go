package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// Open a database connection
	db = connectDatabase()
	autoMigrate()

	// Migrate the schema
	r := setUpRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/vertex/:id", findVertexByIdHandler)
	api.GET("/vertex", searchVerticesHandler)
	api.POST("/vertex", createVertexHandler)
	api.DELETE("/vertex/:id", deleteVertexByIdHandler)
	api.POST("/vertex/:id/property", createVertexPropertyHandler)
	api.POST("/edge", createEdgeHandler)
	api.GET("/edge", searchEdgesHandler)
	return r
}

func autoMigrate() {
	err := db.AutoMigrate(&Vertex{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&VertexProperty{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Edge{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&EdgeProperty{})
	if err != nil {
		return
	}
}

func connectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(`file::memory:?cache=shared`), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func findVertexByIdHandler(c *gin.Context) {
	id := c.Param("id")
	var vertex Vertex
	if result := db.Preload("Properties").First(&vertex, id); result.Error != nil {
		c.JSON(404, ErrorResponse{Error: "vertex not found"})
		return
	}
	c.JSON(200, vertex)
}

func searchVerticesHandler(c *gin.Context) {
	var searchVerticesRequest SearchVerticesRequest
	if err := c.ShouldBindQuery(&searchVerticesRequest); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	var vertices []Vertex
	limit := searchVerticesRequest.Size
	if limit == 0 {
		limit = 10
	}
	offset := (searchVerticesRequest.Page - 1) * limit
	if result := db.Offset(offset).Limit(limit).Where("name LIKE ?", "%"+searchVerticesRequest.Q+"%").Find(&vertices); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	var total int64
	if result := db.Model(&Vertex{}).Where("name LIKE ?", "%"+searchVerticesRequest.Q+"%").Count(&total); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(200, PageResponse[Vertex]{
		Total: total,
		Page:  searchVerticesRequest.Page,
		Size:  searchVerticesRequest.Size,
		Data:  vertices,
	})
}

func createVertexHandler(c *gin.Context) {
	var createVertexRequest CreateVertexRequest
	if err := c.ShouldBindJSON(&createVertexRequest); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	var vertex Vertex
	vertex.Name = createVertexRequest.Name
	vertex.Type = createVertexRequest.Type
	if result := db.Create(&vertex); result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: vertices.name, vertices.type" {
			c.JSON(409, ErrorResponse{Error: "name and type must be unique"})
			return
		}
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, CreateVertexResponse{ID: vertex.ID})
}

func deleteVertexByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if result := db.Delete(&Vertex{}, id); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(204, nil)
}

func createEdgeHandler(c *gin.Context) {
	var createEdgeRequest CreateEdgeRequest
	if err := c.ShouldBindJSON(&createEdgeRequest); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	var edge Edge
	edge.From = createEdgeRequest.From
	edge.To = createEdgeRequest.To
	edge.Type = createEdgeRequest.Type
	if result := db.Create(&edge); result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: edges.from, edges.to, edges.type" {
			c.JSON(409, ErrorResponse{Error: "from, to and type must be unique"})
			return
		}
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, CreateEdgeResponse{ID: edge.ID})
}

func searchEdgesHandler(c *gin.Context) {
	var searchEdgesRequest SearchEdgesRequest
	if err := c.ShouldBindQuery(&searchEdgesRequest); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	var edges []Edge
	limit := searchEdgesRequest.Size
	if limit == 0 {
		limit = 10
	}
	offset := (searchEdgesRequest.Page - 1) * limit
	if result := db.Offset(offset).
		Where("`type` in ?", searchEdgesRequest.Type).
		Where("`from` in ?", searchEdgesRequest.From).
		Where("`to` in ?", searchEdgesRequest.To).
		Limit(limit).
		Find(&edges); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}

	var total int64
	if result := db.Model(&Edge{}).
		Where("`type` in ?", searchEdgesRequest.Type).
		Where("`from` in ?", searchEdgesRequest.From).
		Where("`to` in ?", searchEdgesRequest.To).
		Count(&total); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(200, PageResponse[Edge]{
		Total: total,
		Page:  searchEdgesRequest.Page,
		Size:  searchEdgesRequest.Size,
		Data:  edges,
	})
}

func createVertexPropertyHandler(c *gin.Context) {
	var createVertexPropertyRequest CreateVertexPropertyRequest
	if err := c.ShouldBindJSON(&createVertexPropertyRequest); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	var vertex Vertex
	var id = c.Param("id")
	if result := db.First(&vertex, id); result.Error != nil {
		c.JSON(404, ErrorResponse{Error: "vertex not found"})
		return
	}
	var vertexProperty VertexProperty
	vertexProperty.VertexID = vertex.ID
	vertexProperty.Key = createVertexPropertyRequest.Key
	vertexProperty.Value = createVertexPropertyRequest.Value
	if result := db.Create(&vertexProperty); result.Error != nil {
		c.JSON(500, ErrorResponse{Error: result.Error.Error()})
		return
	}
	c.JSON(201, CreateVertexPropertyResponse{ID: vertexProperty.ID})
}

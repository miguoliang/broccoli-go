package broccoli_go

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// Open a database connection
	db = connectDatabase()

	// Migrate the schema
	r := setUpRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/vertex/:id", findVertexByIdHandler)
	r.GET("/vertex", searchVerticesHandler)
	r.POST("/vertex", createVertexHandler)
	r.DELETE("/vertex/:id", deleteVertexByIdHandler)
	return r
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
	if result := db.First(&vertex, id); result.Error != nil {
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

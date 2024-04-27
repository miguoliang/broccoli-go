package persistence

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	database *gorm.DB
)

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Vertex{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&VertexProperty{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Edge{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&EdgeProperty{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Checkout{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db, err := gorm.Open(sqlite.Open(`file::memory:?cache=shared`), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigrate(db)
	database = db
}

func GetDatabaseConnection() *gorm.DB {
	return database
}

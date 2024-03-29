package persistence

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

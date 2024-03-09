package main

import (
	"gorm.io/gorm"
	"time"
)

type OperationInfo struct {
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

type Vertex struct {
	gorm.Model
	ID            uint             `gorm:"primaryKey"`
	Name          string           `gorm:"uniqueIndex:vertex_name_type_idx"`
	Type          string           `gorm:"uniqueIndex:vertex_name_type_idx"`
	Properties    []VertexProperty `gorm:"foreignKey:VertexID"`
	OperationInfo OperationInfo    `gorm:"embedded"`
}

type VertexProperty struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	VertexID      uint
	Key           string
	Value         string
	OperationInfo OperationInfo `gorm:"embedded"`
}

type Edge struct {
	gorm.Model
	ID            uint           `gorm:"primaryKey"`
	FromVertex    uint           `gorm:"uniqueIndex:edge_idx"`
	ToVertex      uint           `gorm:"uniqueIndex:edge_idx"`
	Type          string         `gorm:"uniqueIndex:edge_idx"`
	Properties    []EdgeProperty `gorm:"foreignKey:EdgeID"`
	OperationInfo OperationInfo  `gorm:"embedded"`
}

type EdgeProperty struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	EdgeID        uint
	Key           string
	Value         string
	OperationInfo OperationInfo `gorm:"embedded"`
}

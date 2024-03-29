package persistence

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Vertex struct {
	Model
	Name       string           `json:"name" gorm:"uniqueIndex:vertex_name_type_idx"`
	Type       string           `json:"type" gorm:"uniqueIndex:vertex_name_type_idx"`
	Properties []VertexProperty `json:"properties" gorm:"foreignKey:VertexID"`
}

type VertexProperty struct {
	Model
	VertexID uint   `json:"vertex_id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type Edge struct {
	Model
	From       uint           `json:"from" gorm:"uniqueIndex:edge_idx"`
	To         uint           `json:"to" gorm:"uniqueIndex:edge_idx"`
	Type       string         `json:"type" gorm:"uniqueIndex:edge_idx"`
	Properties []EdgeProperty `json:"properties" gorm:"foreignKey:EdgeID"`
}

type EdgeProperty struct {
	Model
	EdgeID uint   `json:"edge_id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

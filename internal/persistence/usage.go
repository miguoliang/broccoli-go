package persistence

import "time"

type Usage struct {
	Model
	Date              time.Time `gorm:"index" json:"date"`
	VertexNum         uint      `gorm:"index" json:"vertex_num"`
	EdgeNum           uint      `gorm:"index" json:"edge_num"`
	VertexPropertyNum uint      `gorm:"index" json:"vertex_property_num"`
}

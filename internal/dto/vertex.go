package dto

import "github.com/miguoliang/broccoli-go/internal/persistence"

type SearchVerticesRequest struct {
	Q string `form:"q"`
	PageParams
}

type CreateVertexRequest struct {
	Name       string            `json:"name" binding:"required"`
	Type       string            `json:"type" binding:"required"`
	Properties map[string]string `json:"properties"`
}

type SearchVerticesResponse PageResponse[persistence.Vertex]

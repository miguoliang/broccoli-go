package dto

import "github.com/miguoliang/broccoli-go/internal/persistence"

type SearchVerticesRequest struct {
	Q string `form:"q"`
	PageParams
}

type CreateVertexRequest struct {
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
}

type CreateVertexResponse struct {
	ID uint `json:"id"`
}

type SearchVerticesResponse PageResponse[persistence.Vertex]

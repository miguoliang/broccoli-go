package dto

import "github.com/miguoliang/broccoli-go/internal/persistence"

type CreateEdgeRequest struct {
	From uint   `json:"from" min:"1" binding:"required"`
	To   uint   `json:"to" min:"1" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type SearchEdgesRequest struct {
	From []uint   `form:"from" binding:"required"`
	To   []uint   `form:"to" binding:"required"`
	Type []string `form:"type" binding:"required"`
	PageParams
}

type SearchEdgesResponse PageResponse[persistence.Edge]

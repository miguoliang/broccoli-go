package dto

import "github.com/miguoliang/broccoli-go/internal/persistence"

type CreateEdgeRequest struct {
	From uint   `json:"from"`
	To   uint   `json:"to"`
	Type string `json:"type"`
}

type CreateEdgeResponse struct {
	ID uint `json:"id"`
}

type SearchEdgesRequest struct {
	From []uint   `form:"from" binding:"required"`
	To   []uint   `form:"to" binding:"required"`
	Type []string `form:"type" binding:"required"`
	PageParams
}

type SearchEdgesResponse PageResponse[persistence.Edge]

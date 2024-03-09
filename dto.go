package main

type PageParams struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type SearchVerticesRequest struct {
	Q string `json:"q"`
	PageParams
}

type CreateVertexRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateVertexResponse struct {
	ID uint `json:"id"`
}

package broccoli_go

type ErrorResponse struct {
	Error string `json:"error"`
}

type PageParams struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type PageResponse[T any] struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Data  []T   `json:"data"`
}

type SearchVerticesRequest struct {
	Q string `form:"q"`
	PageParams
}

type CreateVertexRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateVertexResponse struct {
	ID uint `json:"id"`
}

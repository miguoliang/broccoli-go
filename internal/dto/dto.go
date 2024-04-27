package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreatedResponse struct {
	ID uint `json:"id"`
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

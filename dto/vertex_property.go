package dto

type CreateVertexPropertyRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateVertexPropertyResponse struct {
	ID uint `json:"id"`
}

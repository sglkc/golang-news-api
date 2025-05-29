package models

type BaseResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DataResponse[T any] struct {
	BaseResponse
	Data T `json:"data"`
}

type PaginatedResponse[T any] struct {
	DataResponse[[]T]
	Count int `json:"count"`
	Pages int `json:"pages"`
}

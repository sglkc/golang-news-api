package models

type BaseResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DataResponse[T any] struct {
	BaseResponse
	Data T `json:"data"`
}

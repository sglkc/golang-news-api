package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/sglkc/golang-news-api/models"
)

func SendJSON(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
) {
	var success bool

	if status >= 200 && status < 400 {
		success = true
	} else {
		success = false
	}

	render.Status(r, status)
	render.JSON(w, r, models.BaseResponse{
		Message: message,
		Success: success,
	})
}

func SendData[T any](
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	data T,
) {
	var success bool

	if status >= 200 && status < 400 {
		success = true
	} else {
		success = false
	}

	render.Status(r, status)
	render.JSON(w, r, models.DataResponse[T]{
		BaseResponse: models.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: data,
	})
}

func SendPage[T any](
	w http.ResponseWriter,
	r *http.Request,
	data []T,
) {
	var items []T
	var page, limit string

	page = r.URL.Query().Get("page")
	limit = r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}

	start := (pageInt - 1) * limitInt
	end := start + limitInt
	count := len(data)

	// if start > count {
	// 	items = []T{}
	// }
	if end > count {
		end = count
	}

	items = data[start:end]

	render.Status(r, http.StatusOK)
	render.JSON(w, r, models.PaginatedResponse[T]{
		DataResponse: models.DataResponse[[]T]{
			BaseResponse: models.BaseResponse{
				Message: "Success",
				Success: true,
			},
			Data: items,
		},
		Count: count,
		Pages: (count + limitInt - 1) / limitInt,
	})
}

package utils

import (
	"net/http"

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

func SendData(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	data any,
) {
	var success bool

	if status >= 200 && status < 400 {
		success = true
	} else {
		success = false
	}

	render.Status(r, status)
	render.JSON(w, r, models.DataResponse[any]{
		BaseResponse: models.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: data,
	})
}

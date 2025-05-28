package controllers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/sglkc/golang-news-api/database"
	"github.com/sglkc/golang-news-api/models"
	"github.com/sglkc/golang-news-api/utils"
)

func ListNews(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	var news []models.News
	result := db.Find(&news)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot fetch news")
		return
	}

	utils.SendData(w, r, http.StatusOK, "Success", news)
}

func GetNews(w http.ResponseWriter, r *http.Request) {}

func CreateNews(w http.ResponseWriter, r *http.Request) {
	news := models.News{}

	err := render.DecodeJSON(r.Body, &news)
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid body")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	result := db.Create(&news)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot create news")
		return
	}

	utils.SendData(w, r, http.StatusCreated, "Success", news)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {}

func DeleteNews(w http.ResponseWriter, r *http.Request) {}

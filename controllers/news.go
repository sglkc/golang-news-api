package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func GetNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid ID")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	var news models.News
	result := db.First(&news, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find news")
		return
	}

	utils.SendData(w, r, http.StatusOK, "Success", news)
}

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

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid ID")
		return
	}

	newNews := models.News{}
	err = render.DecodeJSON(r.Body, &newNews)
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid body")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	var news models.News
	result := db.First(&news, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find news")
		return
	}

	result = db.Model(&news).Updates(&newNews)

	utils.SendData(w, r, http.StatusOK, "Success", news)
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid ID")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	// TODO: fix double query
	var news models.News
	result := db.First(&news, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find news")
		return
	}

	db.Delete(&models.News{}, id)

	utils.SendJSON(w, r, http.StatusOK, "Success")
}

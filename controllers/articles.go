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

func ListArticles(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	var articles []models.Article
	result := db.Find(&articles)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot fetch articles")
		return
	}

	utils.SendPage(w, r, articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
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

	var article models.Article
	result := db.First(&article, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find article")
		return
	}

	utils.SendData(w, r, http.StatusOK, "Success", article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}
	err := render.DecodeJSON(r.Body, &article)
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid body")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	result := db.Create(&article)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot create article")
		return
	}

	utils.SendData(w, r, http.StatusCreated, "Success", article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid ID")
		return
	}

	newArticle := models.Article{}
	err = render.DecodeJSON(r.Body, &newArticle)
	if err != nil {
		utils.SendJSON(w, r, http.StatusBadRequest, "Invalid body")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Internal error")
		return
	}

	var article models.Article
	result := db.First(&article, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find article")
		return
	}

	result = db.Model(&article).Updates(&newArticle)

	utils.SendData(w, r, http.StatusOK, "Success", article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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
	var article models.Article
	result := db.First(&article, id)
	if result.Error != nil {
		utils.SendJSON(w, r, http.StatusInternalServerError, "Cannot find article")
		return
	}

	db.Delete(&models.Article{}, id)

	utils.SendJSON(w, r, http.StatusOK, "Success")
}

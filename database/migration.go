package database

import (
	"github.com/sglkc/golang-news-api/models"
)

func Migrate() error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.News{})
	if err != nil {
		return err
	}

	return nil
}

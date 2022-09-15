package database

import (
	"day2/config"
	"day2/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func FindBookById(id string) (interface{}, error) {
	var book models.Book

	if err := config.DB.Find(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(book models.Book) (interface{}, error) {

	if err := config.DB.Model(&book).Updates(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(id string) error {
	var book models.Book

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return err
	}

	return nil
}

func NewBook(book models.Book) error {

	if err := config.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

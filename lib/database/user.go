package database

import (
	"day2/config"
	"day2/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserById(id string) (interface{}, error) {
	var user models.User

	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user models.User) (interface{}, error) {

	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id string) error {
	var user models.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func NewUser(user models.User) error {

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

package repository

import (
	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/models/db"
)

func CreateUser(user *models.User) (*models.User, error) {
	if err := db.DB().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByFieldAndValue(field string, value interface{}) (*models.User, error) {
	var user models.User
	if field == "id" {
		if err := db.DB().First(&user, value).Error; err != nil {
			return nil, err
		}

	} else if err := db.DB().Where(field+"= ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserById(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserProjects(user *models.User) ([]models.Project, error) {
	var projects []models.Project
	if err := db.DB().Preload("User").Preload("Users").Model(user).Association("Projects").Find(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}

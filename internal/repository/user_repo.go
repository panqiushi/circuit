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

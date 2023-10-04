package repository

import (
	"code.circuit.io/circuit/internal/models"
	"code.circuit.io/circuit/internal/models/db"
)

func CreateUser(user *models.User) (*models.User, error) {
	dbInstance := db.DB()
	dbInstance.AutoMigrate(&models.User{})
	if err := dbInstance.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

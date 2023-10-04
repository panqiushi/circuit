package repository

import (
	"code.circuit.io/circuit/internal/models"
	"code.circuit.io/circuit/internal/models/db"
)

func CreateUser(user *models.User) (*models.User, error) {
	if err := db.DB().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

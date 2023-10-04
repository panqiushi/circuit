package service

import (
	"code.circuit.io/circuit/internal/models"
	"code.circuit.io/circuit/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateUserIfNecessary(c *gin.Context) (*models.User, error) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return nil, err
	}

	if user.ID != 0 {
		return &user, nil
	}

	createdUser, err := repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

package service

import (
	"errors"

	consts "code.circuit.io/circuit/internal/const"
	"code.circuit.io/circuit/internal/models"
	"code.circuit.io/circuit/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserIfNecessary(c *gin.Context) (*models.User, error) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return nil, err
	}

	if user.ID != 0 {
		return repository.FindUserByFieldAndValue("id", user.ID)
	}

	if err := validateUser(&user); err != nil {
		return nil, err
	}

	exixtedUser, err := repository.FindUserByFieldAndValue("email", user.Email)
	if (err != nil && err.Error() != consts.RECORD_NOT_FOUND) || exixtedUser != nil {
		return nil, err
	}

	exixtedUser, err = repository.FindUserByFieldAndValue("phone", user.Phone)
	if (err != nil && err.Error() != consts.RECORD_NOT_FOUND) || exixtedUser != nil {
		return nil, err
	}

	exixtedUser, err = repository.FindUserByFieldAndValue("name", user.Name)
	if (err != nil && err.Error() != consts.RECORD_NOT_FOUND) || exixtedUser != nil {
		return nil, err
	}

	user.HashPassword, err = HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	createdUser, err := repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func validateUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Phone == "" {
		return errors.New("phone is required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

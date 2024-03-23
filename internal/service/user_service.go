package service

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	consts "circuit.io/circuit/internal/const"
	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Email string `json:"email" form:"email"`
	// Phone    string `json:"phone"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("73RnG8x2hLs6YpDfWvQ9Uz1cVbM5XtJnKw4FgSdAeZoPqCtH6rL0aVtB1iZw3lNfJ7sKgDm1fVhA8tH7wBfYsE9dM2hG8dQ4oZsR")

func CreateUserIfNecessary(c *gin.Context) (*models.User, error) {
	var user models.User
	if err := c.Bind(&user); err != nil {
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

	// exixtedUser, err = repository.FindUserByFieldAndValue("phone", user.Phone)
	// if (err != nil && err.Error() != consts.RECORD_NOT_FOUND) || exixtedUser != nil {
	// 	return nil, err
	// }

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

func LoginHandler(c *gin.Context) (int, error) {
	var loginUser LoginUser
	if err := c.Bind(&loginUser); err != nil {
		return http.StatusBadRequest, err
	}

	field := "email"
	value := loginUser.Email
	// if loginUser.Email == "" {
	// 	field = "phone"
	// 	value = loginUser.Phone
	// }

	user, err := repository.FindUserByFieldAndValue(field, value)
	if err != nil && err.Error() != consts.RECORD_NOT_FOUND {
		return http.StatusNotFound, err
	}

	if !CheckPassword(loginUser.Password, user.HashPassword) {
		return http.StatusUnauthorized, errors.New("invalid password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to generate token")
	}

	c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)
	c.SetCookie("userName", user.Name, 3600*24, "/", "", false, true)
	c.SetCookie("userId", strconv.FormatUint(uint64(user.ID), 10), 3600*24, "/", "", false, true)
	return http.StatusOK, nil
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

	// if user.Phone == "" {
	// 	return errors.New("phone is required")
	// }

	if user.Name == "" {
		return errors.New("name is required")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func GetUserProjects(c *gin.Context) (int, []models.Project, error) {
	userId, err := c.Cookie("userId")
	if err != nil {
		return http.StatusUnauthorized, nil, err
	}

	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	user, err := repository.FindUserById(uint(uid))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	projects, err := repository.FindUserProjects(user)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	log.Default().Println("Projects: ", projects)
	return http.StatusOK, projects, nil
}

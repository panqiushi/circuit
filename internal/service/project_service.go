package service

import (
	"net/http"
	"strconv"

	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) (int, *models.Project, error) {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	userId, err := c.Cookie("userId")
	if err != nil {
		return http.StatusUnauthorized, nil, err
	}

	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	project.UserId = uint(uid)

	user, err := repository.FindUserById(uint(uid))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	createdProject, err := repository.CreateProject(&project)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	createdProject, err = repository.AddMemberToProject(createdProject, user)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, createdProject, nil
}

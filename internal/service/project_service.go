package service

import (
	"log"
	"net/http"
	"strconv"

	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) (int, *models.Project, error) {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	existingProject, err := repository.FindProjectByName(project.Name)

	if existingProject != nil {
		return http.StatusBadRequest, nil, gorm.ErrDuplicatedKey
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, nil, err
	}

	userId, err := c.Cookie("userId")
	if err != nil {
		return http.StatusUnauthorized, nil, err
	}
	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil || uid == 0 {
		return http.StatusInternalServerError, nil, err
	}
	project.UserId = uint(uid)
	user, err := repository.FindUserById(project.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	project.User = *user
	log.Default().Println("User: ", project.User.ID, project.UserId)
	createdProject, err := repository.CreateProject(&project)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	log.Default().Println("Created User: ", createdProject.UserId, createdProject.User.ID)
	createdProject, err = repository.AddMemberToProject(createdProject, user)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, createdProject, nil
}

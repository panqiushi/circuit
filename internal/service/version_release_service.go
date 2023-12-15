package service

import (
	"net/http"
	"strconv"

	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateVersionRelease(c *gin.Context) (int, *models.VersionRelease, error) {
	var version_release models.VersionRelease
	if err := c.Bind(&version_release); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	existingVersionRelease, err := repository.FindVersionReleaseByName(version_release.VersionName)

	if existingVersionRelease != nil {
		return http.StatusBadRequest, nil, gorm.ErrDuplicatedKey
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, nil, err
	}

	project, err := repository.FindProjectById(version_release.ProjectID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	version_release.Project = *project

	userId, err := c.Cookie("userId")
	if err != nil {
		return http.StatusUnauthorized, nil, err
	}
	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil || uid == 0 {
		return http.StatusInternalServerError, nil, err
	}
	version_release.UserId = uint(uid)
	user, err := repository.FindUserById(version_release.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	version_release.User = *user
	createdVersionRelease, err := repository.CreateVersionRelease(&version_release)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	_, err = repository.AddVersionReleaseToProject(project, createdVersionRelease)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, createdVersionRelease, nil
}

package api

import (
	"net/http"
	"strconv"

	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"
	"circuit.io/circuit/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRouters(router *gin.Engine) {
	router.POST("/project", service.AuthMiddleware(), func(context *gin.Context) {
		var project models.Project
		if err := context.Bind(&project); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		userId, err := context.Cookie("userId")
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
		uid, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		project.UserId = uint(uid)
		createdProject, err := repository.CreateProject(&project)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, createdProject)
		}
	})

	router.GET("/project/:name", service.AuthMiddleware(), func(context *gin.Context) {
		name := context.Param("name")
		project, err := repository.FindProjectByName(name)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, project)
		}
	})
}

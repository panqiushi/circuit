package api

import (
	"net/http"

	"circuit.io/circuit/internal/repository"
	"circuit.io/circuit/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRouters(router *gin.Engine) {
	router.POST("/project", service.AuthMiddleware(), func(context *gin.Context) {
		code, _, err := service.CreateProject(context)
		if err != nil || code != http.StatusOK {
			context.JSON(code, gin.H{
				"error": err.Error(),
			})
		} else {
			context.Redirect(http.StatusMovedPermanently, "/projects")
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

	router.GET("/projects/all", service.AuthMiddleware(), func(context *gin.Context) {
		projects, err := repository.FindAllProject()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, projects)
		}
	})
}

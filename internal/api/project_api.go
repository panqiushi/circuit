package api

import (
	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRouters(router *gin.Engine) {
	router.POST("/project", func(context *gin.Context) {
		var project models.Project
		if err := context.Bind(&project); err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		createdProject, err := repository.CreateProject(&project)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(200, createdProject)
		}
	})
}

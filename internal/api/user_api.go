package api

import (
	"encoding/json"
	"net/http"

	"circuit.io/circuit/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/user", func(context *gin.Context) {
		user, err := service.CreateUserIfNecessary(context)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			jsonBytes, err := json.Marshal(user)
			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
			}

			var data map[string]interface{}
			error := json.Unmarshal([]byte(jsonBytes), &data)
			if error != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			delete(data, "Password")
			delete(data, "HashPassword")
			context.Redirect(http.StatusMovedPermanently, "/f/login")
		}
	})

	router.POST("/login", func(context *gin.Context) {
		code, err := service.LoginHandler(context)
		if err != nil {
			context.JSON(code, gin.H{
				"error": err.Error(),
			})
		} else {
			// context.Redirect(http.StatusMovedPermanently, "/f/dashboard")
			context.JSON(code, gin.H{"message": "Login successful"})
		}
	})

	router.GET("/user/projects", service.AuthMiddleware(), func(context *gin.Context) {
		code, projects, err := service.GetUserProjects(context)
		if err != nil {
			context.JSON(code, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(code, projects)
		}
	})
}

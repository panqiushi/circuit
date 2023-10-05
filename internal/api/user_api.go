package api

import (
	"encoding/json"

	"code.circuit.io/circuit/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
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
			context.JSON(200, ResponseBuilder(data, 200, "success"))
		}
	})

	router.POST("/login", func(context *gin.Context) {
		service.LoginHandler(context)
	})
}

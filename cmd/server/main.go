package main

import (
	"code.circuit.io/circuit/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/user", func(context *gin.Context) {
		user, err := service.CreateUserIfNecessary(context)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(200, gin.H{
				"user": user,
			})
		}
	})

	router.Run(":8080")
}

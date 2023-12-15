package api

import (
	"net/http"

	"circuit.io/circuit/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterVersionReleaseRoutes(router *gin.Engine) {
	router.POST("/version_release", service.AuthMiddleware(), func(context *gin.Context) {
		code, versionRelease, err := service.CreateVersionRelease(context)
		if err != nil || code != http.StatusOK {
			context.JSON(code, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, versionRelease)
		}
	})
}

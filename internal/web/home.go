package web

import (
	"circuit.io/circuit/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RegisterHomePageRoutes(router *gin.Engine) {

	router.GET("/", service.AuthMiddleware(), func(context *gin.Context) {
		// localizer := context.MustGet("localizer").(*i18n.Localizer)
		// context.HTML(http.StatusOK, "home.html", getData(*localizer))
		context.Redirect(http.StatusMovedPermanently, "/dash")
	})

	router.GET("/dash", service.AuthMiddleware(), func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "home.html", getData(*localizer))
	})
}

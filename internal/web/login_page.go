package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterLoginPageRoutes(router *gin.Engine) {

	data := struct {
		Title string
	}{Title: "登录"}

	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", data)
	})
}

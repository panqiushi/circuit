package router

import (
	"code.circuit.io/circuit/internal/api"
	"code.circuit.io/circuit/internal/web"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()

	Router.LoadHTMLGlob("templates/pages/*")

	api.RegisterUserRoutes(Router)
	web.RegisterLoginPageRoutes(Router)
}

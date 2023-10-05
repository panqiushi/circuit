package router

import (
	"code.circuit.io/circuit/internal/api"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()

	api.RegisterUserRoutes(Router)
}

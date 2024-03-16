package router

import (
	"encoding/json"

	"circuit.io/circuit/internal/api"
	consts "circuit.io/circuit/internal/const"
	"circuit.io/circuit/internal/web"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()

	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(consts.MUST_LOAD_I18N_FILE)

	Router.Use(func(c *gin.Context) {
		acceptLanguage := c.GetHeader("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, acceptLanguage)
		c.Set("localizer", localizer)
		c.Next()
	})

	// Router.LoadHTMLGlob("templates/**/*")
	// Router.Static("/assets", "./assets")

	apiGroup := Router.Group("/a")
	apiGroup.GET("", func(context *gin.Context) {
		context.String(200, "hello circuit!")
	})
	api.RegisterUserRoutes(apiGroup)
	api.RegisterProjectRouters(apiGroup)
	api.RegisterVersionReleaseRoutes(apiGroup)
	web.RegisterWebPageRoutes(Router)

	Router.StaticFS("/f", gin.Dir("./frontend/.output/public", true))
}

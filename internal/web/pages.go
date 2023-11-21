package web

import (
	"circuit.io/circuit/internal/repository"
	"circuit.io/circuit/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RegisterWebPageRoutes(router *gin.Engine) {

	router.GET("/login", func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "login.html", getData(*localizer))
	})

	router.GET("/signup", func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "signup.html", getData(*localizer))
	})

	router.GET("/", service.AuthMiddleware(), func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/dash")
	})

	router.GET("/dash", service.AuthMiddleware(), func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "home.html", getData(*localizer))
	})

	router.GET("/projects", service.AuthMiddleware(), func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		data := getData(*localizer)
		projects, err := repository.FindAllProject()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		data["Projects"] = projects
		context.HTML(http.StatusOK, "projects.html", data)
	})
}

func getData(localizer i18n.Localizer) gin.H {
	return gin.H{
		"Theme":             "dark",
		"SignIn":            localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_in"}),
		"SignUp":            localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_up"}),
		"Welcome":           localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "welcome"}),
		"Or":                localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "or"}),
		"Email":             localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "email"}),
		"Password":          localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"}),
		"AlreadHaveAccount": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "already_have_account"}),
		"Home":              localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "home"}),
		"Dashboard":         localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "dashboard"}),
		"VersionRelease":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "version_release"}),
		"Username":          localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username"}),
		"ProjectName":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "project_name"}),
		"ProjectDesc":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "project_description"}),
		"CreateProject":     localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "create_project"}),
		"Project":           localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "project"}),
		"ProjectList":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "project_list"}),
		"Confirm":           localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "confirm"}),
	}

}

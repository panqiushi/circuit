package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func RegisterLoginPageRoutes(router *gin.Engine) {

	router.GET("/login", func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "login.html", gin.H{
			"Title":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_in"}),
			"SignIn":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_in"}),
			"SignUp":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_up"}),
			"Welcome":  localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "welcome"}),
			"Or":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "or"}),
			"Email":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "email"}),
			"Password": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"}),
		})
	})

	router.GET("/signup", func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		context.HTML(http.StatusOK, "signup.html", gin.H{
			"Title":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_up"}),
			"SignIn":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_in"}),
			"SignUp":   localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sign_up"}),
			"Welcome":  localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "welcome"}),
			"Or":       localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "or"}),
			"Email":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "email"}),
			"Password": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"}),
		})
	})

}

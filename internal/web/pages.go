package web

import (
	"encoding/json"
	"log"
	"os"

	consts "circuit.io/circuit/internal/const"
	"circuit.io/circuit/internal/repository"
	"circuit.io/circuit/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var i18nMessage gin.H = nil

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

	router.GET("/version_release/create", service.AuthMiddleware(), func(context *gin.Context) {
		localizer := context.MustGet("localizer").(*i18n.Localizer)
		data := getData(*localizer)
		projects, err := repository.FindAllProject()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		data["Projects"] = projects
		context.HTML(http.StatusOK, "version_release_create.html", data)
	})
}

func getData(localizer i18n.Localizer) gin.H {
	if i18nMessage != nil {
		return i18nMessage
	}

	data, err := os.ReadFile(consts.MUST_LOAD_I18N_FILE)
	if err != nil {
		log.Fatalf("Failed to load file: %s", err)
	}

	var jsonData map[string]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err)
	}

	hData := gin.H{}
	for key := range jsonData {
		hData[key] = localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
	}

	i18nMessage = hData
	return i18nMessage
}

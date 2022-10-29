package controller

import (
	"mnc-bank-api/manager"
	"mnc-bank-api/middleware"
	"mnc-bank-api/utils/jsonrw"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func NewController(usecaseManager manager.UsecaseManager, router *gin.Engine) *Controller {
	controller := Controller{
		ucMan:  usecaseManager,
		router: router,
	}

	router.GET("/", func(ctx *gin.Context) {

		jsonrw.JsonUpdateList("activity_log", nil)
		ctx.String(http.StatusOK, "hello world")
	})

	router.GET("/session", middleware.IsLogin(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello logged in person")
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.String(http.StatusPermanentRedirect, "you are already logged in")
	})

	return &controller
}

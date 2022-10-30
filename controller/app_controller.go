package controller

import (
	"mnc-bank-api/manager"
	response "mnc-bank-api/utils/common_response"
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
		ctx.String(http.StatusOK, "hello world")
	})

	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusPermanentRedirect, gin.H{
			"message": "you are already logged in",
		})
	})

	router.GET("/activity_log", func(ctx *gin.Context) {
		list, _ := jsonrw.JsonReadData("activity_log")
		response.JsonDataResponse(ctx, list)
	})

	return &controller
}

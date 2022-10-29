package controller

import (
	"log"
	"mnc-bank-api/config"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/authenticator"
	"mnc-bank-api/utils/jsonrw"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	usecase usecase.CustomerUsecase
	router  *gin.Engine
}

func (lc *LoginController) Login(ctx *gin.Context) {
	customerId, _ := ctx.Cookie("session")
	if customerId != "" {
		ctx.Redirect(300, "/")
		return
	}

	var credential model.Credential

	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	customer, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
	if err != nil {
		utils.JsonErrorNotFound(ctx, err, "customer not found")
		return
	}

	ctx.SetCookie("session", customer.Id, 3600, "/", "localhost", true, true)

	err = jsonrw.JsonWriteData("activity_log", model.Activity{
		Id:         utils.GenerateId(),
		CustomerId: customer.Id,
		Activity:   "login",
		Time:       time.Now(),
	})
	if err != nil {
		log.Println("unable to log login:", err)
	}

	utils.JsonSuccessMessage(ctx, "login success")
}

func (lc *LoginController) Logout(ctx *gin.Context) {
	customerId, err := ctx.Cookie("session")
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "error getting session")
		return
	}

	if customerId == "" {
		utils.JsonBadRequestMessage(ctx, "Not logged in")
		return
	}

	err = jsonrw.JsonWriteData("activity_log", model.Activity{
		Id:         utils.GenerateId(),
		CustomerId: customerId,
		Activity:   "logout",
		Time:       time.Now(),
	})
	if err != nil {
		log.Println("unable to log logout:", err)
	}

	ctx.SetCookie("session", "", -1, "/", "localhost", true, true)
	utils.JsonSuccessMessage(ctx, "logout success")
}

func (lc *LoginController) LoginTest(ctx *gin.Context) {
	var credential model.Credential
	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)

	err := ctx.ShouldBindJSON(&credential)

	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	customer, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "invalid credentials")
		return
	}

	_, err = accessToken.GenerateAccessToken(&customer)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot generate token")
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func NewLoginController(usecase usecase.CustomerUsecase, router *gin.Engine) *LoginController {
	controller := LoginController{
		usecase: usecase,
		router:  router,
	}

	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)
	router.POST("/test/login", controller.LoginTest)

	return &controller
}

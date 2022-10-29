package controller

// import (
// 	"mnc-bank-api/config"
// 	"mnc-bank-api/model"
// 	"mnc-bank-api/usecase"
// 	"mnc-bank-api/utils"
// 	"mnc-bank-api/utils/authenticator"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type LoginController struct {
// 	usecase usecase.CustomerUsecase
// 	router  *gin.Engine
// }

// func (lc *LoginController) Login(ctx *gin.Context) {
// 	var credential model.Credential
// 	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)

// 	err := ctx.ShouldBindJSON(&credential)
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
// 		return
// 	}

// 	customer, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "invalid credentials")
// 		return
// 	}

// 	token, err := accessToken.GenerateAccessToken(&customer)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "cannot generate token")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "you are logged in",
// 		"token":   token,
// 	})
// }

// func (lc *LoginController) LoginTest(ctx *gin.Context) {
// 	var credential model.Credential
// 	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)

// 	err := ctx.ShouldBindJSON(&credential)

// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
// 		return
// 	}

// 	customer, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "invalid credentials")
// 		return
// 	}

// 	_, err = accessToken.GenerateAccessToken(&customer)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "cannot generate token")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, customer)
// }

// func NewLoginController(usecase usecase.CustomerUsecase, router *gin.Engine) *LoginController {
// 	controller := LoginController{
// 		usecase: usecase,
// 		router:  router,
// 	}

// 	router.POST("/login", controller.Login)
// 	router.POST("/test/login", controller.LoginTest)

// 	return &controller
// }

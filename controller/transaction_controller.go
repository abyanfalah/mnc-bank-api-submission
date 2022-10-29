package controller

import (
	"mnc-bank-api/config"
	"mnc-bank-api/middleware"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	usecase         usecase.TransactionUsecase
	customerUsecase usecase.CustomerUsecase
	router          *gin.Engine
}

func (c *TransactionController) ListTransaction(ctx *gin.Context) {
	list, err := c.usecase.GetAll()
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot get transaction list")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *TransactionController) GetById(ctx *gin.Context) {
	transaction, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorNotFound(ctx, err, "cannot get transaction")
		return
	}

	utils.JsonDataResponse(ctx, transaction)
}

func (c *TransactionController) CreateNewTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	transaction.Id = utils.GenerateId()

	newTransaction, err := c.usecase.Insert(&transaction)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, newTransaction, "transaction created")
}

func NewTransactionController(usecase usecase.TransactionUsecase, CustomerUsecase usecase.CustomerUsecase, router *gin.Engine) *TransactionController {
	controller := TransactionController{
		usecase:         usecase,
		customerUsecase: CustomerUsecase,
		router:          router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	protectedRoute := router.Group("/transaction", authMiddleware.RequireToken())
	protectedRoute.GET("", controller.ListTransaction)
	protectedRoute.GET("/:id", controller.GetById)
	protectedRoute.POST("", controller.CreateNewTransaction)

	return &controller
}

package controller

import (
	"mnc-bank-api/middleware"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	"mnc-bank-api/utils"

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

	if len(list) == 0 {
		utils.JsonSuccessMessage(ctx, "list empty")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

// func (c *TransactionController) GetById(ctx *gin.Context) {
// 	transaction, err := c.usecase.GetById(ctx.Param("id"))
// 	if err != nil {
// 		utils.JsonErrorNotFound(ctx, err, "cannot get transaction")
// 		return
// 	}

// 	utils.JsonDataResponse(ctx, transaction)
// }

func (c *TransactionController) CreateNewTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	customerId, _ := ctx.Cookie("session")
	customer, _ := c.customerUsecase.GetById(customerId)

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	if transaction.Amount > customer.Balance {
		utils.JsonBadRequestMessage(ctx, "payment amount exceed your balance")
		return
	}

	// create transaction
	transaction.SenderId = customer.Id
	newTransaction, err := c.usecase.Insert(&transaction)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	// update customer balance

	utils.JsonDataMessageResponse(ctx, newTransaction, "transaction created")
}

func NewTransactionController(usecase usecase.TransactionUsecase, CustomerUsecase usecase.CustomerUsecase, router *gin.Engine) *TransactionController {
	controller := TransactionController{
		usecase:         usecase,
		customerUsecase: CustomerUsecase,
		router:          router,
	}

	router.GET("/transaction", middleware.IsLogin(), controller.ListTransaction)
	// sessionMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	// protectedRoute := router.Group("/transaction", authMiddleware.RequireToken())
	// protectedRoute.GET("", controller.ListTransaction)
	// protectedRoute.GET("/:id", controller.GetById)
	// protectedRoute.POST("", controller.CreateNewTransaction)

	return &controller
}

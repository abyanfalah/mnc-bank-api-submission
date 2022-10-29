package controller

import (
	"mnc-bank-api/middleware"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	response "mnc-bank-api/utils/common_response"

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
		response.JsonErrorInternalServerError(ctx, err, "cannot get transaction list")
		return
	}

	if len(list) == 0 {
		response.JsonSuccessMessage(ctx, "list empty")
		return
	}

	response.JsonDataResponse(ctx, list)
}

// func (c *TransactionController) GetById(ctx *gin.Context) {
// 	transaction, err := c.usecase.GetById(ctx.Param("id"))
// 	if err != nil {
// 		response.JsonErrorNotFound(ctx, err, "cannot get transaction")
// 		return
// 	}

// 	response.JsonDataResponse(ctx, transaction)
// }

func (c *TransactionController) CreateNewTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	customerId, _ := ctx.Cookie("session")
	customer, _ := c.customerUsecase.GetById(customerId)

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		response.JsonErrorBadRequestMessage(ctx, err, "cant bind struct")
		return
	}

	if transaction.ReceiverId == customerId {
		response.JsonBadRequestMessage(ctx, "invalid receiver. nice joke")
		return
	}

	_, err = c.customerUsecase.GetById(transaction.ReceiverId)
	if err != nil {
		response.JsonErrorBadRequest(ctx, err)
		return
	}

	if transaction.Amount < 0 {
		response.JsonBadRequestMessage(ctx, "invalid payment amount")
		return
	}

	if transaction.Amount > customer.Balance {
		response.JsonBadRequestMessage(ctx, "payment amount exceed your balance")
		return
	}

	// update both customers balance
	err = c.customerUsecase.UpdateBothBalance(transaction.Amount, customerId, transaction.ReceiverId)
	if err != nil {
		response.JsonErrorInternalServerError(ctx, err, "transaction failed, cannot update both balance")
		return
	}

	// create transaction
	transaction.SenderId = customer.Id
	newTransaction, err := c.usecase.Insert(&transaction)
	if err != nil {
		response.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	response.JsonDataMessageResponse(ctx, newTransaction, "transaction created")
}

func NewTransactionController(usecase usecase.TransactionUsecase, CustomerUsecase usecase.CustomerUsecase, router *gin.Engine) *TransactionController {
	controller := TransactionController{
		usecase:         usecase,
		customerUsecase: CustomerUsecase,
		router:          router,
	}

	protectedRoute := router.Group("/transaction", middleware.IsLogin())
	protectedRoute.GET("", controller.ListTransaction)
	// protectedRoute.GET("/:id", controller.GetById)
	protectedRoute.POST("", controller.CreateNewTransaction)

	return &controller
}

package controller

import (
	"log"
	"mnc-bank-api/model"
	"mnc-bank-api/usecase"
	response "mnc-bank-api/utils/common_response"
	"mnc-bank-api/utils/jsonrw"
	"net/http"
	"time"

	"mnc-bank-api/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	usecase usecase.CustomerUsecase
	router  *gin.Engine
}

func (c *CustomerController) ListCustomer(ctx *gin.Context) {
	list, err := c.usecase.GetAll()
	if err != nil {
		response.JsonErrorInternalServerError(ctx, err, "cannot get customer list")
		return
	}

	response.JsonDataResponse(ctx, list)
}

func (c *CustomerController) GetById(ctx *gin.Context) {
	customer, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		response.JsonErrorNotFound(ctx, err)
		return
	}

	response.JsonDataResponse(ctx, customer)
}

func (c *CustomerController) CreateNewCustomer(ctx *gin.Context) {
	custId, _ := ctx.Cookie("session")
	if custId == "" {
		response.JsonBadRequestMessage(ctx, "you are registered!")
		return
	}

	var customer model.Customer

	err := ctx.ShouldBind(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  customer,
		})
		return
	}

	find := c.usecase.GetByUsername(customer.Username)
	if find.Id != "" {
		response.JsonBadRequestMessage(ctx, "username taken!")
		return
	}

	customer, err = c.usecase.Insert(&customer)
	if err != nil {
		response.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	err = jsonrw.JsonWriteData("activity_log", model.Activity{
		Id:         utils.GenerateId(),
		CustomerId: customer.Id,
		Activity:   "register",
		Time:       time.Now(),
	})
	if err != nil {
		log.Println("unable to log registration:", err)
	}

	response.JsonDataMessageResponse(ctx, customer, "customer created")
}

func NewCustomerController(usecase usecase.CustomerUsecase, router *gin.Engine) *CustomerController {
	controller := CustomerController{
		usecase: usecase,
		router:  router,
	}

	router.GET("/customer", controller.ListCustomer)
	router.GET("/customer/:id", controller.GetById)
	router.POST("/customer", controller.CreateNewCustomer)

	// protectedRoute := router.Group("/customer", authMiddleware.RequireToken())
	// protectedRoute.POST("/", controller.CreateNewCustomer)
	// protectedRoute.PUT("/:id", controller.UpdateCustomer)
	// protectedRoute.DELETE("/:id", controller.DeleteCustomer)

	return &controller
}

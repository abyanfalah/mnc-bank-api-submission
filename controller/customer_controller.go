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
	var customer model.Customer

	err := ctx.ShouldBind(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  customer,
		})
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
		// response.JsonErrorInternalServerError(ctx, err, "unable to log registration")
		// return
		log.Println("unable to log registration:", err)
	}

	response.JsonDataMessageResponse(ctx, customer, "customer created")
}

// func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
// 	var customer model.Customer

// 	err := ctx.ShouldBindJSON(&customer)
// 	if err != nil {
// 		response.JsonErrorBadRequest(ctx, err, "cant bind struct")
// 		return
// 	}

// 	customer.Id = ctx.Param("id")
// 	updatedCustomer, err := c.usecase.Update(&customer)
// 	if err != nil {
// 		response.JsonErrorInternalServerError(ctx, err, "update failed")
// 		return
// 	}

// 	response.JsonDataResponse(ctx, updatedCustomer)
// }

// func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
// 	customer, err := c.usecase.GetById(ctx.Param("id"))
// 	if err != nil {
// 		response.JsonErrorNotFound(ctx, err, "customer not found")
// 		return
// 	}

// 	err = c.usecase.Delete(customer.Id)
// 	if err != nil {
// 		response.JsonErrorInternalServerError(ctx, err, "cannot delete customer")
// 		return
// 	}

// 	err = os.Remove("./images/customer/" + customer.Id + ".jpg")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	response.JsonSuccessMessage(ctx, "Customer deleted")
// }

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

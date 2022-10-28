package controller

import (
	"net/http"
	"warung-makan/config"
	"warung-makan/middleware"
	"warung-makan/model"
	"warung-makan/usecase"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	usecase usecase.CustomerUsecase
	router  *gin.Engine
}

func (c *CustomerController) ListCustomer(ctx *gin.Context) {
	if name := ctx.Query("name"); name != "" {
		customer, err := c.usecase.GetByName(ctx.Query("name"))

		if err != nil {
			utils.JsonErrorNotFound(ctx, err, "cannot get list")
			return
		}

		if len(customer) == 0 {
			ctx.String(http.StatusBadRequest, "no customer with name like "+name)
			return
		}

		utils.JsonDataResponse(ctx, customer)
		return
	}

	list, err := c.usecase.GetAll()
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot get customer list")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *CustomerController) GetById(ctx *gin.Context) {
	customer, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorNotFound(ctx, err, "cannot get customer")
		return
	}

	utils.JsonDataResponse(ctx, customer)
}

func (c *CustomerController) CreateNewCustomer(ctx *gin.Context) {
	var customer model.Customer
	c.router.MaxMultipartMemory = 8 << 20

	err := ctx.ShouldBind(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  customer,
		})
		return
	}

	imageFile, err := ctx.FormFile("image_file")
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant get image")
		return
	}

	id := utils.GenerateId()

	imagePath := "./images/customer/" + id + ".jpg"
	err = ctx.SaveUploadedFile(imageFile, imagePath)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot save image")
		return
	}

	customer.Id = id
	customer.Image = id + ".jpg"
	customer, err = c.usecase.Insert(&customer)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, customer, "customer created")
}

// func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
// 	var customer model.Customer

// 	err := ctx.ShouldBindJSON(&customer)
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
// 		return
// 	}

// 	customer.Id = ctx.Param("id")
// 	updatedCustomer, err := c.usecase.Update(&customer)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "update failed")
// 		return
// 	}

// 	utils.JsonDataResponse(ctx, updatedCustomer)
// }

// func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
// 	customer, err := c.usecase.GetById(ctx.Param("id"))
// 	if err != nil {
// 		utils.JsonErrorNotFound(ctx, err, "customer not found")
// 		return
// 	}

// 	err = c.usecase.Delete(customer.Id)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "cannot delete customer")
// 		return
// 	}

// 	err = os.Remove("./images/customer/" + customer.Id + ".jpg")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	utils.JsonSuccessMessage(ctx, "Customer deleted")
// }

func NewCustomerController(usecase usecase.CustomerUsecase, router *gin.Engine) *CustomerController {
	controller := CustomerController{
		usecase: usecase,
		router:  router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	router.GET("/customer", controller.ListCustomer)
	router.GET("/customer/:id", controller.GetById)

	protectedRoute := router.Group("/customer", authMiddleware.RequireToken())
	protectedRoute.POST("/", controller.CreateNewCustomer)
	// protectedRoute.PUT("/:id", controller.UpdateCustomer)
	// protectedRoute.DELETE("/:id", controller.DeleteCustomer)

	return &controller
}

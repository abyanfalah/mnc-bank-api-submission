package server

import (
	"mnc-bank-api/config"
	"mnc-bank-api/controller"
	"mnc-bank-api/manager"
	"strconv"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

type appServer struct {
	ucMan  manager.UsecaseManager
	engine *gin.Engine
	config config.Config
}

func NewAppServer() *appServer {
	config := config.NewConfig()
	// infraMan := manager.NewInfraManager()
	repoMan := manager.NewRepoManager()

	return &appServer{
		ucMan:  manager.NewUsecaseManager(repoMan),
		engine: gin.Default(),
		config: config,
	}
}

func (a *appServer) initHandlers() {
	controller.NewController(a.ucMan, a.engine)

	controller.NewCustomerController(a.ucMan.CustomerUsecase(), a.engine)
	controller.NewTransactionController(a.ucMan.TransactionUsecase(), a.ucMan.CustomerUsecase(), a.engine)
	controller.NewLoginController(a.ucMan.CustomerUsecase(), a.engine)
}

func (a *appServer) Run() {
	a.initHandlers()
	apiPort := a.config.ApiConfig.Port
	if apiPort == "" {
		apiPort = "8000"
	}

	for {
		err := a.engine.Run(":" + apiPort)
		if err != nil {
			apiPortInt, _ := strconv.Atoi(apiPort)
			apiPort = strconv.Itoa(apiPortInt + 1)
		} else {
			break
		}
	}

}

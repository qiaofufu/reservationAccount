package main

import (
	"ReservationAccount/config"
	_ "ReservationAccount/docs"
	"ReservationAccount/initialize"
	"ReservationAccount/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @title Swagger API
// @version 1.1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	initialization()
	start()
}

func initialization() {
	config.LoadConfig()
	initialize.InitDatabase()
	initialize.InitDatabaseTables()
	initialize.InitRedis()
	initialize.InitBase()
	go initialize.InitCronJob()
	utils.InitWechatPay()
	utils.InitCredential()
}

func start() {
	engine := gin.Default()
	port := ":" + viper.GetString("server.port")
	pemPath := viper.GetString("server.pemPath")
	keyPath := viper.GetString("server.keyPath")

	// swagger page
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	initialize.InitRouter(engine)

	err := engine.RunTLS(port, pemPath, keyPath)

	if err != nil {
		panic(fmt.Errorf("start server error [%w]\n", err))
	}
}

package initialize

import (
	"ReservationAccount/middleware"
	"ReservationAccount/router"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("", func(context *gin.Context) {
		context.String(200, "success")
	})
	Router := engine.Group("api/v1")

	// Tls
	Router.Use(middleware.Tls())

	// CORS 跨域
	Router.Use(middleware.Cors())

	// Log 日志
	// Router.Use(middleware.Log())

	systemRouter := router.RouterGroupAPP.SystemRouter

	// 路由初始化
	systemRouter.ReservationRouter.MountRouter(Router)
	systemRouter.WechatRouter.MountRouter(Router)
	systemRouter.BaseRouter.MountRouter(Router)
}

package system

import (
	v1 "ReservationAccount/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (receiver BaseRouter) MountRouter(group *gin.RouterGroup) {
	r := group.Group("base")

	r.POST("captchaPhone", v1.APIGroup.CaptchaAPI.CaptchaPhone)
}

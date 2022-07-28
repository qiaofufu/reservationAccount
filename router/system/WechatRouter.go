package system

import (
	v1 "ReservationAccount/api/v1"
	"github.com/gin-gonic/gin"
)

type WechatRouter struct {
}

func (receiver WechatRouter) MountRouter(group *gin.RouterGroup) {
	wechatAPI := v1.APIGroup.WechatAPI
	r := group.Group("wechat")
	{
		r.POST("getOpenID", wechatAPI.GetOpenID)
	}
}

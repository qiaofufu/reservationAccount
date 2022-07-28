package v1

import (
	"ReservationAccount/models/request"
	"ReservationAccount/models/response"
	"github.com/gin-gonic/gin"
)

type WechatAPI struct {
}

// GetOpenID
// @Summary 获取微信openid
// @Tags 微信
// @Accept json
// @Produce json
// @Param data body request.GetOpenID true "参数"
// @Success 200 {object} response.GetOpenID "查询成功"
// @Router /wechat/getOpenID [post]
func (w WechatAPI) GetOpenID(ctx *gin.Context) {
	var reqDTO request.GetOpenID
	if err := ctx.ShouldBindJSON(&reqDTO); err != nil {
		response.BindJSONError(err, ctx)
		return
	}
	openID, err := WechatService.GetOpenID(reqDTO.Code)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	var resDTO response.GetOpenID
	resDTO.OpenID = openID

	response.Success(resDTO, "获取成功", ctx)
}

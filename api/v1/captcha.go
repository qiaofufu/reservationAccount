package v1

import (
	"ReservationAccount/models/request"
	"ReservationAccount/models/response"
	"ReservationAccount/utils"
	"github.com/gin-gonic/gin"
)

type CaptchaAPI struct{}

// CaptchaPhone
// @Summary 获取手机验证码
// @Tags 验证
// @Accept json
// @Product json
// @Param data body request.ShotMessageCode true "传入参数"
// @Success 200 {object} response.DTO{} "获取成功"
// @Router /base/captchaPhone [post]
func (receiver CaptchaAPI) CaptchaPhone(ctx *gin.Context) {
	var phoneCaptcha request.ShotMessageCode
	if err := ctx.ShouldBindJSON(&phoneCaptcha); err != nil {
		response.BindJSONError(err, ctx)
		return
	}
	phoneCaptcha.Number = utils.DecryptByAes(phoneCaptcha.Number)
	err := CaptchaService.CaptchaPhone(phoneCaptcha.Number)
	if err != nil {
		response.FailWithMessage("获取失败!"+err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("获取成功", ctx)
}

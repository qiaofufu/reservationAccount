package system

import (
	"ReservationAccount/global"
	"ReservationAccount/utils"
	"context"
	"errors"
	"time"
)

type CaptchaService struct{}

// CaptchaPhone
// 生成手机验证码
func (receiver CaptchaService) CaptchaPhone(phone string) error {
	code := utils.RandomCode(6)
	err := utils.SendSMS(phone, code)
	ctx := context.Background()
	if err != nil {
		return errors.New("发送手机验证码失败！ " + err.Error())
	}
	global.Redis.Set(ctx, "phone:"+phone, code, time.Second*300)
	return nil
}

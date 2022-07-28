package models

import (
	"ReservationAccount/global"
)

const (
	UnPaid = 0 // 未支付
	Paid   = 1 // 支付成功
	Error  = 2 // 错误
)

type ReservationRecord struct {
	global.Model
	RealName           string `json:"real_name"`               // 真实名字
	IdentifyCard       string `json:"identify_card"`           // 身份证
	ContactPhone       string `json:"contact_phone"`           // 联系电话
	ReservationPhoneID uint   `json:"reservation_phone_id"`    // 预定电话id
	OpenID             string `json:"open_id"`                 // 微信openid
	School             string `json:"school"`                  // 预约学校
	SalesmanID         uint   `json:"salesman_id"`             // 推销员id
	Status             uint   `json:"status" gorm:"default:0"` // 订单状态 0-未支付 1-支付成功 2-错误
}

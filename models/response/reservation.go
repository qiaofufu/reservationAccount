package response

import (
	"ReservationAccount/models"
	"time"
)

type GetsRandomPhoneNumbers struct {
	Numbers []models.PhoneNumber `json:"numbers"`
}

type SearchPhone struct {
	Number []models.PhoneNumber `json:"number"`
}

type GetReservationRecord struct {
	CreatedAt        time.Time `json:"created_at"`              // 订单时间
	RealName         string    `json:"real_name"`               // 真实名字
	IdentifyCard     string    `json:"identify_card"`           // 身份证 加密
	ContactPhone     string    `json:"contact_phone"`           // 联系电话 加密
	ReservationPhone string    `json:"reservation_phone"`       // 预定电话 加密
	School           string    `json:"school"`                  // 预约学校
	SalesmanID       uint      `json:"salesman_id"`             // 推销员id
	SalesmanName     string    `json:"salesman_name"`           // 推销员名字
	SalesmanPhone    string    `json:"salesman_phone"`          // 推销员电话 加密
	Status           uint      `json:"status" gorm:"default:0"` // 订单状态 0-未支付 1-支付成功 2-错误
}

type ReservationPhone struct {
	PrepayId  string `json:"prepay_id"`
	Appid     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

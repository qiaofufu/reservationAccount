package models

import "ReservationAccount/global"

type Salesman struct {
	global.Model
	RealName     string // 推销员真实姓名
	ContactPhone string // 联系电话
}

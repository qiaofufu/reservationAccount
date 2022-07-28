package models

import (
	"ReservationAccount/global"
)

const (
	Sold   = 1
	UnSold = 0
)

type PhoneNumber struct {
	global.Model
	Phone   string `json:"phone"`                   // 手机号
	Status  int    `json:"status" gorm:"default:0"` // 手机号状态 0-未售 1-已售
	Version int    `json:"-" gorm:"default:1"`      // 乐观锁 version
}

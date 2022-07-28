package v1

import "ReservationAccount/service"

type Group struct {
	ReservationAPI
	WechatAPI
	CaptchaAPI
}

var APIGroup Group

var (
	ReservationService = service.ServiceGroupAPP.SystemService.ReservationService
	WechatService      = service.ServiceGroupAPP.SystemService.WechatService
	CaptchaService     = service.ServiceGroupAPP.SystemService.CaptchaService
	SalesmanService    = service.ServiceGroupAPP.SystemService.SalesmanService
)

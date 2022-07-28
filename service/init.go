package service

import systemSer "ReservationAccount/service/system"

type ServiceGroup struct {
	systemSer.SystemService
}

var ServiceGroupAPP = new(ServiceGroup)

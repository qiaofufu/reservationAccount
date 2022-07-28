package initialize

import (
	"ReservationAccount/service"
	"github.com/jasonlvhit/gocron"
)

func InitCronJob() {
	gocron.Every(1).Minute().Do(service.ServiceGroupAPP.ReservationService.DelOvertimeOrder)
	<-gocron.Start()
}

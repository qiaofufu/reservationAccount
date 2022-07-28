package system

import (
	v1 "ReservationAccount/api/v1"
	"github.com/gin-gonic/gin"
)

type ReservationRouter struct {
}

func (receiver ReservationRouter) MountRouter(group *gin.RouterGroup) {
	reservationAPI := v1.APIGroup.ReservationAPI
	r := group.Group("reservation")
	{
		r.POST("reservationPhone", reservationAPI.ReservationPhone)
		r.POST("searchPhone", reservationAPI.SearchPhone)
		r.POST("getReservationRecord", reservationAPI.GetReservationRecord)
		r.GET("getRandomPhoneNumbers", reservationAPI.GetsRandomPhoneNumbers)
		r.POST("payCallback/v1", reservationAPI.PayCallback)
	}
}

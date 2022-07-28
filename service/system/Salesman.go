package system

import (
	"ReservationAccount/global"
	"ReservationAccount/models"
)

type SalesmanService struct {
}

func (s SalesmanService) getSalesmanRecordByID(id uint) (salesman models.Salesman) {
	global.DB.Where("id = ?", id).First(&salesman)
	return
}

func (s SalesmanService) GetSalesmanRealNameByID(id uint) string {
	record := s.getSalesmanRecordByID(id)
	return record.RealName
}

func (s SalesmanService) GetSalesmanContactPhoneByID(id uint) string {
	record := s.getSalesmanRecordByID(id)
	return record.ContactPhone
}

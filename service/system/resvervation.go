package system

import (
	"ReservationAccount/global"
	"ReservationAccount/models"
	"ReservationAccount/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

const prefix = "1854730"

type ReservationService struct {
}

// GetRandomPhoneNumbers
// 获取随机未售手机号
func (receiver ReservationService) GetRandomPhoneNumbers() (phones []models.PhoneNumber, err error) {
	var cnt int64
	global.DB.Model(&models.PhoneNumber{}).Where("status = ?", models.UnSold).Count(&cnt)
	off := rand.Intn(int(cnt + 1))
	err = global.DB.Model(&models.PhoneNumber{}).Where("status = ?", models.UnSold).Offset(off).Limit(10).Find(&phones).Error
	return
}

// SearchPhone
// 模糊查询后四位手机号
func (receiver ReservationService) SearchPhone(suffix string) (phone []models.PhoneNumber, err error) {
	if len(suffix) > 4 {
		err = errors.New("尾数超出4位限制")
		return
	}
	phones := prefix + "%" + suffix + "%"
	log.Println(phones)
	err = global.DB.Where("phone LIKE ?", phones).Find(&phone).Error
	return
}

// GetReservationRecord 获取预约记录
func (receiver ReservationService) GetReservationRecord(contactPhone string) (record models.ReservationRecord, err error) {
	err = global.DB.Where("contact_phone = ?", contactPhone).First(&record).Error
	return
}

// ExistRecord
// 检查记录是否存在
func (receiver ReservationService) ExistRecord(identifyCard string, openID string) bool {
	var cnt int64
	err := global.DB.Model(&models.ReservationRecord{}).Where("identify_card = ? or open_id = ? and status != ?", identifyCard, openID, models.Error).Count(&cnt).Error
	if err != nil {
		return true
	}
	if cnt >= 1 {
		return true
	} else {
		return false
	}
}

// Reservation
// 预约号码
func (receiver ReservationService) Reservation(realName string, identifyCard string, contactPhone string, school string, reservationPhoneID uint, salesmanID uint, openID string) (prePayRes jsapi.PrepayWithRequestPaymentResponse, err error) {

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 检验被预约的号码是否可以购买，并更新记录
		var phone models.PhoneNumber
		err1 := tx.Where("id = ?", reservationPhoneID).First(&phone).Error
		if err1 != nil {
			return errors.New("该账号不存")
		}
		sql := fmt.Sprintf("update phone_numbers "+
			"set status = 1, version = version + 1 "+
			"where id = %d and version = %d and status = %d", phone.ID, phone.Version, models.UnSold)
		result := tx.Exec(sql)
		if result.Error != nil {
			return errors.New("内部错误001")
		}
		if result.RowsAffected == 0 {
			return errors.New("该账号不存在或已被预购")
		}

		record := models.ReservationRecord{
			RealName:           realName,
			IdentifyCard:       identifyCard,
			ContactPhone:       contactPhone,
			ReservationPhoneID: reservationPhoneID,
			OpenID:             openID,
			SalesmanID:         salesmanID,
			School:             school,
			Status:             models.UnPaid,
		}

		// 生成订单缓存
		ctx := context.Background()
		data, _ := json.Marshal(record)
		global.Redis.Set(ctx, openID+identifyCard, data, time.Minute*15)
		// 生成订单记录
		err1 = tx.Create(&record).Error
		if err1 != nil {
			global.Redis.Del(ctx, openID+identifyCard)
			return errors.New("生成内部订单错误")
		}
		// 生成微信支付订单
		log.Println(record)
		prepayData, err1 := utils.PrePay("手机号预约", fmt.Sprintf("%d", record.ID), time.Now().Add(time.Minute*10), openID)
		if err1 != nil {
			global.Redis.Del(ctx, openID+identifyCard)
			return err1
		}
		prePayRes = prepayData
		return nil
	})

	return
}

// DelOvertimeOrder
// 删除超时订单
func (receiver ReservationService) DelOvertimeOrder() {
	log.Println("del overtime order")
	global.DB.Transaction(func(tx *gorm.DB) error {
		var records []models.ReservationRecord
		err := tx.Where("status = ? &&  created_at < DATE_SUB(NOW(),INTERVAL 30 MINUTE)", models.UnPaid).Find(&records).Error
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for _, v := range records {
			sql := "update phone_numbers " +
				"set status = ? where id = ?"
			err = tx.Exec(sql, models.UnSold, v.ReservationPhoneID).Error
			if err != nil {
				return err
			}
			err = tx.Model(&v).Update("status", models.Error).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// CompleteOrder
// 完成订单
func (receiver ReservationService) CompleteOrder(transaction *payments.Transaction) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		sql := "update reservation_records set status = ? where id = ?"
		result := tx.Exec(sql, models.Sold, transaction.OutTradeNo)
		if result.RowsAffected == 0 {
			return errors.New(fmt.Sprintf("当前订单不存在, data[%s]", transaction))
		}
		return nil
	})
	return err
}

// GetPhoneNumberByID
// 根据id获取手机号
func (receiver ReservationService) GetPhoneNumberByID(id uint) string {
	var phone models.PhoneNumber
	global.DB.Where("id = ?", id).Find(&phone)
	return phone.Phone
}

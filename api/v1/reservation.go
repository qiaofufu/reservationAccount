package v1

import (
	"ReservationAccount/models/request"
	"ReservationAccount/models/response"
	"ReservationAccount/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"log"
	"net/http"
)

type ReservationAPI struct {
}

// ReservationPhone
// @Summary 预约手机号
// @Tags 预约前台
// @Accept json
// @Produce json
// @Param data body request.ReservationPhone true "参数"
// @Success 200 {object} response.ReservationPhone "查询成功"
// @Router /reservation/reservationPhone [post]
func (receiver ReservationAPI) ReservationPhone(ctx *gin.Context) {
	var reqDTO request.ReservationPhone
	if err := ctx.ShouldBindJSON(&reqDTO); err != nil {
		response.BindJSONError(err, ctx)
		return
	}

	// 判断是否有预约记录，限制为每人只能预约一个手机号
	if ReservationService.ExistRecord(reqDTO.IdentifyCard, reqDTO.OpenID) {
		response.FailWithMessage("该用户以参与预约，请勿多次参与", ctx)
		return
	}
	// 信息解密
	data := utils.DecryptByAes(reqDTO.IdentifyCard)
	log.Println(data)
	reqDTO.IdentifyCard = data
	data = utils.DecryptByAes(reqDTO.ContactPhone)
	log.Println(data)
	reqDTO.ContactPhone = data
	// 进行预约操作
	prepay, err := ReservationService.Reservation(reqDTO.RealName, reqDTO.IdentifyCard, reqDTO.ContactPhone, reqDTO.School, reqDTO.ReservationPhoneID, reqDTO.SalesmanID, reqDTO.OpenID)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resDTO := response.ReservationPhone{
		PrepayId:  *prepay.PrepayId,
		Appid:     *prepay.Appid,
		TimeStamp: *prepay.TimeStamp,
		NonceStr:  *prepay.NonceStr,
		Package:   *prepay.Package,
		SignType:  *prepay.SignType,
		PaySign:   *prepay.PaySign,
	}
	response.Success(resDTO, "预约成功", ctx)
}

// GetReservationRecord
// @Summary 获取预约记录
// @Tags 预约前台
// @Accept json
// @Produce json
// @Param data body request.GetReservationRecord true "参数"
// @Success 200 {object} response.GetReservationRecord "查询成功"
// @Router /reservation/getReservationRecord [post]
func (receiver ReservationAPI) GetReservationRecord(ctx *gin.Context) {
	var reqDTO request.GetReservationRecord
	if err := ctx.ShouldBindJSON(&reqDTO); err != nil {
		response.BindJSONError(err, ctx)
		return
	}

	data := utils.DecryptByAes(reqDTO.Phone)
	log.Println(data)
	reqDTO.Phone = string(data)

	if utils.VerifySMS(reqDTO.Phone, reqDTO.VerifyCode) == false {
		response.FailWithMessage("短信验证失败", ctx)
		return
	}

	record, err := ReservationService.GetReservationRecord(reqDTO.Phone)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	record.ContactPhone = utils.EncryptByAes(record.ContactPhone)
	record.IdentifyCard = utils.EncryptByAes(record.IdentifyCard)

	reservationPhone := ReservationService.GetPhoneNumberByID(record.ReservationPhoneID)
	reservationPhone = utils.EncryptByAes(reservationPhone)
	salesmanPhone := SalesmanService.GetSalesmanContactPhoneByID(record.SalesmanID)
	salesmanPhone = utils.EncryptByAes(salesmanPhone)
	resDTO := response.GetReservationRecord{
		CreatedAt:        record.CreatedAt,
		RealName:         record.RealName,
		IdentifyCard:     record.IdentifyCard,
		ContactPhone:     record.ContactPhone,
		ReservationPhone: reservationPhone,
		School:           record.School,
		SalesmanID:       record.SalesmanID,
		SalesmanName:     SalesmanService.GetSalesmanRealNameByID(record.SalesmanID),
		SalesmanPhone:    salesmanPhone,
		Status:           record.Status,
	}
	response.Success(resDTO, "查询成功", ctx)
}

// GetsRandomPhoneNumbers
// @Summary 获取随机手机号
// @Tags 预约前台
// @Accept json
// @Produce json
// @Success 200 {object} response.GetsRandomPhoneNumbers "获取成功"
// @Router /reservation/getRandomPhoneNumbers [get]
func (receiver ReservationAPI) GetsRandomPhoneNumbers(ctx *gin.Context) {
	phones, err := ReservationService.GetRandomPhoneNumbers()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Success(phones, "获取成功", ctx)
}

// SearchPhone
// @Summary 模糊查找手机号
// @Tags 预约前台
// @Accept json
// @Produce json
// @Param data body request.SearchPhone true "参数"
// @Success 200 {object} response.SearchPhone "获取成功"
// @Router /reservation/searchPhone [post]
func (receiver ReservationAPI) SearchPhone(ctx *gin.Context) {
	var reqDTO request.SearchPhone
	if err := ctx.ShouldBindJSON(&reqDTO); err != nil {
		response.BindJSONError(err, ctx)
		return
	}

	phone, err := ReservationService.SearchPhone(reqDTO.Suffix)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resDTO := response.SearchPhone{Number: phone}
	response.Success(resDTO, "查询成功", ctx)
}

// PayCallback
// 支付回调函数
func (receiver ReservationAPI) PayCallback(ctx *gin.Context) {

	// 解析resource
	transaction := new(payments.Transaction)
	notifyReq, err := utils.WechatHandler.ParseNotifyRequest(context.Background(), ctx.Request, transaction)
	if err != nil {
		log.Println("验签失败")
		log.Println(err)
		ctx.XML(http.StatusOK, gin.H{
			"return_code": "FAIL",
			"return_msg":  err.Error(),
		})
		return
	}

	if notifyReq.Summary != "支付成功" {
		log.Println("支付失败")
		ctx.XML(http.StatusOK, gin.H{
			"return_code": "SUCCESS",
			"return_msg":  "OK",
		})
		return
	}

	// 进行后续业务处理
	err = ReservationService.CompleteOrder(transaction)
	if err != nil {
		log.Println(err)
		ctx.XML(http.StatusOK, gin.H{
			"return_code": "FAIL",
			"return_msg":  err.Error(),
		})
		return
	}
	ctx.XML(http.StatusOK, gin.H{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
	})

}

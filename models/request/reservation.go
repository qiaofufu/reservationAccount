package request

type SearchPhone struct {
	Suffix string `json:"suffix"`
}

type GetReservationRecord struct {
	Phone      string `json:"phone"`       // 手机号 加密
	VerifyCode string `json:"verify_code"` // 验证码
}

type ReservationPhone struct {
	RealName           string `json:"real_name"`
	School             string `json:"school"`
	IdentifyCard       string `json:"identify_card"` // 加密
	ContactPhone       string `json:"contact_phone"` // 加密
	ReservationPhoneID uint   `json:"reservation_phone_id"`
	SalesmanID         uint   `json:"salesman_id"`
	OpenID             string `json:"open_id"`
}

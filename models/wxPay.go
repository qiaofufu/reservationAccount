package models

import "time"

type WxPayRecord struct {
	ID             string    // 通知的唯一id
	CreatedAt      time.Time // 通知创建时间
	EventType      string    // 通知类型 支付成功的通知类型为TRANSACTION.SUCCESS
	ResourceType   string    //通知数据类型 - 支付成功通知为encrypt-resource
	Algorithm      string    // 加密算法类型
	Ciphertext     string    // Base64编码后数据密文
	AssociatedData string    // 附加数据
	OriginalType   string    // 原始回调类型， transaction
	Nonce          string    // 加密使用的随机串
	Summary        string    // 回调摘要
}

type PayRecord struct {
}

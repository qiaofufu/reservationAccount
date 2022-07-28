package utils

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"time"
)

var WechatClient *core.Client
var WechatHandler *notify.Handler

func InitWechatPay() {
	// 从本地目录加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(viper.GetString("wechat.keyPath"))
	if err != nil {
		log.Fatalf("load merchant private key error : %w", err)
		return
	}
	ctx := context.Background()
	mchID := viper.GetString("wechat.mchID")
	mchCertificateSerialNumber := viper.GetString("wechat.mchCertificateSerialNumber")
	mchAPIv3Key := viper.GetString("wechat.mchAPIv3Key")
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err : %s\n", err)
		return
	}

	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	WechatClient = client
	WechatHandler = handler

	return
}

// PrePay
// 预支付
func PrePay(description string, orderID string, timeExpire time.Time, openID string) (prepay jsapi.PrepayWithRequestPaymentResponse, err error) {
	svc := jsapi.JsapiApiService{Client: WechatClient}
	ctx := context.Background()
	resp, _, err1 := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(viper.GetString("wechat.appID")),      // 由微信生成的应用ID，全局唯一。请求基础下单接口时请注意APPID的应用属性，例如公众号场景下，需使用应用属性为公众号的服务号APPID
			Mchid:       core.String(viper.GetString("wechat.mchID")),      // 直连商户的商户号，由微信支付生成并下发。
			Description: core.String(description),                          // 商品描述
			OutTradeNo:  core.String(orderID),                              // 平台内部订单id
			TimeExpire:  core.Time(timeExpire),                             // 订单失效时间
			Attach:      nil,                                               // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。
			NotifyUrl:   core.String(viper.GetString("wechat.notify_url")), // 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 公网域名必须为https，如果是走专线接入，使用专线NAT IP或者私有回调域名可使用http
			Amount: &jsapi.Amount{
				Total:    core.Int64(1),
				Currency: core.String("CNY"),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openID),
			},
		})
	if err1 != nil {
		log.Println(err1)
		err = errors.New("微信订单生成失败")
	} else {
		prepay = *resp
	}
	return
}

//

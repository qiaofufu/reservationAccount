package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

type WechatService struct {
}

// GetOpenID
// 获取 微信OpenID
func (receiver WechatService) GetOpenID(code string) (openid string, err error) {
	appID := viper.GetString("wechat.appID")
	wxSecret := viper.GetString("wechat.wxSecret")
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appID, wxSecret, code)
	resp, err1 := http.DefaultClient.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var wxMap map[string]interface{}
	err = json.Unmarshal(data, &wxMap)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if wxMap["openid"] == nil {
		err = errors.New("获取openid失败")
		return
	}
	log.Println(wxMap["openid"].(string))
	return wxMap["openid"].(string), nil
}

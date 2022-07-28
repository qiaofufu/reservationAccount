package initialize

import (
	"ReservationAccount/utils"
	"github.com/spf13/viper"
)

func InitBase() {
	utils.Key = []byte(viper.GetString("server.key"))
}

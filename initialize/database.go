package initialize

import (
	"ReservationAccount/global"
	"ReservationAccount/models"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	charset := viper.GetString("database.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, dbname, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("Database init error: " + err.Error())
	}
	global.DB = db
}

func InitDatabaseTables() {
	err := global.DB.Set("gorm:table_options", "AUTO_INCREMENT=100000").AutoMigrate(
		&models.PhoneNumber{},
		&models.Salesman{},
		&models.ReservationRecord{},
	)
	if err != nil {
		panic(fmt.Errorf("database auto migrate fail err: %w\n", err))
	}
}

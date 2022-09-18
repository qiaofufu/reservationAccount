package initialize

import (
	"ReservationAccount/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

func InitRedis() {

	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	if err != nil {
		log.Fatal(err)
	}
	global.Redis = client
}

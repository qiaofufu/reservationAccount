package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Logger *Log
	Redis  *redis.Client
)

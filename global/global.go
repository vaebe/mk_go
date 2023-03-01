package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"mk/config"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
	JWTInfo     *config.JWTConfig
	MysqlConfig *config.MysqlConfig
	RedisConfig *config.RedisConfig
)

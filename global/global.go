package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"mk/config"
)

var (
	Trans       ut.Translator
	DB          *gorm.DB
	RedisClient *redis.Client
	JWTConfig   *config.JWTConfig
	MysqlConfig *config.MysqlConfig
	RedisConfig *config.RedisConfig
	EmailConfig *config.EmailConfig
)

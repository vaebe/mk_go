package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mk/global"
)

// GetEnvInfo 获取设置的env变量, 变量设置完成需要重启ide
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// InitConfig 初始化config配置
func InitConfig() {
	//debug := GetEnvInfo("Mk_DEBUG")
	configFilePrefix := "config"

	// 配置文件路径
	configFileName := fmt.Sprintf("./%s-dev.yaml", configFilePrefix)

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// mysqlConfig - 全局变量
	if err := v.UnmarshalKey("mysqlConfig", &global.MysqlConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("MysqlConfig配置信息: %v", global.MysqlConfig)

	// redisConfig - 全局变量
	if err := v.UnmarshalKey("redisConfig", &global.RedisConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("UserSrvConfig配置信息: %v", global.RedisConfig)
}

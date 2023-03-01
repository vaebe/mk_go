package config

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

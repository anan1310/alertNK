package config

// Config 主配置
type Config struct {
	MySQL      mysql      `mapstructure:"mysql"`
	Redis      redis      `mapstructure:"redis"`
	Clickhouse clickhouse `mapstructure:"clickhouse"`
	Zap        zap        `mapstructure:"zap"`
	Server     server     `mapstructure:"server"`
	JWT        jWT        `mapstructure:"jwt"`
}

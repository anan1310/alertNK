package config

import "time"

// 服务端
type server struct {
	Env           string        `mapstructure:"env"` //当前是开发环境还是生成环境
	RunMode       string        `mapstructure:"runMode"`
	HttpPort      int           `mapstructure:"httpPort"`
	ReadTimeout   time.Duration `mapstructure:"readTimeout"`  //读
	WriteTimeout  time.Duration `mapstructure:"writeTimeout"` //写
	GroupWait     int           `mapstructure:"groupWait" json:"groupWait"`
	GroupInterval int           `mapstructure:"groupInterval" json:"groupInterval"`
	RecoverWait   int           `mapstructure:"recoverWait"  json:"recoverWait"`
}

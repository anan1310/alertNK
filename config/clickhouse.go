package config

import "fmt"

// Clickhouse连接信息
type clickhouse struct {
	UserName     string `mapstructure:"userName"`
	PassWord     string `mapstructure:"passWord"`
	Host         string `mapstructure:"host"` //主机地址
	DataBase     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

func (c *clickhouse) Dsn() string {
	tcpInfo := "tcp://%s?username=%s&password=%s&database=%s&read_timeout=10s&compress=true"
	return fmt.Sprintf(tcpInfo, c.Host, c.UserName, c.PassWord, c.DataBase)
}

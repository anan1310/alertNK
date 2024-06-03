package config

// Mysql 数据库配置
type mysql struct {
	DriverName   string `mapstructure:"driverName"`
	UserName     string `mapstructure:"username"`
	PassWord     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	DataBase     string `mapstructure:"database"`
	Config       string `mapstructure:"config"` // 高级配置
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

func (m *mysql) Dsn() string {
	return m.UserName + ":" + m.PassWord + "@tcp(" + m.Host + ")/" + m.DataBase + "?" + m.Config
}

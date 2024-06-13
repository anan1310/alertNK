package config

type mail struct {
	Host     string   `mapstructure:"host" json:"host" yaml:"host"`
	Pass     string   `mapstructure:"pass" json:"pass" yaml:"pass"`
	Port     int      `mapstructure:"port" json:"port" yaml:"port"`
	To       []string `mapstructure:"to" json:"to" yaml:"to"`
	SmtpUser string   `mapstructure:"smtpUser" json:"smtpUser" yaml:"smtpUser"`
}

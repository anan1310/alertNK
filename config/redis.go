package config

type redis struct {
	Host string `mapstructure:"host"  json:"host"`
	Port string `mapstructure:"port"  json:"port"`
	Pass string `mapstructure:"pass" json:"pass"`
}

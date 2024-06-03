package config

type zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`                      // 级别
	Format       string `mapstructure:"format" json:"format" yaml:"format"`                   // 输出
	ShowLine     bool   `mapstructure:"showLine" json:"showLine" yaml:"showLine"`             // 显示行
	EncodeLevel  string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"`    // 编码级
	Path         string `mapstructure:"path" json:"path" yaml:"path" `                        //输出路径
	LogInConsole bool   `mapstructure:"logInConsole" json:"logInConsole" yaml:"logInConsole"` // 输出控制台
	MaxAge       int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`                   //最长保存天数
	MaxBackups   int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups" `      //最多备份几个
	Maxsize      int    `mapstructure:"maxsize" json:"maxsize" yaml:"maxsize" `               // 日志文件大小
	Compress     bool   `mapstructure:"compress"  json:"compress" yaml:"compress"`            // 是否压缩文件，使用gzip
}

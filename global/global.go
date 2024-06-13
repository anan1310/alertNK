package global

import (
	"alarm_collector/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Layout = "2006-01-02 15:04:05"
	Config config.Config
	Viper  *viper.Viper
	Logger *zap.Logger
)

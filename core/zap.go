package core

import (
	"alarm_collector/global"
	"alarm_collector/pkg/utils/file_util"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func Zap() (logger *zap.Logger) {
	//判断日志路径是否存在
	_ = file_util.IsNotExistMkDir(global.Config.Zap.Path)
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.Config.Zap.Path), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.Config.Zap.Path), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.Config.Zap.Path), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.Config.Zap.Path), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",                   // 添加堆栈信息
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     CustomTimeEncoder,              //自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.FullCallerEncoder,      // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}
	switch {
	case global.Config.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Config.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	maxSize := global.Config.Zap.Maxsize
	maxAge := global.Config.Zap.MaxAge
	maxBackups := global.Config.Zap.MaxBackups
	compress := global.Config.Zap.Compress
	writer := getWriteSyncer(fileName, maxSize, maxBackups, maxAge, compress) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

/*
1、指定将日志写到哪里去：文件打开的是io.writer类型，需要通过zapcore.AddSync()函数转换为WriteSyncer
2、使用lumberjack进行分割
*/
func getWriteSyncer(filename string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 日志文件的位置
		MaxSize:    maxSize,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackup, // 保留旧文件的最大个数
		MaxAge:     maxAge,    // 保留旧文件的最大天数
		Compress:   compress,  // 是否压缩/归档旧文件
	}

	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

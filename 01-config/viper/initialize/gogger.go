package initialize

import (
	"github.com/tanshouke/silents/01-config/viper/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.Logger {
	// Create logs directory if it doesn't exist
	if global.Cfg.Zap.Director != "" {
		if err := os.MkdirAll(global.Cfg.Zap.Director, 0755); err != nil {
			panic(err)
		}
	}

	// Determine log level
	/*
		log  level config
	*/
	var level zapcore.Level
	switch global.Cfg.Zap.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// Configure encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  global.Cfg.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	/* 日志输出样式配置 */
	// Use color encoder for console in development
	if global.Cfg.Zap.Format == "console" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Create encoder
	var encoder zapcore.Encoder
	if global.Cfg.Zap.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Create core
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	// Build logger
	logger := zap.New(core)

	if global.Cfg.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

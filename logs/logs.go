package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	logger = buildLogger("debug")
}

func GetLogger() *zap.Logger {
	return logger
}

func buildLogger(level string) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(zapLevel)
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
	return logger
}

func SetLogLevel(level string) {
	logger = buildLogger(level)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

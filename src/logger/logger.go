package logger

import (
	"go.uber.org/zap"
)

type Log struct {
	logger *zap.Logger
}

func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger
}

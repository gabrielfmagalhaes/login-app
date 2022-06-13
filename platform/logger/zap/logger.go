package zap

import (
	"login-app/platform/logger"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() logger.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	return &Logger{log: sugar}
}

func (logger *Logger) Errorf(message string, args ...interface{}) {
	logger.log.Errorw(message, args)
}
func (logger *Logger) Fatalf(message string, args ...interface{}) {
	logger.log.Fatalw(message, args)
}
func (logger *Logger) Infof(message string, args ...interface{}) {
	logger.log.Infow(message, args)
}
func (logger *Logger) Warnf(message string, args ...interface{}) {
	logger.log.Warnw(message, args)
}
func (logger *Logger) Debugf(message string, args ...interface{}) {
	logger.log.Debugw(message, args)
}
func (logger *Logger) Printf(message string, args ...interface{}) {
	logger.log.Infow(message, args)
}

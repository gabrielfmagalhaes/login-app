package zap

import (
	"login-app/platform/logger"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() logger.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	return &Logger{logger: sugar}
}

func (logger *Logger) Errorf(message string, args ...interface{}) {
	logger.logger.Errorln(message, args)
}
func (logger *Logger) Fatalf(message string, args ...interface{}) {
	logger.logger.Fatalln(message, args)
}
func (logger *Logger) Infof(message string, args ...interface{}) {
	logger.logger.Infoln(message, args)
}
func (logger *Logger) Warnf(message string, args ...interface{}) {
	logger.logger.Warnln(message, args)
}
func (logger *Logger) Debugf(message string, args ...interface{}) {
	logger.logger.Debugln(message, args)
}

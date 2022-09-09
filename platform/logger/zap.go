package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	return &zapLogger{logger: sugar}
}

func (l *zapLogger) Errorf(message string, args ...interface{}) {
	l.logger.Errorln(message, args)
}
func (l *zapLogger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatalln(message, args)
}
func (l *zapLogger) Infof(message string, args ...interface{}) {
	l.logger.Infoln(message, args)
}
func (l *zapLogger) Warnf(message string, args ...interface{}) {
	l.logger.Warnln(message, args)
}
func (l *zapLogger) Debugf(message string, args ...interface{}) {
	l.logger.Debugln(message, args)
}

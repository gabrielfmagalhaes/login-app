package logger

type Logger interface {
	Errorf(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
	Infof(message string, args ...interface{})
	Warnf(message string, args ...interface{})
	Debugf(message string, args ...interface{})
}

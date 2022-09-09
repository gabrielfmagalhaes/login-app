package logger

type LoggerMock struct{}

func (l LoggerMock) Infof(_ string, _ ...interface{})  {}
func (l LoggerMock) Debugf(_ string, _ ...interface{}) {}
func (l LoggerMock) Warnf(_ string, _ ...interface{})  {}
func (l LoggerMock) Errorf(_ string, _ ...interface{}) {}
func (l LoggerMock) Fatalf(_ string, _ ...interface{}) {}

package logger_test

import (
	"github.com/SGchuyue/logger/logger"
	"testing"
)

// 测试logger日志包的功能
func TestInitLogger(t *testing.T) {
	logger.InitLogger("test.log", 2, 3, 5, true)
	for i := 0; i < 50000; i++ {
		logger.Error("error test")
		logger.Errorf("test errorf %s", "test")
		logger.Debug("debug test")
		logger.Debugf("test debugf %s", "test")
		logger.Info("info test")
		logger.Infof("test infof %s", "test")
		logger.Warn("warn test")
		logger.Warnf("test warnf %s", "test")
	}
}
